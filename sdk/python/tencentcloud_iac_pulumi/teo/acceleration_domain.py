# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AccelerationDomainArgs', 'AccelerationDomain']

@pulumi.input_type
class AccelerationDomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 origin_info: pulumi.Input['AccelerationDomainOriginInfoArgs'],
                 zone_id: pulumi.Input[str],
                 http_origin_port: Optional[pulumi.Input[int]] = None,
                 https_origin_port: Optional[pulumi.Input[int]] = None,
                 ipv6_status: Optional[pulumi.Input[str]] = None,
                 origin_protocol: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccelerationDomain resource.
        :param pulumi.Input[str] domain_name: Accelerated domain name.
        :param pulumi.Input['AccelerationDomainOriginInfoArgs'] origin_info: Details of the origin.
        :param pulumi.Input[str] zone_id: ID of the site related with the accelerated domain name.
        :param pulumi.Input[int] http_origin_port: HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        :param pulumi.Input[int] https_origin_port: HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        :param pulumi.Input[str] ipv6_status: IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        :param pulumi.Input[str] origin_protocol: Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        :param pulumi.Input[str] status: Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "origin_info", origin_info)
        pulumi.set(__self__, "zone_id", zone_id)
        if http_origin_port is not None:
            pulumi.set(__self__, "http_origin_port", http_origin_port)
        if https_origin_port is not None:
            pulumi.set(__self__, "https_origin_port", https_origin_port)
        if ipv6_status is not None:
            pulumi.set(__self__, "ipv6_status", ipv6_status)
        if origin_protocol is not None:
            pulumi.set(__self__, "origin_protocol", origin_protocol)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        Accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="originInfo")
    def origin_info(self) -> pulumi.Input['AccelerationDomainOriginInfoArgs']:
        """
        Details of the origin.
        """
        return pulumi.get(self, "origin_info")

    @origin_info.setter
    def origin_info(self, value: pulumi.Input['AccelerationDomainOriginInfoArgs']):
        pulumi.set(self, "origin_info", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        ID of the site related with the accelerated domain name.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="httpOriginPort")
    def http_origin_port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        """
        return pulumi.get(self, "http_origin_port")

    @http_origin_port.setter
    def http_origin_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_origin_port", value)

    @property
    @pulumi.getter(name="httpsOriginPort")
    def https_origin_port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        """
        return pulumi.get(self, "https_origin_port")

    @https_origin_port.setter
    def https_origin_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_origin_port", value)

    @property
    @pulumi.getter(name="ipv6Status")
    def ipv6_status(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        """
        return pulumi.get(self, "ipv6_status")

    @ipv6_status.setter
    def ipv6_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_status", value)

    @property
    @pulumi.getter(name="originProtocol")
    def origin_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        """
        return pulumi.get(self, "origin_protocol")

    @origin_protocol.setter
    def origin_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_protocol", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _AccelerationDomainState:
    def __init__(__self__, *,
                 cname: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 http_origin_port: Optional[pulumi.Input[int]] = None,
                 https_origin_port: Optional[pulumi.Input[int]] = None,
                 ipv6_status: Optional[pulumi.Input[str]] = None,
                 origin_info: Optional[pulumi.Input['AccelerationDomainOriginInfoArgs']] = None,
                 origin_protocol: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccelerationDomain resources.
        :param pulumi.Input[str] cname: CNAME address.
        :param pulumi.Input[str] domain_name: Accelerated domain name.
        :param pulumi.Input[int] http_origin_port: HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        :param pulumi.Input[int] https_origin_port: HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        :param pulumi.Input[str] ipv6_status: IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        :param pulumi.Input['AccelerationDomainOriginInfoArgs'] origin_info: Details of the origin.
        :param pulumi.Input[str] origin_protocol: Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        :param pulumi.Input[str] status: Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        :param pulumi.Input[str] zone_id: ID of the site related with the accelerated domain name.
        """
        if cname is not None:
            pulumi.set(__self__, "cname", cname)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if http_origin_port is not None:
            pulumi.set(__self__, "http_origin_port", http_origin_port)
        if https_origin_port is not None:
            pulumi.set(__self__, "https_origin_port", https_origin_port)
        if ipv6_status is not None:
            pulumi.set(__self__, "ipv6_status", ipv6_status)
        if origin_info is not None:
            pulumi.set(__self__, "origin_info", origin_info)
        if origin_protocol is not None:
            pulumi.set(__self__, "origin_protocol", origin_protocol)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def cname(self) -> Optional[pulumi.Input[str]]:
        """
        CNAME address.
        """
        return pulumi.get(self, "cname")

    @cname.setter
    def cname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cname", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        Accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="httpOriginPort")
    def http_origin_port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        """
        return pulumi.get(self, "http_origin_port")

    @http_origin_port.setter
    def http_origin_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_origin_port", value)

    @property
    @pulumi.getter(name="httpsOriginPort")
    def https_origin_port(self) -> Optional[pulumi.Input[int]]:
        """
        HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        """
        return pulumi.get(self, "https_origin_port")

    @https_origin_port.setter
    def https_origin_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_origin_port", value)

    @property
    @pulumi.getter(name="ipv6Status")
    def ipv6_status(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        """
        return pulumi.get(self, "ipv6_status")

    @ipv6_status.setter
    def ipv6_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_status", value)

    @property
    @pulumi.getter(name="originInfo")
    def origin_info(self) -> Optional[pulumi.Input['AccelerationDomainOriginInfoArgs']]:
        """
        Details of the origin.
        """
        return pulumi.get(self, "origin_info")

    @origin_info.setter
    def origin_info(self, value: Optional[pulumi.Input['AccelerationDomainOriginInfoArgs']]):
        pulumi.set(self, "origin_info", value)

    @property
    @pulumi.getter(name="originProtocol")
    def origin_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        """
        return pulumi.get(self, "origin_protocol")

    @origin_protocol.setter
    def origin_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "origin_protocol", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the site related with the accelerated domain name.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class AccelerationDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 http_origin_port: Optional[pulumi.Input[int]] = None,
                 https_origin_port: Optional[pulumi.Input[int]] = None,
                 ipv6_status: Optional[pulumi.Input[str]] = None,
                 origin_info: Optional[pulumi.Input[pulumi.InputType['AccelerationDomainOriginInfoArgs']]] = None,
                 origin_protocol: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a teo acceleration_domain

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        acceleration_domain = tencentcloud.teo.AccelerationDomain("accelerationDomain",
            domain_name="aaa.makn.cn",
            origin_info=tencentcloud.teo.AccelerationDomainOriginInfoArgs(
                origin="150.109.8.1",
                origin_type="IP_DOMAIN",
            ),
            zone_id="zone-2o0i41pv2h8c")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo acceleration_domain can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/accelerationDomain:AccelerationDomain acceleration_domain acceleration_domain_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: Accelerated domain name.
        :param pulumi.Input[int] http_origin_port: HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        :param pulumi.Input[int] https_origin_port: HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        :param pulumi.Input[str] ipv6_status: IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        :param pulumi.Input[pulumi.InputType['AccelerationDomainOriginInfoArgs']] origin_info: Details of the origin.
        :param pulumi.Input[str] origin_protocol: Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        :param pulumi.Input[str] status: Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        :param pulumi.Input[str] zone_id: ID of the site related with the accelerated domain name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccelerationDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a teo acceleration_domain

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        acceleration_domain = tencentcloud.teo.AccelerationDomain("accelerationDomain",
            domain_name="aaa.makn.cn",
            origin_info=tencentcloud.teo.AccelerationDomainOriginInfoArgs(
                origin="150.109.8.1",
                origin_type="IP_DOMAIN",
            ),
            zone_id="zone-2o0i41pv2h8c")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo acceleration_domain can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/accelerationDomain:AccelerationDomain acceleration_domain acceleration_domain_id
        ```

        :param str resource_name: The name of the resource.
        :param AccelerationDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccelerationDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 http_origin_port: Optional[pulumi.Input[int]] = None,
                 https_origin_port: Optional[pulumi.Input[int]] = None,
                 ipv6_status: Optional[pulumi.Input[str]] = None,
                 origin_info: Optional[pulumi.Input[pulumi.InputType['AccelerationDomainOriginInfoArgs']]] = None,
                 origin_protocol: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccelerationDomainArgs.__new__(AccelerationDomainArgs)

            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["http_origin_port"] = http_origin_port
            __props__.__dict__["https_origin_port"] = https_origin_port
            __props__.__dict__["ipv6_status"] = ipv6_status
            if origin_info is None and not opts.urn:
                raise TypeError("Missing required property 'origin_info'")
            __props__.__dict__["origin_info"] = origin_info
            __props__.__dict__["origin_protocol"] = origin_protocol
            __props__.__dict__["status"] = status
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["cname"] = None
        super(AccelerationDomain, __self__).__init__(
            'tencentcloud:Teo/accelerationDomain:AccelerationDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cname: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            http_origin_port: Optional[pulumi.Input[int]] = None,
            https_origin_port: Optional[pulumi.Input[int]] = None,
            ipv6_status: Optional[pulumi.Input[str]] = None,
            origin_info: Optional[pulumi.Input[pulumi.InputType['AccelerationDomainOriginInfoArgs']]] = None,
            origin_protocol: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'AccelerationDomain':
        """
        Get an existing AccelerationDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cname: CNAME address.
        :param pulumi.Input[str] domain_name: Accelerated domain name.
        :param pulumi.Input[int] http_origin_port: HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        :param pulumi.Input[int] https_origin_port: HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        :param pulumi.Input[str] ipv6_status: IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        :param pulumi.Input[pulumi.InputType['AccelerationDomainOriginInfoArgs']] origin_info: Details of the origin.
        :param pulumi.Input[str] origin_protocol: Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        :param pulumi.Input[str] status: Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        :param pulumi.Input[str] zone_id: ID of the site related with the accelerated domain name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccelerationDomainState.__new__(_AccelerationDomainState)

        __props__.__dict__["cname"] = cname
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["http_origin_port"] = http_origin_port
        __props__.__dict__["https_origin_port"] = https_origin_port
        __props__.__dict__["ipv6_status"] = ipv6_status
        __props__.__dict__["origin_info"] = origin_info
        __props__.__dict__["origin_protocol"] = origin_protocol
        __props__.__dict__["status"] = status
        __props__.__dict__["zone_id"] = zone_id
        return AccelerationDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cname(self) -> pulumi.Output[str]:
        """
        CNAME address.
        """
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Accelerated domain name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="httpOriginPort")
    def http_origin_port(self) -> pulumi.Output[int]:
        """
        HTTP back-to-origin port, the value is 1-65535, effective when OriginProtocol=FOLLOW/HTTP, if not filled in, the default value is 80.
        """
        return pulumi.get(self, "http_origin_port")

    @property
    @pulumi.getter(name="httpsOriginPort")
    def https_origin_port(self) -> pulumi.Output[int]:
        """
        HTTPS back-to-origin port. The value range is 1-65535. It takes effect when OriginProtocol=FOLLOW/HTTPS. If it is not filled in, the default value is 443.
        """
        return pulumi.get(self, "https_origin_port")

    @property
    @pulumi.getter(name="ipv6Status")
    def ipv6_status(self) -> pulumi.Output[str]:
        """
        IPv6 status, the value is: `follow`: follow the site IPv6 configuration; `on`: on; `off`: off. If not filled in, the default is: `follow`.
        """
        return pulumi.get(self, "ipv6_status")

    @property
    @pulumi.getter(name="originInfo")
    def origin_info(self) -> pulumi.Output['outputs.AccelerationDomainOriginInfo']:
        """
        Details of the origin.
        """
        return pulumi.get(self, "origin_info")

    @property
    @pulumi.getter(name="originProtocol")
    def origin_protocol(self) -> pulumi.Output[str]:
        """
        Origin return protocol, possible values are: `FOLLOW`: protocol follow; `HTTP`: HTTP protocol back to source; `HTTPS`: HTTPS protocol back to source. If not filled in, the default is: `FOLLOW`.
        """
        return pulumi.get(self, "origin_protocol")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Accelerated domain name status, the values are: `online`: enabled; `offline`: disabled.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        ID of the site related with the accelerated domain name.
        """
        return pulumi.get(self, "zone_id")

