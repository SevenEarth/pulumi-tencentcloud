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
    'GetInstanceDiskNumResult',
    'AwaitableGetInstanceDiskNumResult',
    'get_instance_disk_num',
    'get_instance_disk_num_output',
]

@pulumi.output_type
class GetInstanceDiskNumResult:
    """
    A collection of values returned by getInstanceDiskNum.
    """
    def __init__(__self__, attach_detail_sets=None, id=None, instance_ids=None, result_output_file=None):
        if attach_detail_sets and not isinstance(attach_detail_sets, list):
            raise TypeError("Expected argument 'attach_detail_sets' to be a list")
        pulumi.set(__self__, "attach_detail_sets", attach_detail_sets)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_ids and not isinstance(instance_ids, list):
            raise TypeError("Expected argument 'instance_ids' to be a list")
        pulumi.set(__self__, "instance_ids", instance_ids)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="attachDetailSets")
    def attach_detail_sets(self) -> Sequence['outputs.GetInstanceDiskNumAttachDetailSetResult']:
        """
        Mount information list.
        """
        return pulumi.get(self, "attach_detail_sets")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceIds")
    def instance_ids(self) -> Sequence[str]:
        return pulumi.get(self, "instance_ids")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceDiskNumResult(GetInstanceDiskNumResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceDiskNumResult(
            attach_detail_sets=self.attach_detail_sets,
            id=self.id,
            instance_ids=self.instance_ids,
            result_output_file=self.result_output_file)


def get_instance_disk_num(instance_ids: Optional[Sequence[str]] = None,
                          result_output_file: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceDiskNumResult:
    """
    Use this data source to query detailed information of lighthouse instance_disk_num

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_disk_num = tencentcloud.Lighthouse.get_instance_disk_num(instance_ids=["lhins-xxxxxx"])
    ```


    :param Sequence[str] instance_ids: List of instance IDs.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceIds'] = instance_ids
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Lighthouse/getInstanceDiskNum:getInstanceDiskNum', __args__, opts=opts, typ=GetInstanceDiskNumResult).value

    return AwaitableGetInstanceDiskNumResult(
        attach_detail_sets=__ret__.attach_detail_sets,
        id=__ret__.id,
        instance_ids=__ret__.instance_ids,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instance_disk_num)
def get_instance_disk_num_output(instance_ids: Optional[pulumi.Input[Sequence[str]]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceDiskNumResult]:
    """
    Use this data source to query detailed information of lighthouse instance_disk_num

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_disk_num = tencentcloud.Lighthouse.get_instance_disk_num(instance_ids=["lhins-xxxxxx"])
    ```


    :param Sequence[str] instance_ids: List of instance IDs.
    :param str result_output_file: Used to save results.
    """
    ...
