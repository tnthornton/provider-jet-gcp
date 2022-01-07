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

type GuestAcceleratorsObservation struct {
}

type GuestAcceleratorsParameters struct {

	// The number of the guest accelerator cards exposed to
	// this instance.
	// +kubebuilder:validation:Required
	AcceleratorCount *int64 `json:"acceleratorCount" tf:"accelerator_count,omitempty"`

	// The full or partial URL of the accelerator type to
	// attach to this instance. For example:
	// 'projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100'
	//
	// If you are creating an instance template, specify only the accelerator name.
	// +kubebuilder:validation:Required
	AcceleratorType *string `json:"acceleratorType" tf:"accelerator_type,omitempty"`
}

type InstancePropertiesObservation struct {
}

type InstancePropertiesParameters struct {

	// Guest accelerator type and count.
	// +kubebuilder:validation:Optional
	GuestAccelerators []GuestAcceleratorsParameters `json:"guestAccelerators,omitempty" tf:"guest_accelerators,omitempty"`

	// The amount of local ssd to reserve with each instance. This
	// reserves disks of type 'local-ssd'.
	// +kubebuilder:validation:Optional
	LocalSsds []LocalSsdsParameters `json:"localSsds,omitempty" tf:"local_ssds,omitempty"`

	// The name of the machine type to reserve.
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// The minimum CPU platform for the reservation. For example,
	// '"Intel Skylake"'. See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`
}

type LocalSsdsObservation struct {
}

type LocalSsdsParameters struct {

	// The size of the disk in base-2 GB.
	// +kubebuilder:validation:Required
	DiskSizeGb *int64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"]
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`
}

type ReservationObservation struct {
	Commitment *string `json:"commitment,omitempty" tf:"commitment,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReservationParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reservation for instances with specific machine shapes.
	// +kubebuilder:validation:Required
	SpecificReservation []SpecificReservationParameters `json:"specificReservation" tf:"specific_reservation,omitempty"`

	// When set to true, only VMs that target this reservation by name can
	// consume this reservation. Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	// +kubebuilder:validation:Optional
	SpecificReservationRequired *bool `json:"specificReservationRequired,omitempty" tf:"specific_reservation_required,omitempty"`

	// The zone where the reservation is made.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type SpecificReservationObservation struct {
	InUseCount *int64 `json:"inUseCount,omitempty" tf:"in_use_count,omitempty"`
}

type SpecificReservationParameters struct {

	// The number of resources that are allocated.
	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// The instance properties for the reservation.
	// +kubebuilder:validation:Required
	InstanceProperties []InstancePropertiesParameters `json:"instanceProperties" tf:"instance_properties,omitempty"`
}

// ReservationSpec defines the desired state of Reservation
type ReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationParameters `json:"forProvider"`
}

// ReservationStatus defines the observed state of Reservation.
type ReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reservation is the Schema for the Reservations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Reservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReservationSpec   `json:"spec"`
	Status            ReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationList contains a list of Reservations
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reservation `json:"items"`
}

// Repository type metadata.
var (
	Reservation_Kind             = "Reservation"
	Reservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reservation_Kind}.String()
	Reservation_KindAPIVersion   = Reservation_Kind + "." + CRDGroupVersion.String()
	Reservation_GroupVersionKind = CRDGroupVersion.WithKind(Reservation_Kind)
)

func init() {
	SchemeBuilder.Register(&Reservation{}, &ReservationList{})
}
