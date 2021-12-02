//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSet) DeepCopyInto(out *KustomizationSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSet.
func (in *KustomizationSet) DeepCopy() *KustomizationSet {
	if in == nil {
		return nil
	}
	out := new(KustomizationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KustomizationSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetGenerator) DeepCopyInto(out *KustomizationSetGenerator) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = new(ListGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetGenerator.
func (in *KustomizationSetGenerator) DeepCopy() *KustomizationSetGenerator {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetList) DeepCopyInto(out *KustomizationSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KustomizationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetList.
func (in *KustomizationSetList) DeepCopy() *KustomizationSetList {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KustomizationSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetSpec) DeepCopyInto(out *KustomizationSetSpec) {
	*out = *in
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]KustomizationSetGenerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetSpec.
func (in *KustomizationSetSpec) DeepCopy() *KustomizationSetSpec {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetStatus) DeepCopyInto(out *KustomizationSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetStatus.
func (in *KustomizationSetStatus) DeepCopy() *KustomizationSetStatus {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetTemplate) DeepCopyInto(out *KustomizationSetTemplate) {
	*out = *in
	in.KustomizationSetTemplateMeta.DeepCopyInto(&out.KustomizationSetTemplateMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetTemplate.
func (in *KustomizationSetTemplate) DeepCopy() *KustomizationSetTemplate {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizationSetTemplateMeta) DeepCopyInto(out *KustomizationSetTemplateMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizationSetTemplateMeta.
func (in *KustomizationSetTemplateMeta) DeepCopy() *KustomizationSetTemplateMeta {
	if in == nil {
		return nil
	}
	out := new(KustomizationSetTemplateMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGenerator) DeepCopyInto(out *ListGenerator) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make([]v1.JSON, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(KustomizationSetTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGenerator.
func (in *ListGenerator) DeepCopy() *ListGenerator {
	if in == nil {
		return nil
	}
	out := new(ListGenerator)
	in.DeepCopyInto(out)
	return out
}
