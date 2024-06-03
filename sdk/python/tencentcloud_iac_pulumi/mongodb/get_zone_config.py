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
    'GetZoneConfigResult',
    'AwaitableGetZoneConfigResult',
    'get_zone_config',
    'get_zone_config_output',
]

@pulumi.output_type
class GetZoneConfigResult:
    """
    A collection of values returned by getZoneConfig.
    """
    def __init__(__self__, available_zone=None, id=None, lists=None, result_output_file=None):
        if available_zone and not isinstance(available_zone, str):
            raise TypeError("Expected argument 'available_zone' to be a str")
        pulumi.set(__self__, "available_zone", available_zone)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="availableZone")
    def available_zone(self) -> Optional[str]:
        """
        The available zone of the Mongodb.
        """
        return pulumi.get(self, "available_zone")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetZoneConfigListResult']:
        """
        A list of zone config. Each element contains the following attributes:
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetZoneConfigResult(GetZoneConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneConfigResult(
            available_zone=self.available_zone,
            id=self.id,
            lists=self.lists,
            result_output_file=self.result_output_file)


def get_zone_config(available_zone: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneConfigResult:
    """
    Use this data source to query the available mongodb specifications for different zone.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    mongodb = tencentcloud.Mongodb.get_zone_config(available_zone="ap-guangzhou-2")
    ```
    <!--End PulumiCodeChooser -->


    :param str available_zone: The available zone of the Mongodb.
    :param str result_output_file: Used to store results.
    """
    __args__ = dict()
    __args__['availableZone'] = available_zone
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mongodb/getZoneConfig:getZoneConfig', __args__, opts=opts, typ=GetZoneConfigResult).value

    return AwaitableGetZoneConfigResult(
        available_zone=pulumi.get(__ret__, 'available_zone'),
        id=pulumi.get(__ret__, 'id'),
        lists=pulumi.get(__ret__, 'lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_zone_config)
def get_zone_config_output(available_zone: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneConfigResult]:
    """
    Use this data source to query the available mongodb specifications for different zone.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    mongodb = tencentcloud.Mongodb.get_zone_config(available_zone="ap-guangzhou-2")
    ```
    <!--End PulumiCodeChooser -->


    :param str available_zone: The available zone of the Mongodb.
    :param str result_output_file: Used to store results.
    """
    ...
