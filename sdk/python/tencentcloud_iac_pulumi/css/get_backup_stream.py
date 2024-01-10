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
    'GetBackupStreamResult',
    'AwaitableGetBackupStreamResult',
    'get_backup_stream',
    'get_backup_stream_output',
]

@pulumi.output_type
class GetBackupStreamResult:
    """
    A collection of values returned by getBackupStream.
    """
    def __init__(__self__, id=None, result_output_file=None, stream_info_lists=None, stream_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if stream_info_lists and not isinstance(stream_info_lists, list):
            raise TypeError("Expected argument 'stream_info_lists' to be a list")
        pulumi.set(__self__, "stream_info_lists", stream_info_lists)
        if stream_name and not isinstance(stream_name, str):
            raise TypeError("Expected argument 'stream_name' to be a str")
        pulumi.set(__self__, "stream_name", stream_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="streamInfoLists")
    def stream_info_lists(self) -> Sequence['outputs.GetBackupStreamStreamInfoListResult']:
        """
        Backup stream group info.
        """
        return pulumi.get(self, "stream_info_lists")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[str]:
        """
        Stream name.
        """
        return pulumi.get(self, "stream_name")


class AwaitableGetBackupStreamResult(GetBackupStreamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupStreamResult(
            id=self.id,
            result_output_file=self.result_output_file,
            stream_info_lists=self.stream_info_lists,
            stream_name=self.stream_name)


def get_backup_stream(result_output_file: Optional[str] = None,
                      stream_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupStreamResult:
    """
    Use this data source to query detailed information of css backup_stream

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_stream = tencentcloud.Css.get_backup_stream(stream_name="live")
    ```


    :param str result_output_file: Used to save results.
    :param str stream_name: Stream id.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    __args__['streamName'] = stream_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Css/getBackupStream:getBackupStream', __args__, opts=opts, typ=GetBackupStreamResult).value

    return AwaitableGetBackupStreamResult(
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        stream_info_lists=__ret__.stream_info_lists,
        stream_name=__ret__.stream_name)


@_utilities.lift_output_func(get_backup_stream)
def get_backup_stream_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             stream_name: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupStreamResult]:
    """
    Use this data source to query detailed information of css backup_stream

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_stream = tencentcloud.Css.get_backup_stream(stream_name="live")
    ```


    :param str result_output_file: Used to save results.
    :param str stream_name: Stream id.
    """
    ...
