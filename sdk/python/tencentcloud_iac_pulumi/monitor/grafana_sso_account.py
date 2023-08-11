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

__all__ = ['GrafanaSsoAccountArgs', 'GrafanaSsoAccount']

@pulumi.input_type
class GrafanaSsoAccountArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 notes: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]] = None):
        """
        The set of arguments for constructing a GrafanaSsoAccount resource.
        :param pulumi.Input[str] instance_id: grafana instance id.
        :param pulumi.Input[str] user_id: sub account uin of specific user.
        :param pulumi.Input[str] notes: account related description.
        :param pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]] roles: grafana role.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "user_id", user_id)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        grafana instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        sub account uin of specific user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        account related description.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]]:
        """
        grafana role.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]]):
        pulumi.set(self, "roles", value)


@pulumi.input_type
class _GrafanaSsoAccountState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GrafanaSsoAccount resources.
        :param pulumi.Input[str] instance_id: grafana instance id.
        :param pulumi.Input[str] notes: account related description.
        :param pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]] roles: grafana role.
        :param pulumi.Input[str] user_id: sub account uin of specific user.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        grafana instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        """
        account related description.
        """
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]]:
        """
        grafana role.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GrafanaSsoAccountRoleArgs']]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        sub account uin of specific user.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class GrafanaSsoAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrafanaSsoAccountRoleArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor grafana ssoAccount

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-6"
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=availability_zone,
            cidr_block="10.0.1.0/24")
        foo = tencentcloud.monitor.GrafanaInstance("foo",
            instance_name="test-grafana",
            vpc_id=vpc.id,
            subnet_ids=[subnet.id],
            grafana_init_password="1234567890",
            enable_internet=False,
            tags={
                "createdBy": "test",
            })
        sso_account = tencentcloud.monitor.GrafanaSsoAccount("ssoAccount",
            instance_id=foo.id,
            user_id="111",
            notes="desc12222",
            roles=[tencentcloud.monitor.GrafanaSsoAccountRoleArgs(
                organization="Main Org.",
                role="Admin",
            )])
        ```

        ## Import

        monitor grafana ssoAccount can be imported using the instance_id#user_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount ssoAccount grafana-50nj6v00#111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: grafana instance id.
        :param pulumi.Input[str] notes: account related description.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrafanaSsoAccountRoleArgs']]]] roles: grafana role.
        :param pulumi.Input[str] user_id: sub account uin of specific user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrafanaSsoAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor grafana ssoAccount

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-6"
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=availability_zone,
            cidr_block="10.0.1.0/24")
        foo = tencentcloud.monitor.GrafanaInstance("foo",
            instance_name="test-grafana",
            vpc_id=vpc.id,
            subnet_ids=[subnet.id],
            grafana_init_password="1234567890",
            enable_internet=False,
            tags={
                "createdBy": "test",
            })
        sso_account = tencentcloud.monitor.GrafanaSsoAccount("ssoAccount",
            instance_id=foo.id,
            user_id="111",
            notes="desc12222",
            roles=[tencentcloud.monitor.GrafanaSsoAccountRoleArgs(
                organization="Main Org.",
                role="Admin",
            )])
        ```

        ## Import

        monitor grafana ssoAccount can be imported using the instance_id#user_id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount ssoAccount grafana-50nj6v00#111
        ```

        :param str resource_name: The name of the resource.
        :param GrafanaSsoAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrafanaSsoAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrafanaSsoAccountRoleArgs']]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GrafanaSsoAccountArgs.__new__(GrafanaSsoAccountArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["notes"] = notes
            __props__.__dict__["roles"] = roles
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(GrafanaSsoAccount, __self__).__init__(
            'tencentcloud:Monitor/grafanaSsoAccount:GrafanaSsoAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrafanaSsoAccountRoleArgs']]]]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'GrafanaSsoAccount':
        """
        Get an existing GrafanaSsoAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: grafana instance id.
        :param pulumi.Input[str] notes: account related description.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GrafanaSsoAccountRoleArgs']]]] roles: grafana role.
        :param pulumi.Input[str] user_id: sub account uin of specific user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GrafanaSsoAccountState.__new__(_GrafanaSsoAccountState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["notes"] = notes
        __props__.__dict__["roles"] = roles
        __props__.__dict__["user_id"] = user_id
        return GrafanaSsoAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        grafana instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[str]:
        """
        account related description.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence['outputs.GrafanaSsoAccountRole']]:
        """
        grafana role.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        sub account uin of specific user.
        """
        return pulumi.get(self, "user_id")

