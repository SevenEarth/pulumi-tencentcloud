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

__all__ = ['CustomHeaderArgs', 'CustomHeader']

@pulumi.input_type
class CustomHeaderArgs:
    def __init__(__self__, *,
                 rule_id: pulumi.Input[str],
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]] = None):
        """
        The set of arguments for constructing a CustomHeader resource.
        :param pulumi.Input[str] rule_id: Rule id.
        :param pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]] headers: Headers.
        """
        pulumi.set(__self__, "rule_id", rule_id)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        """
        Rule id.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]]:
        """
        Headers.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]]):
        pulumi.set(self, "headers", value)


@pulumi.input_type
class _CustomHeaderState:
    def __init__(__self__, *,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomHeader resources.
        :param pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]] headers: Headers.
        :param pulumi.Input[str] rule_id: Rule id.
        """
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]]:
        """
        Headers.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomHeaderHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule id.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)


class CustomHeader(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomHeaderHeaderArgs']]]]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a gaap custom_header

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        custom_header = tencentcloud.gaap.CustomHeader("customHeader",
            headers=[
                tencentcloud.gaap.CustomHeaderHeaderArgs(
                    header_name="HeaderName1",
                    header_value="HeaderValue1",
                ),
                tencentcloud.gaap.CustomHeaderHeaderArgs(
                    header_name="HeaderName2",
                    header_value="HeaderValue2",
                ),
            ],
            rule_id="rule-xxxxxx")
        ```

        ## Import

        gaap custom_header can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/customHeader:CustomHeader custom_header ruleId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomHeaderHeaderArgs']]]] headers: Headers.
        :param pulumi.Input[str] rule_id: Rule id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomHeaderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a gaap custom_header

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        custom_header = tencentcloud.gaap.CustomHeader("customHeader",
            headers=[
                tencentcloud.gaap.CustomHeaderHeaderArgs(
                    header_name="HeaderName1",
                    header_value="HeaderValue1",
                ),
                tencentcloud.gaap.CustomHeaderHeaderArgs(
                    header_name="HeaderName2",
                    header_value="HeaderValue2",
                ),
            ],
            rule_id="rule-xxxxxx")
        ```

        ## Import

        gaap custom_header can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/customHeader:CustomHeader custom_header ruleId
        ```

        :param str resource_name: The name of the resource.
        :param CustomHeaderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomHeaderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomHeaderHeaderArgs']]]]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CustomHeaderArgs.__new__(CustomHeaderArgs)

            __props__.__dict__["headers"] = headers
            if rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'rule_id'")
            __props__.__dict__["rule_id"] = rule_id
        super(CustomHeader, __self__).__init__(
            'tencentcloud:Gaap/customHeader:CustomHeader',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomHeaderHeaderArgs']]]]] = None,
            rule_id: Optional[pulumi.Input[str]] = None) -> 'CustomHeader':
        """
        Get an existing CustomHeader resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CustomHeaderHeaderArgs']]]] headers: Headers.
        :param pulumi.Input[str] rule_id: Rule id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomHeaderState.__new__(_CustomHeaderState)

        __props__.__dict__["headers"] = headers
        __props__.__dict__["rule_id"] = rule_id
        return CustomHeader(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def headers(self) -> pulumi.Output[Optional[Sequence['outputs.CustomHeaderHeader']]]:
        """
        Headers.
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        Rule id.
        """
        return pulumi.get(self, "rule_id")

