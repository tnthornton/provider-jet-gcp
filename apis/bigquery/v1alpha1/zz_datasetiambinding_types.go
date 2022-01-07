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

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type DatasetIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatasetIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	DatasetID *string `json:"datasetId" tf:"dataset_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// DatasetIAMBindingSpec defines the desired state of DatasetIAMBinding
type DatasetIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetIAMBindingParameters `json:"forProvider"`
}

// DatasetIAMBindingStatus defines the observed state of DatasetIAMBinding.
type DatasetIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMBinding is the Schema for the DatasetIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DatasetIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasetIAMBindingSpec   `json:"spec"`
	Status            DatasetIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMBindingList contains a list of DatasetIAMBindings
type DatasetIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasetIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	DatasetIAMBinding_Kind             = "DatasetIAMBinding"
	DatasetIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatasetIAMBinding_Kind}.String()
	DatasetIAMBinding_KindAPIVersion   = DatasetIAMBinding_Kind + "." + CRDGroupVersion.String()
	DatasetIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(DatasetIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&DatasetIAMBinding{}, &DatasetIAMBindingList{})
}
