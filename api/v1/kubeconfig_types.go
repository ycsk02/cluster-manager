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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubeconfigUser is a struct used to create a kubectl configuration YAML file
type KubeconfigUser struct {
	Name string                `json:"name" yaml:"name"`
	User KubeconfigUserKeyPair `json:"user" yaml:"user"`
}

// KubeconfigUserKeyPair is a struct used to create a kubectl configuration YAML file
type KubeconfigUserKeyPair struct {
	ClientCertificateData string                 `json:"client-certificate-data" yaml:"client-certificate-data"`
	ClientKeyData         string                 `json:"client-key-data" yaml:"client-key-data"`
	AuthProvider          KubeconfigAuthProvider `json:"auth-provider" yaml:"auth-provider,omitempty"`
}

// KubeconfigAuthProvider is a struct used to create a kubectl authentication provider
type KubeconfigAuthProvider struct {
	Name   string            `json:"name" yaml:"name"`
	Config map[string]string `json:"config" yaml:"config"`
}

// KubeconfigNamedCluster is a struct used to create a kubectl configuration YAML file
type KubeconfigNamedCluster struct {
	Name    string            `json:"name" yaml:"name"`
	Cluster KubeconfigCluster `json:"cluster" yaml:"cluster"`
}

// KubeconfigCluster is a struct used to create a kubectl configuration YAML file
type KubeconfigCluster struct {
	Server                   string `json:"server" yaml:"server"`
	CertificateAuthorityData string `json:"certificate-authority-data" yaml:"certificate-authority-data"`
	CertificateAuthority     string `json:"certificate-authority" yaml:"certificate-authority,omitempty"`
}

// KubeconfigNamedContext is a struct used to create a kubectl configuration YAML file
type KubeconfigNamedContext struct {
	Name    string            `json:"name" yaml:"name"`
	Context KubeconfigContext `json:"context" yaml:"context"`
}

// KubeconfigContext is a struct used to create a kubectl configuration YAML file
type KubeconfigContext struct {
	Cluster   string `json:"cluster" yaml:"cluster"`
	Namespace string `json:"namespace" yaml:"namespace,omitempty"`
	User      string `json:"user" yaml:"user"`
}

type KubeconfigPreferences struct {
}

type KubeConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Clusters       []KubeconfigNamedCluster `json:"clusters" yaml:"clusters"`
	Users          []KubeconfigUser         `json:"users" yaml:"users"`
	Contexts       []KubeconfigNamedContext `json:"contexts" yaml:"contexts"`
	CurrentContext string                   `json:"current-context" yaml:"current-context"`
	Preferences    KubeconfigPreferences    `json:"preferences,omitempty" yaml:"preferences,omitempty"`
}

// KubeConfigStatus defines the observed state of KubeConfig
type KubeConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// KubeConfig is the Schema for the kubeconfigs API
type KubeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeConfigSpec   `json:"spec,omitempty"`
	Status KubeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeConfigList contains a list of KubeConfig
type KubeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeConfig{}, &KubeConfigList{})
}
