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

type TLSProfileSelector string

// Versions 1.0 and 1.1 are not provided as they are considered vulnerable
// +kubebuilder:validation:Enum=TLSv1.2;TLSv1.3
type TLSProtocolVersion string

const (
	// VersionTLSv12 is version 1.2 of the TLS security protocol.
	VersionTLS12 TLSProtocolVersion = "TLSv1.2"
	// VersionTLSv13 is version 1.3 of the TLS security protocol.
	VersionTLS13 TLSProtocolVersion = "TLSv1.3"
)

type TLSProfileConfig struct {
	MinVersion TLSProtocolVersion `json:"minVersion,omitzero"`
	Ciphers    []string           `json:"ciphers,omitzero"`
	Groups     []string           `json:"groups,omitzero"`
}

type TLSProfileRules struct {
	// +optional
	Selector []TLSProfileSelector `json:"selector,omitzero"`
	// +optional
	Config TLSProfileConfig `json:"config,omitzero"`
}

// TLSProfileSpec defines the desired state of TLSProfile
type TLSProfileSpec struct {
	// +optional
	Rules []TLSProfileRules `json:"rules,omitzero"`
}

// TLSProfileComponents defines the actual state of TLSProfile agains each
// of the component
type TLSProfileComponents struct {
	Component          string `json:"component"`
	ObservedGeneration int64  `json:"ObservedGeneration"`
	TLSProfileConfig   `json:",inline"`
}

// TLSProfileStatus defines the observed state of TLSProfile
type TLSProfileStatus struct {
	// +optional
	Componenets []TLSProfileComponents `json:"componenets,omitzero"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TLSProfile is the Schema for the tlsprofiles API
type TLSProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitzero"`

	Spec   TLSProfileSpec   `json:"spec,omitzero"`
	Status TLSProfileStatus `json:"status,omitzero"`
}

//+kubebuilder:object:root=true

// TLSProfileList contains a list of TLSProfile
type TLSProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []TLSProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSProfile{}, &TLSProfileList{})
}
