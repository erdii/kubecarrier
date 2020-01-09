// +build !ignore_autogenerated

/*
Copyright 2020 The KubeCarrier Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCarrier) DeepCopyInto(out *KubeCarrier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCarrier.
func (in *KubeCarrier) DeepCopy() *KubeCarrier {
	if in == nil {
		return nil
	}
	out := new(KubeCarrier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeCarrier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCarrierCondition) DeepCopyInto(out *KubeCarrierCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCarrierCondition.
func (in *KubeCarrierCondition) DeepCopy() *KubeCarrierCondition {
	if in == nil {
		return nil
	}
	out := new(KubeCarrierCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCarrierList) DeepCopyInto(out *KubeCarrierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeCarrier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCarrierList.
func (in *KubeCarrierList) DeepCopy() *KubeCarrierList {
	if in == nil {
		return nil
	}
	out := new(KubeCarrierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeCarrierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCarrierSpec) DeepCopyInto(out *KubeCarrierSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCarrierSpec.
func (in *KubeCarrierSpec) DeepCopy() *KubeCarrierSpec {
	if in == nil {
		return nil
	}
	out := new(KubeCarrierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeCarrierStatus) DeepCopyInto(out *KubeCarrierStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KubeCarrierCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeCarrierStatus.
func (in *KubeCarrierStatus) DeepCopy() *KubeCarrierStatus {
	if in == nil {
		return nil
	}
	out := new(KubeCarrierStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterRegistration) DeepCopyInto(out *ServiceClusterRegistration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterRegistration.
func (in *ServiceClusterRegistration) DeepCopy() *ServiceClusterRegistration {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterRegistration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceClusterRegistration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterRegistrationCondition) DeepCopyInto(out *ServiceClusterRegistrationCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterRegistrationCondition.
func (in *ServiceClusterRegistrationCondition) DeepCopy() *ServiceClusterRegistrationCondition {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterRegistrationCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterRegistrationList) DeepCopyInto(out *ServiceClusterRegistrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceClusterRegistration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterRegistrationList.
func (in *ServiceClusterRegistrationList) DeepCopy() *ServiceClusterRegistrationList {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterRegistrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceClusterRegistrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterRegistrationSpec) DeepCopyInto(out *ServiceClusterRegistrationSpec) {
	*out = *in
	out.KubeconfigSecret = in.KubeconfigSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterRegistrationSpec.
func (in *ServiceClusterRegistrationSpec) DeepCopy() *ServiceClusterRegistrationSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterRegistrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterRegistrationStatus) DeepCopyInto(out *ServiceClusterRegistrationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceClusterRegistrationCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterRegistrationStatus.
func (in *ServiceClusterRegistrationStatus) DeepCopy() *ServiceClusterRegistrationStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterRegistrationStatus)
	in.DeepCopyInto(out)
	return out
}
