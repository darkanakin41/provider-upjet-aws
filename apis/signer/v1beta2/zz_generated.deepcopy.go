//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationInitParameters) DeepCopyInto(out *DestinationInitParameters) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3InitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationInitParameters.
func (in *DestinationInitParameters) DeepCopy() *DestinationInitParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationObservation) DeepCopyInto(out *DestinationObservation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Observation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationObservation.
func (in *DestinationObservation) DeepCopy() *DestinationObservation {
	if in == nil {
		return nil
	}
	out := new(DestinationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationParameters) DeepCopyInto(out *DestinationParameters) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Parameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationParameters.
func (in *DestinationParameters) DeepCopy() *DestinationParameters {
	if in == nil {
		return nil
	}
	out := new(DestinationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationRecordInitParameters) DeepCopyInto(out *RevocationRecordInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordInitParameters.
func (in *RevocationRecordInitParameters) DeepCopy() *RevocationRecordInitParameters {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationRecordObservation) DeepCopyInto(out *RevocationRecordObservation) {
	*out = *in
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.RevokedBy != nil {
		in, out := &in.RevokedBy, &out.RevokedBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordObservation.
func (in *RevocationRecordObservation) DeepCopy() *RevocationRecordObservation {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevocationRecordParameters) DeepCopyInto(out *RevocationRecordParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevocationRecordParameters.
func (in *RevocationRecordParameters) DeepCopy() *RevocationRecordParameters {
	if in == nil {
		return nil
	}
	out := new(RevocationRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3InitParameters) DeepCopyInto(out *S3InitParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3InitParameters.
func (in *S3InitParameters) DeepCopy() *S3InitParameters {
	if in == nil {
		return nil
	}
	out := new(S3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Observation) DeepCopyInto(out *S3Observation) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Observation.
func (in *S3Observation) DeepCopy() *S3Observation {
	if in == nil {
		return nil
	}
	out := new(S3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Parameters) DeepCopyInto(out *S3Parameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Parameters.
func (in *S3Parameters) DeepCopy() *S3Parameters {
	if in == nil {
		return nil
	}
	out := new(S3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodInitParameters) DeepCopyInto(out *SignatureValidityPeriodInitParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodInitParameters.
func (in *SignatureValidityPeriodInitParameters) DeepCopy() *SignatureValidityPeriodInitParameters {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodObservation) DeepCopyInto(out *SignatureValidityPeriodObservation) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodObservation.
func (in *SignatureValidityPeriodObservation) DeepCopy() *SignatureValidityPeriodObservation {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureValidityPeriodParameters) DeepCopyInto(out *SignatureValidityPeriodParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureValidityPeriodParameters.
func (in *SignatureValidityPeriodParameters) DeepCopy() *SignatureValidityPeriodParameters {
	if in == nil {
		return nil
	}
	out := new(SignatureValidityPeriodParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectInitParameters) DeepCopyInto(out *SignedObjectInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectInitParameters.
func (in *SignedObjectInitParameters) DeepCopy() *SignedObjectInitParameters {
	if in == nil {
		return nil
	}
	out := new(SignedObjectInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectObservation) DeepCopyInto(out *SignedObjectObservation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]SignedObjectS3Observation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectObservation.
func (in *SignedObjectObservation) DeepCopy() *SignedObjectObservation {
	if in == nil {
		return nil
	}
	out := new(SignedObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectParameters) DeepCopyInto(out *SignedObjectParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectParameters.
func (in *SignedObjectParameters) DeepCopy() *SignedObjectParameters {
	if in == nil {
		return nil
	}
	out := new(SignedObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectS3InitParameters) DeepCopyInto(out *SignedObjectS3InitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectS3InitParameters.
func (in *SignedObjectS3InitParameters) DeepCopy() *SignedObjectS3InitParameters {
	if in == nil {
		return nil
	}
	out := new(SignedObjectS3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectS3Observation) DeepCopyInto(out *SignedObjectS3Observation) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectS3Observation.
func (in *SignedObjectS3Observation) DeepCopy() *SignedObjectS3Observation {
	if in == nil {
		return nil
	}
	out := new(SignedObjectS3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignedObjectS3Parameters) DeepCopyInto(out *SignedObjectS3Parameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignedObjectS3Parameters.
func (in *SignedObjectS3Parameters) DeepCopy() *SignedObjectS3Parameters {
	if in == nil {
		return nil
	}
	out := new(SignedObjectS3Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJob) DeepCopyInto(out *SigningJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJob.
func (in *SigningJob) DeepCopy() *SigningJob {
	if in == nil {
		return nil
	}
	out := new(SigningJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobInitParameters) DeepCopyInto(out *SigningJobInitParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnoreSigningJobFailure != nil {
		in, out := &in.IgnoreSigningJobFailure, &out.IgnoreSigningJobFailure
		*out = new(bool)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ProfileNameRef != nil {
		in, out := &in.ProfileNameRef, &out.ProfileNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProfileNameSelector != nil {
		in, out := &in.ProfileNameSelector, &out.ProfileNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobInitParameters.
func (in *SigningJobInitParameters) DeepCopy() *SigningJobInitParameters {
	if in == nil {
		return nil
	}
	out := new(SigningJobInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobList) DeepCopyInto(out *SigningJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobList.
func (in *SigningJobList) DeepCopy() *SigningJobList {
	if in == nil {
		return nil
	}
	out := new(SigningJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobObservation) DeepCopyInto(out *SigningJobObservation) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IgnoreSigningJobFailure != nil {
		in, out := &in.IgnoreSigningJobFailure, &out.IgnoreSigningJobFailure
		*out = new(bool)
		**out = **in
	}
	if in.JobID != nil {
		in, out := &in.JobID, &out.JobID
		*out = new(string)
		**out = **in
	}
	if in.JobInvoker != nil {
		in, out := &in.JobInvoker, &out.JobInvoker
		*out = new(string)
		**out = **in
	}
	if in.JobOwner != nil {
		in, out := &in.JobOwner, &out.JobOwner
		*out = new(string)
		**out = **in
	}
	if in.PlatformDisplayName != nil {
		in, out := &in.PlatformDisplayName, &out.PlatformDisplayName
		*out = new(string)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ProfileVersion != nil {
		in, out := &in.ProfileVersion, &out.ProfileVersion
		*out = new(string)
		**out = **in
	}
	if in.RequestedBy != nil {
		in, out := &in.RequestedBy, &out.RequestedBy
		*out = new(string)
		**out = **in
	}
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]RevocationRecordObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignatureExpiresAt != nil {
		in, out := &in.SignatureExpiresAt, &out.SignatureExpiresAt
		*out = new(string)
		**out = **in
	}
	if in.SignedObject != nil {
		in, out := &in.SignedObject, &out.SignedObject
		*out = make([]SignedObjectObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusReason != nil {
		in, out := &in.StatusReason, &out.StatusReason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobObservation.
func (in *SigningJobObservation) DeepCopy() *SigningJobObservation {
	if in == nil {
		return nil
	}
	out := new(SigningJobObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobParameters) DeepCopyInto(out *SigningJobParameters) {
	*out = *in
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(DestinationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IgnoreSigningJobFailure != nil {
		in, out := &in.IgnoreSigningJobFailure, &out.IgnoreSigningJobFailure
		*out = new(bool)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ProfileNameRef != nil {
		in, out := &in.ProfileNameRef, &out.ProfileNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProfileNameSelector != nil {
		in, out := &in.ProfileNameSelector, &out.ProfileNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobParameters.
func (in *SigningJobParameters) DeepCopy() *SigningJobParameters {
	if in == nil {
		return nil
	}
	out := new(SigningJobParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobSpec) DeepCopyInto(out *SigningJobSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobSpec.
func (in *SigningJobSpec) DeepCopy() *SigningJobSpec {
	if in == nil {
		return nil
	}
	out := new(SigningJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningJobStatus) DeepCopyInto(out *SigningJobStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningJobStatus.
func (in *SigningJobStatus) DeepCopy() *SigningJobStatus {
	if in == nil {
		return nil
	}
	out := new(SigningJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningMaterialInitParameters) DeepCopyInto(out *SigningMaterialInitParameters) {
	*out = *in
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningMaterialInitParameters.
func (in *SigningMaterialInitParameters) DeepCopy() *SigningMaterialInitParameters {
	if in == nil {
		return nil
	}
	out := new(SigningMaterialInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningMaterialObservation) DeepCopyInto(out *SigningMaterialObservation) {
	*out = *in
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningMaterialObservation.
func (in *SigningMaterialObservation) DeepCopy() *SigningMaterialObservation {
	if in == nil {
		return nil
	}
	out := new(SigningMaterialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningMaterialParameters) DeepCopyInto(out *SigningMaterialParameters) {
	*out = *in
	if in.CertificateArn != nil {
		in, out := &in.CertificateArn, &out.CertificateArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningMaterialParameters.
func (in *SigningMaterialParameters) DeepCopy() *SigningMaterialParameters {
	if in == nil {
		return nil
	}
	out := new(SigningMaterialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfile) DeepCopyInto(out *SigningProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfile.
func (in *SigningProfile) DeepCopy() *SigningProfile {
	if in == nil {
		return nil
	}
	out := new(SigningProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileInitParameters) DeepCopyInto(out *SigningProfileInitParameters) {
	*out = *in
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = new(SignatureValidityPeriodInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SigningMaterial != nil {
		in, out := &in.SigningMaterial, &out.SigningMaterial
		*out = new(SigningMaterialInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileInitParameters.
func (in *SigningProfileInitParameters) DeepCopy() *SigningProfileInitParameters {
	if in == nil {
		return nil
	}
	out := new(SigningProfileInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileList) DeepCopyInto(out *SigningProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SigningProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileList.
func (in *SigningProfileList) DeepCopy() *SigningProfileList {
	if in == nil {
		return nil
	}
	out := new(SigningProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SigningProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileObservation) DeepCopyInto(out *SigningProfileObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PlatformDisplayName != nil {
		in, out := &in.PlatformDisplayName, &out.PlatformDisplayName
		*out = new(string)
		**out = **in
	}
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.RevocationRecord != nil {
		in, out := &in.RevocationRecord, &out.RevocationRecord
		*out = make([]SigningProfileRevocationRecordObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = new(SignatureValidityPeriodObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.SigningMaterial != nil {
		in, out := &in.SigningMaterial, &out.SigningMaterial
		*out = new(SigningMaterialObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.VersionArn != nil {
		in, out := &in.VersionArn, &out.VersionArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileObservation.
func (in *SigningProfileObservation) DeepCopy() *SigningProfileObservation {
	if in == nil {
		return nil
	}
	out := new(SigningProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileParameters) DeepCopyInto(out *SigningProfileParameters) {
	*out = *in
	if in.PlatformID != nil {
		in, out := &in.PlatformID, &out.PlatformID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SignatureValidityPeriod != nil {
		in, out := &in.SignatureValidityPeriod, &out.SignatureValidityPeriod
		*out = new(SignatureValidityPeriodParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SigningMaterial != nil {
		in, out := &in.SigningMaterial, &out.SigningMaterial
		*out = new(SigningMaterialParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileParameters.
func (in *SigningProfileParameters) DeepCopy() *SigningProfileParameters {
	if in == nil {
		return nil
	}
	out := new(SigningProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileRevocationRecordInitParameters) DeepCopyInto(out *SigningProfileRevocationRecordInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileRevocationRecordInitParameters.
func (in *SigningProfileRevocationRecordInitParameters) DeepCopy() *SigningProfileRevocationRecordInitParameters {
	if in == nil {
		return nil
	}
	out := new(SigningProfileRevocationRecordInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileRevocationRecordObservation) DeepCopyInto(out *SigningProfileRevocationRecordObservation) {
	*out = *in
	if in.RevocationEffectiveFrom != nil {
		in, out := &in.RevocationEffectiveFrom, &out.RevocationEffectiveFrom
		*out = new(string)
		**out = **in
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.RevokedBy != nil {
		in, out := &in.RevokedBy, &out.RevokedBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileRevocationRecordObservation.
func (in *SigningProfileRevocationRecordObservation) DeepCopy() *SigningProfileRevocationRecordObservation {
	if in == nil {
		return nil
	}
	out := new(SigningProfileRevocationRecordObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileRevocationRecordParameters) DeepCopyInto(out *SigningProfileRevocationRecordParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileRevocationRecordParameters.
func (in *SigningProfileRevocationRecordParameters) DeepCopy() *SigningProfileRevocationRecordParameters {
	if in == nil {
		return nil
	}
	out := new(SigningProfileRevocationRecordParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileSpec) DeepCopyInto(out *SigningProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileSpec.
func (in *SigningProfileSpec) DeepCopy() *SigningProfileSpec {
	if in == nil {
		return nil
	}
	out := new(SigningProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SigningProfileStatus) DeepCopyInto(out *SigningProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SigningProfileStatus.
func (in *SigningProfileStatus) DeepCopy() *SigningProfileStatus {
	if in == nil {
		return nil
	}
	out := new(SigningProfileStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceInitParameters) DeepCopyInto(out *SourceInitParameters) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(SourceS3InitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceInitParameters.
func (in *SourceInitParameters) DeepCopy() *SourceInitParameters {
	if in == nil {
		return nil
	}
	out := new(SourceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceObservation) DeepCopyInto(out *SourceObservation) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(SourceS3Observation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceObservation.
func (in *SourceObservation) DeepCopy() *SourceObservation {
	if in == nil {
		return nil
	}
	out := new(SourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceParameters) DeepCopyInto(out *SourceParameters) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(SourceS3Parameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceParameters.
func (in *SourceParameters) DeepCopy() *SourceParameters {
	if in == nil {
		return nil
	}
	out := new(SourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceS3InitParameters) DeepCopyInto(out *SourceS3InitParameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceS3InitParameters.
func (in *SourceS3InitParameters) DeepCopy() *SourceS3InitParameters {
	if in == nil {
		return nil
	}
	out := new(SourceS3InitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceS3Observation) DeepCopyInto(out *SourceS3Observation) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceS3Observation.
func (in *SourceS3Observation) DeepCopy() *SourceS3Observation {
	if in == nil {
		return nil
	}
	out := new(SourceS3Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceS3Parameters) DeepCopyInto(out *SourceS3Parameters) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceS3Parameters.
func (in *SourceS3Parameters) DeepCopy() *SourceS3Parameters {
	if in == nil {
		return nil
	}
	out := new(SourceS3Parameters)
	in.DeepCopyInto(out)
	return out
}