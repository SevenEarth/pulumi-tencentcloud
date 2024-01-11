# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainLockArgs', 'DomainLock']

@pulumi.input_type
class DomainLockArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 lock_days: pulumi.Input[int]):
        """
        The set of arguments for constructing a DomainLock resource.
        :param pulumi.Input[str] domain: Domain name.
        :param pulumi.Input[int] lock_days: The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "lock_days", lock_days)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="lockDays")
    def lock_days(self) -> pulumi.Input[int]:
        """
        The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        return pulumi.get(self, "lock_days")

    @lock_days.setter
    def lock_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "lock_days", value)


@pulumi.input_type
class _DomainLockState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 lock_code: Optional[pulumi.Input[str]] = None,
                 lock_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering DomainLock resources.
        :param pulumi.Input[str] domain: Domain name.
        :param pulumi.Input[str] lock_code: Domain unlock code, can be obtained through the ModifyDomainLock interface.
        :param pulumi.Input[int] lock_days: The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if lock_code is not None:
            pulumi.set(__self__, "lock_code", lock_code)
        if lock_days is not None:
            pulumi.set(__self__, "lock_days", lock_days)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="lockCode")
    def lock_code(self) -> Optional[pulumi.Input[str]]:
        """
        Domain unlock code, can be obtained through the ModifyDomainLock interface.
        """
        return pulumi.get(self, "lock_code")

    @lock_code.setter
    def lock_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lock_code", value)

    @property
    @pulumi.getter(name="lockDays")
    def lock_days(self) -> Optional[pulumi.Input[int]]:
        """
        The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        return pulumi.get(self, "lock_days")

    @lock_days.setter
    def lock_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lock_days", value)


class DomainLock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 lock_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a dnspod domain_lock

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        domain_lock = tencentcloud.dnspod.DomainLock("domainLock",
            domain="dnspod.cn",
            lock_days=30)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain name.
        :param pulumi.Input[int] lock_days: The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainLockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dnspod domain_lock

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        domain_lock = tencentcloud.dnspod.DomainLock("domainLock",
            domain="dnspod.cn",
            lock_days=30)
        ```

        :param str resource_name: The name of the resource.
        :param DomainLockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainLockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 lock_days: Optional[pulumi.Input[int]] = None,
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
            __props__ = DomainLockArgs.__new__(DomainLockArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if lock_days is None and not opts.urn:
                raise TypeError("Missing required property 'lock_days'")
            __props__.__dict__["lock_days"] = lock_days
            __props__.__dict__["lock_code"] = None
        super(DomainLock, __self__).__init__(
            'tencentcloud:Dnspod/domainLock:DomainLock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            lock_code: Optional[pulumi.Input[str]] = None,
            lock_days: Optional[pulumi.Input[int]] = None) -> 'DomainLock':
        """
        Get an existing DomainLock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain name.
        :param pulumi.Input[str] lock_code: Domain unlock code, can be obtained through the ModifyDomainLock interface.
        :param pulumi.Input[int] lock_days: The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainLockState.__new__(_DomainLockState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["lock_code"] = lock_code
        __props__.__dict__["lock_days"] = lock_days
        return DomainLock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="lockCode")
    def lock_code(self) -> pulumi.Output[str]:
        """
        Domain unlock code, can be obtained through the ModifyDomainLock interface.
        """
        return pulumi.get(self, "lock_code")

    @property
    @pulumi.getter(name="lockDays")
    def lock_days(self) -> pulumi.Output[int]:
        """
        The number of max days to lock the domain+ Old packages: D_FREE 30 days, D_PLUS 90 days, D_EXTRA 30 days, D_EXPERT 60 days, D_ULTRA 365 days+ New packages: DP_FREE 365 days, DP_PLUS 365 days, DP_EXTRA 365 days, DP_EXPERT 365 days, DP_ULTRA 365 days.
        """
        return pulumi.get(self, "lock_days")
