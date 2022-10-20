# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] instance_id: Instance ID that the account belongs to.
        :param pulumi.Input[str] password: Password of the SQL Server account.
        :param pulumi.Input[bool] is_admin: Indicate that the account is root account or not.
        :param pulumi.Input[str] name: Name of the SQL Server account.
        :param pulumi.Input[str] remark: Remark of the SQL Server account.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "password", password)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID that the account belongs to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password of the SQL Server account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate that the account is root account or not.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SQL Server account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of the SQL Server account.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] create_time: Create time of the SQL Server account.
        :param pulumi.Input[str] instance_id: Instance ID that the account belongs to.
        :param pulumi.Input[bool] is_admin: Indicate that the account is root account or not.
        :param pulumi.Input[str] name: Name of the SQL Server account.
        :param pulumi.Input[str] password: Password of the SQL Server account.
        :param pulumi.Input[str] remark: Remark of the SQL Server account.
        :param pulumi.Input[int] status: Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
        :param pulumi.Input[str] update_time: Last updated time of the SQL Server account.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the SQL Server account.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID that the account belongs to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate that the account is root account or not.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SQL Server account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password of the SQL Server account.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remark of the SQL Server account.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last updated time of the SQL Server account.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create SQL Server account

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.sqlserver.Account("foo",
            instance_id=tencentcloud_sqlserver_instance["example"]["id"],
            password="test1233",
            remark="testt")
        ```

        ## Import

        SQL Server account can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/account:Account foo mssql-3cdq7kx5#tf_sqlserver_account
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance ID that the account belongs to.
        :param pulumi.Input[bool] is_admin: Indicate that the account is root account or not.
        :param pulumi.Input[str] name: Name of the SQL Server account.
        :param pulumi.Input[str] password: Password of the SQL Server account.
        :param pulumi.Input[str] remark: Remark of the SQL Server account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create SQL Server account

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.sqlserver.Account("foo",
            instance_id=tencentcloud_sqlserver_instance["example"]["id"],
            password="test1233",
            remark="testt")
        ```

        ## Import

        SQL Server account can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/account:Account foo mssql-3cdq7kx5#tf_sqlserver_account
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
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
            __props__ = AccountArgs.__new__(AccountArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["is_admin"] = is_admin
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["remark"] = remark
            __props__.__dict__["create_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
        super(Account, __self__).__init__(
            'tencentcloud:Sqlserver/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_admin: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time of the SQL Server account.
        :param pulumi.Input[str] instance_id: Instance ID that the account belongs to.
        :param pulumi.Input[bool] is_admin: Indicate that the account is root account or not.
        :param pulumi.Input[str] name: Name of the SQL Server account.
        :param pulumi.Input[str] password: Password of the SQL Server account.
        :param pulumi.Input[str] remark: Remark of the SQL Server account.
        :param pulumi.Input[int] status: Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
        :param pulumi.Input[str] update_time: Last updated time of the SQL Server account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_admin"] = is_admin
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["remark"] = remark
        __props__.__dict__["status"] = status
        __props__.__dict__["update_time"] = update_time
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the SQL Server account.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID that the account belongs to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate that the account is root account or not.
        """
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the SQL Server account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password of the SQL Server account.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remark of the SQL Server account.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        Status of the SQL Server account. Valid values: 1, 2, 3, 4. 1 for creating, 2 for running, 3 for modifying, 4 for resetting password, -1 for deleting.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last updated time of the SQL Server account.
        """
        return pulumi.get(self, "update_time")

