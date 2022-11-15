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

type TestGridProjectObservation struct {

	// The Amazon Resource Name of this Test Grid Project.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TestGridProjectParameters struct {

	// Human-readable description of the project.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the Selenium testing project.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC security groups and subnets that are attached to a project. See VPC Config below.
	// +kubebuilder:validation:Optional
	VPCConfig []VPCConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type VPCConfigObservation struct {
}

type VPCConfigParameters struct {

	// A list of VPC security group IDs in your Amazon VPC.
	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// A list of VPC subnet IDs in your Amazon VPC.
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// The ID of the Amazon VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// TestGridProjectSpec defines the desired state of TestGridProject
type TestGridProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TestGridProjectParameters `json:"forProvider"`
}

// TestGridProjectStatus defines the observed state of TestGridProject.
type TestGridProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TestGridProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TestGridProject is the Schema for the TestGridProjects API. Provides a Devicefarm test_grid_project
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TestGridProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TestGridProjectSpec   `json:"spec"`
	Status            TestGridProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TestGridProjectList contains a list of TestGridProjects
type TestGridProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestGridProject `json:"items"`
}

// Repository type metadata.
var (
	TestGridProject_Kind             = "TestGridProject"
	TestGridProject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TestGridProject_Kind}.String()
	TestGridProject_KindAPIVersion   = TestGridProject_Kind + "." + CRDGroupVersion.String()
	TestGridProject_GroupVersionKind = CRDGroupVersion.WithKind(TestGridProject_Kind)
)

func init() {
	SchemeBuilder.Register(&TestGridProject{}, &TestGridProjectList{})
}
