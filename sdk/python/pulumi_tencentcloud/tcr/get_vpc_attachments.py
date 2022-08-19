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
    'GetVpcAttachmentsResult',
    'AwaitableGetVpcAttachmentsResult',
    'get_vpc_attachments',
    'get_vpc_attachments_output',
]

@pulumi.output_type
class GetVpcAttachmentsResult:
    """
    A collection of values returned by getVpcAttachments.
    """
    def __init__(__self__, id=None, instance_id=None, result_output_file=None, subnet_id=None, vpc_attachment_lists=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_attachment_lists and not isinstance(vpc_attachment_lists, list):
            raise TypeError("Expected argument 'vpc_attachment_lists' to be a list")
        pulumi.set(__self__, "vpc_attachment_lists", vpc_attachment_lists)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        ID of subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcAttachmentLists")
    def vpc_attachment_lists(self) -> Sequence['outputs.GetVpcAttachmentsVpcAttachmentListResult']:
        """
        Information list of the dedicated TCR namespaces.
        """
        return pulumi.get(self, "vpc_attachment_lists")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        ID of VPC.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetVpcAttachmentsResult(GetVpcAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcAttachmentsResult(
            id=self.id,
            instance_id=self.instance_id,
            result_output_file=self.result_output_file,
            subnet_id=self.subnet_id,
            vpc_attachment_lists=self.vpc_attachment_lists,
            vpc_id=self.vpc_id)


def get_vpc_attachments(instance_id: Optional[str] = None,
                        result_output_file: Optional[str] = None,
                        subnet_id: Optional[str] = None,
                        vpc_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcAttachmentsResult:
    """
    Use this data source to query detailed information of TCR VPC attachment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    id = tencentcloud.Tcr.get_vpc_attachments(instance_id="cls-satg5125")
    ```


    :param str instance_id: ID of the instance to query.
    :param str result_output_file: Used to save results.
    :param str subnet_id: ID of subnet to query.
    :param str vpc_id: ID of VPC to query.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['resultOutputFile'] = result_output_file
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tcr/getVpcAttachments:getVpcAttachments', __args__, opts=opts, typ=GetVpcAttachmentsResult).value

    return AwaitableGetVpcAttachmentsResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        result_output_file=__ret__.result_output_file,
        subnet_id=__ret__.subnet_id,
        vpc_attachment_lists=__ret__.vpc_attachment_lists,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(get_vpc_attachments)
def get_vpc_attachments_output(instance_id: Optional[pulumi.Input[str]] = None,
                               result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                               vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcAttachmentsResult]:
    """
    Use this data source to query detailed information of TCR VPC attachment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    id = tencentcloud.Tcr.get_vpc_attachments(instance_id="cls-satg5125")
    ```


    :param str instance_id: ID of the instance to query.
    :param str result_output_file: Used to save results.
    :param str subnet_id: ID of subnet to query.
    :param str vpc_id: ID of VPC to query.
    """
    ...
