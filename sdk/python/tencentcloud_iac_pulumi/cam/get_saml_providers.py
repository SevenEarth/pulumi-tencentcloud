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
    'GetSamlProvidersResult',
    'AwaitableGetSamlProvidersResult',
    'get_saml_providers',
    'get_saml_providers_output',
]

@pulumi.output_type
class GetSamlProvidersResult:
    """
    A collection of values returned by getSamlProviders.
    """
    def __init__(__self__, description=None, id=None, name=None, provider_lists=None, result_output_file=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provider_lists and not isinstance(provider_lists, list):
            raise TypeError("Expected argument 'provider_lists' to be a list")
        pulumi.set(__self__, "provider_lists", provider_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of CAM SAML provider.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of CAM SAML provider.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="providerLists")
    def provider_lists(self) -> Sequence['outputs.GetSamlProvidersProviderListResult']:
        """
        A list of CAM SAML providers. Each element contains the following attributes:
        """
        return pulumi.get(self, "provider_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetSamlProvidersResult(GetSamlProvidersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSamlProvidersResult(
            description=self.description,
            id=self.id,
            name=self.name,
            provider_lists=self.provider_lists,
            result_output_file=self.result_output_file)


def get_saml_providers(description: Optional[str] = None,
                       name: Optional[str] = None,
                       result_output_file: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSamlProvidersResult:
    """
    Use this data source to query detailed information of CAM SAML providers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_saml_providers(name="cam-test-provider")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the CAM SAML provider.
    :param str name: Name of the CAM SAML provider to be queried.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cam/getSamlProviders:getSamlProviders', __args__, opts=opts, typ=GetSamlProvidersResult).value

    return AwaitableGetSamlProvidersResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provider_lists=pulumi.get(__ret__, 'provider_lists'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_saml_providers)
def get_saml_providers_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                              name: Optional[pulumi.Input[Optional[str]]] = None,
                              result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSamlProvidersResult]:
    """
    Use this data source to query detailed information of CAM SAML providers

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_saml_providers(name="cam-test-provider")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the CAM SAML provider.
    :param str name: Name of the CAM SAML provider to be queried.
    :param str result_output_file: Used to save results.
    """
    ...
