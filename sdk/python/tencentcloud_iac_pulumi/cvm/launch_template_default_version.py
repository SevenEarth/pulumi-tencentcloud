# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LaunchTemplateDefaultVersionArgs', 'LaunchTemplateDefaultVersion']

@pulumi.input_type
class LaunchTemplateDefaultVersionArgs:
    def __init__(__self__, *,
                 default_version: pulumi.Input[int],
                 launch_template_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a LaunchTemplateDefaultVersion resource.
        :param pulumi.Input[int] default_version: The number of the version that you want to set as the default version.
        :param pulumi.Input[str] launch_template_id: Instance launch template ID.
        """
        pulumi.set(__self__, "default_version", default_version)
        pulumi.set(__self__, "launch_template_id", launch_template_id)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> pulumi.Input[int]:
        """
        The number of the version that you want to set as the default version.
        """
        return pulumi.get(self, "default_version")

    @default_version.setter
    def default_version(self, value: pulumi.Input[int]):
        pulumi.set(self, "default_version", value)

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> pulumi.Input[str]:
        """
        Instance launch template ID.
        """
        return pulumi.get(self, "launch_template_id")

    @launch_template_id.setter
    def launch_template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "launch_template_id", value)


@pulumi.input_type
class _LaunchTemplateDefaultVersionState:
    def __init__(__self__, *,
                 default_version: Optional[pulumi.Input[int]] = None,
                 launch_template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LaunchTemplateDefaultVersion resources.
        :param pulumi.Input[int] default_version: The number of the version that you want to set as the default version.
        :param pulumi.Input[str] launch_template_id: Instance launch template ID.
        """
        if default_version is not None:
            pulumi.set(__self__, "default_version", default_version)
        if launch_template_id is not None:
            pulumi.set(__self__, "launch_template_id", launch_template_id)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> Optional[pulumi.Input[int]]:
        """
        The number of the version that you want to set as the default version.
        """
        return pulumi.get(self, "default_version")

    @default_version.setter
    def default_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_version", value)

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance launch template ID.
        """
        return pulumi.get(self, "launch_template_id")

    @launch_template_id.setter
    def launch_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "launch_template_id", value)


class LaunchTemplateDefaultVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_version: Optional[pulumi.Input[int]] = None,
                 launch_template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cvm launch_template_default_version

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        launch_template_default_version = tencentcloud.cvm.LaunchTemplateDefaultVersion("launchTemplateDefaultVersion",
            default_version=2,
            launch_template_id="lt-34vaef8fe")
        ```

        ## Import

        cvm launch_template_default_version can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion launch_template_default_version launch_template_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_version: The number of the version that you want to set as the default version.
        :param pulumi.Input[str] launch_template_id: Instance launch template ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LaunchTemplateDefaultVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cvm launch_template_default_version

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        launch_template_default_version = tencentcloud.cvm.LaunchTemplateDefaultVersion("launchTemplateDefaultVersion",
            default_version=2,
            launch_template_id="lt-34vaef8fe")
        ```

        ## Import

        cvm launch_template_default_version can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion launch_template_default_version launch_template_id
        ```

        :param str resource_name: The name of the resource.
        :param LaunchTemplateDefaultVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LaunchTemplateDefaultVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_version: Optional[pulumi.Input[int]] = None,
                 launch_template_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LaunchTemplateDefaultVersionArgs.__new__(LaunchTemplateDefaultVersionArgs)

            if default_version is None and not opts.urn:
                raise TypeError("Missing required property 'default_version'")
            __props__.__dict__["default_version"] = default_version
            if launch_template_id is None and not opts.urn:
                raise TypeError("Missing required property 'launch_template_id'")
            __props__.__dict__["launch_template_id"] = launch_template_id
        super(LaunchTemplateDefaultVersion, __self__).__init__(
            'tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_version: Optional[pulumi.Input[int]] = None,
            launch_template_id: Optional[pulumi.Input[str]] = None) -> 'LaunchTemplateDefaultVersion':
        """
        Get an existing LaunchTemplateDefaultVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_version: The number of the version that you want to set as the default version.
        :param pulumi.Input[str] launch_template_id: Instance launch template ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LaunchTemplateDefaultVersionState.__new__(_LaunchTemplateDefaultVersionState)

        __props__.__dict__["default_version"] = default_version
        __props__.__dict__["launch_template_id"] = launch_template_id
        return LaunchTemplateDefaultVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultVersion")
    def default_version(self) -> pulumi.Output[int]:
        """
        The number of the version that you want to set as the default version.
        """
        return pulumi.get(self, "default_version")

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> pulumi.Output[str]:
        """
        Instance launch template ID.
        """
        return pulumi.get(self, "launch_template_id")

