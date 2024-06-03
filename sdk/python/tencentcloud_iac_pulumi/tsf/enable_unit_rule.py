# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnableUnitRuleArgs', 'EnableUnitRule']

@pulumi.input_type
class EnableUnitRuleArgs:
    def __init__(__self__, *,
                 rule_id: pulumi.Input[str],
                 switch: pulumi.Input[str]):
        """
        The set of arguments for constructing a EnableUnitRule resource.
        :param pulumi.Input[str] rule_id: api ID.
        :param pulumi.Input[str] switch: switch, on: `enabled`, off: `disabled`.
        """
        pulumi.set(__self__, "rule_id", rule_id)
        pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        """
        api ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Input[str]:
        """
        switch, on: `enabled`, off: `disabled`.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: pulumi.Input[str]):
        pulumi.set(self, "switch", value)


@pulumi.input_type
class _EnableUnitRuleState:
    def __init__(__self__, *,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnableUnitRule resources.
        :param pulumi.Input[str] rule_id: api ID.
        :param pulumi.Input[str] switch: switch, on: `enabled`, off: `disabled`.
        """
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        api ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter
    def switch(self) -> Optional[pulumi.Input[str]]:
        """
        switch, on: `enabled`, off: `disabled`.
        """
        return pulumi.get(self, "switch")

    @switch.setter
    def switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "switch", value)


class EnableUnitRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf enable_unit_rule

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        enable_unit_rule = tencentcloud.tsf.EnableUnitRule("enableUnitRule",
            rule_id="unit-rl-is9m4nxz",
            switch="enabled")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf enable_unit_rule can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/enableUnitRule:EnableUnitRule enable_unit_rule enable_unit_rule_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] rule_id: api ID.
        :param pulumi.Input[str] switch: switch, on: `enabled`, off: `disabled`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnableUnitRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf enable_unit_rule

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        enable_unit_rule = tencentcloud.tsf.EnableUnitRule("enableUnitRule",
            rule_id="unit-rl-is9m4nxz",
            switch="enabled")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf enable_unit_rule can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/enableUnitRule:EnableUnitRule enable_unit_rule enable_unit_rule_id
        ```

        :param str resource_name: The name of the resource.
        :param EnableUnitRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnableUnitRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 switch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnableUnitRuleArgs.__new__(EnableUnitRuleArgs)

            if rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'rule_id'")
            __props__.__dict__["rule_id"] = rule_id
            if switch is None and not opts.urn:
                raise TypeError("Missing required property 'switch'")
            __props__.__dict__["switch"] = switch
        super(EnableUnitRule, __self__).__init__(
            'tencentcloud:Tsf/enableUnitRule:EnableUnitRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rule_id: Optional[pulumi.Input[str]] = None,
            switch: Optional[pulumi.Input[str]] = None) -> 'EnableUnitRule':
        """
        Get an existing EnableUnitRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] rule_id: api ID.
        :param pulumi.Input[str] switch: switch, on: `enabled`, off: `disabled`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnableUnitRuleState.__new__(_EnableUnitRuleState)

        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["switch"] = switch
        return EnableUnitRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        api ID.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter
    def switch(self) -> pulumi.Output[str]:
        """
        switch, on: `enabled`, off: `disabled`.
        """
        return pulumi.get(self, "switch")

