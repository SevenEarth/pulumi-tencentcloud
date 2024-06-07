# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AuditLogFileFilterArgs',
    'BackupDownloadRestrictionLimitVpcArgs',
    'ParamTemplateParamListArgs',
    'PasswordComplexityParamListArgs',
    'PrivilegeColumnArgs',
    'PrivilegeDatabaseArgs',
    'PrivilegeTableArgs',
    'ProxyProxyNodeCustomArgs',
    'RoGroupRoGroupInfoArgs',
    'RoGroupRoWeightValueArgs',
    'RollbackDatabaseArgs',
    'RollbackTableArgs',
    'RollbackTableTableArgs',
]

@pulumi.input_type
class AuditLogFileFilterArgs:
    def __init__(__self__, *,
                 affect_rows: Optional[pulumi.Input[int]] = None,
                 db_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exec_time: Optional[pulumi.Input[int]] = None,
                 hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sql: Optional[pulumi.Input[str]] = None,
                 sql_type: Optional[pulumi.Input[str]] = None,
                 sql_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sqls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 table_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[int] affect_rows: Affects the number of rows. Indicates to filter audit logs whose number of affected rows is greater than this value.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] db_names: Database name.
        :param pulumi.Input[int] exec_time: Execution time. The unit is: ms. Indicates to filter audit logs whose execution time is greater than this value.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] hosts: Client address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_names: The name of policy.
        :param pulumi.Input[str] sql: SQL statement. support fuzzy matching.
        :param pulumi.Input[str] sql_type: SQL type. Currently supported: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, ALTER, SET, REPLACE, EXECUTE.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sql_types: SQL type. Supports simultaneous query of multiple types. Currently supported: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, ALTER, SET, REPLACE, EXECUTE.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sqls: SQL statement. Support passing multiple sql statements.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] table_names: Table name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: User name.
        """
        if affect_rows is not None:
            pulumi.set(__self__, "affect_rows", affect_rows)
        if db_names is not None:
            pulumi.set(__self__, "db_names", db_names)
        if exec_time is not None:
            pulumi.set(__self__, "exec_time", exec_time)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if policy_names is not None:
            pulumi.set(__self__, "policy_names", policy_names)
        if sql is not None:
            pulumi.set(__self__, "sql", sql)
        if sql_type is not None:
            pulumi.set(__self__, "sql_type", sql_type)
        if sql_types is not None:
            pulumi.set(__self__, "sql_types", sql_types)
        if sqls is not None:
            pulumi.set(__self__, "sqls", sqls)
        if table_names is not None:
            pulumi.set(__self__, "table_names", table_names)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="affectRows")
    def affect_rows(self) -> Optional[pulumi.Input[int]]:
        """
        Affects the number of rows. Indicates to filter audit logs whose number of affected rows is greater than this value.
        """
        return pulumi.get(self, "affect_rows")

    @affect_rows.setter
    def affect_rows(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "affect_rows", value)

    @property
    @pulumi.getter(name="dbNames")
    def db_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Database name.
        """
        return pulumi.get(self, "db_names")

    @db_names.setter
    def db_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "db_names", value)

    @property
    @pulumi.getter(name="execTime")
    def exec_time(self) -> Optional[pulumi.Input[int]]:
        """
        Execution time. The unit is: ms. Indicates to filter audit logs whose execution time is greater than this value.
        """
        return pulumi.get(self, "exec_time")

    @exec_time.setter
    def exec_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "exec_time", value)

    @property
    @pulumi.getter
    def hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Client address.
        """
        return pulumi.get(self, "hosts")

    @hosts.setter
    def hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hosts", value)

    @property
    @pulumi.getter(name="policyNames")
    def policy_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The name of policy.
        """
        return pulumi.get(self, "policy_names")

    @policy_names.setter
    def policy_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policy_names", value)

    @property
    @pulumi.getter
    def sql(self) -> Optional[pulumi.Input[str]]:
        """
        SQL statement. support fuzzy matching.
        """
        return pulumi.get(self, "sql")

    @sql.setter
    def sql(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql", value)

    @property
    @pulumi.getter(name="sqlType")
    def sql_type(self) -> Optional[pulumi.Input[str]]:
        """
        SQL type. Currently supported: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, ALTER, SET, REPLACE, EXECUTE.
        """
        return pulumi.get(self, "sql_type")

    @sql_type.setter
    def sql_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sql_type", value)

    @property
    @pulumi.getter(name="sqlTypes")
    def sql_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        SQL type. Supports simultaneous query of multiple types. Currently supported: SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, ALTER, SET, REPLACE, EXECUTE.
        """
        return pulumi.get(self, "sql_types")

    @sql_types.setter
    def sql_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sql_types", value)

    @property
    @pulumi.getter
    def sqls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        SQL statement. Support passing multiple sql statements.
        """
        return pulumi.get(self, "sqls")

    @sqls.setter
    def sqls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sqls", value)

    @property
    @pulumi.getter(name="tableNames")
    def table_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Table name.
        """
        return pulumi.get(self, "table_names")

    @table_names.setter
    def table_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "table_names", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User name.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users", value)


@pulumi.input_type
class BackupDownloadRestrictionLimitVpcArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 vpc_lists: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] region: Restrict downloads from regions. Currently only the current region is supported.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vpc_lists: List of vpcs to limit downloads.
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "vpc_lists", vpc_lists)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Restrict downloads from regions. Currently only the current region is supported.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="vpcLists")
    def vpc_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of vpcs to limit downloads.
        """
        return pulumi.get(self, "vpc_lists")

    @vpc_lists.setter
    def vpc_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vpc_lists", value)


@pulumi.input_type
class ParamTemplateParamListArgs:
    def __init__(__self__, *,
                 current_value: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] current_value: The value of parameter.
        :param pulumi.Input[str] name: The name of parameter.
        """
        if current_value is not None:
            pulumi.set(__self__, "current_value", current_value)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="currentValue")
    def current_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of parameter.
        """
        return pulumi.get(self, "current_value")

    @current_value.setter
    def current_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "current_value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PasswordComplexityParamListArgs:
    def __init__(__self__, *,
                 current_value: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] current_value: Parameter value.
        :param pulumi.Input[str] name: Parameter name.
        """
        if current_value is not None:
            pulumi.set(__self__, "current_value", current_value)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="currentValue")
    def current_value(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter value.
        """
        return pulumi.get(self, "current_value")

    @current_value.setter
    def current_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "current_value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class PrivilegeColumnArgs:
    def __init__(__self__, *,
                 column_name: pulumi.Input[str],
                 database_name: pulumi.Input[str],
                 privileges: pulumi.Input[Sequence[pulumi.Input[str]]],
                 table_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] column_name: Column name.
        :param pulumi.Input[str] database_name: Database name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.
        :param pulumi.Input[str] table_name: Table name.
        """
        pulumi.set(__self__, "column_name", column_name)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "privileges", privileges)
        pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> pulumi.Input[str]:
        """
        Column name.
        """
        return pulumi.get(self, "column_name")

    @column_name.setter
    def column_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "column_name", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Database name.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        Table name.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)


@pulumi.input_type
class PrivilegeDatabaseArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 privileges: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] database_name: Database name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: Database privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "privileges", privileges)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Database name.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Database privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "privileges", value)


@pulumi.input_type
class PrivilegeTableArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 privileges: pulumi.Input[Sequence[pulumi.Input[str]]],
                 table_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] database_name: Database name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] privileges: Table privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW,TRIGGER.
        :param pulumi.Input[str] table_name: Table name.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "privileges", privileges)
        pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Database name.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def privileges(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Table privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW,TRIGGER.
        """
        return pulumi.get(self, "privileges")

    @privileges.setter
    def privileges(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "privileges", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        Table name.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)


@pulumi.input_type
class ProxyProxyNodeCustomArgs:
    def __init__(__self__, *,
                 cpu: pulumi.Input[int],
                 mem: pulumi.Input[int],
                 node_count: pulumi.Input[int],
                 region: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        """
        :param pulumi.Input[int] cpu: Number of CPU cores.
        :param pulumi.Input[int] mem: Memory size.
        :param pulumi.Input[int] node_count: Number of nodes.
        :param pulumi.Input[str] region: Region.
        :param pulumi.Input[str] zone: Zone.
        """
        pulumi.set(__self__, "cpu", cpu)
        pulumi.set(__self__, "mem", mem)
        pulumi.set(__self__, "node_count", node_count)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Input[int]:
        """
        Number of CPU cores.
        """
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: pulumi.Input[int]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter
    def mem(self) -> pulumi.Input[int]:
        """
        Memory size.
        """
        return pulumi.get(self, "mem")

    @mem.setter
    def mem(self, value: pulumi.Input[int]):
        pulumi.set(self, "mem", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Input[int]:
        """
        Number of nodes.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        Zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class RoGroupRoGroupInfoArgs:
    def __init__(__self__, *,
                 min_ro_in_group: Optional[pulumi.Input[int]] = None,
                 replication_delay_time: Optional[pulumi.Input[int]] = None,
                 ro_group_name: Optional[pulumi.Input[str]] = None,
                 ro_max_delay_time: Optional[pulumi.Input[int]] = None,
                 ro_offline_delay: Optional[pulumi.Input[int]] = None,
                 weight_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] min_ro_in_group: The minimum number of reserved instances. It can be set to any value less than or equal to the number of RO instances under this RO group. Note that if the setting value is greater than the number of RO instances, it will not be removed; if it is set to 0, all instances whose latency exceeds the limit will be removed.
        :param pulumi.Input[int] replication_delay_time: Delayed replication time.
        :param pulumi.Input[str] ro_group_name: RO group name.
        :param pulumi.Input[int] ro_max_delay_time: RO instance maximum latency threshold. The unit is seconds, the minimum value is 1. Note that the RO group must have enabled instance delay culling policy for this value to be valid.
        :param pulumi.Input[int] ro_offline_delay: Whether to enable delayed culling of instances. Supported values are: 1 - on; 0 - not on. Note that if you enable instance delay culling, you must set the delay threshold (RoMaxDelayTime) parameter.
        :param pulumi.Input[str] weight_mode: weight mode. Supported values include: `system` - automatically assigned by the system; `custom` - user-defined settings. Note that if the `custom` mode is set, the RO instance weight configuration (RoWeightValues) parameter must be set.
        """
        if min_ro_in_group is not None:
            pulumi.set(__self__, "min_ro_in_group", min_ro_in_group)
        if replication_delay_time is not None:
            pulumi.set(__self__, "replication_delay_time", replication_delay_time)
        if ro_group_name is not None:
            pulumi.set(__self__, "ro_group_name", ro_group_name)
        if ro_max_delay_time is not None:
            pulumi.set(__self__, "ro_max_delay_time", ro_max_delay_time)
        if ro_offline_delay is not None:
            pulumi.set(__self__, "ro_offline_delay", ro_offline_delay)
        if weight_mode is not None:
            pulumi.set(__self__, "weight_mode", weight_mode)

    @property
    @pulumi.getter(name="minRoInGroup")
    def min_ro_in_group(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of reserved instances. It can be set to any value less than or equal to the number of RO instances under this RO group. Note that if the setting value is greater than the number of RO instances, it will not be removed; if it is set to 0, all instances whose latency exceeds the limit will be removed.
        """
        return pulumi.get(self, "min_ro_in_group")

    @min_ro_in_group.setter
    def min_ro_in_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ro_in_group", value)

    @property
    @pulumi.getter(name="replicationDelayTime")
    def replication_delay_time(self) -> Optional[pulumi.Input[int]]:
        """
        Delayed replication time.
        """
        return pulumi.get(self, "replication_delay_time")

    @replication_delay_time.setter
    def replication_delay_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "replication_delay_time", value)

    @property
    @pulumi.getter(name="roGroupName")
    def ro_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        RO group name.
        """
        return pulumi.get(self, "ro_group_name")

    @ro_group_name.setter
    def ro_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ro_group_name", value)

    @property
    @pulumi.getter(name="roMaxDelayTime")
    def ro_max_delay_time(self) -> Optional[pulumi.Input[int]]:
        """
        RO instance maximum latency threshold. The unit is seconds, the minimum value is 1. Note that the RO group must have enabled instance delay culling policy for this value to be valid.
        """
        return pulumi.get(self, "ro_max_delay_time")

    @ro_max_delay_time.setter
    def ro_max_delay_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ro_max_delay_time", value)

    @property
    @pulumi.getter(name="roOfflineDelay")
    def ro_offline_delay(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to enable delayed culling of instances. Supported values are: 1 - on; 0 - not on. Note that if you enable instance delay culling, you must set the delay threshold (RoMaxDelayTime) parameter.
        """
        return pulumi.get(self, "ro_offline_delay")

    @ro_offline_delay.setter
    def ro_offline_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ro_offline_delay", value)

    @property
    @pulumi.getter(name="weightMode")
    def weight_mode(self) -> Optional[pulumi.Input[str]]:
        """
        weight mode. Supported values include: `system` - automatically assigned by the system; `custom` - user-defined settings. Note that if the `custom` mode is set, the RO instance weight configuration (RoWeightValues) parameter must be set.
        """
        return pulumi.get(self, "weight_mode")

    @weight_mode.setter
    def weight_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "weight_mode", value)


@pulumi.input_type
class RoGroupRoWeightValueArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 weight: pulumi.Input[int]):
        """
        :param pulumi.Input[str] instance_id: RO instance ID.
        :param pulumi.Input[int] weight: Weights. The value range is [0, 100].
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        RO instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[int]:
        """
        Weights. The value range is [0, 100].
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[int]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class RollbackDatabaseArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 new_database_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] database_name: The original database name before rollback.
        :param pulumi.Input[str] new_database_name: The new database name after rollback.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "new_database_name", new_database_name)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The original database name before rollback.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="newDatabaseName")
    def new_database_name(self) -> pulumi.Input[str]:
        """
        The new database name after rollback.
        """
        return pulumi.get(self, "new_database_name")

    @new_database_name.setter
    def new_database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "new_database_name", value)


@pulumi.input_type
class RollbackTableArgs:
    def __init__(__self__, *,
                 database: pulumi.Input[str],
                 tables: pulumi.Input[Sequence[pulumi.Input['RollbackTableTableArgs']]]):
        """
        :param pulumi.Input[str] database: Database name.
        :param pulumi.Input[Sequence[pulumi.Input['RollbackTableTableArgs']]] tables: Database table details.
        """
        pulumi.set(__self__, "database", database)
        pulumi.set(__self__, "tables", tables)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Input[str]:
        """
        Database name.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: pulumi.Input[str]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def tables(self) -> pulumi.Input[Sequence[pulumi.Input['RollbackTableTableArgs']]]:
        """
        Database table details.
        """
        return pulumi.get(self, "tables")

    @tables.setter
    def tables(self, value: pulumi.Input[Sequence[pulumi.Input['RollbackTableTableArgs']]]):
        pulumi.set(self, "tables", value)


@pulumi.input_type
class RollbackTableTableArgs:
    def __init__(__self__, *,
                 new_table_name: pulumi.Input[str],
                 table_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] new_table_name: New database table name after rollback.
        :param pulumi.Input[str] table_name: The original database table name before rollback.
        """
        pulumi.set(__self__, "new_table_name", new_table_name)
        pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="newTableName")
    def new_table_name(self) -> pulumi.Input[str]:
        """
        New database table name after rollback.
        """
        return pulumi.get(self, "new_table_name")

    @new_table_name.setter
    def new_table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "new_table_name", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        The original database table name before rollback.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)


