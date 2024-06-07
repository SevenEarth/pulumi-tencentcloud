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

__all__ = [
    'GetTablesResult',
    'AwaitableGetTablesResult',
    'get_tables',
    'get_tables_output',
]

@pulumi.output_type
class GetTablesResult:
    """
    A collection of values returned by getTables.
    """
    def __init__(__self__, cluster_id=None, id=None, lists=None, result_output_file=None, table_id=None, table_name=None, tablegroup_id=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if table_id and not isinstance(table_id, str):
            raise TypeError("Expected argument 'table_id' to be a str")
        pulumi.set(__self__, "table_id", table_id)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)
        if tablegroup_id and not isinstance(tablegroup_id, str):
            raise TypeError("Expected argument 'tablegroup_id' to be a str")
        pulumi.set(__self__, "tablegroup_id", tablegroup_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetTablesListResult']:
        """
        A list of TcaplusDB tables. Each element contains the following attributes.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="tableId")
    def table_id(self) -> Optional[str]:
        """
        ID of the TcaplusDB table.
        """
        return pulumi.get(self, "table_id")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[str]:
        """
        Name of the TcaplusDB table.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter(name="tablegroupId")
    def tablegroup_id(self) -> Optional[str]:
        """
        Table group id of the TcaplusDB table.
        """
        return pulumi.get(self, "tablegroup_id")


class AwaitableGetTablesResult(GetTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTablesResult(
            cluster_id=self.cluster_id,
            id=self.id,
            lists=self.lists,
            result_output_file=self.result_output_file,
            table_id=self.table_id,
            table_name=self.table_name,
            tablegroup_id=self.tablegroup_id)


def get_tables(cluster_id: Optional[str] = None,
               result_output_file: Optional[str] = None,
               table_id: Optional[str] = None,
               table_name: Optional[str] = None,
               tablegroup_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTablesResult:
    """
    Use this data source to query TcaplusDB tables.


    :param str cluster_id: ID of the TcaplusDB cluster to be query.
    :param str result_output_file: File for saving results.
    :param str table_id: Table ID to be query.
    :param str table_name: Table name to be query.
    :param str tablegroup_id: ID of the table group to be query.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['resultOutputFile'] = result_output_file
    __args__['tableId'] = table_id
    __args__['tableName'] = table_name
    __args__['tablegroupId'] = tablegroup_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tcaplus/getTables:getTables', __args__, opts=opts, typ=GetTablesResult).value

    return AwaitableGetTablesResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        lists=pulumi.get(__ret__, 'lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        table_id=pulumi.get(__ret__, 'table_id'),
        table_name=pulumi.get(__ret__, 'table_name'),
        tablegroup_id=pulumi.get(__ret__, 'tablegroup_id'))


@_utilities.lift_output_func(get_tables)
def get_tables_output(cluster_id: Optional[pulumi.Input[str]] = None,
                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      table_id: Optional[pulumi.Input[Optional[str]]] = None,
                      table_name: Optional[pulumi.Input[Optional[str]]] = None,
                      tablegroup_id: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTablesResult]:
    """
    Use this data source to query TcaplusDB tables.


    :param str cluster_id: ID of the TcaplusDB cluster to be query.
    :param str result_output_file: File for saving results.
    :param str table_id: Table ID to be query.
    :param str table_name: Table name to be query.
    :param str tablegroup_id: ID of the table group to be query.
    """
    ...
