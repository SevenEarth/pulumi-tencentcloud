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
    'GetOrgFinancialByMonthResult',
    'AwaitableGetOrgFinancialByMonthResult',
    'get_org_financial_by_month',
    'get_org_financial_by_month_output',
]

@pulumi.output_type
class GetOrgFinancialByMonthResult:
    """
    A collection of values returned by getOrgFinancialByMonth.
    """
    def __init__(__self__, end_month=None, id=None, items=None, member_uins=None, product_codes=None, result_output_file=None):
        if end_month and not isinstance(end_month, str):
            raise TypeError("Expected argument 'end_month' to be a str")
        pulumi.set(__self__, "end_month", end_month)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)
        if member_uins and not isinstance(member_uins, list):
            raise TypeError("Expected argument 'member_uins' to be a list")
        pulumi.set(__self__, "member_uins", member_uins)
        if product_codes and not isinstance(product_codes, list):
            raise TypeError("Expected argument 'product_codes' to be a list")
        pulumi.set(__self__, "product_codes", product_codes)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="endMonth")
    def end_month(self) -> Optional[str]:
        return pulumi.get(self, "end_month")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.GetOrgFinancialByMonthItemResult']:
        """
        Organization financial info by month.
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter(name="memberUins")
    def member_uins(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "member_uins")

    @property
    @pulumi.getter(name="productCodes")
    def product_codes(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "product_codes")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetOrgFinancialByMonthResult(GetOrgFinancialByMonthResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgFinancialByMonthResult(
            end_month=self.end_month,
            id=self.id,
            items=self.items,
            member_uins=self.member_uins,
            product_codes=self.product_codes,
            result_output_file=self.result_output_file)


def get_org_financial_by_month(end_month: Optional[str] = None,
                               member_uins: Optional[Sequence[int]] = None,
                               product_codes: Optional[Sequence[str]] = None,
                               result_output_file: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgFinancialByMonthResult:
    """
    Use this data source to query detailed information of organization org_financial_by_month

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    org_financial_by_month = tencentcloud.Organization.get_org_financial_by_month(end_month="2023-05",
        member_uins=[100026517717])
    ```
    <!--End PulumiCodeChooser -->


    :param str end_month: Query for the end month. Format:yyyy-mm, for example:2021-01.
    :param Sequence[int] member_uins: Member uin list. Up to 100.
    :param Sequence[str] product_codes: Product code list. Up to 100.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['endMonth'] = end_month
    __args__['memberUins'] = member_uins
    __args__['productCodes'] = product_codes
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Organization/getOrgFinancialByMonth:getOrgFinancialByMonth', __args__, opts=opts, typ=GetOrgFinancialByMonthResult).value

    return AwaitableGetOrgFinancialByMonthResult(
        end_month=pulumi.get(__ret__, 'end_month'),
        id=pulumi.get(__ret__, 'id'),
        items=pulumi.get(__ret__, 'items'),
        member_uins=pulumi.get(__ret__, 'member_uins'),
        product_codes=pulumi.get(__ret__, 'product_codes'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_org_financial_by_month)
def get_org_financial_by_month_output(end_month: Optional[pulumi.Input[Optional[str]]] = None,
                                      member_uins: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                      product_codes: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgFinancialByMonthResult]:
    """
    Use this data source to query detailed information of organization org_financial_by_month

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    org_financial_by_month = tencentcloud.Organization.get_org_financial_by_month(end_month="2023-05",
        member_uins=[100026517717])
    ```
    <!--End PulumiCodeChooser -->


    :param str end_month: Query for the end month. Format:yyyy-mm, for example:2021-01.
    :param Sequence[int] member_uins: Member uin list. Up to 100.
    :param Sequence[str] product_codes: Product code list. Up to 100.
    :param str result_output_file: Used to save results.
    """
    ...
