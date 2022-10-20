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
    'GetGroupPolicyAttachmentsResult',
    'AwaitableGetGroupPolicyAttachmentsResult',
    'get_group_policy_attachments',
    'get_group_policy_attachments_output',
]

@pulumi.output_type
class GetGroupPolicyAttachmentsResult:
    """
    A collection of values returned by getGroupPolicyAttachments.
    """
    def __init__(__self__, create_mode=None, group_id=None, group_policy_attachment_lists=None, id=None, policy_id=None, policy_type=None, result_output_file=None):
        if create_mode and not isinstance(create_mode, int):
            raise TypeError("Expected argument 'create_mode' to be a int")
        pulumi.set(__self__, "create_mode", create_mode)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if group_policy_attachment_lists and not isinstance(group_policy_attachment_lists, list):
            raise TypeError("Expected argument 'group_policy_attachment_lists' to be a list")
        pulumi.set(__self__, "group_policy_attachment_lists", group_policy_attachment_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[int]:
        """
        Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        """
        ID of CAM group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupPolicyAttachmentLists")
    def group_policy_attachment_lists(self) -> Sequence['outputs.GetGroupPolicyAttachmentsGroupPolicyAttachmentListResult']:
        """
        A list of CAM group policy attachments. Each element contains the following attributes:
        """
        return pulumi.get(self, "group_policy_attachment_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        Name of CAM group.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        """
        Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetGroupPolicyAttachmentsResult(GetGroupPolicyAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupPolicyAttachmentsResult(
            create_mode=self.create_mode,
            group_id=self.group_id,
            group_policy_attachment_lists=self.group_policy_attachment_lists,
            id=self.id,
            policy_id=self.policy_id,
            policy_type=self.policy_type,
            result_output_file=self.result_output_file)


def get_group_policy_attachments(create_mode: Optional[int] = None,
                                 group_id: Optional[str] = None,
                                 policy_id: Optional[str] = None,
                                 policy_type: Optional[str] = None,
                                 result_output_file: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupPolicyAttachmentsResult:
    """
    Use this data source to query detailed information of CAM group policy attachments

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_group_policy_attachments(group_id=tencentcloud_cam_group["foo"]["id"])
    bar = tencentcloud.Cam.get_group_policy_attachments(group_id=tencentcloud_cam_group["foo"]["id"],
        policy_id=tencentcloud_cam_policy["foo"]["id"])
    ```


    :param int create_mode: Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
    :param str group_id: ID of the attached CAM group to be queried.
    :param str policy_id: ID of CAM policy to be queried.
    :param str policy_type: Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['createMode'] = create_mode
    __args__['groupId'] = group_id
    __args__['policyId'] = policy_id
    __args__['policyType'] = policy_type
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cam/getGroupPolicyAttachments:getGroupPolicyAttachments', __args__, opts=opts, typ=GetGroupPolicyAttachmentsResult).value

    return AwaitableGetGroupPolicyAttachmentsResult(
        create_mode=__ret__.create_mode,
        group_id=__ret__.group_id,
        group_policy_attachment_lists=__ret__.group_policy_attachment_lists,
        id=__ret__.id,
        policy_id=__ret__.policy_id,
        policy_type=__ret__.policy_type,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(get_group_policy_attachments)
def get_group_policy_attachments_output(create_mode: Optional[pulumi.Input[Optional[int]]] = None,
                                        group_id: Optional[pulumi.Input[str]] = None,
                                        policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                        policy_type: Optional[pulumi.Input[Optional[str]]] = None,
                                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupPolicyAttachmentsResult]:
    """
    Use this data source to query detailed information of CAM group policy attachments

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    foo = tencentcloud.Cam.get_group_policy_attachments(group_id=tencentcloud_cam_group["foo"]["id"])
    bar = tencentcloud.Cam.get_group_policy_attachments(group_id=tencentcloud_cam_group["foo"]["id"],
        policy_id=tencentcloud_cam_policy["foo"]["id"])
    ```


    :param int create_mode: Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
    :param str group_id: ID of the attached CAM group to be queried.
    :param str policy_id: ID of CAM policy to be queried.
    :param str policy_type: Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
    :param str result_output_file: Used to save results.
    """
    ...
