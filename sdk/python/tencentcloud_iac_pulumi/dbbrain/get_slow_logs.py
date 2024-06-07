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
    'GetSlowLogsResult',
    'AwaitableGetSlowLogsResult',
    'get_slow_logs',
    'get_slow_logs_output',
]

@pulumi.output_type
class GetSlowLogsResult:
    """
    A collection of values returned by getSlowLogs.
    """
    def __init__(__self__, dbs=None, end_time=None, id=None, instance_id=None, ips=None, keys=None, md5=None, product=None, result_output_file=None, rows=None, start_time=None, times=None, users=None):
        if dbs and not isinstance(dbs, list):
            raise TypeError("Expected argument 'dbs' to be a list")
        pulumi.set(__self__, "dbs", dbs)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if md5 and not isinstance(md5, str):
            raise TypeError("Expected argument 'md5' to be a str")
        pulumi.set(__self__, "md5", md5)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if rows and not isinstance(rows, list):
            raise TypeError("Expected argument 'rows' to be a list")
        pulumi.set(__self__, "rows", rows)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if times and not isinstance(times, list):
            raise TypeError("Expected argument 'times' to be a list")
        pulumi.set(__self__, "times", times)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def dbs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "dbs")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        return pulumi.get(self, "end_time")

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
    def ips(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def md5(self) -> str:
        return pulumi.get(self, "md5")

    @property
    @pulumi.getter
    def product(self) -> str:
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def rows(self) -> Sequence['outputs.GetSlowLogsRowResult']:
        """
        Slow log details.
        """
        return pulumi.get(self, "rows")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def times(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "times")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


class AwaitableGetSlowLogsResult(GetSlowLogsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSlowLogsResult(
            dbs=self.dbs,
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            ips=self.ips,
            keys=self.keys,
            md5=self.md5,
            product=self.product,
            result_output_file=self.result_output_file,
            rows=self.rows,
            start_time=self.start_time,
            times=self.times,
            users=self.users)


def get_slow_logs(dbs: Optional[Sequence[str]] = None,
                  end_time: Optional[str] = None,
                  instance_id: Optional[str] = None,
                  ips: Optional[Sequence[str]] = None,
                  keys: Optional[Sequence[str]] = None,
                  md5: Optional[str] = None,
                  product: Optional[str] = None,
                  result_output_file: Optional[str] = None,
                  start_time: Optional[str] = None,
                  times: Optional[Sequence[int]] = None,
                  users: Optional[Sequence[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSlowLogsResult:
    """
    Use this data source to query detailed information of dbbrain slow_logs

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    slow_logs = tencentcloud.Dbbrain.get_slow_logs(end_time="%s",
        instance_id="%s",
        md5="4961208426639258265",
        product="mysql",
        start_time="%s")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] dbs: database list.
    :param str end_time: The deadline, such as 2019-09-11 10:13:14, the interval between the deadline and the start time is less than 7 days.
    :param str instance_id: instance Id.
    :param Sequence[str] ips: ip.
    :param Sequence[str] keys: keywords.
    :param str md5: md5 value of sql template.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str start_time: Start time, such as 2019-09-10 12:13:14.
    :param Sequence[int] times: Time-consuming interval, the left and right boundaries of the time-consuming interval correspond to the 0th element and the first element of the array respectively.
    :param Sequence[str] users: user.
    """
    __args__ = dict()
    __args__['dbs'] = dbs
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['ips'] = ips
    __args__['keys'] = keys
    __args__['md5'] = md5
    __args__['product'] = product
    __args__['resultOutputFile'] = result_output_file
    __args__['startTime'] = start_time
    __args__['times'] = times
    __args__['users'] = users
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dbbrain/getSlowLogs:getSlowLogs', __args__, opts=opts, typ=GetSlowLogsResult).value

    return AwaitableGetSlowLogsResult(
        dbs=pulumi.get(__ret__, 'dbs'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        ips=pulumi.get(__ret__, 'ips'),
        keys=pulumi.get(__ret__, 'keys'),
        md5=pulumi.get(__ret__, 'md5'),
        product=pulumi.get(__ret__, 'product'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        rows=pulumi.get(__ret__, 'rows'),
        start_time=pulumi.get(__ret__, 'start_time'),
        times=pulumi.get(__ret__, 'times'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_slow_logs)
def get_slow_logs_output(dbs: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         end_time: Optional[pulumi.Input[str]] = None,
                         instance_id: Optional[pulumi.Input[str]] = None,
                         ips: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         keys: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         md5: Optional[pulumi.Input[str]] = None,
                         product: Optional[pulumi.Input[str]] = None,
                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         start_time: Optional[pulumi.Input[str]] = None,
                         times: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                         users: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSlowLogsResult]:
    """
    Use this data source to query detailed information of dbbrain slow_logs

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    slow_logs = tencentcloud.Dbbrain.get_slow_logs(end_time="%s",
        instance_id="%s",
        md5="4961208426639258265",
        product="mysql",
        start_time="%s")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] dbs: database list.
    :param str end_time: The deadline, such as 2019-09-11 10:13:14, the interval between the deadline and the start time is less than 7 days.
    :param str instance_id: instance Id.
    :param Sequence[str] ips: ip.
    :param Sequence[str] keys: keywords.
    :param str md5: md5 value of sql template.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str start_time: Start time, such as 2019-09-10 12:13:14.
    :param Sequence[int] times: Time-consuming interval, the left and right boundaries of the time-consuming interval correspond to the 0th element and the first element of the array respectively.
    :param Sequence[str] users: user.
    """
    ...
