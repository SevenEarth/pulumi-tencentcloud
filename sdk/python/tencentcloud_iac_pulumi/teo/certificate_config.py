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

__all__ = ['CertificateConfigArgs', 'CertificateConfig']

@pulumi.input_type
class CertificateConfigArgs:
    def __init__(__self__, *,
                 host: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 mode: Optional[pulumi.Input[str]] = None,
                 server_cert_infos: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]] = None):
        """
        The set of arguments for constructing a CertificateConfig resource.
        :param pulumi.Input[str] host: Acceleration domain name that needs to modify the certificate configuration.
        :param pulumi.Input[str] zone_id: Site ID.
        :param pulumi.Input[str] mode: Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]] server_cert_infos: SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        """
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "zone_id", zone_id)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if server_cert_infos is not None:
            pulumi.set(__self__, "server_cert_infos", server_cert_infos)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        Acceleration domain name that needs to modify the certificate configuration.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

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
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="serverCertInfos")
    def server_cert_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]]:
        """
        SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        """
        return pulumi.get(self, "server_cert_infos")

    @server_cert_infos.setter
    def server_cert_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]]):
        pulumi.set(self, "server_cert_infos", value)


@pulumi.input_type
class _CertificateConfigState:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 server_cert_infos: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CertificateConfig resources.
        :param pulumi.Input[str] host: Acceleration domain name that needs to modify the certificate configuration.
        :param pulumi.Input[str] mode: Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]] server_cert_infos: SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        if host is not None:
            pulumi.set(__self__, "host", host)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if server_cert_infos is not None:
            pulumi.set(__self__, "server_cert_infos", server_cert_infos)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Acceleration domain name that needs to modify the certificate configuration.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="serverCertInfos")
    def server_cert_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]]:
        """
        SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        """
        return pulumi.get(self, "server_cert_infos")

    @server_cert_infos.setter
    def server_cert_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CertificateConfigServerCertInfoArgs']]]]):
        pulumi.set(self, "server_cert_infos", value)

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


class CertificateConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 server_cert_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateConfigServerCertInfoArgs']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a teo certificate

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        certificate = tencentcloud.teo.CertificateConfig("certificate",
            host="test.tencentcloud-terraform-provider.cn",
            mode="eofreecert",
            zone_id="zone-2o1t24kgy362")
        ```
        <!--End PulumiCodeChooser -->

        ### Configure SSL certificate

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        certificate = tencentcloud.teo.CertificateConfig("certificate",
            host="test.tencentcloud-terraform-provider.cn",
            mode="sslcert",
            server_cert_infos=[tencentcloud.teo.CertificateConfigServerCertInfoArgs(
                cert_id="8xiUJIJd",
            )],
            zone_id="zone-2o1t24kgy362")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo certificate can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/certificateConfig:CertificateConfig certificate zone_id#host#cert_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: Acceleration domain name that needs to modify the certificate configuration.
        :param pulumi.Input[str] mode: Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateConfigServerCertInfoArgs']]]] server_cert_infos: SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a teo certificate

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        certificate = tencentcloud.teo.CertificateConfig("certificate",
            host="test.tencentcloud-terraform-provider.cn",
            mode="eofreecert",
            zone_id="zone-2o1t24kgy362")
        ```
        <!--End PulumiCodeChooser -->

        ### Configure SSL certificate

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        certificate = tencentcloud.teo.CertificateConfig("certificate",
            host="test.tencentcloud-terraform-provider.cn",
            mode="sslcert",
            server_cert_infos=[tencentcloud.teo.CertificateConfigServerCertInfoArgs(
                cert_id="8xiUJIJd",
            )],
            zone_id="zone-2o1t24kgy362")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        teo certificate can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Teo/certificateConfig:CertificateConfig certificate zone_id#host#cert_id
        ```

        :param str resource_name: The name of the resource.
        :param CertificateConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 server_cert_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateConfigServerCertInfoArgs']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertificateConfigArgs.__new__(CertificateConfigArgs)

            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            __props__.__dict__["mode"] = mode
            __props__.__dict__["server_cert_infos"] = server_cert_infos
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(CertificateConfig, __self__).__init__(
            'tencentcloud:Teo/certificateConfig:CertificateConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            server_cert_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateConfigServerCertInfoArgs']]]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'CertificateConfig':
        """
        Get an existing CertificateConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: Acceleration domain name that needs to modify the certificate configuration.
        :param pulumi.Input[str] mode: Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CertificateConfigServerCertInfoArgs']]]] server_cert_infos: SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        :param pulumi.Input[str] zone_id: Site ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateConfigState.__new__(_CertificateConfigState)

        __props__.__dict__["host"] = host
        __props__.__dict__["mode"] = mode
        __props__.__dict__["server_cert_infos"] = server_cert_infos
        __props__.__dict__["zone_id"] = zone_id
        return CertificateConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Acceleration domain name that needs to modify the certificate configuration.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Mode of configuring the certificate, the values are: `disable`: Do not configure the certificate; `eofreecert`: Configure EdgeOne free certificate; `sslcert`: Configure SSL certificate. If not filled in, the default value is `disable`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter(name="serverCertInfos")
    def server_cert_infos(self) -> pulumi.Output[Sequence['outputs.CertificateConfigServerCertInfo']]:
        """
        SSL certificate configuration, this parameter takes effect only when mode = sslcert, just enter the corresponding CertId. You can go to the SSL certificate list to view the CertId.
        """
        return pulumi.get(self, "server_cert_infos")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        Site ID.
        """
        return pulumi.get(self, "zone_id")

