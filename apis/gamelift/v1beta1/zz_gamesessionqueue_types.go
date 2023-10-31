// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GameSessionQueueInitParameters struct {

	// Information to be added to all events that are related to this game session queue.
	CustomEventData *string `json:"customEventData,omitempty" tf:"custom_event_data,omitempty"`

	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicy []PlayerLatencyPolicyInitParameters `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`
}

type GameSessionQueueObservation struct {

	// Game Session Queue ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Information to be added to all events that are related to this game session queue.
	CustomEventData *string `json:"customEventData,omitempty" tf:"custom_event_data,omitempty"`

	// List of fleet/alias ARNs used by session queue for placing game sessions.
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An SNS topic ARN that is set up to receive game session placement notifications.
	NotificationTarget *string `json:"notificationTarget,omitempty" tf:"notification_target,omitempty"`

	// One or more policies used to choose fleet based on player latency. See below.
	PlayerLatencyPolicy []PlayerLatencyPolicyObservation `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Maximum time a game session request can remain in the queue.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`
}

type GameSessionQueueParameters struct {

	// Information to be added to all events that are related to this game session queue.
	// +kubebuilder:validation:Optional
	CustomEventData *string `json:"customEventData,omitempty" tf:"custom_event_data,omitempty"`

	// List of fleet/alias ARNs used by session queue for placing game sessions.
	// +kubebuilder:validation:Optional
	Destinations []*string `json:"destinations,omitempty" tf:"destinations,omitempty"`

	// An SNS topic ARN that is set up to receive game session placement notifications.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	NotificationTarget *string `json:"notificationTarget,omitempty" tf:"notification_target,omitempty"`

	// Reference to a Topic in sns to populate notificationTarget.
	// +kubebuilder:validation:Optional
	NotificationTargetRef *v1.Reference `json:"notificationTargetRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate notificationTarget.
	// +kubebuilder:validation:Optional
	NotificationTargetSelector *v1.Selector `json:"notificationTargetSelector,omitempty" tf:"-"`

	// One or more policies used to choose fleet based on player latency. See below.
	// +kubebuilder:validation:Optional
	PlayerLatencyPolicy []PlayerLatencyPolicyParameters `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum time a game session request can remain in the queue.
	// +kubebuilder:validation:Optional
	TimeoutInSeconds *float64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`
}

type PlayerLatencyPolicyInitParameters struct {

	// Maximum latency value that is allowed for any player.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `json:"maximumIndividualPlayerLatencyMilliseconds,omitempty" tf:"maximum_individual_player_latency_milliseconds,omitempty"`

	// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
	PolicyDurationSeconds *float64 `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds,omitempty"`
}

type PlayerLatencyPolicyObservation struct {

	// Maximum latency value that is allowed for any player.
	MaximumIndividualPlayerLatencyMilliseconds *float64 `json:"maximumIndividualPlayerLatencyMilliseconds,omitempty" tf:"maximum_individual_player_latency_milliseconds,omitempty"`

	// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
	PolicyDurationSeconds *float64 `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds,omitempty"`
}

type PlayerLatencyPolicyParameters struct {

	// Maximum latency value that is allowed for any player.
	// +kubebuilder:validation:Optional
	MaximumIndividualPlayerLatencyMilliseconds *float64 `json:"maximumIndividualPlayerLatencyMilliseconds" tf:"maximum_individual_player_latency_milliseconds,omitempty"`

	// Length of time that the policy is enforced while placing a new game session. Absence of value for this attribute means that the policy is enforced until the queue times out.
	// +kubebuilder:validation:Optional
	PolicyDurationSeconds *float64 `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds,omitempty"`
}

// GameSessionQueueSpec defines the desired state of GameSessionQueue
type GameSessionQueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GameSessionQueueParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GameSessionQueueInitParameters `json:"initProvider,omitempty"`
}

// GameSessionQueueStatus defines the observed state of GameSessionQueue.
type GameSessionQueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GameSessionQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameSessionQueue is the Schema for the GameSessionQueues API. Provides a GameLift Game Session Queue resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GameSessionQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameSessionQueueSpec   `json:"spec"`
	Status            GameSessionQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameSessionQueueList contains a list of GameSessionQueues
type GameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameSessionQueue `json:"items"`
}

// Repository type metadata.
var (
	GameSessionQueue_Kind             = "GameSessionQueue"
	GameSessionQueue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GameSessionQueue_Kind}.String()
	GameSessionQueue_KindAPIVersion   = GameSessionQueue_Kind + "." + CRDGroupVersion.String()
	GameSessionQueue_GroupVersionKind = CRDGroupVersion.WithKind(GameSessionQueue_Kind)
)

func init() {
	SchemeBuilder.Register(&GameSessionQueue{}, &GameSessionQueueList{})
}
