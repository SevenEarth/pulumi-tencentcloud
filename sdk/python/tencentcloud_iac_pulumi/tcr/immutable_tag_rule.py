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

__all__ = ['ImmutableTagRuleArgs', 'ImmutableTagRule']

@pulumi.input_type
class ImmutableTagRuleArgs:
    def __init__(__self__, *,
                 namespace_name: pulumi.Input[str],
                 registry_id: pulumi.Input[str],
                 rule: pulumi.Input['ImmutableTagRuleRuleArgs'],
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a ImmutableTagRule resource.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input['ImmutableTagRuleRuleArgs'] rule: rule.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "namespace_name", namespace_name)
        pulumi.set(__self__, "registry_id", registry_id)
        pulumi.set(__self__, "rule", rule)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        namespace name.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[str]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Input['ImmutableTagRuleRuleArgs']:
        """
        rule.
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: pulumi.Input['ImmutableTagRuleRuleArgs']):
        pulumi.set(self, "rule", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ImmutableTagRuleState:
    def __init__(__self__, *,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input['ImmutableTagRuleRuleArgs']] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering ImmutableTagRule resources.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input['ImmutableTagRuleRuleArgs'] rule: rule.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if rule is not None:
            pulumi.set(__self__, "rule", rule)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        namespace name.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[str]]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter
    def rule(self) -> Optional[pulumi.Input['ImmutableTagRuleRuleArgs']]:
        """
        rule.
        """
        return pulumi.get(self, "rule")

    @rule.setter
    def rule(self, value: Optional[pulumi.Input['ImmutableTagRuleRuleArgs']]):
        pulumi.set(self, "rule", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class ImmutableTagRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['ImmutableTagRuleRuleArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a tcr immutable_tag_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_rule = tencentcloud.tcr.ImmutableTagRule("myRule",
            namespace_name="%s",
            registry_id="%s",
            rule=tencentcloud.tcr.ImmutableTagRuleRuleArgs(
                disabled=False,
                repository_decoration="repoMatches",
                repository_pattern="**",
                tag_decoration="matches",
                tag_pattern="**",
            ),
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        tcr immutable_tag_rule can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/immutableTagRule:ImmutableTagRule immutable_tag_rule immutable_tag_rule_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[pulumi.InputType['ImmutableTagRuleRuleArgs']] rule: rule.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImmutableTagRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tcr immutable_tag_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_rule = tencentcloud.tcr.ImmutableTagRule("myRule",
            namespace_name="%s",
            registry_id="%s",
            rule=tencentcloud.tcr.ImmutableTagRuleRuleArgs(
                disabled=False,
                repository_decoration="repoMatches",
                repository_pattern="**",
                tag_decoration="matches",
                tag_pattern="**",
            ),
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        tcr immutable_tag_rule can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/immutableTagRule:ImmutableTagRule immutable_tag_rule immutable_tag_rule_id
        ```

        :param str resource_name: The name of the resource.
        :param ImmutableTagRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImmutableTagRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[pulumi.InputType['ImmutableTagRuleRuleArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = ImmutableTagRuleArgs.__new__(ImmutableTagRuleArgs)

            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            if rule is None and not opts.urn:
                raise TypeError("Missing required property 'rule'")
            __props__.__dict__["rule"] = rule
            __props__.__dict__["tags"] = tags
        super(ImmutableTagRule, __self__).__init__(
            'tencentcloud:Tcr/immutableTagRule:ImmutableTagRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            registry_id: Optional[pulumi.Input[str]] = None,
            rule: Optional[pulumi.Input[pulumi.InputType['ImmutableTagRuleRuleArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'ImmutableTagRule':
        """
        Get an existing ImmutableTagRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[pulumi.InputType['ImmutableTagRuleRuleArgs']] rule: rule.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImmutableTagRuleState.__new__(_ImmutableTagRuleState)

        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["rule"] = rule
        __props__.__dict__["tags"] = tags
        return ImmutableTagRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        namespace name.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter
    def rule(self) -> pulumi.Output['outputs.ImmutableTagRuleRule']:
        """
        rule.
        """
        return pulumi.get(self, "rule")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

