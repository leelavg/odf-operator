/*
Copyright 2021 Red Hat OpenShift Data Foundation.

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

// TLSProfileSpec defines the desired state of TLSProfile
type TLSProfileSpec struct {
}

// TLSProfileStatus defines the observed state of TLSProfile
type TLSProfileStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TLSProfile is the Schema for the tlsprofiles API
type TLSProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TLSProfileSpec   `json:"spec,omitempty"`
	Status TLSProfileStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TLSProfileList contains a list of TLSProfile
type TLSProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TLSProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSProfile{}, &TLSProfileList{})
}
