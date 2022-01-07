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

type KMSSettingsObservation struct {
}

type KMSSettingsParameters struct {

	// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
	// +kubebuilder:validation:Required
	NextRotationTime *string `json:"nextRotationTime" tf:"next_rotation_time,omitempty"`

	// Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
	// +kubebuilder:validation:Required
	RotationPeriod *string `json:"rotationPeriod" tf:"rotation_period,omitempty"`
}

type ResourceSettingsObservation struct {
}

type ResourceSettingsParameters struct {

	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type ResourcesObservation struct {
	ResourceID *int64 `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type ResourcesParameters struct {
}

type WorkloadObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`
}

type WorkloadParameters struct {

	// Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
	// +kubebuilder:validation:Required
	BillingAccount *string `json:"billingAccount" tf:"billing_account,omitempty"`

	// Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS
	// +kubebuilder:validation:Required
	ComplianceRegime *string `json:"complianceRegime" tf:"compliance_regime,omitempty"`

	// Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	// +kubebuilder:validation:Optional
	KMSSettings []KMSSettingsParameters `json:"kmsSettings,omitempty" tf:"kms_settings,omitempty"`

	// Optional. Labels applied to the workload.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The organization for the resource
	// +kubebuilder:validation:Required
	Organization *string `json:"organization" tf:"organization,omitempty"`

	// Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
	// +kubebuilder:validation:Optional
	ProvisionedResourcesParent *string `json:"provisionedResourcesParent,omitempty" tf:"provisioned_resources_parent,omitempty"`

	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	// +kubebuilder:validation:Optional
	ResourceSettings []ResourceSettingsParameters `json:"resourceSettings,omitempty" tf:"resource_settings,omitempty"`
}

// WorkloadSpec defines the desired state of Workload
type WorkloadSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkloadParameters `json:"forProvider"`
}

// WorkloadStatus defines the observed state of Workload.
type WorkloadStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkloadObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workload is the Schema for the Workloads API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Workload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkloadSpec   `json:"spec"`
	Status            WorkloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkloadList contains a list of Workloads
type WorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workload `json:"items"`
}

// Repository type metadata.
var (
	Workload_Kind             = "Workload"
	Workload_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workload_Kind}.String()
	Workload_KindAPIVersion   = Workload_Kind + "." + CRDGroupVersion.String()
	Workload_GroupVersionKind = CRDGroupVersion.WithKind(Workload_Kind)
)

func init() {
	SchemeBuilder.Register(&Workload{}, &WorkloadList{})
}
