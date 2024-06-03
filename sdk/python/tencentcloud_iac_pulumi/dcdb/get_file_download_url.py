# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFileDownloadUrlResult',
    'AwaitableGetFileDownloadUrlResult',
    'get_file_download_url',
    'get_file_download_url_output',
]

@pulumi.output_type
class GetFileDownloadUrlResult:
    """
    A collection of values returned by getFileDownloadUrl.
    """
    def __init__(__self__, file_path=None, id=None, instance_id=None, pre_signed_url=None, result_output_file=None, shard_id=None):
        if file_path and not isinstance(file_path, str):
            raise TypeError("Expected argument 'file_path' to be a str")
        pulumi.set(__self__, "file_path", file_path)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if pre_signed_url and not isinstance(pre_signed_url, str):
            raise TypeError("Expected argument 'pre_signed_url' to be a str")
        pulumi.set(__self__, "pre_signed_url", pre_signed_url)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if shard_id and not isinstance(shard_id, str):
            raise TypeError("Expected argument 'shard_id' to be a str")
        pulumi.set(__self__, "shard_id", shard_id)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> str:
        return pulumi.get(self, "file_path")

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
    @pulumi.getter(name="preSignedUrl")
    def pre_signed_url(self) -> str:
        """
        Signed download URL.
        """
        return pulumi.get(self, "pre_signed_url")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="shardId")
    def shard_id(self) -> str:
        return pulumi.get(self, "shard_id")


class AwaitableGetFileDownloadUrlResult(GetFileDownloadUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileDownloadUrlResult(
            file_path=self.file_path,
            id=self.id,
            instance_id=self.instance_id,
            pre_signed_url=self.pre_signed_url,
            result_output_file=self.result_output_file,
            shard_id=self.shard_id)


def get_file_download_url(file_path: Optional[str] = None,
                          instance_id: Optional[str] = None,
                          result_output_file: Optional[str] = None,
                          shard_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileDownloadUrlResult:
    """
    Use this data source to query detailed information of dcdb file_download_url

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    file_download_url = tencentcloud.Dcdb.get_file_download_url(instance_id=local["dcdb_id"],
        shard_id="shard-1b5r04az",
        file_path="/cos_backup/test.txt")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_path: Unsigned file path.
    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    :param str shard_id: Instance Shard ID.
    """
    __args__ = dict()
    __args__['filePath'] = file_path
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['shardId'] = shard_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dcdb/getFileDownloadUrl:getFileDownloadUrl', __args__, opts=opts, typ=GetFileDownloadUrlResult).value

    return AwaitableGetFileDownloadUrlResult(
        file_path=pulumi.get(__ret__, 'file_path'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        pre_signed_url=pulumi.get(__ret__, 'pre_signed_url'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        shard_id=pulumi.get(__ret__, 'shard_id'))


@_utilities.lift_output_func(get_file_download_url)
def get_file_download_url_output(file_path: Optional[pulumi.Input[str]] = None,
                                 instance_id: Optional[pulumi.Input[str]] = None,
                                 result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                 shard_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileDownloadUrlResult]:
    """
    Use this data source to query detailed information of dcdb file_download_url

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    file_download_url = tencentcloud.Dcdb.get_file_download_url(instance_id=local["dcdb_id"],
        shard_id="shard-1b5r04az",
        file_path="/cos_backup/test.txt")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_path: Unsigned file path.
    :param str instance_id: Instance ID.
    :param str result_output_file: Used to save results.
    :param str shard_id: Instance Shard ID.
    """
    ...
