# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewayCcnRoutesArgs', 'GatewayCcnRoutes']

@pulumi.input_type
class GatewayCcnRoutesArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[str],
                 route_id: pulumi.Input[str],
                 status: pulumi.Input[str],
                 vpn_gateway_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a GatewayCcnRoutes resource.
        :param pulumi.Input[str] destination_cidr_block: Routing CIDR.
        :param pulumi.Input[str] route_id: Route Id.
        :param pulumi.Input[str] status: Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        :param pulumi.Input[str] vpn_gateway_id: VPN GATEWAY INSTANCE ID.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "route_id", route_id)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        Routing CIDR.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Input[str]:
        """
        Route Id.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Input[str]:
        """
        VPN GATEWAY INSTANCE ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_gateway_id", value)


@pulumi.input_type
class _GatewayCcnRoutesState:
    def __init__(__self__, *,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayCcnRoutes resources.
        :param pulumi.Input[str] destination_cidr_block: Routing CIDR.
        :param pulumi.Input[str] route_id: Route Id.
        :param pulumi.Input[str] status: Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        :param pulumi.Input[str] vpn_gateway_id: VPN GATEWAY INSTANCE ID.
        """
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if route_id is not None:
            pulumi.set(__self__, "route_id", route_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        Routing CIDR.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> Optional[pulumi.Input[str]]:
        """
        Route Id.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPN GATEWAY INSTANCE ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)


class GatewayCcnRoutes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpn_gateway_ccn_routes

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        vpn_gateway_ccn_routes = tencentcloud.vpn.GatewayCcnRoutes("vpnGatewayCcnRoutes",
            destination_cidr_block="192.168.1.0/24",
            route_id="vpnr-akdy0757",
            status="DISABLE",
            vpn_gateway_id="vpngw-lie1a4u7")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        vpc vpn_gateway_ccn_routes can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Vpn/gatewayCcnRoutes:GatewayCcnRoutes vpn_gateway_ccn_routes vpn_gateway_id#ccn_routes_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: Routing CIDR.
        :param pulumi.Input[str] route_id: Route Id.
        :param pulumi.Input[str] status: Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        :param pulumi.Input[str] vpn_gateway_id: VPN GATEWAY INSTANCE ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayCcnRoutesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpn_gateway_ccn_routes

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        vpn_gateway_ccn_routes = tencentcloud.vpn.GatewayCcnRoutes("vpnGatewayCcnRoutes",
            destination_cidr_block="192.168.1.0/24",
            route_id="vpnr-akdy0757",
            status="DISABLE",
            vpn_gateway_id="vpngw-lie1a4u7")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        vpc vpn_gateway_ccn_routes can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Vpn/gatewayCcnRoutes:GatewayCcnRoutes vpn_gateway_ccn_routes vpn_gateway_id#ccn_routes_id
        ```

        :param str resource_name: The name of the resource.
        :param GatewayCcnRoutesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayCcnRoutesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayCcnRoutesArgs.__new__(GatewayCcnRoutesArgs)

            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            if route_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_id'")
            __props__.__dict__["route_id"] = route_id
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            if vpn_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        super(GatewayCcnRoutes, __self__).__init__(
            'tencentcloud:Vpn/gatewayCcnRoutes:GatewayCcnRoutes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            route_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None) -> 'GatewayCcnRoutes':
        """
        Get an existing GatewayCcnRoutes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: Routing CIDR.
        :param pulumi.Input[str] route_id: Route Id.
        :param pulumi.Input[str] status: Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        :param pulumi.Input[str] vpn_gateway_id: VPN GATEWAY INSTANCE ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayCcnRoutesState.__new__(_GatewayCcnRoutesState)

        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["route_id"] = route_id
        __props__.__dict__["status"] = status
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        return GatewayCcnRoutes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        Routing CIDR.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[str]:
        """
        Route Id.
        """
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Whether routing information is enabled. `ENABLE`: Enable Route, `DISABLE`: Disable Route.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        VPN GATEWAY INSTANCE ID.
        """
        return pulumi.get(self, "vpn_gateway_id")

