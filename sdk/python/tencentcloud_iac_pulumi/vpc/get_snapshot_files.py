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
    'GetSnapshotFilesResult',
    'AwaitableGetSnapshotFilesResult',
    'get_snapshot_files',
    'get_snapshot_files_output',
]

@pulumi.output_type
class GetSnapshotFilesResult:
    """
    A collection of values returned by getSnapshotFiles.
    """
    def __init__(__self__, business_type=None, end_date=None, id=None, instance_id=None, result_output_file=None, snapshot_file_sets=None, start_date=None):
        if business_type and not isinstance(business_type, str):
            raise TypeError("Expected argument 'business_type' to be a str")
        pulumi.set(__self__, "business_type", business_type)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if snapshot_file_sets and not isinstance(snapshot_file_sets, list):
            raise TypeError("Expected argument 'snapshot_file_sets' to be a list")
        pulumi.set(__self__, "snapshot_file_sets", snapshot_file_sets)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)

    @property
    @pulumi.getter(name="businessType")
    def business_type(self) -> str:
        return pulumi.get(self, "business_type")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        return pulumi.get(self, "end_date")

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
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="snapshotFileSets")
    def snapshot_file_sets(self) -> Sequence['outputs.GetSnapshotFilesSnapshotFileSetResult']:
        """
        snap shot file set.
        """
        return pulumi.get(self, "snapshot_file_sets")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        return pulumi.get(self, "start_date")


class AwaitableGetSnapshotFilesResult(GetSnapshotFilesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotFilesResult(
            business_type=self.business_type,
            end_date=self.end_date,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            snapshot_file_sets=self.snapshot_file_sets,
            start_date=self.start_date)


def get_snapshot_files(business_type: Optional[str] = None,
                       end_date: Optional[str] = None,
                       instance_id: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       start_date: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotFilesResult:
    """
    Use this data source to query detailed information of vpc snapshot_files

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    snapshot_files = tencentcloud.Vpc.get_snapshot_files(business_type="securitygroup",
        end_date="2023-10-30 19:00:00",
        instance_id="sg-902tl7t7",
        start_date="2022-10-10 00:00:00")
    ```
    <!--End PulumiCodeChooser -->


    :param str business_type: Business type, currently supports security group:securitygroup.
    :param str end_date: End date in the format %Y-%m-%d %H:%M:%S.
    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    :param str start_date: Start date in the format %Y-%m-%d %H:%M:%S.
    """
    __args__ = dict()
    __args__['businessType'] = business_type
    __args__['endDate'] = end_date
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['startDate'] = start_date
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/getSnapshotFiles:getSnapshotFiles', __args__, opts=opts, typ=GetSnapshotFilesResult).value

    return AwaitableGetSnapshotFilesResult(
        business_type=pulumi.get(__ret__, 'business_type'),
        end_date=pulumi.get(__ret__, 'end_date'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        snapshot_file_sets=pulumi.get(__ret__, 'snapshot_file_sets'),
        start_date=pulumi.get(__ret__, 'start_date'))


@_utilities.lift_output_func(get_snapshot_files)
def get_snapshot_files_output(business_type: Optional[pulumi.Input[str]] = None,
                              end_date: Optional[pulumi.Input[str]] = None,
                              instance_id: Optional[pulumi.Input[str]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              start_date: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotFilesResult]:
    """
    Use this data source to query detailed information of vpc snapshot_files

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    snapshot_files = tencentcloud.Vpc.get_snapshot_files(business_type="securitygroup",
        end_date="2023-10-30 19:00:00",
        instance_id="sg-902tl7t7",
        start_date="2022-10-10 00:00:00")
    ```
    <!--End PulumiCodeChooser -->


    :param str business_type: Business type, currently supports security group:securitygroup.
    :param str end_date: End date in the format %Y-%m-%d %H:%M:%S.
    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    :param str start_date: Start date in the format %Y-%m-%d %H:%M:%S.
    """
    ...
