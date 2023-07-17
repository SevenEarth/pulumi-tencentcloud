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
    'GetShardSpecResult',
    'AwaitableGetShardSpecResult',
    'get_shard_spec',
    'get_shard_spec_output',
]

@pulumi.output_type
class GetShardSpecResult:
    """
    A collection of values returned by getShardSpec.
    """
    def __init__(__self__, id=None, result_output_file=None, spec_configs=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if spec_configs and not isinstance(spec_configs, list):
            raise TypeError("Expected argument 'spec_configs' to be a list")
        pulumi.set(__self__, "spec_configs", spec_configs)

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
    @pulumi.getter(name="specConfigs")
    def spec_configs(self) -> Sequence['outputs.GetShardSpecSpecConfigResult']:
        """
        list of instance specifications.
        """
        return pulumi.get(self, "spec_configs")


class AwaitableGetShardSpecResult(GetShardSpecResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetShardSpecResult(
            id=self.id,
            result_output_file=self.result_output_file,
            spec_configs=self.spec_configs)


def get_shard_spec(result_output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetShardSpecResult:
    """
    Use this data source to query detailed information of dcdb shard_spec

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    shard_spec = tencentcloud.Dcdb.get_shard_spec()
    ```


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dcdb/getShardSpec:getShardSpec', __args__, opts=opts, typ=GetShardSpecResult).value

    return AwaitableGetShardSpecResult(
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        spec_configs=__ret__.spec_configs)


@_utilities.lift_output_func(get_shard_spec)
def get_shard_spec_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetShardSpecResult]:
    """
    Use this data source to query detailed information of dcdb shard_spec

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    shard_spec = tencentcloud.Dcdb.get_shard_spec()
    ```


    :param str result_output_file: Used to save results.
    """
    ...
