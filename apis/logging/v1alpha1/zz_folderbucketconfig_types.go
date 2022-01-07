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

type FolderBucketConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FolderBucketConfigParameters struct {

	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	// +kubebuilder:validation:Required
	BucketID *string `json:"bucketId" tf:"bucket_id,omitempty"`

	// An optional description for this bucket.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parent resource that contains the logging bucket.
	// +kubebuilder:validation:Required
	Folder *string `json:"folder" tf:"folder,omitempty"`

	// The location of the bucket.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
}

// FolderBucketConfigSpec defines the desired state of FolderBucketConfig
type FolderBucketConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderBucketConfigParameters `json:"forProvider"`
}

// FolderBucketConfigStatus defines the observed state of FolderBucketConfig.
type FolderBucketConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderBucketConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderBucketConfig is the Schema for the FolderBucketConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FolderBucketConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderBucketConfigSpec   `json:"spec"`
	Status            FolderBucketConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderBucketConfigList contains a list of FolderBucketConfigs
type FolderBucketConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderBucketConfig `json:"items"`
}

// Repository type metadata.
var (
	FolderBucketConfig_Kind             = "FolderBucketConfig"
	FolderBucketConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderBucketConfig_Kind}.String()
	FolderBucketConfig_KindAPIVersion   = FolderBucketConfig_Kind + "." + CRDGroupVersion.String()
	FolderBucketConfig_GroupVersionKind = CRDGroupVersion.WithKind(FolderBucketConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderBucketConfig{}, &FolderBucketConfigList{})
}
