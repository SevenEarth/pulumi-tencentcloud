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

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 blueprint_id: pulumi.Input[str],
                 bundle_id: pulumi.Input[str],
                 instance_name: pulumi.Input[str],
                 period: pulumi.Input[int],
                 renew_flag: pulumi.Input[str],
                 client_token: Optional[pulumi.Input[str]] = None,
                 containers: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 login_configuration: Optional[pulumi.Input['InstanceLoginConfigurationArgs']] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] blueprint_id: ID of the Lighthouse image.
        :param pulumi.Input[str] bundle_id: ID of the Lighthouse package.
        :param pulumi.Input[str] instance_name: The display name of the Lighthouse instance.
        :param pulumi.Input[int] period: Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        :param pulumi.Input[str] renew_flag: Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        :param pulumi.Input[str] client_token: A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]] containers: Configuration of the containers to create.
        :param pulumi.Input[bool] dry_run: Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        :param pulumi.Input['InstanceLoginConfigurationArgs'] login_configuration: Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        :param pulumi.Input[str] zone: List of availability zones. A random AZ is selected by default.
        """
        pulumi.set(__self__, "blueprint_id", blueprint_id)
        pulumi.set(__self__, "bundle_id", bundle_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "period", period)
        pulumi.set(__self__, "renew_flag", renew_flag)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if containers is not None:
            pulumi.set(__self__, "containers", containers)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if login_configuration is not None:
            pulumi.set(__self__, "login_configuration", login_configuration)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="blueprintId")
    def blueprint_id(self) -> pulumi.Input[str]:
        """
        ID of the Lighthouse image.
        """
        return pulumi.get(self, "blueprint_id")

    @blueprint_id.setter
    def blueprint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "blueprint_id", value)

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> pulumi.Input[str]:
        """
        ID of the Lighthouse package.
        """
        return pulumi.get(self, "bundle_id")

    @bundle_id.setter
    def bundle_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bundle_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        The display name of the Lighthouse instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[int]:
        """
        Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[int]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="renewFlag")
    def renew_flag(self) -> pulumi.Input[str]:
        """
        Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        """
        return pulumi.get(self, "renew_flag")

    @renew_flag.setter
    def renew_flag(self, value: pulumi.Input[str]):
        pulumi.set(self, "renew_flag", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter
    def containers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]]:
        """
        Configuration of the containers to create.
        """
        return pulumi.get(self, "containers")

    @containers.setter
    def containers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]]):
        pulumi.set(self, "containers", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="loginConfiguration")
    def login_configuration(self) -> Optional[pulumi.Input['InstanceLoginConfigurationArgs']]:
        """
        Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        """
        return pulumi.get(self, "login_configuration")

    @login_configuration.setter
    def login_configuration(self, value: Optional[pulumi.Input['InstanceLoginConfigurationArgs']]):
        pulumi.set(self, "login_configuration", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        List of availability zones. A random AZ is selected by default.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 blueprint_id: Optional[pulumi.Input[str]] = None,
                 bundle_id: Optional[pulumi.Input[str]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 containers: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 login_configuration: Optional[pulumi.Input['InstanceLoginConfigurationArgs']] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_flag: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] blueprint_id: ID of the Lighthouse image.
        :param pulumi.Input[str] bundle_id: ID of the Lighthouse package.
        :param pulumi.Input[str] client_token: A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]] containers: Configuration of the containers to create.
        :param pulumi.Input[bool] dry_run: Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        :param pulumi.Input[str] instance_name: The display name of the Lighthouse instance.
        :param pulumi.Input['InstanceLoginConfigurationArgs'] login_configuration: Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        :param pulumi.Input[int] period: Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        :param pulumi.Input[str] renew_flag: Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        :param pulumi.Input[str] zone: List of availability zones. A random AZ is selected by default.
        """
        if blueprint_id is not None:
            pulumi.set(__self__, "blueprint_id", blueprint_id)
        if bundle_id is not None:
            pulumi.set(__self__, "bundle_id", bundle_id)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if containers is not None:
            pulumi.set(__self__, "containers", containers)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if login_configuration is not None:
            pulumi.set(__self__, "login_configuration", login_configuration)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if renew_flag is not None:
            pulumi.set(__self__, "renew_flag", renew_flag)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="blueprintId")
    def blueprint_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Lighthouse image.
        """
        return pulumi.get(self, "blueprint_id")

    @blueprint_id.setter
    def blueprint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "blueprint_id", value)

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Lighthouse package.
        """
        return pulumi.get(self, "bundle_id")

    @bundle_id.setter
    def bundle_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bundle_id", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter
    def containers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]]:
        """
        Configuration of the containers to create.
        """
        return pulumi.get(self, "containers")

    @containers.setter
    def containers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceContainerArgs']]]]):
        pulumi.set(self, "containers", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Lighthouse instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="loginConfiguration")
    def login_configuration(self) -> Optional[pulumi.Input['InstanceLoginConfigurationArgs']]:
        """
        Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        """
        return pulumi.get(self, "login_configuration")

    @login_configuration.setter
    def login_configuration(self, value: Optional[pulumi.Input['InstanceLoginConfigurationArgs']]):
        pulumi.set(self, "login_configuration", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="renewFlag")
    def renew_flag(self) -> Optional[pulumi.Input[str]]:
        """
        Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        """
        return pulumi.get(self, "renew_flag")

    @renew_flag.setter
    def renew_flag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "renew_flag", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        List of availability zones. A random AZ is selected by default.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blueprint_id: Optional[pulumi.Input[str]] = None,
                 bundle_id: Optional[pulumi.Input[str]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 containers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceContainerArgs']]]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 login_configuration: Optional[pulumi.Input[pulumi.InputType['InstanceLoginConfigurationArgs']]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_flag: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a lighthouse instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        lighthouse = tencentcloud.lighthouse.Instance("lighthouse",
            blueprint_id="lhbp-f1lkcd41",
            bundle_id="bundle2022_gen_01",
            containers=[
                tencentcloud.lighthouse.InstanceContainerArgs(
                    command="ls -l",
                    container_image="ccr.ccs.tencentyun.com/qcloud/nginx",
                    container_name="nginx",
                    envs=[
                        tencentcloud.lighthouse.InstanceContainerEnvArgs(
                            key="key",
                            value="value",
                        ),
                        tencentcloud.lighthouse.InstanceContainerEnvArgs(
                            key="key2",
                            value="value2",
                        ),
                    ],
                    publish_ports=[
                        tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                            container_port=80,
                            host_port=80,
                            ip="127.0.0.1",
                            protocol="tcp",
                        ),
                        tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                            container_port=36000,
                            host_port=36000,
                            ip="127.0.0.1",
                            protocol="tcp",
                        ),
                    ],
                    volumes=[
                        tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                            container_path="/data",
                            host_path="/tmp",
                        ),
                        tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                            container_path="/var",
                            host_path="/tmp",
                        ),
                    ],
                ),
                tencentcloud.lighthouse.InstanceContainerArgs(
                    command="echo \"hello\"",
                    container_image="ccr.ccs.tencentyun.com/qcloud/resty",
                    container_name="resty",
                    envs=[tencentcloud.lighthouse.InstanceContainerEnvArgs(
                        key="key2",
                        value="value2",
                    )],
                    publish_ports=[tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                        container_port=80,
                        host_port=80,
                        ip="127.0.0.1",
                        protocol="udp",
                    )],
                    volumes=[tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                        container_path="/var",
                        host_path="/tmp",
                    )],
                ),
            ],
            instance_name="hello world",
            period=1,
            renew_flag="NOTIFY_AND_AUTO_RENEW",
            zone="ap-guangzhou-3")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blueprint_id: ID of the Lighthouse image.
        :param pulumi.Input[str] bundle_id: ID of the Lighthouse package.
        :param pulumi.Input[str] client_token: A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceContainerArgs']]]] containers: Configuration of the containers to create.
        :param pulumi.Input[bool] dry_run: Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        :param pulumi.Input[str] instance_name: The display name of the Lighthouse instance.
        :param pulumi.Input[pulumi.InputType['InstanceLoginConfigurationArgs']] login_configuration: Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        :param pulumi.Input[int] period: Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        :param pulumi.Input[str] renew_flag: Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        :param pulumi.Input[str] zone: List of availability zones. A random AZ is selected by default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a lighthouse instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        lighthouse = tencentcloud.lighthouse.Instance("lighthouse",
            blueprint_id="lhbp-f1lkcd41",
            bundle_id="bundle2022_gen_01",
            containers=[
                tencentcloud.lighthouse.InstanceContainerArgs(
                    command="ls -l",
                    container_image="ccr.ccs.tencentyun.com/qcloud/nginx",
                    container_name="nginx",
                    envs=[
                        tencentcloud.lighthouse.InstanceContainerEnvArgs(
                            key="key",
                            value="value",
                        ),
                        tencentcloud.lighthouse.InstanceContainerEnvArgs(
                            key="key2",
                            value="value2",
                        ),
                    ],
                    publish_ports=[
                        tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                            container_port=80,
                            host_port=80,
                            ip="127.0.0.1",
                            protocol="tcp",
                        ),
                        tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                            container_port=36000,
                            host_port=36000,
                            ip="127.0.0.1",
                            protocol="tcp",
                        ),
                    ],
                    volumes=[
                        tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                            container_path="/data",
                            host_path="/tmp",
                        ),
                        tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                            container_path="/var",
                            host_path="/tmp",
                        ),
                    ],
                ),
                tencentcloud.lighthouse.InstanceContainerArgs(
                    command="echo \"hello\"",
                    container_image="ccr.ccs.tencentyun.com/qcloud/resty",
                    container_name="resty",
                    envs=[tencentcloud.lighthouse.InstanceContainerEnvArgs(
                        key="key2",
                        value="value2",
                    )],
                    publish_ports=[tencentcloud.lighthouse.InstanceContainerPublishPortArgs(
                        container_port=80,
                        host_port=80,
                        ip="127.0.0.1",
                        protocol="udp",
                    )],
                    volumes=[tencentcloud.lighthouse.InstanceContainerVolumeArgs(
                        container_path="/var",
                        host_path="/tmp",
                    )],
                ),
            ],
            instance_name="hello world",
            period=1,
            renew_flag="NOTIFY_AND_AUTO_RENEW",
            zone="ap-guangzhou-3")
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 blueprint_id: Optional[pulumi.Input[str]] = None,
                 bundle_id: Optional[pulumi.Input[str]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 containers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceContainerArgs']]]]] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 login_configuration: Optional[pulumi.Input[pulumi.InputType['InstanceLoginConfigurationArgs']]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 renew_flag: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstanceArgs.__new__(InstanceArgs)

            if blueprint_id is None and not opts.urn:
                raise TypeError("Missing required property 'blueprint_id'")
            __props__.__dict__["blueprint_id"] = blueprint_id
            if bundle_id is None and not opts.urn:
                raise TypeError("Missing required property 'bundle_id'")
            __props__.__dict__["bundle_id"] = bundle_id
            __props__.__dict__["client_token"] = client_token
            __props__.__dict__["containers"] = containers
            __props__.__dict__["dry_run"] = dry_run
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            __props__.__dict__["login_configuration"] = login_configuration
            if period is None and not opts.urn:
                raise TypeError("Missing required property 'period'")
            __props__.__dict__["period"] = period
            if renew_flag is None and not opts.urn:
                raise TypeError("Missing required property 'renew_flag'")
            __props__.__dict__["renew_flag"] = renew_flag
            __props__.__dict__["zone"] = zone
        super(Instance, __self__).__init__(
            'tencentcloud:Lighthouse/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            blueprint_id: Optional[pulumi.Input[str]] = None,
            bundle_id: Optional[pulumi.Input[str]] = None,
            client_token: Optional[pulumi.Input[str]] = None,
            containers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceContainerArgs']]]]] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            login_configuration: Optional[pulumi.Input[pulumi.InputType['InstanceLoginConfigurationArgs']]] = None,
            period: Optional[pulumi.Input[int]] = None,
            renew_flag: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] blueprint_id: ID of the Lighthouse image.
        :param pulumi.Input[str] bundle_id: ID of the Lighthouse package.
        :param pulumi.Input[str] client_token: A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceContainerArgs']]]] containers: Configuration of the containers to create.
        :param pulumi.Input[bool] dry_run: Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        :param pulumi.Input[str] instance_name: The display name of the Lighthouse instance.
        :param pulumi.Input[pulumi.InputType['InstanceLoginConfigurationArgs']] login_configuration: Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        :param pulumi.Input[int] period: Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        :param pulumi.Input[str] renew_flag: Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        :param pulumi.Input[str] zone: List of availability zones. A random AZ is selected by default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["blueprint_id"] = blueprint_id
        __props__.__dict__["bundle_id"] = bundle_id
        __props__.__dict__["client_token"] = client_token
        __props__.__dict__["containers"] = containers
        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["login_configuration"] = login_configuration
        __props__.__dict__["period"] = period
        __props__.__dict__["renew_flag"] = renew_flag
        __props__.__dict__["zone"] = zone
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blueprintId")
    def blueprint_id(self) -> pulumi.Output[str]:
        """
        ID of the Lighthouse image.
        """
        return pulumi.get(self, "blueprint_id")

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> pulumi.Output[str]:
        """
        ID of the Lighthouse package.
        """
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[Optional[str]]:
        """
        A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def containers(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceContainer']]]:
        """
        Configuration of the containers to create.
        """
        return pulumi.get(self, "containers")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        The display name of the Lighthouse instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="loginConfiguration")
    def login_configuration(self) -> pulumi.Output[Optional['outputs.InstanceLoginConfiguration']]:
        """
        Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
        """
        return pulumi.get(self, "login_configuration")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[int]:
        """
        Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="renewFlag")
    def renew_flag(self) -> pulumi.Output[str]:
        """
        Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
        """
        return pulumi.get(self, "renew_flag")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        """
        List of availability zones. A random AZ is selected by default.
        """
        return pulumi.get(self, "zone")

