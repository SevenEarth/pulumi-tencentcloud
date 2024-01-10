# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GlobalDomainDnsArgs', 'GlobalDomainDns']

@pulumi.input_type
class GlobalDomainDnsArgs:
    def __init__(__self__, *,
                 domain_id: pulumi.Input[str],
                 nation_country_inner_codes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 proxy_id_lists: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a GlobalDomainDns resource.
        :param pulumi.Input[str] domain_id: Domain Id.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nation_country_inner_codes: Nation Country Inner Codes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxy_id_lists: Proxy Id List.
        """
        pulumi.set(__self__, "domain_id", domain_id)
        pulumi.set(__self__, "nation_country_inner_codes", nation_country_inner_codes)
        pulumi.set(__self__, "proxy_id_lists", proxy_id_lists)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Input[str]:
        """
        Domain Id.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter(name="nationCountryInnerCodes")
    def nation_country_inner_codes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Nation Country Inner Codes.
        """
        return pulumi.get(self, "nation_country_inner_codes")

    @nation_country_inner_codes.setter
    def nation_country_inner_codes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "nation_country_inner_codes", value)

    @property
    @pulumi.getter(name="proxyIdLists")
    def proxy_id_lists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Proxy Id List.
        """
        return pulumi.get(self, "proxy_id_lists")

    @proxy_id_lists.setter
    def proxy_id_lists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "proxy_id_lists", value)


@pulumi.input_type
class _GlobalDomainDnsState:
    def __init__(__self__, *,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 nation_country_inner_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 proxy_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering GlobalDomainDns resources.
        :param pulumi.Input[str] domain_id: Domain Id.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nation_country_inner_codes: Nation Country Inner Codes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxy_id_lists: Proxy Id List.
        """
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if nation_country_inner_codes is not None:
            pulumi.set(__self__, "nation_country_inner_codes", nation_country_inner_codes)
        if proxy_id_lists is not None:
            pulumi.set(__self__, "proxy_id_lists", proxy_id_lists)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        Domain Id.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter(name="nationCountryInnerCodes")
    def nation_country_inner_codes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Nation Country Inner Codes.
        """
        return pulumi.get(self, "nation_country_inner_codes")

    @nation_country_inner_codes.setter
    def nation_country_inner_codes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nation_country_inner_codes", value)

    @property
    @pulumi.getter(name="proxyIdLists")
    def proxy_id_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Proxy Id List.
        """
        return pulumi.get(self, "proxy_id_lists")

    @proxy_id_lists.setter
    def proxy_id_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "proxy_id_lists", value)


class GlobalDomainDns(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 nation_country_inner_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 proxy_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a gaap global domain dns

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        global_domain_dns = tencentcloud.gaap.GlobalDomainDns("globalDomainDns",
            domain_id="dm-xxxxxx",
            nation_country_inner_codes=["101001"],
            proxy_id_lists=["link-xxxxxx"])
        ```

        ## Import

        gaap global_domain_dns can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/globalDomainDns:GlobalDomainDns global_domain_dns ${domainId}#${dnsRecordId}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: Domain Id.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nation_country_inner_codes: Nation Country Inner Codes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxy_id_lists: Proxy Id List.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalDomainDnsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a gaap global domain dns

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        global_domain_dns = tencentcloud.gaap.GlobalDomainDns("globalDomainDns",
            domain_id="dm-xxxxxx",
            nation_country_inner_codes=["101001"],
            proxy_id_lists=["link-xxxxxx"])
        ```

        ## Import

        gaap global_domain_dns can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Gaap/globalDomainDns:GlobalDomainDns global_domain_dns ${domainId}#${dnsRecordId}
        ```

        :param str resource_name: The name of the resource.
        :param GlobalDomainDnsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalDomainDnsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 nation_country_inner_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 proxy_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = GlobalDomainDnsArgs.__new__(GlobalDomainDnsArgs)

            if domain_id is None and not opts.urn:
                raise TypeError("Missing required property 'domain_id'")
            __props__.__dict__["domain_id"] = domain_id
            if nation_country_inner_codes is None and not opts.urn:
                raise TypeError("Missing required property 'nation_country_inner_codes'")
            __props__.__dict__["nation_country_inner_codes"] = nation_country_inner_codes
            if proxy_id_lists is None and not opts.urn:
                raise TypeError("Missing required property 'proxy_id_lists'")
            __props__.__dict__["proxy_id_lists"] = proxy_id_lists
        super(GlobalDomainDns, __self__).__init__(
            'tencentcloud:Gaap/globalDomainDns:GlobalDomainDns',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_id: Optional[pulumi.Input[str]] = None,
            nation_country_inner_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            proxy_id_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'GlobalDomainDns':
        """
        Get an existing GlobalDomainDns resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: Domain Id.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nation_country_inner_codes: Nation Country Inner Codes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] proxy_id_lists: Proxy Id List.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalDomainDnsState.__new__(_GlobalDomainDnsState)

        __props__.__dict__["domain_id"] = domain_id
        __props__.__dict__["nation_country_inner_codes"] = nation_country_inner_codes
        __props__.__dict__["proxy_id_lists"] = proxy_id_lists
        return GlobalDomainDns(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        Domain Id.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="nationCountryInnerCodes")
    def nation_country_inner_codes(self) -> pulumi.Output[Sequence[str]]:
        """
        Nation Country Inner Codes.
        """
        return pulumi.get(self, "nation_country_inner_codes")

    @property
    @pulumi.getter(name="proxyIdLists")
    def proxy_id_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        Proxy Id List.
        """
        return pulumi.get(self, "proxy_id_lists")

