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
    'GetFileSystemClientsResult',
    'AwaitableGetFileSystemClientsResult',
    'get_file_system_clients',
    'get_file_system_clients_output',
]

@pulumi.output_type
class GetFileSystemClientsResult:
    """
    A collection of values returned by getFileSystemClients.
    """
    def __init__(__self__, client_lists=None, file_system_id=None, id=None, result_output_file=None):
        if client_lists and not isinstance(client_lists, list):
            raise TypeError("Expected argument 'client_lists' to be a list")
        pulumi.set(__self__, "client_lists", client_lists)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clientLists")
    def client_lists(self) -> Sequence['outputs.GetFileSystemClientsClientListResult']:
        """
        Client list.
        """
        return pulumi.get(self, "client_lists")

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        return pulumi.get(self, "file_system_id")

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


class AwaitableGetFileSystemClientsResult(GetFileSystemClientsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileSystemClientsResult(
            client_lists=self.client_lists,
            file_system_id=self.file_system_id,
            id=self.id,
            result_output_file=self.result_output_file)


def get_file_system_clients(file_system_id: Optional[str] = None,
                            result_output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileSystemClientsResult:
    """
    Use this data source to query detailed information of cfs file_system_clients

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    file_system_clients = tencentcloud.Cfs.get_file_system_clients(file_system_id="cfs-iobiaxtj")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_system_id: File system ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['fileSystemId'] = file_system_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cfs/getFileSystemClients:getFileSystemClients', __args__, opts=opts, typ=GetFileSystemClientsResult).value

    return AwaitableGetFileSystemClientsResult(
        client_lists=pulumi.get(__ret__, 'client_lists'),
        file_system_id=pulumi.get(__ret__, 'file_system_id'),
        id=pulumi.get(__ret__, 'id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_file_system_clients)
def get_file_system_clients_output(file_system_id: Optional[pulumi.Input[str]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileSystemClientsResult]:
    """
    Use this data source to query detailed information of cfs file_system_clients

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    file_system_clients = tencentcloud.Cfs.get_file_system_clients(file_system_id="cfs-iobiaxtj")
    ```
    <!--End PulumiCodeChooser -->


    :param str file_system_id: File system ID.
    :param str result_output_file: Used to save results.
    """
    ...
