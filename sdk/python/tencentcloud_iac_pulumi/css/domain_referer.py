# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainRefererArgs', 'DomainReferer']

@pulumi.input_type
class DomainRefererArgs:
    def __init__(__self__, *,
                 allow_empty: pulumi.Input[int],
                 domain_name: pulumi.Input[str],
                 enable: pulumi.Input[int],
                 rules: pulumi.Input[str],
                 type: pulumi.Input[int]):
        """
        The set of arguments for constructing a DomainReferer resource.
        :param pulumi.Input[int] allow_empty: Allow blank referers, 0: not allowed, 1: allowed.
        :param pulumi.Input[str] domain_name: Domain Name.
        :param pulumi.Input[int] enable: Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        :param pulumi.Input[str] rules: The list of referers to; separate.
        :param pulumi.Input[int] type: List type: 0: blacklist, 1: whitelist.
        """
        pulumi.set(__self__, "allow_empty", allow_empty)
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "enable", enable)
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="allowEmpty")
    def allow_empty(self) -> pulumi.Input[int]:
        """
        Allow blank referers, 0: not allowed, 1: allowed.
        """
        return pulumi.get(self, "allow_empty")

    @allow_empty.setter
    def allow_empty(self, value: pulumi.Input[int]):
        pulumi.set(self, "allow_empty", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        Domain Name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[int]:
        """
        Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[int]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[str]:
        """
        The list of referers to; separate.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[str]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[int]:
        """
        List type: 0: blacklist, 1: whitelist.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[int]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _DomainRefererState:
    def __init__(__self__, *,
                 allow_empty: Optional[pulumi.Input[int]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[int]] = None,
                 rules: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering DomainReferer resources.
        :param pulumi.Input[int] allow_empty: Allow blank referers, 0: not allowed, 1: allowed.
        :param pulumi.Input[str] domain_name: Domain Name.
        :param pulumi.Input[int] enable: Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        :param pulumi.Input[str] rules: The list of referers to; separate.
        :param pulumi.Input[int] type: List type: 0: blacklist, 1: whitelist.
        """
        if allow_empty is not None:
            pulumi.set(__self__, "allow_empty", allow_empty)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="allowEmpty")
    def allow_empty(self) -> Optional[pulumi.Input[int]]:
        """
        Allow blank referers, 0: not allowed, 1: allowed.
        """
        return pulumi.get(self, "allow_empty")

    @allow_empty.setter
    def allow_empty(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "allow_empty", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Domain Name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[str]]:
        """
        The list of referers to; separate.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[int]]:
        """
        List type: 0: blacklist, 1: whitelist.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "type", value)


class DomainReferer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_empty: Optional[pulumi.Input[int]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[int]] = None,
                 rules: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a css domain_referer

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        domain_referer = tencentcloud.css.DomainReferer("domainReferer",
            allow_empty=1,
            domain_name="test122.jingxhu.top",
            enable=0,
            rules="example.com",
            type=1)
        ```

        ## Import

        css domain_referer can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Css/domainReferer:DomainReferer domain_referer domainName
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] allow_empty: Allow blank referers, 0: not allowed, 1: allowed.
        :param pulumi.Input[str] domain_name: Domain Name.
        :param pulumi.Input[int] enable: Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        :param pulumi.Input[str] rules: The list of referers to; separate.
        :param pulumi.Input[int] type: List type: 0: blacklist, 1: whitelist.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainRefererArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a css domain_referer

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        domain_referer = tencentcloud.css.DomainReferer("domainReferer",
            allow_empty=1,
            domain_name="test122.jingxhu.top",
            enable=0,
            rules="example.com",
            type=1)
        ```

        ## Import

        css domain_referer can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Css/domainReferer:DomainReferer domain_referer domainName
        ```

        :param str resource_name: The name of the resource.
        :param DomainRefererArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainRefererArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_empty: Optional[pulumi.Input[int]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[int]] = None,
                 rules: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[int]] = None,
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
            __props__ = DomainRefererArgs.__new__(DomainRefererArgs)

            if allow_empty is None and not opts.urn:
                raise TypeError("Missing required property 'allow_empty'")
            __props__.__dict__["allow_empty"] = allow_empty
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            if enable is None and not opts.urn:
                raise TypeError("Missing required property 'enable'")
            __props__.__dict__["enable"] = enable
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(DomainReferer, __self__).__init__(
            'tencentcloud:Css/domainReferer:DomainReferer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_empty: Optional[pulumi.Input[int]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[int]] = None,
            rules: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[int]] = None) -> 'DomainReferer':
        """
        Get an existing DomainReferer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] allow_empty: Allow blank referers, 0: not allowed, 1: allowed.
        :param pulumi.Input[str] domain_name: Domain Name.
        :param pulumi.Input[int] enable: Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        :param pulumi.Input[str] rules: The list of referers to; separate.
        :param pulumi.Input[int] type: List type: 0: blacklist, 1: whitelist.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainRefererState.__new__(_DomainRefererState)

        __props__.__dict__["allow_empty"] = allow_empty
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["enable"] = enable
        __props__.__dict__["rules"] = rules
        __props__.__dict__["type"] = type
        return DomainReferer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowEmpty")
    def allow_empty(self) -> pulumi.Output[int]:
        """
        Allow blank referers, 0: not allowed, 1: allowed.
        """
        return pulumi.get(self, "allow_empty")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Domain Name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[int]:
        """
        Whether to enable the referer blacklist authentication of the current domain name,`0`: off, `1`: on.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[str]:
        """
        The list of referers to; separate.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[int]:
        """
        List type: 0: blacklist, 1: whitelist.
        """
        return pulumi.get(self, "type")

