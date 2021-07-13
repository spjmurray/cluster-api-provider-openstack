// +build !ignore_autogenerated

/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha4

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPair) DeepCopyInto(out *AddressPair) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPair.
func (in *AddressPair) DeepCopy() *AddressPair {
	if in == nil {
		return nil
	}
	out := new(AddressPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bastion) DeepCopyInto(out *Bastion) {
	*out = *in
	in.Instance.DeepCopyInto(&out.Instance)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bastion.
func (in *Bastion) DeepCopy() *Bastion {
	if in == nil {
		return nil
	}
	out := new(Bastion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalRouterIPParam) DeepCopyInto(out *ExternalRouterIPParam) {
	*out = *in
	in.Subnet.DeepCopyInto(&out.Subnet)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalRouterIPParam.
func (in *ExternalRouterIPParam) DeepCopy() *ExternalRouterIPParam {
	if in == nil {
		return nil
	}
	out := new(ExternalRouterIPParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedIP) DeepCopyInto(out *FixedIP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedIP.
func (in *FixedIP) DeepCopy() *FixedIP {
	if in == nil {
		return nil
	}
	out := new(FixedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = new([]Network)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Network, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigDrive != nil {
		in, out := &in.ConfigDrive, &out.ConfigDrive
		*out = new(bool)
		**out = **in
	}
	if in.RootVolume != nil {
		in, out := &in.RootVolume, &out.RootVolume
		*out = new(RootVolume)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(Subnet)
		(*in).DeepCopyInto(*out)
	}
	if in.PortOpts != nil {
		in, out := &in.PortOpts, &out.PortOpts
		*out = new(PortOpts)
		(*in).DeepCopyInto(*out)
	}
	if in.Router != nil {
		in, out := &in.Router, &out.Router
		*out = new(Router)
		(*in).DeepCopyInto(*out)
	}
	if in.APIServerLoadBalancer != nil {
		in, out := &in.APIServerLoadBalancer, &out.APIServerLoadBalancer
		*out = new(LoadBalancer)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParam) DeepCopyInto(out *NetworkParam) {
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParam.
func (in *NetworkParam) DeepCopy() *NetworkParam {
	if in == nil {
		return nil
	}
	out := new(NetworkParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackCluster) DeepCopyInto(out *OpenStackCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackCluster.
func (in *OpenStackCluster) DeepCopy() *OpenStackCluster {
	if in == nil {
		return nil
	}
	out := new(OpenStackCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClusterList) DeepCopyInto(out *OpenStackClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClusterList.
func (in *OpenStackClusterList) DeepCopy() *OpenStackClusterList {
	if in == nil {
		return nil
	}
	out := new(OpenStackClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClusterSpec) DeepCopyInto(out *OpenStackClusterSpec) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	in.Subnet.DeepCopyInto(&out.Subnet)
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalRouterIPs != nil {
		in, out := &in.ExternalRouterIPs, &out.ExternalRouterIPs
		*out = make([]ExternalRouterIPParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIServerLoadBalancerAdditionalPorts != nil {
		in, out := &in.APIServerLoadBalancerAdditionalPorts, &out.APIServerLoadBalancerAdditionalPorts
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.ControlPlaneAvailabilityZones != nil {
		in, out := &in.ControlPlaneAvailabilityZones, &out.ControlPlaneAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(Bastion)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(OpenStackIdentityReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClusterSpec.
func (in *OpenStackClusterSpec) DeepCopy() *OpenStackClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClusterStatus) DeepCopyInto(out *OpenStackClusterStatus) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalNetwork != nil {
		in, out := &in.ExternalNetwork, &out.ExternalNetwork
		*out = new(Network)
		(*in).DeepCopyInto(*out)
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1alpha4.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ControlPlaneSecurityGroup != nil {
		in, out := &in.ControlPlaneSecurityGroup, &out.ControlPlaneSecurityGroup
		*out = new(SecurityGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerSecurityGroup != nil {
		in, out := &in.WorkerSecurityGroup, &out.WorkerSecurityGroup
		*out = new(SecurityGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.BastionSecurityGroup != nil {
		in, out := &in.BastionSecurityGroup, &out.BastionSecurityGroup
		*out = new(SecurityGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(Instance)
		(*in).DeepCopyInto(*out)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.ClusterStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClusterStatus.
func (in *OpenStackClusterStatus) DeepCopy() *OpenStackClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackIdentityReference) DeepCopyInto(out *OpenStackIdentityReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackIdentityReference.
func (in *OpenStackIdentityReference) DeepCopy() *OpenStackIdentityReference {
	if in == nil {
		return nil
	}
	out := new(OpenStackIdentityReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachine) DeepCopyInto(out *OpenStackMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachine.
func (in *OpenStackMachine) DeepCopy() *OpenStackMachine {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineList) DeepCopyInto(out *OpenStackMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineList.
func (in *OpenStackMachineList) DeepCopy() *OpenStackMachineList {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineSpec) DeepCopyInto(out *OpenStackMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortOpts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroupParam, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServerMetadata != nil {
		in, out := &in.ServerMetadata, &out.ServerMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigDrive != nil {
		in, out := &in.ConfigDrive, &out.ConfigDrive
		*out = new(bool)
		**out = **in
	}
	if in.RootVolume != nil {
		in, out := &in.RootVolume, &out.RootVolume
		*out = new(RootVolume)
		**out = **in
	}
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(OpenStackIdentityReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineSpec.
func (in *OpenStackMachineSpec) DeepCopy() *OpenStackMachineSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineStatus) DeepCopyInto(out *OpenStackMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(InstanceState)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineStatus.
func (in *OpenStackMachineStatus) DeepCopy() *OpenStackMachineStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineTemplate) DeepCopyInto(out *OpenStackMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineTemplate.
func (in *OpenStackMachineTemplate) DeepCopy() *OpenStackMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineTemplateList) DeepCopyInto(out *OpenStackMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineTemplateList.
func (in *OpenStackMachineTemplateList) DeepCopy() *OpenStackMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineTemplateResource) DeepCopyInto(out *OpenStackMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineTemplateResource.
func (in *OpenStackMachineTemplateResource) DeepCopy() *OpenStackMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackMachineTemplateSpec) DeepCopyInto(out *OpenStackMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackMachineTemplateSpec.
func (in *OpenStackMachineTemplateSpec) DeepCopy() *OpenStackMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortOpts) DeepCopyInto(out *PortOpts) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.FixedIPs != nil {
		in, out := &in.FixedIPs, &out.FixedIPs
		*out = make([]FixedIP, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedAddressPairs != nil {
		in, out := &in.AllowedAddressPairs, &out.AllowedAddressPairs
		*out = make([]AddressPair, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortOpts.
func (in *PortOpts) DeepCopy() *PortOpts {
	if in == nil {
		return nil
	}
	out := new(PortOpts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootVolume) DeepCopyInto(out *RootVolume) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootVolume.
func (in *RootVolume) DeepCopy() *RootVolume {
	if in == nil {
		return nil
	}
	out := new(RootVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Router) DeepCopyInto(out *Router) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Router.
func (in *Router) DeepCopy() *Router {
	if in == nil {
		return nil
	}
	out := new(Router)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupFilter) DeepCopyInto(out *SecurityGroupFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupFilter.
func (in *SecurityGroupFilter) DeepCopy() *SecurityGroupFilter {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupParam) DeepCopyInto(out *SecurityGroupParam) {
	*out = *in
	out.Filter = in.Filter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupParam.
func (in *SecurityGroupParam) DeepCopy() *SecurityGroupParam {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupRule.
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetFilter) DeepCopyInto(out *SubnetFilter) {
	*out = *in
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetFilter.
func (in *SubnetFilter) DeepCopy() *SubnetFilter {
	if in == nil {
		return nil
	}
	out := new(SubnetFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetParam) DeepCopyInto(out *SubnetParam) {
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetParam.
func (in *SubnetParam) DeepCopy() *SubnetParam {
	if in == nil {
		return nil
	}
	out := new(SubnetParam)
	in.DeepCopyInto(out)
	return out
}
