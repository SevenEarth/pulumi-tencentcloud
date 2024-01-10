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
    def __init__(__self__, domain=None, domains=None, id=None, instance_id=None, result_output_file=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetDomainsDomainResult']:
        """
        Domain info list.
        """
        return pulumi.get(self, "domains")

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
        """
        Instance unique ID.
        """
        return pulumi.get(self, "instance_id")

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
            domain=self.domain,
            domains=self.domains,
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file)


def get_domains(domain: Optional[str] = None,
                instance_id: Optional[str] = None,
                result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    Use this data source to query detailed information of waf domains

    ## Example Usage
    ### Find all domains

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_domains()
    ```
    ### Find domains by filter

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_domains(domain="tf.example.com",
        instance_id="waf_2kxtlbky01b3wceb")
    ```


    :param str domain: Domain name.
    :param str instance_id: Unique ID of Instance.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Waf/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        domain=__ret__.domain,
        domains=__ret__.domains,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_domains)
def get_domains_output(domain: Optional[pulumi.Input[Optional[str]]] = None,
                       instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainsResult]:
    """
    Use this data source to query detailed information of waf domains

    ## Example Usage
    ### Find all domains

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_domains()
    ```
    ### Find domains by filter

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Waf.get_domains(domain="tf.example.com",
        instance_id="waf_2kxtlbky01b3wceb")
    ```


    :param str domain: Domain name.
    :param str instance_id: Unique ID of Instance.
    :param str result_output_file: Used to save results.
    """
    ...
