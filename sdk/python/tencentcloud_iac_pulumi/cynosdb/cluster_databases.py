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

__all__ = ['ClusterDatabasesArgs', 'ClusterDatabases']

@pulumi.input_type
class ClusterDatabasesArgs:
    def __init__(__self__, *,
                 character_set: pulumi.Input[str],
                 cluster_id: pulumi.Input[str],
                 collate_rule: pulumi.Input[str],
                 db_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 user_host_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]] = None):
        """
        The set of arguments for constructing a ClusterDatabases resource.
        :param pulumi.Input[str] character_set: Character Set Type.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] collate_rule: Sort Rules.
        :param pulumi.Input[str] db_name: Database name.
        :param pulumi.Input[str] description: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]] user_host_privileges: Authorize user host permissions.
        """
        pulumi.set(__self__, "character_set", character_set)
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "collate_rule", collate_rule)
        pulumi.set(__self__, "db_name", db_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if user_host_privileges is not None:
            pulumi.set(__self__, "user_host_privileges", user_host_privileges)

    @property
    @pulumi.getter(name="characterSet")
    def character_set(self) -> pulumi.Input[str]:
        """
        Character Set Type.
        """
        return pulumi.get(self, "character_set")

    @character_set.setter
    def character_set(self, value: pulumi.Input[str]):
        pulumi.set(self, "character_set", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="collateRule")
    def collate_rule(self) -> pulumi.Input[str]:
        """
        Sort Rules.
        """
        return pulumi.get(self, "collate_rule")

    @collate_rule.setter
    def collate_rule(self, value: pulumi.Input[str]):
        pulumi.set(self, "collate_rule", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        Database name.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="userHostPrivileges")
    def user_host_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]]:
        """
        Authorize user host permissions.
        """
        return pulumi.get(self, "user_host_privileges")

    @user_host_privileges.setter
    def user_host_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]]):
        pulumi.set(self, "user_host_privileges", value)


@pulumi.input_type
class _ClusterDatabasesState:
    def __init__(__self__, *,
                 character_set: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 collate_rule: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 user_host_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]] = None):
        """
        Input properties used for looking up and filtering ClusterDatabases resources.
        :param pulumi.Input[str] character_set: Character Set Type.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] collate_rule: Sort Rules.
        :param pulumi.Input[str] db_name: Database name.
        :param pulumi.Input[str] description: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]] user_host_privileges: Authorize user host permissions.
        """
        if character_set is not None:
            pulumi.set(__self__, "character_set", character_set)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if collate_rule is not None:
            pulumi.set(__self__, "collate_rule", collate_rule)
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if user_host_privileges is not None:
            pulumi.set(__self__, "user_host_privileges", user_host_privileges)

    @property
    @pulumi.getter(name="characterSet")
    def character_set(self) -> Optional[pulumi.Input[str]]:
        """
        Character Set Type.
        """
        return pulumi.get(self, "character_set")

    @character_set.setter
    def character_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="collateRule")
    def collate_rule(self) -> Optional[pulumi.Input[str]]:
        """
        Sort Rules.
        """
        return pulumi.get(self, "collate_rule")

    @collate_rule.setter
    def collate_rule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collate_rule", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database name.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="userHostPrivileges")
    def user_host_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]]:
        """
        Authorize user host permissions.
        """
        return pulumi.get(self, "user_host_privileges")

    @user_host_privileges.setter
    def user_host_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterDatabasesUserHostPrivilegeArgs']]]]):
        pulumi.set(self, "user_host_privileges", value)


class ClusterDatabases(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 character_set: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 collate_rule: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 user_host_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDatabasesUserHostPrivilegeArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a cynosdb cluster_databases

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cluster_databases = tencentcloud.cynosdb.ClusterDatabases("clusterDatabases",
            character_set="utf8",
            cluster_id="cynosdbmysql-bws8h88b",
            collate_rule="utf8_general_ci",
            db_name="terraform-test",
            description="terraform test",
            user_host_privileges=[tencentcloud.cynosdb.ClusterDatabasesUserHostPrivilegeArgs(
                db_host="%",
                db_privilege="READ_ONLY",
                db_user_name="root",
            )])
        ```

        ## Import

        cynosdb cluster_databases can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases cluster_databases cluster_databases_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set: Character Set Type.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] collate_rule: Sort Rules.
        :param pulumi.Input[str] db_name: Database name.
        :param pulumi.Input[str] description: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDatabasesUserHostPrivilegeArgs']]]] user_host_privileges: Authorize user host permissions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterDatabasesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cynosdb cluster_databases

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cluster_databases = tencentcloud.cynosdb.ClusterDatabases("clusterDatabases",
            character_set="utf8",
            cluster_id="cynosdbmysql-bws8h88b",
            collate_rule="utf8_general_ci",
            db_name="terraform-test",
            description="terraform test",
            user_host_privileges=[tencentcloud.cynosdb.ClusterDatabasesUserHostPrivilegeArgs(
                db_host="%",
                db_privilege="READ_ONLY",
                db_user_name="root",
            )])
        ```

        ## Import

        cynosdb cluster_databases can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases cluster_databases cluster_databases_id
        ```

        :param str resource_name: The name of the resource.
        :param ClusterDatabasesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterDatabasesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 character_set: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 collate_rule: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 user_host_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDatabasesUserHostPrivilegeArgs']]]]] = None,
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
            __props__ = ClusterDatabasesArgs.__new__(ClusterDatabasesArgs)

            if character_set is None and not opts.urn:
                raise TypeError("Missing required property 'character_set'")
            __props__.__dict__["character_set"] = character_set
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if collate_rule is None and not opts.urn:
                raise TypeError("Missing required property 'collate_rule'")
            __props__.__dict__["collate_rule"] = collate_rule
            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            __props__.__dict__["description"] = description
            __props__.__dict__["user_host_privileges"] = user_host_privileges
        super(ClusterDatabases, __self__).__init__(
            'tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            character_set: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            collate_rule: Optional[pulumi.Input[str]] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            user_host_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDatabasesUserHostPrivilegeArgs']]]]] = None) -> 'ClusterDatabases':
        """
        Get an existing ClusterDatabases resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set: Character Set Type.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] collate_rule: Sort Rules.
        :param pulumi.Input[str] db_name: Database name.
        :param pulumi.Input[str] description: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterDatabasesUserHostPrivilegeArgs']]]] user_host_privileges: Authorize user host permissions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterDatabasesState.__new__(_ClusterDatabasesState)

        __props__.__dict__["character_set"] = character_set
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["collate_rule"] = collate_rule
        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["description"] = description
        __props__.__dict__["user_host_privileges"] = user_host_privileges
        return ClusterDatabases(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="characterSet")
    def character_set(self) -> pulumi.Output[str]:
        """
        Character Set Type.
        """
        return pulumi.get(self, "character_set")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="collateRule")
    def collate_rule(self) -> pulumi.Output[str]:
        """
        Sort Rules.
        """
        return pulumi.get(self, "collate_rule")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        Database name.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="userHostPrivileges")
    def user_host_privileges(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterDatabasesUserHostPrivilege']]]:
        """
        Authorize user host permissions.
        """
        return pulumi.get(self, "user_host_privileges")

