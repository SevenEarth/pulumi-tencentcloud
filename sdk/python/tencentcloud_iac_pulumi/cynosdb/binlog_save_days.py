# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BinlogSaveDaysArgs', 'BinlogSaveDays']

@pulumi.input_type
class BinlogSaveDaysArgs:
    def __init__(__self__, *,
                 binlog_save_days: pulumi.Input[int],
                 cluster_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a BinlogSaveDays resource.
        :param pulumi.Input[int] binlog_save_days: Binlog retention days.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        """
        pulumi.set(__self__, "binlog_save_days", binlog_save_days)
        pulumi.set(__self__, "cluster_id", cluster_id)

    @property
    @pulumi.getter(name="binlogSaveDays")
    def binlog_save_days(self) -> pulumi.Input[int]:
        """
        Binlog retention days.
        """
        return pulumi.get(self, "binlog_save_days")

    @binlog_save_days.setter
    def binlog_save_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "binlog_save_days", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)


@pulumi.input_type
class _BinlogSaveDaysState:
    def __init__(__self__, *,
                 binlog_save_days: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BinlogSaveDays resources.
        :param pulumi.Input[int] binlog_save_days: Binlog retention days.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        """
        if binlog_save_days is not None:
            pulumi.set(__self__, "binlog_save_days", binlog_save_days)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)

    @property
    @pulumi.getter(name="binlogSaveDays")
    def binlog_save_days(self) -> Optional[pulumi.Input[int]]:
        """
        Binlog retention days.
        """
        return pulumi.get(self, "binlog_save_days")

    @binlog_save_days.setter
    def binlog_save_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "binlog_save_days", value)

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


class BinlogSaveDays(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binlog_save_days: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a cynosdb binlog_save_days

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        binlog_save_days = tencentcloud.cynosdb.BinlogSaveDays("binlogSaveDays",
            binlog_save_days=7,
            cluster_id="cynosdbmysql-123")
        ```

        ## Import

        cynosdb binlog_save_days can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays binlog_save_days binlog_save_days_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] binlog_save_days: Binlog retention days.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BinlogSaveDaysArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cynosdb binlog_save_days

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        binlog_save_days = tencentcloud.cynosdb.BinlogSaveDays("binlogSaveDays",
            binlog_save_days=7,
            cluster_id="cynosdbmysql-123")
        ```

        ## Import

        cynosdb binlog_save_days can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays binlog_save_days binlog_save_days_id
        ```

        :param str resource_name: The name of the resource.
        :param BinlogSaveDaysArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BinlogSaveDaysArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 binlog_save_days: Optional[pulumi.Input[int]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = BinlogSaveDaysArgs.__new__(BinlogSaveDaysArgs)

            if binlog_save_days is None and not opts.urn:
                raise TypeError("Missing required property 'binlog_save_days'")
            __props__.__dict__["binlog_save_days"] = binlog_save_days
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
        super(BinlogSaveDays, __self__).__init__(
            'tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            binlog_save_days: Optional[pulumi.Input[int]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None) -> 'BinlogSaveDays':
        """
        Get an existing BinlogSaveDays resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] binlog_save_days: Binlog retention days.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BinlogSaveDaysState.__new__(_BinlogSaveDaysState)

        __props__.__dict__["binlog_save_days"] = binlog_save_days
        __props__.__dict__["cluster_id"] = cluster_id
        return BinlogSaveDays(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="binlogSaveDays")
    def binlog_save_days(self) -> pulumi.Output[int]:
        """
        Binlog retention days.
        """
        return pulumi.get(self, "binlog_save_days")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

