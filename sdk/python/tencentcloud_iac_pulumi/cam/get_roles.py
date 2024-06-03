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
    'GetRolesResult',
    'AwaitableGetRolesResult',
    'get_roles',
    'get_roles_output',
]

@pulumi.output_type
class GetRolesResult:
    """
    A collection of values returned by getRoles.
    """
    def __init__(__self__, description=None, id=None, name=None, result_output_file=None, role_id=None, role_lists=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if role_id and not isinstance(role_id, str):
            raise TypeError("Expected argument 'role_id' to be a str")
        pulumi.set(__self__, "role_id", role_id)
        if role_lists and not isinstance(role_lists, list):
            raise TypeError("Expected argument 'role_lists' to be a list")
        pulumi.set(__self__, "role_lists", role_lists)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of CAM role.
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
        Name of CAM role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[str]:
        """
        Id of CAM role.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="roleLists")
    def role_lists(self) -> Sequence['outputs.GetRolesRoleListResult']:
        """
        A list of CAM roles. Each element contains the following attributes:
        """
        return pulumi.get(self, "role_lists")


class AwaitableGetRolesResult(GetRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRolesResult(
            description=self.description,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            role_id=self.role_id,
            role_lists=self.role_lists)


def get_roles(description: Optional[str] = None,
              name: Optional[str] = None,
              result_output_file: Optional[str] = None,
              role_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRolesResult:
    """
    Use this data source to query detailed information of CAM roles

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_roles(role_id=tencentcloud_cam_role["foo"]["id"])
    bar = tencentcloud.Cam.get_roles(name="cam-role-test")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the CAM role to be queried.
    :param str name: Name of the CAM policy to be queried.
    :param str result_output_file: Used to save results.
    :param str role_id: ID of the CAM role to be queried.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['roleId'] = role_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cam/getRoles:getRoles', __args__, opts=opts, typ=GetRolesResult).value

    return AwaitableGetRolesResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        role_id=pulumi.get(__ret__, 'role_id'),
        role_lists=pulumi.get(__ret__, 'role_lists'))


@_utilities.lift_output_func(get_roles)
def get_roles_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     role_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRolesResult]:
    """
    Use this data source to query detailed information of CAM roles

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_roles(role_id=tencentcloud_cam_role["foo"]["id"])
    bar = tencentcloud.Cam.get_roles(name="cam-role-test")
    ```
    <!--End PulumiCodeChooser -->


    :param str description: The description of the CAM role to be queried.
    :param str name: Name of the CAM policy to be queried.
    :param str result_output_file: Used to save results.
    :param str role_id: ID of the CAM role to be queried.
    """
    ...
