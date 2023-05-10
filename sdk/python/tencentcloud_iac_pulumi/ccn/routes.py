# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoutesArgs', 'Routes']

@pulumi.input_type
class RoutesArgs:
    def __init__(__self__, *,
                 ccn_id: pulumi.Input[str],
                 route_id: pulumi.Input[str],
                 switch: pulumi.Input[str]):
        """
        The set of arguments for constructing a Routes resource.
        :param pulumi.Input[str] ccn_id: CCN Instance ID.
        :param pulumi.Input[str] route_id: CCN Route Id List.
        :param pulumi.Input[str] switch: `on`: Enable, `off`: Disable.
        """
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "route_id", route_id)
        pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Input[str]:
        """
        CCN Instance ID.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Input[str]:
        """
        CCN Route Id List.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Input[str]:
        """
        `on`: Enable, `off`: Disable.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: pulumi.Input[str]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class _RoutesState:
    def __init__(__self__, *,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Routes resources.
        :param pulumi.Input[str] ccn_id: CCN Instance ID.
        :param pulumi.Input[str] route_id: CCN Route Id List.
        :param pulumi.Input[str] switch: `on`: Enable, `off`: Disable.
        """
        if ccn_id is not None:
            pulumi.set(__self__, "ccn_id", ccn_id)
        if route_id is not None:
            pulumi.set(__self__, "route_id", route_id)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> Optional[pulumi.Input[str]]:
        """
        CCN Instance ID.
        """
        return pulumi.get(self, "ccn_id")

    @ccn_id.setter
    def ccn_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ccn_id", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> Optional[pulumi.Input[str]]:
        """
        CCN Route Id List.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        """
        `on`: Enable, `off`: Disable.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


class Routes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc ccn_routes

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        ccn_routes = tencentcloud.ccn.Routes("ccnRoutes",
            ccn_id="ccn-39lqkygf",
            route_id="ccnr-3o0dfyuw",
            switch="on")
        ```

        ## Import

        vpc ccn_routes can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ccn/routes:Routes ccn_routes ccnId#routesId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ccn_id: CCN Instance ID.
        :param pulumi.Input[str] route_id: CCN Route Id List.
        :param pulumi.Input[str] switch: `on`: Enable, `off`: Disable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoutesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc ccn_routes

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        ccn_routes = tencentcloud.ccn.Routes("ccnRoutes",
            ccn_id="ccn-39lqkygf",
            route_id="ccnr-3o0dfyuw",
            switch="on")
        ```

        ## Import

        vpc ccn_routes can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ccn/routes:Routes ccn_routes ccnId#routesId
        ```

        :param str resource_name: The name of the resource.
        :param RoutesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoutesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ccn_id: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoutesArgs.__new__(RoutesArgs)

            if ccn_id is None and not opts.urn:
                raise TypeError("Missing required property 'ccn_id'")
            __props__.__dict__["ccn_id"] = ccn_id
            if route_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_id'")
            __props__.__dict__["route_id"] = route_id
            if switch is None and not opts.urn:
                raise TypeError("Missing required property 'switch'")
            __props__.__dict__["switch"] = switch
        super(Routes, __self__).__init__(
            'tencentcloud:Ccn/routes:Routes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ccn_id: Optional[pulumi.Input[str]] = None,
            route_id: Optional[pulumi.Input[str]] = None,
            switch: Optional[pulumi.Input[str]] = None) -> 'Routes':
        """
        Get an existing Routes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ccn_id: CCN Instance ID.
        :param pulumi.Input[str] route_id: CCN Route Id List.
        :param pulumi.Input[str] switch: `on`: Enable, `off`: Disable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoutesState.__new__(_RoutesState)

        __props__.__dict__["ccn_id"] = ccn_id
        __props__.__dict__["route_id"] = route_id
        __props__.__dict__["switch"] = switch
        return Routes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> pulumi.Output[str]:
        """
        CCN Instance ID.
        """
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[str]:
        """
        CCN Route Id List.
        """
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Output[str]:
        """
        `on`: Enable, `off`: Disable.
        """
        return pulumi.get(self, "switch")

