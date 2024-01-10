# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['KeyvalConfigArgs', 'KeyvalConfig']

@pulumi.input_type
class KeyvalConfigArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 items: pulumi.Input['KeyvalConfigItemsArgs']):
        """
        The set of arguments for constructing a KeyvalConfig resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input['KeyvalConfigItemsArgs'] items: configuration list.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def items(self) -> pulumi.Input['KeyvalConfigItemsArgs']:
        """
        configuration list.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: pulumi.Input['KeyvalConfigItemsArgs']):
        pulumi.set(self, "items", value)


@pulumi.input_type
class _KeyvalConfigState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input['KeyvalConfigItemsArgs']] = None):
        """
        Input properties used for looking up and filtering KeyvalConfig resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input['KeyvalConfigItemsArgs'] items: configuration list.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if items is not None:
            pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def items(self) -> Optional[pulumi.Input['KeyvalConfigItemsArgs']]:
        """
        configuration list.
        """
        return pulumi.get(self, "items")

    @items.setter
    def items(self, value: Optional[pulumi.Input['KeyvalConfigItemsArgs']]):
        pulumi.set(self, "items", value)


class KeyvalConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[pulumi.InputType['KeyvalConfigItemsArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a clickhouse keyval_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        keyval_config = tencentcloud.clickhouse.KeyvalConfig("keyvalConfig",
            instance_id="cdwch-datuhk3z",
            items=tencentcloud.clickhouse.KeyvalConfigItemsArgs(
                conf_key="max_open_files",
                conf_value="50000",
            ))
        ```

        ## Import

        clickhouse config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clickhouse/keyvalConfig:KeyvalConfig config cdwch-datuhk3z#max_open_files#50000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[pulumi.InputType['KeyvalConfigItemsArgs']] items: configuration list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyvalConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a clickhouse keyval_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        keyval_config = tencentcloud.clickhouse.KeyvalConfig("keyvalConfig",
            instance_id="cdwch-datuhk3z",
            items=tencentcloud.clickhouse.KeyvalConfigItemsArgs(
                conf_key="max_open_files",
                conf_value="50000",
            ))
        ```

        ## Import

        clickhouse config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clickhouse/keyvalConfig:KeyvalConfig config cdwch-datuhk3z#max_open_files#50000
        ```

        :param str resource_name: The name of the resource.
        :param KeyvalConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyvalConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 items: Optional[pulumi.Input[pulumi.InputType['KeyvalConfigItemsArgs']]] = None,
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
            __props__ = KeyvalConfigArgs.__new__(KeyvalConfigArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if items is None and not opts.urn:
                raise TypeError("Missing required property 'items'")
            __props__.__dict__["items"] = items
        super(KeyvalConfig, __self__).__init__(
            'tencentcloud:Clickhouse/keyvalConfig:KeyvalConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            items: Optional[pulumi.Input[pulumi.InputType['KeyvalConfigItemsArgs']]] = None) -> 'KeyvalConfig':
        """
        Get an existing KeyvalConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[pulumi.InputType['KeyvalConfigItemsArgs']] items: configuration list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyvalConfigState.__new__(_KeyvalConfigState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["items"] = items
        return KeyvalConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def items(self) -> pulumi.Output['outputs.KeyvalConfigItems']:
        """
        configuration list.
        """
        return pulumi.get(self, "items")

