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
    'GetInstanceParamsResult',
    'AwaitableGetInstanceParamsResult',
    'get_instance_params',
    'get_instance_params_output',
]

@pulumi.output_type
class GetInstanceParamsResult:
    """
    A collection of values returned by getInstanceParams.
    """
    def __init__(__self__, id=None, instance_enum_params=None, instance_id=None, instance_integer_params=None, instance_multi_params=None, instance_text_params=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_enum_params and not isinstance(instance_enum_params, list):
            raise TypeError("Expected argument 'instance_enum_params' to be a list")
        pulumi.set(__self__, "instance_enum_params", instance_enum_params)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_integer_params and not isinstance(instance_integer_params, list):
            raise TypeError("Expected argument 'instance_integer_params' to be a list")
        pulumi.set(__self__, "instance_integer_params", instance_integer_params)
        if instance_multi_params and not isinstance(instance_multi_params, list):
            raise TypeError("Expected argument 'instance_multi_params' to be a list")
        pulumi.set(__self__, "instance_multi_params", instance_multi_params)
        if instance_text_params and not isinstance(instance_text_params, list):
            raise TypeError("Expected argument 'instance_text_params' to be a list")
        pulumi.set(__self__, "instance_text_params", instance_text_params)
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
    @pulumi.getter(name="instanceEnumParams")
    def instance_enum_params(self) -> Sequence['outputs.GetInstanceParamsInstanceEnumParamResult']:
        """
        Enum parameter.
        """
        return pulumi.get(self, "instance_enum_params")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceIntegerParams")
    def instance_integer_params(self) -> Sequence['outputs.GetInstanceParamsInstanceIntegerParamResult']:
        """
        Integer parameter.
        """
        return pulumi.get(self, "instance_integer_params")

    @property
    @pulumi.getter(name="instanceMultiParams")
    def instance_multi_params(self) -> Sequence['outputs.GetInstanceParamsInstanceMultiParamResult']:
        """
        multi parameter.
        """
        return pulumi.get(self, "instance_multi_params")

    @property
    @pulumi.getter(name="instanceTextParams")
    def instance_text_params(self) -> Sequence['outputs.GetInstanceParamsInstanceTextParamResult']:
        """
        text parameter.
        """
        return pulumi.get(self, "instance_text_params")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetInstanceParamsResult(GetInstanceParamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceParamsResult(
            id=self.id,
            instance_enum_params=self.instance_enum_params,
            instance_id=self.instance_id,
            instance_integer_params=self.instance_integer_params,
            instance_multi_params=self.instance_multi_params,
            instance_text_params=self.instance_text_params,
            result_output_file=self.result_output_file)


def get_instance_params(instance_id: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceParamsResult:
    """
    Use this data source to query detailed information of mongodb instance_params

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_params = tencentcloud.Mongodb.get_instance_params(instance_id="cmgo-gwqk8669")
    ```


    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Mongodb/getInstanceParams:getInstanceParams', __args__, opts=opts, typ=GetInstanceParamsResult).value

    return AwaitableGetInstanceParamsResult(
        id=__ret__.id,
        instance_enum_params=__ret__.instance_enum_params,
        instance_id=__ret__.instance_id,
        instance_integer_params=__ret__.instance_integer_params,
        instance_multi_params=__ret__.instance_multi_params,
        instance_text_params=__ret__.instance_text_params,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_instance_params)
def get_instance_params_output(instance_id: Optional[pulumi.Input[str]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceParamsResult]:
    """
    Use this data source to query detailed information of mongodb instance_params

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    instance_params = tencentcloud.Mongodb.get_instance_params(instance_id="cmgo-gwqk8669")
    ```


    :param str instance_id: InstanceId.
    :param str result_output_file: Used to save results.
    """
    ...
