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
    'GetAclsResult',
    'AwaitableGetAclsResult',
    'get_acls',
    'get_acls_output',
]

@pulumi.output_type
class GetAclsResult:
    """
    A collection of values returned by getAcls.
    """
    def __init__(__self__, acl_lists=None, id=None, name=None, result_output_file=None, vpc_id=None):
        if acl_lists and not isinstance(acl_lists, list):
            raise TypeError("Expected argument 'acl_lists' to be a list")
        pulumi.set(__self__, "acl_lists", acl_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="aclLists")
    def acl_lists(self) -> Sequence['outputs.GetAclsAclListResult']:
        """
        The information list of the VPC. Each element contains the following attributes:
        """
        return pulumi.get(self, "acl_lists")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        ID of the network ACL instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the network ACL.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        ID of the VPC instance.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetAclsResult(GetAclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAclsResult(
            acl_lists=self.acl_lists,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            vpc_id=self.vpc_id)


def get_acls(id: Optional[str] = None,
             name: Optional[str] = None,
             result_output_file: Optional[str] = None,
             vpc_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAclsResult:
    """
    Use this data source to query VPC Network ACL information.


    :param str id: ID of the network ACL instance.
    :param str name: Name of the network ACL.
    :param str result_output_file: Used to save results.
    :param str vpc_id: ID of the VPC instance.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Vpc/getAcls:getAcls', __args__, opts=opts, typ=GetAclsResult).value

    return AwaitableGetAclsResult(
        acl_lists=__ret__.acl_lists,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_acls)
def get_acls_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAclsResult]:
    """
    Use this data source to query VPC Network ACL information.


    :param str id: ID of the network ACL instance.
    :param str name: Name of the network ACL.
    :param str result_output_file: Used to save results.
    :param str vpc_id: ID of the VPC instance.
    """
    ...
