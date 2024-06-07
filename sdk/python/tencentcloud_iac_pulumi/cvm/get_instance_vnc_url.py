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
    'GetInstanceVncUrlResult',
    'AwaitableGetInstanceVncUrlResult',
    'get_instance_vnc_url',
    'get_instance_vnc_url_output',
]

@pulumi.output_type
class GetInstanceVncUrlResult:
    """
    A collection of values returned by getInstanceVncUrl.
    """
    def __init__(__self__, id=None, instance_id=None, instance_vnc_url=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_vnc_url and not isinstance(instance_vnc_url, str):
            raise TypeError("Expected argument 'instance_vnc_url' to be a str")
        pulumi.set(__self__, "instance_vnc_url", instance_vnc_url)
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
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceVncUrl")
    def instance_vnc_url(self) -> str:
        """
        Instance VNC URL.
        """
        return pulumi.get(self, "instance_vnc_url")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceVncUrlResult(GetInstanceVncUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceVncUrlResult(
            id=self.id,
            instance_id=self.instance_id,
            instance_vnc_url=self.instance_vnc_url,
            result_output_file=self.result_output_file)


def get_instance_vnc_url(instance_id: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceVncUrlResult:
    """
    Use this data source to query detailed information of cvm instance_vnc_url

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_vnc_url = tencentcloud.Cvm.get_instance_vnc_url(instance_id="ins-xxxxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: Instance ID. To obtain the instance IDs, you can call `DescribeInstances` and look for `InstanceId` in the response.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cvm/getInstanceVncUrl:getInstanceVncUrl', __args__, opts=opts, typ=GetInstanceVncUrlResult).value

    return AwaitableGetInstanceVncUrlResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        instance_vnc_url=pulumi.get(__ret__, 'instance_vnc_url'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_instance_vnc_url)
def get_instance_vnc_url_output(instance_id: Optional[pulumi.Input[str]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceVncUrlResult]:
    """
    Use this data source to query detailed information of cvm instance_vnc_url

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_vnc_url = tencentcloud.Cvm.get_instance_vnc_url(instance_id="ins-xxxxxxxx")
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: Instance ID. To obtain the instance IDs, you can call `DescribeInstances` and look for `InstanceId` in the response.
    :param str result_output_file: Used to save results.
    """
    ...
