# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TmpGrafanaConfigArgs', 'TmpGrafanaConfig']

@pulumi.input_type
class TmpGrafanaConfigArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TmpGrafanaConfig resource.
        :param pulumi.Input[str] config: JSON encoded string.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[str]:
        """
        JSON encoded string.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[str]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _TmpGrafanaConfigState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpGrafanaConfig resources.
        :param pulumi.Input[str] config: JSON encoded string.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        JSON encoded string.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class TmpGrafanaConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor tmp_grafana_config

        ## Example Usage

        ```python
        import pulumi
        import json
        import tencentcloud_iac_pulumi as tencentcloud

        tmp_grafana_config = tencentcloud.monitor.TmpGrafanaConfig("tmpGrafanaConfig",
            config=json.dumps({
                "server": {
                    "http_port": 8080,
                    "root_url": "https://cloud-grafana.woa.com/grafana-ffrdnrfa/",
                    "serve_from_sub_path": True,
                },
            }),
            instance_id="grafana-29phe08q")
        ```

        ## Import

        monitor tmp_grafana_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpGrafanaConfig:TmpGrafanaConfig tmp_grafana_config tmp_grafana_config_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: JSON encoded string.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpGrafanaConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor tmp_grafana_config

        ## Example Usage

        ```python
        import pulumi
        import json
        import tencentcloud_iac_pulumi as tencentcloud

        tmp_grafana_config = tencentcloud.monitor.TmpGrafanaConfig("tmpGrafanaConfig",
            config=json.dumps({
                "server": {
                    "http_port": 8080,
                    "root_url": "https://cloud-grafana.woa.com/grafana-ffrdnrfa/",
                    "serve_from_sub_path": True,
                },
            }),
            instance_id="grafana-29phe08q")
        ```

        ## Import

        monitor tmp_grafana_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpGrafanaConfig:TmpGrafanaConfig tmp_grafana_config tmp_grafana_config_id
        ```

        :param str resource_name: The name of the resource.
        :param TmpGrafanaConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpGrafanaConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TmpGrafanaConfigArgs.__new__(TmpGrafanaConfigArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(TmpGrafanaConfig, __self__).__init__(
            'tencentcloud:Monitor/tmpGrafanaConfig:TmpGrafanaConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'TmpGrafanaConfig':
        """
        Get an existing TmpGrafanaConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: JSON encoded string.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpGrafanaConfigState.__new__(_TmpGrafanaConfigState)

        __props__.__dict__["config"] = config
        __props__.__dict__["instance_id"] = instance_id
        return TmpGrafanaConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        JSON encoded string.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

