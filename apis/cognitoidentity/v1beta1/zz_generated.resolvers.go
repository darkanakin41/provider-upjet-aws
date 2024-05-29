// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CognitoIdentityPoolProviderPrincipalTag) ResolveReferences( // ResolveReferences of this CognitoIdentityPoolProviderPrincipalTag.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidentity.aws.upbound.io", "v1beta1", "Pool", "PoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IdentityPoolIDRef,
			Selector:     mg.Spec.ForProvider.IdentityPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityPoolID")
	}
	mg.Spec.ForProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityPoolIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta2", "UserPool", "UserPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityProviderName),
			Extract:      resource.ExtractParamPath("endpoint", true),
			Reference:    mg.Spec.ForProvider.IdentityProviderNameRef,
			Selector:     mg.Spec.ForProvider.IdentityProviderNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityProviderName")
	}
	mg.Spec.ForProvider.IdentityProviderName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityProviderNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidentity.aws.upbound.io", "v1beta1", "Pool", "PoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IdentityPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IdentityPoolIDRef,
			Selector:     mg.Spec.InitProvider.IdentityPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IdentityPoolID")
	}
	mg.Spec.InitProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IdentityPoolIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta2", "UserPool", "UserPoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IdentityProviderName),
			Extract:      resource.ExtractParamPath("endpoint", true),
			Reference:    mg.Spec.InitProvider.IdentityProviderNameRef,
			Selector:     mg.Spec.InitProvider.IdentityProviderNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IdentityProviderName")
	}
	mg.Spec.InitProvider.IdentityProviderName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IdentityProviderNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CognitoIdentityProviders); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolClient", "UserPoolClientList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDRef,
				Selector:     mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID")
		}
		mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "SAMLProvider", "SAMLProviderList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SAMLProviderArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.ForProvider.SAMLProviderArnsRefs,
			Selector:      mg.Spec.ForProvider.SAMLProviderArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SAMLProviderArns")
	}
	mg.Spec.ForProvider.SAMLProviderArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SAMLProviderArnsRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.InitProvider.CognitoIdentityProviders); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolClient", "UserPoolClientList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientIDRef,
				Selector:     mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientID")
		}
		mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.CognitoIdentityProviders[i3].ClientIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "SAMLProvider", "SAMLProviderList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SAMLProviderArns),
			Extract:       common.ARNExtractor(),
			References:    mg.Spec.InitProvider.SAMLProviderArnsRefs,
			Selector:      mg.Spec.InitProvider.SAMLProviderArnsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SAMLProviderArns")
	}
	mg.Spec.InitProvider.SAMLProviderArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SAMLProviderArnsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this PoolRolesAttachment.
func (mg *PoolRolesAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidentity.aws.upbound.io", "v1beta1", "Pool", "PoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IdentityPoolIDRef,
			Selector:     mg.Spec.ForProvider.IdentityPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityPoolID")
	}
	mg.Spec.ForProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityPoolIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RoleMapping); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RoleMapping[i3].MappingRule); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef,
					Selector:     mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn")
			}
			mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("cognitoidentity.aws.upbound.io", "v1beta1", "Pool", "PoolList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IdentityPoolID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IdentityPoolIDRef,
			Selector:     mg.Spec.InitProvider.IdentityPoolIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IdentityPoolID")
	}
	mg.Spec.InitProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IdentityPoolIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.RoleMapping); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.RoleMapping[i3].MappingRule); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef,
					Selector:     mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArn")
			}
			mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
