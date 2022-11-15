/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DevicePool.
func (mg *DevicePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ProjectArnRef,
		Selector:     mg.Spec.ForProvider.ProjectArnSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectArn")
	}
	mg.Spec.ForProvider.ProjectArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NetworkProfile.
func (mg *NetworkProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ProjectArnRef,
		Selector:     mg.Spec.ForProvider.ProjectArnSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectArn")
	}
	mg.Spec.ForProvider.ProjectArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TestGridProject.
func (mg *TestGridProject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCConfig[i3].VPCID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VPCConfig[i3].VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPCConfig[i3].VPCIDSelector,
			To: reference.To{
				List:    &v1beta1.VPCList{},
				Managed: &v1beta1.VPC{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfig[i3].VPCID")
		}
		mg.Spec.ForProvider.VPCConfig[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCConfig[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Upload.
func (mg *Upload) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProjectArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.ProjectArnRef,
		Selector:     mg.Spec.ForProvider.ProjectArnSelector,
		To: reference.To{
			List:    &ProjectList{},
			Managed: &Project{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProjectArn")
	}
	mg.Spec.ForProvider.ProjectArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProjectArnRef = rsp.ResolvedReference

	return nil
}
