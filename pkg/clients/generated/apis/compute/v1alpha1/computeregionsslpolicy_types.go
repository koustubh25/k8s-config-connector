// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeRegionSSLPolicySpec struct {
	/* A list of features enabled when the selected profile is CUSTOM. The
	method returns the set of features that can be specified in this
	list. This field must be empty if the profile is not CUSTOM.

	See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	for which ciphers are available to use. **Note**: this argument
	*must* be present when using the 'CUSTOM' profile. This argument
	*must not* be present when using any other profile. */
	// +optional
	CustomFeatures []string `json:"customFeatures,omitempty"`

	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The minimum version of SSL protocol that can be used by the clients
	to establish a connection with the load balancer. Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]. */
	// +optional
	MinTlsVersion *string `json:"minTlsVersion,omitempty"`

	/* Profile specifies the set of SSL features that can be used by the
	load balancer when negotiating SSL with clients. If using 'CUSTOM',
	the set of SSL features to enable must be specified in the
	'customFeatures' field.

	See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	for information on what cipher suites each profile provides. If
	'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]. */
	// +optional
	Profile *string `json:"profile,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region where the regional SSL policy resides. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeRegionSSLPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRegionSSLPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The list of features enabled in the SSL policy. */
	// +optional
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`

	/* Fingerprint of this resource. A hash of the contents stored in this
	object. This field is used in optimistic locking. */
	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRegionSSLPolicy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRegionSSLPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionSSLPolicySpec   `json:"spec,omitempty"`
	Status ComputeRegionSSLPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRegionSSLPolicyList contains a list of ComputeRegionSSLPolicy
type ComputeRegionSSLPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionSSLPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionSSLPolicy{}, &ComputeRegionSSLPolicyList{})
}