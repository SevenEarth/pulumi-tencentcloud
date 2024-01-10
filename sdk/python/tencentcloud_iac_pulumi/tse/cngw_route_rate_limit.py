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

__all__ = ['CngwRouteRateLimitArgs', 'CngwRouteRateLimit']

@pulumi.input_type
class CngwRouteRateLimitArgs:
    def __init__(__self__, *,
                 gateway_id: pulumi.Input[str],
                 limit_detail: pulumi.Input['CngwRouteRateLimitLimitDetailArgs'],
                 route_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a CngwRouteRateLimit resource.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input['CngwRouteRateLimitLimitDetailArgs'] limit_detail: rate limit configuration.
        :param pulumi.Input[str] route_id: Route id, or route name.
        """
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "limit_detail", limit_detail)
        pulumi.set(__self__, "route_id", route_id)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="limitDetail")
    def limit_detail(self) -> pulumi.Input['CngwRouteRateLimitLimitDetailArgs']:
        """
        rate limit configuration.
        """
        return pulumi.get(self, "limit_detail")

    @limit_detail.setter
    def limit_detail(self, value: pulumi.Input['CngwRouteRateLimitLimitDetailArgs']):
        pulumi.set(self, "limit_detail", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Input[str]:
        """
        Route id, or route name.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_id", value)


@pulumi.input_type
class _CngwRouteRateLimitState:
    def __init__(__self__, *,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 limit_detail: Optional[pulumi.Input['CngwRouteRateLimitLimitDetailArgs']] = None,
                 route_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CngwRouteRateLimit resources.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input['CngwRouteRateLimitLimitDetailArgs'] limit_detail: rate limit configuration.
        :param pulumi.Input[str] route_id: Route id, or route name.
        """
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if limit_detail is not None:
            pulumi.set(__self__, "limit_detail", limit_detail)
        if route_id is not None:
            pulumi.set(__self__, "route_id", route_id)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="limitDetail")
    def limit_detail(self) -> Optional[pulumi.Input['CngwRouteRateLimitLimitDetailArgs']]:
        """
        rate limit configuration.
        """
        return pulumi.get(self, "limit_detail")

    @limit_detail.setter
    def limit_detail(self, value: Optional[pulumi.Input['CngwRouteRateLimitLimitDetailArgs']]):
        pulumi.set(self, "limit_detail", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> Optional[pulumi.Input[str]]:
        """
        Route id, or route name.
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_id", value)


class CngwRouteRateLimit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 limit_detail: Optional[pulumi.Input[pulumi.InputType['CngwRouteRateLimitLimitDetailArgs']]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tse cngw_route_rate_limit

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
        cngw_service = tencentcloud.tse.CngwService("cngwService",
            gateway_id=cngw_gateway.id,
            path="/test",
            protocol="http",
            retries=5,
            timeout=60000,
            upstream_type="HostIP",
            upstream_info=tencentcloud.tse.CngwServiceUpstreamInfoArgs(
                algorithm="round-robin",
                auto_scaling_cvm_port=0,
                host="arunma.cn",
                port=8012,
                slow_start=0,
            ))
        cngw_route = tencentcloud.tse.CngwRoute("cngwRoute",
            destination_ports=[],
            force_https=False,
            gateway_id=cngw_gateway.id,
            hosts=["192.168.0.1:9090"],
            https_redirect_status_code=426,
            paths=["/user"],
            headers=[tencentcloud.tse.CngwRouteHeaderArgs(
                key="req",
                value="terraform",
            )],
            preserve_host=False,
            protocols=[
                "http",
                "https",
            ],
            route_name="terraform-route",
            service_id=cngw_service.service_id,
            strip_path=True)
        cngw_route_rate_limit = tencentcloud.tse.CngwRouteRateLimit("cngwRouteRateLimit",
            gateway_id=cngw_gateway.id,
            route_id=cngw_route.route_id,
            limit_detail=tencentcloud.tse.CngwRouteRateLimitLimitDetailArgs(
                enabled=True,
                header="req",
                hide_client_headers=True,
                is_delay=True,
                limit_by="header",
                line_up_time=10,
                policy="redis",
                response_type="default",
                qps_thresholds=[tencentcloud.tse.CngwRouteRateLimitLimitDetailQpsThresholdArgs(
                    max=10,
                    unit="minute",
                )],
            ))
        ```

        ## Import

        tse cngw_route_rate_limit can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tse/cngwRouteRateLimit:CngwRouteRateLimit cngw_route_rate_limit gatewayId#routeId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[pulumi.InputType['CngwRouteRateLimitLimitDetailArgs']] limit_detail: rate limit configuration.
        :param pulumi.Input[str] route_id: Route id, or route name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CngwRouteRateLimitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tse cngw_route_rate_limit

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
        cngw_service = tencentcloud.tse.CngwService("cngwService",
            gateway_id=cngw_gateway.id,
            path="/test",
            protocol="http",
            retries=5,
            timeout=60000,
            upstream_type="HostIP",
            upstream_info=tencentcloud.tse.CngwServiceUpstreamInfoArgs(
                algorithm="round-robin",
                auto_scaling_cvm_port=0,
                host="arunma.cn",
                port=8012,
                slow_start=0,
            ))
        cngw_route = tencentcloud.tse.CngwRoute("cngwRoute",
            destination_ports=[],
            force_https=False,
            gateway_id=cngw_gateway.id,
            hosts=["192.168.0.1:9090"],
            https_redirect_status_code=426,
            paths=["/user"],
            headers=[tencentcloud.tse.CngwRouteHeaderArgs(
                key="req",
                value="terraform",
            )],
            preserve_host=False,
            protocols=[
                "http",
                "https",
            ],
            route_name="terraform-route",
            service_id=cngw_service.service_id,
            strip_path=True)
        cngw_route_rate_limit = tencentcloud.tse.CngwRouteRateLimit("cngwRouteRateLimit",
            gateway_id=cngw_gateway.id,
            route_id=cngw_route.route_id,
            limit_detail=tencentcloud.tse.CngwRouteRateLimitLimitDetailArgs(
                enabled=True,
                header="req",
                hide_client_headers=True,
                is_delay=True,
                limit_by="header",
                line_up_time=10,
                policy="redis",
                response_type="default",
                qps_thresholds=[tencentcloud.tse.CngwRouteRateLimitLimitDetailQpsThresholdArgs(
                    max=10,
                    unit="minute",
                )],
            ))
        ```

        ## Import

        tse cngw_route_rate_limit can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tse/cngwRouteRateLimit:CngwRouteRateLimit cngw_route_rate_limit gatewayId#routeId
        ```

        :param str resource_name: The name of the resource.
        :param CngwRouteRateLimitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CngwRouteRateLimitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 limit_detail: Optional[pulumi.Input[pulumi.InputType['CngwRouteRateLimitLimitDetailArgs']]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = CngwRouteRateLimitArgs.__new__(CngwRouteRateLimitArgs)

            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            if limit_detail is None and not opts.urn:
                raise TypeError("Missing required property 'limit_detail'")
            __props__.__dict__["limit_detail"] = limit_detail
            if route_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_id'")
            __props__.__dict__["route_id"] = route_id
        super(CngwRouteRateLimit, __self__).__init__(
            'tencentcloud:Tse/cngwRouteRateLimit:CngwRouteRateLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            limit_detail: Optional[pulumi.Input[pulumi.InputType['CngwRouteRateLimitLimitDetailArgs']]] = None,
            route_id: Optional[pulumi.Input[str]] = None) -> 'CngwRouteRateLimit':
        """
        Get an existing CngwRouteRateLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: gateway ID.
        :param pulumi.Input[pulumi.InputType['CngwRouteRateLimitLimitDetailArgs']] limit_detail: rate limit configuration.
        :param pulumi.Input[str] route_id: Route id, or route name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CngwRouteRateLimitState.__new__(_CngwRouteRateLimitState)

        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["limit_detail"] = limit_detail
        __props__.__dict__["route_id"] = route_id
        return CngwRouteRateLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        gateway ID.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="limitDetail")
    def limit_detail(self) -> pulumi.Output['outputs.CngwRouteRateLimitLimitDetail']:
        """
        rate limit configuration.
        """
        return pulumi.get(self, "limit_detail")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[str]:
        """
        Route id, or route name.
        """
        return pulumi.get(self, "route_id")

