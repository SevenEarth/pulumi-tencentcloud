# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetDetectArgs', 'NetDetect']

@pulumi.input_type
class NetDetectArgs:
    def __init__(__self__, *,
                 detect_destination_ips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 net_detect_name: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 net_detect_description: Optional[pulumi.Input[str]] = None,
                 next_hop_destination: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetDetect resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] detect_destination_ips: An array of probe destination IPv4 addresses. Up to two.
        :param pulumi.Input[str] net_detect_name: Network probe name, the maximum length cannot exceed 60 bytes.
        :param pulumi.Input[str] subnet_id: Subnet instance ID. Such as:subnet-12345678.
        :param pulumi.Input[str] vpc_id: `VPC` instance `ID`. Such as:`vpc-12345678`.
        :param pulumi.Input[str] net_detect_description: Network probe description.
        :param pulumi.Input[str] next_hop_destination: The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        :param pulumi.Input[str] next_hop_type: The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        """
        pulumi.set(__self__, "detect_destination_ips", detect_destination_ips)
        pulumi.set(__self__, "net_detect_name", net_detect_name)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if net_detect_description is not None:
            pulumi.set(__self__, "net_detect_description", net_detect_description)
        if next_hop_destination is not None:
            pulumi.set(__self__, "next_hop_destination", next_hop_destination)
        if next_hop_type is not None:
            pulumi.set(__self__, "next_hop_type", next_hop_type)

    @property
    @pulumi.getter(name="detectDestinationIps")
    def detect_destination_ips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of probe destination IPv4 addresses. Up to two.
        """
        return pulumi.get(self, "detect_destination_ips")

    @detect_destination_ips.setter
    def detect_destination_ips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "detect_destination_ips", value)

    @property
    @pulumi.getter(name="netDetectName")
    def net_detect_name(self) -> pulumi.Input[str]:
        """
        Network probe name, the maximum length cannot exceed 60 bytes.
        """
        return pulumi.get(self, "net_detect_name")

    @net_detect_name.setter
    def net_detect_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "net_detect_name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet instance ID. Such as:subnet-12345678.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="netDetectDescription")
    def net_detect_description(self) -> Optional[pulumi.Input[str]]:
        """
        Network probe description.
        """
        return pulumi.get(self, "net_detect_description")

    @net_detect_description.setter
    def net_detect_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_detect_description", value)

    @property
    @pulumi.getter(name="nextHopDestination")
    def next_hop_destination(self) -> Optional[pulumi.Input[str]]:
        """
        The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        """
        return pulumi.get(self, "next_hop_destination")

    @next_hop_destination.setter
    def next_hop_destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_destination", value)

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        """
        return pulumi.get(self, "next_hop_type")

    @next_hop_type.setter
    def next_hop_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_type", value)


@pulumi.input_type
class _NetDetectState:
    def __init__(__self__, *,
                 detect_destination_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 net_detect_description: Optional[pulumi.Input[str]] = None,
                 net_detect_name: Optional[pulumi.Input[str]] = None,
                 next_hop_destination: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetDetect resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] detect_destination_ips: An array of probe destination IPv4 addresses. Up to two.
        :param pulumi.Input[str] net_detect_description: Network probe description.
        :param pulumi.Input[str] net_detect_name: Network probe name, the maximum length cannot exceed 60 bytes.
        :param pulumi.Input[str] next_hop_destination: The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        :param pulumi.Input[str] next_hop_type: The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        :param pulumi.Input[str] subnet_id: Subnet instance ID. Such as:subnet-12345678.
        :param pulumi.Input[str] vpc_id: `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        if detect_destination_ips is not None:
            pulumi.set(__self__, "detect_destination_ips", detect_destination_ips)
        if net_detect_description is not None:
            pulumi.set(__self__, "net_detect_description", net_detect_description)
        if net_detect_name is not None:
            pulumi.set(__self__, "net_detect_name", net_detect_name)
        if next_hop_destination is not None:
            pulumi.set(__self__, "next_hop_destination", next_hop_destination)
        if next_hop_type is not None:
            pulumi.set(__self__, "next_hop_type", next_hop_type)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="detectDestinationIps")
    def detect_destination_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of probe destination IPv4 addresses. Up to two.
        """
        return pulumi.get(self, "detect_destination_ips")

    @detect_destination_ips.setter
    def detect_destination_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "detect_destination_ips", value)

    @property
    @pulumi.getter(name="netDetectDescription")
    def net_detect_description(self) -> Optional[pulumi.Input[str]]:
        """
        Network probe description.
        """
        return pulumi.get(self, "net_detect_description")

    @net_detect_description.setter
    def net_detect_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_detect_description", value)

    @property
    @pulumi.getter(name="netDetectName")
    def net_detect_name(self) -> Optional[pulumi.Input[str]]:
        """
        Network probe name, the maximum length cannot exceed 60 bytes.
        """
        return pulumi.get(self, "net_detect_name")

    @net_detect_name.setter
    def net_detect_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "net_detect_name", value)

    @property
    @pulumi.getter(name="nextHopDestination")
    def next_hop_destination(self) -> Optional[pulumi.Input[str]]:
        """
        The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        """
        return pulumi.get(self, "next_hop_destination")

    @next_hop_destination.setter
    def next_hop_destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_destination", value)

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        """
        return pulumi.get(self, "next_hop_type")

    @next_hop_type.setter
    def next_hop_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_type", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet instance ID. Such as:subnet-12345678.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class NetDetect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 detect_destination_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 net_detect_description: Optional[pulumi.Input[str]] = None,
                 net_detect_name: Optional[pulumi.Input[str]] = None,
                 next_hop_destination: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc net_detect

        ## Example Usage

        ## Import

        vpc net_detect can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Vpc/netDetect:NetDetect net_detect net_detect_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] detect_destination_ips: An array of probe destination IPv4 addresses. Up to two.
        :param pulumi.Input[str] net_detect_description: Network probe description.
        :param pulumi.Input[str] net_detect_name: Network probe name, the maximum length cannot exceed 60 bytes.
        :param pulumi.Input[str] next_hop_destination: The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        :param pulumi.Input[str] next_hop_type: The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        :param pulumi.Input[str] subnet_id: Subnet instance ID. Such as:subnet-12345678.
        :param pulumi.Input[str] vpc_id: `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetDetectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc net_detect

        ## Example Usage

        ## Import

        vpc net_detect can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Vpc/netDetect:NetDetect net_detect net_detect_id
        ```

        :param str resource_name: The name of the resource.
        :param NetDetectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetDetectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 detect_destination_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 net_detect_description: Optional[pulumi.Input[str]] = None,
                 net_detect_name: Optional[pulumi.Input[str]] = None,
                 next_hop_destination: Optional[pulumi.Input[str]] = None,
                 next_hop_type: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetDetectArgs.__new__(NetDetectArgs)

            if detect_destination_ips is None and not opts.urn:
                raise TypeError("Missing required property 'detect_destination_ips'")
            __props__.__dict__["detect_destination_ips"] = detect_destination_ips
            __props__.__dict__["net_detect_description"] = net_detect_description
            if net_detect_name is None and not opts.urn:
                raise TypeError("Missing required property 'net_detect_name'")
            __props__.__dict__["net_detect_name"] = net_detect_name
            __props__.__dict__["next_hop_destination"] = next_hop_destination
            __props__.__dict__["next_hop_type"] = next_hop_type
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(NetDetect, __self__).__init__(
            'tencentcloud:Vpc/netDetect:NetDetect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            detect_destination_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            net_detect_description: Optional[pulumi.Input[str]] = None,
            net_detect_name: Optional[pulumi.Input[str]] = None,
            next_hop_destination: Optional[pulumi.Input[str]] = None,
            next_hop_type: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'NetDetect':
        """
        Get an existing NetDetect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] detect_destination_ips: An array of probe destination IPv4 addresses. Up to two.
        :param pulumi.Input[str] net_detect_description: Network probe description.
        :param pulumi.Input[str] net_detect_name: Network probe name, the maximum length cannot exceed 60 bytes.
        :param pulumi.Input[str] next_hop_destination: The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        :param pulumi.Input[str] next_hop_type: The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        :param pulumi.Input[str] subnet_id: Subnet instance ID. Such as:subnet-12345678.
        :param pulumi.Input[str] vpc_id: `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetDetectState.__new__(_NetDetectState)

        __props__.__dict__["detect_destination_ips"] = detect_destination_ips
        __props__.__dict__["net_detect_description"] = net_detect_description
        __props__.__dict__["net_detect_name"] = net_detect_name
        __props__.__dict__["next_hop_destination"] = next_hop_destination
        __props__.__dict__["next_hop_type"] = next_hop_type
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["vpc_id"] = vpc_id
        return NetDetect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="detectDestinationIps")
    def detect_destination_ips(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of probe destination IPv4 addresses. Up to two.
        """
        return pulumi.get(self, "detect_destination_ips")

    @property
    @pulumi.getter(name="netDetectDescription")
    def net_detect_description(self) -> pulumi.Output[Optional[str]]:
        """
        Network probe description.
        """
        return pulumi.get(self, "net_detect_description")

    @property
    @pulumi.getter(name="netDetectName")
    def net_detect_name(self) -> pulumi.Output[str]:
        """
        Network probe name, the maximum length cannot exceed 60 bytes.
        """
        return pulumi.get(self, "net_detect_name")

    @property
    @pulumi.getter(name="nextHopDestination")
    def next_hop_destination(self) -> pulumi.Output[Optional[str]]:
        """
        The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        """
        return pulumi.get(self, "next_hop_destination")

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> pulumi.Output[Optional[str]]:
        """
        The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        """
        return pulumi.get(self, "next_hop_type")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Subnet instance ID. Such as:subnet-12345678.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        `VPC` instance `ID`. Such as:`vpc-12345678`.
        """
        return pulumi.get(self, "vpc_id")

