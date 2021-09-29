// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020-2021 Intel Corporation

package daemon

import (
	"context"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"time"

	"github.com/smart-edge-open/openshift-operator/common/pkg/utils"
	fec "github.com/smart-edge-open/openshift-operator/sriov-fec/api/v2"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ConfigurationConditionReason string

const (
	ConditionConfigured       string                       = "Configured"
	ConfigurationInProgress   ConfigurationConditionReason = "InProgress"
	ConfigurationFailed       ConfigurationConditionReason = "Failed"
	ConfigurationNotRequested ConfigurationConditionReason = "NotRequested"
	ConfigurationSucceeded    ConfigurationConditionReason = "Succeeded"
)

var (
	resyncPeriod          = time.Minute
	configPath            = "/sriov_config/config/accelerators.json"
	getSriovInventory     = GetSriovInventory
	supportedAccelerators utils.AcceleratorDiscoveryConfig
)

type NodeConfigReconciler struct {
	client.Client
	log              *logrus.Logger
	nodeName         string
	namespace        string
	drainAndExecute  Drainer
	nodeConfigurator *NodeConfigurator
}

type Drainer func(configurer func(ctx context.Context) bool, drain bool) error
type configurer func(ctx context.Context) bool

type configurationStatus struct {
	isRebootRequested  bool
	configurationError error
}

func NewNodeConfigReconciler(k8sClient client.Client, drainer Drainer, log *logrus.Logger, nodeName, ns string) (r *NodeConfigReconciler, err error) {

	if supportedAccelerators, err = utils.LoadDiscoveryConfig(configPath); err != nil {
		return nil, err
	}

	kc, err := createKernelController(log)
	if err != nil {
		return nil, err
	}

	nc := &NodeConfigurator{Log: log, kernelController: kc}

	return &NodeConfigReconciler{
		Client:           k8sClient,
		log:              log,
		nodeName:         nodeName,
		namespace:        ns,
		drainAndExecute:  drainer,
		nodeConfigurator: nc,
	}, nil
}

func (r *NodeConfigReconciler) Reconcile(_ context.Context, req ctrl.Request) (ctrl.Result, error) {

	nc, err := r.readSriovFecNodeConfig(req.NamespacedName)
	if err != nil {
		return requeueNowWithError(err)
	}

	detectedInventory, err := r.readExistingInventory()
	if err != nil {
		return requeueNowWithError(err)
	}

	if isConfigurationOfNonExistingInventoryRequested(nc.Spec.PhysicalFunctions, detectedInventory) {
		return requeueLaterOrNowIfError(r.updateStatus(nc, metav1.ConditionFalse, ConfigurationFailed, "requested configuration reffers to not existing accelerator"))
	}

	if !isReconcileRequired(nc) {
		r.log.Info("Nothing to do")
		return requeueLater()
	}

	if err := r.updateStatus(nc, metav1.ConditionFalse, ConfigurationInProgress, "Configuration started"); err != nil {
		return requeueNowWithError(err)
	}

	if rebootRequested, err := r.configureNode(nc); err != nil {
		r.log.WithError(err).Error("error occurred during configuring node")
		return requeueNowWithError(r.updateStatus(nc, metav1.ConditionFalse, ConfigurationFailed, err.Error()))
	} else if rebootRequested {
		r.log.Info("status update skipped - CR will be handled again after node reboot")
		return requeueLater()
	} else {
		return requeueLaterOrNowIfError(r.updateStatus(nc, metav1.ConditionTrue, ConfigurationSucceeded, "Configured successfully"))
	}
}

// CreateEmptyNodeConfigIfNeeded creates empty CR to be Reconciled in near future and filled with Status.
// If invoked before manager's Start, it'll need a direct API client
// (Manager's/Controller's client is cached and cache is not initialized yet).
func (r *NodeConfigReconciler) CreateEmptyNodeConfigIfNeeded(c client.Client) error {
	nodeConfig := &fec.SriovFecNodeConfig{}

	err := c.Get(context.Background(), client.ObjectKey{Name: r.nodeName, Namespace: r.namespace}, nodeConfig)
	if err == nil {
		r.log.Info("already exists")
		return nil
	}

	if !k8serrors.IsNotFound(err) {
		return err
	}

	r.log.Infof("SriovFecNodeConfig{name: %s, namespace: %s} not found - creating", r.nodeName, r.namespace)

	nodeConfig = &fec.SriovFecNodeConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      r.nodeName,
			Namespace: r.namespace,
		},
		Spec: fec.SriovFecNodeConfigSpec{
			PhysicalFunctions: []fec.PhysicalFunctionConfigExt{},
		},
	}

	if createErr := c.Create(context.Background(), nodeConfig); createErr != nil {
		r.log.WithError(createErr).Error("failed to create")
		return createErr
	}

	meta.SetStatusCondition(&nodeConfig.Status.Conditions, metav1.Condition{
		Type:               ConditionConfigured,
		Status:             metav1.ConditionFalse,
		Reason:             string(ConfigurationNotRequested),
		Message:            "",
		ObservedGeneration: nodeConfig.GetGeneration(),
	})

	if inv, err := r.readExistingInventory(); err != nil {
		return err
	} else {
		nodeConfig.Status.Inventory = *inv
	}

	if updateErr := c.Status().Update(context.Background(), nodeConfig); updateErr != nil {
		r.log.WithError(updateErr).Error("failed to update cr status")
		return updateErr
	}

	return nil
}

func (r *NodeConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&fec.SriovFecNodeConfig{}).
		WithEventFilter(
			predicate.And(
				ResourceNamePredicate{
					requiredName: r.nodeName,
					log:          r.log,
				},
				predicate.GenerationChangedPredicate{},
			),
		).Complete(r)
}

func (r *NodeConfigReconciler) updateStatus(nc *fec.SriovFecNodeConfig, status metav1.ConditionStatus, reason ConfigurationConditionReason, msg string) error {
	previousCondition := findOrCreateConfigurationStatusCondition(nc)

	condition := metav1.Condition{
		Type:    ConditionConfigured,
		Status:  status,
		Reason:  string(reason),
		Message: msg,
		ObservedGeneration: func() int64 {
			if reason != ConfigurationInProgress {
				return nc.GetGeneration()
			} else {
				return previousCondition.ObservedGeneration
			}
		}(),
	}

	meta.SetStatusCondition(&nc.Status.Conditions, condition)
	if inv, err := getSriovInventory(r.log); err != nil {
		r.log.WithError(err).
			WithField("reason", condition.Reason).
			WithField("message", condition.Message).
			Error("failed to obtain sriov inventory for the node")
	} else {
		nc.Status.Inventory = *inv
	}

	if err := r.Status().Update(context.Background(), nc); err != nil {
		return err
	}

	r.log.WithField("previous", previousCondition).
		WithField("current", condition).
		Infof("%s condition transition", ConditionConfigured)

	return nil
}

func (r *NodeConfigReconciler) readExistingInventory() (*fec.NodeInventory, error) {
	inv, err := getSriovInventory(r.log)
	if err != nil {
		r.log.WithError(err).Error("failed to obtain sriov inventory for the node")
	}
	return inv, err
}

func (r *NodeConfigReconciler) readSriovFecNodeConfig(nn types.NamespacedName) (nc *fec.SriovFecNodeConfig, err error) {
	getSriovFecNodeConfig := func() (*fec.SriovFecNodeConfig, error) {
		sfnc := new(fec.SriovFecNodeConfig)
		if err := r.Client.Get(context.TODO(), nn, sfnc); err != nil {
			return nil, err
		}
		return sfnc, nil
	}

	if nc, err = getSriovFecNodeConfig(); err != nil {
		if !k8serrors.IsNotFound(err) {
			r.log.WithError(err).Error("Get() failed")
			return nil, err
		}

		r.log.Info("SriovFecNodeConfig not found - creating")
		if err := r.CreateEmptyNodeConfigIfNeeded(r.Client); err != nil {
			r.log.WithError(err).Error("Couldn't create SriovFecNodeConfig")
			return nil, err
		}

		if nc, err = getSriovFecNodeConfig(); err != nil {
			return nil, err
		}
	}

	return nc, nil
}

func (r *NodeConfigReconciler) configureNode(nodeConfig *fec.SriovFecNodeConfig) (isRebootRequested bool, err error) {
	configurer, status := r.createNodeConfigurer(nodeConfig)

	if err = r.drainAndExecute(configurer, !nodeConfig.Spec.DrainSkip); err != nil {
		return false, err
	}

	return status.isRebootRequested, status.configurationError
}

func (r *NodeConfigReconciler) createNodeConfigurer(nodeConfig *fec.SriovFecNodeConfig) (configurer, *configurationStatus) {
	status := new(configurationStatus)

	configurer := func(ctx context.Context) bool {
		missingParams, err := r.nodeConfigurator.isAnyKernelParamsMissing()
		if err != nil {
			r.log.WithError(err).Error("failed to check for missing params")
			status.configurationError = err
			return true
		}

		if missingParams {
			r.log.Info("missing kernel params")

			err := r.nodeConfigurator.addMissingKernelParams()
			if err != nil {
				r.log.WithError(err).Error("failed to add missing params")
				status.configurationError = err
				return true
			}

			r.log.Info("added kernel params - rebooting")
			if err := r.nodeConfigurator.rebootNode(); err != nil {
				r.log.WithError(err).Error("failed to request a node reboot")
				status.configurationError = err
				return true
			}
			status.isRebootRequested = true
			return false // leave node cordoned & keep the leadership
		}
		if err := r.nodeConfigurator.applyConfig(nodeConfig.Spec); err != nil {
			r.log.WithError(err).Error("failed applying new PF/VF configuration")
			status.configurationError = err
			return true
		}

		status.configurationError = r.restartDevicePlugin()
		return true
	}

	return configurer, status
}

func (r *NodeConfigReconciler) restartDevicePlugin() error {
	pods := &corev1.PodList{}
	err := r.Client.List(context.TODO(), pods,
		client.InNamespace(r.namespace),
		&client.MatchingLabels{"app": "sriov-device-plugin-daemonset"})

	if err != nil {
		return errors.Wrap(err, "failed to get pods")
	}
	if len(pods.Items) == 0 {
		r.log.Info("there is no running instance of device plugin, nothing to restart")
	}

	for _, p := range pods.Items {
		if p.Spec.NodeName != r.nodeName {
			continue
		}
		d := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: p.Namespace,
				Name:      p.Name,
			},
		}
		if err := r.Delete(context.TODO(), d, &client.DeleteOptions{}); err != nil {
			return errors.Wrap(err, "failed to delete sriov-device-plugin-daemonset pod")
		}

	}

	return nil
}

func isReconcileRequired(nc *fec.SriovFecNodeConfig) bool {

	isGenerationChanged := func() bool {
		return nc.GetGeneration() != findOrCreateConfigurationStatusCondition(nc).ObservedGeneration
	}

	isReconfigurationRequested := func() bool {
		return len(nc.Spec.PhysicalFunctions) > 0
	}

	isSriovFecConfigured := func() bool {
		for _, accelerator := range nc.Status.Inventory.SriovAccelerators {
			if len(accelerator.VFs) > 0 {
				return true
			}
		}
		return false
	}

	return isGenerationChanged() && (isReconfigurationRequested() || isSriovFecConfigured())
}

func findOrCreateConfigurationStatusCondition(nc *fec.SriovFecNodeConfig) metav1.Condition {
	configurationStatusCondition := nc.FindCondition(ConditionConfigured)
	if configurationStatusCondition == nil {
		return metav1.Condition{
			Type:               ConditionConfigured,
			Status:             metav1.ConditionTrue,
			Reason:             string(ConfigurationNotRequested),
			ObservedGeneration: 0,
		}
	}

	return *configurationStatusCondition
}

//returns error if requested configuration refers to not existing inventory/accelerator
func isConfigurationOfNonExistingInventoryRequested(requestedConfiguration []fec.PhysicalFunctionConfigExt, existingInventory *fec.NodeInventory) bool {
OUTER:
	for _, pf := range requestedConfiguration {
		for _, acc := range existingInventory.SriovAccelerators {
			if acc.PCIAddress == pf.PCIAddress {
				continue OUTER
			}
		}

		return true
	}
	return false
}

//returns result indicating necessity of re-queuing Reconcile after configured resyncPeriod
func requeueLater() (reconcile.Result, error) {
	return reconcile.Result{RequeueAfter: resyncPeriod}, nil
}

//returns result indicating necessity of re-queuing Reconcile(...) immediately; non-nil err will be logged by controller
func requeueNowWithError(e error) (reconcile.Result, error) {
	return reconcile.Result{Requeue: true}, e
}

//returns result indicating necessity of re-queuing Reconcile(...):
//immediately - in case when given err is non-nil;
//on configured schedule, when err is nil
func requeueLaterOrNowIfError(e error) (reconcile.Result, error) {
	return reconcile.Result{RequeueAfter: resyncPeriod}, e
}

type ResourceNamePredicate struct {
	predicate.Funcs
	requiredName string
	log          *logrus.Logger
}

func (r ResourceNamePredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectNew.GetName() != r.requiredName {
		r.log.WithField("expected name", r.requiredName).Info("CR intended for another node - ignoring")
		return false
	}
	return true
}

func (r ResourceNamePredicate) Create(e event.CreateEvent) bool {
	if e.Object.GetName() != r.requiredName {
		r.log.WithField("expected name", r.requiredName).Info("CR intended for another node - ignoring")
		return false
	}
	return true
}

func CreateManager(config *rest.Config, namespace string, scheme *runtime.Scheme) (manager.Manager, error) {
	return ctrl.NewManager(config, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0",
		LeaderElection:     false,
		Namespace:          namespace,
	})
}
