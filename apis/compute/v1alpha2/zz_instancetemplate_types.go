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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskEncryptionKeyObservation struct {
}

type DiskEncryptionKeyParameters struct {

	// The self link of the encryption key that is stored in Google Cloud KMS.
	// +kubebuilder:validation:Required
	KMSKeySelfLink *string `json:"kmsKeySelfLink" tf:"kms_key_self_link,omitempty"`
}

type DiskObservation struct {
}

type DiskParameters struct {

	// Whether or not the disk should be auto-deleted. This defaults to true.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Indicates that this is a boot disk.
	// +kubebuilder:validation:Optional
	Boot *bool `json:"boot,omitempty" tf:"boot,omitempty"`

	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Encrypts or decrypts a disk using a customer-supplied encryption key.
	// +kubebuilder:validation:Optional
	DiskEncryptionKey []DiskEncryptionKeyParameters `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`

	// Name of the disk. When not provided, this defaults to the name of the instance.
	// +kubebuilder:validation:Optional
	DiskName *string `json:"diskName,omitempty" tf:"disk_name,omitempty"`

	// The size of the image in gigabytes. If not specified, it will inherit the size of its base image. For SCRATCH disks, the size must be exactly 375GB.
	// +kubebuilder:validation:Optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// The Google Compute Engine disk type. Can be either "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard".
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// Specifies the disk interface to use for attaching this disk.
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// A set of key/value label pairs to assign to disks,
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// A list (short name or id) of resource policies to attach to this disk. Currently a max of 1 resource policy is supported.
	// +kubebuilder:validation:Optional
	ResourcePolicies []*string `json:"resourcePolicies,omitempty" tf:"resource_policies,omitempty"`

	// The name (not self_link) of the disk (such as those managed by google_compute_disk) to attach. ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The image from which to initialize this disk. This can be one of: the image's self_link, projects/{project}/global/images/{image}, projects/{project}/global/images/family/{family}, global/images/{image}, global/images/family/{family}, family/{family}, {project}/{family}, {project}/{image}, {family}, or {image}. ~> Note: Either source or source_image is required when creating a new instance except for when creating a local SSD.
	// +kubebuilder:validation:Optional
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// The type of Google Compute Engine disk, can be either "SCRATCH" or "PERSISTENT".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceTemplateAdvancedMachineFeaturesObservation struct {
}

type InstanceTemplateAdvancedMachineFeaturesParameters struct {

	// Whether to enable nested virtualization or not.
	// +kubebuilder:validation:Optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty" tf:"enable_nested_virtualization,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kubebuilder:validation:Optional
	ThreadsPerCore *int64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type InstanceTemplateConfidentialInstanceConfigObservation struct {
}

type InstanceTemplateConfidentialInstanceConfigParameters struct {

	// Defines whether the instance should have confidential compute enabled.
	// +kubebuilder:validation:Required
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute,omitempty"`
}

type InstanceTemplateGuestAcceleratorObservation struct {
}

type InstanceTemplateGuestAcceleratorParameters struct {

	// The number of the guest accelerator cards exposed to this instance.
	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceTemplateNetworkInterfaceAccessConfigObservation struct {
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type InstanceTemplateNetworkInterfaceAccessConfigParameters struct {

	// The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.
	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip,omitempty"`

	// The networking tier used for configuring this instance template. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM.
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`
}

type InstanceTemplateNetworkInterfaceAliasIPRangeObservation struct {
}

type InstanceTemplateNetworkInterfaceAliasIPRangeParameters struct {

	// The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.
	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used.
	// +kubebuilder:validation:Optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type InstanceTemplateNetworkInterfaceIPv6AccessConfigObservation struct {
	ExternalIPv6 *string `json:"externalIpv6,omitempty" tf:"external_ipv6,omitempty"`

	ExternalIPv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty" tf:"external_ipv6_prefix_length,omitempty"`

	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type InstanceTemplateNetworkInterfaceIPv6AccessConfigParameters struct {

	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6
	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`
}

type InstanceTemplateNetworkInterfaceObservation struct {
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InstanceTemplateNetworkInterfaceParameters struct {

	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet (this means that ssh provisioners will not work unless you are running Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times.
	// +kubebuilder:validation:Optional
	AccessConfig []InstanceTemplateNetworkInterfaceAccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks.
	// +kubebuilder:validation:Optional
	AliasIPRange []InstanceTemplateNetworkInterfaceAliasIPRangeParameters `json:"aliasIpRange,omitempty" tf:"alias_ip_range,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []InstanceTemplateNetworkInterfaceIPv6AccessConfigParameters `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// The name or self_link of the network to attach this interface to. Use network attribute for Legacy or Auto subnetted networks and subnetwork for custom subnetted networks.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The private IP address to assign to the instance. If empty, the address will be automatically assigned.
	// +kubebuilder:validation:Optional
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// The name of the subnetwork to attach this interface to. The subnetwork must exist in the same region this instance will be created in. Either network or subnetwork must be provided.
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`
}

type InstanceTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type InstanceTemplateParameters struct {

	// Controls for advanced machine-related behavior features.
	// +kubebuilder:validation:Optional
	AdvancedMachineFeatures []InstanceTemplateAdvancedMachineFeaturesParameters `json:"advancedMachineFeatures,omitempty" tf:"advanced_machine_features,omitempty"`

	// Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.
	// +kubebuilder:validation:Optional
	CanIPForward *bool `json:"canIpForward,omitempty" tf:"can_ip_forward,omitempty"`

	// The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail to create.
	// +kubebuilder:validation:Optional
	ConfidentialInstanceConfig []InstanceTemplateConfidentialInstanceConfigParameters `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config,omitempty"`

	// A brief description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disks to attach to instances created from this template. This can be specified multiple times for multiple disks.
	// +kubebuilder:validation:Required
	Disk []DiskParameters `json:"disk" tf:"disk,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	// +kubebuilder:validation:Optional
	GuestAccelerator []InstanceTemplateGuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// A description of the instance.
	// +kubebuilder:validation:Optional
	InstanceDescription *string `json:"instanceDescription,omitempty" tf:"instance_description,omitempty"`

	// A set of key/value label pairs to assign to instances created from this template,
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to create. To create a machine with a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB like custom-6-20480 for 6 vCPU and 20GB of RAM.
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// Metadata key/value pairs to make available from within instances created from this template.
	// +kubebuilder:validation:Optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.
	// +kubebuilder:validation:Optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`

	// Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell or Intel Skylake.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// Networks to attach to instances created from this template. This can be specified multiple times for multiple networks.
	// +kubebuilder:validation:Optional
	NetworkInterface []InstanceTemplateNetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom subnetwork resource is tied to a specific region. Defaults to the region of the Provider if no value is given.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the reservations that this instance can consume from.
	// +kubebuilder:validation:Optional
	ReservationAffinity []InstanceTemplateReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// The scheduling strategy to use.
	// +kubebuilder:validation:Optional
	Scheduling []InstanceTemplateSchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// Service account to attach to the instance.
	// +kubebuilder:validation:Optional
	ServiceAccount []InstanceTemplateServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Enable Shielded VM on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Note: shielded_instance_config can only be used with boot images with shielded vm support.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []InstanceTemplateShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// Tags to attach to the instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InstanceTemplateReservationAffinityObservation struct {
}

type InstanceTemplateReservationAffinityParameters struct {

	// Specifies the label selector for the reservation to use.
	// +kubebuilder:validation:Optional
	SpecificReservation []InstanceTemplateReservationAffinitySpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceTemplateReservationAffinitySpecificReservationObservation struct {
}

type InstanceTemplateReservationAffinitySpecificReservationParameters struct {

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Corresponds to the label values of a reservation resource.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type InstanceTemplateSchedulingNodeAffinitiesObservation struct {
}

type InstanceTemplateSchedulingNodeAffinitiesParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type InstanceTemplateSchedulingObservation struct {
}

type InstanceTemplateSchedulingParameters struct {

	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true.
	// +kubebuilder:validation:Optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`

	// Minimum number of cpus for the instance.
	// +kubebuilder:validation:Optional
	MinNodeCpus *int64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus,omitempty"`

	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	// +kubebuilder:validation:Optional
	NodeAffinities []InstanceTemplateSchedulingNodeAffinitiesParameters `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`

	// Defines the maintenance behavior for this instance.
	// +kubebuilder:validation:Optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`

	// Allows instance to be preempted. This defaults to false.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type InstanceTemplateServiceAccountObservation struct {
}

type InstanceTemplateServiceAccountParameters struct {

	// The service account e-mail address. If not given, the default Google Compute Engine service account is used.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the cloud-platform scope.
	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`
}

type InstanceTemplateShieldedInstanceConfigObservation struct {
}

type InstanceTemplateShieldedInstanceConfigParameters struct {

	// Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// Verify the digital signature of all boot components, and halt the boot process if signature verification fails. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

// InstanceTemplateSpec defines the desired state of InstanceTemplate
type InstanceTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceTemplateParameters `json:"forProvider"`
}

// InstanceTemplateStatus defines the observed state of InstanceTemplate.
type InstanceTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceTemplate is the Schema for the InstanceTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type InstanceTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceTemplateSpec   `json:"spec"`
	Status            InstanceTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceTemplateList contains a list of InstanceTemplates
type InstanceTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceTemplate `json:"items"`
}

// Repository type metadata.
var (
	InstanceTemplate_Kind             = "InstanceTemplate"
	InstanceTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceTemplate_Kind}.String()
	InstanceTemplate_KindAPIVersion   = InstanceTemplate_Kind + "." + CRDGroupVersion.String()
	InstanceTemplate_GroupVersionKind = CRDGroupVersion.WithKind(InstanceTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceTemplate{}, &InstanceTemplateList{})
}
