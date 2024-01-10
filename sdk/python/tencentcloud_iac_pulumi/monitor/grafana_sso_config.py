# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GrafanaSsoConfigArgs', 'GrafanaSsoConfig']

@pulumi.input_type
class GrafanaSsoConfigArgs:
    def __init__(__self__, *,
                 enable_sso: pulumi.Input[bool],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a GrafanaSsoConfig resource.
        :param pulumi.Input[bool] enable_sso: Whether to enable SSO: `true` for enabling; `false` for disabling.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        pulumi.set(__self__, "enable_sso", enable_sso)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> pulumi.Input[bool]:
        """
        Whether to enable SSO: `true` for enabling; `false` for disabling.
        """
        return pulumi.get(self, "enable_sso")

    @enable_sso.setter
    def enable_sso(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_sso", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _GrafanaSsoConfigState:
    def __init__(__self__, *,
                 enable_sso: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GrafanaSsoConfig resources.
        :param pulumi.Input[bool] enable_sso: Whether to enable SSO: `true` for enabling; `false` for disabling.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        if enable_sso is not None:
            pulumi.set(__self__, "enable_sso", enable_sso)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable SSO: `true` for enabling; `false` for disabling.
        """
        return pulumi.get(self, "enable_sso")

    @enable_sso.setter
    def enable_sso(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_sso", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class GrafanaSsoConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_sso: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor grafana_sso_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        grafana_sso_config = tencentcloud.monitor.GrafanaSsoConfig("grafanaSsoConfig",
            enable_sso=False,
            instance_id="grafana-dp2hnnfa")
        ```

        ## Import

        monitor grafana_sso_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig grafana_sso_config instance_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_sso: Whether to enable SSO: `true` for enabling; `false` for disabling.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrafanaSsoConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor grafana_sso_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        grafana_sso_config = tencentcloud.monitor.GrafanaSsoConfig("grafanaSsoConfig",
            enable_sso=False,
            instance_id="grafana-dp2hnnfa")
        ```

        ## Import

        monitor grafana_sso_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig grafana_sso_config instance_id
        ```

        :param str resource_name: The name of the resource.
        :param GrafanaSsoConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrafanaSsoConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enable_sso: Optional[pulumi.Input[bool]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GrafanaSsoConfigArgs.__new__(GrafanaSsoConfigArgs)

            if enable_sso is None and not opts.urn:
                raise TypeError("Missing required property 'enable_sso'")
            __props__.__dict__["enable_sso"] = enable_sso
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(GrafanaSsoConfig, __self__).__init__(
            'tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enable_sso: Optional[pulumi.Input[bool]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'GrafanaSsoConfig':
        """
        Get an existing GrafanaSsoConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enable_sso: Whether to enable SSO: `true` for enabling; `false` for disabling.
        :param pulumi.Input[str] instance_id: Grafana instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GrafanaSsoConfigState.__new__(_GrafanaSsoConfigState)

        __props__.__dict__["enable_sso"] = enable_sso
        __props__.__dict__["instance_id"] = instance_id
        return GrafanaSsoConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> pulumi.Output[bool]:
        """
        Whether to enable SSO: `true` for enabling; `false` for disabling.
        """
        return pulumi.get(self, "enable_sso")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Grafana instance ID.
        """
        return pulumi.get(self, "instance_id")

