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

__all__ = ['CngwCanaryRuleArgs', 'CngwCanaryRule']

@pulumi.input_type
class CngwCanaryRuleArgs:
    def __init__(__self__, *,
                 canary_rule: pulumi.Input['CngwCanaryRuleCanaryRuleArgs'],
                 gateway_id: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a CngwCanaryRule resource.
        :param pulumi.Input['CngwCanaryRuleCanaryRuleArgs'] canary_rule: canary rule configuration.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] service_id: service ID.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "canary_rule", canary_rule)
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "service_id", service_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="canaryRule")
    def canary_rule(self) -> pulumi.Input['CngwCanaryRuleCanaryRuleArgs']:
        """
        canary rule configuration.
        """
        return pulumi.get(self, "canary_rule")

    @canary_rule.setter
    def canary_rule(self, value: pulumi.Input['CngwCanaryRuleCanaryRuleArgs']):
        pulumi.set(self, "canary_rule", value)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

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
class _CngwCanaryRuleState:
    def __init__(__self__, *,
                 canary_rule: Optional[pulumi.Input['CngwCanaryRuleCanaryRuleArgs']] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering CngwCanaryRule resources.
        :param pulumi.Input['CngwCanaryRuleCanaryRuleArgs'] canary_rule: canary rule configuration.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] service_id: service ID.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if canary_rule is not None:
            pulumi.set(__self__, "canary_rule", canary_rule)
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="canaryRule")
    def canary_rule(self) -> Optional[pulumi.Input['CngwCanaryRuleCanaryRuleArgs']]:
        """
        canary rule configuration.
        """
        return pulumi.get(self, "canary_rule")

    @canary_rule.setter
    def canary_rule(self, value: Optional[pulumi.Input['CngwCanaryRuleCanaryRuleArgs']]):
        pulumi.set(self, "canary_rule", value)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

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


class CngwCanaryRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 canary_rule: Optional[pulumi.Input[pulumi.InputType['CngwCanaryRuleCanaryRuleArgs']]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a tse cngw_canary_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cngw_service = tencentcloud.tse.CngwService("cngwService",
            gateway_id="gateway-ddbb709b",
            path="/test",
            protocol="http",
            retries=5,
            tags={
                "created": "terraform",
            },
            timeout=6000,
            upstream_type="IPList",
            upstream_info=tencentcloud.tse.CngwServiceUpstreamInfoArgs(
                algorithm="round-robin",
                auto_scaling_cvm_port=80,
                auto_scaling_group_id="asg-519acdug",
                auto_scaling_hook_status="Normal",
                auto_scaling_tat_cmd_status="Normal",
                port=0,
                slow_start=20,
                targets=[tencentcloud.tse.CngwServiceUpstreamInfoTargetArgs(
                    health="HEALTHCHECKS_OFF",
                    host="192.168.0.1",
                    port=80,
                    weight=100,
                )],
            ))
        cngw_canary_rule = tencentcloud.tse.CngwCanaryRule("cngwCanaryRule",
            gateway_id=cngw_service.gateway_id,
            service_id=cngw_service.service_id,
            tags={
                "created": "terraform",
            },
            canary_rule=tencentcloud.tse.CngwCanaryRuleCanaryRuleArgs(
                enabled=True,
                priority=100,
                balanced_service_lists=[tencentcloud.tse.CngwCanaryRuleCanaryRuleBalancedServiceListArgs(
                    percent=100,
                    service_id=cngw_service.service_id,
                    service_name=cngw_service.name,
                )],
                condition_lists=[tencentcloud.tse.CngwCanaryRuleCanaryRuleConditionListArgs(
                    key="test",
                    operator="eq",
                    type="query",
                    value="1",
                )],
            ))
        ```

        ## Import

        tse cngw_canary_rule can be imported using the gatewayId#serviceId#priority, e.g.

        ```sh
         $ pulumi import tencentcloud:Tse/cngwCanaryRule:CngwCanaryRule cngw_canary_rule gateway-ddbb709b#b6017eaf-2363-481e-9e93-8d65aaf498cd#100
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CngwCanaryRuleCanaryRuleArgs']] canary_rule: canary rule configuration.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] service_id: service ID.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CngwCanaryRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tse cngw_canary_rule

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cngw_service = tencentcloud.tse.CngwService("cngwService",
            gateway_id="gateway-ddbb709b",
            path="/test",
            protocol="http",
            retries=5,
            tags={
                "created": "terraform",
            },
            timeout=6000,
            upstream_type="IPList",
            upstream_info=tencentcloud.tse.CngwServiceUpstreamInfoArgs(
                algorithm="round-robin",
                auto_scaling_cvm_port=80,
                auto_scaling_group_id="asg-519acdug",
                auto_scaling_hook_status="Normal",
                auto_scaling_tat_cmd_status="Normal",
                port=0,
                slow_start=20,
                targets=[tencentcloud.tse.CngwServiceUpstreamInfoTargetArgs(
                    health="HEALTHCHECKS_OFF",
                    host="192.168.0.1",
                    port=80,
                    weight=100,
                )],
            ))
        cngw_canary_rule = tencentcloud.tse.CngwCanaryRule("cngwCanaryRule",
            gateway_id=cngw_service.gateway_id,
            service_id=cngw_service.service_id,
            tags={
                "created": "terraform",
            },
            canary_rule=tencentcloud.tse.CngwCanaryRuleCanaryRuleArgs(
                enabled=True,
                priority=100,
                balanced_service_lists=[tencentcloud.tse.CngwCanaryRuleCanaryRuleBalancedServiceListArgs(
                    percent=100,
                    service_id=cngw_service.service_id,
                    service_name=cngw_service.name,
                )],
                condition_lists=[tencentcloud.tse.CngwCanaryRuleCanaryRuleConditionListArgs(
                    key="test",
                    operator="eq",
                    type="query",
                    value="1",
                )],
            ))
        ```

        ## Import

        tse cngw_canary_rule can be imported using the gatewayId#serviceId#priority, e.g.

        ```sh
         $ pulumi import tencentcloud:Tse/cngwCanaryRule:CngwCanaryRule cngw_canary_rule gateway-ddbb709b#b6017eaf-2363-481e-9e93-8d65aaf498cd#100
        ```

        :param str resource_name: The name of the resource.
        :param CngwCanaryRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CngwCanaryRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 canary_rule: Optional[pulumi.Input[pulumi.InputType['CngwCanaryRuleCanaryRuleArgs']]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CngwCanaryRuleArgs.__new__(CngwCanaryRuleArgs)

            if canary_rule is None and not opts.urn:
                raise TypeError("Missing required property 'canary_rule'")
            __props__.__dict__["canary_rule"] = canary_rule
            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["tags"] = tags
        super(CngwCanaryRule, __self__).__init__(
            'tencentcloud:Tse/cngwCanaryRule:CngwCanaryRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            canary_rule: Optional[pulumi.Input[pulumi.InputType['CngwCanaryRuleCanaryRuleArgs']]] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'CngwCanaryRule':
        """
        Get an existing CngwCanaryRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CngwCanaryRuleCanaryRuleArgs']] canary_rule: canary rule configuration.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[str] service_id: service ID.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CngwCanaryRuleState.__new__(_CngwCanaryRuleState)

        __props__.__dict__["canary_rule"] = canary_rule
        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["tags"] = tags
        return CngwCanaryRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canaryRule")
    def canary_rule(self) -> pulumi.Output['outputs.CngwCanaryRuleCanaryRule']:
        """
        canary rule configuration.
        """
        return pulumi.get(self, "canary_rule")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        service ID.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

