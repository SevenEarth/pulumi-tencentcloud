# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetBinlogDownloadUrlResult',
    'AwaitableGetBinlogDownloadUrlResult',
    'get_binlog_download_url',
    'get_binlog_download_url_output',
]

@pulumi.output_type
class GetBinlogDownloadUrlResult:
    """
    A collection of values returned by getBinlogDownloadUrl.
    """
    def __init__(__self__, binlog_id=None, cluster_id=None, download_url=None, id=None, result_output_file=None):
        if binlog_id and not isinstance(binlog_id, int):
            raise TypeError("Expected argument 'binlog_id' to be a int")
        pulumi.set(__self__, "binlog_id", binlog_id)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if download_url and not isinstance(download_url, str):
            raise TypeError("Expected argument 'download_url' to be a str")
        pulumi.set(__self__, "download_url", download_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="binlogId")
    def binlog_id(self) -> int:
        return pulumi.get(self, "binlog_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="downloadUrl")
    def download_url(self) -> str:
        """
        Download address.
        """
        return pulumi.get(self, "download_url")

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


class AwaitableGetBinlogDownloadUrlResult(GetBinlogDownloadUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBinlogDownloadUrlResult(
            binlog_id=self.binlog_id,
            cluster_id=self.cluster_id,
            download_url=self.download_url,
            id=self.id,
            result_output_file=self.result_output_file)


def get_binlog_download_url(binlog_id: Optional[int] = None,
                            cluster_id: Optional[str] = None,
                            result_output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBinlogDownloadUrlResult:
    """
    Use this data source to query detailed information of cynosdb binlog_download_url

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    binlog_download_url = tencentcloud.Cynosdb.get_binlog_download_url(binlog_id=6202249,
        cluster_id="cynosdbmysql-bws8h88b")
    ```


    :param int binlog_id: Binlog file ID.
    :param str cluster_id: Cluster ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['binlogId'] = binlog_id
    __args__['clusterId'] = cluster_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cynosdb/getBinlogDownloadUrl:getBinlogDownloadUrl', __args__, opts=opts, typ=GetBinlogDownloadUrlResult).value

    return AwaitableGetBinlogDownloadUrlResult(
        binlog_id=__ret__.binlog_id,
        cluster_id=__ret__.cluster_id,
        download_url=__ret__.download_url,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_binlog_download_url)
def get_binlog_download_url_output(binlog_id: Optional[pulumi.Input[int]] = None,
                                   cluster_id: Optional[pulumi.Input[str]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBinlogDownloadUrlResult]:
    """
    Use this data source to query detailed information of cynosdb binlog_download_url

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    binlog_download_url = tencentcloud.Cynosdb.get_binlog_download_url(binlog_id=6202249,
        cluster_id="cynosdbmysql-bws8h88b")
    ```


    :param int binlog_id: Binlog file ID.
    :param str cluster_id: Cluster ID.
    :param str result_output_file: Used to save results.
    """
    ...