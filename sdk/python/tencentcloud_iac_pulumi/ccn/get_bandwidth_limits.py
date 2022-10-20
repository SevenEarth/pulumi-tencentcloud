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
    'GetBandwidthLimitsResult',
    'AwaitableGetBandwidthLimitsResult',
    'get_bandwidth_limits',
    'get_bandwidth_limits_output',
]

@pulumi.output_type
class GetBandwidthLimitsResult:
    """
    A collection of values returned by getBandwidthLimits.
    """
    def __init__(__self__, ccn_id=None, id=None, limits=None, result_output_file=None):
        if ccn_id and not isinstance(ccn_id, str):
            raise TypeError("Expected argument 'ccn_id' to be a str")
        pulumi.set(__self__, "ccn_id", ccn_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limits and not isinstance(limits, list):
            raise TypeError("Expected argument 'limits' to be a list")
        pulumi.set(__self__, "limits", limits)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> str:
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def limits(self) -> Sequence['outputs.GetBandwidthLimitsLimitResult']:
        """
        The bandwidth limits of regions:
        """
        return pulumi.get(self, "limits")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetBandwidthLimitsResult(GetBandwidthLimitsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBandwidthLimitsResult(
            ccn_id=self.ccn_id,
            id=self.id,
            limits=self.limits,
            result_output_file=self.result_output_file)


def get_bandwidth_limits(ccn_id: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBandwidthLimitsResult:
    """
    Use this data source to query detailed information of CCN bandwidth limits.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    config = pulumi.Config()
    other_region1 = config.get("otherRegion1")
    if other_region1 is None:
        other_region1 = "ap-shanghai"
    main = tencentcloud.ccn.Instance("main",
        description="ci-temp-test-ccn-des",
        qos="AG")
    limit = tencentcloud.Ccn.get_bandwidth_limits_output(ccn_id=main.id)
    limit1 = tencentcloud.ccn.BandwidthLimit("limit1",
        ccn_id=main.id,
        region=other_region1,
        bandwidth_limit=500)
    ```


    :param str ccn_id: ID of the CCN to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['ccnId'] = ccn_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ccn/getBandwidthLimits:getBandwidthLimits', __args__, opts=opts, typ=GetBandwidthLimitsResult).value

    return AwaitableGetBandwidthLimitsResult(
        ccn_id=__ret__.ccn_id,
        id=__ret__.id,
        limits=__ret__.limits,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_bandwidth_limits)
def get_bandwidth_limits_output(ccn_id: Optional[pulumi.Input[str]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBandwidthLimitsResult]:
    """
    Use this data source to query detailed information of CCN bandwidth limits.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    config = pulumi.Config()
    other_region1 = config.get("otherRegion1")
    if other_region1 is None:
        other_region1 = "ap-shanghai"
    main = tencentcloud.ccn.Instance("main",
        description="ci-temp-test-ccn-des",
        qos="AG")
    limit = tencentcloud.Ccn.get_bandwidth_limits_output(ccn_id=main.id)
    limit1 = tencentcloud.ccn.BandwidthLimit("limit1",
        ccn_id=main.id,
        region=other_region1,
        bandwidth_limit=500)
    ```


    :param str ccn_id: ID of the CCN to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
