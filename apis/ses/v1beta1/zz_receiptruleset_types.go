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

type ReceiptRuleSetObservation struct {

	// SES receipt rule set ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// SES receipt rule set name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReceiptRuleSetParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the rule set.
	// +kubebuilder:validation:Required
	RuleSetName *string `json:"ruleSetName" tf:"rule_set_name,omitempty"`
}

// ReceiptRuleSetSpec defines the desired state of ReceiptRuleSet
type ReceiptRuleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReceiptRuleSetParameters `json:"forProvider"`
}

// ReceiptRuleSetStatus defines the observed state of ReceiptRuleSet.
type ReceiptRuleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReceiptRuleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptRuleSet is the Schema for the ReceiptRuleSets API. Provides an SES receipt rule set resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReceiptRuleSetSpec   `json:"spec"`
	Status            ReceiptRuleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptRuleSetList contains a list of ReceiptRuleSets
type ReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReceiptRuleSet `json:"items"`
}

// Repository type metadata.
var (
	ReceiptRuleSet_Kind             = "ReceiptRuleSet"
	ReceiptRuleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReceiptRuleSet_Kind}.String()
	ReceiptRuleSet_KindAPIVersion   = ReceiptRuleSet_Kind + "." + CRDGroupVersion.String()
	ReceiptRuleSet_GroupVersionKind = CRDGroupVersion.WithKind(ReceiptRuleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ReceiptRuleSet{}, &ReceiptRuleSetList{})
}