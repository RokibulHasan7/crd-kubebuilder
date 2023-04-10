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

const (
	SUCCESS = "Success"
	FAILED  = "Failed"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RokibulHasanSpec defines the desired state of RokibulHasan
type RokibulHasanSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RokibulHasan. Edit rokibulhasan_types.go to remove/update

	// +kubebuilder:validation:Maximum=23
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Required
	StartTime int `json:"startTime"`

	// +kubebuilder:validation:Maximum=23
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Required
	EndTime int `json:"endTime"`
	// +kubebuilder:validation:Required
	Replicas    int              `json:"replicas"`
	Deployments []DeploymentList `json:"deployments"`
}

type DeploymentList struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// RokibulHasanStatus defines the observed state of RokibulHasan
type RokibulHasanStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RokibulHasan is the Schema for the rokibulhasans API
type RokibulHasan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RokibulHasanSpec   `json:"spec,omitempty"`
	Status RokibulHasanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RokibulHasanList contains a list of RokibulHasan
type RokibulHasanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RokibulHasan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RokibulHasan{}, &RokibulHasanList{})
}
