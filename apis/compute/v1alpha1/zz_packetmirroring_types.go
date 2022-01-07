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

type CollectorIlbObservation struct {
}

type CollectorIlbParameters struct {

	// The URL of the forwarding rule.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// IP CIDR ranges that apply as a filter on the source (ingress) or
	// destination (egress) IP in the IP header. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	CidrRanges []*string `json:"cidrRanges,omitempty" tf:"cidr_ranges,omitempty"`

	// Direction of traffic to mirror. Default value: "BOTH" Possible values: ["INGRESS", "EGRESS", "BOTH"]
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Protocols that apply as a filter on mirrored traffic. Possible values: ["tcp", "udp", "icmp"]
	// +kubebuilder:validation:Optional
	IPProtocols []*string `json:"ipProtocols,omitempty" tf:"ip_protocols,omitempty"`
}

type InstancesObservation struct {
}

type InstancesParameters struct {

	// The URL of the instances where this rule should be active.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type MirroredResourcesObservation struct {
}

type MirroredResourcesParameters struct {

	// All the listed instances will be mirrored.  Specify at most 50.
	// +kubebuilder:validation:Optional
	Instances []InstancesParameters `json:"instances,omitempty" tf:"instances,omitempty"`

	// All instances in one of these subnetworks will be mirrored.
	// +kubebuilder:validation:Optional
	Subnetworks []SubnetworksParameters `json:"subnetworks,omitempty" tf:"subnetworks,omitempty"`

	// All instances with these tags will be mirrored.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// The full self_link URL of the network where this rule is active.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type PacketMirroringObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PacketMirroringParameters struct {

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have is_mirroring_collector
	// set to true.
	// +kubebuilder:validation:Required
	CollectorIlb []CollectorIlbParameters `json:"collectorIlb" tf:"collector_ilb,omitempty"`

	// A human-readable description of the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// A means of specifying which resources to mirror.
	// +kubebuilder:validation:Required
	MirroredResources []MirroredResourcesParameters `json:"mirroredResources" tf:"mirrored_resources,omitempty"`

	// The name of the packet mirroring rule
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// +kubebuilder:validation:Required
	Network []NetworkParameters `json:"network" tf:"network,omitempty"`

	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type SubnetworksObservation struct {
}

type SubnetworksParameters struct {

	// The URL of the subnetwork where this rule should be active.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// PacketMirroringSpec defines the desired state of PacketMirroring
type PacketMirroringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PacketMirroringParameters `json:"forProvider"`
}

// PacketMirroringStatus defines the observed state of PacketMirroring.
type PacketMirroringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PacketMirroringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PacketMirroring is the Schema for the PacketMirrorings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type PacketMirroring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PacketMirroringSpec   `json:"spec"`
	Status            PacketMirroringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PacketMirroringList contains a list of PacketMirrorings
type PacketMirroringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PacketMirroring `json:"items"`
}

// Repository type metadata.
var (
	PacketMirroring_Kind             = "PacketMirroring"
	PacketMirroring_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PacketMirroring_Kind}.String()
	PacketMirroring_KindAPIVersion   = PacketMirroring_Kind + "." + CRDGroupVersion.String()
	PacketMirroring_GroupVersionKind = CRDGroupVersion.WithKind(PacketMirroring_Kind)
)

func init() {
	SchemeBuilder.Register(&PacketMirroring{}, &PacketMirroringList{})
}
