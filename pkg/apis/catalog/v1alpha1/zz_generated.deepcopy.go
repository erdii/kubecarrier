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
	"k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCondition) DeepCopyInto(out *AccountCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCondition.
func (in *AccountCondition) DeepCopy() *AccountCondition {
	if in == nil {
		return nil
	}
	out := new(AccountCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountMetadata) DeepCopyInto(out *AccountMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountMetadata.
func (in *AccountMetadata) DeepCopy() *AccountMetadata {
	if in == nil {
		return nil
	}
	out := new(AccountMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	out.Metadata = in.Metadata
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]AccountRole, len(*in))
		copy(*out, *in)
	}
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]v1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	out.Namespace = in.Namespace
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccountCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDInformation) DeepCopyInto(out *CRDInformation) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]CRDVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ServiceCluster = in.ServiceCluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDInformation.
func (in *CRDInformation) DeepCopy() *CRDInformation {
	if in == nil {
		return nil
	}
	out := new(CRDInformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRDVersion) DeepCopyInto(out *CRDVersion) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensionsv1.CustomResourceValidation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRDVersion.
func (in *CRDVersion) DeepCopy() *CRDVersion {
	if in == nil {
		return nil
	}
	out := new(CRDVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Catalog) DeepCopyInto(out *Catalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Catalog.
func (in *Catalog) DeepCopy() *Catalog {
	if in == nil {
		return nil
	}
	out := new(Catalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Catalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogCondition) DeepCopyInto(out *CatalogCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogCondition.
func (in *CatalogCondition) DeepCopy() *CatalogCondition {
	if in == nil {
		return nil
	}
	out := new(CatalogCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntry) DeepCopyInto(out *CatalogEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntry.
func (in *CatalogEntry) DeepCopy() *CatalogEntry {
	if in == nil {
		return nil
	}
	out := new(CatalogEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntryCondition) DeepCopyInto(out *CatalogEntryCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntryCondition.
func (in *CatalogEntryCondition) DeepCopy() *CatalogEntryCondition {
	if in == nil {
		return nil
	}
	out := new(CatalogEntryCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntryList) DeepCopyInto(out *CatalogEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntryList.
func (in *CatalogEntryList) DeepCopy() *CatalogEntryList {
	if in == nil {
		return nil
	}
	out := new(CatalogEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntryMetadata) DeepCopyInto(out *CatalogEntryMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntryMetadata.
func (in *CatalogEntryMetadata) DeepCopy() *CatalogEntryMetadata {
	if in == nil {
		return nil
	}
	out := new(CatalogEntryMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySet) DeepCopyInto(out *CatalogEntrySet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySet.
func (in *CatalogEntrySet) DeepCopy() *CatalogEntrySet {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogEntrySet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySetCondition) DeepCopyInto(out *CatalogEntrySetCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySetCondition.
func (in *CatalogEntrySetCondition) DeepCopy() *CatalogEntrySetCondition {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySetList) DeepCopyInto(out *CatalogEntrySetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogEntrySet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySetList.
func (in *CatalogEntrySetList) DeepCopy() *CatalogEntrySetList {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogEntrySetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySetMetadata) DeepCopyInto(out *CatalogEntrySetMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySetMetadata.
func (in *CatalogEntrySetMetadata) DeepCopy() *CatalogEntrySetMetadata {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySetMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySetSpec) DeepCopyInto(out *CatalogEntrySetSpec) {
	*out = *in
	out.Metadata = in.Metadata
	if in.Derive != nil {
		in, out := &in.Derive, &out.Derive
		*out = new(DerivedConfig)
		(*in).DeepCopyInto(*out)
	}
	in.Discover.DeepCopyInto(&out.Discover)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySetSpec.
func (in *CatalogEntrySetSpec) DeepCopy() *CatalogEntrySetSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySetStatus) DeepCopyInto(out *CatalogEntrySetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CatalogEntrySetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySetStatus.
func (in *CatalogEntrySetStatus) DeepCopy() *CatalogEntrySetStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntrySpec) DeepCopyInto(out *CatalogEntrySpec) {
	*out = *in
	out.Metadata = in.Metadata
	out.BaseCRD = in.BaseCRD
	if in.Derive != nil {
		in, out := &in.Derive, &out.Derive
		*out = new(DerivedConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntrySpec.
func (in *CatalogEntrySpec) DeepCopy() *CatalogEntrySpec {
	if in == nil {
		return nil
	}
	out := new(CatalogEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogEntryStatus) DeepCopyInto(out *CatalogEntryStatus) {
	*out = *in
	if in.TenantCRD != nil {
		in, out := &in.TenantCRD, &out.TenantCRD
		*out = new(CRDInformation)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderCRD != nil {
		in, out := &in.ProviderCRD, &out.ProviderCRD
		*out = new(CRDInformation)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CatalogEntryCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogEntryStatus.
func (in *CatalogEntryStatus) DeepCopy() *CatalogEntryStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogList) DeepCopyInto(out *CatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Catalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogList.
func (in *CatalogList) DeepCopy() *CatalogList {
	if in == nil {
		return nil
	}
	out := new(CatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSpec) DeepCopyInto(out *CatalogSpec) {
	*out = *in
	if in.CatalogEntrySelector != nil {
		in, out := &in.CatalogEntrySelector, &out.CatalogEntrySelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TenantSelector != nil {
		in, out := &in.TenantSelector, &out.TenantSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSpec.
func (in *CatalogSpec) DeepCopy() *CatalogSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogStatus) DeepCopyInto(out *CatalogStatus) {
	*out = *in
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make([]ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CatalogCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogStatus.
func (in *CatalogStatus) DeepCopy() *CatalogStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDiscoverySetConfig) DeepCopyInto(out *CustomResourceDiscoverySetConfig) {
	*out = *in
	out.CRD = in.CRD
	in.ServiceClusterSelector.DeepCopyInto(&out.ServiceClusterSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDiscoverySetConfig.
func (in *CustomResourceDiscoverySetConfig) DeepCopy() *CustomResourceDiscoverySetConfig {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDiscoverySetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedConfig) DeepCopyInto(out *DerivedConfig) {
	*out = *in
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = make([]VersionExposeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedConfig.
func (in *DerivedConfig) DeepCopy() *DerivedConfig {
	if in == nil {
		return nil
	}
	out := new(DerivedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResource) DeepCopyInto(out *DerivedCustomResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResource.
func (in *DerivedCustomResource) DeepCopy() *DerivedCustomResource {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DerivedCustomResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResourceCondition) DeepCopyInto(out *DerivedCustomResourceCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResourceCondition.
func (in *DerivedCustomResourceCondition) DeepCopy() *DerivedCustomResourceCondition {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResourceList) DeepCopyInto(out *DerivedCustomResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DerivedCustomResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResourceList.
func (in *DerivedCustomResourceList) DeepCopy() *DerivedCustomResourceList {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DerivedCustomResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResourceReference) DeepCopyInto(out *DerivedCustomResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResourceReference.
func (in *DerivedCustomResourceReference) DeepCopy() *DerivedCustomResourceReference {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResourceSpec) DeepCopyInto(out *DerivedCustomResourceSpec) {
	*out = *in
	out.BaseCRD = in.BaseCRD
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = make([]VersionExposeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResourceSpec.
func (in *DerivedCustomResourceSpec) DeepCopy() *DerivedCustomResourceSpec {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DerivedCustomResourceStatus) DeepCopyInto(out *DerivedCustomResourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DerivedCustomResourceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DerivedCR != nil {
		in, out := &in.DerivedCR, &out.DerivedCR
		*out = new(DerivedCustomResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DerivedCustomResourceStatus.
func (in *DerivedCustomResourceStatus) DeepCopy() *DerivedCustomResourceStatus {
	if in == nil {
		return nil
	}
	out := new(DerivedCustomResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldPath) DeepCopyInto(out *FieldPath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldPath.
func (in *FieldPath) DeepCopy() *FieldPath {
	if in == nil {
		return nil
	}
	out := new(FieldPath)
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
func (in *Offering) DeepCopyInto(out *Offering) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Offering.
func (in *Offering) DeepCopy() *Offering {
	if in == nil {
		return nil
	}
	out := new(Offering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Offering) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfferingList) DeepCopyInto(out *OfferingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Offering, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfferingList.
func (in *OfferingList) DeepCopy() *OfferingList {
	if in == nil {
		return nil
	}
	out := new(OfferingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OfferingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfferingMetadata) DeepCopyInto(out *OfferingMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfferingMetadata.
func (in *OfferingMetadata) DeepCopy() *OfferingMetadata {
	if in == nil {
		return nil
	}
	out := new(OfferingMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfferingSpec) DeepCopyInto(out *OfferingSpec) {
	*out = *in
	out.Metadata = in.Metadata
	out.Provider = in.Provider
	in.CRD.DeepCopyInto(&out.CRD)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfferingSpec.
func (in *OfferingSpec) DeepCopy() *OfferingSpec {
	if in == nil {
		return nil
	}
	out := new(OfferingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Provider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderList) DeepCopyInto(out *ProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderList.
func (in *ProviderList) DeepCopy() *ProviderList {
	if in == nil {
		return nil
	}
	out := new(ProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSpec) DeepCopyInto(out *ProviderSpec) {
	*out = *in
	out.Metadata = in.Metadata
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSpec.
func (in *ProviderSpec) DeepCopy() *ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterReference) DeepCopyInto(out *ServiceClusterReference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterReference.
func (in *ServiceClusterReference) DeepCopy() *ServiceClusterReference {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceClusterReference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterReferenceList) DeepCopyInto(out *ServiceClusterReferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceClusterReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterReferenceList.
func (in *ServiceClusterReferenceList) DeepCopy() *ServiceClusterReferenceList {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterReferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceClusterReferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceClusterReferenceSpec) DeepCopyInto(out *ServiceClusterReferenceSpec) {
	*out = *in
	out.Metadata = in.Metadata
	out.Provider = in.Provider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceClusterReferenceSpec.
func (in *ServiceClusterReferenceSpec) DeepCopy() *ServiceClusterReferenceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceClusterReferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenant) DeepCopyInto(out *Tenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenant.
func (in *Tenant) DeepCopy() *Tenant {
	if in == nil {
		return nil
	}
	out := new(Tenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantList) DeepCopyInto(out *TenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantList.
func (in *TenantList) DeepCopy() *TenantList {
	if in == nil {
		return nil
	}
	out := new(TenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantSpec) DeepCopyInto(out *TenantSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantSpec.
func (in *TenantSpec) DeepCopy() *TenantSpec {
	if in == nil {
		return nil
	}
	out := new(TenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionExposeConfig) DeepCopyInto(out *VersionExposeConfig) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]FieldPath, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionExposeConfig.
func (in *VersionExposeConfig) DeepCopy() *VersionExposeConfig {
	if in == nil {
		return nil
	}
	out := new(VersionExposeConfig)
	in.DeepCopyInto(out)
	return out
}
