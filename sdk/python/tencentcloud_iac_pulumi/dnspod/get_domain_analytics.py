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
    'GetDomainAnalyticsResult',
    'AwaitableGetDomainAnalyticsResult',
    'get_domain_analytics',
    'get_domain_analytics_output',
]

@pulumi.output_type
class GetDomainAnalyticsResult:
    """
    A collection of values returned by getDomainAnalytics.
    """
    def __init__(__self__, alias_datas=None, datas=None, dns_format=None, domain=None, domain_id=None, end_date=None, id=None, infos=None, result_output_file=None, start_date=None):
        if alias_datas and not isinstance(alias_datas, list):
            raise TypeError("Expected argument 'alias_datas' to be a list")
        pulumi.set(__self__, "alias_datas", alias_datas)
        if datas and not isinstance(datas, list):
            raise TypeError("Expected argument 'datas' to be a list")
        pulumi.set(__self__, "datas", datas)
        if dns_format and not isinstance(dns_format, str):
            raise TypeError("Expected argument 'dns_format' to be a str")
        pulumi.set(__self__, "dns_format", dns_format)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domain_id and not isinstance(domain_id, int):
            raise TypeError("Expected argument 'domain_id' to be a int")
        pulumi.set(__self__, "domain_id", domain_id)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if infos and not isinstance(infos, list):
            raise TypeError("Expected argument 'infos' to be a list")
        pulumi.set(__self__, "infos", infos)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)

    @property
    @pulumi.getter(name="aliasDatas")
    def alias_datas(self) -> Sequence['outputs.GetDomainAnalyticsAliasDataResult']:
        """
        Domain alias resolution volume statistics information.
        """
        return pulumi.get(self, "alias_datas")

    @property
    @pulumi.getter
    def datas(self) -> Sequence['outputs.GetDomainAnalyticsDataResult']:
        """
        Subtotal of resolution volume for the current statistical dimension.
        """
        return pulumi.get(self, "datas")

    @property
    @pulumi.getter(name="dnsFormat")
    def dns_format(self) -> Optional[str]:
        """
        DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
        """
        return pulumi.get(self, "dns_format")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The domain name currently being queried.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[int]:
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        """
        End time of the current statistical period.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def infos(self) -> Sequence['outputs.GetDomainAnalyticsInfoResult']:
        """
        Domain resolution volume statistics query information.
        """
        return pulumi.get(self, "infos")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        """
        Start time of the current statistical period.
        """
        return pulumi.get(self, "start_date")


class AwaitableGetDomainAnalyticsResult(GetDomainAnalyticsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainAnalyticsResult(
            alias_datas=self.alias_datas,
            datas=self.datas,
            dns_format=self.dns_format,
            domain=self.domain,
            domain_id=self.domain_id,
            end_date=self.end_date,
            id=self.id,
            infos=self.infos,
            result_output_file=self.result_output_file,
            start_date=self.start_date)


def get_domain_analytics(dns_format: Optional[str] = None,
                         domain: Optional[str] = None,
                         domain_id: Optional[int] = None,
                         end_date: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         start_date: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainAnalyticsResult:
    """
    Use this data source to query detailed information of dnspod domain_analytics

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_analytics = tencentcloud.Dnspod.get_domain_analytics(dns_format="HOUR",
        domain="dnspod.cn",
        end_date="2023-10-12",
        start_date="2023-10-07")
    ```
    <!--End PulumiCodeChooser -->


    :param str dns_format: DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
    :param str domain: The domain name to query for resolution volume.
    :param int domain_id: Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
    :param str end_date: The end date of the query, format: YYYY-MM-DD.
    :param str result_output_file: Used to save results.
    :param str start_date: The start date of the query, format: YYYY-MM-DD.
    """
    __args__ = dict()
    __args__['dnsFormat'] = dns_format
    __args__['domain'] = domain
    __args__['domainId'] = domain_id
    __args__['endDate'] = end_date
    __args__['resultOutputFile'] = result_output_file
    __args__['startDate'] = start_date
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dnspod/getDomainAnalytics:getDomainAnalytics', __args__, opts=opts, typ=GetDomainAnalyticsResult).value

    return AwaitableGetDomainAnalyticsResult(
        alias_datas=pulumi.get(__ret__, 'alias_datas'),
        datas=pulumi.get(__ret__, 'datas'),
        dns_format=pulumi.get(__ret__, 'dns_format'),
        domain=pulumi.get(__ret__, 'domain'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        end_date=pulumi.get(__ret__, 'end_date'),
        id=pulumi.get(__ret__, 'id'),
        infos=pulumi.get(__ret__, 'infos'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        start_date=pulumi.get(__ret__, 'start_date'))


@_utilities.lift_output_func(get_domain_analytics)
def get_domain_analytics_output(dns_format: Optional[pulumi.Input[Optional[str]]] = None,
                                domain: Optional[pulumi.Input[str]] = None,
                                domain_id: Optional[pulumi.Input[Optional[int]]] = None,
                                end_date: Optional[pulumi.Input[str]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                start_date: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainAnalyticsResult]:
    """
    Use this data source to query detailed information of dnspod domain_analytics

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_analytics = tencentcloud.Dnspod.get_domain_analytics(dns_format="HOUR",
        domain="dnspod.cn",
        end_date="2023-10-12",
        start_date="2023-10-07")
    ```
    <!--End PulumiCodeChooser -->


    :param str dns_format: DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
    :param str domain: The domain name to query for resolution volume.
    :param int domain_id: Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
    :param str end_date: The end date of the query, format: YYYY-MM-DD.
    :param str result_output_file: Used to save results.
    :param str start_date: The start date of the query, format: YYYY-MM-DD.
    """
    ...
