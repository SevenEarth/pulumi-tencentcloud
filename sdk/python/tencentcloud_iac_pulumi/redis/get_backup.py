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
    'GetBackupResult',
    'AwaitableGetBackupResult',
    'get_backup',
    'get_backup_output',
]

@pulumi.output_type
class GetBackupResult:
    """
    A collection of values returned by getBackup.
    """
    def __init__(__self__, backup_sets=None, begin_time=None, end_time=None, id=None, instance_id=None, instance_name=None, result_output_file=None, statuses=None):
        if backup_sets and not isinstance(backup_sets, list):
            raise TypeError("Expected argument 'backup_sets' to be a list")
        pulumi.set(__self__, "backup_sets", backup_sets)
        if begin_time and not isinstance(begin_time, str):
            raise TypeError("Expected argument 'begin_time' to be a str")
        pulumi.set(__self__, "begin_time", begin_time)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)

    @property
    @pulumi.getter(name="backupSets")
    def backup_sets(self) -> Sequence['outputs.GetBackupBackupSetResult']:
        """
        An array of backups for the instance.
        """
        return pulumi.get(self, "backup_sets")

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> Optional[str]:
        return pulumi.get(self, "begin_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        """
        Backup end time.
        """
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
    def instance_id(self) -> Optional[str]:
        """
        The ID of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        """
        The name of instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[int]]:
        """
        Backup status.1: The backup is locked by another process.2: The backup is normal and not locked by any process.-1: The backup has expired.3: The backup is being exported.4: The backup export is successful.
        """
        return pulumi.get(self, "statuses")


class AwaitableGetBackupResult(GetBackupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupResult(
            backup_sets=self.backup_sets,
            begin_time=self.begin_time,
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            instance_name=self.instance_name,
            result_output_file=self.result_output_file,
            statuses=self.statuses)


def get_backup(begin_time: Optional[str] = None,
               end_time: Optional[str] = None,
               instance_id: Optional[str] = None,
               instance_name: Optional[str] = None,
               result_output_file: Optional[str] = None,
               statuses: Optional[Sequence[int]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupResult:
    """
    Use this data source to query detailed information of redis backup

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup = tencentcloud.Redis.get_backup(begin_time="2023-04-07 03:57:30",
        end_time="2023-04-07 03:57:56",
        instance_id="crs-c1nl9rpv",
        instance_name="Keep-terraform",
        statuses=[2])
    ```


    :param str begin_time: start time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
    :param str end_time: End time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
    :param str instance_id: The ID of instance.
    :param str instance_name: Instance name, which supports fuzzy search based on instance name.
    :param str result_output_file: Used to save results.
    :param Sequence[int] statuses: Status of the backup task:1: Backup is in the process.2: The backup is normal.3: Backup to RDB file processing.4: RDB conversion completed.-1: The backup has expired.-2: Backup deleted.
    """
    __args__ = dict()
    __args__['beginTime'] = begin_time
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['instanceName'] = instance_name
    __args__['resultOutputFile'] = result_output_file
    __args__['statuses'] = statuses
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Redis/getBackup:getBackup', __args__, opts=opts, typ=GetBackupResult).value

    return AwaitableGetBackupResult(
        backup_sets=__ret__.backup_sets,
        begin_time=__ret__.begin_time,
        end_time=__ret__.end_time,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_name=__ret__.instance_name,
        result_output_file=__ret__.result_output_file,
        statuses=__ret__.statuses)


@_utilities.lift_output_func(get_backup)
def get_backup_output(begin_time: Optional[pulumi.Input[Optional[str]]] = None,
                      end_time: Optional[pulumi.Input[Optional[str]]] = None,
                      instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                      instance_name: Optional[pulumi.Input[Optional[str]]] = None,
                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      statuses: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupResult]:
    """
    Use this data source to query detailed information of redis backup

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup = tencentcloud.Redis.get_backup(begin_time="2023-04-07 03:57:30",
        end_time="2023-04-07 03:57:56",
        instance_id="crs-c1nl9rpv",
        instance_name="Keep-terraform",
        statuses=[2])
    ```


    :param str begin_time: start time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
    :param str end_time: End time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.
    :param str instance_id: The ID of instance.
    :param str instance_name: Instance name, which supports fuzzy search based on instance name.
    :param str result_output_file: Used to save results.
    :param Sequence[int] statuses: Status of the backup task:1: Backup is in the process.2: The backup is normal.3: Backup to RDB file processing.4: RDB conversion completed.-1: The backup has expired.-2: Backup deleted.
    """
    ...