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
func (in *Directory) DeepCopyInto(out *Directory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Directory.
func (in *Directory) DeepCopy() *Directory {
	if in == nil {
		return nil
	}
	out := new(Directory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Directory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryInitParameters) DeepCopyInto(out *DirectoryInitParameters) {
	*out = *in
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.DirectoryIDRef != nil {
		in, out := &in.DirectoryIDRef, &out.DirectoryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DirectoryIDSelector != nil {
		in, out := &in.DirectoryIDSelector, &out.DirectoryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPGroupIds != nil {
		in, out := &in.IPGroupIds, &out.IPGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SelfServicePermissions != nil {
		in, out := &in.SelfServicePermissions, &out.SelfServicePermissions
		*out = new(SelfServicePermissionsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.WorkspaceAccessProperties != nil {
		in, out := &in.WorkspaceAccessProperties, &out.WorkspaceAccessProperties
		*out = new(WorkspaceAccessPropertiesInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceCreationProperties != nil {
		in, out := &in.WorkspaceCreationProperties, &out.WorkspaceCreationProperties
		*out = new(WorkspaceCreationPropertiesInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryInitParameters.
func (in *DirectoryInitParameters) DeepCopy() *DirectoryInitParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryList) DeepCopyInto(out *DirectoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Directory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryList.
func (in *DirectoryList) DeepCopy() *DirectoryList {
	if in == nil {
		return nil
	}
	out := new(DirectoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DirectoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryObservation) DeepCopyInto(out *DirectoryObservation) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.CustomerUserName != nil {
		in, out := &in.CustomerUserName, &out.CustomerUserName
		*out = new(string)
		**out = **in
	}
	if in.DNSIPAddresses != nil {
		in, out := &in.DNSIPAddresses, &out.DNSIPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.DirectoryName != nil {
		in, out := &in.DirectoryName, &out.DirectoryName
		*out = new(string)
		**out = **in
	}
	if in.DirectoryType != nil {
		in, out := &in.DirectoryType, &out.DirectoryType
		*out = new(string)
		**out = **in
	}
	if in.IAMRoleID != nil {
		in, out := &in.IAMRoleID, &out.IAMRoleID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPGroupIds != nil {
		in, out := &in.IPGroupIds, &out.IPGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RegistrationCode != nil {
		in, out := &in.RegistrationCode, &out.RegistrationCode
		*out = new(string)
		**out = **in
	}
	if in.SelfServicePermissions != nil {
		in, out := &in.SelfServicePermissions, &out.SelfServicePermissions
		*out = new(SelfServicePermissionsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.WorkspaceAccessProperties != nil {
		in, out := &in.WorkspaceAccessProperties, &out.WorkspaceAccessProperties
		*out = new(WorkspaceAccessPropertiesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceCreationProperties != nil {
		in, out := &in.WorkspaceCreationProperties, &out.WorkspaceCreationProperties
		*out = new(WorkspaceCreationPropertiesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceSecurityGroupID != nil {
		in, out := &in.WorkspaceSecurityGroupID, &out.WorkspaceSecurityGroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryObservation.
func (in *DirectoryObservation) DeepCopy() *DirectoryObservation {
	if in == nil {
		return nil
	}
	out := new(DirectoryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryParameters) DeepCopyInto(out *DirectoryParameters) {
	*out = *in
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.DirectoryIDRef != nil {
		in, out := &in.DirectoryIDRef, &out.DirectoryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DirectoryIDSelector != nil {
		in, out := &in.DirectoryIDSelector, &out.DirectoryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.IPGroupIds != nil {
		in, out := &in.IPGroupIds, &out.IPGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SelfServicePermissions != nil {
		in, out := &in.SelfServicePermissions, &out.SelfServicePermissions
		*out = new(SelfServicePermissionsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.WorkspaceAccessProperties != nil {
		in, out := &in.WorkspaceAccessProperties, &out.WorkspaceAccessProperties
		*out = new(WorkspaceAccessPropertiesParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkspaceCreationProperties != nil {
		in, out := &in.WorkspaceCreationProperties, &out.WorkspaceCreationProperties
		*out = new(WorkspaceCreationPropertiesParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryParameters.
func (in *DirectoryParameters) DeepCopy() *DirectoryParameters {
	if in == nil {
		return nil
	}
	out := new(DirectoryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectorySpec) DeepCopyInto(out *DirectorySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectorySpec.
func (in *DirectorySpec) DeepCopy() *DirectorySpec {
	if in == nil {
		return nil
	}
	out := new(DirectorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DirectoryStatus) DeepCopyInto(out *DirectoryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DirectoryStatus.
func (in *DirectoryStatus) DeepCopy() *DirectoryStatus {
	if in == nil {
		return nil
	}
	out := new(DirectoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfServicePermissionsInitParameters) DeepCopyInto(out *SelfServicePermissionsInitParameters) {
	*out = *in
	if in.ChangeComputeType != nil {
		in, out := &in.ChangeComputeType, &out.ChangeComputeType
		*out = new(bool)
		**out = **in
	}
	if in.IncreaseVolumeSize != nil {
		in, out := &in.IncreaseVolumeSize, &out.IncreaseVolumeSize
		*out = new(bool)
		**out = **in
	}
	if in.RebuildWorkspace != nil {
		in, out := &in.RebuildWorkspace, &out.RebuildWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.RestartWorkspace != nil {
		in, out := &in.RestartWorkspace, &out.RestartWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.SwitchRunningMode != nil {
		in, out := &in.SwitchRunningMode, &out.SwitchRunningMode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissionsInitParameters.
func (in *SelfServicePermissionsInitParameters) DeepCopy() *SelfServicePermissionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfServicePermissionsObservation) DeepCopyInto(out *SelfServicePermissionsObservation) {
	*out = *in
	if in.ChangeComputeType != nil {
		in, out := &in.ChangeComputeType, &out.ChangeComputeType
		*out = new(bool)
		**out = **in
	}
	if in.IncreaseVolumeSize != nil {
		in, out := &in.IncreaseVolumeSize, &out.IncreaseVolumeSize
		*out = new(bool)
		**out = **in
	}
	if in.RebuildWorkspace != nil {
		in, out := &in.RebuildWorkspace, &out.RebuildWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.RestartWorkspace != nil {
		in, out := &in.RestartWorkspace, &out.RestartWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.SwitchRunningMode != nil {
		in, out := &in.SwitchRunningMode, &out.SwitchRunningMode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissionsObservation.
func (in *SelfServicePermissionsObservation) DeepCopy() *SelfServicePermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfServicePermissionsParameters) DeepCopyInto(out *SelfServicePermissionsParameters) {
	*out = *in
	if in.ChangeComputeType != nil {
		in, out := &in.ChangeComputeType, &out.ChangeComputeType
		*out = new(bool)
		**out = **in
	}
	if in.IncreaseVolumeSize != nil {
		in, out := &in.IncreaseVolumeSize, &out.IncreaseVolumeSize
		*out = new(bool)
		**out = **in
	}
	if in.RebuildWorkspace != nil {
		in, out := &in.RebuildWorkspace, &out.RebuildWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.RestartWorkspace != nil {
		in, out := &in.RestartWorkspace, &out.RestartWorkspace
		*out = new(bool)
		**out = **in
	}
	if in.SwitchRunningMode != nil {
		in, out := &in.SwitchRunningMode, &out.SwitchRunningMode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfServicePermissionsParameters.
func (in *SelfServicePermissionsParameters) DeepCopy() *SelfServicePermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(SelfServicePermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAccessPropertiesInitParameters) DeepCopyInto(out *WorkspaceAccessPropertiesInitParameters) {
	*out = *in
	if in.DeviceTypeAndroid != nil {
		in, out := &in.DeviceTypeAndroid, &out.DeviceTypeAndroid
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeChromeos != nil {
		in, out := &in.DeviceTypeChromeos, &out.DeviceTypeChromeos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeIos != nil {
		in, out := &in.DeviceTypeIos, &out.DeviceTypeIos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeLinux != nil {
		in, out := &in.DeviceTypeLinux, &out.DeviceTypeLinux
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeOsx != nil {
		in, out := &in.DeviceTypeOsx, &out.DeviceTypeOsx
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWeb != nil {
		in, out := &in.DeviceTypeWeb, &out.DeviceTypeWeb
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWindows != nil {
		in, out := &in.DeviceTypeWindows, &out.DeviceTypeWindows
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeZeroclient != nil {
		in, out := &in.DeviceTypeZeroclient, &out.DeviceTypeZeroclient
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAccessPropertiesInitParameters.
func (in *WorkspaceAccessPropertiesInitParameters) DeepCopy() *WorkspaceAccessPropertiesInitParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAccessPropertiesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAccessPropertiesObservation) DeepCopyInto(out *WorkspaceAccessPropertiesObservation) {
	*out = *in
	if in.DeviceTypeAndroid != nil {
		in, out := &in.DeviceTypeAndroid, &out.DeviceTypeAndroid
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeChromeos != nil {
		in, out := &in.DeviceTypeChromeos, &out.DeviceTypeChromeos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeIos != nil {
		in, out := &in.DeviceTypeIos, &out.DeviceTypeIos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeLinux != nil {
		in, out := &in.DeviceTypeLinux, &out.DeviceTypeLinux
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeOsx != nil {
		in, out := &in.DeviceTypeOsx, &out.DeviceTypeOsx
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWeb != nil {
		in, out := &in.DeviceTypeWeb, &out.DeviceTypeWeb
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWindows != nil {
		in, out := &in.DeviceTypeWindows, &out.DeviceTypeWindows
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeZeroclient != nil {
		in, out := &in.DeviceTypeZeroclient, &out.DeviceTypeZeroclient
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAccessPropertiesObservation.
func (in *WorkspaceAccessPropertiesObservation) DeepCopy() *WorkspaceAccessPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAccessPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceAccessPropertiesParameters) DeepCopyInto(out *WorkspaceAccessPropertiesParameters) {
	*out = *in
	if in.DeviceTypeAndroid != nil {
		in, out := &in.DeviceTypeAndroid, &out.DeviceTypeAndroid
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeChromeos != nil {
		in, out := &in.DeviceTypeChromeos, &out.DeviceTypeChromeos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeIos != nil {
		in, out := &in.DeviceTypeIos, &out.DeviceTypeIos
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeLinux != nil {
		in, out := &in.DeviceTypeLinux, &out.DeviceTypeLinux
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeOsx != nil {
		in, out := &in.DeviceTypeOsx, &out.DeviceTypeOsx
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWeb != nil {
		in, out := &in.DeviceTypeWeb, &out.DeviceTypeWeb
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeWindows != nil {
		in, out := &in.DeviceTypeWindows, &out.DeviceTypeWindows
		*out = new(string)
		**out = **in
	}
	if in.DeviceTypeZeroclient != nil {
		in, out := &in.DeviceTypeZeroclient, &out.DeviceTypeZeroclient
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceAccessPropertiesParameters.
func (in *WorkspaceAccessPropertiesParameters) DeepCopy() *WorkspaceAccessPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceAccessPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCreationPropertiesInitParameters) DeepCopyInto(out *WorkspaceCreationPropertiesInitParameters) {
	*out = *in
	if in.CustomSecurityGroupID != nil {
		in, out := &in.CustomSecurityGroupID, &out.CustomSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.CustomSecurityGroupIDRef != nil {
		in, out := &in.CustomSecurityGroupIDRef, &out.CustomSecurityGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSecurityGroupIDSelector != nil {
		in, out := &in.CustomSecurityGroupIDSelector, &out.CustomSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultOu != nil {
		in, out := &in.DefaultOu, &out.DefaultOu
		*out = new(string)
		**out = **in
	}
	if in.EnableInternetAccess != nil {
		in, out := &in.EnableInternetAccess, &out.EnableInternetAccess
		*out = new(bool)
		**out = **in
	}
	if in.EnableMaintenanceMode != nil {
		in, out := &in.EnableMaintenanceMode, &out.EnableMaintenanceMode
		*out = new(bool)
		**out = **in
	}
	if in.UserEnabledAsLocalAdministrator != nil {
		in, out := &in.UserEnabledAsLocalAdministrator, &out.UserEnabledAsLocalAdministrator
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCreationPropertiesInitParameters.
func (in *WorkspaceCreationPropertiesInitParameters) DeepCopy() *WorkspaceCreationPropertiesInitParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCreationPropertiesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCreationPropertiesObservation) DeepCopyInto(out *WorkspaceCreationPropertiesObservation) {
	*out = *in
	if in.CustomSecurityGroupID != nil {
		in, out := &in.CustomSecurityGroupID, &out.CustomSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.DefaultOu != nil {
		in, out := &in.DefaultOu, &out.DefaultOu
		*out = new(string)
		**out = **in
	}
	if in.EnableInternetAccess != nil {
		in, out := &in.EnableInternetAccess, &out.EnableInternetAccess
		*out = new(bool)
		**out = **in
	}
	if in.EnableMaintenanceMode != nil {
		in, out := &in.EnableMaintenanceMode, &out.EnableMaintenanceMode
		*out = new(bool)
		**out = **in
	}
	if in.UserEnabledAsLocalAdministrator != nil {
		in, out := &in.UserEnabledAsLocalAdministrator, &out.UserEnabledAsLocalAdministrator
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCreationPropertiesObservation.
func (in *WorkspaceCreationPropertiesObservation) DeepCopy() *WorkspaceCreationPropertiesObservation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCreationPropertiesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceCreationPropertiesParameters) DeepCopyInto(out *WorkspaceCreationPropertiesParameters) {
	*out = *in
	if in.CustomSecurityGroupID != nil {
		in, out := &in.CustomSecurityGroupID, &out.CustomSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.CustomSecurityGroupIDRef != nil {
		in, out := &in.CustomSecurityGroupIDRef, &out.CustomSecurityGroupIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomSecurityGroupIDSelector != nil {
		in, out := &in.CustomSecurityGroupIDSelector, &out.CustomSecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultOu != nil {
		in, out := &in.DefaultOu, &out.DefaultOu
		*out = new(string)
		**out = **in
	}
	if in.EnableInternetAccess != nil {
		in, out := &in.EnableInternetAccess, &out.EnableInternetAccess
		*out = new(bool)
		**out = **in
	}
	if in.EnableMaintenanceMode != nil {
		in, out := &in.EnableMaintenanceMode, &out.EnableMaintenanceMode
		*out = new(bool)
		**out = **in
	}
	if in.UserEnabledAsLocalAdministrator != nil {
		in, out := &in.UserEnabledAsLocalAdministrator, &out.UserEnabledAsLocalAdministrator
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceCreationPropertiesParameters.
func (in *WorkspaceCreationPropertiesParameters) DeepCopy() *WorkspaceCreationPropertiesParameters {
	if in == nil {
		return nil
	}
	out := new(WorkspaceCreationPropertiesParameters)
	in.DeepCopyInto(out)
	return out
}
