# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TmpScrapeJobArgs', 'TmpScrapeJob']

@pulumi.input_type
class TmpScrapeJobArgs:
    def __init__(__self__, *,
                 agent_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 config: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TmpScrapeJob resource.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] config: Job content.
        """
        pulumi.set(__self__, "agent_id", agent_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if config is not None:
            pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Input[str]:
        """
        Agent id.
        """
        return pulumi.get(self, "agent_id")

    @agent_id.setter
    def agent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agent_id", value)

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

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        Job content.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)


@pulumi.input_type
class _TmpScrapeJobState:
    def __init__(__self__, *,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpScrapeJob resources.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] config: Job content.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        if agent_id is not None:
            pulumi.set(__self__, "agent_id", agent_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> Optional[pulumi.Input[str]]:
        """
        Agent id.
        """
        return pulumi.get(self, "agent_id")

    @agent_id.setter
    def agent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        Job content.
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


class TmpScrapeJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor tmpScrapeJob

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        tmp_scrape_job = tencentcloud.monitor.TmpScrapeJob("tmpScrapeJob",
            instance_id="prom-dko9d0nu",
            agent_id="agent-6a7g40k2",
            config=\"\"\"job_name: demo-config
        honor_timestamps: true
        metrics_path: /metrics
        scheme: https
        \"\"\")
        ```

        ## Import

        monitor tmpScrapeJob can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpScrapeJob:TmpScrapeJob tmpScrapeJob tmpScrapeJob_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] config: Job content.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpScrapeJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor tmpScrapeJob

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        tmp_scrape_job = tencentcloud.monitor.TmpScrapeJob("tmpScrapeJob",
            instance_id="prom-dko9d0nu",
            agent_id="agent-6a7g40k2",
            config=\"\"\"job_name: demo-config
        honor_timestamps: true
        metrics_path: /metrics
        scheme: https
        \"\"\")
        ```

        ## Import

        monitor tmpScrapeJob can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Monitor/tmpScrapeJob:TmpScrapeJob tmpScrapeJob tmpScrapeJob_id
        ```

        :param str resource_name: The name of the resource.
        :param TmpScrapeJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpScrapeJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TmpScrapeJobArgs.__new__(TmpScrapeJobArgs)

            if agent_id is None and not opts.urn:
                raise TypeError("Missing required property 'agent_id'")
            __props__.__dict__["agent_id"] = agent_id
            __props__.__dict__["config"] = config
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(TmpScrapeJob, __self__).__init__(
            'tencentcloud:Monitor/tmpScrapeJob:TmpScrapeJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_id: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'TmpScrapeJob':
        """
        Get an existing TmpScrapeJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] config: Job content.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpScrapeJobState.__new__(_TmpScrapeJobState)

        __props__.__dict__["agent_id"] = agent_id
        __props__.__dict__["config"] = config
        __props__.__dict__["instance_id"] = instance_id
        return TmpScrapeJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Output[str]:
        """
        Agent id.
        """
        return pulumi.get(self, "agent_id")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[str]]:
        """
        Job content.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

