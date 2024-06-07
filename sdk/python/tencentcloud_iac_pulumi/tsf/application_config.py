# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationConfigArgs', 'ApplicationConfig']

@pulumi.input_type
class ApplicationConfigArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 config_name: pulumi.Input[str],
                 config_value: pulumi.Input[str],
                 config_version: pulumi.Input[str],
                 config_type: Optional[pulumi.Input[str]] = None,
                 config_version_desc: Optional[pulumi.Input[str]] = None,
                 encode_with_base64: Optional[pulumi.Input[bool]] = None,
                 program_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ApplicationConfig resource.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_value: configuration item value.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] config_type: configuration item value type.
        :param pulumi.Input[str] config_version_desc: configuration item version description.
        :param pulumi.Input[bool] encode_with_base64: Base64 encoded configuration items.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] program_id_lists: Program id list.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "config_name", config_name)
        pulumi.set(__self__, "config_value", config_value)
        pulumi.set(__self__, "config_version", config_version)
        if config_type is not None:
            pulumi.set(__self__, "config_type", config_type)
        if config_version_desc is not None:
            pulumi.set(__self__, "config_version_desc", config_version_desc)
        if encode_with_base64 is not None:
            pulumi.set(__self__, "encode_with_base64", encode_with_base64)
        if program_id_lists is not None:
            pulumi.set(__self__, "program_id_lists", program_id_lists)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        Application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> pulumi.Input[str]:
        """
        configuration item name.
        """
        return pulumi.get(self, "config_name")

    @config_name.setter
    def config_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_name", value)

    @property
    @pulumi.getter(name="configValue")
    def config_value(self) -> pulumi.Input[str]:
        """
        configuration item value.
        """
        return pulumi.get(self, "config_value")

    @config_value.setter
    def config_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_value", value)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> pulumi.Input[str]:
        """
        configuration item version.
        """
        return pulumi.get(self, "config_version")

    @config_version.setter
    def config_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_version", value)

    @property
    @pulumi.getter(name="configType")
    def config_type(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item value type.
        """
        return pulumi.get(self, "config_type")

    @config_type.setter
    def config_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_type", value)

    @property
    @pulumi.getter(name="configVersionDesc")
    def config_version_desc(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item version description.
        """
        return pulumi.get(self, "config_version_desc")

    @config_version_desc.setter
    def config_version_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version_desc", value)

    @property
    @pulumi.getter(name="encodeWithBase64")
    def encode_with_base64(self) -> Optional[pulumi.Input[bool]]:
        """
        Base64 encoded configuration items.
        """
        return pulumi.get(self, "encode_with_base64")

    @encode_with_base64.setter
    def encode_with_base64(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encode_with_base64", value)

    @property
    @pulumi.getter(name="programIdLists")
    def program_id_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Program id list.
        """
        return pulumi.get(self, "program_id_lists")

    @program_id_lists.setter
    def program_id_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "program_id_lists", value)


@pulumi.input_type
class _ApplicationConfigState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 config_type: Optional[pulumi.Input[str]] = None,
                 config_value: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 config_version_desc: Optional[pulumi.Input[str]] = None,
                 encode_with_base64: Optional[pulumi.Input[bool]] = None,
                 program_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ApplicationConfig resources.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_type: configuration item value type.
        :param pulumi.Input[str] config_value: configuration item value.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] config_version_desc: configuration item version description.
        :param pulumi.Input[bool] encode_with_base64: Base64 encoded configuration items.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] program_id_lists: Program id list.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if config_name is not None:
            pulumi.set(__self__, "config_name", config_name)
        if config_type is not None:
            pulumi.set(__self__, "config_type", config_type)
        if config_value is not None:
            pulumi.set(__self__, "config_value", config_value)
        if config_version is not None:
            pulumi.set(__self__, "config_version", config_version)
        if config_version_desc is not None:
            pulumi.set(__self__, "config_version_desc", config_version_desc)
        if encode_with_base64 is not None:
            pulumi.set(__self__, "encode_with_base64", encode_with_base64)
        if program_id_lists is not None:
            pulumi.set(__self__, "program_id_lists", program_id_lists)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item name.
        """
        return pulumi.get(self, "config_name")

    @config_name.setter
    def config_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_name", value)

    @property
    @pulumi.getter(name="configType")
    def config_type(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item value type.
        """
        return pulumi.get(self, "config_type")

    @config_type.setter
    def config_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_type", value)

    @property
    @pulumi.getter(name="configValue")
    def config_value(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item value.
        """
        return pulumi.get(self, "config_value")

    @config_value.setter
    def config_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_value", value)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item version.
        """
        return pulumi.get(self, "config_version")

    @config_version.setter
    def config_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version", value)

    @property
    @pulumi.getter(name="configVersionDesc")
    def config_version_desc(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item version description.
        """
        return pulumi.get(self, "config_version_desc")

    @config_version_desc.setter
    def config_version_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version_desc", value)

    @property
    @pulumi.getter(name="encodeWithBase64")
    def encode_with_base64(self) -> Optional[pulumi.Input[bool]]:
        """
        Base64 encoded configuration items.
        """
        return pulumi.get(self, "encode_with_base64")

    @encode_with_base64.setter
    def encode_with_base64(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encode_with_base64", value)

    @property
    @pulumi.getter(name="programIdLists")
    def program_id_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Program id list.
        """
        return pulumi.get(self, "program_id_lists")

    @program_id_lists.setter
    def program_id_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "program_id_lists", value)


class ApplicationConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 config_type: Optional[pulumi.Input[str]] = None,
                 config_value: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 config_version_desc: Optional[pulumi.Input[str]] = None,
                 encode_with_base64: Optional[pulumi.Input[bool]] = None,
                 program_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf application_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_config = tencentcloud.tsf.ApplicationConfig("applicationConfig",
            application_id="application-ym9mxmza",
            config_name="test-2",
            config_value="name: \\"name\\"",
            config_version="1.0",
            config_version_desc="test2",
            encode_with_base64=False)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_type: configuration item value type.
        :param pulumi.Input[str] config_value: configuration item value.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] config_version_desc: configuration item version description.
        :param pulumi.Input[bool] encode_with_base64: Base64 encoded configuration items.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] program_id_lists: Program id list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf application_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_config = tencentcloud.tsf.ApplicationConfig("applicationConfig",
            application_id="application-ym9mxmza",
            config_name="test-2",
            config_value="name: \\"name\\"",
            config_version="1.0",
            config_version_desc="test2",
            encode_with_base64=False)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param ApplicationConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 config_type: Optional[pulumi.Input[str]] = None,
                 config_value: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 config_version_desc: Optional[pulumi.Input[str]] = None,
                 encode_with_base64: Optional[pulumi.Input[bool]] = None,
                 program_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationConfigArgs.__new__(ApplicationConfigArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if config_name is None and not opts.urn:
                raise TypeError("Missing required property 'config_name'")
            __props__.__dict__["config_name"] = config_name
            __props__.__dict__["config_type"] = config_type
            if config_value is None and not opts.urn:
                raise TypeError("Missing required property 'config_value'")
            __props__.__dict__["config_value"] = config_value
            if config_version is None and not opts.urn:
                raise TypeError("Missing required property 'config_version'")
            __props__.__dict__["config_version"] = config_version
            __props__.__dict__["config_version_desc"] = config_version_desc
            __props__.__dict__["encode_with_base64"] = encode_with_base64
            __props__.__dict__["program_id_lists"] = program_id_lists
        super(ApplicationConfig, __self__).__init__(
            'tencentcloud:Tsf/applicationConfig:ApplicationConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            config_name: Optional[pulumi.Input[str]] = None,
            config_type: Optional[pulumi.Input[str]] = None,
            config_value: Optional[pulumi.Input[str]] = None,
            config_version: Optional[pulumi.Input[str]] = None,
            config_version_desc: Optional[pulumi.Input[str]] = None,
            encode_with_base64: Optional[pulumi.Input[bool]] = None,
            program_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ApplicationConfig':
        """
        Get an existing ApplicationConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_type: configuration item value type.
        :param pulumi.Input[str] config_value: configuration item value.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] config_version_desc: configuration item version description.
        :param pulumi.Input[bool] encode_with_base64: Base64 encoded configuration items.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] program_id_lists: Program id list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationConfigState.__new__(_ApplicationConfigState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["config_name"] = config_name
        __props__.__dict__["config_type"] = config_type
        __props__.__dict__["config_value"] = config_value
        __props__.__dict__["config_version"] = config_version
        __props__.__dict__["config_version_desc"] = config_version_desc
        __props__.__dict__["encode_with_base64"] = encode_with_base64
        __props__.__dict__["program_id_lists"] = program_id_lists
        return ApplicationConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        Application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> pulumi.Output[str]:
        """
        configuration item name.
        """
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter(name="configType")
    def config_type(self) -> pulumi.Output[Optional[str]]:
        """
        configuration item value type.
        """
        return pulumi.get(self, "config_type")

    @property
    @pulumi.getter(name="configValue")
    def config_value(self) -> pulumi.Output[str]:
        """
        configuration item value.
        """
        return pulumi.get(self, "config_value")

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> pulumi.Output[str]:
        """
        configuration item version.
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter(name="configVersionDesc")
    def config_version_desc(self) -> pulumi.Output[Optional[str]]:
        """
        configuration item version description.
        """
        return pulumi.get(self, "config_version_desc")

    @property
    @pulumi.getter(name="encodeWithBase64")
    def encode_with_base64(self) -> pulumi.Output[Optional[bool]]:
        """
        Base64 encoded configuration items.
        """
        return pulumi.get(self, "encode_with_base64")

    @property
    @pulumi.getter(name="programIdLists")
    def program_id_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Program id list.
        """
        return pulumi.get(self, "program_id_lists")

