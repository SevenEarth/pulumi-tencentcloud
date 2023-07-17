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

__all__ = ['RenewDiskArgs', 'RenewDisk']

@pulumi.input_type
class RenewDiskArgs:
    def __init__(__self__, *,
                 disk_id: pulumi.Input[str],
                 renew_disk_charge_prepaid: pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs'],
                 auto_voucher: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RenewDisk resource.
        :param pulumi.Input[str] disk_id: List of disk ID.
        :param pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs'] renew_disk_charge_prepaid: Renew cloud hard disk subscription related parameter settings.
        :param pulumi.Input[bool] auto_voucher: Whether to automatically use the voucher. Not used by default.
        """
        pulumi.set(__self__, "disk_id", disk_id)
        pulumi.set(__self__, "renew_disk_charge_prepaid", renew_disk_charge_prepaid)
        if auto_voucher is not None:
            pulumi.set(__self__, "auto_voucher", auto_voucher)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Input[str]:
        """
        List of disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="renewDiskChargePrepaid")
    def renew_disk_charge_prepaid(self) -> pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs']:
        """
        Renew cloud hard disk subscription related parameter settings.
        """
        return pulumi.get(self, "renew_disk_charge_prepaid")

    @renew_disk_charge_prepaid.setter
    def renew_disk_charge_prepaid(self, value: pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs']):
        pulumi.set(self, "renew_disk_charge_prepaid", value)

    @property
    @pulumi.getter(name="autoVoucher")
    def auto_voucher(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically use the voucher. Not used by default.
        """
        return pulumi.get(self, "auto_voucher")

    @auto_voucher.setter
    def auto_voucher(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_voucher", value)


@pulumi.input_type
class _RenewDiskState:
    def __init__(__self__, *,
                 auto_voucher: Optional[pulumi.Input[bool]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 renew_disk_charge_prepaid: Optional[pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs']] = None):
        """
        Input properties used for looking up and filtering RenewDisk resources.
        :param pulumi.Input[bool] auto_voucher: Whether to automatically use the voucher. Not used by default.
        :param pulumi.Input[str] disk_id: List of disk ID.
        :param pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs'] renew_disk_charge_prepaid: Renew cloud hard disk subscription related parameter settings.
        """
        if auto_voucher is not None:
            pulumi.set(__self__, "auto_voucher", auto_voucher)
        if disk_id is not None:
            pulumi.set(__self__, "disk_id", disk_id)
        if renew_disk_charge_prepaid is not None:
            pulumi.set(__self__, "renew_disk_charge_prepaid", renew_disk_charge_prepaid)

    @property
    @pulumi.getter(name="autoVoucher")
    def auto_voucher(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically use the voucher. Not used by default.
        """
        return pulumi.get(self, "auto_voucher")

    @auto_voucher.setter
    def auto_voucher(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_voucher", value)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> Optional[pulumi.Input[str]]:
        """
        List of disk ID.
        """
        return pulumi.get(self, "disk_id")

    @disk_id.setter
    def disk_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_id", value)

    @property
    @pulumi.getter(name="renewDiskChargePrepaid")
    def renew_disk_charge_prepaid(self) -> Optional[pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs']]:
        """
        Renew cloud hard disk subscription related parameter settings.
        """
        return pulumi.get(self, "renew_disk_charge_prepaid")

    @renew_disk_charge_prepaid.setter
    def renew_disk_charge_prepaid(self, value: Optional[pulumi.Input['RenewDiskRenewDiskChargePrepaidArgs']]):
        pulumi.set(self, "renew_disk_charge_prepaid", value)


class RenewDisk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_voucher: Optional[pulumi.Input[bool]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 renew_disk_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewDiskRenewDiskChargePrepaidArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a lighthouse renew_disk

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        renew_disk = tencentcloud.lighthouse.RenewDisk("renewDisk",
            auto_voucher=True,
            disk_id="lhdisk-xxxxxx",
            renew_disk_charge_prepaid=tencentcloud.lighthouse.RenewDiskRenewDiskChargePrepaidArgs(
                period=1,
                renew_flag="NOTIFY_AND_AUTO_RENEW",
                time_unit="m",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_voucher: Whether to automatically use the voucher. Not used by default.
        :param pulumi.Input[str] disk_id: List of disk ID.
        :param pulumi.Input[pulumi.InputType['RenewDiskRenewDiskChargePrepaidArgs']] renew_disk_charge_prepaid: Renew cloud hard disk subscription related parameter settings.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RenewDiskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a lighthouse renew_disk

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        renew_disk = tencentcloud.lighthouse.RenewDisk("renewDisk",
            auto_voucher=True,
            disk_id="lhdisk-xxxxxx",
            renew_disk_charge_prepaid=tencentcloud.lighthouse.RenewDiskRenewDiskChargePrepaidArgs(
                period=1,
                renew_flag="NOTIFY_AND_AUTO_RENEW",
                time_unit="m",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param RenewDiskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RenewDiskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_voucher: Optional[pulumi.Input[bool]] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 renew_disk_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewDiskRenewDiskChargePrepaidArgs']]] = None,
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
            __props__ = RenewDiskArgs.__new__(RenewDiskArgs)

            __props__.__dict__["auto_voucher"] = auto_voucher
            if disk_id is None and not opts.urn:
                raise TypeError("Missing required property 'disk_id'")
            __props__.__dict__["disk_id"] = disk_id
            if renew_disk_charge_prepaid is None and not opts.urn:
                raise TypeError("Missing required property 'renew_disk_charge_prepaid'")
            __props__.__dict__["renew_disk_charge_prepaid"] = renew_disk_charge_prepaid
        super(RenewDisk, __self__).__init__(
            'tencentcloud:Lighthouse/renewDisk:RenewDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_voucher: Optional[pulumi.Input[bool]] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            renew_disk_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewDiskRenewDiskChargePrepaidArgs']]] = None) -> 'RenewDisk':
        """
        Get an existing RenewDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_voucher: Whether to automatically use the voucher. Not used by default.
        :param pulumi.Input[str] disk_id: List of disk ID.
        :param pulumi.Input[pulumi.InputType['RenewDiskRenewDiskChargePrepaidArgs']] renew_disk_charge_prepaid: Renew cloud hard disk subscription related parameter settings.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RenewDiskState.__new__(_RenewDiskState)

        __props__.__dict__["auto_voucher"] = auto_voucher
        __props__.__dict__["disk_id"] = disk_id
        __props__.__dict__["renew_disk_charge_prepaid"] = renew_disk_charge_prepaid
        return RenewDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoVoucher")
    def auto_voucher(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to automatically use the voucher. Not used by default.
        """
        return pulumi.get(self, "auto_voucher")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        List of disk ID.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="renewDiskChargePrepaid")
    def renew_disk_charge_prepaid(self) -> pulumi.Output['outputs.RenewDiskRenewDiskChargePrepaid']:
        """
        Renew cloud hard disk subscription related parameter settings.
        """
        return pulumi.get(self, "renew_disk_charge_prepaid")

