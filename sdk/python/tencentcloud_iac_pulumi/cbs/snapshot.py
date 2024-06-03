# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnapshotArgs', 'Snapshot']

@pulumi.input_type
class SnapshotArgs:
    def __init__(__self__, *,
                 snapshot_name: pulumi.Input[str],
                 storage_id: pulumi.Input[str],
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Snapshot resource.
        :param pulumi.Input[str] snapshot_name: Name of the snapshot.
        :param pulumi.Input[str] storage_id: ID of the the CBS which this snapshot created from.
        :param pulumi.Input[Mapping[str, Any]] tags: cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        pulumi.set(__self__, "snapshot_name", snapshot_name)
        pulumi.set(__self__, "storage_id", storage_id)
        if tags is not None:
            warnings.warn("""cbs snapshot do not support tag now.""", DeprecationWarning)
            pulumi.log.warn("""tags is deprecated: cbs snapshot do not support tag now.""")
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Input[str]:
        """
        Name of the snapshot.
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_name", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Input[str]:
        """
        ID of the the CBS which this snapshot created from.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        warnings.warn("""cbs snapshot do not support tag now.""", DeprecationWarning)
        pulumi.log.warn("""tags is deprecated: cbs snapshot do not support tag now.""")

        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SnapshotState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 percent: Optional[pulumi.Input[int]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 snapshot_status: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Snapshot resources.
        :param pulumi.Input[str] create_time: Creation time of snapshot.
        :param pulumi.Input[str] disk_type: Types of CBS which this snapshot created from.
        :param pulumi.Input[int] percent: Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.
        :param pulumi.Input[str] snapshot_name: Name of the snapshot.
        :param pulumi.Input[str] snapshot_status: Status of the snapshot.
        :param pulumi.Input[str] storage_id: ID of the the CBS which this snapshot created from.
        :param pulumi.Input[int] storage_size: Volume of storage which this snapshot created from.
        :param pulumi.Input[Mapping[str, Any]] tags: cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if percent is not None:
            pulumi.set(__self__, "percent", percent)
        if snapshot_name is not None:
            pulumi.set(__self__, "snapshot_name", snapshot_name)
        if snapshot_status is not None:
            pulumi.set(__self__, "snapshot_status", snapshot_status)
        if storage_id is not None:
            pulumi.set(__self__, "storage_id", storage_id)
        if storage_size is not None:
            pulumi.set(__self__, "storage_size", storage_size)
        if tags is not None:
            warnings.warn("""cbs snapshot do not support tag now.""", DeprecationWarning)
            pulumi.log.warn("""tags is deprecated: cbs snapshot do not support tag now.""")
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of snapshot.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        Types of CBS which this snapshot created from.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter
    def percent(self) -> Optional[pulumi.Input[int]]:
        """
        Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.
        """
        return pulumi.get(self, "percent")

    @percent.setter
    def percent(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "percent", value)

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the snapshot.
        """
        return pulumi.get(self, "snapshot_name")

    @snapshot_name.setter
    def snapshot_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_name", value)

    @property
    @pulumi.getter(name="snapshotStatus")
    def snapshot_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the snapshot.
        """
        return pulumi.get(self, "snapshot_status")

    @snapshot_status.setter
    def snapshot_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_status", value)

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the the CBS which this snapshot created from.
        """
        return pulumi.get(self, "storage_id")

    @storage_id.setter
    def storage_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_id", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Volume of storage which this snapshot created from.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        warnings.warn("""cbs snapshot do not support tag now.""", DeprecationWarning)
        pulumi.log.warn("""tags is deprecated: cbs snapshot do not support tag now.""")

        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Snapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a CBS snapshot.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot = tencentcloud.cbs.Snapshot("snapshot",
            snapshot_name="unnamed",
            storage_id="disk-kdt0sq6m")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CBS snapshot can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cbs/snapshot:Snapshot snapshot snap-3sa3f39b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] snapshot_name: Name of the snapshot.
        :param pulumi.Input[str] storage_id: ID of the the CBS which this snapshot created from.
        :param pulumi.Input[Mapping[str, Any]] tags: cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CBS snapshot.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        snapshot = tencentcloud.cbs.Snapshot("snapshot",
            snapshot_name="unnamed",
            storage_id="disk-kdt0sq6m")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        CBS snapshot can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cbs/snapshot:Snapshot snapshot snap-3sa3f39b
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 storage_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotArgs.__new__(SnapshotArgs)

            if snapshot_name is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_name'")
            __props__.__dict__["snapshot_name"] = snapshot_name
            if storage_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_id'")
            __props__.__dict__["storage_id"] = storage_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["create_time"] = None
            __props__.__dict__["disk_type"] = None
            __props__.__dict__["percent"] = None
            __props__.__dict__["snapshot_status"] = None
            __props__.__dict__["storage_size"] = None
        super(Snapshot, __self__).__init__(
            'tencentcloud:Cbs/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            disk_type: Optional[pulumi.Input[str]] = None,
            percent: Optional[pulumi.Input[int]] = None,
            snapshot_name: Optional[pulumi.Input[str]] = None,
            snapshot_status: Optional[pulumi.Input[str]] = None,
            storage_id: Optional[pulumi.Input[str]] = None,
            storage_size: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creation time of snapshot.
        :param pulumi.Input[str] disk_type: Types of CBS which this snapshot created from.
        :param pulumi.Input[int] percent: Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.
        :param pulumi.Input[str] snapshot_name: Name of the snapshot.
        :param pulumi.Input[str] snapshot_status: Status of the snapshot.
        :param pulumi.Input[str] storage_id: ID of the the CBS which this snapshot created from.
        :param pulumi.Input[int] storage_size: Volume of storage which this snapshot created from.
        :param pulumi.Input[Mapping[str, Any]] tags: cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotState.__new__(_SnapshotState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["disk_type"] = disk_type
        __props__.__dict__["percent"] = percent
        __props__.__dict__["snapshot_name"] = snapshot_name
        __props__.__dict__["snapshot_status"] = snapshot_status
        __props__.__dict__["storage_id"] = storage_id
        __props__.__dict__["storage_size"] = storage_size
        __props__.__dict__["tags"] = tags
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of snapshot.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> pulumi.Output[str]:
        """
        Types of CBS which this snapshot created from.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter
    def percent(self) -> pulumi.Output[int]:
        """
        Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.
        """
        return pulumi.get(self, "percent")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> pulumi.Output[str]:
        """
        Name of the snapshot.
        """
        return pulumi.get(self, "snapshot_name")

    @property
    @pulumi.getter(name="snapshotStatus")
    def snapshot_status(self) -> pulumi.Output[str]:
        """
        Status of the snapshot.
        """
        return pulumi.get(self, "snapshot_status")

    @property
    @pulumi.getter(name="storageId")
    def storage_id(self) -> pulumi.Output[str]:
        """
        ID of the the CBS which this snapshot created from.
        """
        return pulumi.get(self, "storage_id")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Output[int]:
        """
        Volume of storage which this snapshot created from.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        cbs snapshot do not support tag now. The available tags within this CBS Snapshot.
        """
        warnings.warn("""cbs snapshot do not support tag now.""", DeprecationWarning)
        pulumi.log.warn("""tags is deprecated: cbs snapshot do not support tag now.""")

        return pulumi.get(self, "tags")

