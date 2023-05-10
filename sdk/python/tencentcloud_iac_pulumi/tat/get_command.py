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
    'GetCommandResult',
    'AwaitableGetCommandResult',
    'get_command',
    'get_command_output',
]

@pulumi.output_type
class GetCommandResult:
    """
    A collection of values returned by getCommand.
    """
    def __init__(__self__, command_id=None, command_name=None, command_sets=None, command_type=None, created_by=None, id=None, result_output_file=None):
        if command_id and not isinstance(command_id, str):
            raise TypeError("Expected argument 'command_id' to be a str")
        pulumi.set(__self__, "command_id", command_id)
        if command_name and not isinstance(command_name, str):
            raise TypeError("Expected argument 'command_name' to be a str")
        pulumi.set(__self__, "command_name", command_name)
        if command_sets and not isinstance(command_sets, list):
            raise TypeError("Expected argument 'command_sets' to be a list")
        pulumi.set(__self__, "command_sets", command_sets)
        if command_type and not isinstance(command_type, str):
            raise TypeError("Expected argument 'command_type' to be a str")
        pulumi.set(__self__, "command_type", command_type)
        if created_by and not isinstance(created_by, str):
            raise TypeError("Expected argument 'created_by' to be a str")
        pulumi.set(__self__, "created_by", created_by)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> Optional[str]:
        """
        Command ID.
        """
        return pulumi.get(self, "command_id")

    @property
    @pulumi.getter(name="commandName")
    def command_name(self) -> Optional[str]:
        """
        Command name.
        """
        return pulumi.get(self, "command_name")

    @property
    @pulumi.getter(name="commandSets")
    def command_sets(self) -> Sequence['outputs.GetCommandCommandSetResult']:
        """
        List of command details.
        """
        return pulumi.get(self, "command_sets")

    @property
    @pulumi.getter(name="commandType")
    def command_type(self) -> Optional[str]:
        """
        Command type.
        """
        return pulumi.get(self, "command_type")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        """
        return pulumi.get(self, "created_by")

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


class AwaitableGetCommandResult(GetCommandResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCommandResult(
            command_id=self.command_id,
            command_name=self.command_name,
            command_sets=self.command_sets,
            command_type=self.command_type,
            created_by=self.created_by,
            id=self.id,
            result_output_file=self.result_output_file)


def get_command(command_id: Optional[str] = None,
                command_name: Optional[str] = None,
                command_type: Optional[str] = None,
                created_by: Optional[str] = None,
                result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCommandResult:
    """
    Use this data source to query detailed information of tat command

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    command = tencentcloud.Tat.get_command(command_type="SHELL",
        created_by="TAT")
    ```


    :param str command_id: Command ID.
    :param str command_name: Command name.
    :param str command_type: Command type, Value is `SHELL` or `POWERSHELL`.
    :param str created_by: Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['commandId'] = command_id
    __args__['commandName'] = command_name
    __args__['commandType'] = command_type
    __args__['createdBy'] = created_by
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tat/getCommand:getCommand', __args__, opts=opts, typ=GetCommandResult).value

    return AwaitableGetCommandResult(
        command_id=__ret__.command_id,
        command_name=__ret__.command_name,
        command_sets=__ret__.command_sets,
        command_type=__ret__.command_type,
        created_by=__ret__.created_by,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_command)
def get_command_output(command_id: Optional[pulumi.Input[Optional[str]]] = None,
                       command_name: Optional[pulumi.Input[Optional[str]]] = None,
                       command_type: Optional[pulumi.Input[Optional[str]]] = None,
                       created_by: Optional[pulumi.Input[Optional[str]]] = None,
                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCommandResult]:
    """
    Use this data source to query detailed information of tat command

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    command = tencentcloud.Tat.get_command(command_type="SHELL",
        created_by="TAT")
    ```


    :param str command_id: Command ID.
    :param str command_name: Command name.
    :param str command_type: Command type, Value is `SHELL` or `POWERSHELL`.
    :param str created_by: Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
    :param str result_output_file: Used to save results.
    """
    ...