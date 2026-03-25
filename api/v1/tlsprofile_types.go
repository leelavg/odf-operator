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

// TLSProfileSelector identifies which components should use a specific TLS configuration.
// TODO: Define Enums for component selectors (e.g., "noobaa/s3", "ceph/rgw")
type TLSProfileSelector string

// TLSProtocolVersion represents a TLS protocol version.
// Only TLS 1.2 and 1.3 are supported as TLS 1.0 and 1.1 are considered vulnerable.
// +kubebuilder:validation:Enum=TLSv1.2;TLSv1.3
type TLSProtocolVersion string

const (
	// VersionTLS12 is version 1.2 of the TLS security protocol.
	VersionTLS12 TLSProtocolVersion = "TLSv1.2"
	// VersionTLS13 is version 1.3 of the TLS security protocol.
	VersionTLS13 TLSProtocolVersion = "TLSv1.3"
)

// TLSProfileConfig defines the TLS security configuration including protocol version,
// cipher suites, and elliptic curve groups.
type TLSProfileConfig struct {
	// MinVersion specifies the minimum TLS protocol version to accept.
	// Must be either "TLSv1.2" or "TLSv1.3".
	// +optional
	MinVersion TLSProtocolVersion `json:"minVersion,omitzero"`

	// Ciphers is a list of IANA cipher suite names to enable.
	// Example: ["TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
	// +optional
	Ciphers []string `json:"ciphers,omitzero"`

	// Groups is a list of elliptic curve group names to enable for key exchange.
	// Use GroupToID and GroupToOpenSSL mappings to translate to Go/OpenSSL formats.
	// Example: ["X25519", "prime256v1", "secp384r1"]
	// Supports post-quantum hybrid groups like "X25519MLKEM768".
	// +optional
	Groups []string `json:"groups,omitzero"`
}

// TLSProfileRules defines a TLS configuration rule that applies to specific components.
type TLSProfileRules struct {
	// Selector identifies which components this rule applies to.
	// +optional
	Selector []TLSProfileSelector `json:"selector,omitzero"`

	// Config is the TLS configuration to apply to the selected components.
	// +optional
	Config TLSProfileConfig `json:"config,omitzero"`
}

// TLSProfileSpec defines the desired state of TLSProfile.
// It contains a list of rules that map TLS configurations to specific components.
type TLSProfileSpec struct {
	// Rules is a list of TLS configuration rules.
	// Each rule specifies which components should use a particular TLS configuration.
	// Overrides happend based on Selectors in each rule
	// +optional
	Rules []TLSProfileRules `json:"rules,omitzero"`
}

// TLSProfileComponents represents the observed TLS configuration for a specific component.
type TLSProfileComponents struct {
	// Component is the name of the component (e.g., "noobaa/s3", "ceph/rgw").
	Component string `json:"component"`

	// ObservedGeneration is the generation of the TLSProfile that was last processed
	// by this component. Used to track whether the component has applied the latest config.
	ObservedGeneration int64 `json:"ObservedGeneration"`

	// TLSProfileConfig is the actual TLS configuration applied to this component.
	TLSProfileConfig `json:",inline"`
}

// TLSProfileStatus defines the observed state of TLSProfile.
// It tracks which components have applied the TLS configuration and what settings they're using.
type TLSProfileStatus struct {
	// Components is a list of components and their observed TLS configurations.
	// Each entry shows what TLS settings a component is currently using.
	// +optional
	Componenets []TLSProfileComponents `json:"componenets,omitzero"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TLSProfile is the Schema for the tlsprofiles API.
// It allows administrators to configure TLS settings (protocol versions, ciphers, curves)
// for ODF components in a centralized way.
type TLSProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// Spec defines the desired TLS configuration rules.
	Spec TLSProfileSpec `json:"spec,omitzero"`

	// Status shows the observed TLS configuration state across components.
	Status TLSProfileStatus `json:"status,omitzero"`
}

//+kubebuilder:object:root=true

// TLSProfileList contains a list of TLSProfile resources.
type TLSProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []TLSProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TLSProfile{}, &TLSProfileList{})
}
