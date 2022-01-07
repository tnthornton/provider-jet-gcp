/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LocationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type LocationParameters struct {

	// Name of the Location resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// LocationSpec defines the desired state of Location
type LocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationParameters `json:"forProvider"`
}

// LocationStatus defines the observed state of Location.
type LocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Location is the Schema for the Locations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Location struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationSpec   `json:"spec"`
	Status            LocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationList contains a list of Locations
type LocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Location `json:"items"`
}

// Repository type metadata.
var (
	Location_Kind             = "Location"
	Location_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Location_Kind}.String()
	Location_KindAPIVersion   = Location_Kind + "." + CRDGroupVersion.String()
	Location_GroupVersionKind = CRDGroupVersion.WithKind(Location_Kind)
)

func init() {
	SchemeBuilder.Register(&Location{}, &LocationList{})
}
