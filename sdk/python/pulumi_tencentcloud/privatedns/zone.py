# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 account_vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]] = None,
                 dns_forward_status: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tag_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input[str] domain: Domain name, which must be in the format of standard TLD.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]] account_vpc_sets: List of authorized accounts' VPCs to associate with the private domain.
        :param pulumi.Input[str] dns_forward_status: Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        :param pulumi.Input[str] remark: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]] tag_sets: It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of the private dns zone.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]] vpc_sets: Associates the private domain to a VPC when it is created.
        """
        pulumi.set(__self__, "domain", domain)
        if account_vpc_sets is not None:
            pulumi.set(__self__, "account_vpc_sets", account_vpc_sets)
        if dns_forward_status is not None:
            pulumi.set(__self__, "dns_forward_status", dns_forward_status)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if tag_sets is not None:
            warnings.warn("""It has been deprecated from version 1.72.4. Use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tag_sets is deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.""")
        if tag_sets is not None:
            pulumi.set(__self__, "tag_sets", tag_sets)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_sets is not None:
            pulumi.set(__self__, "vpc_sets", vpc_sets)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain name, which must be in the format of standard TLD.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="accountVpcSets")
    def account_vpc_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]]:
        """
        List of authorized accounts' VPCs to associate with the private domain.
        """
        return pulumi.get(self, "account_vpc_sets")

    @account_vpc_sets.setter
    def account_vpc_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]]):
        pulumi.set(self, "account_vpc_sets", value)

    @property
    @pulumi.getter(name="dnsForwardStatus")
    def dns_forward_status(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        """
        return pulumi.get(self, "dns_forward_status")

    @dns_forward_status.setter
    def dns_forward_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_forward_status", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="tagSets")
    def tag_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]]:
        """
        It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        """
        return pulumi.get(self, "tag_sets")

    @tag_sets.setter
    def tag_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]]):
        pulumi.set(self, "tag_sets", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of the private dns zone.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcSets")
    def vpc_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]]:
        """
        Associates the private domain to a VPC when it is created.
        """
        return pulumi.get(self, "vpc_sets")

    @vpc_sets.setter
    def vpc_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]]):
        pulumi.set(self, "vpc_sets", value)


@pulumi.input_type
class _ZoneState:
    def __init__(__self__, *,
                 account_vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]] = None,
                 dns_forward_status: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tag_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]] account_vpc_sets: List of authorized accounts' VPCs to associate with the private domain.
        :param pulumi.Input[str] dns_forward_status: Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        :param pulumi.Input[str] domain: Domain name, which must be in the format of standard TLD.
        :param pulumi.Input[str] remark: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]] tag_sets: It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of the private dns zone.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]] vpc_sets: Associates the private domain to a VPC when it is created.
        """
        if account_vpc_sets is not None:
            pulumi.set(__self__, "account_vpc_sets", account_vpc_sets)
        if dns_forward_status is not None:
            pulumi.set(__self__, "dns_forward_status", dns_forward_status)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if tag_sets is not None:
            warnings.warn("""It has been deprecated from version 1.72.4. Use `tags` instead.""", DeprecationWarning)
            pulumi.log.warn("""tag_sets is deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.""")
        if tag_sets is not None:
            pulumi.set(__self__, "tag_sets", tag_sets)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_sets is not None:
            pulumi.set(__self__, "vpc_sets", vpc_sets)

    @property
    @pulumi.getter(name="accountVpcSets")
    def account_vpc_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]]:
        """
        List of authorized accounts' VPCs to associate with the private domain.
        """
        return pulumi.get(self, "account_vpc_sets")

    @account_vpc_sets.setter
    def account_vpc_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAccountVpcSetArgs']]]]):
        pulumi.set(self, "account_vpc_sets", value)

    @property
    @pulumi.getter(name="dnsForwardStatus")
    def dns_forward_status(self) -> Optional[pulumi.Input[str]]:
        """
        Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        """
        return pulumi.get(self, "dns_forward_status")

    @dns_forward_status.setter
    def dns_forward_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_forward_status", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name, which must be in the format of standard TLD.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="tagSets")
    def tag_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]]:
        """
        It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        """
        return pulumi.get(self, "tag_sets")

    @tag_sets.setter
    def tag_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneTagSetArgs']]]]):
        pulumi.set(self, "tag_sets", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of the private dns zone.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcSets")
    def vpc_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]]:
        """
        Associates the private domain to a VPC when it is created.
        """
        return pulumi.get(self, "vpc_sets")

    @vpc_sets.setter
    def vpc_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneVpcSetArgs']]]]):
        pulumi.set(self, "vpc_sets", value)


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAccountVpcSetArgs']]]]] = None,
                 dns_forward_status: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tag_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTagSetArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcSetArgs']]]]] = None,
                 __props__=None):
        """
        Provide a resource to create a Private Dns Zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.private_dns.Zone("foo",
            account_vpc_sets=[tencentcloud.private.dns.ZoneAccountVpcSetArgs(
                region="ap-guangzhou",
                uin="454xxxxxxx",
                uniq_vpc_id="vpc-xxxxx",
                vpc_name="test-redis",
            )],
            dns_forward_status="DISABLED",
            domain="domain.com",
            remark="test",
            tags={
                "created_by": [{}],
                "terraform": [{}],
            },
            vpc_sets=[tencentcloud.private.dns.ZoneVpcSetArgs(
                region="ap-guangzhou",
                uniq_vpc_id="vpc-xxxxx",
            )])
        ```

        ## Import

        Private Dns Zone can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:PrivateDns/zone:Zone foo zone_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAccountVpcSetArgs']]]] account_vpc_sets: List of authorized accounts' VPCs to associate with the private domain.
        :param pulumi.Input[str] dns_forward_status: Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        :param pulumi.Input[str] domain: Domain name, which must be in the format of standard TLD.
        :param pulumi.Input[str] remark: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTagSetArgs']]]] tag_sets: It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of the private dns zone.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcSetArgs']]]] vpc_sets: Associates the private domain to a VPC when it is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a Private Dns Zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        foo = tencentcloud.private_dns.Zone("foo",
            account_vpc_sets=[tencentcloud.private.dns.ZoneAccountVpcSetArgs(
                region="ap-guangzhou",
                uin="454xxxxxxx",
                uniq_vpc_id="vpc-xxxxx",
                vpc_name="test-redis",
            )],
            dns_forward_status="DISABLED",
            domain="domain.com",
            remark="test",
            tags={
                "created_by": [{}],
                "terraform": [{}],
            },
            vpc_sets=[tencentcloud.private.dns.ZoneVpcSetArgs(
                region="ap-guangzhou",
                uniq_vpc_id="vpc-xxxxx",
            )])
        ```

        ## Import

        Private Dns Zone can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:PrivateDns/zone:Zone foo zone_id
        ```

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAccountVpcSetArgs']]]]] = None,
                 dns_forward_status: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tag_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTagSetArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcSetArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneArgs.__new__(ZoneArgs)

            __props__.__dict__["account_vpc_sets"] = account_vpc_sets
            __props__.__dict__["dns_forward_status"] = dns_forward_status
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["remark"] = remark
            if tag_sets is not None and not opts.urn:
                warnings.warn("""It has been deprecated from version 1.72.4. Use `tags` instead.""", DeprecationWarning)
                pulumi.log.warn("""tag_sets is deprecated: It has been deprecated from version 1.72.4. Use `tags` instead.""")
            __props__.__dict__["tag_sets"] = tag_sets
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_sets"] = vpc_sets
        super(Zone, __self__).__init__(
            'tencentcloud:PrivateDns/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAccountVpcSetArgs']]]]] = None,
            dns_forward_status: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            tag_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTagSetArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_sets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcSetArgs']]]]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAccountVpcSetArgs']]]] account_vpc_sets: List of authorized accounts' VPCs to associate with the private domain.
        :param pulumi.Input[str] dns_forward_status: Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        :param pulumi.Input[str] domain: Domain name, which must be in the format of standard TLD.
        :param pulumi.Input[str] remark: Remarks.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneTagSetArgs']]]] tag_sets: It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of the private dns zone.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneVpcSetArgs']]]] vpc_sets: Associates the private domain to a VPC when it is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["account_vpc_sets"] = account_vpc_sets
        __props__.__dict__["dns_forward_status"] = dns_forward_status
        __props__.__dict__["domain"] = domain
        __props__.__dict__["remark"] = remark
        __props__.__dict__["tag_sets"] = tag_sets
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_sets"] = vpc_sets
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountVpcSets")
    def account_vpc_sets(self) -> pulumi.Output[Optional[Sequence['outputs.ZoneAccountVpcSet']]]:
        """
        List of authorized accounts' VPCs to associate with the private domain.
        """
        return pulumi.get(self, "account_vpc_sets")

    @property
    @pulumi.getter(name="dnsForwardStatus")
    def dns_forward_status(self) -> pulumi.Output[Optional[str]]:
        """
        Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        """
        return pulumi.get(self, "dns_forward_status")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain name, which must be in the format of standard TLD.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Remarks.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="tagSets")
    def tag_sets(self) -> pulumi.Output[Sequence['outputs.ZoneTagSet']]:
        """
        It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        """
        return pulumi.get(self, "tag_sets")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tags of the private dns zone.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcSets")
    def vpc_sets(self) -> pulumi.Output[Optional[Sequence['outputs.ZoneVpcSet']]]:
        """
        Associates the private domain to a VPC when it is created.
        """
        return pulumi.get(self, "vpc_sets")

