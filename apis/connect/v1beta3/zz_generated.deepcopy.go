//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta3

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigInitParameters) DeepCopyInto(out *ConfigInitParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(EndTimeInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(StartTimeInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigInitParameters.
func (in *ConfigInitParameters) DeepCopy() *ConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(EndTimeObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(StartTimeObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(EndTimeParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(StartTimeParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndTimeInitParameters) DeepCopyInto(out *EndTimeInitParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndTimeInitParameters.
func (in *EndTimeInitParameters) DeepCopy() *EndTimeInitParameters {
	if in == nil {
		return nil
	}
	out := new(EndTimeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndTimeObservation) DeepCopyInto(out *EndTimeObservation) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndTimeObservation.
func (in *EndTimeObservation) DeepCopy() *EndTimeObservation {
	if in == nil {
		return nil
	}
	out := new(EndTimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndTimeParameters) DeepCopyInto(out *EndTimeParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndTimeParameters.
func (in *EndTimeParameters) DeepCopy() *EndTimeParameters {
	if in == nil {
		return nil
	}
	out := new(EndTimeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperation) DeepCopyInto(out *HoursOfOperation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperation.
func (in *HoursOfOperation) DeepCopy() *HoursOfOperation {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HoursOfOperation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationInitParameters) DeepCopyInto(out *HoursOfOperationInitParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceIDRef != nil {
		in, out := &in.InstanceIDRef, &out.InstanceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceIDSelector != nil {
		in, out := &in.InstanceIDSelector, &out.InstanceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
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
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationInitParameters.
func (in *HoursOfOperationInitParameters) DeepCopy() *HoursOfOperationInitParameters {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationList) DeepCopyInto(out *HoursOfOperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HoursOfOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationList.
func (in *HoursOfOperationList) DeepCopy() *HoursOfOperationList {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HoursOfOperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationObservation) DeepCopyInto(out *HoursOfOperationObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationID != nil {
		in, out := &in.HoursOfOperationID, &out.HoursOfOperationID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
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
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationObservation.
func (in *HoursOfOperationObservation) DeepCopy() *HoursOfOperationObservation {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationParameters) DeepCopyInto(out *HoursOfOperationParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceIDRef != nil {
		in, out := &in.InstanceIDRef, &out.InstanceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceIDSelector != nil {
		in, out := &in.InstanceIDSelector, &out.InstanceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationParameters.
func (in *HoursOfOperationParameters) DeepCopy() *HoursOfOperationParameters {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationSpec) DeepCopyInto(out *HoursOfOperationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationSpec.
func (in *HoursOfOperationSpec) DeepCopy() *HoursOfOperationSpec {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HoursOfOperationStatus) DeepCopyInto(out *HoursOfOperationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HoursOfOperationStatus.
func (in *HoursOfOperationStatus) DeepCopy() *HoursOfOperationStatus {
	if in == nil {
		return nil
	}
	out := new(HoursOfOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutboundCallerConfigInitParameters) DeepCopyInto(out *OutboundCallerConfigInitParameters) {
	*out = *in
	if in.OutboundCallerIDName != nil {
		in, out := &in.OutboundCallerIDName, &out.OutboundCallerIDName
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerIDNumberID != nil {
		in, out := &in.OutboundCallerIDNumberID, &out.OutboundCallerIDNumberID
		*out = new(string)
		**out = **in
	}
	if in.OutboundFlowID != nil {
		in, out := &in.OutboundFlowID, &out.OutboundFlowID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutboundCallerConfigInitParameters.
func (in *OutboundCallerConfigInitParameters) DeepCopy() *OutboundCallerConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(OutboundCallerConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutboundCallerConfigObservation) DeepCopyInto(out *OutboundCallerConfigObservation) {
	*out = *in
	if in.OutboundCallerIDName != nil {
		in, out := &in.OutboundCallerIDName, &out.OutboundCallerIDName
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerIDNumberID != nil {
		in, out := &in.OutboundCallerIDNumberID, &out.OutboundCallerIDNumberID
		*out = new(string)
		**out = **in
	}
	if in.OutboundFlowID != nil {
		in, out := &in.OutboundFlowID, &out.OutboundFlowID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutboundCallerConfigObservation.
func (in *OutboundCallerConfigObservation) DeepCopy() *OutboundCallerConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OutboundCallerConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutboundCallerConfigParameters) DeepCopyInto(out *OutboundCallerConfigParameters) {
	*out = *in
	if in.OutboundCallerIDName != nil {
		in, out := &in.OutboundCallerIDName, &out.OutboundCallerIDName
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerIDNumberID != nil {
		in, out := &in.OutboundCallerIDNumberID, &out.OutboundCallerIDNumberID
		*out = new(string)
		**out = **in
	}
	if in.OutboundFlowID != nil {
		in, out := &in.OutboundFlowID, &out.OutboundFlowID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutboundCallerConfigParameters.
func (in *OutboundCallerConfigParameters) DeepCopy() *OutboundCallerConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OutboundCallerConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueInitParameters) DeepCopyInto(out *QueueInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationID != nil {
		in, out := &in.HoursOfOperationID, &out.HoursOfOperationID
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationIDRef != nil {
		in, out := &in.HoursOfOperationIDRef, &out.HoursOfOperationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.HoursOfOperationIDSelector != nil {
		in, out := &in.HoursOfOperationIDSelector, &out.HoursOfOperationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceIDRef != nil {
		in, out := &in.InstanceIDRef, &out.InstanceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceIDSelector != nil {
		in, out := &in.InstanceIDSelector, &out.InstanceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxContacts != nil {
		in, out := &in.MaxContacts, &out.MaxContacts
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerConfig != nil {
		in, out := &in.OutboundCallerConfig, &out.OutboundCallerConfig
		*out = new(OutboundCallerConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.QuickConnectIds != nil {
		in, out := &in.QuickConnectIds, &out.QuickConnectIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueInitParameters.
func (in *QueueInitParameters) DeepCopy() *QueueInitParameters {
	if in == nil {
		return nil
	}
	out := new(QueueInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueObservation) DeepCopyInto(out *QueueObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationID != nil {
		in, out := &in.HoursOfOperationID, &out.HoursOfOperationID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.MaxContacts != nil {
		in, out := &in.MaxContacts, &out.MaxContacts
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerConfig != nil {
		in, out := &in.OutboundCallerConfig, &out.OutboundCallerConfig
		*out = new(OutboundCallerConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueID != nil {
		in, out := &in.QueueID, &out.QueueID
		*out = new(string)
		**out = **in
	}
	if in.QuickConnectIds != nil {
		in, out := &in.QuickConnectIds, &out.QuickConnectIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueObservation.
func (in *QueueObservation) DeepCopy() *QueueObservation {
	if in == nil {
		return nil
	}
	out := new(QueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueParameters) DeepCopyInto(out *QueueParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationID != nil {
		in, out := &in.HoursOfOperationID, &out.HoursOfOperationID
		*out = new(string)
		**out = **in
	}
	if in.HoursOfOperationIDRef != nil {
		in, out := &in.HoursOfOperationIDRef, &out.HoursOfOperationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.HoursOfOperationIDSelector != nil {
		in, out := &in.HoursOfOperationIDSelector, &out.HoursOfOperationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceIDRef != nil {
		in, out := &in.InstanceIDRef, &out.InstanceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceIDSelector != nil {
		in, out := &in.InstanceIDSelector, &out.InstanceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxContacts != nil {
		in, out := &in.MaxContacts, &out.MaxContacts
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OutboundCallerConfig != nil {
		in, out := &in.OutboundCallerConfig, &out.OutboundCallerConfig
		*out = new(OutboundCallerConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.QuickConnectIds != nil {
		in, out := &in.QuickConnectIds, &out.QuickConnectIds
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueParameters.
func (in *QueueParameters) DeepCopy() *QueueParameters {
	if in == nil {
		return nil
	}
	out := new(QueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeInitParameters) DeepCopyInto(out *StartTimeInitParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeInitParameters.
func (in *StartTimeInitParameters) DeepCopy() *StartTimeInitParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeObservation) DeepCopyInto(out *StartTimeObservation) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeObservation.
func (in *StartTimeObservation) DeepCopy() *StartTimeObservation {
	if in == nil {
		return nil
	}
	out := new(StartTimeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartTimeParameters) DeepCopyInto(out *StartTimeParameters) {
	*out = *in
	if in.Hours != nil {
		in, out := &in.Hours, &out.Hours
		*out = new(float64)
		**out = **in
	}
	if in.Minutes != nil {
		in, out := &in.Minutes, &out.Minutes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartTimeParameters.
func (in *StartTimeParameters) DeepCopy() *StartTimeParameters {
	if in == nil {
		return nil
	}
	out := new(StartTimeParameters)
	in.DeepCopyInto(out)
	return out
}
