# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StorageSetArgs', 'StorageSet']

@pulumi.input_type
class StorageSetArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 storage_name: pulumi.Input[str],
                 storage_size: pulumi.Input[int],
                 storage_type: pulumi.Input[str],
                 charge_type: Optional[pulumi.Input[str]] = None,
                 disk_count: Optional[pulumi.Input[int]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 throughput_performance: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a StorageSet resource.
        :param pulumi.Input[str] availability_zone: The available zone that the CBS instance locates at.
        :param pulumi.Input[str] storage_name: Name of CBS. The maximum length can not exceed 60 bytes.
        :param pulumi.Input[int] storage_size: Volume of CBS, and unit is GB.
        :param pulumi.Input[str] storage_type: Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        :param pulumi.Input[str] charge_type: The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        :param pulumi.Input[int] disk_count: The number of disks to be purchased. Default 1.
        :param pulumi.Input[bool] encrypt: Indicates whether CBS is encrypted.
        :param pulumi.Input[int] project_id: ID of the project to which the instance belongs.
        :param pulumi.Input[str] snapshot_id: ID of the snapshot. If specified, created the CBS by this snapshot.
        :param pulumi.Input[int] throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        pulumi.set(__self__, "storage_name", storage_name)
        pulumi.set(__self__, "storage_size", storage_size)
        pulumi.set(__self__, "storage_type", storage_type)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if disk_count is not None:
            pulumi.set(__self__, "disk_count", disk_count)
        if encrypt is not None:
            pulumi.set(__self__, "encrypt", encrypt)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if throughput_performance is not None:
            pulumi.set(__self__, "throughput_performance", throughput_performance)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> pulumi.Input[str]:
        """
        Name of CBS. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "storage_name")

    @storage_name.setter
    def storage_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_name", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Input[int]:
        """
        Volume of CBS, and unit is GB.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Input[str]:
        """
        Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="diskCount")
    def disk_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of disks to be purchased. Default 1.
        """
        return pulumi.get(self, "disk_count")

    @disk_count.setter
    def disk_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_count", value)

    @property
    @pulumi.getter
    def encrypt(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether CBS is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @encrypt.setter
    def encrypt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypt", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID of the project to which the instance belongs.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the snapshot. If specified, created the CBS by this snapshot.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> Optional[pulumi.Input[int]]:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")

    @throughput_performance.setter
    def throughput_performance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "throughput_performance", value)


@pulumi.input_type
class _StorageSetState:
    def __init__(__self__, *,
                 attached: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 disk_count: Optional[pulumi.Input[int]] = None,
                 disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 storage_name: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 storage_status: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 throughput_performance: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering StorageSet resources.
        :param pulumi.Input[bool] attached: Indicates whether the CBS is mounted the CVM.
        :param pulumi.Input[str] availability_zone: The available zone that the CBS instance locates at.
        :param pulumi.Input[str] charge_type: The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        :param pulumi.Input[int] disk_count: The number of disks to be purchased. Default 1.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: disk id list.
        :param pulumi.Input[bool] encrypt: Indicates whether CBS is encrypted.
        :param pulumi.Input[int] project_id: ID of the project to which the instance belongs.
        :param pulumi.Input[str] snapshot_id: ID of the snapshot. If specified, created the CBS by this snapshot.
        :param pulumi.Input[str] storage_name: Name of CBS. The maximum length can not exceed 60 bytes.
        :param pulumi.Input[int] storage_size: Volume of CBS, and unit is GB.
        :param pulumi.Input[str] storage_status: Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
        :param pulumi.Input[str] storage_type: Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        :param pulumi.Input[int] throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        if attached is not None:
            pulumi.set(__self__, "attached", attached)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if disk_count is not None:
            pulumi.set(__self__, "disk_count", disk_count)
        if disk_ids is not None:
            pulumi.set(__self__, "disk_ids", disk_ids)
        if encrypt is not None:
            pulumi.set(__self__, "encrypt", encrypt)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if storage_name is not None:
            pulumi.set(__self__, "storage_name", storage_name)
        if storage_size is not None:
            pulumi.set(__self__, "storage_size", storage_size)
        if storage_status is not None:
            pulumi.set(__self__, "storage_status", storage_status)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if throughput_performance is not None:
            pulumi.set(__self__, "throughput_performance", throughput_performance)

    @property
    @pulumi.getter
    def attached(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the CBS is mounted the CVM.
        """
        return pulumi.get(self, "attached")

    @attached.setter
    def attached(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "attached", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="diskCount")
    def disk_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of disks to be purchased. Default 1.
        """
        return pulumi.get(self, "disk_count")

    @disk_count.setter
    def disk_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_count", value)

    @property
    @pulumi.getter(name="diskIds")
    def disk_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        disk id list.
        """
        return pulumi.get(self, "disk_ids")

    @disk_ids.setter
    def disk_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "disk_ids", value)

    @property
    @pulumi.getter
    def encrypt(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether CBS is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @encrypt.setter
    def encrypt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "encrypt", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID of the project to which the instance belongs.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the snapshot. If specified, created the CBS by this snapshot.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of CBS. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "storage_name")

    @storage_name.setter
    def storage_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_name", value)

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> Optional[pulumi.Input[int]]:
        """
        Volume of CBS, and unit is GB.
        """
        return pulumi.get(self, "storage_size")

    @storage_size.setter
    def storage_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_size", value)

    @property
    @pulumi.getter(name="storageStatus")
    def storage_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
        """
        return pulumi.get(self, "storage_status")

    @storage_status.setter
    def storage_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_status", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> Optional[pulumi.Input[int]]:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")

    @throughput_performance.setter
    def throughput_performance(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "throughput_performance", value)


class StorageSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 disk_count: Optional[pulumi.Input[int]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 storage_name: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 throughput_performance: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create CBS set.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        storage = tencentcloud.cbs.StorageSet("storage",
            availability_zone="ap-guangzhou-3",
            disk_count=10,
            encrypt=False,
            project_id=0,
            storage_name="mystorage",
            storage_size=100,
            storage_type="CLOUD_SSD")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The available zone that the CBS instance locates at.
        :param pulumi.Input[str] charge_type: The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        :param pulumi.Input[int] disk_count: The number of disks to be purchased. Default 1.
        :param pulumi.Input[bool] encrypt: Indicates whether CBS is encrypted.
        :param pulumi.Input[int] project_id: ID of the project to which the instance belongs.
        :param pulumi.Input[str] snapshot_id: ID of the snapshot. If specified, created the CBS by this snapshot.
        :param pulumi.Input[str] storage_name: Name of CBS. The maximum length can not exceed 60 bytes.
        :param pulumi.Input[int] storage_size: Volume of CBS, and unit is GB.
        :param pulumi.Input[str] storage_type: Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        :param pulumi.Input[int] throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create CBS set.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        storage = tencentcloud.cbs.StorageSet("storage",
            availability_zone="ap-guangzhou-3",
            disk_count=10,
            encrypt=False,
            project_id=0,
            storage_name="mystorage",
            storage_size=100,
            storage_type="CLOUD_SSD")
        ```

        :param str resource_name: The name of the resource.
        :param StorageSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 disk_count: Optional[pulumi.Input[int]] = None,
                 encrypt: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 storage_name: Optional[pulumi.Input[str]] = None,
                 storage_size: Optional[pulumi.Input[int]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 throughput_performance: Optional[pulumi.Input[int]] = None,
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
            __props__ = StorageSetArgs.__new__(StorageSetArgs)

            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["charge_type"] = charge_type
            __props__.__dict__["disk_count"] = disk_count
            __props__.__dict__["encrypt"] = encrypt
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["snapshot_id"] = snapshot_id
            if storage_name is None and not opts.urn:
                raise TypeError("Missing required property 'storage_name'")
            __props__.__dict__["storage_name"] = storage_name
            if storage_size is None and not opts.urn:
                raise TypeError("Missing required property 'storage_size'")
            __props__.__dict__["storage_size"] = storage_size
            if storage_type is None and not opts.urn:
                raise TypeError("Missing required property 'storage_type'")
            __props__.__dict__["storage_type"] = storage_type
            __props__.__dict__["throughput_performance"] = throughput_performance
            __props__.__dict__["attached"] = None
            __props__.__dict__["disk_ids"] = None
            __props__.__dict__["storage_status"] = None
        super(StorageSet, __self__).__init__(
            'tencentcloud:Cbs/storageSet:StorageSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attached: Optional[pulumi.Input[bool]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            disk_count: Optional[pulumi.Input[int]] = None,
            disk_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            encrypt: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            storage_name: Optional[pulumi.Input[str]] = None,
            storage_size: Optional[pulumi.Input[int]] = None,
            storage_status: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            throughput_performance: Optional[pulumi.Input[int]] = None) -> 'StorageSet':
        """
        Get an existing StorageSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] attached: Indicates whether the CBS is mounted the CVM.
        :param pulumi.Input[str] availability_zone: The available zone that the CBS instance locates at.
        :param pulumi.Input[str] charge_type: The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        :param pulumi.Input[int] disk_count: The number of disks to be purchased. Default 1.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] disk_ids: disk id list.
        :param pulumi.Input[bool] encrypt: Indicates whether CBS is encrypted.
        :param pulumi.Input[int] project_id: ID of the project to which the instance belongs.
        :param pulumi.Input[str] snapshot_id: ID of the snapshot. If specified, created the CBS by this snapshot.
        :param pulumi.Input[str] storage_name: Name of CBS. The maximum length can not exceed 60 bytes.
        :param pulumi.Input[int] storage_size: Volume of CBS, and unit is GB.
        :param pulumi.Input[str] storage_status: Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
        :param pulumi.Input[str] storage_type: Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        :param pulumi.Input[int] throughput_performance: Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageSetState.__new__(_StorageSetState)

        __props__.__dict__["attached"] = attached
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["disk_count"] = disk_count
        __props__.__dict__["disk_ids"] = disk_ids
        __props__.__dict__["encrypt"] = encrypt
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["storage_name"] = storage_name
        __props__.__dict__["storage_size"] = storage_size
        __props__.__dict__["storage_status"] = storage_status
        __props__.__dict__["storage_type"] = storage_type
        __props__.__dict__["throughput_performance"] = throughput_performance
        return StorageSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attached(self) -> pulumi.Output[bool]:
        """
        Indicates whether the CBS is mounted the CVM.
        """
        return pulumi.get(self, "attached")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The available zone that the CBS instance locates at.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        The charge type of CBS instance. Only support `POSTPAID_BY_HOUR`.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="diskCount")
    def disk_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of disks to be purchased. Default 1.
        """
        return pulumi.get(self, "disk_count")

    @property
    @pulumi.getter(name="diskIds")
    def disk_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        disk id list.
        """
        return pulumi.get(self, "disk_ids")

    @property
    @pulumi.getter
    def encrypt(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether CBS is encrypted.
        """
        return pulumi.get(self, "encrypt")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[int]]:
        """
        ID of the project to which the instance belongs.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[str]:
        """
        ID of the snapshot. If specified, created the CBS by this snapshot.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="storageName")
    def storage_name(self) -> pulumi.Output[str]:
        """
        Name of CBS. The maximum length can not exceed 60 bytes.
        """
        return pulumi.get(self, "storage_name")

    @property
    @pulumi.getter(name="storageSize")
    def storage_size(self) -> pulumi.Output[int]:
        """
        Volume of CBS, and unit is GB.
        """
        return pulumi.get(self, "storage_size")

    @property
    @pulumi.getter(name="storageStatus")
    def storage_status(self) -> pulumi.Output[str]:
        """
        Status of CBS. Valid values: UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING.
        """
        return pulumi.get(self, "storage_status")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[str]:
        """
        Type of CBS medium. Valid values: CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_BSSD: General Purpose SSD, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="throughputPerformance")
    def throughput_performance(self) -> pulumi.Output[Optional[int]]:
        """
        Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
        """
        return pulumi.get(self, "throughput_performance")

