# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FileSystemArgs', 'FileSystem']

@pulumi.input_type
class FileSystemArgs:
    def __init__(__self__, *,
                 capacity_quota: pulumi.Input[int],
                 file_system_name: pulumi.Input[str],
                 posix_acl: pulumi.Input[bool],
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ranger: Optional[pulumi.Input[bool]] = None,
                 ranger_service_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 super_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a FileSystem resource.
        :param pulumi.Input[int] capacity_quota: file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        :param pulumi.Input[str] file_system_name: file system name.
        :param pulumi.Input[bool] posix_acl: check POSIX ACL or not.
        :param pulumi.Input[str] description: desc of the file system.
        :param pulumi.Input[bool] enable_ranger: check the ranger address or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ranger_service_addresses: ranger address list, default empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] super_users: super users of the file system, default empty.
        """
        pulumi.set(__self__, "capacity_quota", capacity_quota)
        pulumi.set(__self__, "file_system_name", file_system_name)
        pulumi.set(__self__, "posix_acl", posix_acl)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_ranger is not None:
            pulumi.set(__self__, "enable_ranger", enable_ranger)
        if ranger_service_addresses is not None:
            pulumi.set(__self__, "ranger_service_addresses", ranger_service_addresses)
        if super_users is not None:
            pulumi.set(__self__, "super_users", super_users)

    @property
    @pulumi.getter(name="capacityQuota")
    def capacity_quota(self) -> pulumi.Input[int]:
        """
        file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        """
        return pulumi.get(self, "capacity_quota")

    @capacity_quota.setter
    def capacity_quota(self, value: pulumi.Input[int]):
        pulumi.set(self, "capacity_quota", value)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Input[str]:
        """
        file system name.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="posixAcl")
    def posix_acl(self) -> pulumi.Input[bool]:
        """
        check POSIX ACL or not.
        """
        return pulumi.get(self, "posix_acl")

    @posix_acl.setter
    def posix_acl(self, value: pulumi.Input[bool]):
        pulumi.set(self, "posix_acl", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        desc of the file system.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableRanger")
    def enable_ranger(self) -> Optional[pulumi.Input[bool]]:
        """
        check the ranger address or not.
        """
        return pulumi.get(self, "enable_ranger")

    @enable_ranger.setter
    def enable_ranger(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ranger", value)

    @property
    @pulumi.getter(name="rangerServiceAddresses")
    def ranger_service_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ranger address list, default empty.
        """
        return pulumi.get(self, "ranger_service_addresses")

    @ranger_service_addresses.setter
    def ranger_service_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ranger_service_addresses", value)

    @property
    @pulumi.getter(name="superUsers")
    def super_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        super users of the file system, default empty.
        """
        return pulumi.get(self, "super_users")

    @super_users.setter
    def super_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "super_users", value)


@pulumi.input_type
class _FileSystemState:
    def __init__(__self__, *,
                 capacity_quota: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ranger: Optional[pulumi.Input[bool]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 posix_acl: Optional[pulumi.Input[bool]] = None,
                 ranger_service_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 super_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering FileSystem resources.
        :param pulumi.Input[int] capacity_quota: file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        :param pulumi.Input[str] description: desc of the file system.
        :param pulumi.Input[bool] enable_ranger: check the ranger address or not.
        :param pulumi.Input[str] file_system_name: file system name.
        :param pulumi.Input[bool] posix_acl: check POSIX ACL or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ranger_service_addresses: ranger address list, default empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] super_users: super users of the file system, default empty.
        """
        if capacity_quota is not None:
            pulumi.set(__self__, "capacity_quota", capacity_quota)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_ranger is not None:
            pulumi.set(__self__, "enable_ranger", enable_ranger)
        if file_system_name is not None:
            pulumi.set(__self__, "file_system_name", file_system_name)
        if posix_acl is not None:
            pulumi.set(__self__, "posix_acl", posix_acl)
        if ranger_service_addresses is not None:
            pulumi.set(__self__, "ranger_service_addresses", ranger_service_addresses)
        if super_users is not None:
            pulumi.set(__self__, "super_users", super_users)

    @property
    @pulumi.getter(name="capacityQuota")
    def capacity_quota(self) -> Optional[pulumi.Input[int]]:
        """
        file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        """
        return pulumi.get(self, "capacity_quota")

    @capacity_quota.setter
    def capacity_quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity_quota", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        desc of the file system.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableRanger")
    def enable_ranger(self) -> Optional[pulumi.Input[bool]]:
        """
        check the ranger address or not.
        """
        return pulumi.get(self, "enable_ranger")

    @enable_ranger.setter
    def enable_ranger(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ranger", value)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> Optional[pulumi.Input[str]]:
        """
        file system name.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="posixAcl")
    def posix_acl(self) -> Optional[pulumi.Input[bool]]:
        """
        check POSIX ACL or not.
        """
        return pulumi.get(self, "posix_acl")

    @posix_acl.setter
    def posix_acl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "posix_acl", value)

    @property
    @pulumi.getter(name="rangerServiceAddresses")
    def ranger_service_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        ranger address list, default empty.
        """
        return pulumi.get(self, "ranger_service_addresses")

    @ranger_service_addresses.setter
    def ranger_service_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ranger_service_addresses", value)

    @property
    @pulumi.getter(name="superUsers")
    def super_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        super users of the file system, default empty.
        """
        return pulumi.get(self, "super_users")

    @super_users.setter
    def super_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "super_users", value)


class FileSystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_quota: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ranger: Optional[pulumi.Input[bool]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 posix_acl: Optional[pulumi.Input[bool]] = None,
                 ranger_service_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 super_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a chdfs file_system

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        file_system = tencentcloud.chdfs.FileSystem("fileSystem",
            capacity_quota=10995116277760,
            description="file system for terraform test",
            enable_ranger=True,
            file_system_name="terraform-test",
            posix_acl=False,
            ranger_service_addresses=[
                "127.0.0.1:80",
                "127.0.0.1:8000",
            ],
            super_users=[
                "terraform",
                "iac",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        chdfs file_system can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Chdfs/fileSystem:FileSystem file_system file_system_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity_quota: file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        :param pulumi.Input[str] description: desc of the file system.
        :param pulumi.Input[bool] enable_ranger: check the ranger address or not.
        :param pulumi.Input[str] file_system_name: file system name.
        :param pulumi.Input[bool] posix_acl: check POSIX ACL or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ranger_service_addresses: ranger address list, default empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] super_users: super users of the file system, default empty.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a chdfs file_system

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        file_system = tencentcloud.chdfs.FileSystem("fileSystem",
            capacity_quota=10995116277760,
            description="file system for terraform test",
            enable_ranger=True,
            file_system_name="terraform-test",
            posix_acl=False,
            ranger_service_addresses=[
                "127.0.0.1:80",
                "127.0.0.1:8000",
            ],
            super_users=[
                "terraform",
                "iac",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        chdfs file_system can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Chdfs/fileSystem:FileSystem file_system file_system_id
        ```

        :param str resource_name: The name of the resource.
        :param FileSystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity_quota: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ranger: Optional[pulumi.Input[bool]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 posix_acl: Optional[pulumi.Input[bool]] = None,
                 ranger_service_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 super_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileSystemArgs.__new__(FileSystemArgs)

            if capacity_quota is None and not opts.urn:
                raise TypeError("Missing required property 'capacity_quota'")
            __props__.__dict__["capacity_quota"] = capacity_quota
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_ranger"] = enable_ranger
            if file_system_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_name'")
            __props__.__dict__["file_system_name"] = file_system_name
            if posix_acl is None and not opts.urn:
                raise TypeError("Missing required property 'posix_acl'")
            __props__.__dict__["posix_acl"] = posix_acl
            __props__.__dict__["ranger_service_addresses"] = ranger_service_addresses
            __props__.__dict__["super_users"] = super_users
        super(FileSystem, __self__).__init__(
            'tencentcloud:Chdfs/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capacity_quota: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_ranger: Optional[pulumi.Input[bool]] = None,
            file_system_name: Optional[pulumi.Input[str]] = None,
            posix_acl: Optional[pulumi.Input[bool]] = None,
            ranger_service_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            super_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'FileSystem':
        """
        Get an existing FileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity_quota: file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        :param pulumi.Input[str] description: desc of the file system.
        :param pulumi.Input[bool] enable_ranger: check the ranger address or not.
        :param pulumi.Input[str] file_system_name: file system name.
        :param pulumi.Input[bool] posix_acl: check POSIX ACL or not.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ranger_service_addresses: ranger address list, default empty.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] super_users: super users of the file system, default empty.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileSystemState.__new__(_FileSystemState)

        __props__.__dict__["capacity_quota"] = capacity_quota
        __props__.__dict__["description"] = description
        __props__.__dict__["enable_ranger"] = enable_ranger
        __props__.__dict__["file_system_name"] = file_system_name
        __props__.__dict__["posix_acl"] = posix_acl
        __props__.__dict__["ranger_service_addresses"] = ranger_service_addresses
        __props__.__dict__["super_users"] = super_users
        return FileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="capacityQuota")
    def capacity_quota(self) -> pulumi.Output[int]:
        """
        file system capacity. min 1GB, max 1PB, CapacityQuota is N * 1073741824.
        """
        return pulumi.get(self, "capacity_quota")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        desc of the file system.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableRanger")
    def enable_ranger(self) -> pulumi.Output[Optional[bool]]:
        """
        check the ranger address or not.
        """
        return pulumi.get(self, "enable_ranger")

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Output[str]:
        """
        file system name.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="posixAcl")
    def posix_acl(self) -> pulumi.Output[bool]:
        """
        check POSIX ACL or not.
        """
        return pulumi.get(self, "posix_acl")

    @property
    @pulumi.getter(name="rangerServiceAddresses")
    def ranger_service_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        ranger address list, default empty.
        """
        return pulumi.get(self, "ranger_service_addresses")

    @property
    @pulumi.getter(name="superUsers")
    def super_users(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        super users of the file system, default empty.
        """
        return pulumi.get(self, "super_users")

