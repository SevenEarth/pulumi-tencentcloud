# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AntiInfoLeakStrategyArgs',
    'ClbDomainLoadBalancerSetArgs',
    'CustomRuleStrategyArgs',
    'CustomWhiteRuleStrategyArgs',
    'IpAccessControlItemArgs',
    'SaasDomainPortArgs',
    'GetWafInfosParamArgs',
]

@pulumi.input_type
class AntiInfoLeakStrategyArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 field: pulumi.Input[str]):
        """
        :param pulumi.Input[str] content: Matching Content. If field is returncode support: 400, 403, 404, 4xx, 500, 501, 502, 504, 5xx; If field is information support: idcard, phone, bankcard; If field is keywords users input matching content themselves.
        :param pulumi.Input[str] field: Matching Fields. support: returncode, keywords, information.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        Matching Content. If field is returncode support: 400, 403, 404, 4xx, 500, 501, 502, 504, 5xx; If field is information support: idcard, phone, bankcard; If field is keywords users input matching content themselves.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        Matching Fields. support: returncode, keywords, information.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)


@pulumi.input_type
class ClbDomainLoadBalancerSetArgs:
    def __init__(__self__, *,
                 listener_id: pulumi.Input[str],
                 listener_name: pulumi.Input[str],
                 load_balancer_id: pulumi.Input[str],
                 load_balancer_name: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 region: pulumi.Input[str],
                 vip: pulumi.Input[str],
                 vport: pulumi.Input[int],
                 zone: pulumi.Input[str],
                 load_balancer_type: Optional[pulumi.Input[str]] = None,
                 numerical_vpc_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] listener_id: Unique ID of listener in LB.
        :param pulumi.Input[str] listener_name: Listener name.
        :param pulumi.Input[str] load_balancer_id: LoadBalancer unique ID.
        :param pulumi.Input[str] load_balancer_name: LoadBalancer name.
        :param pulumi.Input[str] protocol: Protocol of listener, http or https.
        :param pulumi.Input[str] region: LoadBalancer region.
        :param pulumi.Input[str] vip: LoadBalancer IP.
        :param pulumi.Input[int] vport: LoadBalancer port.
        :param pulumi.Input[str] zone: LoadBalancer zone.
        :param pulumi.Input[str] load_balancer_type: Network type for load balancer.
        :param pulumi.Input[int] numerical_vpc_id: VPCID for load balancer, public network is -1, and internal network is filled in according to actual conditions.
        """
        pulumi.set(__self__, "listener_id", listener_id)
        pulumi.set(__self__, "listener_name", listener_name)
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "vip", vip)
        pulumi.set(__self__, "vport", vport)
        pulumi.set(__self__, "zone", zone)
        if load_balancer_type is not None:
            pulumi.set(__self__, "load_balancer_type", load_balancer_type)
        if numerical_vpc_id is not None:
            pulumi.set(__self__, "numerical_vpc_id", numerical_vpc_id)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Input[str]:
        """
        Unique ID of listener in LB.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> pulumi.Input[str]:
        """
        Listener name.
        """
        return pulumi.get(self, "listener_name")

    @listener_name.setter
    def listener_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "listener_name", value)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> pulumi.Input[str]:
        """
        LoadBalancer unique ID.
        """
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Input[str]:
        """
        LoadBalancer name.
        """
        return pulumi.get(self, "load_balancer_name")

    @load_balancer_name.setter
    def load_balancer_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_name", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Protocol of listener, http or https.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        LoadBalancer region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def vip(self) -> pulumi.Input[str]:
        """
        LoadBalancer IP.
        """
        return pulumi.get(self, "vip")

    @vip.setter
    def vip(self, value: pulumi.Input[str]):
        pulumi.set(self, "vip", value)

    @property
    @pulumi.getter
    def vport(self) -> pulumi.Input[int]:
        """
        LoadBalancer port.
        """
        return pulumi.get(self, "vport")

    @vport.setter
    def vport(self, value: pulumi.Input[int]):
        pulumi.set(self, "vport", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        LoadBalancer zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="loadBalancerType")
    def load_balancer_type(self) -> Optional[pulumi.Input[str]]:
        """
        Network type for load balancer.
        """
        return pulumi.get(self, "load_balancer_type")

    @load_balancer_type.setter
    def load_balancer_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer_type", value)

    @property
    @pulumi.getter(name="numericalVpcId")
    def numerical_vpc_id(self) -> Optional[pulumi.Input[int]]:
        """
        VPCID for load balancer, public network is -1, and internal network is filled in according to actual conditions.
        """
        return pulumi.get(self, "numerical_vpc_id")

    @numerical_vpc_id.setter
    def numerical_vpc_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "numerical_vpc_id", value)


@pulumi.input_type
class CustomRuleStrategyArgs:
    def __init__(__self__, *,
                 arg: pulumi.Input[str],
                 compare_func: pulumi.Input[str],
                 content: pulumi.Input[str],
                 field: pulumi.Input[str]):
        """
        :param pulumi.Input[str] arg: Matching parameters.
        :param pulumi.Input[str] compare_func: Logical symbol.
        :param pulumi.Input[str] content: Matching Content.
        :param pulumi.Input[str] field: Matching Fields.
        """
        pulumi.set(__self__, "arg", arg)
        pulumi.set(__self__, "compare_func", compare_func)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def arg(self) -> pulumi.Input[str]:
        """
        Matching parameters.
        """
        return pulumi.get(self, "arg")

    @arg.setter
    def arg(self, value: pulumi.Input[str]):
        pulumi.set(self, "arg", value)

    @property
    @pulumi.getter(name="compareFunc")
    def compare_func(self) -> pulumi.Input[str]:
        """
        Logical symbol.
        """
        return pulumi.get(self, "compare_func")

    @compare_func.setter
    def compare_func(self, value: pulumi.Input[str]):
        pulumi.set(self, "compare_func", value)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        Matching Content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        Matching Fields.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)


@pulumi.input_type
class CustomWhiteRuleStrategyArgs:
    def __init__(__self__, *,
                 arg: pulumi.Input[str],
                 compare_func: pulumi.Input[str],
                 content: pulumi.Input[str],
                 field: pulumi.Input[str]):
        """
        :param pulumi.Input[str] arg: Matching parameters.
        :param pulumi.Input[str] compare_func: Logical symbol.
        :param pulumi.Input[str] content: Matching Content.
        :param pulumi.Input[str] field: Matching Fields.
        """
        pulumi.set(__self__, "arg", arg)
        pulumi.set(__self__, "compare_func", compare_func)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def arg(self) -> pulumi.Input[str]:
        """
        Matching parameters.
        """
        return pulumi.get(self, "arg")

    @arg.setter
    def arg(self, value: pulumi.Input[str]):
        pulumi.set(self, "arg", value)

    @property
    @pulumi.getter(name="compareFunc")
    def compare_func(self) -> pulumi.Input[str]:
        """
        Logical symbol.
        """
        return pulumi.get(self, "compare_func")

    @compare_func.setter
    def compare_func(self, value: pulumi.Input[str]):
        pulumi.set(self, "compare_func", value)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        Matching Content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        Matching Fields.
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)


@pulumi.input_type
class IpAccessControlItemArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[int],
                 ip: pulumi.Input[str],
                 note: pulumi.Input[str],
                 valid_ts: pulumi.Input[int],
                 id: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 valid_status: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] action: Action value 40 is whitelist, 42 is blacklist.
        :param pulumi.Input[str] ip: IP address.
        :param pulumi.Input[str] note: Note info.
        :param pulumi.Input[int] valid_ts: Effective date, with a second level timestamp value. For example, 1680570420 represents 2023-04-04 09:07:00; 2019571199 means permanently effective.
        :param pulumi.Input[str] id: ID of the resource.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "note", note)
        pulumi.set(__self__, "valid_ts", valid_ts)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if valid_status is not None:
            pulumi.set(__self__, "valid_status", valid_status)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[int]:
        """
        Action value 40 is whitelist, 42 is blacklist.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[int]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def note(self) -> pulumi.Input[str]:
        """
        Note info.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: pulumi.Input[str]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter(name="validTs")
    def valid_ts(self) -> pulumi.Input[int]:
        """
        Effective date, with a second level timestamp value. For example, 1680570420 represents 2023-04-04 09:07:00; 2019571199 means permanently effective.
        """
        return pulumi.get(self, "valid_ts")

    @valid_ts.setter
    def valid_ts(self, value: pulumi.Input[int]):
        pulumi.set(self, "valid_ts", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="validStatus")
    def valid_status(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "valid_status")

    @valid_status.setter
    def valid_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "valid_status", value)


@pulumi.input_type
class SaasDomainPortArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 upstream_port: pulumi.Input[str],
                 upstream_protocol: pulumi.Input[str],
                 nginx_server_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] port: Listening port.
        :param pulumi.Input[str] protocol: The listening protocol of listening port.
        :param pulumi.Input[str] upstream_port: The upstream port for listening port.
        :param pulumi.Input[str] upstream_protocol: The upstream protocol for listening port.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "upstream_port", upstream_port)
        pulumi.set(__self__, "upstream_protocol", upstream_protocol)
        if nginx_server_id is not None:
            pulumi.set(__self__, "nginx_server_id", nginx_server_id)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[str]:
        """
        Listening port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[str]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The listening protocol of listening port.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="upstreamPort")
    def upstream_port(self) -> pulumi.Input[str]:
        """
        The upstream port for listening port.
        """
        return pulumi.get(self, "upstream_port")

    @upstream_port.setter
    def upstream_port(self, value: pulumi.Input[str]):
        pulumi.set(self, "upstream_port", value)

    @property
    @pulumi.getter(name="upstreamProtocol")
    def upstream_protocol(self) -> pulumi.Input[str]:
        """
        The upstream protocol for listening port.
        """
        return pulumi.get(self, "upstream_protocol")

    @upstream_protocol.setter
    def upstream_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "upstream_protocol", value)

    @property
    @pulumi.getter(name="nginxServerId")
    def nginx_server_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nginx_server_id")

    @nginx_server_id.setter
    def nginx_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nginx_server_id", value)


@pulumi.input_type
class GetWafInfosParamArgs:
    def __init__(__self__, *,
                 load_balancer_id: str,
                 domain_id: Optional[str] = None,
                 listener_id: Optional[str] = None):
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)

    @property
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> str:
        return pulumi.get(self, "load_balancer_id")

    @load_balancer_id.setter
    def load_balancer_id(self, value: str):
        pulumi.set(self, "load_balancer_id", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[str]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[str]:
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[str]):
        pulumi.set(self, "listener_id", value)


