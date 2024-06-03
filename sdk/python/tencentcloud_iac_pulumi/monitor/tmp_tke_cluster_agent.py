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

__all__ = ['TmpTkeClusterAgentArgs', 'TmpTkeClusterAgent']

@pulumi.input_type
class TmpTkeClusterAgentArgs:
    def __init__(__self__, *,
                 agents: pulumi.Input['TmpTkeClusterAgentAgentsArgs'],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a TmpTkeClusterAgent resource.
        :param pulumi.Input['TmpTkeClusterAgentAgentsArgs'] agents: agent list.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        pulumi.set(__self__, "agents", agents)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def agents(self) -> pulumi.Input['TmpTkeClusterAgentAgentsArgs']:
        """
        agent list.
        """
        return pulumi.get(self, "agents")

    @agents.setter
    def agents(self, value: pulumi.Input['TmpTkeClusterAgentAgentsArgs']):
        pulumi.set(self, "agents", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _TmpTkeClusterAgentState:
    def __init__(__self__, *,
                 agents: Optional[pulumi.Input['TmpTkeClusterAgentAgentsArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TmpTkeClusterAgent resources.
        :param pulumi.Input['TmpTkeClusterAgentAgentsArgs'] agents: agent list.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        if agents is not None:
            pulumi.set(__self__, "agents", agents)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def agents(self) -> Optional[pulumi.Input['TmpTkeClusterAgentAgentsArgs']]:
        """
        agent list.
        """
        return pulumi.get(self, "agents")

    @agents.setter
    def agents(self, value: Optional[pulumi.Input['TmpTkeClusterAgentAgentsArgs']]):
        pulumi.set(self, "agents", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class TmpTkeClusterAgent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents: Optional[pulumi.Input[pulumi.InputType['TmpTkeClusterAgentAgentsArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tmp tke cluster agent

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TmpTkeClusterAgentAgentsArgs']] agents: agent list.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TmpTkeClusterAgentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tmp tke cluster agent

        :param str resource_name: The name of the resource.
        :param TmpTkeClusterAgentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TmpTkeClusterAgentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents: Optional[pulumi.Input[pulumi.InputType['TmpTkeClusterAgentAgentsArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TmpTkeClusterAgentArgs.__new__(TmpTkeClusterAgentArgs)

            if agents is None and not opts.urn:
                raise TypeError("Missing required property 'agents'")
            __props__.__dict__["agents"] = agents
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(TmpTkeClusterAgent, __self__).__init__(
            'tencentcloud:Monitor/tmpTkeClusterAgent:TmpTkeClusterAgent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            agents: Optional[pulumi.Input[pulumi.InputType['TmpTkeClusterAgentAgentsArgs']]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'TmpTkeClusterAgent':
        """
        Get an existing TmpTkeClusterAgent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TmpTkeClusterAgentAgentsArgs']] agents: agent list.
        :param pulumi.Input[str] instance_id: Instance Id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TmpTkeClusterAgentState.__new__(_TmpTkeClusterAgentState)

        __props__.__dict__["agents"] = agents
        __props__.__dict__["instance_id"] = instance_id
        return TmpTkeClusterAgent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def agents(self) -> pulumi.Output['outputs.TmpTkeClusterAgentAgents']:
        """
        agent list.
        """
        return pulumi.get(self, "agents")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance Id.
        """
        return pulumi.get(self, "instance_id")

