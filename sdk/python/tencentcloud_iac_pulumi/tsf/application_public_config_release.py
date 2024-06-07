# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationPublicConfigReleaseArgs', 'ApplicationPublicConfigRelease']

@pulumi.input_type
class ApplicationPublicConfigReleaseArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 namespace_id: pulumi.Input[str],
                 release_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationPublicConfigRelease resource.
        :param pulumi.Input[str] config_id: ConfigId.
        :param pulumi.Input[str] namespace_id: namespace-id.
        :param pulumi.Input[str] release_desc: Release description.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "namespace_id", namespace_id)
        if release_desc is not None:
            pulumi.set(__self__, "release_desc", release_desc)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        """
        ConfigId.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Input[str]:
        """
        namespace-id.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Release description.
        """
        return pulumi.get(self, "release_desc")

    @release_desc.setter
    def release_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_desc", value)


@pulumi.input_type
class _ApplicationPublicConfigReleaseState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationPublicConfigRelease resources.
        :param pulumi.Input[str] config_id: ConfigId.
        :param pulumi.Input[str] namespace_id: namespace-id.
        :param pulumi.Input[str] release_desc: Release description.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if release_desc is not None:
            pulumi.set(__self__, "release_desc", release_desc)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        ConfigId.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        namespace-id.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Release description.
        """
        return pulumi.get(self, "release_desc")

    @release_desc.setter
    def release_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_desc", value)


class ApplicationPublicConfigRelease(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf application_public_config_release

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_public_config_release = tencentcloud.tsf.ApplicationPublicConfigRelease("applicationPublicConfigRelease",
            config_id="dcfg-p-123456",
            namespace_id="namespace-123456",
            release_desc="product version")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf application_public_config_release can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease application_public_config_release application_public_config_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: ConfigId.
        :param pulumi.Input[str] namespace_id: namespace-id.
        :param pulumi.Input[str] release_desc: Release description.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationPublicConfigReleaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf application_public_config_release

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_public_config_release = tencentcloud.tsf.ApplicationPublicConfigRelease("applicationPublicConfigRelease",
            config_id="dcfg-p-123456",
            namespace_id="namespace-123456",
            release_desc="product version")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf application_public_config_release can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease application_public_config_release application_public_config_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationPublicConfigReleaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationPublicConfigReleaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationPublicConfigReleaseArgs.__new__(ApplicationPublicConfigReleaseArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_id'")
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["release_desc"] = release_desc
        super(ApplicationPublicConfigRelease, __self__).__init__(
            'tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            release_desc: Optional[pulumi.Input[str]] = None) -> 'ApplicationPublicConfigRelease':
        """
        Get an existing ApplicationPublicConfigRelease resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: ConfigId.
        :param pulumi.Input[str] namespace_id: namespace-id.
        :param pulumi.Input[str] release_desc: Release description.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationPublicConfigReleaseState.__new__(_ApplicationPublicConfigReleaseState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["release_desc"] = release_desc
        return ApplicationPublicConfigRelease(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        ConfigId.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        namespace-id.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> pulumi.Output[Optional[str]]:
        """
        Release description.
        """
        return pulumi.get(self, "release_desc")

