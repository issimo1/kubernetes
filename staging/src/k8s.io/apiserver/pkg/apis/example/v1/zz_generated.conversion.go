//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	example "k8s.io/apiserver/pkg/apis/example"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Pod)(nil), (*example.Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Pod_To_example_Pod(a.(*Pod), b.(*example.Pod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.Pod)(nil), (*Pod)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_Pod_To_v1_Pod(a.(*example.Pod), b.(*Pod), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodCondition)(nil), (*example.PodCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodCondition_To_example_PodCondition(a.(*PodCondition), b.(*example.PodCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodCondition)(nil), (*PodCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodCondition_To_v1_PodCondition(a.(*example.PodCondition), b.(*PodCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodList)(nil), (*example.PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodList_To_example_PodList(a.(*PodList), b.(*example.PodList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodList)(nil), (*PodList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodList_To_v1_PodList(a.(*example.PodList), b.(*PodList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodSpec)(nil), (*example.PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodSpec_To_example_PodSpec(a.(*PodSpec), b.(*example.PodSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodSpec)(nil), (*PodSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodSpec_To_v1_PodSpec(a.(*example.PodSpec), b.(*PodSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodStatus)(nil), (*example.PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PodStatus_To_example_PodStatus(a.(*PodStatus), b.(*example.PodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.PodStatus)(nil), (*PodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_PodStatus_To_v1_PodStatus(a.(*example.PodStatus), b.(*PodStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Pod_To_example_Pod(in *Pod, out *example.Pod, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_PodSpec_To_example_PodSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_PodStatus_To_example_PodStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Pod_To_example_Pod is an autogenerated conversion function.
func Convert_v1_Pod_To_example_Pod(in *Pod, out *example.Pod, s conversion.Scope) error {
	return autoConvert_v1_Pod_To_example_Pod(in, out, s)
}

func autoConvert_example_Pod_To_v1_Pod(in *example.Pod, out *Pod, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_example_PodSpec_To_v1_PodSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_example_PodStatus_To_v1_PodStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_Pod_To_v1_Pod is an autogenerated conversion function.
func Convert_example_Pod_To_v1_Pod(in *example.Pod, out *Pod, s conversion.Scope) error {
	return autoConvert_example_Pod_To_v1_Pod(in, out, s)
}

func autoConvert_v1_PodCondition_To_example_PodCondition(in *PodCondition, out *example.PodCondition, s conversion.Scope) error {
	out.Type = example.PodConditionType(in.Type)
	out.Status = example.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_PodCondition_To_example_PodCondition is an autogenerated conversion function.
func Convert_v1_PodCondition_To_example_PodCondition(in *PodCondition, out *example.PodCondition, s conversion.Scope) error {
	return autoConvert_v1_PodCondition_To_example_PodCondition(in, out, s)
}

func autoConvert_example_PodCondition_To_v1_PodCondition(in *example.PodCondition, out *PodCondition, s conversion.Scope) error {
	out.Type = PodConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_example_PodCondition_To_v1_PodCondition is an autogenerated conversion function.
func Convert_example_PodCondition_To_v1_PodCondition(in *example.PodCondition, out *PodCondition, s conversion.Scope) error {
	return autoConvert_example_PodCondition_To_v1_PodCondition(in, out, s)
}

func autoConvert_v1_PodList_To_example_PodList(in *PodList, out *example.PodList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]example.Pod, len(*in))
		for i := range *in {
			if err := Convert_v1_Pod_To_example_Pod(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_PodList_To_example_PodList is an autogenerated conversion function.
func Convert_v1_PodList_To_example_PodList(in *PodList, out *example.PodList, s conversion.Scope) error {
	return autoConvert_v1_PodList_To_example_PodList(in, out, s)
}

func autoConvert_example_PodList_To_v1_PodList(in *example.PodList, out *PodList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pod, len(*in))
		for i := range *in {
			if err := Convert_example_Pod_To_v1_Pod(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_example_PodList_To_v1_PodList is an autogenerated conversion function.
func Convert_example_PodList_To_v1_PodList(in *example.PodList, out *PodList, s conversion.Scope) error {
	return autoConvert_example_PodList_To_v1_PodList(in, out, s)
}

func autoConvert_v1_PodSpec_To_example_PodSpec(in *PodSpec, out *example.PodSpec, s conversion.Scope) error {
	out.RestartPolicy = example.RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	// INFO: in.DeprecatedServiceAccount opted out of conversion generation
	out.NodeName = in.NodeName
	// INFO: in.HostNetwork opted out of conversion generation
	// INFO: in.HostPID opted out of conversion generation
	// INFO: in.HostIPC opted out of conversion generation
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_v1_PodSpec_To_example_PodSpec is an autogenerated conversion function.
func Convert_v1_PodSpec_To_example_PodSpec(in *PodSpec, out *example.PodSpec, s conversion.Scope) error {
	return autoConvert_v1_PodSpec_To_example_PodSpec(in, out, s)
}

func autoConvert_example_PodSpec_To_v1_PodSpec(in *example.PodSpec, out *PodSpec, s conversion.Scope) error {
	out.RestartPolicy = RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	out.NodeName = in.NodeName
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_example_PodSpec_To_v1_PodSpec is an autogenerated conversion function.
func Convert_example_PodSpec_To_v1_PodSpec(in *example.PodSpec, out *PodSpec, s conversion.Scope) error {
	return autoConvert_example_PodSpec_To_v1_PodSpec(in, out, s)
}

func autoConvert_v1_PodStatus_To_example_PodStatus(in *PodStatus, out *example.PodStatus, s conversion.Scope) error {
	out.Phase = example.PodPhase(in.Phase)
	out.Conditions = *(*[]example.PodCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.PodIP = in.PodIP
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_v1_PodStatus_To_example_PodStatus is an autogenerated conversion function.
func Convert_v1_PodStatus_To_example_PodStatus(in *PodStatus, out *example.PodStatus, s conversion.Scope) error {
	return autoConvert_v1_PodStatus_To_example_PodStatus(in, out, s)
}

func autoConvert_example_PodStatus_To_v1_PodStatus(in *example.PodStatus, out *PodStatus, s conversion.Scope) error {
	out.Phase = PodPhase(in.Phase)
	out.Conditions = *(*[]PodCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.PodIP = in.PodIP
	out.StartTime = (*metav1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_example_PodStatus_To_v1_PodStatus is an autogenerated conversion function.
func Convert_example_PodStatus_To_v1_PodStatus(in *example.PodStatus, out *PodStatus, s conversion.Scope) error {
	return autoConvert_example_PodStatus_To_v1_PodStatus(in, out, s)
}