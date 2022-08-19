# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewaySnatArgs', 'GatewaySnat']

@pulumi.input_type
class GatewaySnatArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 nat_gateway_id: pulumi.Input[str],
                 public_ip_addrs: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_type: pulumi.Input[str],
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_private_ip_addr: Optional[pulumi.Input[str]] = None,
                 subnet_cidr_block: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GatewaySnat resource.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] nat_gateway_id: NAT gateway ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addrs: Elastic IP address pool.
        :param pulumi.Input[str] resource_type: Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        :param pulumi.Input[str] instance_id: Instance ID, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] instance_private_ip_addr: Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] subnet_cidr_block: The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        :param pulumi.Input[str] subnet_id: Subnet instance ID, required when `resource_type` is SUBNET.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        pulumi.set(__self__, "public_ip_addrs", public_ip_addrs)
        pulumi.set(__self__, "resource_type", resource_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_private_ip_addr is not None:
            pulumi.set(__self__, "instance_private_ip_addr", instance_private_ip_addr)
        if subnet_cidr_block is not None:
            pulumi.set(__self__, "subnet_cidr_block", subnet_cidr_block)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Input[str]:
        """
        NAT gateway ID.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="publicIpAddrs")
    def public_ip_addrs(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Elastic IP address pool.
        """
        return pulumi.get(self, "public_ip_addrs")

    @public_ip_addrs.setter
    def public_ip_addrs(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "public_ip_addrs", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        """
        Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instancePrivateIpAddr")
    def instance_private_ip_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_private_ip_addr")

    @instance_private_ip_addr.setter
    def instance_private_ip_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_private_ip_addr", value)

    @property
    @pulumi.getter(name="subnetCidrBlock")
    def subnet_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_cidr_block")

    @subnet_cidr_block.setter
    def subnet_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_cidr_block", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet instance ID, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class _GatewaySnatState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_private_ip_addr: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 public_ip_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 snat_id: Optional[pulumi.Input[str]] = None,
                 subnet_cidr_block: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewaySnat resources.
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_id: Instance ID, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] instance_private_ip_addr: Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] nat_gateway_id: NAT gateway ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addrs: Elastic IP address pool.
        :param pulumi.Input[str] resource_type: Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        :param pulumi.Input[str] snat_id: SNAT rule ID.
        :param pulumi.Input[str] subnet_cidr_block: The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        :param pulumi.Input[str] subnet_id: Subnet instance ID, required when `resource_type` is SUBNET.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_private_ip_addr is not None:
            pulumi.set(__self__, "instance_private_ip_addr", instance_private_ip_addr)
        if nat_gateway_id is not None:
            pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if public_ip_addrs is not None:
            pulumi.set(__self__, "public_ip_addrs", public_ip_addrs)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if snat_id is not None:
            pulumi.set(__self__, "snat_id", snat_id)
        if subnet_cidr_block is not None:
            pulumi.set(__self__, "subnet_cidr_block", subnet_cidr_block)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instancePrivateIpAddr")
    def instance_private_ip_addr(self) -> Optional[pulumi.Input[str]]:
        """
        Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_private_ip_addr")

    @instance_private_ip_addr.setter
    def instance_private_ip_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_private_ip_addr", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        NAT gateway ID.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="publicIpAddrs")
    def public_ip_addrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Elastic IP address pool.
        """
        return pulumi.get(self, "public_ip_addrs")

    @public_ip_addrs.setter
    def public_ip_addrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "public_ip_addrs", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter(name="snatId")
    def snat_id(self) -> Optional[pulumi.Input[str]]:
        """
        SNAT rule ID.
        """
        return pulumi.get(self, "snat_id")

    @snat_id.setter
    def snat_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snat_id", value)

    @property
    @pulumi.getter(name="subnetCidrBlock")
    def subnet_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_cidr_block")

    @subnet_cidr_block.setter
    def subnet_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_cidr_block", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet instance ID, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class GatewaySnat(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_private_ip_addr: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 public_ip_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 subnet_cidr_block: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a NAT Gateway SNat rule.

        ## Import

        VPN gateway route can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}', resource_id range `subnet_id`, `instance_id`, e.g. SUBNET SNat

        ```sh
         $ pulumi import tencentcloud:Nat/gatewaySnat:GatewaySnat my_snat nat-r4ip1cwt#subnet-2ap74y35
        ```

         NETWORKINTERFACT SNat

        ```sh
         $ pulumi import tencentcloud:Nat/gatewaySnat:GatewaySnat my_snat nat-r4ip1cwt#ins-da412f5a
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_id: Instance ID, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] instance_private_ip_addr: Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] nat_gateway_id: NAT gateway ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addrs: Elastic IP address pool.
        :param pulumi.Input[str] resource_type: Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        :param pulumi.Input[str] subnet_cidr_block: The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        :param pulumi.Input[str] subnet_id: Subnet instance ID, required when `resource_type` is SUBNET.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewaySnatArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a NAT Gateway SNat rule.

        ## Import

        VPN gateway route can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}', resource_id range `subnet_id`, `instance_id`, e.g. SUBNET SNat

        ```sh
         $ pulumi import tencentcloud:Nat/gatewaySnat:GatewaySnat my_snat nat-r4ip1cwt#subnet-2ap74y35
        ```

         NETWORKINTERFACT SNat

        ```sh
         $ pulumi import tencentcloud:Nat/gatewaySnat:GatewaySnat my_snat nat-r4ip1cwt#ins-da412f5a
        ```

        :param str resource_name: The name of the resource.
        :param GatewaySnatArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewaySnatArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_private_ip_addr: Optional[pulumi.Input[str]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 public_ip_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 subnet_cidr_block: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewaySnatArgs.__new__(GatewaySnatArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["instance_private_ip_addr"] = instance_private_ip_addr
            if nat_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'nat_gateway_id'")
            __props__.__dict__["nat_gateway_id"] = nat_gateway_id
            if public_ip_addrs is None and not opts.urn:
                raise TypeError("Missing required property 'public_ip_addrs'")
            __props__.__dict__["public_ip_addrs"] = public_ip_addrs
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["subnet_cidr_block"] = subnet_cidr_block
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["snat_id"] = None
        super(GatewaySnat, __self__).__init__(
            'tencentcloud:Nat/gatewaySnat:GatewaySnat',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_private_ip_addr: Optional[pulumi.Input[str]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            public_ip_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_type: Optional[pulumi.Input[str]] = None,
            snat_id: Optional[pulumi.Input[str]] = None,
            subnet_cidr_block: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'GatewaySnat':
        """
        Get an existing GatewaySnat resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description.
        :param pulumi.Input[str] instance_id: Instance ID, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] instance_private_ip_addr: Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        :param pulumi.Input[str] nat_gateway_id: NAT gateway ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addrs: Elastic IP address pool.
        :param pulumi.Input[str] resource_type: Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        :param pulumi.Input[str] snat_id: SNAT rule ID.
        :param pulumi.Input[str] subnet_cidr_block: The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        :param pulumi.Input[str] subnet_id: Subnet instance ID, required when `resource_type` is SUBNET.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewaySnatState.__new__(_GatewaySnatState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_private_ip_addr"] = instance_private_ip_addr
        __props__.__dict__["nat_gateway_id"] = nat_gateway_id
        __props__.__dict__["public_ip_addrs"] = public_ip_addrs
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["snat_id"] = snat_id
        __props__.__dict__["subnet_cidr_block"] = subnet_cidr_block
        __props__.__dict__["subnet_id"] = subnet_id
        return GatewaySnat(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        Instance ID, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instancePrivateIpAddr")
    def instance_private_ip_addr(self) -> pulumi.Output[Optional[str]]:
        """
        Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
        """
        return pulumi.get(self, "instance_private_ip_addr")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[str]:
        """
        NAT gateway ID.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="publicIpAddrs")
    def public_ip_addrs(self) -> pulumi.Output[Sequence[str]]:
        """
        Elastic IP address pool.
        """
        return pulumi.get(self, "public_ip_addrs")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        """
        Resource type. Valid values: SUBNET, NETWORKINTERFACE.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="snatId")
    def snat_id(self) -> pulumi.Output[str]:
        """
        SNAT rule ID.
        """
        return pulumi.get(self, "snat_id")

    @property
    @pulumi.getter(name="subnetCidrBlock")
    def subnet_cidr_block(self) -> pulumi.Output[Optional[str]]:
        """
        The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_cidr_block")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        Subnet instance ID, required when `resource_type` is SUBNET.
        """
        return pulumi.get(self, "subnet_id")

