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
    'GetBackupCommandsResult',
    'AwaitableGetBackupCommandsResult',
    'get_backup_commands',
    'get_backup_commands_output',
]

@pulumi.output_type
class GetBackupCommandsResult:
    """
    A collection of values returned by getBackupCommands.
    """
    def __init__(__self__, backup_file_type=None, data_base_name=None, id=None, is_recovery=None, lists=None, local_path=None, result_output_file=None):
        if backup_file_type and not isinstance(backup_file_type, str):
            raise TypeError("Expected argument 'backup_file_type' to be a str")
        pulumi.set(__self__, "backup_file_type", backup_file_type)
        if data_base_name and not isinstance(data_base_name, str):
            raise TypeError("Expected argument 'data_base_name' to be a str")
        pulumi.set(__self__, "data_base_name", data_base_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_recovery and not isinstance(is_recovery, str):
            raise TypeError("Expected argument 'is_recovery' to be a str")
        pulumi.set(__self__, "is_recovery", is_recovery)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if local_path and not isinstance(local_path, str):
            raise TypeError("Expected argument 'local_path' to be a str")
        pulumi.set(__self__, "local_path", local_path)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="backupFileType")
    def backup_file_type(self) -> str:
        return pulumi.get(self, "backup_file_type")

    @property
    @pulumi.getter(name="dataBaseName")
    def data_base_name(self) -> str:
        return pulumi.get(self, "data_base_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isRecovery")
    def is_recovery(self) -> str:
        return pulumi.get(self, "is_recovery")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.GetBackupCommandsListResult']:
        """
        Command list.
        """
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter(name="localPath")
    def local_path(self) -> Optional[str]:
        return pulumi.get(self, "local_path")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetBackupCommandsResult(GetBackupCommandsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupCommandsResult(
            backup_file_type=self.backup_file_type,
            data_base_name=self.data_base_name,
            id=self.id,
            is_recovery=self.is_recovery,
            lists=self.lists,
            local_path=self.local_path,
            result_output_file=self.result_output_file)


def get_backup_commands(backup_file_type: Optional[str] = None,
                        data_base_name: Optional[str] = None,
                        is_recovery: Optional[str] = None,
                        local_path: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupCommandsResult:
    """
    Use this data source to query detailed information of sqlserver datasource_backup_command

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_command = tencentcloud.Sqlserver.get_backup_commands(backup_file_type="FULL",
        data_base_name="keep-publish-instance",
        is_recovery="NO",
        local_path="")
    ```


    :param str backup_file_type: Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
    :param str data_base_name: Database name.
    :param str is_recovery: Whether restoration is required. No: not required. Yes: required.
    :param str local_path: Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['backupFileType'] = backup_file_type
    __args__['dataBaseName'] = data_base_name
    __args__['isRecovery'] = is_recovery
    __args__['localPath'] = local_path
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getBackupCommands:getBackupCommands', __args__, opts=opts, typ=GetBackupCommandsResult).value

    return AwaitableGetBackupCommandsResult(
        backup_file_type=__ret__.backup_file_type,
        data_base_name=__ret__.data_base_name,
        id=__ret__.id,
        is_recovery=__ret__.is_recovery,
        lists=__ret__.lists,
        local_path=__ret__.local_path,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_backup_commands)
def get_backup_commands_output(backup_file_type: Optional[pulumi.Input[str]] = None,
                               data_base_name: Optional[pulumi.Input[str]] = None,
                               is_recovery: Optional[pulumi.Input[str]] = None,
                               local_path: Optional[pulumi.Input[Optional[str]]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupCommandsResult]:
    """
    Use this data source to query detailed information of sqlserver datasource_backup_command

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    backup_command = tencentcloud.Sqlserver.get_backup_commands(backup_file_type="FULL",
        data_base_name="keep-publish-instance",
        is_recovery="NO",
        local_path="")
    ```


    :param str backup_file_type: Backup file type. Full: full backup. FULL_LOG: full backup which needs log increments. FULL_DIFF: full backup which needs differential increments. LOG: log backup. DIFF: differential backup.
    :param str data_base_name: Database name.
    :param str is_recovery: Whether restoration is required. No: not required. Yes: required.
    :param str local_path: Storage path of backup files. If this parameter is left empty, the default storage path will be D:.
    :param str result_output_file: Used to save results.
    """
    ...
