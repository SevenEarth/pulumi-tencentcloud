# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TmpCvmAgentArgs', 'TmpCvmAgent']

@pulumi.input_type
class TmpCvmAgentArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TmpCvmAgent resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] name: Agent name.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Agent name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TmpCvmAgentState:
    def __init__(__self__, *,
                 agent_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpCvmAgent resources.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] name: Agent name.
        """
        if agent_id is not None:
            pulumi.set(__self__, "agent_id", agent_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Agent name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class TmpCvmAgent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a monitor tmpCvmAgent

        ## Example Usage

        <!--Start PulumiCodeChooser -->
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
        foo_tmp_instance = tencentcloud.monitor.TmpInstance("fooTmpInstance",
            instance_name="tf-tmp-instance",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            data_retention_time=30,
            zone=availability_zone,
            tags={
                "createdBy": "terraform",
            })
        foo_tmp_cvm_agent = tencentcloud.monitor.TmpCvmAgent("fooTmpCvmAgent", instance_id=foo_tmp_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        monitor tmpCvmAgent can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent tmpCvmAgent instance_id#agent_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] name: Agent name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpCvmAgentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a monitor tmpCvmAgent

        ## Example Usage

        <!--Start PulumiCodeChooser -->
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
        foo_tmp_instance = tencentcloud.monitor.TmpInstance("fooTmpInstance",
            instance_name="tf-tmp-instance",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            data_retention_time=30,
            zone=availability_zone,
            tags={
                "createdBy": "terraform",
            })
        foo_tmp_cvm_agent = tencentcloud.monitor.TmpCvmAgent("fooTmpCvmAgent", instance_id=foo_tmp_instance.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        monitor tmpCvmAgent can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent tmpCvmAgent instance_id#agent_id
        ```

        :param str resource_name: The name of the resource.
        :param TmpCvmAgentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpCvmAgentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TmpCvmAgentArgs.__new__(TmpCvmAgentArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["agent_id"] = None
        super(TmpCvmAgent, __self__).__init__(
            'tencentcloud:Monitor/tmpCvmAgent:TmpCvmAgent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agent_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'TmpCvmAgent':
        """
        Get an existing TmpCvmAgent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] agent_id: Agent id.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] name: Agent name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpCvmAgentState.__new__(_TmpCvmAgentState)

        __props__.__dict__["agent_id"] = agent_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        return TmpCvmAgent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="agentId")
    def agent_id(self) -> pulumi.Output[str]:
        """
        Agent id.
        """
        return pulumi.get(self, "agent_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Agent name.
        """
        return pulumi.get(self, "name")

