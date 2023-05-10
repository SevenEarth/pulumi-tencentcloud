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
    'GetHealthScoresResult',
    'AwaitableGetHealthScoresResult',
    'get_health_scores',
    'get_health_scores_output',
]

@pulumi.output_type
class GetHealthScoresResult:
    """
    A collection of values returned by getHealthScores.
    """
    def __init__(__self__, datas=None, id=None, instance_id=None, product=None, result_output_file=None, time=None):
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if product and not isinstance(product, str):
            raise TypeError("Expected argument 'product' to be a str")
        pulumi.set(__self__, "product", product)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if time and not isinstance(time, str):
            raise TypeError("Expected argument 'time' to be a str")
        pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetHealthScoresDataResult']:
        """
        Health score and abnormal deduction items.
        """
        return pulumi.get(self, "datas")

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
    @pulumi.getter
    def product(self) -> str:
        return pulumi.get(self, "product")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def time(self) -> str:
        return pulumi.get(self, "time")


class AwaitableGetHealthScoresResult(GetHealthScoresResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHealthScoresResult(
            datas=self.datas,
            id=self.id,
            instance_id=self.instance_id,
            product=self.product,
            result_output_file=self.result_output_file,
            time=self.time)


def get_health_scores(instance_id: Optional[str] = None,
                      product: Optional[str] = None,
                      result_output_file: Optional[str] = None,
                      time: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHealthScoresResult:
    """
    Use this data source to query detailed information of dbbrain health_scores

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    health_scores = tencentcloud.Dbbrain.get_health_scores(instance_id="",
        product="",
        time="")
    ```


    :param str instance_id: The ID of the instance whose health score needs to be obtained.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database TDSQL-C for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str time: The time to obtain the health score, the time format is as follows: 2019-09-10 12:13:14.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['product'] = product
    __args__['resultOutputFile'] = result_output_file
    __args__['time'] = time
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dbbrain/getHealthScores:getHealthScores', __args__, opts=opts, typ=GetHealthScoresResult).value

    return AwaitableGetHealthScoresResult(
        datas=__ret__.datas,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        product=__ret__.product,
        result_output_file=__ret__.result_output_file,
        time=__ret__.time)


@_utilities.lift_output_func(get_health_scores)
def get_health_scores_output(instance_id: Optional[pulumi.Input[str]] = None,
                             product: Optional[pulumi.Input[str]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             time: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHealthScoresResult]:
    """
    Use this data source to query detailed information of dbbrain health_scores

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    health_scores = tencentcloud.Dbbrain.get_health_scores(instance_id="",
        product="",
        time="")
    ```


    :param str instance_id: The ID of the instance whose health score needs to be obtained.
    :param str product: Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database TDSQL-C for MySQL, the default is mysql.
    :param str result_output_file: Used to save results.
    :param str time: The time to obtain the health score, the time format is as follows: 2019-09-10 12:13:14.
    """
    ...
