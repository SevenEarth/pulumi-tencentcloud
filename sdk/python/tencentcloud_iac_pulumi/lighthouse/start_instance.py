# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StartInstanceArgs', 'StartInstance']

@pulumi.input_type
class StartInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a StartInstance resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        pulumi.set(__self__, "instance_id", instance_id)

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


@pulumi.input_type
class _StartInstanceState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StartInstance resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

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


class StartInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a lighthouse start_instance

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        start_instance = tencentcloud.lighthouse.StartInstance("startInstance", instance_id="lhins-xxxxxx")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StartInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a lighthouse start_instance

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        start_instance = tencentcloud.lighthouse.StartInstance("startInstance", instance_id="lhins-xxxxxx")
        ```

        :param str resource_name: The name of the resource.
        :param StartInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StartInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
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
            __props__ = StartInstanceArgs.__new__(StartInstanceArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(StartInstance, __self__).__init__(
            'tencentcloud:Lighthouse/startInstance:StartInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'StartInstance':
        """
        Get an existing StartInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StartInstanceState.__new__(_StartInstanceState)

        __props__.__dict__["instance_id"] = instance_id
        return StartInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

