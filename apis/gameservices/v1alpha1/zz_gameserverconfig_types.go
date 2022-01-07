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

type FleetConfigsObservation struct {
}

type FleetConfigsParameters struct {

	// The fleet spec, which is sent to Agones to configure fleet.
	// The spec can be passed as inline json but it is recommended to use a file reference
	// instead. File references can contain the json or yaml format of the fleet spec. Eg:
	//
	// * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
	// * fleet_spec = file("fleet_configs.json")
	//
	// The format of the spec can be found :
	// 'https://agones.dev/site/docs/reference/fleet/'.
	// +kubebuilder:validation:Required
	FleetSpec *string `json:"fleetSpec" tf:"fleet_spec,omitempty"`

	// The name of the FleetConfig.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GameServerConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GameServerConfigParameters struct {

	// A unique id for the deployment config.
	// +kubebuilder:validation:Required
	ConfigID *string `json:"configId" tf:"config_id,omitempty"`

	// A unique id for the deployment.
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// The description of the game server config.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The fleet config contains list of fleet specs. In the Single Cloud, there
	// will be only one.
	// +kubebuilder:validation:Required
	FleetConfigs []FleetConfigsParameters `json:"fleetConfigs" tf:"fleet_configs,omitempty"`

	// The labels associated with this game server config. Each label is a
	// key-value pair.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location of the Deployment.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. This contains the autoscaling settings.
	// +kubebuilder:validation:Optional
	ScalingConfigs []ScalingConfigsParameters `json:"scalingConfigs,omitempty" tf:"scaling_configs,omitempty"`
}

type ScalingConfigsObservation struct {
}

type ScalingConfigsParameters struct {

	// Fleet autoscaler spec, which is sent to Agones.
	// Example spec can be found :
	// https://agones.dev/site/docs/reference/fleetautoscaler/
	// +kubebuilder:validation:Required
	FleetAutoscalerSpec *string `json:"fleetAutoscalerSpec" tf:"fleet_autoscaler_spec,omitempty"`

	// The name of the ScalingConfig
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The schedules to which this scaling config applies.
	// +kubebuilder:validation:Optional
	Schedules []SchedulesParameters `json:"schedules,omitempty" tf:"schedules,omitempty"`

	// Labels used to identify the clusters to which this scaling config
	// applies. A cluster is subject to this scaling config if its labels match
	// any of the selector entries.
	// +kubebuilder:validation:Optional
	Selectors []SelectorsParameters `json:"selectors,omitempty" tf:"selectors,omitempty"`
}

type SchedulesObservation struct {
}

type SchedulesParameters struct {

	// The duration for the cron job event. The duration of the event is effective
	// after the cron job's start time.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	CronJobDuration *string `json:"cronJobDuration,omitempty" tf:"cron_job_duration,omitempty"`

	// The cron definition of the scheduled event. See
	// https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
	// defined by the realm.
	// +kubebuilder:validation:Optional
	CronSpec *string `json:"cronSpec,omitempty" tf:"cron_spec,omitempty"`

	// The end time of the event.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The start time of the event.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type SelectorsObservation struct {
}

type SelectorsParameters struct {

	// Set of labels to group by.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

// GameServerConfigSpec defines the desired state of GameServerConfig
type GameServerConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GameServerConfigParameters `json:"forProvider"`
}

// GameServerConfigStatus defines the observed state of GameServerConfig.
type GameServerConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GameServerConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerConfig is the Schema for the GameServerConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type GameServerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameServerConfigSpec   `json:"spec"`
	Status            GameServerConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerConfigList contains a list of GameServerConfigs
type GameServerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameServerConfig `json:"items"`
}

// Repository type metadata.
var (
	GameServerConfig_Kind             = "GameServerConfig"
	GameServerConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GameServerConfig_Kind}.String()
	GameServerConfig_KindAPIVersion   = GameServerConfig_Kind + "." + CRDGroupVersion.String()
	GameServerConfig_GroupVersionKind = CRDGroupVersion.WithKind(GameServerConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&GameServerConfig{}, &GameServerConfigList{})
}
