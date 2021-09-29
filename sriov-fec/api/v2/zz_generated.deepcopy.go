// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020-2021 Intel Corporation

// Code generated by controller-gen. DO NOT EDIT.

package v2

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACC100BBDevConfig) DeepCopyInto(out *ACC100BBDevConfig) {
	*out = *in
	out.Uplink4G = in.Uplink4G
	out.Downlink4G = in.Downlink4G
	out.Uplink5G = in.Uplink5G
	out.Downlink5G = in.Downlink5G
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACC100BBDevConfig.
func (in *ACC100BBDevConfig) DeepCopy() *ACC100BBDevConfig {
	if in == nil {
		return nil
	}
	out := new(ACC100BBDevConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorSelector) DeepCopyInto(out *AcceleratorSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorSelector.
func (in *AcceleratorSelector) DeepCopy() *AcceleratorSelector {
	if in == nil {
		return nil
	}
	out := new(AcceleratorSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BBDevConfig) DeepCopyInto(out *BBDevConfig) {
	*out = *in
	if in.N3000 != nil {
		in, out := &in.N3000, &out.N3000
		*out = new(N3000BBDevConfig)
		**out = **in
	}
	if in.ACC100 != nil {
		in, out := &in.ACC100, &out.ACC100
		*out = new(ACC100BBDevConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BBDevConfig.
func (in *BBDevConfig) DeepCopy() *BBDevConfig {
	if in == nil {
		return nil
	}
	out := new(BBDevConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ByPriority) DeepCopyInto(out *ByPriority) {
	{
		in := &in
		*out = make(ByPriority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByPriority.
func (in ByPriority) DeepCopy() ByPriority {
	if in == nil {
		return nil
	}
	out := new(ByPriority)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *N3000BBDevConfig) DeepCopyInto(out *N3000BBDevConfig) {
	*out = *in
	out.Downlink = in.Downlink
	out.Uplink = in.Uplink
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new N3000BBDevConfig.
func (in *N3000BBDevConfig) DeepCopy() *N3000BBDevConfig {
	if in == nil {
		return nil
	}
	out := new(N3000BBDevConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.PhysicalFunctions != nil {
		in, out := &in.PhysicalFunctions, &out.PhysicalFunctions
		*out = make([]PhysicalFunctionConfigExt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeInventory) DeepCopyInto(out *NodeInventory) {
	*out = *in
	if in.SriovAccelerators != nil {
		in, out := &in.SriovAccelerators, &out.SriovAccelerators
		*out = make([]SriovAccelerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeInventory.
func (in *NodeInventory) DeepCopy() *NodeInventory {
	if in == nil {
		return nil
	}
	out := new(NodeInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalFunctionConfig) DeepCopyInto(out *PhysicalFunctionConfig) {
	*out = *in
	in.BBDevConfig.DeepCopyInto(&out.BBDevConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalFunctionConfig.
func (in *PhysicalFunctionConfig) DeepCopy() *PhysicalFunctionConfig {
	if in == nil {
		return nil
	}
	out := new(PhysicalFunctionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PhysicalFunctionConfigExt) DeepCopyInto(out *PhysicalFunctionConfigExt) {
	*out = *in
	in.BBDevConfig.DeepCopyInto(&out.BBDevConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PhysicalFunctionConfigExt.
func (in *PhysicalFunctionConfigExt) DeepCopy() *PhysicalFunctionConfigExt {
	if in == nil {
		return nil
	}
	out := new(PhysicalFunctionConfigExt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueGroupConfig) DeepCopyInto(out *QueueGroupConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueGroupConfig.
func (in *QueueGroupConfig) DeepCopy() *QueueGroupConfig {
	if in == nil {
		return nil
	}
	out := new(QueueGroupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovAccelerator) DeepCopyInto(out *SriovAccelerator) {
	*out = *in
	if in.VFs != nil {
		in, out := &in.VFs, &out.VFs
		*out = make([]VF, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovAccelerator.
func (in *SriovAccelerator) DeepCopy() *SriovAccelerator {
	if in == nil {
		return nil
	}
	out := new(SriovAccelerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecClusterConfig) DeepCopyInto(out *SriovFecClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecClusterConfig.
func (in *SriovFecClusterConfig) DeepCopy() *SriovFecClusterConfig {
	if in == nil {
		return nil
	}
	out := new(SriovFecClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovFecClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecClusterConfigList) DeepCopyInto(out *SriovFecClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovFecClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecClusterConfigList.
func (in *SriovFecClusterConfigList) DeepCopy() *SriovFecClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(SriovFecClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovFecClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecClusterConfigSpec) DeepCopyInto(out *SriovFecClusterConfigSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.AcceleratorSelector = in.AcceleratorSelector
	in.PhysicalFunction.DeepCopyInto(&out.PhysicalFunction)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecClusterConfigSpec.
func (in *SriovFecClusterConfigSpec) DeepCopy() *SriovFecClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SriovFecClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecClusterConfigStatus) DeepCopyInto(out *SriovFecClusterConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecClusterConfigStatus.
func (in *SriovFecClusterConfigStatus) DeepCopy() *SriovFecClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SriovFecClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecNodeConfig) DeepCopyInto(out *SriovFecNodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecNodeConfig.
func (in *SriovFecNodeConfig) DeepCopy() *SriovFecNodeConfig {
	if in == nil {
		return nil
	}
	out := new(SriovFecNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovFecNodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecNodeConfigList) DeepCopyInto(out *SriovFecNodeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovFecNodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecNodeConfigList.
func (in *SriovFecNodeConfigList) DeepCopy() *SriovFecNodeConfigList {
	if in == nil {
		return nil
	}
	out := new(SriovFecNodeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovFecNodeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecNodeConfigSpec) DeepCopyInto(out *SriovFecNodeConfigSpec) {
	*out = *in
	if in.PhysicalFunctions != nil {
		in, out := &in.PhysicalFunctions, &out.PhysicalFunctions
		*out = make([]PhysicalFunctionConfigExt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecNodeConfigSpec.
func (in *SriovFecNodeConfigSpec) DeepCopy() *SriovFecNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SriovFecNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovFecNodeConfigStatus) DeepCopyInto(out *SriovFecNodeConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Inventory.DeepCopyInto(&out.Inventory)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovFecNodeConfigStatus.
func (in *SriovFecNodeConfigStatus) DeepCopy() *SriovFecNodeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SriovFecNodeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkDownlink) DeepCopyInto(out *UplinkDownlink) {
	*out = *in
	out.Queues = in.Queues
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkDownlink.
func (in *UplinkDownlink) DeepCopy() *UplinkDownlink {
	if in == nil {
		return nil
	}
	out := new(UplinkDownlink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UplinkDownlinkQueues) DeepCopyInto(out *UplinkDownlinkQueues) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UplinkDownlinkQueues.
func (in *UplinkDownlinkQueues) DeepCopy() *UplinkDownlinkQueues {
	if in == nil {
		return nil
	}
	out := new(UplinkDownlinkQueues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VF) DeepCopyInto(out *VF) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VF.
func (in *VF) DeepCopy() *VF {
	if in == nil {
		return nil
	}
	out := new(VF)
	in.DeepCopyInto(out)
	return out
}
