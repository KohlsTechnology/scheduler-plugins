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

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1 "k8s.io/kube-scheduler/config/v1"
	config "sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CapacitySchedulingArgs)(nil), (*config.CapacitySchedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CapacitySchedulingArgs_To_config_CapacitySchedulingArgs(a.(*CapacitySchedulingArgs), b.(*config.CapacitySchedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CapacitySchedulingArgs)(nil), (*CapacitySchedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CapacitySchedulingArgs_To_v1beta1_CapacitySchedulingArgs(a.(*config.CapacitySchedulingArgs), b.(*CapacitySchedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoschedulingArgs)(nil), (*config.CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(a.(*CoschedulingArgs), b.(*config.CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CoschedulingArgs)(nil), (*CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(a.(*config.CoschedulingArgs), b.(*CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeResourcesAllocatableArgs)(nil), (*config.NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(a.(*NodeResourcesAllocatableArgs), b.(*config.NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeResourcesAllocatableArgs)(nil), (*NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(a.(*config.NodeResourcesAllocatableArgs), b.(*NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_CapacitySchedulingArgs_To_config_CapacitySchedulingArgs(in *CapacitySchedulingArgs, out *config.CapacitySchedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CapacitySchedulingArgs_To_config_CapacitySchedulingArgs is an autogenerated conversion function.
func Convert_v1beta1_CapacitySchedulingArgs_To_config_CapacitySchedulingArgs(in *CapacitySchedulingArgs, out *config.CapacitySchedulingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_CapacitySchedulingArgs_To_config_CapacitySchedulingArgs(in, out, s)
}

func autoConvert_config_CapacitySchedulingArgs_To_v1beta1_CapacitySchedulingArgs(in *config.CapacitySchedulingArgs, out *CapacitySchedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_CapacitySchedulingArgs_To_v1beta1_CapacitySchedulingArgs is an autogenerated conversion function.
func Convert_config_CapacitySchedulingArgs_To_v1beta1_CapacitySchedulingArgs(in *config.CapacitySchedulingArgs, out *CapacitySchedulingArgs, s conversion.Scope) error {
	return autoConvert_config_CapacitySchedulingArgs_To_v1beta1_CapacitySchedulingArgs(in, out, s)
}

func autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int64_To_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.DeniedPGExpirationTimeSeconds, &out.DeniedPGExpirationTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.KubeMaster, &out.KubeMaster, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs is an autogenerated conversion function.
func Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in, out, s)
}

func autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_int64_To_Pointer_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.DeniedPGExpirationTimeSeconds, &out.DeniedPGExpirationTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.KubeMaster, &out.KubeMaster, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs is an autogenerated conversion function.
func Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in, out, s)
}

func autoConvert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = config.ModeType(in.Mode)
	return nil
}

// Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in, out, s)
}

func autoConvert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = ModeType(in.Mode)
	return nil
}

// Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in, out, s)
}