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
    'GetDomainsResult',
    'AwaitableGetDomainsResult',
    'get_domains',
    'get_domains_output',
]

@pulumi.output_type
class GetDomainsResult:
    """
    A collection of values returned by getDomains.
    """
    def __init__(__self__, domain_lists=None, domain_prefix=None, domain_status=None, domain_type=None, id=None, is_delay_live=None, play_type=None, result_output_file=None):
        if domain_lists and not isinstance(domain_lists, list):
            raise TypeError("Expected argument 'domain_lists' to be a list")
        pulumi.set(__self__, "domain_lists", domain_lists)
        if domain_prefix and not isinstance(domain_prefix, str):
            raise TypeError("Expected argument 'domain_prefix' to be a str")
        pulumi.set(__self__, "domain_prefix", domain_prefix)
        if domain_status and not isinstance(domain_status, int):
            raise TypeError("Expected argument 'domain_status' to be a int")
        pulumi.set(__self__, "domain_status", domain_status)
        if domain_type and not isinstance(domain_type, int):
            raise TypeError("Expected argument 'domain_type' to be a int")
        pulumi.set(__self__, "domain_type", domain_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_delay_live and not isinstance(is_delay_live, int):
            raise TypeError("Expected argument 'is_delay_live' to be a int")
        pulumi.set(__self__, "is_delay_live", is_delay_live)
        if play_type and not isinstance(play_type, int):
            raise TypeError("Expected argument 'play_type' to be a int")
        pulumi.set(__self__, "play_type", play_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="domainLists")
    def domain_lists(self) -> Sequence['outputs.GetDomainsDomainListResult']:
        """
        A list of domain name details.
        """
        return pulumi.get(self, "domain_lists")

    @property
    @pulumi.getter(name="domainPrefix")
    def domain_prefix(self) -> Optional[str]:
        return pulumi.get(self, "domain_prefix")

    @property
    @pulumi.getter(name="domainStatus")
    def domain_status(self) -> Optional[int]:
        return pulumi.get(self, "domain_status")

    @property
    @pulumi.getter(name="domainType")
    def domain_type(self) -> Optional[int]:
        return pulumi.get(self, "domain_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDelayLive")
    def is_delay_live(self) -> Optional[int]:
        """
        Whether to slow live broadcast: 0: normal live broadcast. 1: Slow live broadcast.
        """
        return pulumi.get(self, "is_delay_live")

    @property
    @pulumi.getter(name="playType")
    def play_type(self) -> Optional[int]:
        """
        Playing area, this parameter is meaningful only when Type=1. 1: Domestic. 2: Global. 3: Overseas.
        """
        return pulumi.get(self, "play_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDomainsResult(GetDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsResult(
            domain_lists=self.domain_lists,
            domain_prefix=self.domain_prefix,
            domain_status=self.domain_status,
            domain_type=self.domain_type,
            id=self.id,
            is_delay_live=self.is_delay_live,
            play_type=self.play_type,
            result_output_file=self.result_output_file)


def get_domains(domain_prefix: Optional[str] = None,
                domain_status: Optional[int] = None,
                domain_type: Optional[int] = None,
                is_delay_live: Optional[int] = None,
                play_type: Optional[int] = None,
                result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    Use this data source to query detailed information of css domains

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domains = tencentcloud.Css.get_domains(domain_type=0,
        is_delay_live=0,
        play_type=1)
    ```


    :param str domain_prefix: domain name prefix.
    :param int domain_status: domain name status filter. 0-disable, 1-enable.
    :param int domain_type: Domain name type filtering. 0-push, 1-play.
    :param int is_delay_live: 0 normal live broadcast 1 slow live broadcast default 0.
    :param int play_type: Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['domainPrefix'] = domain_prefix
    __args__['domainStatus'] = domain_status
    __args__['domainType'] = domain_type
    __args__['isDelayLive'] = is_delay_live
    __args__['playType'] = play_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Css/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        domain_lists=__ret__.domain_lists,
        domain_prefix=__ret__.domain_prefix,
        domain_status=__ret__.domain_status,
        domain_type=__ret__.domain_type,
        id=__ret__.id,
        is_delay_live=__ret__.is_delay_live,
        play_type=__ret__.play_type,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_domains)
def get_domains_output(domain_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                       domain_status: Optional[pulumi.Input[Optional[int]]] = None,
                       domain_type: Optional[pulumi.Input[Optional[int]]] = None,
                       is_delay_live: Optional[pulumi.Input[Optional[int]]] = None,
                       play_type: Optional[pulumi.Input[Optional[int]]] = None,
                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainsResult]:
    """
    Use this data source to query detailed information of css domains

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domains = tencentcloud.Css.get_domains(domain_type=0,
        is_delay_live=0,
        play_type=1)
    ```


    :param str domain_prefix: domain name prefix.
    :param int domain_status: domain name status filter. 0-disable, 1-enable.
    :param int domain_type: Domain name type filtering. 0-push, 1-play.
    :param int is_delay_live: 0 normal live broadcast 1 slow live broadcast default 0.
    :param int play_type: Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
    :param str result_output_file: Used to save results.
    """
    ...