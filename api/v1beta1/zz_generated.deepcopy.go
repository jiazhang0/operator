// +build !ignore_autogenerated

/*
Copyright 2021 CNCF.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcCompletedStatus) DeepCopyInto(out *CcCompletedStatus) {
	*out = *in
	if in.CompletedNodesList != nil {
		in, out := &in.CompletedNodesList, &out.CompletedNodesList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcCompletedStatus.
func (in *CcCompletedStatus) DeepCopy() *CcCompletedStatus {
	if in == nil {
		return nil
	}
	out := new(CcCompletedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcFailedNodeStatus) DeepCopyInto(out *CcFailedNodeStatus) {
	*out = *in
	if in.FailedNodesList != nil {
		in, out := &in.FailedNodesList, &out.FailedNodesList
		*out = make([]FailedNodeStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcFailedNodeStatus.
func (in *CcFailedNodeStatus) DeepCopy() *CcFailedNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CcFailedNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcInstallConfig) DeepCopyInto(out *CcInstallConfig) {
	*out = *in
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcInstallConfig.
func (in *CcInstallConfig) DeepCopy() *CcInstallConfig {
	if in == nil {
		return nil
	}
	out := new(CcInstallConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcInstallationInProgressStatus) DeepCopyInto(out *CcInstallationInProgressStatus) {
	*out = *in
	if in.BinariesInstalledNodesList != nil {
		in, out := &in.BinariesInstalledNodesList, &out.BinariesInstalledNodesList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcInstallationInProgressStatus.
func (in *CcInstallationInProgressStatus) DeepCopy() *CcInstallationInProgressStatus {
	if in == nil {
		return nil
	}
	out := new(CcInstallationInProgressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcInstallationStatus) DeepCopyInto(out *CcInstallationStatus) {
	*out = *in
	in.InProgress.DeepCopyInto(&out.InProgress)
	in.Completed.DeepCopyInto(&out.Completed)
	in.Failed.DeepCopyInto(&out.Failed)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcInstallationStatus.
func (in *CcInstallationStatus) DeepCopy() *CcInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(CcInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcRuntime) DeepCopyInto(out *CcRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcRuntime.
func (in *CcRuntime) DeepCopy() *CcRuntime {
	if in == nil {
		return nil
	}
	out := new(CcRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CcRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcRuntimeList) DeepCopyInto(out *CcRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CcRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcRuntimeList.
func (in *CcRuntimeList) DeepCopy() *CcRuntimeList {
	if in == nil {
		return nil
	}
	out := new(CcRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CcRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcRuntimeSpec) DeepCopyInto(out *CcRuntimeSpec) {
	*out = *in
	if in.CcNodeSelector != nil {
		in, out := &in.CcNodeSelector, &out.CcNodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcRuntimeSpec.
func (in *CcRuntimeSpec) DeepCopy() *CcRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(CcRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcRuntimeStatus) DeepCopyInto(out *CcRuntimeStatus) {
	*out = *in
	in.InstallationStatus.DeepCopyInto(&out.InstallationStatus)
	in.UnInstallationStatus.DeepCopyInto(&out.UnInstallationStatus)
	out.Upgradestatus = in.Upgradestatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcRuntimeStatus.
func (in *CcRuntimeStatus) DeepCopy() *CcRuntimeStatus {
	if in == nil {
		return nil
	}
	out := new(CcRuntimeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcUnInstallationInProgressStatus) DeepCopyInto(out *CcUnInstallationInProgressStatus) {
	*out = *in
	if in.BinariesUnInstalledNodesList != nil {
		in, out := &in.BinariesUnInstalledNodesList, &out.BinariesUnInstalledNodesList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcUnInstallationInProgressStatus.
func (in *CcUnInstallationInProgressStatus) DeepCopy() *CcUnInstallationInProgressStatus {
	if in == nil {
		return nil
	}
	out := new(CcUnInstallationInProgressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcUnInstallationStatus) DeepCopyInto(out *CcUnInstallationStatus) {
	*out = *in
	in.InProgress.DeepCopyInto(&out.InProgress)
	in.Completed.DeepCopyInto(&out.Completed)
	in.Failed.DeepCopyInto(&out.Failed)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcUnInstallationStatus.
func (in *CcUnInstallationStatus) DeepCopy() *CcUnInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(CcUnInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcUpgradeStatus) DeepCopyInto(out *CcUpgradeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcUpgradeStatus.
func (in *CcUpgradeStatus) DeepCopy() *CcUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(CcUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailedNodeStatus) DeepCopyInto(out *FailedNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailedNodeStatus.
func (in *FailedNodeStatus) DeepCopy() *FailedNodeStatus {
	if in == nil {
		return nil
	}
	out := new(FailedNodeStatus)
	in.DeepCopyInto(out)
	return out
}
