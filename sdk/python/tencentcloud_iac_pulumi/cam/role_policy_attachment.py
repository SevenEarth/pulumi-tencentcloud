# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RolePolicyAttachmentArgs', 'RolePolicyAttachment']

@pulumi.input_type
class RolePolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 role_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a RolePolicyAttachment resource.
        :param pulumi.Input[str] policy_id: ID of the policy.
        :param pulumi.Input[str] role_id: ID of the attached CAM role.
        """
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        ID of the policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        ID of the attached CAM role.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)


@pulumi.input_type
class _RolePolicyAttachmentState:
    def __init__(__self__, *,
                 create_mode: Optional[pulumi.Input[int]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RolePolicyAttachment resources.
        :param pulumi.Input[int] create_mode: Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        :param pulumi.Input[str] create_time: The create time of the CAM role policy attachment.
        :param pulumi.Input[str] policy_id: ID of the policy.
        :param pulumi.Input[str] policy_name: The name of the policy.
        :param pulumi.Input[str] policy_type: Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        :param pulumi.Input[str] role_id: ID of the attached CAM role.
        """
        if create_mode is not None:
            pulumi.set(__self__, "create_mode", create_mode)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[pulumi.Input[int]]:
        """
        Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        """
        return pulumi.get(self, "create_mode")

    @create_mode.setter
    def create_mode(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "create_mode", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the CAM role policy attachment.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the policy.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the attached CAM role.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class RolePolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a CAM role policy attachment.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        cam_policy_basic = config.get("camPolicyBasic")
        if cam_policy_basic is None:
            cam_policy_basic = "keep-cam-policy"
        cam_role_basic = config.get("camRoleBasic")
        if cam_role_basic is None:
            cam_role_basic = "keep-cam-role"
        policy = tencentcloud.Cam.get_policies(name=cam_policy_basic)
        roles = tencentcloud.Cam.get_roles(name=cam_role_basic)
        role_policy_attachment_basic = tencentcloud.cam.RolePolicyAttachment("rolePolicyAttachmentBasic",
            role_id=roles.role_lists[0].role_id,
            policy_id=policy.policy_lists[0].policy_id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CAM role policy attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment foo 4611686018427922725#26800353
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_id: ID of the policy.
        :param pulumi.Input[str] role_id: ID of the attached CAM role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RolePolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CAM role policy attachment.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        cam_policy_basic = config.get("camPolicyBasic")
        if cam_policy_basic is None:
            cam_policy_basic = "keep-cam-policy"
        cam_role_basic = config.get("camRoleBasic")
        if cam_role_basic is None:
            cam_role_basic = "keep-cam-role"
        policy = tencentcloud.Cam.get_policies(name=cam_policy_basic)
        roles = tencentcloud.Cam.get_roles(name=cam_role_basic)
        role_policy_attachment_basic = tencentcloud.cam.RolePolicyAttachment("rolePolicyAttachmentBasic",
            role_id=roles.role_lists[0].role_id,
            policy_id=policy.policy_lists[0].policy_id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CAM role policy attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment foo 4611686018427922725#26800353
        ```

        :param str resource_name: The name of the resource.
        :param RolePolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RolePolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RolePolicyAttachmentArgs.__new__(RolePolicyAttachmentArgs)

            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["create_mode"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["policy_name"] = None
            __props__.__dict__["policy_type"] = None
        super(RolePolicyAttachment, __self__).__init__(
            'tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_mode: Optional[pulumi.Input[int]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'RolePolicyAttachment':
        """
        Get an existing RolePolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] create_mode: Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        :param pulumi.Input[str] create_time: The create time of the CAM role policy attachment.
        :param pulumi.Input[str] policy_id: ID of the policy.
        :param pulumi.Input[str] policy_name: The name of the policy.
        :param pulumi.Input[str] policy_type: Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        :param pulumi.Input[str] role_id: ID of the attached CAM role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RolePolicyAttachmentState.__new__(_RolePolicyAttachmentState)

        __props__.__dict__["create_mode"] = create_mode
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["role_id"] = role_id
        return RolePolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> pulumi.Output[int]:
        """
        Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The create time of the CAM role policy attachment.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        ID of the policy.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        The name of the policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        ID of the attached CAM role.
        """
        return pulumi.get(self, "role_id")

