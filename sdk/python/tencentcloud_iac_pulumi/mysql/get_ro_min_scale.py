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
    'GetRoMinScaleResult',
    'AwaitableGetRoMinScaleResult',
    'get_ro_min_scale',
    'get_ro_min_scale_output',
]

@pulumi.output_type
class GetRoMinScaleResult:
    """
    A collection of values returned by getRoMinScale.
    """
    def __init__(__self__, id=None, master_instance_id=None, memory=None, result_output_file=None, ro_instance_id=None, volume=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if master_instance_id and not isinstance(master_instance_id, str):
            raise TypeError("Expected argument 'master_instance_id' to be a str")
        pulumi.set(__self__, "master_instance_id", master_instance_id)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if ro_instance_id and not isinstance(ro_instance_id, str):
            raise TypeError("Expected argument 'ro_instance_id' to be a str")
        pulumi.set(__self__, "ro_instance_id", ro_instance_id)
        if volume and not isinstance(volume, int):
            raise TypeError("Expected argument 'volume' to be a int")
        pulumi.set(__self__, "volume", volume)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="masterInstanceId")
    def master_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "master_instance_id")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        Memory specification size, unit: MB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="roInstanceId")
    def ro_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "ro_instance_id")

    @property
    @pulumi.getter
    def volume(self) -> int:
        """
        Disk specification size, unit: GB.
        """
        return pulumi.get(self, "volume")


class AwaitableGetRoMinScaleResult(GetRoMinScaleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoMinScaleResult(
            id=self.id,
            master_instance_id=self.master_instance_id,
            memory=self.memory,
            result_output_file=self.result_output_file,
            ro_instance_id=self.ro_instance_id,
            volume=self.volume)


def get_ro_min_scale(master_instance_id: Optional[str] = None,
                     result_output_file: Optional[str] = None,
                     ro_instance_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoMinScaleResult:
    """
    Use this data source to query detailed information of mysql ro_min_scale

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ro_min_scale = tencentcloud.Mysql.get_ro_min_scale(master_instance_id="cdb-fitq5t9h")
    ```
    <!--End PulumiCodeChooser -->


    :param str master_instance_id: The primary instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the RoInstanceId parameter cannot be empty at the same time. Note that when the input parameter contains RoInstanceId, the return value is the minimum specification when the read-only instance is upgraded; when the input parameter only contains MasterInstanceId, the return value is the minimum specification when the read-only instance is purchased.
    :param str result_output_file: Used to save results.
    :param str ro_instance_id: The read-only instance ID, in the format: cdbro-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the MasterInstanceId parameter cannot be empty at the same time.
    """
    __args__ = dict()
    __args__['masterInstanceId'] = master_instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['roInstanceId'] = ro_instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mysql/getRoMinScale:getRoMinScale', __args__, opts=opts, typ=GetRoMinScaleResult).value

    return AwaitableGetRoMinScaleResult(
        id=pulumi.get(__ret__, 'id'),
        master_instance_id=pulumi.get(__ret__, 'master_instance_id'),
        memory=pulumi.get(__ret__, 'memory'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        ro_instance_id=pulumi.get(__ret__, 'ro_instance_id'),
        volume=pulumi.get(__ret__, 'volume'))


@_utilities.lift_output_func(get_ro_min_scale)
def get_ro_min_scale_output(master_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                            result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            ro_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoMinScaleResult]:
    """
    Use this data source to query detailed information of mysql ro_min_scale

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    ro_min_scale = tencentcloud.Mysql.get_ro_min_scale(master_instance_id="cdb-fitq5t9h")
    ```
    <!--End PulumiCodeChooser -->


    :param str master_instance_id: The primary instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the RoInstanceId parameter cannot be empty at the same time. Note that when the input parameter contains RoInstanceId, the return value is the minimum specification when the read-only instance is upgraded; when the input parameter only contains MasterInstanceId, the return value is the minimum specification when the read-only instance is purchased.
    :param str result_output_file: Used to save results.
    :param str ro_instance_id: The read-only instance ID, in the format: cdbro-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the MasterInstanceId parameter cannot be empty at the same time.
    """
    ...
