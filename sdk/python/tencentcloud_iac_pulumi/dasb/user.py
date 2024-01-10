# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 real_name: pulumi.Input[str],
                 user_name: pulumi.Input[str],
                 auth_type: Optional[pulumi.Input[int]] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 validate_from: Optional[pulumi.Input[str]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validate_to: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] real_name: Real name, maximum length 20 characters, cannot contain blank characters.
        :param pulumi.Input[str] user_name: Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        :param pulumi.Input[int] auth_type: Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        :param pulumi.Input[str] department_id: Department ID, such as: 1.2.3.
        :param pulumi.Input[str] email: Email.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_id_sets: The set of user group IDs to which it belongs.
        :param pulumi.Input[str] phone: Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        :param pulumi.Input[str] validate_from: User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        :param pulumi.Input[str] validate_time: Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        :param pulumi.Input[str] validate_to: User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        pulumi.set(__self__, "real_name", real_name)
        pulumi.set(__self__, "user_name", user_name)
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if department_id is not None:
            pulumi.set(__self__, "department_id", department_id)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if group_id_sets is not None:
            pulumi.set(__self__, "group_id_sets", group_id_sets)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if validate_from is not None:
            pulumi.set(__self__, "validate_from", validate_from)
        if validate_time is not None:
            pulumi.set(__self__, "validate_time", validate_time)
        if validate_to is not None:
            pulumi.set(__self__, "validate_to", validate_to)

    @property
    @pulumi.getter(name="realName")
    def real_name(self) -> pulumi.Input[str]:
        """
        Real name, maximum length 20 characters, cannot contain blank characters.
        """
        return pulumi.get(self, "real_name")

    @real_name.setter
    def real_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "real_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[int]]:
        """
        Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> Optional[pulumi.Input[str]]:
        """
        Department ID, such as: 1.2.3.
        """
        return pulumi.get(self, "department_id")

    @department_id.setter
    def department_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "department_id", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="groupIdSets")
    def group_id_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The set of user group IDs to which it belongs.
        """
        return pulumi.get(self, "group_id_sets")

    @group_id_sets.setter
    def group_id_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "group_id_sets", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="validateFrom")
    def validate_from(self) -> Optional[pulumi.Input[str]]:
        """
        User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_from")

    @validate_from.setter
    def validate_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_from", value)

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> Optional[pulumi.Input[str]]:
        """
        Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        """
        return pulumi.get(self, "validate_time")

    @validate_time.setter
    def validate_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_time", value)

    @property
    @pulumi.getter(name="validateTo")
    def validate_to(self) -> Optional[pulumi.Input[str]]:
        """
        User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_to")

    @validate_to.setter
    def validate_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_to", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 auth_type: Optional[pulumi.Input[int]] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 real_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_from: Optional[pulumi.Input[str]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validate_to: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[int] auth_type: Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        :param pulumi.Input[str] department_id: Department ID, such as: 1.2.3.
        :param pulumi.Input[str] email: Email.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_id_sets: The set of user group IDs to which it belongs.
        :param pulumi.Input[str] phone: Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        :param pulumi.Input[str] real_name: Real name, maximum length 20 characters, cannot contain blank characters.
        :param pulumi.Input[str] user_name: Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        :param pulumi.Input[str] validate_from: User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        :param pulumi.Input[str] validate_time: Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        :param pulumi.Input[str] validate_to: User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        if auth_type is not None:
            pulumi.set(__self__, "auth_type", auth_type)
        if department_id is not None:
            pulumi.set(__self__, "department_id", department_id)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if group_id_sets is not None:
            pulumi.set(__self__, "group_id_sets", group_id_sets)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if real_name is not None:
            pulumi.set(__self__, "real_name", real_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)
        if validate_from is not None:
            pulumi.set(__self__, "validate_from", validate_from)
        if validate_time is not None:
            pulumi.set(__self__, "validate_time", validate_time)
        if validate_to is not None:
            pulumi.set(__self__, "validate_to", validate_to)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> Optional[pulumi.Input[int]]:
        """
        Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        """
        return pulumi.get(self, "auth_type")

    @auth_type.setter
    def auth_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auth_type", value)

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> Optional[pulumi.Input[str]]:
        """
        Department ID, such as: 1.2.3.
        """
        return pulumi.get(self, "department_id")

    @department_id.setter
    def department_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "department_id", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="groupIdSets")
    def group_id_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The set of user group IDs to which it belongs.
        """
        return pulumi.get(self, "group_id_sets")

    @group_id_sets.setter
    def group_id_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "group_id_sets", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter(name="realName")
    def real_name(self) -> Optional[pulumi.Input[str]]:
        """
        Real name, maximum length 20 characters, cannot contain blank characters.
        """
        return pulumi.get(self, "real_name")

    @real_name.setter
    def real_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "real_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)

    @property
    @pulumi.getter(name="validateFrom")
    def validate_from(self) -> Optional[pulumi.Input[str]]:
        """
        User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_from")

    @validate_from.setter
    def validate_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_from", value)

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> Optional[pulumi.Input[str]]:
        """
        Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        """
        return pulumi.get(self, "validate_time")

    @validate_time.setter
    def validate_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_time", value)

    @property
    @pulumi.getter(name="validateTo")
    def validate_to(self) -> Optional[pulumi.Input[str]]:
        """
        User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_to")

    @validate_to.setter
    def validate_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_to", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input[int]] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 real_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_from: Optional[pulumi.Input[str]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validate_to: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dasb user

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.User("example",
            auth_type=0,
            department_id="1.2",
            email="demo@tencent.com",
            phone="+86|18345678782",
            real_name="terraform",
            user_name="tf_example",
            validate_from="2023-09-22T02:00:00+08:00",
            validate_to="2023-09-23T03:00:00+08:00")
        ```

        ## Import

        dasb user can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dasb/user:User example 134
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auth_type: Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        :param pulumi.Input[str] department_id: Department ID, such as: 1.2.3.
        :param pulumi.Input[str] email: Email.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_id_sets: The set of user group IDs to which it belongs.
        :param pulumi.Input[str] phone: Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        :param pulumi.Input[str] real_name: Real name, maximum length 20 characters, cannot contain blank characters.
        :param pulumi.Input[str] user_name: Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        :param pulumi.Input[str] validate_from: User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        :param pulumi.Input[str] validate_time: Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        :param pulumi.Input[str] validate_to: User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dasb user

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.dasb.User("example",
            auth_type=0,
            department_id="1.2",
            email="demo@tencent.com",
            phone="+86|18345678782",
            real_name="terraform",
            user_name="tf_example",
            validate_from="2023-09-22T02:00:00+08:00",
            validate_to="2023-09-23T03:00:00+08:00")
        ```

        ## Import

        dasb user can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dasb/user:User example 134
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input[int]] = None,
                 department_id: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 real_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 validate_from: Optional[pulumi.Input[str]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validate_to: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["auth_type"] = auth_type
            __props__.__dict__["department_id"] = department_id
            __props__.__dict__["email"] = email
            __props__.__dict__["group_id_sets"] = group_id_sets
            __props__.__dict__["phone"] = phone
            if real_name is None and not opts.urn:
                raise TypeError("Missing required property 'real_name'")
            __props__.__dict__["real_name"] = real_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
            __props__.__dict__["validate_from"] = validate_from
            __props__.__dict__["validate_time"] = validate_time
            __props__.__dict__["validate_to"] = validate_to
        super(User, __self__).__init__(
            'tencentcloud:Dasb/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_type: Optional[pulumi.Input[int]] = None,
            department_id: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            group_id_sets: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            phone: Optional[pulumi.Input[str]] = None,
            real_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None,
            validate_from: Optional[pulumi.Input[str]] = None,
            validate_time: Optional[pulumi.Input[str]] = None,
            validate_to: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auth_type: Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        :param pulumi.Input[str] department_id: Department ID, such as: 1.2.3.
        :param pulumi.Input[str] email: Email.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] group_id_sets: The set of user group IDs to which it belongs.
        :param pulumi.Input[str] phone: Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        :param pulumi.Input[str] real_name: Real name, maximum length 20 characters, cannot contain blank characters.
        :param pulumi.Input[str] user_name: Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        :param pulumi.Input[str] validate_from: User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        :param pulumi.Input[str] validate_time: Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        :param pulumi.Input[str] validate_to: User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["auth_type"] = auth_type
        __props__.__dict__["department_id"] = department_id
        __props__.__dict__["email"] = email
        __props__.__dict__["group_id_sets"] = group_id_sets
        __props__.__dict__["phone"] = phone
        __props__.__dict__["real_name"] = real_name
        __props__.__dict__["user_name"] = user_name
        __props__.__dict__["validate_from"] = validate_from
        __props__.__dict__["validate_time"] = validate_time
        __props__.__dict__["validate_to"] = validate_to
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output[Optional[int]]:
        """
        Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. If not passed, the default is 0.
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter(name="departmentId")
    def department_id(self) -> pulumi.Output[Optional[str]]:
        """
        Department ID, such as: 1.2.3.
        """
        return pulumi.get(self, "department_id")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        Email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="groupIdSets")
    def group_id_sets(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        The set of user group IDs to which it belongs.
        """
        return pulumi.get(self, "group_id_sets")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional[str]]:
        """
        Fill in the mainland mobile phone number directly. If it is a number from other countries or regions, enter it in the format of country area code|mobile phone number. For example: +852|xxxxxxxx.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter(name="realName")
    def real_name(self) -> pulumi.Output[str]:
        """
        Real name, maximum length 20 characters, cannot contain blank characters.
        """
        return pulumi.get(self, "real_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        Username, 3-20 characters, must start with an English letter and cannot contain characters other than letters, numbers, '.', '_', '-'.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="validateFrom")
    def validate_from(self) -> pulumi.Output[Optional[str]]:
        """
        User effective time, such as: 2021-09-22T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_from")

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> pulumi.Output[Optional[str]]:
        """
        Access time period limit, a string composed of 0 and 1, length 168 (7 * 24), representing the time period the user is allowed to access in a week. The Nth character in the string represents the Nth hour of the week, 0 - means access is not allowed, 1 - means access is allowed.
        """
        return pulumi.get(self, "validate_time")

    @property
    @pulumi.getter(name="validateTo")
    def validate_to(self) -> pulumi.Output[Optional[str]]:
        """
        User expiration time, such as: 2021-09-23T00:00:00+00:00If the effective and expiry time are not filled in, the user will be valid for a long time.
        """
        return pulumi.get(self, "validate_to")

