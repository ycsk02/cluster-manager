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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfigList) DeepCopyInto(out *KubeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfigList.
func (in *KubeConfigList) DeepCopy() *KubeConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfigSpec) DeepCopyInto(out *KubeConfigSpec) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]KubeconfigNamedCluster, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]KubeconfigUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make([]KubeconfigNamedContext, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfigSpec.
func (in *KubeConfigSpec) DeepCopy() *KubeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfigStatus) DeepCopyInto(out *KubeConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfigStatus.
func (in *KubeConfigStatus) DeepCopy() *KubeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigAuthProvider) DeepCopyInto(out *KubeconfigAuthProvider) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigAuthProvider.
func (in *KubeconfigAuthProvider) DeepCopy() *KubeconfigAuthProvider {
	if in == nil {
		return nil
	}
	out := new(KubeconfigAuthProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigCluster) DeepCopyInto(out *KubeconfigCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigCluster.
func (in *KubeconfigCluster) DeepCopy() *KubeconfigCluster {
	if in == nil {
		return nil
	}
	out := new(KubeconfigCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigContext) DeepCopyInto(out *KubeconfigContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigContext.
func (in *KubeconfigContext) DeepCopy() *KubeconfigContext {
	if in == nil {
		return nil
	}
	out := new(KubeconfigContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigNamedCluster) DeepCopyInto(out *KubeconfigNamedCluster) {
	*out = *in
	out.Cluster = in.Cluster
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigNamedCluster.
func (in *KubeconfigNamedCluster) DeepCopy() *KubeconfigNamedCluster {
	if in == nil {
		return nil
	}
	out := new(KubeconfigNamedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigNamedContext) DeepCopyInto(out *KubeconfigNamedContext) {
	*out = *in
	out.Context = in.Context
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigNamedContext.
func (in *KubeconfigNamedContext) DeepCopy() *KubeconfigNamedContext {
	if in == nil {
		return nil
	}
	out := new(KubeconfigNamedContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigUser) DeepCopyInto(out *KubeconfigUser) {
	*out = *in
	in.User.DeepCopyInto(&out.User)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigUser.
func (in *KubeconfigUser) DeepCopy() *KubeconfigUser {
	if in == nil {
		return nil
	}
	out := new(KubeconfigUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeconfigUserKeyPair) DeepCopyInto(out *KubeconfigUserKeyPair) {
	*out = *in
	in.AuthProvider.DeepCopyInto(&out.AuthProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeconfigUserKeyPair.
func (in *KubeconfigUserKeyPair) DeepCopy() *KubeconfigUserKeyPair {
	if in == nil {
		return nil
	}
	out := new(KubeconfigUserKeyPair)
	in.DeepCopyInto(out)
	return out
}
