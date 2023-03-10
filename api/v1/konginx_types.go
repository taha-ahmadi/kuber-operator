/*
Copyright 2023.

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

func init() {
	SchemeBuilder.Register(&KONginx{}, &KONginxList{})
}

// KONginxSpec defines the desired state of OvhNginx
type KONginxSpec struct {
	// Number of replicas for the Nginx Pods
	ReplicaCount int32 `json:"replicaCount"`
	// Exposed port for the Nginx server
	Port int32 `json:"port"`
}

// KONginxStatus defines the observed state of OvhNginx
type KONginxStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KONginx is the Schema for the ovhnginxes API
type KONginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KONginxSpec   `json:"spec,omitempty"`
	Status KONginxStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KONginxList contains a list of OvhNginx
type KONginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KONginx `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KONginx{}, &KONginxList{})
}