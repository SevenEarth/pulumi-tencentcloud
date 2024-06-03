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
    'GetRegionsResult',
    'AwaitableGetRegionsResult',
    'get_regions',
    'get_regions_output',
]

@pulumi.output_type
class GetRegionsResult:
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, id=None, region_sets=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region_sets and not isinstance(region_sets, list):
            raise TypeError("Expected argument 'region_sets' to be a list")
        pulumi.set(__self__, "region_sets", region_sets)
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
    @pulumi.getter(name="regionSets")
    def region_sets(self) -> Sequence['outputs.GetRegionsRegionSetResult']:
        """
        Region information set.
        """
        return pulumi.get(self, "region_sets")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetRegionsResult(GetRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionsResult(
            id=self.id,
            region_sets=self.region_sets,
            result_output_file=self.result_output_file)


def get_regions(result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionsResult:
    """
    Use this data source to query detailed information of postgresql regions

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    regions = tencentcloud.Postgresql.get_regions()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Postgresql/getRegions:getRegions', __args__, opts=opts, typ=GetRegionsResult).value

    return AwaitableGetRegionsResult(
        id=pulumi.get(__ret__, 'id'),
        region_sets=pulumi.get(__ret__, 'region_sets'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_regions)
def get_regions_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionsResult]:
    """
    Use this data source to query detailed information of postgresql regions

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    regions = tencentcloud.Postgresql.get_regions()
    ```
    <!--End PulumiCodeChooser -->


    :param str result_output_file: Used to save results.
    """
    ...
