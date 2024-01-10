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

__all__ = ['CngwGatewayArgs', 'CngwGateway']

@pulumi.input_type
class CngwGatewayArgs:
    def __init__(__self__, *,
                 gateway_version: pulumi.Input[str],
                 node_config: pulumi.Input['CngwGatewayNodeConfigArgs'],
                 type: pulumi.Input[str],
                 vpc_config: pulumi.Input['CngwGatewayVpcConfigArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 enable_cls: Optional[pulumi.Input[bool]] = None,
                 engine_region: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 ingress_class_name: Optional[pulumi.Input[str]] = None,
                 internet_config: Optional[pulumi.Input['CngwGatewayInternetConfigArgs']] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 trade_type: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a CngwGateway resource.
        :param pulumi.Input[str] gateway_version: gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        :param pulumi.Input['CngwGatewayNodeConfigArgs'] node_config: gateway node configration.
        :param pulumi.Input[str] type: gateway type,currently only supports kong.
        :param pulumi.Input['CngwGatewayVpcConfigArgs'] vpc_config: vpc information.
        :param pulumi.Input[str] description: description information, up to 120 characters.
        :param pulumi.Input[bool] enable_cls: whether to enable CLS log. Default value: fasle.
        :param pulumi.Input[str] engine_region: engine region of gateway.
        :param pulumi.Input[str] feature_version: product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        :param pulumi.Input[str] ingress_class_name: ingress class name.
        :param pulumi.Input['CngwGatewayInternetConfigArgs'] internet_config: internet configration.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network outbound traffic bandwidth,[1,2048]Mbps.
        :param pulumi.Input[str] name: gateway name, supports up to 60 characters.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] trade_type: trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        """
        pulumi.set(__self__, "gateway_version", gateway_version)
        pulumi.set(__self__, "node_config", node_config)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vpc_config", vpc_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_cls is not None:
            pulumi.set(__self__, "enable_cls", enable_cls)
        if engine_region is not None:
            pulumi.set(__self__, "engine_region", engine_region)
        if feature_version is not None:
            pulumi.set(__self__, "feature_version", feature_version)
        if ingress_class_name is not None:
            pulumi.set(__self__, "ingress_class_name", ingress_class_name)
        if internet_config is not None:
            pulumi.set(__self__, "internet_config", internet_config)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trade_type is not None:
            pulumi.set(__self__, "trade_type", trade_type)

    @property
    @pulumi.getter(name="gatewayVersion")
    def gateway_version(self) -> pulumi.Input[str]:
        """
        gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        """
        return pulumi.get(self, "gateway_version")

    @gateway_version.setter
    def gateway_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_version", value)

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Input['CngwGatewayNodeConfigArgs']:
        """
        gateway node configration.
        """
        return pulumi.get(self, "node_config")

    @node_config.setter
    def node_config(self, value: pulumi.Input['CngwGatewayNodeConfigArgs']):
        pulumi.set(self, "node_config", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        gateway type,currently only supports kong.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> pulumi.Input['CngwGatewayVpcConfigArgs']:
        """
        vpc information.
        """
        return pulumi.get(self, "vpc_config")

    @vpc_config.setter
    def vpc_config(self, value: pulumi.Input['CngwGatewayVpcConfigArgs']):
        pulumi.set(self, "vpc_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description information, up to 120 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableCls")
    def enable_cls(self) -> Optional[pulumi.Input[bool]]:
        """
        whether to enable CLS log. Default value: fasle.
        """
        return pulumi.get(self, "enable_cls")

    @enable_cls.setter
    def enable_cls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_cls", value)

    @property
    @pulumi.getter(name="engineRegion")
    def engine_region(self) -> Optional[pulumi.Input[str]]:
        """
        engine region of gateway.
        """
        return pulumi.get(self, "engine_region")

    @engine_region.setter
    def engine_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_region", value)

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> Optional[pulumi.Input[str]]:
        """
        product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        """
        return pulumi.get(self, "feature_version")

    @feature_version.setter
    def feature_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_version", value)

    @property
    @pulumi.getter(name="ingressClassName")
    def ingress_class_name(self) -> Optional[pulumi.Input[str]]:
        """
        ingress class name.
        """
        return pulumi.get(self, "ingress_class_name")

    @ingress_class_name.setter
    def ingress_class_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ingress_class_name", value)

    @property
    @pulumi.getter(name="internetConfig")
    def internet_config(self) -> Optional[pulumi.Input['CngwGatewayInternetConfigArgs']]:
        """
        internet configration.
        """
        return pulumi.get(self, "internet_config")

    @internet_config.setter
    def internet_config(self, value: Optional[pulumi.Input['CngwGatewayInternetConfigArgs']]):
        pulumi.set(self, "internet_config", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        public network outbound traffic bandwidth,[1,2048]Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        gateway name, supports up to 60 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tradeType")
    def trade_type(self) -> Optional[pulumi.Input[int]]:
        """
        trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        """
        return pulumi.get(self, "trade_type")

    @trade_type.setter
    def trade_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trade_type", value)


@pulumi.input_type
class _CngwGatewayState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_cls: Optional[pulumi.Input[bool]] = None,
                 engine_region: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 gateway_version: Optional[pulumi.Input[str]] = None,
                 ingress_class_name: Optional[pulumi.Input[str]] = None,
                 instance_ports: Optional[pulumi.Input[Sequence[pulumi.Input['CngwGatewayInstancePortArgs']]]] = None,
                 internet_config: Optional[pulumi.Input['CngwGatewayInternetConfigArgs']] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input['CngwGatewayNodeConfigArgs']] = None,
                 public_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 trade_type: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_config: Optional[pulumi.Input['CngwGatewayVpcConfigArgs']] = None):
        """
        Input properties used for looking up and filtering CngwGateway resources.
        :param pulumi.Input[str] description: description information, up to 120 characters.
        :param pulumi.Input[bool] enable_cls: whether to enable CLS log. Default value: fasle.
        :param pulumi.Input[str] engine_region: engine region of gateway.
        :param pulumi.Input[str] feature_version: product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        :param pulumi.Input[str] gateway_version: gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        :param pulumi.Input[str] ingress_class_name: ingress class name.
        :param pulumi.Input[Sequence[pulumi.Input['CngwGatewayInstancePortArgs']]] instance_ports: Port information that the instance listens to.
        :param pulumi.Input['CngwGatewayInternetConfigArgs'] internet_config: internet configration.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network outbound traffic bandwidth,[1,2048]Mbps.
        :param pulumi.Input[str] name: gateway name, supports up to 60 characters.
        :param pulumi.Input['CngwGatewayNodeConfigArgs'] node_config: gateway node configration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addresses: Public IP address list.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] trade_type: trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        :param pulumi.Input[str] type: gateway type,currently only supports kong.
        :param pulumi.Input['CngwGatewayVpcConfigArgs'] vpc_config: vpc information.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_cls is not None:
            pulumi.set(__self__, "enable_cls", enable_cls)
        if engine_region is not None:
            pulumi.set(__self__, "engine_region", engine_region)
        if feature_version is not None:
            pulumi.set(__self__, "feature_version", feature_version)
        if gateway_version is not None:
            pulumi.set(__self__, "gateway_version", gateway_version)
        if ingress_class_name is not None:
            pulumi.set(__self__, "ingress_class_name", ingress_class_name)
        if instance_ports is not None:
            pulumi.set(__self__, "instance_ports", instance_ports)
        if internet_config is not None:
            pulumi.set(__self__, "internet_config", internet_config)
        if internet_max_bandwidth_out is not None:
            pulumi.set(__self__, "internet_max_bandwidth_out", internet_max_bandwidth_out)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_config is not None:
            pulumi.set(__self__, "node_config", node_config)
        if public_ip_addresses is not None:
            pulumi.set(__self__, "public_ip_addresses", public_ip_addresses)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trade_type is not None:
            pulumi.set(__self__, "trade_type", trade_type)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpc_config is not None:
            pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description information, up to 120 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableCls")
    def enable_cls(self) -> Optional[pulumi.Input[bool]]:
        """
        whether to enable CLS log. Default value: fasle.
        """
        return pulumi.get(self, "enable_cls")

    @enable_cls.setter
    def enable_cls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_cls", value)

    @property
    @pulumi.getter(name="engineRegion")
    def engine_region(self) -> Optional[pulumi.Input[str]]:
        """
        engine region of gateway.
        """
        return pulumi.get(self, "engine_region")

    @engine_region.setter
    def engine_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_region", value)

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> Optional[pulumi.Input[str]]:
        """
        product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        """
        return pulumi.get(self, "feature_version")

    @feature_version.setter
    def feature_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_version", value)

    @property
    @pulumi.getter(name="gatewayVersion")
    def gateway_version(self) -> Optional[pulumi.Input[str]]:
        """
        gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        """
        return pulumi.get(self, "gateway_version")

    @gateway_version.setter
    def gateway_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_version", value)

    @property
    @pulumi.getter(name="ingressClassName")
    def ingress_class_name(self) -> Optional[pulumi.Input[str]]:
        """
        ingress class name.
        """
        return pulumi.get(self, "ingress_class_name")

    @ingress_class_name.setter
    def ingress_class_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ingress_class_name", value)

    @property
    @pulumi.getter(name="instancePorts")
    def instance_ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CngwGatewayInstancePortArgs']]]]:
        """
        Port information that the instance listens to.
        """
        return pulumi.get(self, "instance_ports")

    @instance_ports.setter
    def instance_ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CngwGatewayInstancePortArgs']]]]):
        pulumi.set(self, "instance_ports", value)

    @property
    @pulumi.getter(name="internetConfig")
    def internet_config(self) -> Optional[pulumi.Input['CngwGatewayInternetConfigArgs']]:
        """
        internet configration.
        """
        return pulumi.get(self, "internet_config")

    @internet_config.setter
    def internet_config(self, value: Optional[pulumi.Input['CngwGatewayInternetConfigArgs']]):
        pulumi.set(self, "internet_config", value)

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> Optional[pulumi.Input[int]]:
        """
        public network outbound traffic bandwidth,[1,2048]Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @internet_max_bandwidth_out.setter
    def internet_max_bandwidth_out(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth_out", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        gateway name, supports up to 60 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> Optional[pulumi.Input['CngwGatewayNodeConfigArgs']]:
        """
        gateway node configration.
        """
        return pulumi.get(self, "node_config")

    @node_config.setter
    def node_config(self, value: Optional[pulumi.Input['CngwGatewayNodeConfigArgs']]):
        pulumi.set(self, "node_config", value)

    @property
    @pulumi.getter(name="publicIpAddresses")
    def public_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Public IP address list.
        """
        return pulumi.get(self, "public_ip_addresses")

    @public_ip_addresses.setter
    def public_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "public_ip_addresses", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tradeType")
    def trade_type(self) -> Optional[pulumi.Input[int]]:
        """
        trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        """
        return pulumi.get(self, "trade_type")

    @trade_type.setter
    def trade_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trade_type", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        gateway type,currently only supports kong.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> Optional[pulumi.Input['CngwGatewayVpcConfigArgs']]:
        """
        vpc information.
        """
        return pulumi.get(self, "vpc_config")

    @vpc_config.setter
    def vpc_config(self, value: Optional[pulumi.Input['CngwGatewayVpcConfigArgs']]):
        pulumi.set(self, "vpc_config", value)


class CngwGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_cls: Optional[pulumi.Input[bool]] = None,
                 engine_region: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 gateway_version: Optional[pulumi.Input[str]] = None,
                 ingress_class_name: Optional[pulumi.Input[str]] = None,
                 internet_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayInternetConfigArgs']]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayNodeConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 trade_type: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayVpcConfigArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a tse cngw_gateway

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-4"
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=availability_zone,
            cidr_block="10.0.1.0/24")
        cngw_gateway = tencentcloud.tse.CngwGateway("cngwGateway",
            description="terraform test1",
            enable_cls=True,
            engine_region="ap-guangzhou",
            feature_version="STANDARD",
            gateway_version="2.5.1",
            ingress_class_name="tse-nginx-ingress",
            internet_max_bandwidth_out=0,
            trade_type=0,
            type="kong",
            node_config=tencentcloud.tse.CngwGatewayNodeConfigArgs(
                number=2,
                specification="1c2g",
            ),
            vpc_config=tencentcloud.tse.CngwGatewayVpcConfigArgs(
                subnet_id=subnet.id,
                vpc_id=vpc.id,
            ),
            tags={
                "createdBy": "terraform",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description information, up to 120 characters.
        :param pulumi.Input[bool] enable_cls: whether to enable CLS log. Default value: fasle.
        :param pulumi.Input[str] engine_region: engine region of gateway.
        :param pulumi.Input[str] feature_version: product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        :param pulumi.Input[str] gateway_version: gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        :param pulumi.Input[str] ingress_class_name: ingress class name.
        :param pulumi.Input[pulumi.InputType['CngwGatewayInternetConfigArgs']] internet_config: internet configration.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network outbound traffic bandwidth,[1,2048]Mbps.
        :param pulumi.Input[str] name: gateway name, supports up to 60 characters.
        :param pulumi.Input[pulumi.InputType['CngwGatewayNodeConfigArgs']] node_config: gateway node configration.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] trade_type: trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        :param pulumi.Input[str] type: gateway type,currently only supports kong.
        :param pulumi.Input[pulumi.InputType['CngwGatewayVpcConfigArgs']] vpc_config: vpc information.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CngwGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tse cngw_gateway

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-4"
        vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
        subnet = tencentcloud.subnet.Instance("subnet",
            vpc_id=vpc.id,
            availability_zone=availability_zone,
            cidr_block="10.0.1.0/24")
        cngw_gateway = tencentcloud.tse.CngwGateway("cngwGateway",
            description="terraform test1",
            enable_cls=True,
            engine_region="ap-guangzhou",
            feature_version="STANDARD",
            gateway_version="2.5.1",
            ingress_class_name="tse-nginx-ingress",
            internet_max_bandwidth_out=0,
            trade_type=0,
            type="kong",
            node_config=tencentcloud.tse.CngwGatewayNodeConfigArgs(
                number=2,
                specification="1c2g",
            ),
            vpc_config=tencentcloud.tse.CngwGatewayVpcConfigArgs(
                subnet_id=subnet.id,
                vpc_id=vpc.id,
            ),
            tags={
                "createdBy": "terraform",
            })
        ```

        :param str resource_name: The name of the resource.
        :param CngwGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CngwGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_cls: Optional[pulumi.Input[bool]] = None,
                 engine_region: Optional[pulumi.Input[str]] = None,
                 feature_version: Optional[pulumi.Input[str]] = None,
                 gateway_version: Optional[pulumi.Input[str]] = None,
                 ingress_class_name: Optional[pulumi.Input[str]] = None,
                 internet_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayInternetConfigArgs']]] = None,
                 internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayNodeConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 trade_type: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayVpcConfigArgs']]] = None,
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
            __props__ = CngwGatewayArgs.__new__(CngwGatewayArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["enable_cls"] = enable_cls
            __props__.__dict__["engine_region"] = engine_region
            __props__.__dict__["feature_version"] = feature_version
            if gateway_version is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_version'")
            __props__.__dict__["gateway_version"] = gateway_version
            __props__.__dict__["ingress_class_name"] = ingress_class_name
            __props__.__dict__["internet_config"] = internet_config
            __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
            __props__.__dict__["name"] = name
            if node_config is None and not opts.urn:
                raise TypeError("Missing required property 'node_config'")
            __props__.__dict__["node_config"] = node_config
            __props__.__dict__["tags"] = tags
            __props__.__dict__["trade_type"] = trade_type
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if vpc_config is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_config'")
            __props__.__dict__["vpc_config"] = vpc_config
            __props__.__dict__["instance_ports"] = None
            __props__.__dict__["public_ip_addresses"] = None
        super(CngwGateway, __self__).__init__(
            'tencentcloud:Tse/cngwGateway:CngwGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_cls: Optional[pulumi.Input[bool]] = None,
            engine_region: Optional[pulumi.Input[str]] = None,
            feature_version: Optional[pulumi.Input[str]] = None,
            gateway_version: Optional[pulumi.Input[str]] = None,
            ingress_class_name: Optional[pulumi.Input[str]] = None,
            instance_ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CngwGatewayInstancePortArgs']]]]] = None,
            internet_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayInternetConfigArgs']]] = None,
            internet_max_bandwidth_out: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayNodeConfigArgs']]] = None,
            public_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            trade_type: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vpc_config: Optional[pulumi.Input[pulumi.InputType['CngwGatewayVpcConfigArgs']]] = None) -> 'CngwGateway':
        """
        Get an existing CngwGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description information, up to 120 characters.
        :param pulumi.Input[bool] enable_cls: whether to enable CLS log. Default value: fasle.
        :param pulumi.Input[str] engine_region: engine region of gateway.
        :param pulumi.Input[str] feature_version: product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        :param pulumi.Input[str] gateway_version: gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        :param pulumi.Input[str] ingress_class_name: ingress class name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CngwGatewayInstancePortArgs']]]] instance_ports: Port information that the instance listens to.
        :param pulumi.Input[pulumi.InputType['CngwGatewayInternetConfigArgs']] internet_config: internet configration.
        :param pulumi.Input[int] internet_max_bandwidth_out: public network outbound traffic bandwidth,[1,2048]Mbps.
        :param pulumi.Input[str] name: gateway name, supports up to 60 characters.
        :param pulumi.Input[pulumi.InputType['CngwGatewayNodeConfigArgs']] node_config: gateway node configration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] public_ip_addresses: Public IP address list.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] trade_type: trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        :param pulumi.Input[str] type: gateway type,currently only supports kong.
        :param pulumi.Input[pulumi.InputType['CngwGatewayVpcConfigArgs']] vpc_config: vpc information.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CngwGatewayState.__new__(_CngwGatewayState)

        __props__.__dict__["description"] = description
        __props__.__dict__["enable_cls"] = enable_cls
        __props__.__dict__["engine_region"] = engine_region
        __props__.__dict__["feature_version"] = feature_version
        __props__.__dict__["gateway_version"] = gateway_version
        __props__.__dict__["ingress_class_name"] = ingress_class_name
        __props__.__dict__["instance_ports"] = instance_ports
        __props__.__dict__["internet_config"] = internet_config
        __props__.__dict__["internet_max_bandwidth_out"] = internet_max_bandwidth_out
        __props__.__dict__["name"] = name
        __props__.__dict__["node_config"] = node_config
        __props__.__dict__["public_ip_addresses"] = public_ip_addresses
        __props__.__dict__["tags"] = tags
        __props__.__dict__["trade_type"] = trade_type
        __props__.__dict__["type"] = type
        __props__.__dict__["vpc_config"] = vpc_config
        return CngwGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        description information, up to 120 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableCls")
    def enable_cls(self) -> pulumi.Output[Optional[bool]]:
        """
        whether to enable CLS log. Default value: fasle.
        """
        return pulumi.get(self, "enable_cls")

    @property
    @pulumi.getter(name="engineRegion")
    def engine_region(self) -> pulumi.Output[str]:
        """
        engine region of gateway.
        """
        return pulumi.get(self, "engine_region")

    @property
    @pulumi.getter(name="featureVersion")
    def feature_version(self) -> pulumi.Output[str]:
        """
        product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        """
        return pulumi.get(self, "feature_version")

    @property
    @pulumi.getter(name="gatewayVersion")
    def gateway_version(self) -> pulumi.Output[str]:
        """
        gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        """
        return pulumi.get(self, "gateway_version")

    @property
    @pulumi.getter(name="ingressClassName")
    def ingress_class_name(self) -> pulumi.Output[str]:
        """
        ingress class name.
        """
        return pulumi.get(self, "ingress_class_name")

    @property
    @pulumi.getter(name="instancePorts")
    def instance_ports(self) -> pulumi.Output[Sequence['outputs.CngwGatewayInstancePort']]:
        """
        Port information that the instance listens to.
        """
        return pulumi.get(self, "instance_ports")

    @property
    @pulumi.getter(name="internetConfig")
    def internet_config(self) -> pulumi.Output[Optional['outputs.CngwGatewayInternetConfig']]:
        """
        internet configration.
        """
        return pulumi.get(self, "internet_config")

    @property
    @pulumi.getter(name="internetMaxBandwidthOut")
    def internet_max_bandwidth_out(self) -> pulumi.Output[Optional[int]]:
        """
        public network outbound traffic bandwidth,[1,2048]Mbps.
        """
        return pulumi.get(self, "internet_max_bandwidth_out")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        gateway name, supports up to 60 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> pulumi.Output['outputs.CngwGatewayNodeConfig']:
        """
        gateway node configration.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="publicIpAddresses")
    def public_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        Public IP address list.
        """
        return pulumi.get(self, "public_ip_addresses")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tradeType")
    def trade_type(self) -> pulumi.Output[Optional[int]]:
        """
        trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        """
        return pulumi.get(self, "trade_type")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        gateway type,currently only supports kong.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> pulumi.Output['outputs.CngwGatewayVpcConfig']:
        """
        vpc information.
        """
        return pulumi.get(self, "vpc_config")

