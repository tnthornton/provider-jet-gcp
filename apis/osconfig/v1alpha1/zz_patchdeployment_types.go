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

type AptObservation struct {
}

type AptParameters struct {

	// List of packages to exclude from update. These packages will be excluded.
	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	// +kubebuilder:validation:Optional
	ExclusivePackages []*string `json:"exclusivePackages,omitempty" tf:"exclusive_packages,omitempty"`

	// By changing the type to DIST, the patching is performed using apt-get dist-upgrade instead. Possible values: ["DIST", "UPGRADE"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DisruptionBudgetObservation struct {
}

type DisruptionBudgetParameters struct {

	// Specifies a fixed value.
	// +kubebuilder:validation:Optional
	Fixed *int64 `json:"fixed,omitempty" tf:"fixed,omitempty"`

	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	// +kubebuilder:validation:Optional
	Percentage *int64 `json:"percentage,omitempty" tf:"percentage,omitempty"`
}

type GcsObjectObservation struct {
}

type GcsObjectParameters struct {

	// Bucket of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	// +kubebuilder:validation:Required
	GenerationNumber *string `json:"generationNumber" tf:"generation_number,omitempty"`

	// Name of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Object *string `json:"object" tf:"object,omitempty"`
}

type GooObservation struct {
}

type GooParameters struct {

	// goo update settings. Use this setting to override the default goo patch rules.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type GroupLabelsObservation struct {
}

type GroupLabelsParameters struct {

	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter
	// +kubebuilder:validation:Required
	Labels map[string]*string `json:"labels" tf:"labels,omitempty"`
}

type InstanceFilterObservation struct {
}

type InstanceFilterParameters struct {

	// Target all VM instances in the project. If true, no other criteria is permitted.
	// +kubebuilder:validation:Optional
	All *bool `json:"all,omitempty" tf:"all,omitempty"`

	// Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.
	// +kubebuilder:validation:Optional
	GroupLabels []GroupLabelsParameters `json:"groupLabels,omitempty" tf:"group_labels,omitempty"`

	// Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
	// VMs when targeting configs, for example prefix="prod-".
	// +kubebuilder:validation:Optional
	InstanceNamePrefixes []*string `json:"instanceNamePrefixes,omitempty" tf:"instance_name_prefixes,omitempty"`

	// Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}',
	// 'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or
	// 'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'
	// +kubebuilder:validation:Optional
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LinuxExecStepConfigGcsObjectObservation struct {
}

type LinuxExecStepConfigGcsObjectParameters struct {

	// Bucket of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	// +kubebuilder:validation:Required
	GenerationNumber *string `json:"generationNumber" tf:"generation_number,omitempty"`

	// Name of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Object *string `json:"object" tf:"object,omitempty"`
}

type LinuxExecStepConfigObservation struct {
}

type LinuxExecStepConfigParameters struct {

	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +kubebuilder:validation:Optional
	AllowedSuccessCodes []*int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes,omitempty"`

	// A Cloud Storage object containing the executable.
	// +kubebuilder:validation:Optional
	GcsObject []GcsObjectParameters `json:"gcsObject,omitempty" tf:"gcs_object,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +kubebuilder:validation:Optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter,omitempty"`

	// An absolute path to the executable on the VM.
	// +kubebuilder:validation:Optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path,omitempty"`
}

type MonthlyObservation struct {
}

type MonthlyParameters struct {

	// One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
	// Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
	// will not run in February, April, June, etc.
	// +kubebuilder:validation:Optional
	MonthDay *int64 `json:"monthDay,omitempty" tf:"month_day,omitempty"`

	// Week day in a month.
	// +kubebuilder:validation:Optional
	WeekDayOfMonth []WeekDayOfMonthParameters `json:"weekDayOfMonth,omitempty" tf:"week_day_of_month,omitempty"`
}

type OneTimeScheduleObservation struct {
}

type OneTimeScheduleParameters struct {

	// The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Required
	ExecuteTime *string `json:"executeTime" tf:"execute_time,omitempty"`
}

type PatchConfigObservation struct {
}

type PatchConfigParameters struct {

	// Apt update settings. Use this setting to override the default apt patch rules.
	// +kubebuilder:validation:Optional
	Apt []AptParameters `json:"apt,omitempty" tf:"apt,omitempty"`

	// goo update settings. Use this setting to override the default goo patch rules.
	// +kubebuilder:validation:Optional
	Goo []GooParameters `json:"goo,omitempty" tf:"goo,omitempty"`

	// The ExecStep to run after the patch update.
	// +kubebuilder:validation:Optional
	PostStep []PostStepParameters `json:"postStep,omitempty" tf:"post_step,omitempty"`

	// The ExecStep to run before the patch update.
	// +kubebuilder:validation:Optional
	PreStep []PreStepParameters `json:"preStep,omitempty" tf:"pre_step,omitempty"`

	// Post-patch reboot settings. Possible values: ["DEFAULT", "ALWAYS", "NEVER"]
	// +kubebuilder:validation:Optional
	RebootConfig *string `json:"rebootConfig,omitempty" tf:"reboot_config,omitempty"`

	// Windows update settings. Use this setting to override the default Windows patch rules.
	// +kubebuilder:validation:Optional
	WindowsUpdate []WindowsUpdateParameters `json:"windowsUpdate,omitempty" tf:"windows_update,omitempty"`

	// Yum update settings. Use this setting to override the default yum patch rules.
	// +kubebuilder:validation:Optional
	Yum []YumParameters `json:"yum,omitempty" tf:"yum,omitempty"`

	// zypper update settings. Use this setting to override the default zypper patch rules.
	// +kubebuilder:validation:Optional
	Zypper []ZypperParameters `json:"zypper,omitempty" tf:"zypper,omitempty"`
}

type PatchDeploymentObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastExecuteTime *string `json:"lastExecuteTime,omitempty" tf:"last_execute_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type PatchDeploymentParameters struct {

	// Description of the patch deployment. Length of the description is limited to 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Duration of the patch. After the duration ends, the patch times out.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// VM instances to patch.
	// +kubebuilder:validation:Required
	InstanceFilter []InstanceFilterParameters `json:"instanceFilter" tf:"instance_filter,omitempty"`

	// Schedule a one-time execution.
	// +kubebuilder:validation:Optional
	OneTimeSchedule []OneTimeScheduleParameters `json:"oneTimeSchedule,omitempty" tf:"one_time_schedule,omitempty"`

	// Patch configuration that is applied.
	// +kubebuilder:validation:Optional
	PatchConfig []PatchConfigParameters `json:"patchConfig,omitempty" tf:"patch_config,omitempty"`

	// A name for the patch deployment in the project. When creating a name the following rules apply:
	// * Must contain only lowercase letters, numbers, and hyphens.
	// * Must start with a letter.
	// * Must be between 1-63 characters.
	// * Must end with a number or a letter.
	// * Must be unique within the project.
	// +kubebuilder:validation:Required
	PatchDeploymentID *string `json:"patchDeploymentId" tf:"patch_deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Schedule recurring executions.
	// +kubebuilder:validation:Optional
	RecurringSchedule []RecurringScheduleParameters `json:"recurringSchedule,omitempty" tf:"recurring_schedule,omitempty"`

	// Rollout strategy of the patch job.
	// +kubebuilder:validation:Optional
	Rollout []RolloutParameters `json:"rollout,omitempty" tf:"rollout,omitempty"`
}

type PostStepObservation struct {
}

type PostStepParameters struct {

	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +kubebuilder:validation:Optional
	LinuxExecStepConfig []LinuxExecStepConfigParameters `json:"linuxExecStepConfig,omitempty" tf:"linux_exec_step_config,omitempty"`

	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +kubebuilder:validation:Optional
	WindowsExecStepConfig []WindowsExecStepConfigParameters `json:"windowsExecStepConfig,omitempty" tf:"windows_exec_step_config,omitempty"`
}

type PreStepLinuxExecStepConfigObservation struct {
}

type PreStepLinuxExecStepConfigParameters struct {

	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +kubebuilder:validation:Optional
	AllowedSuccessCodes []*int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes,omitempty"`

	// A Cloud Storage object containing the executable.
	// +kubebuilder:validation:Optional
	GcsObject []LinuxExecStepConfigGcsObjectParameters `json:"gcsObject,omitempty" tf:"gcs_object,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +kubebuilder:validation:Optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter,omitempty"`

	// An absolute path to the executable on the VM.
	// +kubebuilder:validation:Optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path,omitempty"`
}

type PreStepObservation struct {
}

type PreStepParameters struct {

	// The ExecStepConfig for all Linux VMs targeted by the PatchJob.
	// +kubebuilder:validation:Optional
	LinuxExecStepConfig []PreStepLinuxExecStepConfigParameters `json:"linuxExecStepConfig,omitempty" tf:"linux_exec_step_config,omitempty"`

	// The ExecStepConfig for all Windows VMs targeted by the PatchJob.
	// +kubebuilder:validation:Optional
	WindowsExecStepConfig []PreStepWindowsExecStepConfigParameters `json:"windowsExecStepConfig,omitempty" tf:"windows_exec_step_config,omitempty"`
}

type PreStepWindowsExecStepConfigGcsObjectObservation struct {
}

type PreStepWindowsExecStepConfigGcsObjectParameters struct {

	// Bucket of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	// +kubebuilder:validation:Required
	GenerationNumber *string `json:"generationNumber" tf:"generation_number,omitempty"`

	// Name of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Object *string `json:"object" tf:"object,omitempty"`
}

type PreStepWindowsExecStepConfigObservation struct {
}

type PreStepWindowsExecStepConfigParameters struct {

	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +kubebuilder:validation:Optional
	AllowedSuccessCodes []*int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes,omitempty"`

	// A Cloud Storage object containing the executable.
	// +kubebuilder:validation:Optional
	GcsObject []PreStepWindowsExecStepConfigGcsObjectParameters `json:"gcsObject,omitempty" tf:"gcs_object,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +kubebuilder:validation:Optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter,omitempty"`

	// An absolute path to the executable on the VM.
	// +kubebuilder:validation:Optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path,omitempty"`
}

type RecurringScheduleObservation struct {
	LastExecuteTime *string `json:"lastExecuteTime,omitempty" tf:"last_execute_time,omitempty"`

	NextExecuteTime *string `json:"nextExecuteTime,omitempty" tf:"next_execute_time,omitempty"`
}

type RecurringScheduleParameters struct {

	// The end time at which a recurring patch deployment schedule is no longer active.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Schedule with monthly executions.
	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Time of the day to run a recurring deployment.
	// +kubebuilder:validation:Required
	TimeOfDay []TimeOfDayParameters `json:"timeOfDay" tf:"time_of_day,omitempty"`

	// Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
	// determined by the chosen time zone.
	// +kubebuilder:validation:Required
	TimeZone []TimeZoneParameters `json:"timeZone" tf:"time_zone,omitempty"`

	// Schedule with weekly executions.
	// +kubebuilder:validation:Optional
	Weekly []WeeklyParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type RolloutObservation struct {
}

type RolloutParameters struct {

	// The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up.
	// During patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps.
	// A VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget.
	// For zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone.
	// For example, if the disruption budget has a fixed value of 10, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops.
	// +kubebuilder:validation:Required
	DisruptionBudget []DisruptionBudgetParameters `json:"disruptionBudget" tf:"disruption_budget,omitempty"`

	// Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE", "CONCURRENT_ZONES"]
	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type TimeOfDayObservation struct {
}

type TimeOfDayParameters struct {

	// Hours of day in 24 hour format. Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	// +kubebuilder:validation:Optional
	Hours *int64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Minutes of hour of day. Must be from 0 to 59.
	// +kubebuilder:validation:Optional
	Minutes *int64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	// +kubebuilder:validation:Optional
	Nanos *int64 `json:"nanos,omitempty" tf:"nanos,omitempty"`

	// Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	// +kubebuilder:validation:Optional
	Seconds *int64 `json:"seconds,omitempty" tf:"seconds,omitempty"`
}

type TimeZoneObservation struct {
}

type TimeZoneParameters struct {

	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// IANA Time Zone Database version number, e.g. "2019a".
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type WeekDayOfMonthObservation struct {
}

type WeekDayOfMonthParameters struct {

	// A day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
	// +kubebuilder:validation:Required
	WeekOrdinal *int64 `json:"weekOrdinal" tf:"week_ordinal,omitempty"`
}

type WeeklyObservation struct {
}

type WeeklyParameters struct {

	// IANA Time Zone Database time zone, e.g. "America/New_York". Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`
}

type WindowsExecStepConfigGcsObjectObservation struct {
}

type WindowsExecStepConfigGcsObjectParameters struct {

	// Bucket of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.
	// +kubebuilder:validation:Required
	GenerationNumber *string `json:"generationNumber" tf:"generation_number,omitempty"`

	// Name of the Cloud Storage object.
	// +kubebuilder:validation:Required
	Object *string `json:"object" tf:"object,omitempty"`
}

type WindowsExecStepConfigObservation struct {
}

type WindowsExecStepConfigParameters struct {

	// Defaults to [0]. A list of possible return values that the execution can return to indicate a success.
	// +kubebuilder:validation:Optional
	AllowedSuccessCodes []*int64 `json:"allowedSuccessCodes,omitempty" tf:"allowed_success_codes,omitempty"`

	// A Cloud Storage object containing the executable.
	// +kubebuilder:validation:Optional
	GcsObject []WindowsExecStepConfigGcsObjectParameters `json:"gcsObject,omitempty" tf:"gcs_object,omitempty"`

	// The script interpreter to use to run the script. If no interpreter is specified the script will
	// be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"]
	// +kubebuilder:validation:Optional
	Interpreter *string `json:"interpreter,omitempty" tf:"interpreter,omitempty"`

	// An absolute path to the executable on the VM.
	// +kubebuilder:validation:Optional
	LocalPath *string `json:"localPath,omitempty" tf:"local_path,omitempty"`
}

type WindowsUpdateObservation struct {
}

type WindowsUpdateParameters struct {

	// Only apply updates of these windows update classifications. If empty, all updates are applied. Possible values: ["CRITICAL", "SECURITY", "DEFINITION", "DRIVER", "FEATURE_PACK", "SERVICE_PACK", "TOOL", "UPDATE_ROLLUP", "UPDATE"]
	// +kubebuilder:validation:Optional
	Classifications []*string `json:"classifications,omitempty" tf:"classifications,omitempty"`

	// List of KBs to exclude from update.
	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// An exclusive list of kbs to be updated. These are the only patches that will be updated.
	// This field must not be used with other patch configurations.
	// +kubebuilder:validation:Optional
	ExclusivePatches []*string `json:"exclusivePatches,omitempty" tf:"exclusive_patches,omitempty"`
}

type YumObservation struct {
}

type YumParameters struct {

	// List of packages to exclude from update. These packages will be excluded.
	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// An exclusive list of packages to be updated. These are the only packages that will be updated.
	// If these packages are not installed, they will be ignored. This field cannot be specified with
	// any other patch configuration fields.
	// +kubebuilder:validation:Optional
	ExclusivePackages []*string `json:"exclusivePackages,omitempty" tf:"exclusive_packages,omitempty"`

	// Will cause patch to run yum update-minimal instead.
	// +kubebuilder:validation:Optional
	Minimal *bool `json:"minimal,omitempty" tf:"minimal,omitempty"`

	// Adds the --security flag to yum update. Not supported on all platforms.
	// +kubebuilder:validation:Optional
	Security *bool `json:"security,omitempty" tf:"security,omitempty"`
}

type ZypperObservation struct {
}

type ZypperParameters struct {

	// Install only patches with these categories. Common categories include security, recommended, and feature.
	// +kubebuilder:validation:Optional
	Categories []*string `json:"categories,omitempty" tf:"categories,omitempty"`

	// List of packages to exclude from update.
	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.
	// This field must not be used with any other patch configuration fields.
	// +kubebuilder:validation:Optional
	ExclusivePatches []*string `json:"exclusivePatches,omitempty" tf:"exclusive_patches,omitempty"`

	// Install only patches with these severities. Common severities include critical, important, moderate, and low.
	// +kubebuilder:validation:Optional
	Severities []*string `json:"severities,omitempty" tf:"severities,omitempty"`

	// Adds the --with-optional flag to zypper patch.
	// +kubebuilder:validation:Optional
	WithOptional *bool `json:"withOptional,omitempty" tf:"with_optional,omitempty"`

	// Adds the --with-update flag, to zypper patch.
	// +kubebuilder:validation:Optional
	WithUpdate *bool `json:"withUpdate,omitempty" tf:"with_update,omitempty"`
}

// PatchDeploymentSpec defines the desired state of PatchDeployment
type PatchDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PatchDeploymentParameters `json:"forProvider"`
}

// PatchDeploymentStatus defines the observed state of PatchDeployment.
type PatchDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PatchDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PatchDeployment is the Schema for the PatchDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type PatchDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PatchDeploymentSpec   `json:"spec"`
	Status            PatchDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PatchDeploymentList contains a list of PatchDeployments
type PatchDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PatchDeployment `json:"items"`
}

// Repository type metadata.
var (
	PatchDeployment_Kind             = "PatchDeployment"
	PatchDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PatchDeployment_Kind}.String()
	PatchDeployment_KindAPIVersion   = PatchDeployment_Kind + "." + CRDGroupVersion.String()
	PatchDeployment_GroupVersionKind = CRDGroupVersion.WithKind(PatchDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&PatchDeployment{}, &PatchDeploymentList{})
}
