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

__all__ = ['PasswordComplexityArgs', 'PasswordComplexity']

@pulumi.input_type
class PasswordComplexityArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]] = None):
        """
        The set of arguments for constructing a PasswordComplexity resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]] param_lists: List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if param_lists is not None:
            pulumi.set(__self__, "param_lists", param_lists)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]]:
        """
        List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        return pulumi.get(self, "param_lists")

    @param_lists.setter
    def param_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]]):
        pulumi.set(self, "param_lists", value)


@pulumi.input_type
class _PasswordComplexityState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]] = None):
        """
        Input properties used for looking up and filtering PasswordComplexity resources.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]] param_lists: List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if param_lists is not None:
            pulumi.set(__self__, "param_lists", param_lists)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]]:
        """
        List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        return pulumi.get(self, "param_lists")

    @param_lists.setter
    def param_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PasswordComplexityParamListArgs']]]]):
        pulumi.set(self, "param_lists", value)


class PasswordComplexity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PasswordComplexityParamListArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql password_complexity

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="cdb")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[0].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="mysql test")
        example_instance = tencentcloud.mysql.Instance("exampleInstance",
            internet_service=1,
            engine_version="5.7",
            charge_type="POSTPAID",
            root_password="PassWord123",
            slave_deploy_mode=0,
            availability_zone=zones.zones[0].name,
            slave_sync_mode=1,
            instance_name="tf-example-mysql",
            mem_size=4000,
            volume_size=200,
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            intranet_port=3306,
            security_groups=[security_group.id],
            tags={
                "name": "test",
            },
            parameters={
                "character_set_server": "utf8",
                "max_connections": "1000",
            })
        example_password_complexity = tencentcloud.mysql.PasswordComplexity("examplePasswordComplexity",
            instance_id=example_instance.id,
            param_lists=[
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_length",
                    current_value="8",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_mixed_case_count",
                    current_value="2",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_number_count",
                    current_value="2",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_special_char_count",
                    current_value="2",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PasswordComplexityParamListArgs']]]] param_lists: List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PasswordComplexityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql password_complexity

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        zones = tencentcloud.Availability.get_zones_by_product(product="cdb")
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            availability_zone=zones.zones[0].name,
            vpc_id=vpc.id,
            cidr_block="10.0.0.0/16",
            is_multicast=False)
        security_group = tencentcloud.security.Group("securityGroup", description="mysql test")
        example_instance = tencentcloud.mysql.Instance("exampleInstance",
            internet_service=1,
            engine_version="5.7",
            charge_type="POSTPAID",
            root_password="PassWord123",
            slave_deploy_mode=0,
            availability_zone=zones.zones[0].name,
            slave_sync_mode=1,
            instance_name="tf-example-mysql",
            mem_size=4000,
            volume_size=200,
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            intranet_port=3306,
            security_groups=[security_group.id],
            tags={
                "name": "test",
            },
            parameters={
                "character_set_server": "utf8",
                "max_connections": "1000",
            })
        example_password_complexity = tencentcloud.mysql.PasswordComplexity("examplePasswordComplexity",
            instance_id=example_instance.id,
            param_lists=[
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_length",
                    current_value="8",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_mixed_case_count",
                    current_value="2",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_number_count",
                    current_value="2",
                ),
                tencentcloud.mysql.PasswordComplexityParamListArgs(
                    name="validate_password_special_char_count",
                    current_value="2",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param PasswordComplexityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PasswordComplexityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PasswordComplexityParamListArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PasswordComplexityArgs.__new__(PasswordComplexityArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["param_lists"] = param_lists
        super(PasswordComplexity, __self__).__init__(
            'tencentcloud:Mysql/passwordComplexity:PasswordComplexity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            param_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PasswordComplexityParamListArgs']]]]] = None) -> 'PasswordComplexity':
        """
        Get an existing PasswordComplexity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PasswordComplexityParamListArgs']]]] param_lists: List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PasswordComplexityState.__new__(_PasswordComplexityState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["param_lists"] = param_lists
        return PasswordComplexity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="paramLists")
    def param_lists(self) -> pulumi.Output[Optional[Sequence['outputs.PasswordComplexityParamList']]]:
        """
        List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validate_password_polic`, `validate_password_lengt` `validate_password_mixed_case_coun`, `validate_password_number_coun`, `validate_password_special_char_coun`.
        """
        return pulumi.get(self, "param_lists")

