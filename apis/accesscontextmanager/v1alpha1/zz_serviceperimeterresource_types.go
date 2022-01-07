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

type ServicePerimeterResourceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServicePerimeterResourceParameters struct {

	// The name of the Service Perimeter to add this resource to.
	// +kubebuilder:validation:Required
	PerimeterName *string `json:"perimeterName" tf:"perimeter_name,omitempty"`

	// A GCP resource that is inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	// +kubebuilder:validation:Required
	Resource *string `json:"resource" tf:"resource,omitempty"`
}

// ServicePerimeterResourceSpec defines the desired state of ServicePerimeterResource
type ServicePerimeterResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicePerimeterResourceParameters `json:"forProvider"`
}

// ServicePerimeterResourceStatus defines the observed state of ServicePerimeterResource.
type ServicePerimeterResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicePerimeterResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicePerimeterResource is the Schema for the ServicePerimeterResources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ServicePerimeterResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicePerimeterResourceSpec   `json:"spec"`
	Status            ServicePerimeterResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicePerimeterResourceList contains a list of ServicePerimeterResources
type ServicePerimeterResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePerimeterResource `json:"items"`
}

// Repository type metadata.
var (
	ServicePerimeterResource_Kind             = "ServicePerimeterResource"
	ServicePerimeterResource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicePerimeterResource_Kind}.String()
	ServicePerimeterResource_KindAPIVersion   = ServicePerimeterResource_Kind + "." + CRDGroupVersion.String()
	ServicePerimeterResource_GroupVersionKind = CRDGroupVersion.WithKind(ServicePerimeterResource_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicePerimeterResource{}, &ServicePerimeterResourceList{})
}
