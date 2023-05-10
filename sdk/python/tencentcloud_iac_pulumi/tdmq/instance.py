# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 bind_cluster_id: Optional[pulumi.Input[int]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] cluster_name: The name of tdmq cluster to be created.
        :param pulumi.Input[int] bind_cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] remark: Description of the tdmq cluster.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        if bind_cluster_id is not None:
            pulumi.set(__self__, "bind_cluster_id", bind_cluster_id)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The name of tdmq cluster to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="bindClusterId")
    def bind_cluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "bind_cluster_id")

    @bind_cluster_id.setter
    def bind_cluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bind_cluster_id", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the tdmq cluster.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

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


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 bind_cluster_id: Optional[pulumi.Input[int]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[int] bind_cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] cluster_name: The name of tdmq cluster to be created.
        :param pulumi.Input[str] remark: Description of the tdmq cluster.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if bind_cluster_id is not None:
            pulumi.set(__self__, "bind_cluster_id", bind_cluster_id)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="bindClusterId")
    def bind_cluster_id(self) -> Optional[pulumi.Input[int]]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "bind_cluster_id")

    @bind_cluster_id.setter
    def bind_cluster_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bind_cluster_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of tdmq cluster to be created.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the tdmq cluster.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

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


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bind_cluster_id: Optional[pulumi.Input[int]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provide a resource to create a TDMQ instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.tdmq.Instance("foo",
            cluster_name="example",
            remark="this is description.",
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        Tdmq instance can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/instance:Instance test tdmq_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bind_cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] cluster_name: The name of tdmq cluster to be created.
        :param pulumi.Input[str] remark: Description of the tdmq cluster.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a TDMQ instance.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.tdmq.Instance("foo",
            cluster_name="example",
            remark="this is description.",
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        Tdmq instance can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/instance:Instance test tdmq_id
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
                 bind_cluster_id: Optional[pulumi.Input[int]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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

            __props__.__dict__["bind_cluster_id"] = bind_cluster_id
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["remark"] = remark
            __props__.__dict__["tags"] = tags
        super(Instance, __self__).__init__(
            'tencentcloud:Tdmq/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bind_cluster_id: Optional[pulumi.Input[int]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bind_cluster_id: The Dedicated Cluster Id.
        :param pulumi.Input[str] cluster_name: The name of tdmq cluster to be created.
        :param pulumi.Input[str] remark: Description of the tdmq cluster.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["bind_cluster_id"] = bind_cluster_id
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["remark"] = remark
        __props__.__dict__["tags"] = tags
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bindClusterId")
    def bind_cluster_id(self) -> pulumi.Output[Optional[int]]:
        """
        The Dedicated Cluster Id.
        """
        return pulumi.get(self, "bind_cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of tdmq cluster to be created.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the tdmq cluster.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

