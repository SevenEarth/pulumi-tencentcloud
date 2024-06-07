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
    'GetInstanceZoneInfoResult',
    'AwaitableGetInstanceZoneInfoResult',
    'get_instance_zone_info',
    'get_instance_zone_info_output',
]

@pulumi.output_type
class GetInstanceZoneInfoResult:
    """
    A collection of values returned by getInstanceZoneInfo.
    """
    def __init__(__self__, id=None, instance_id=None, replica_groups=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if replica_groups and not isinstance(replica_groups, list):
            raise TypeError("Expected argument 'replica_groups' to be a list")
        pulumi.set(__self__, "replica_groups", replica_groups)
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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="replicaGroups")
    def replica_groups(self) -> Sequence['outputs.GetInstanceZoneInfoReplicaGroupResult']:
        """
        List of instance node groups.
        """
        return pulumi.get(self, "replica_groups")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceZoneInfoResult(GetInstanceZoneInfoResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceZoneInfoResult(
            id=self.id,
            instance_id=self.instance_id,
            replica_groups=self.replica_groups,
            result_output_file=self.result_output_file)


def get_instance_zone_info(instance_id: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceZoneInfoResult:
    """
    Use this data source to query detailed information of redis instance_zone_info

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_zone_info = tencentcloud.Redis.get_instance_zone_info(instance_id="crs-c1nl9rpv")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: The ID of instance.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Redis/getInstanceZoneInfo:getInstanceZoneInfo', __args__, opts=opts, typ=GetInstanceZoneInfoResult).value

    return AwaitableGetInstanceZoneInfoResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        replica_groups=pulumi.get(__ret__, 'replica_groups'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_instance_zone_info)
def get_instance_zone_info_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceZoneInfoResult]:
    """
    Use this data source to query detailed information of redis instance_zone_info

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_zone_info = tencentcloud.Redis.get_instance_zone_info(instance_id="crs-c1nl9rpv")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: The ID of instance.
    :param str result_output_file: Used to save results.
    """
    ...
