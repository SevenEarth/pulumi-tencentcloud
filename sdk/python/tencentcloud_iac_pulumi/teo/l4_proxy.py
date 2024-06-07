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

__all__ = ['L4ProxyArgs', 'L4Proxy']

@pulumi.input_type
class L4ProxyArgs:
    def __init__(__self__, *,
                 proxy_name: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 accelerate_mainland: Optional[pulumi.Input[str]] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 ddos_protection_config: Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 static_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a L4Proxy resource.
        :param pulumi.Input[str] proxy_name: Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        :param pulumi.Input[str] zone_id: Site ID.
        :param pulumi.Input[str] accelerate_mainland: Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] area: Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        :param pulumi.Input['L4ProxyDdosProtectionConfigArgs'] ddos_protection_config: Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        :param pulumi.Input[str] ipv6: Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] static_ip: Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        pulumi.set(__self__, "proxy_name", proxy_name)
        pulumi.set(__self__, "zone_id", zone_id)
        if accelerate_mainland is not None:
            pulumi.set(__self__, "accelerate_mainland", accelerate_mainland)
        if area is not None:
            pulumi.set(__self__, "area", area)
        if ddos_protection_config is not None:
            pulumi.set(__self__, "ddos_protection_config", ddos_protection_config)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if static_ip is not None:
            pulumi.set(__self__, "static_ip", static_ip)

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> pulumi.Input[str]:
        """
        Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        """
        return pulumi.get(self, "proxy_name")

    @proxy_name.setter
    def proxy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "proxy_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="accelerateMainland")
    def accelerate_mainland(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "accelerate_mainland")

    @accelerate_mainland.setter
    def accelerate_mainland(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accelerate_mainland", value)

    @property
    @pulumi.getter
    def area(self) -> Optional[pulumi.Input[str]]:
        """
        Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        """
        return pulumi.get(self, "area")

    @area.setter
    def area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "area", value)

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']]:
        """
        Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        """
        return pulumi.get(self, "ddos_protection_config")

    @ddos_protection_config.setter
    def ddos_protection_config(self, value: Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']]):
        pulumi.set(self, "ddos_protection_config", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="staticIp")
    def static_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "static_ip")

    @static_ip.setter
    def static_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "static_ip", value)


@pulumi.input_type
class _L4ProxyState:
    def __init__(__self__, *,
                 accelerate_mainland: Optional[pulumi.Input[str]] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 ddos_protection_config: Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 static_ip: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering L4Proxy resources.
        :param pulumi.Input[str] accelerate_mainland: Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] area: Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        :param pulumi.Input['L4ProxyDdosProtectionConfigArgs'] ddos_protection_config: Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        :param pulumi.Input[str] ipv6: Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] proxy_name: Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        :param pulumi.Input[str] static_ip: Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        if accelerate_mainland is not None:
            pulumi.set(__self__, "accelerate_mainland", accelerate_mainland)
        if area is not None:
            pulumi.set(__self__, "area", area)
        if ddos_protection_config is not None:
            pulumi.set(__self__, "ddos_protection_config", ddos_protection_config)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if proxy_name is not None:
            pulumi.set(__self__, "proxy_name", proxy_name)
        if static_ip is not None:
            pulumi.set(__self__, "static_ip", static_ip)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="accelerateMainland")
    def accelerate_mainland(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "accelerate_mainland")

    @accelerate_mainland.setter
    def accelerate_mainland(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accelerate_mainland", value)

    @property
    @pulumi.getter
    def area(self) -> Optional[pulumi.Input[str]]:
        """
        Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        """
        return pulumi.get(self, "area")

    @area.setter
    def area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "area", value)

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']]:
        """
        Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        """
        return pulumi.get(self, "ddos_protection_config")

    @ddos_protection_config.setter
    def ddos_protection_config(self, value: Optional[pulumi.Input['L4ProxyDdosProtectionConfigArgs']]):
        pulumi.set(self, "ddos_protection_config", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        """
        return pulumi.get(self, "proxy_name")

    @proxy_name.setter
    def proxy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_name", value)

    @property
    @pulumi.getter(name="staticIp")
    def static_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "static_ip")

    @static_ip.setter
    def static_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "static_ip", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class L4Proxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerate_mainland: Optional[pulumi.Input[str]] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 ddos_protection_config: Optional[pulumi.Input[pulumi.InputType['L4ProxyDdosProtectionConfigArgs']]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 static_ip: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a teo teo_l4_proxy

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        proxy = tencentcloud.teo.L4Proxy("proxy",
            accelerate_mainland="off",
            area="overseas",
            ipv6="off",
            proxy_name="proxy-test",
            static_ip="off",
            zone_id="zone-2qtuhspy6cr7")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo teo_l4_proxy can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/l4Proxy:L4Proxy teo_l4_proxy teo_l4_proxy_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerate_mainland: Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] area: Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        :param pulumi.Input[pulumi.InputType['L4ProxyDdosProtectionConfigArgs']] ddos_protection_config: Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        :param pulumi.Input[str] ipv6: Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] proxy_name: Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        :param pulumi.Input[str] static_ip: Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: L4ProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a teo teo_l4_proxy

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        proxy = tencentcloud.teo.L4Proxy("proxy",
            accelerate_mainland="off",
            area="overseas",
            ipv6="off",
            proxy_name="proxy-test",
            static_ip="off",
            zone_id="zone-2qtuhspy6cr7")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo teo_l4_proxy can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/l4Proxy:L4Proxy teo_l4_proxy teo_l4_proxy_id
        ```

        :param str resource_name: The name of the resource.
        :param L4ProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(L4ProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerate_mainland: Optional[pulumi.Input[str]] = None,
                 area: Optional[pulumi.Input[str]] = None,
                 ddos_protection_config: Optional[pulumi.Input[pulumi.InputType['L4ProxyDdosProtectionConfigArgs']]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 proxy_name: Optional[pulumi.Input[str]] = None,
                 static_ip: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = L4ProxyArgs.__new__(L4ProxyArgs)

            __props__.__dict__["accelerate_mainland"] = accelerate_mainland
            __props__.__dict__["area"] = area
            __props__.__dict__["ddos_protection_config"] = ddos_protection_config
            __props__.__dict__["ipv6"] = ipv6
            if proxy_name is None and not opts.urn:
                raise TypeError("Missing required property 'proxy_name'")
            __props__.__dict__["proxy_name"] = proxy_name
            __props__.__dict__["static_ip"] = static_ip
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(L4Proxy, __self__).__init__(
            'tencentcloud:Teo/l4Proxy:L4Proxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerate_mainland: Optional[pulumi.Input[str]] = None,
            area: Optional[pulumi.Input[str]] = None,
            ddos_protection_config: Optional[pulumi.Input[pulumi.InputType['L4ProxyDdosProtectionConfigArgs']]] = None,
            ipv6: Optional[pulumi.Input[str]] = None,
            proxy_name: Optional[pulumi.Input[str]] = None,
            static_ip: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'L4Proxy':
        """
        Get an existing L4Proxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerate_mainland: Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] area: Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        :param pulumi.Input[pulumi.InputType['L4ProxyDdosProtectionConfigArgs']] ddos_protection_config: Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        :param pulumi.Input[str] ipv6: Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] proxy_name: Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        :param pulumi.Input[str] static_ip: Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _L4ProxyState.__new__(_L4ProxyState)

        __props__.__dict__["accelerate_mainland"] = accelerate_mainland
        __props__.__dict__["area"] = area
        __props__.__dict__["ddos_protection_config"] = ddos_protection_config
        __props__.__dict__["ipv6"] = ipv6
        __props__.__dict__["proxy_name"] = proxy_name
        __props__.__dict__["static_ip"] = static_ip
        __props__.__dict__["zone_id"] = zone_id
        return L4Proxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accelerateMainland")
    def accelerate_mainland(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to enable network optimization in the Chinese mainland. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "accelerate_mainland")

    @property
    @pulumi.getter
    def area(self) -> pulumi.Output[Optional[str]]:
        """
        Acceleration zone of the Layer 4 proxy instance. `mainland`: Availability zone in the Chinese mainland; `overseas`: Global availability zone (excluding the Chinese mainland); `global`: Global availability zone.
        """
        return pulumi.get(self, "area")

    @property
    @pulumi.getter(name="ddosProtectionConfig")
    def ddos_protection_config(self) -> pulumi.Output[Optional['outputs.L4ProxyDdosProtectionConfig']]:
        """
        Layer 3/Layer 4 DDoS protection. The default protection option of the platform will be used if it is left empty. For details, see [Exclusive DDoS Protection Usage](https://intl.cloud.tencent.com/document/product/1552/95994?from_cn_redirect=1).
        """
        return pulumi.get(self, "ddos_protection_config")

    @property
    @pulumi.getter
    def ipv6(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to enable IPv6 access. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter(name="proxyName")
    def proxy_name(self) -> pulumi.Output[str]:
        """
        Layer 4 proxy instance name. You can enter 1-50 characters. Valid characters are a-z, 0-9, and hyphens (-). However, hyphens (-) cannot be used individually or consecutively and should not be placed at the beginning or end of the name. Modifications are not allowed after creation.
        """
        return pulumi.get(self, "proxy_name")

    @property
    @pulumi.getter(name="staticIp")
    def static_ip(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to enable the fixed IP address. The default value off is used if left empty. This configuration can only be enabled in certain acceleration zones and security protection configurations. For details, see [Creating an L4 Proxy Instance](https://intl.cloud.tencent.com/document/product/1552/90025?from_cn_redirect=1). Valid values: `on`: Enable; `off`: Disable.
        """
        return pulumi.get(self, "static_ip")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

