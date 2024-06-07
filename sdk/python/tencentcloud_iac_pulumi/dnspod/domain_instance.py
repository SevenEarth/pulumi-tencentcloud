# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainInstanceArgs', 'DomainInstance']

@pulumi.input_type
class DomainInstanceArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 group_id: Optional[pulumi.Input[int]] = None,
                 is_mark: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DomainInstance resource.
        :param pulumi.Input[str] domain: The Domain.
        :param pulumi.Input[int] group_id: The Group Id of Domain.
        :param pulumi.Input[str] is_mark: Whether to Mark the Domain.
        :param pulumi.Input[str] remark: The remark of Domain.
        :param pulumi.Input[str] status: The status of Domain.
        """
        pulumi.set(__self__, "domain", domain)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if is_mark is not None:
            pulumi.set(__self__, "is_mark", is_mark)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The Group Id of Domain.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="isMark")
    def is_mark(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to Mark the Domain.
        """
        return pulumi.get(self, "is_mark")

    @is_mark.setter
    def is_mark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "is_mark", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark of Domain.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of Domain.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _DomainInstanceState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 is_mark: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 slave_dns: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DomainInstance resources.
        :param pulumi.Input[str] create_time: Create time of the domain.
        :param pulumi.Input[str] domain: The Domain.
        :param pulumi.Input[int] group_id: The Group Id of Domain.
        :param pulumi.Input[str] is_mark: Whether to Mark the Domain.
        :param pulumi.Input[str] remark: The remark of Domain.
        :param pulumi.Input[str] slave_dns: Is secondary DNS enabled.
        :param pulumi.Input[str] status: The status of Domain.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if is_mark is not None:
            pulumi.set(__self__, "is_mark", is_mark)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if slave_dns is not None:
            pulumi.set(__self__, "slave_dns", slave_dns)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time of the domain.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The Group Id of Domain.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="isMark")
    def is_mark(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to Mark the Domain.
        """
        return pulumi.get(self, "is_mark")

    @is_mark.setter
    def is_mark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "is_mark", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark of Domain.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="slaveDns")
    def slave_dns(self) -> Optional[pulumi.Input[str]]:
        """
        Is secondary DNS enabled.
        """
        return pulumi.get(self, "slave_dns")

    @slave_dns.setter
    def slave_dns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slave_dns", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of Domain.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class DomainInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 is_mark: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create a DnsPod Domain instance.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.dnspod.DomainInstance("foo",
            domain="hello.com",
            remark="this is demo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DnsPod Domain instance can be imported, e.g.

        ```sh
        $ pulumi import tencentcloud:Dnspod/domainInstance:DomainInstance foo domain
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: The Domain.
        :param pulumi.Input[int] group_id: The Group Id of Domain.
        :param pulumi.Input[str] is_mark: Whether to Mark the Domain.
        :param pulumi.Input[str] remark: The remark of Domain.
        :param pulumi.Input[str] status: The status of Domain.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a DnsPod Domain instance.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.dnspod.DomainInstance("foo",
            domain="hello.com",
            remark="this is demo")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        DnsPod Domain instance can be imported, e.g.

        ```sh
        $ pulumi import tencentcloud:Dnspod/domainInstance:DomainInstance foo domain
        ```

        :param str resource_name: The name of the resource.
        :param DomainInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 is_mark: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainInstanceArgs.__new__(DomainInstanceArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["is_mark"] = is_mark
            __props__.__dict__["remark"] = remark
            __props__.__dict__["status"] = status
            __props__.__dict__["create_time"] = None
            __props__.__dict__["slave_dns"] = None
        super(DomainInstance, __self__).__init__(
            'tencentcloud:Dnspod/domainInstance:DomainInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            is_mark: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            slave_dns: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'DomainInstance':
        """
        Get an existing DomainInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Create time of the domain.
        :param pulumi.Input[str] domain: The Domain.
        :param pulumi.Input[int] group_id: The Group Id of Domain.
        :param pulumi.Input[str] is_mark: Whether to Mark the Domain.
        :param pulumi.Input[str] remark: The remark of Domain.
        :param pulumi.Input[str] slave_dns: Is secondary DNS enabled.
        :param pulumi.Input[str] status: The status of Domain.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainInstanceState.__new__(_DomainInstanceState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["domain"] = domain
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["is_mark"] = is_mark
        __props__.__dict__["remark"] = remark
        __props__.__dict__["slave_dns"] = slave_dns
        __props__.__dict__["status"] = status
        return DomainInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time of the domain.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[Optional[int]]:
        """
        The Group Id of Domain.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="isMark")
    def is_mark(self) -> pulumi.Output[str]:
        """
        Whether to Mark the Domain.
        """
        return pulumi.get(self, "is_mark")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        The remark of Domain.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="slaveDns")
    def slave_dns(self) -> pulumi.Output[str]:
        """
        Is secondary DNS enabled.
        """
        return pulumi.get(self, "slave_dns")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of Domain.
        """
        return pulumi.get(self, "status")

