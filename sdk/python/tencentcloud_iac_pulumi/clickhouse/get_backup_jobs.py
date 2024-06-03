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
    'GetBackupJobsResult',
    'AwaitableGetBackupJobsResult',
    'get_backup_jobs',
    'get_backup_jobs_output',
]

@pulumi.output_type
class GetBackupJobsResult:
    """
    A collection of values returned by getBackupJobs.
    """
    def __init__(__self__, back_up_jobs=None, begin_time=None, end_time=None, id=None, instance_id=None, result_output_file=None):
        if back_up_jobs and not isinstance(back_up_jobs, list):
            raise TypeError("Expected argument 'back_up_jobs' to be a list")
        pulumi.set(__self__, "back_up_jobs", back_up_jobs)
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
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="backUpJobs")
    def back_up_jobs(self) -> Sequence['outputs.GetBackupJobsBackUpJobResult']:
        """
        Back up jobs.
        """
        return pulumi.get(self, "back_up_jobs")

    @property
    @pulumi.getter(name="beginTime")
    def begin_time(self) -> Optional[str]:
        return pulumi.get(self, "begin_time")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
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
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetBackupJobsResult(GetBackupJobsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupJobsResult(
            back_up_jobs=self.back_up_jobs,
            begin_time=self.begin_time,
            end_time=self.end_time,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file)


def get_backup_jobs(begin_time: Optional[str] = None,
                    end_time: Optional[str] = None,
                    instance_id: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupJobsResult:
    """
    Use this data source to query detailed information of clickhouse backup jobs

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_jobs = tencentcloud.Clickhouse.get_backup_jobs(instance_id="cdwch-xxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str begin_time: Begin time.
    :param str end_time: End time.
    :param str instance_id: Instance id.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['beginTime'] = begin_time
    __args__['endTime'] = end_time
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Clickhouse/getBackupJobs:getBackupJobs', __args__, opts=opts, typ=GetBackupJobsResult).value

    return AwaitableGetBackupJobsResult(
        back_up_jobs=pulumi.get(__ret__, 'back_up_jobs'),
        begin_time=pulumi.get(__ret__, 'begin_time'),
        end_time=pulumi.get(__ret__, 'end_time'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_backup_jobs)
def get_backup_jobs_output(begin_time: Optional[pulumi.Input[Optional[str]]] = None,
                           end_time: Optional[pulumi.Input[Optional[str]]] = None,
                           instance_id: Optional[pulumi.Input[str]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupJobsResult]:
    """
    Use this data source to query detailed information of clickhouse backup jobs

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_jobs = tencentcloud.Clickhouse.get_backup_jobs(instance_id="cdwch-xxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str begin_time: Begin time.
    :param str end_time: End time.
    :param str instance_id: Instance id.
    :param str result_output_file: Used to save results.
    """
    ...
