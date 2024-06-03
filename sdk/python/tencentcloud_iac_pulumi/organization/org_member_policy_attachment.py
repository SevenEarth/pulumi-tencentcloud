# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrgMemberPolicyAttachmentArgs', 'OrgMemberPolicyAttachment']

@pulumi.input_type
class OrgMemberPolicyAttachmentArgs:
    def __init__(__self__, *,
                 identity_id: pulumi.Input[int],
                 member_uins: pulumi.Input[Sequence[pulumi.Input[int]]],
                 policy_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrgMemberPolicyAttachment resource.
        :param pulumi.Input[int] identity_id: Organization identity ID.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] member_uins: Member Uin list. Up to 10.
        :param pulumi.Input[str] policy_name: Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        :param pulumi.Input[str] description: Notes.The maximum length is 128 characters.
        """
        pulumi.set(__self__, "identity_id", identity_id)
        pulumi.set(__self__, "member_uins", member_uins)
        pulumi.set(__self__, "policy_name", policy_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Input[int]:
        """
        Organization identity ID.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="memberUins")
    def member_uins(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        Member Uin list. Up to 10.
        """
        return pulumi.get(self, "member_uins")

    @member_uins.setter
    def member_uins(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "member_uins", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        """
        Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Notes.The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _OrgMemberPolicyAttachmentState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_id: Optional[pulumi.Input[int]] = None,
                 member_uins: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrgMemberPolicyAttachment resources.
        :param pulumi.Input[str] description: Notes.The maximum length is 128 characters.
        :param pulumi.Input[int] identity_id: Organization identity ID.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] member_uins: Member Uin list. Up to 10.
        :param pulumi.Input[str] policy_name: Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identity_id is not None:
            pulumi.set(__self__, "identity_id", identity_id)
        if member_uins is not None:
            pulumi.set(__self__, "member_uins", member_uins)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Notes.The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> Optional[pulumi.Input[int]]:
        """
        Organization identity ID.
        """
        return pulumi.get(self, "identity_id")

    @identity_id.setter
    def identity_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "identity_id", value)

    @property
    @pulumi.getter(name="memberUins")
    def member_uins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        Member Uin list. Up to 10.
        """
        return pulumi.get(self, "member_uins")

    @member_uins.setter
    def member_uins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "member_uins", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)


class OrgMemberPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_id: Optional[pulumi.Input[int]] = None,
                 member_uins: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a organization org_member_policy_attachment

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        org_member_policy_attachment = tencentcloud.organization.OrgMemberPolicyAttachment("orgMemberPolicyAttachment",
            identity_id=1,
            member_uins=[
                100033905366,
                100033905356,
            ],
            policy_name="example-iac")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        organization org_member_policy_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Organization/orgMemberPolicyAttachment:OrgMemberPolicyAttachment org_member_policy_attachment org_member_policy_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Notes.The maximum length is 128 characters.
        :param pulumi.Input[int] identity_id: Organization identity ID.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] member_uins: Member Uin list. Up to 10.
        :param pulumi.Input[str] policy_name: Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrgMemberPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a organization org_member_policy_attachment

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        org_member_policy_attachment = tencentcloud.organization.OrgMemberPolicyAttachment("orgMemberPolicyAttachment",
            identity_id=1,
            member_uins=[
                100033905366,
                100033905356,
            ],
            policy_name="example-iac")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        organization org_member_policy_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Organization/orgMemberPolicyAttachment:OrgMemberPolicyAttachment org_member_policy_attachment org_member_policy_attachment_id
        ```

        :param str resource_name: The name of the resource.
        :param OrgMemberPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrgMemberPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_id: Optional[pulumi.Input[int]] = None,
                 member_uins: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrgMemberPolicyAttachmentArgs.__new__(OrgMemberPolicyAttachmentArgs)

            __props__.__dict__["description"] = description
            if identity_id is None and not opts.urn:
                raise TypeError("Missing required property 'identity_id'")
            __props__.__dict__["identity_id"] = identity_id
            if member_uins is None and not opts.urn:
                raise TypeError("Missing required property 'member_uins'")
            __props__.__dict__["member_uins"] = member_uins
            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
        super(OrgMemberPolicyAttachment, __self__).__init__(
            'tencentcloud:Organization/orgMemberPolicyAttachment:OrgMemberPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            identity_id: Optional[pulumi.Input[int]] = None,
            member_uins: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            policy_name: Optional[pulumi.Input[str]] = None) -> 'OrgMemberPolicyAttachment':
        """
        Get an existing OrgMemberPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Notes.The maximum length is 128 characters.
        :param pulumi.Input[int] identity_id: Organization identity ID.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] member_uins: Member Uin list. Up to 10.
        :param pulumi.Input[str] policy_name: Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrgMemberPolicyAttachmentState.__new__(_OrgMemberPolicyAttachmentState)

        __props__.__dict__["description"] = description
        __props__.__dict__["identity_id"] = identity_id
        __props__.__dict__["member_uins"] = member_uins
        __props__.__dict__["policy_name"] = policy_name
        return OrgMemberPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Notes.The maximum length is 128 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="identityId")
    def identity_id(self) -> pulumi.Output[int]:
        """
        Organization identity ID.
        """
        return pulumi.get(self, "identity_id")

    @property
    @pulumi.getter(name="memberUins")
    def member_uins(self) -> pulumi.Output[Sequence[int]]:
        """
        Member Uin list. Up to 10.
        """
        return pulumi.get(self, "member_uins")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        Policy name.The maximum length is 128 characters, supporting English letters, numbers, and symbols +=,.@_-.
        """
        return pulumi.get(self, "policy_name")

