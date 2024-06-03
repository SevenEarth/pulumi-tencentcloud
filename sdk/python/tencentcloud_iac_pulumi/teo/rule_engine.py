# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RuleEngineArgs', 'RuleEngine']

@pulumi.input_type
class RuleEngineArgs:
    def __init__(__self__, *,
                 rule_name: pulumi.Input[str],
                 rules: pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]],
                 status: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RuleEngine resource.
        :param pulumi.Input[str] rule_name: The rule name (1 to 255 characters).
        :param pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]] rules: Rule items list.
        :param pulumi.Input[str] status: Rule status. Values:
        :param pulumi.Input[str] zone_id: ID of the site.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: rule tag list.
        """
        pulumi.set(__self__, "rule_name", rule_name)
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "zone_id", zone_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        """
        The rule name (1 to 255 characters).
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]]:
        """
        Rule items list.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Rule status. Values:
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        rule tag list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RuleEngineState:
    def __init__(__self__, *,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RuleEngine resources.
        :param pulumi.Input[str] rule_id: Rule ID.
        :param pulumi.Input[str] rule_name: The rule name (1 to 255 characters).
        :param pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]] rules: Rule items list.
        :param pulumi.Input[str] status: Rule status. Values:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: rule tag list.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Rule ID.
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        The rule name (1 to 255 characters).
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]]]:
        """
        Rule items list.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleEngineRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Rule status. Values:
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        rule tag list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class RuleEngine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleEngineRuleArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a teo rule_engine

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        rule1 = tencentcloud.teo.RuleEngine("rule1",
            zone_id=tencentcloud_teo_zone["example"]["id"],
            rule_name="test-rule",
            status="disable",
            rules=[tencentcloud.teo.RuleEngineRuleArgs(
                actions=[tencentcloud.teo.RuleEngineRuleActionArgs(
                    normal_action=tencentcloud.teo.RuleEngineRuleActionNormalActionArgs(
                        action="UpstreamUrlRedirect",
                        parameters=[
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Type",
                                values=["Path"],
                            ),
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Action",
                                values=["addPrefix"],
                            ),
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Value",
                                values=["/sss"],
                            ),
                        ],
                    ),
                )],
                ors=[
                    tencentcloud.teo.RuleEngineRuleOrArgs(
                        ands=[
                            tencentcloud.teo.RuleEngineRuleOrAndArgs(
                                operator="equal",
                                target="host",
                                ignore_case=False,
                                values=["a.tf-teo-t.xyz"],
                            ),
                            tencentcloud.teo.RuleEngineRuleOrAndArgs(
                                operator="equal",
                                target="extension",
                                ignore_case=False,
                                values=["jpg"],
                            ),
                        ],
                    ),
                    tencentcloud.teo.RuleEngineRuleOrArgs(
                        ands=[tencentcloud.teo.RuleEngineRuleOrAndArgs(
                            operator="equal",
                            target="filename",
                            ignore_case=False,
                            values=["test.txt"],
                        )],
                    ),
                ],
                sub_rules=[tencentcloud.teo.RuleEngineRuleSubRuleArgs(
                    tags=["png"],
                    rules=[tencentcloud.teo.RuleEngineRuleSubRuleRuleArgs(
                        ors=[
                            tencentcloud.teo.RuleEngineRuleSubRuleRuleOrArgs(
                                ands=[
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                        operator="notequal",
                                        target="host",
                                        ignore_case=False,
                                        values=["a.tf-teo-t.xyz"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                        operator="equal",
                                        target="extension",
                                        ignore_case=False,
                                        values=["png"],
                                    ),
                                ],
                            ),
                            tencentcloud.teo.RuleEngineRuleSubRuleRuleOrArgs(
                                ands=[tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                    operator="notequal",
                                    target="filename",
                                    ignore_case=False,
                                    values=["test.txt"],
                                )],
                            ),
                        ],
                        actions=[tencentcloud.teo.RuleEngineRuleSubRuleRuleActionArgs(
                            normal_action=tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionArgs(
                                action="UpstreamUrlRedirect",
                                parameters=[
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Type",
                                        values=["Path"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Action",
                                        values=["addPrefix"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Value",
                                        values=["/www"],
                                    ),
                                ],
                            ),
                        )],
                    )],
                )],
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo rule_engine can be imported using the id#rule_id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/ruleEngine:RuleEngine rule_engine zone-297z8rf93cfw#rule-ajol584a
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] rule_name: The rule name (1 to 255 characters).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleEngineRuleArgs']]]] rules: Rule items list.
        :param pulumi.Input[str] status: Rule status. Values:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: rule tag list.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleEngineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a teo rule_engine

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        rule1 = tencentcloud.teo.RuleEngine("rule1",
            zone_id=tencentcloud_teo_zone["example"]["id"],
            rule_name="test-rule",
            status="disable",
            rules=[tencentcloud.teo.RuleEngineRuleArgs(
                actions=[tencentcloud.teo.RuleEngineRuleActionArgs(
                    normal_action=tencentcloud.teo.RuleEngineRuleActionNormalActionArgs(
                        action="UpstreamUrlRedirect",
                        parameters=[
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Type",
                                values=["Path"],
                            ),
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Action",
                                values=["addPrefix"],
                            ),
                            tencentcloud.teo.RuleEngineRuleActionNormalActionParameterArgs(
                                name="Value",
                                values=["/sss"],
                            ),
                        ],
                    ),
                )],
                ors=[
                    tencentcloud.teo.RuleEngineRuleOrArgs(
                        ands=[
                            tencentcloud.teo.RuleEngineRuleOrAndArgs(
                                operator="equal",
                                target="host",
                                ignore_case=False,
                                values=["a.tf-teo-t.xyz"],
                            ),
                            tencentcloud.teo.RuleEngineRuleOrAndArgs(
                                operator="equal",
                                target="extension",
                                ignore_case=False,
                                values=["jpg"],
                            ),
                        ],
                    ),
                    tencentcloud.teo.RuleEngineRuleOrArgs(
                        ands=[tencentcloud.teo.RuleEngineRuleOrAndArgs(
                            operator="equal",
                            target="filename",
                            ignore_case=False,
                            values=["test.txt"],
                        )],
                    ),
                ],
                sub_rules=[tencentcloud.teo.RuleEngineRuleSubRuleArgs(
                    tags=["png"],
                    rules=[tencentcloud.teo.RuleEngineRuleSubRuleRuleArgs(
                        ors=[
                            tencentcloud.teo.RuleEngineRuleSubRuleRuleOrArgs(
                                ands=[
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                        operator="notequal",
                                        target="host",
                                        ignore_case=False,
                                        values=["a.tf-teo-t.xyz"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                        operator="equal",
                                        target="extension",
                                        ignore_case=False,
                                        values=["png"],
                                    ),
                                ],
                            ),
                            tencentcloud.teo.RuleEngineRuleSubRuleRuleOrArgs(
                                ands=[tencentcloud.teo.RuleEngineRuleSubRuleRuleOrAndArgs(
                                    operator="notequal",
                                    target="filename",
                                    ignore_case=False,
                                    values=["test.txt"],
                                )],
                            ),
                        ],
                        actions=[tencentcloud.teo.RuleEngineRuleSubRuleRuleActionArgs(
                            normal_action=tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionArgs(
                                action="UpstreamUrlRedirect",
                                parameters=[
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Type",
                                        values=["Path"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Action",
                                        values=["addPrefix"],
                                    ),
                                    tencentcloud.teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs(
                                        name="Value",
                                        values=["/www"],
                                    ),
                                ],
                            ),
                        )],
                    )],
                )],
            )])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo rule_engine can be imported using the id#rule_id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/ruleEngine:RuleEngine rule_engine zone-297z8rf93cfw#rule-ajol584a
        ```

        :param str resource_name: The name of the resource.
        :param RuleEngineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleEngineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleEngineRuleArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RuleEngineArgs.__new__(RuleEngineArgs)

            if rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'rule_name'")
            __props__.__dict__["rule_name"] = rule_name
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["rule_id"] = None
        super(RuleEngine, __self__).__init__(
            'tencentcloud:Teo/ruleEngine:RuleEngine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rule_id: Optional[pulumi.Input[str]] = None,
            rule_name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleEngineRuleArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'RuleEngine':
        """
        Get an existing RuleEngine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] rule_id: Rule ID.
        :param pulumi.Input[str] rule_name: The rule name (1 to 255 characters).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleEngineRuleArgs']]]] rules: Rule items list.
        :param pulumi.Input[str] status: Rule status. Values:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: rule tag list.
        :param pulumi.Input[str] zone_id: ID of the site.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RuleEngineState.__new__(_RuleEngineState)

        __props__.__dict__["rule_id"] = rule_id
        __props__.__dict__["rule_name"] = rule_name
        __props__.__dict__["rules"] = rules
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone_id"] = zone_id
        return RuleEngine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        Rule ID.
        """
        return pulumi.get(self, "rule_id")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[str]:
        """
        The rule name (1 to 255 characters).
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.RuleEngineRule']]:
        """
        Rule items list.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Rule status. Values:
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        rule tag list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        ID of the site.
        """
        return pulumi.get(self, "zone_id")

