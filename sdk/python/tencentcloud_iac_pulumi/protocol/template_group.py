# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TemplateGroupArgs', 'TemplateGroup']

@pulumi.input_type
class TemplateGroupArgs:
    def __init__(__self__, *,
                 template_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TemplateGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Service template ID list.
        :param pulumi.Input[str] name: Name of the protocol template group.
        """
        pulumi.set(__self__, "template_ids", template_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Service template ID list.
        """
        return pulumi.get(self, "template_ids")

    @template_ids.setter
    def template_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "template_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the protocol template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TemplateGroupState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TemplateGroup resources.
        :param pulumi.Input[str] name: Name of the protocol template group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Service template ID list.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if template_ids is not None:
            pulumi.set(__self__, "template_ids", template_ids)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the protocol template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Service template ID list.
        """
        return pulumi.get(self, "template_ids")

    @template_ids.setter
    def template_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "template_ids", value)


class TemplateGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage protocol template group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.protocol.TemplateGroup("foo", template_ids=[
            "ipl-axaf24151",
            "ipl-axaf24152",
        ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Protocol template group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Protocol/templateGroup:TemplateGroup foo ppmg-0np3u974
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the protocol template group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Service template ID list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage protocol template group.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.protocol.TemplateGroup("foo", template_ids=[
            "ipl-axaf24151",
            "ipl-axaf24152",
        ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Protocol template group can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Protocol/templateGroup:TemplateGroup foo ppmg-0np3u974
        ```

        :param str resource_name: The name of the resource.
        :param TemplateGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateGroupArgs.__new__(TemplateGroupArgs)

            __props__.__dict__["name"] = name
            if template_ids is None and not opts.urn:
                raise TypeError("Missing required property 'template_ids'")
            __props__.__dict__["template_ids"] = template_ids
        super(TemplateGroup, __self__).__init__(
            'tencentcloud:Protocol/templateGroup:TemplateGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'TemplateGroup':
        """
        Get an existing TemplateGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the protocol template group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Service template ID list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemplateGroupState.__new__(_TemplateGroupState)

        __props__.__dict__["name"] = name
        __props__.__dict__["template_ids"] = template_ids
        return TemplateGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the protocol template group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Service template ID list.
        """
        return pulumi.get(self, "template_ids")

