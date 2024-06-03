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
    'GetStreamMonitorListResult',
    'AwaitableGetStreamMonitorListResult',
    'get_stream_monitor_list',
    'get_stream_monitor_list_output',
]

@pulumi.output_type
class GetStreamMonitorListResult:
    """
    A collection of values returned by getStreamMonitorList.
    """
    def __init__(__self__, id=None, live_stream_monitors=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if live_stream_monitors and not isinstance(live_stream_monitors, list):
            raise TypeError("Expected argument 'live_stream_monitors' to be a list")
        pulumi.set(__self__, "live_stream_monitors", live_stream_monitors)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="liveStreamMonitors")
    def live_stream_monitors(self) -> Sequence['outputs.GetStreamMonitorListLiveStreamMonitorResult']:
        """
        The list of live stream monitoring tasks.Note: This field may return null, indicating that no valid value is available.
        """
        return pulumi.get(self, "live_stream_monitors")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetStreamMonitorListResult(GetStreamMonitorListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamMonitorListResult(
            id=self.id,
            live_stream_monitors=self.live_stream_monitors,
            result_output_file=self.result_output_file)


def get_stream_monitor_list(result_output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamMonitorListResult:
    """
    Use this data source to query detailed information of css stream_monitor_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    stream_monitor_list = tencentcloud.Css.get_stream_monitor_list()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Css/getStreamMonitorList:getStreamMonitorList', __args__, opts=opts, typ=GetStreamMonitorListResult).value

    return AwaitableGetStreamMonitorListResult(
        id=pulumi.get(__ret__, 'id'),
        live_stream_monitors=pulumi.get(__ret__, 'live_stream_monitors'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_stream_monitor_list)
def get_stream_monitor_list_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStreamMonitorListResult]:
    """
    Use this data source to query detailed information of css stream_monitor_list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    stream_monitor_list = tencentcloud.Css.get_stream_monitor_list()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    ...
