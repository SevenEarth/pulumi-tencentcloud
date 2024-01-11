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

__all__ = ['AccountPermissionArgs', 'AccountPermission']

@pulumi.input_type
class AccountPermissionArgs:
    def __init__(__self__, *,
                 all_database: pulumi.Input[bool],
                 cluster: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 database_privilege_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]] = None,
                 global_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AccountPermission resource.
        :param pulumi.Input[bool] all_database: Whether all database tables.
        :param pulumi.Input[str] cluster: Cluster name.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] user_name: User name.
        :param pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]] database_privilege_lists: Database privilege list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_privileges: Global privileges.
        """
        pulumi.set(__self__, "all_database", all_database)
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "user_name", user_name)
        if database_privilege_lists is not None:
            pulumi.set(__self__, "database_privilege_lists", database_privilege_lists)
        if global_privileges is not None:
            pulumi.set(__self__, "global_privileges", global_privileges)

    @property
    @pulumi.getter(name="allDatabase")
    def all_database(self) -> pulumi.Input[bool]:
        """
        Whether all database tables.
        """
        return pulumi.get(self, "all_database")

    @all_database.setter
    def all_database(self, value: pulumi.Input[bool]):
        pulumi.set(self, "all_database", value)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Input[str]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        User name.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="databasePrivilegeLists")
    def database_privilege_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]]:
        """
        Database privilege list.
        """
        return pulumi.get(self, "database_privilege_lists")

    @database_privilege_lists.setter
    def database_privilege_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]]):
        pulumi.set(self, "database_privilege_lists", value)

    @property
    @pulumi.getter(name="globalPrivileges")
    def global_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Global privileges.
        """
        return pulumi.get(self, "global_privileges")

    @global_privileges.setter
    def global_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_privileges", value)


@pulumi.input_type
class _AccountPermissionState:
    def __init__(__self__, *,
                 all_database: Optional[pulumi.Input[bool]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 database_privilege_lists: Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]] = None,
                 global_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccountPermission resources.
        :param pulumi.Input[bool] all_database: Whether all database tables.
        :param pulumi.Input[str] cluster: Cluster name.
        :param pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]] database_privilege_lists: Database privilege list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_privileges: Global privileges.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] user_name: User name.
        """
        if all_database is not None:
            pulumi.set(__self__, "all_database", all_database)
        if cluster is not None:
            pulumi.set(__self__, "cluster", cluster)
        if database_privilege_lists is not None:
            pulumi.set(__self__, "database_privilege_lists", database_privilege_lists)
        if global_privileges is not None:
            pulumi.set(__self__, "global_privileges", global_privileges)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="allDatabase")
    def all_database(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether all database tables.
        """
        return pulumi.get(self, "all_database")

    @all_database.setter
    def all_database(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "all_database", value)

    @property
    @pulumi.getter
    def cluster(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="databasePrivilegeLists")
    def database_privilege_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]]:
        """
        Database privilege list.
        """
        return pulumi.get(self, "database_privilege_lists")

    @database_privilege_lists.setter
    def database_privilege_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccountPermissionDatabasePrivilegeListArgs']]]]):
        pulumi.set(self, "database_privilege_lists", value)

    @property
    @pulumi.getter(name="globalPrivileges")
    def global_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Global privileges.
        """
        return pulumi.get(self, "global_privileges")

    @global_privileges.setter
    def global_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "global_privileges", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class AccountPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_database: Optional[pulumi.Input[bool]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 database_privilege_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountPermissionDatabasePrivilegeListArgs']]]]] = None,
                 global_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a clickhouse account_permission

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        account_permission_all_database = tencentcloud.clickhouse.AccountPermission("accountPermissionAllDatabase",
            all_database=True,
            cluster="default_cluster",
            global_privileges=[
                "SELECT",
                "ALTER",
            ],
            instance_id="cdwch-xxxxxx",
            user_name="user1")
        account_permission_not_all_database = tencentcloud.clickhouse.AccountPermission("accountPermissionNotAllDatabase",
            all_database=False,
            cluster="default_cluster",
            database_privilege_lists=[tencentcloud.clickhouse.AccountPermissionDatabasePrivilegeListArgs(
                database_name="xxxxxx",
                database_privileges=[
                    "SELECT",
                    "ALTER",
                ],
            )],
            instance_id="cdwch-xxxxxx",
            user_name="user2")
        ```

        ## Import

        clickhouse account_permission can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clickhouse/accountPermission:AccountPermission account_permission ${instanceId}#${cluster}#${userName}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_database: Whether all database tables.
        :param pulumi.Input[str] cluster: Cluster name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountPermissionDatabasePrivilegeListArgs']]]] database_privilege_lists: Database privilege list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_privileges: Global privileges.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] user_name: User name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a clickhouse account_permission

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        account_permission_all_database = tencentcloud.clickhouse.AccountPermission("accountPermissionAllDatabase",
            all_database=True,
            cluster="default_cluster",
            global_privileges=[
                "SELECT",
                "ALTER",
            ],
            instance_id="cdwch-xxxxxx",
            user_name="user1")
        account_permission_not_all_database = tencentcloud.clickhouse.AccountPermission("accountPermissionNotAllDatabase",
            all_database=False,
            cluster="default_cluster",
            database_privilege_lists=[tencentcloud.clickhouse.AccountPermissionDatabasePrivilegeListArgs(
                database_name="xxxxxx",
                database_privileges=[
                    "SELECT",
                    "ALTER",
                ],
            )],
            instance_id="cdwch-xxxxxx",
            user_name="user2")
        ```

        ## Import

        clickhouse account_permission can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Clickhouse/accountPermission:AccountPermission account_permission ${instanceId}#${cluster}#${userName}
        ```

        :param str resource_name: The name of the resource.
        :param AccountPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_database: Optional[pulumi.Input[bool]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 database_privilege_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountPermissionDatabasePrivilegeListArgs']]]]] = None,
                 global_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = AccountPermissionArgs.__new__(AccountPermissionArgs)

            if all_database is None and not opts.urn:
                raise TypeError("Missing required property 'all_database'")
            __props__.__dict__["all_database"] = all_database
            if cluster is None and not opts.urn:
                raise TypeError("Missing required property 'cluster'")
            __props__.__dict__["cluster"] = cluster
            __props__.__dict__["database_privilege_lists"] = database_privilege_lists
            __props__.__dict__["global_privileges"] = global_privileges
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(AccountPermission, __self__).__init__(
            'tencentcloud:Clickhouse/accountPermission:AccountPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_database: Optional[pulumi.Input[bool]] = None,
            cluster: Optional[pulumi.Input[str]] = None,
            database_privilege_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountPermissionDatabasePrivilegeListArgs']]]]] = None,
            global_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'AccountPermission':
        """
        Get an existing AccountPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] all_database: Whether all database tables.
        :param pulumi.Input[str] cluster: Cluster name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountPermissionDatabasePrivilegeListArgs']]]] database_privilege_lists: Database privilege list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] global_privileges: Global privileges.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] user_name: User name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountPermissionState.__new__(_AccountPermissionState)

        __props__.__dict__["all_database"] = all_database
        __props__.__dict__["cluster"] = cluster
        __props__.__dict__["database_privilege_lists"] = database_privilege_lists
        __props__.__dict__["global_privileges"] = global_privileges
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["user_name"] = user_name
        return AccountPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allDatabase")
    def all_database(self) -> pulumi.Output[bool]:
        """
        Whether all database tables.
        """
        return pulumi.get(self, "all_database")

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[str]:
        """
        Cluster name.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="databasePrivilegeLists")
    def database_privilege_lists(self) -> pulumi.Output[Optional[Sequence['outputs.AccountPermissionDatabasePrivilegeList']]]:
        """
        Database privilege list.
        """
        return pulumi.get(self, "database_privilege_lists")

    @property
    @pulumi.getter(name="globalPrivileges")
    def global_privileges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Global privileges.
        """
        return pulumi.get(self, "global_privileges")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        User name.
        """
        return pulumi.get(self, "user_name")
