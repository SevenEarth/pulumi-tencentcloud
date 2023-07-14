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
    'GetInsAttributeResult',
    'AwaitableGetInsAttributeResult',
    'get_ins_attribute',
    'get_ins_attribute_output',
]

@pulumi.output_type
class GetInsAttributeResult:
    """
    A collection of values returned by getInsAttribute.
    """
    def __init__(__self__, blocked_threshold=None, event_save_days=None, id=None, instance_id=None, regular_backup_counts=None, regular_backup_enable=None, regular_backup_save_days=None, regular_backup_start_time=None, regular_backup_strategy=None, result_output_file=None, tde_configs=None):
        if blocked_threshold and not isinstance(blocked_threshold, int):
            raise TypeError("Expected argument 'blocked_threshold' to be a int")
        pulumi.set(__self__, "blocked_threshold", blocked_threshold)
        if event_save_days and not isinstance(event_save_days, int):
            raise TypeError("Expected argument 'event_save_days' to be a int")
        pulumi.set(__self__, "event_save_days", event_save_days)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if regular_backup_counts and not isinstance(regular_backup_counts, int):
            raise TypeError("Expected argument 'regular_backup_counts' to be a int")
        pulumi.set(__self__, "regular_backup_counts", regular_backup_counts)
        if regular_backup_enable and not isinstance(regular_backup_enable, str):
            raise TypeError("Expected argument 'regular_backup_enable' to be a str")
        pulumi.set(__self__, "regular_backup_enable", regular_backup_enable)
        if regular_backup_save_days and not isinstance(regular_backup_save_days, int):
            raise TypeError("Expected argument 'regular_backup_save_days' to be a int")
        pulumi.set(__self__, "regular_backup_save_days", regular_backup_save_days)
        if regular_backup_start_time and not isinstance(regular_backup_start_time, str):
            raise TypeError("Expected argument 'regular_backup_start_time' to be a str")
        pulumi.set(__self__, "regular_backup_start_time", regular_backup_start_time)
        if regular_backup_strategy and not isinstance(regular_backup_strategy, str):
            raise TypeError("Expected argument 'regular_backup_strategy' to be a str")
        pulumi.set(__self__, "regular_backup_strategy", regular_backup_strategy)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if tde_configs and not isinstance(tde_configs, list):
            raise TypeError("Expected argument 'tde_configs' to be a list")
        pulumi.set(__self__, "tde_configs", tde_configs)

    @property
    @pulumi.getter(name="blockedThreshold")
    def blocked_threshold(self) -> int:
        """
        Block process threshold in milliseconds.
        """
        return pulumi.get(self, "blocked_threshold")

    @property
    @pulumi.getter(name="eventSaveDays")
    def event_save_days(self) -> int:
        """
        Retention period for the files of slow SQL, blocking, deadlock, and extended events.
        """
        return pulumi.get(self, "event_save_days")

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
    @pulumi.getter(name="regularBackupCounts")
    def regular_backup_counts(self) -> int:
        """
        The number of retained archive backups.
        """
        return pulumi.get(self, "regular_backup_counts")

    @property
    @pulumi.getter(name="regularBackupEnable")
    def regular_backup_enable(self) -> str:
        """
        Archive backup status. Valid values: enable (enabled), disable (disabled).
        """
        return pulumi.get(self, "regular_backup_enable")

    @property
    @pulumi.getter(name="regularBackupSaveDays")
    def regular_backup_save_days(self) -> int:
        """
        Archive backup retention period: [90-3650] days.
        """
        return pulumi.get(self, "regular_backup_save_days")

    @property
    @pulumi.getter(name="regularBackupStartTime")
    def regular_backup_start_time(self) -> str:
        """
        Archive backup start date in YYYY-MM-DD format, which is the current time by default.
        """
        return pulumi.get(self, "regular_backup_start_time")

    @property
    @pulumi.getter(name="regularBackupStrategy")
    def regular_backup_strategy(self) -> str:
        """
        Archive backup policy. Valid values: years (yearly); quarters (quarterly);months` (monthly).
        """
        return pulumi.get(self, "regular_backup_strategy")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="tdeConfigs")
    def tde_configs(self) -> Sequence['outputs.GetInsAttributeTdeConfigResult']:
        """
        TDE Transparent Data Encryption Configuration.
        """
        return pulumi.get(self, "tde_configs")


class AwaitableGetInsAttributeResult(GetInsAttributeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInsAttributeResult(
            blocked_threshold=self.blocked_threshold,
            event_save_days=self.event_save_days,
            id=self.id,
            instance_id=self.instance_id,
            regular_backup_counts=self.regular_backup_counts,
            regular_backup_enable=self.regular_backup_enable,
            regular_backup_save_days=self.regular_backup_save_days,
            regular_backup_start_time=self.regular_backup_start_time,
            regular_backup_strategy=self.regular_backup_strategy,
            result_output_file=self.result_output_file,
            tde_configs=self.tde_configs)


def get_ins_attribute(instance_id: Optional[str] = None,
                      result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInsAttributeResult:
    """
    Use this data source to query detailed information of sqlserver_ins_attribute

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ins_attribute = tencentcloud.Sqlserver.get_ins_attribute(instance_id="mssql-gyg9xycl")
    ```


    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getInsAttribute:getInsAttribute', __args__, opts=opts, typ=GetInsAttributeResult).value

    return AwaitableGetInsAttributeResult(
        blocked_threshold=__ret__.blocked_threshold,
        event_save_days=__ret__.event_save_days,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        regular_backup_counts=__ret__.regular_backup_counts,
        regular_backup_enable=__ret__.regular_backup_enable,
        regular_backup_save_days=__ret__.regular_backup_save_days,
        regular_backup_start_time=__ret__.regular_backup_start_time,
        regular_backup_strategy=__ret__.regular_backup_strategy,
        result_output_file=__ret__.result_output_file,
        tde_configs=__ret__.tde_configs)


@_utilities.lift_output_func(get_ins_attribute)
def get_ins_attribute_output(instance_id: Optional[pulumi.Input[str]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInsAttributeResult]:
    """
    Use this data source to query detailed information of sqlserver_ins_attribute

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ins_attribute = tencentcloud.Sqlserver.get_ins_attribute(instance_id="mssql-gyg9xycl")
    ```


    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    """
    ...
