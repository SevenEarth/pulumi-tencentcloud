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
    'GetApplicationPublicConfigResult',
    'AwaitableGetApplicationPublicConfigResult',
    'get_application_public_config',
    'get_application_public_config_output',
]

@pulumi.output_type
class GetApplicationPublicConfigResult:
    """
    A collection of values returned by getApplicationPublicConfig.
    """
    def __init__(__self__, config_id=None, config_id_lists=None, config_name=None, config_version=None, id=None, result_output_file=None, results=None):
        if config_id and not isinstance(config_id, str):
            raise TypeError("Expected argument 'config_id' to be a str")
        pulumi.set(__self__, "config_id", config_id)
        if config_id_lists and not isinstance(config_id_lists, list):
            raise TypeError("Expected argument 'config_id_lists' to be a list")
        pulumi.set(__self__, "config_id_lists", config_id_lists)
        if config_name and not isinstance(config_name, str):
            raise TypeError("Expected argument 'config_name' to be a str")
        pulumi.set(__self__, "config_name", config_name)
        if config_version and not isinstance(config_version, str):
            raise TypeError("Expected argument 'config_version' to be a str")
        pulumi.set(__self__, "config_version", config_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[str]:
        """
        Config ID. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="configIdLists")
    def config_id_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "config_id_lists")

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[str]:
        """
        Config name. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[str]:
        """
        Config version. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetApplicationPublicConfigResultResult']:
        """
        Paginated global configuration  list. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "results")


class AwaitableGetApplicationPublicConfigResult(GetApplicationPublicConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationPublicConfigResult(
            config_id=self.config_id,
            config_id_lists=self.config_id_lists,
            config_name=self.config_name,
            config_version=self.config_version,
            id=self.id,
            result_output_file=self.result_output_file,
            results=self.results)


def get_application_public_config(config_id: Optional[str] = None,
                                  config_id_lists: Optional[Sequence[str]] = None,
                                  config_name: Optional[str] = None,
                                  config_version: Optional[str] = None,
                                  result_output_file: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationPublicConfigResult:
    """
    Use this data source to query detailed information of tsf application_public_config

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    application_public_config = tencentcloud.Tsf.get_application_public_config(config_id="dcfg-p-evjrbgly",
        config_name="dsadsa",
        config_version="123")
    ```


    :param str config_id: Config ID. Query all items if not passed, high priority.
    :param Sequence[str] config_id_lists: Config ID list. Query all items if not passed, low priority.
    :param str config_name: Config name. Exact query. Query all items if not passed.
    :param str config_version: Config version. Exact query. Query all items if not passed.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['configIdLists'] = config_id_lists
    __args__['configName'] = config_name
    __args__['configVersion'] = config_version
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tsf/getApplicationPublicConfig:getApplicationPublicConfig', __args__, opts=opts, typ=GetApplicationPublicConfigResult).value

    return AwaitableGetApplicationPublicConfigResult(
        config_id=__ret__.config_id,
        config_id_lists=__ret__.config_id_lists,
        config_name=__ret__.config_name,
        config_version=__ret__.config_version,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        results=__ret__.results)


@_utilities.lift_output_func(get_application_public_config)
def get_application_public_config_output(config_id: Optional[pulumi.Input[Optional[str]]] = None,
                                         config_id_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                         config_name: Optional[pulumi.Input[Optional[str]]] = None,
                                         config_version: Optional[pulumi.Input[Optional[str]]] = None,
                                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationPublicConfigResult]:
    """
    Use this data source to query detailed information of tsf application_public_config

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    application_public_config = tencentcloud.Tsf.get_application_public_config(config_id="dcfg-p-evjrbgly",
        config_name="dsadsa",
        config_version="123")
    ```


    :param str config_id: Config ID. Query all items if not passed, high priority.
    :param Sequence[str] config_id_lists: Config ID list. Query all items if not passed, low priority.
    :param str config_name: Config name. Exact query. Query all items if not passed.
    :param str config_version: Config version. Exact query. Query all items if not passed.
    :param str result_output_file: Used to save results.
    """
    ...
