/*
Copyright 2024.

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

// StandaloneSpec defines the desired state of Standalone
type StandaloneSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	RedisConfig      RedisConfig      `json:"RedisConfig,omitempty"`
	KubernetesConfig KubernetesConfig `json:"KubernetesConfig"`
}

type RedisConfig struct {
	MaxmemoryPolicy string `json:"maxmemoryPolicy,omitempty"`
	Client          string `json:"client,omitempty"`
	MaxMmemory      string `json:"maxMmemory,omitempty"`
	Persistence     string `json:"persistence,omitempty"`
}

// StandaloneStatus defines the observed state of Standalone
type StandaloneStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Standalone is the Schema for the standalones API
type Standalone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StandaloneSpec   `json:"spec"`
	Status            StandaloneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StandaloneList contains a list of Standalone
type StandaloneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Standalone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Standalone{}, &StandaloneList{})
}
