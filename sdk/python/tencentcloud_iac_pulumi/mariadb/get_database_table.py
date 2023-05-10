# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDatabaseTableResult',
    'AwaitableGetDatabaseTableResult',
    'get_database_table',
    'get_database_table_output',
]

@pulumi.output_type
class GetDatabaseTableResult:
    """
    A collection of values returned by getDatabaseTable.
    """
    def __init__(__self__, cols=None, db_name=None, id=None, instance_id=None, result_output_file=None, table=None):
        if cols and not isinstance(cols, list):
            raise TypeError("Expected argument 'cols' to be a list")
        pulumi.set(__self__, "cols", cols)
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if table and not isinstance(table, str):
            raise TypeError("Expected argument 'table' to be a str")
        pulumi.set(__self__, "table", table)

    @property
    @pulumi.getter
    def cols(self) -> Sequence['outputs.GetDatabaseTableColResult']:
        """
        column list.
        """
        return pulumi.get(self, "cols")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def table(self) -> str:
        return pulumi.get(self, "table")


class AwaitableGetDatabaseTableResult(GetDatabaseTableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseTableResult(
            cols=self.cols,
            db_name=self.db_name,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            table=self.table)


def get_database_table(db_name: Optional[str] = None,
                       instance_id: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       table: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseTableResult:
    """
    Use this data source to query detailed information of mariadb database_table

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    database_table = tencentcloud.Mariadb.get_database_table(db_name="mysql",
        instance_id="tdsql-e9tklsgz",
        table="server_cost")
    ```


    :param str db_name: database name.
    :param str instance_id: instance id.
    :param str result_output_file: Used to save results.
    :param str table: table name.
    """
    __args__ = dict()
    __args__['dbName'] = db_name
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['table'] = table
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mariadb/getDatabaseTable:getDatabaseTable', __args__, opts=opts, typ=GetDatabaseTableResult).value

    return AwaitableGetDatabaseTableResult(
        cols=__ret__.cols,
        db_name=__ret__.db_name,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file,
        table=__ret__.table)


@_utilities.lift_output_func(get_database_table)
def get_database_table_output(db_name: Optional[pulumi.Input[str]] = None,
                              instance_id: Optional[pulumi.Input[str]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              table: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseTableResult]:
    """
    Use this data source to query detailed information of mariadb database_table

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    database_table = tencentcloud.Mariadb.get_database_table(db_name="mysql",
        instance_id="tdsql-e9tklsgz",
        table="server_cost")
    ```


    :param str db_name: database name.
    :param str instance_id: instance id.
    :param str result_output_file: Used to save results.
    :param str table: table name.
    """
    ...
