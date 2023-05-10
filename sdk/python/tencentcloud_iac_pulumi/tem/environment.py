# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 environment_name: pulumi.Input[str],
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[str] environment_name: environment name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: subnet IDs.
        :param pulumi.Input[str] vpc: vpc ID.
        :param pulumi.Input[str] description: environment description.
        :param pulumi.Input[Mapping[str, Any]] tags: environment tag list.
        """
        pulumi.set(__self__, "environment_name", environment_name)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc", vpc)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Input[str]:
        """
        environment name.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        subnet IDs.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter
    def vpc(self) -> pulumi.Input[str]:
        """
        vpc ID.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        environment description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        environment tag list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EnvironmentState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Environment resources.
        :param pulumi.Input[str] description: environment description.
        :param pulumi.Input[str] environment_name: environment name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: subnet IDs.
        :param pulumi.Input[Mapping[str, Any]] tags: environment tag list.
        :param pulumi.Input[str] vpc: vpc ID.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment_name is not None:
            pulumi.set(__self__, "environment_name", environment_name)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc is not None:
            pulumi.set(__self__, "vpc", vpc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        environment description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[pulumi.Input[str]]:
        """
        environment name.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        subnet IDs.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        environment tag list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def vpc(self) -> Optional[pulumi.Input[str]]:
        """
        vpc ID.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc", value)


class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tem environment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        environment = tencentcloud.tem.Environment("environment",
            description="demo for test",
            environment_name="demo",
            subnet_ids=[
                "subnet-rdkj0agk",
                "subnet-r1c4pn5m",
                "subnet-02hcj95c",
            ],
            tags={
                "created": "terraform",
            },
            vpc="vpc-2hfyray3")
        ```

        ## Import

        tem environment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tem/environment:Environment environment environment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: environment description.
        :param pulumi.Input[str] environment_name: environment name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: subnet IDs.
        :param pulumi.Input[Mapping[str, Any]] tags: environment tag list.
        :param pulumi.Input[str] vpc: vpc ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tem environment

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        environment = tencentcloud.tem.Environment("environment",
            description="demo for test",
            environment_name="demo",
            subnet_ids=[
                "subnet-rdkj0agk",
                "subnet-r1c4pn5m",
                "subnet-02hcj95c",
            ],
            tags={
                "created": "terraform",
            },
            vpc="vpc-2hfyray3")
        ```

        ## Import

        tem environment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tem/environment:Environment environment environment_id
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment_name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc: Optional[pulumi.Input[str]] = None,
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
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["description"] = description
            if environment_name is None and not opts.urn:
                raise TypeError("Missing required property 'environment_name'")
            __props__.__dict__["environment_name"] = environment_name
            if subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["tags"] = tags
            if vpc is None and not opts.urn:
                raise TypeError("Missing required property 'vpc'")
            __props__.__dict__["vpc"] = vpc
        super(Environment, __self__).__init__(
            'tencentcloud:Tem/environment:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            environment_name: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc: Optional[pulumi.Input[str]] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: environment description.
        :param pulumi.Input[str] environment_name: environment name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: subnet IDs.
        :param pulumi.Input[Mapping[str, Any]] tags: environment tag list.
        :param pulumi.Input[str] vpc: vpc ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentState.__new__(_EnvironmentState)

        __props__.__dict__["description"] = description
        __props__.__dict__["environment_name"] = environment_name
        __props__.__dict__["subnet_ids"] = subnet_ids
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc"] = vpc
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        environment description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Output[str]:
        """
        environment name.
        """
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        subnet IDs.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        environment tag list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vpc(self) -> pulumi.Output[str]:
        """
        vpc ID.
        """
        return pulumi.get(self, "vpc")

