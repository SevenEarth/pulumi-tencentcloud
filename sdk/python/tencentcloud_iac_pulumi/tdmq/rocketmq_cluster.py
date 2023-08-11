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

__all__ = ['RocketmqClusterArgs', 'RocketmqCluster']

@pulumi.input_type
class RocketmqClusterArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 remark: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RocketmqCluster resource.
        :param pulumi.Input[str] cluster_name: Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] remark: Cluster description (up to 128 characters).
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster description (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)


@pulumi.input_type
class _RocketmqClusterState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[int]] = None,
                 is_vip: Optional[pulumi.Input[bool]] = None,
                 public_end_point: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 rocket_mq_flag: Optional[pulumi.Input[bool]] = None,
                 support_namespace_endpoint: Optional[pulumi.Input[bool]] = None,
                 vpc_end_point: Optional[pulumi.Input[str]] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqClusterVpcArgs']]]] = None):
        """
        Input properties used for looking up and filtering RocketmqCluster resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] cluster_name: Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[bool] is_vip: Whether it is an exclusive instance.
        :param pulumi.Input[str] public_end_point: Public network access address.
        :param pulumi.Input[str] region: Region information.
        :param pulumi.Input[str] remark: Cluster description (up to 128 characters).
        :param pulumi.Input[bool] rocket_mq_flag: Rocketmq cluster identification.
        :param pulumi.Input[bool] support_namespace_endpoint: Whether the namespace access point is supported.
        :param pulumi.Input[str] vpc_end_point: VPC access address.
        :param pulumi.Input[Sequence[pulumi.Input['RocketmqClusterVpcArgs']]] vpcs: Vpc list.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if is_vip is not None:
            pulumi.set(__self__, "is_vip", is_vip)
        if public_end_point is not None:
            pulumi.set(__self__, "public_end_point", public_end_point)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if rocket_mq_flag is not None:
            pulumi.set(__self__, "rocket_mq_flag", rocket_mq_flag)
        if support_namespace_endpoint is not None:
            pulumi.set(__self__, "support_namespace_endpoint", support_namespace_endpoint)
        if vpc_end_point is not None:
            pulumi.set(__self__, "vpc_end_point", vpc_end_point)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[int]]:
        """
        Creation time in milliseconds.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="isVip")
    def is_vip(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether it is an exclusive instance.
        """
        return pulumi.get(self, "is_vip")

    @is_vip.setter
    def is_vip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_vip", value)

    @property
    @pulumi.getter(name="publicEndPoint")
    def public_end_point(self) -> Optional[pulumi.Input[str]]:
        """
        Public network access address.
        """
        return pulumi.get(self, "public_end_point")

    @public_end_point.setter
    def public_end_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_end_point", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region information.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster description (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="rocketMQFlag")
    def rocket_mq_flag(self) -> Optional[pulumi.Input[bool]]:
        """
        Rocketmq cluster identification.
        """
        return pulumi.get(self, "rocket_mq_flag")

    @rocket_mq_flag.setter
    def rocket_mq_flag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rocket_mq_flag", value)

    @property
    @pulumi.getter(name="supportNamespaceEndpoint")
    def support_namespace_endpoint(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the namespace access point is supported.
        """
        return pulumi.get(self, "support_namespace_endpoint")

    @support_namespace_endpoint.setter
    def support_namespace_endpoint(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "support_namespace_endpoint", value)

    @property
    @pulumi.getter(name="vpcEndPoint")
    def vpc_end_point(self) -> Optional[pulumi.Input[str]]:
        """
        VPC access address.
        """
        return pulumi.get(self, "vpc_end_point")

    @vpc_end_point.setter
    def vpc_end_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_end_point", value)

    @property
    @pulumi.getter
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqClusterVpcArgs']]]]:
        """
        Vpc list.
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RocketmqClusterVpcArgs']]]]):
        pulumi.set(self, "vpcs", value)


class RocketmqCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tdmqRocketmq cluster

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.tdmq.RocketmqCluster("example",
            cluster_name="tf_example",
            remark="remark.")
        ```

        ## Import

        tdmqRocketmq cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster cluster cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[str] remark: Cluster description (up to 128 characters).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RocketmqClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tdmqRocketmq cluster

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.tdmq.RocketmqCluster("example",
            cluster_name="tf_example",
            remark="remark.")
        ```

        ## Import

        tdmqRocketmq cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster cluster cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param RocketmqClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RocketmqClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
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
            __props__ = RocketmqClusterArgs.__new__(RocketmqClusterArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["remark"] = remark
            __props__.__dict__["cluster_id"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["is_vip"] = None
            __props__.__dict__["public_end_point"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["rocket_mq_flag"] = None
            __props__.__dict__["support_namespace_endpoint"] = None
            __props__.__dict__["vpc_end_point"] = None
            __props__.__dict__["vpcs"] = None
        super(RocketmqCluster, __self__).__init__(
            'tencentcloud:Tdmq/rocketmqCluster:RocketmqCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[int]] = None,
            is_vip: Optional[pulumi.Input[bool]] = None,
            public_end_point: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            rocket_mq_flag: Optional[pulumi.Input[bool]] = None,
            support_namespace_endpoint: Optional[pulumi.Input[bool]] = None,
            vpc_end_point: Optional[pulumi.Input[str]] = None,
            vpcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqClusterVpcArgs']]]]] = None) -> 'RocketmqCluster':
        """
        Get an existing RocketmqCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] cluster_name: Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        :param pulumi.Input[int] create_time: Creation time in milliseconds.
        :param pulumi.Input[bool] is_vip: Whether it is an exclusive instance.
        :param pulumi.Input[str] public_end_point: Public network access address.
        :param pulumi.Input[str] region: Region information.
        :param pulumi.Input[str] remark: Cluster description (up to 128 characters).
        :param pulumi.Input[bool] rocket_mq_flag: Rocketmq cluster identification.
        :param pulumi.Input[bool] support_namespace_endpoint: Whether the namespace access point is supported.
        :param pulumi.Input[str] vpc_end_point: VPC access address.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RocketmqClusterVpcArgs']]]] vpcs: Vpc list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RocketmqClusterState.__new__(_RocketmqClusterState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["is_vip"] = is_vip
        __props__.__dict__["public_end_point"] = public_end_point
        __props__.__dict__["region"] = region
        __props__.__dict__["remark"] = remark
        __props__.__dict__["rocket_mq_flag"] = rocket_mq_flag
        __props__.__dict__["support_namespace_endpoint"] = support_namespace_endpoint
        __props__.__dict__["vpc_end_point"] = vpc_end_point
        __props__.__dict__["vpcs"] = vpcs
        return RocketmqCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Cluster name, which can contain 3-64 letters, digits, hyphens, and underscores.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[int]:
        """
        Creation time in milliseconds.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="isVip")
    def is_vip(self) -> pulumi.Output[bool]:
        """
        Whether it is an exclusive instance.
        """
        return pulumi.get(self, "is_vip")

    @property
    @pulumi.getter(name="publicEndPoint")
    def public_end_point(self) -> pulumi.Output[str]:
        """
        Public network access address.
        """
        return pulumi.get(self, "public_end_point")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region information.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        Cluster description (up to 128 characters).
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="rocketMQFlag")
    def rocket_mq_flag(self) -> pulumi.Output[bool]:
        """
        Rocketmq cluster identification.
        """
        return pulumi.get(self, "rocket_mq_flag")

    @property
    @pulumi.getter(name="supportNamespaceEndpoint")
    def support_namespace_endpoint(self) -> pulumi.Output[bool]:
        """
        Whether the namespace access point is supported.
        """
        return pulumi.get(self, "support_namespace_endpoint")

    @property
    @pulumi.getter(name="vpcEndPoint")
    def vpc_end_point(self) -> pulumi.Output[str]:
        """
        VPC access address.
        """
        return pulumi.get(self, "vpc_end_point")

    @property
    @pulumi.getter
    def vpcs(self) -> pulumi.Output[Sequence['outputs.RocketmqClusterVpc']]:
        """
        Vpc list.
        """
        return pulumi.get(self, "vpcs")

