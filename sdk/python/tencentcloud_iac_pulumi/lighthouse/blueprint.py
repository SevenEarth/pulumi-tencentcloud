# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BlueprintArgs', 'Blueprint']

@pulumi.input_type
class BlueprintArgs:
    def __init__(__self__, *,
                 blueprint_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Blueprint resource.
        :param pulumi.Input[str] blueprint_name: Blueprint name, which can contain up to 60 characters.
        :param pulumi.Input[str] description: Blueprint description, which can contain up to 60 characters.
        :param pulumi.Input[str] instance_id: ID of the instance for which to make a blueprint.
        """
        pulumi.set(__self__, "blueprint_name", blueprint_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="blueprintName")
    def blueprint_name(self) -> pulumi.Input[str]:
        """
        Blueprint name, which can contain up to 60 characters.
        """
        return pulumi.get(self, "blueprint_name")

    @blueprint_name.setter
    def blueprint_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "blueprint_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Blueprint description, which can contain up to 60 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the instance for which to make a blueprint.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _BlueprintState:
    def __init__(__self__, *,
                 blueprint_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Blueprint resources.
        :param pulumi.Input[str] blueprint_name: Blueprint name, which can contain up to 60 characters.
        :param pulumi.Input[str] description: Blueprint description, which can contain up to 60 characters.
        :param pulumi.Input[str] instance_id: ID of the instance for which to make a blueprint.
        """
        if blueprint_name is not None:
            pulumi.set(__self__, "blueprint_name", blueprint_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="blueprintName")
    def blueprint_name(self) -> Optional[pulumi.Input[str]]:
        """
        Blueprint name, which can contain up to 60 characters.
        """
        return pulumi.get(self, "blueprint_name")

    @blueprint_name.setter
    def blueprint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blueprint_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Blueprint description, which can contain up to 60 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the instance for which to make a blueprint.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class Blueprint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blueprint_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a lighthouse blueprint

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        blueprint = tencentcloud.lighthouse.Blueprint("blueprint",
            blueprint_name="blueprint_name_test",
            description="blueprint_description_test",
            instance_id="lhins-xxxxxx")
        ```

        ## Import

        lighthouse blueprint can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Lighthouse/blueprint:Blueprint blueprint blueprint_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blueprint_name: Blueprint name, which can contain up to 60 characters.
        :param pulumi.Input[str] description: Blueprint description, which can contain up to 60 characters.
        :param pulumi.Input[str] instance_id: ID of the instance for which to make a blueprint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlueprintArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a lighthouse blueprint

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        blueprint = tencentcloud.lighthouse.Blueprint("blueprint",
            blueprint_name="blueprint_name_test",
            description="blueprint_description_test",
            instance_id="lhins-xxxxxx")
        ```

        ## Import

        lighthouse blueprint can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Lighthouse/blueprint:Blueprint blueprint blueprint_id
        ```

        :param str resource_name: The name of the resource.
        :param BlueprintArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlueprintArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blueprint_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
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
            __props__ = BlueprintArgs.__new__(BlueprintArgs)

            if blueprint_name is None and not opts.urn:
                raise TypeError("Missing required property 'blueprint_name'")
            __props__.__dict__["blueprint_name"] = blueprint_name
            __props__.__dict__["description"] = description
            __props__.__dict__["instance_id"] = instance_id
        super(Blueprint, __self__).__init__(
            'tencentcloud:Lighthouse/blueprint:Blueprint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blueprint_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Blueprint':
        """
        Get an existing Blueprint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blueprint_name: Blueprint name, which can contain up to 60 characters.
        :param pulumi.Input[str] description: Blueprint description, which can contain up to 60 characters.
        :param pulumi.Input[str] instance_id: ID of the instance for which to make a blueprint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlueprintState.__new__(_BlueprintState)

        __props__.__dict__["blueprint_name"] = blueprint_name
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        return Blueprint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blueprintName")
    def blueprint_name(self) -> pulumi.Output[str]:
        """
        Blueprint name, which can contain up to 60 characters.
        """
        return pulumi.get(self, "blueprint_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Blueprint description, which can contain up to 60 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the instance for which to make a blueprint.
        """
        return pulumi.get(self, "instance_id")

