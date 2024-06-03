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
    'GetSqlFiltersResult',
    'AwaitableGetSqlFiltersResult',
    'get_sql_filters',
    'get_sql_filters_output',
]

@pulumi.output_type
class GetSqlFiltersResult:
    """
    A collection of values returned by getSqlFilters.
    """
    def __init__(__self__, filter_ids=None, id=None, instance_id=None, lists=None, result_output_file=None, statuses=None):
        if filter_ids and not isinstance(filter_ids, list):
            raise TypeError("Expected argument 'filter_ids' to be a list")
        pulumi.set(__self__, "filter_ids", filter_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)

    @property
    @pulumi.getter(name="filterIds")
    def filter_ids(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "filter_ids")

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
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetSqlFiltersListResult']:
        """
        sql filter list.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "statuses")


class AwaitableGetSqlFiltersResult(GetSqlFiltersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSqlFiltersResult(
            filter_ids=self.filter_ids,
            id=self.id,
            instance_id=self.instance_id,
            lists=self.lists,
            result_output_file=self.result_output_file,
            statuses=self.statuses)


def get_sql_filters(filter_ids: Optional[Sequence[int]] = None,
                    instance_id: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    statuses: Optional[Sequence[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSqlFiltersResult:
    """
    Use this data source to query detailed information of dbbrain sqlFilters

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    sql_filter = tencentcloud.dbbrain.SqlFilter("sqlFilter",
        instance_id="mysql_ins_id",
        session_token=tencentcloud.dbbrain.SqlFilterSessionTokenArgs(
            user="user",
            password="password",
        ),
        sql_type="SELECT",
        filter_key="test",
        max_concurrency=10,
        duration=3600)
    sql_filters = tencentcloud.Dbbrain.get_sql_filters_output(instance_id="mysql_ins_id",
        filter_ids=[sql_filter.filter_id])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[int] filter_ids: filter id list.
    :param str instance_id: instance id.
    :param str result_output_file: Used to save results.
    :param Sequence[str] statuses: status list.
    """
    __args__ = dict()
    __args__['filterIds'] = filter_ids
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['statuses'] = statuses
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dbbrain/getSqlFilters:getSqlFilters', __args__, opts=opts, typ=GetSqlFiltersResult).value

    return AwaitableGetSqlFiltersResult(
        filter_ids=pulumi.get(__ret__, 'filter_ids'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        lists=pulumi.get(__ret__, 'lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        statuses=pulumi.get(__ret__, 'statuses'))


@_utilities.lift_output_func(get_sql_filters)
def get_sql_filters_output(filter_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                           instance_id: Optional[pulumi.Input[str]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSqlFiltersResult]:
    """
    Use this data source to query detailed information of dbbrain sqlFilters

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    sql_filter = tencentcloud.dbbrain.SqlFilter("sqlFilter",
        instance_id="mysql_ins_id",
        session_token=tencentcloud.dbbrain.SqlFilterSessionTokenArgs(
            user="user",
            password="password",
        ),
        sql_type="SELECT",
        filter_key="test",
        max_concurrency=10,
        duration=3600)
    sql_filters = tencentcloud.Dbbrain.get_sql_filters_output(instance_id="mysql_ins_id",
        filter_ids=[sql_filter.filter_id])
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[int] filter_ids: filter id list.
    :param str instance_id: instance id.
    :param str result_output_file: Used to save results.
    :param Sequence[str] statuses: status list.
    """
    ...
