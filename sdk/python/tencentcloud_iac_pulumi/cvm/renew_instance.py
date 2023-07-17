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

__all__ = ['RenewInstanceArgs', 'RenewInstance']

@pulumi.input_type
class RenewInstanceArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_charge_prepaid: Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']] = None,
                 renew_portable_data_disk: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RenewInstance resource.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input['RenewInstanceInstanceChargePrepaidArgs'] instance_charge_prepaid: Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        :param pulumi.Input[bool] renew_portable_data_disk: Whether to renew the elastic data disk. Valid values:
               - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
               - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
               Default value: TRUE.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_charge_prepaid is not None:
            pulumi.set(__self__, "instance_charge_prepaid", instance_charge_prepaid)
        if renew_portable_data_disk is not None:
            pulumi.set(__self__, "renew_portable_data_disk", renew_portable_data_disk)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceChargePrepaid")
    def instance_charge_prepaid(self) -> Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']]:
        """
        Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        """
        return pulumi.get(self, "instance_charge_prepaid")

    @instance_charge_prepaid.setter
    def instance_charge_prepaid(self, value: Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']]):
        pulumi.set(self, "instance_charge_prepaid", value)

    @property
    @pulumi.getter(name="renewPortableDataDisk")
    def renew_portable_data_disk(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to renew the elastic data disk. Valid values:
        - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
        - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
        Default value: TRUE.
        """
        return pulumi.get(self, "renew_portable_data_disk")

    @renew_portable_data_disk.setter
    def renew_portable_data_disk(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "renew_portable_data_disk", value)


@pulumi.input_type
class _RenewInstanceState:
    def __init__(__self__, *,
                 instance_charge_prepaid: Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 renew_portable_data_disk: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering RenewInstance resources.
        :param pulumi.Input['RenewInstanceInstanceChargePrepaidArgs'] instance_charge_prepaid: Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[bool] renew_portable_data_disk: Whether to renew the elastic data disk. Valid values:
               - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
               - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
               Default value: TRUE.
        """
        if instance_charge_prepaid is not None:
            pulumi.set(__self__, "instance_charge_prepaid", instance_charge_prepaid)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if renew_portable_data_disk is not None:
            pulumi.set(__self__, "renew_portable_data_disk", renew_portable_data_disk)

    @property
    @pulumi.getter(name="instanceChargePrepaid")
    def instance_charge_prepaid(self) -> Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']]:
        """
        Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        """
        return pulumi.get(self, "instance_charge_prepaid")

    @instance_charge_prepaid.setter
    def instance_charge_prepaid(self, value: Optional[pulumi.Input['RenewInstanceInstanceChargePrepaidArgs']]):
        pulumi.set(self, "instance_charge_prepaid", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="renewPortableDataDisk")
    def renew_portable_data_disk(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to renew the elastic data disk. Valid values:
        - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
        - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
        Default value: TRUE.
        """
        return pulumi.get(self, "renew_portable_data_disk")

    @renew_portable_data_disk.setter
    def renew_portable_data_disk(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "renew_portable_data_disk", value)


class RenewInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewInstanceInstanceChargePrepaidArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 renew_portable_data_disk: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Provides a resource to create a cvm renew_instance

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RenewInstanceInstanceChargePrepaidArgs']] instance_charge_prepaid: Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[bool] renew_portable_data_disk: Whether to renew the elastic data disk. Valid values:
               - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
               - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
               Default value: TRUE.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RenewInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cvm renew_instance

        :param str resource_name: The name of the resource.
        :param RenewInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RenewInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewInstanceInstanceChargePrepaidArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 renew_portable_data_disk: Optional[pulumi.Input[bool]] = None,
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
            __props__ = RenewInstanceArgs.__new__(RenewInstanceArgs)

            __props__.__dict__["instance_charge_prepaid"] = instance_charge_prepaid
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["renew_portable_data_disk"] = renew_portable_data_disk
        super(RenewInstance, __self__).__init__(
            'tencentcloud:Cvm/renewInstance:RenewInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_charge_prepaid: Optional[pulumi.Input[pulumi.InputType['RenewInstanceInstanceChargePrepaidArgs']]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            renew_portable_data_disk: Optional[pulumi.Input[bool]] = None) -> 'RenewInstance':
        """
        Get an existing RenewInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RenewInstanceInstanceChargePrepaidArgs']] instance_charge_prepaid: Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[bool] renew_portable_data_disk: Whether to renew the elastic data disk. Valid values:
               - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
               - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
               Default value: TRUE.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RenewInstanceState.__new__(_RenewInstanceState)

        __props__.__dict__["instance_charge_prepaid"] = instance_charge_prepaid
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["renew_portable_data_disk"] = renew_portable_data_disk
        return RenewInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceChargePrepaid")
    def instance_charge_prepaid(self) -> pulumi.Output[Optional['outputs.RenewInstanceInstanceChargePrepaid']]:
        """
        Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.
        """
        return pulumi.get(self, "instance_charge_prepaid")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="renewPortableDataDisk")
    def renew_portable_data_disk(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to renew the elastic data disk. Valid values:
        - `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time
        - `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed
        Default value: TRUE.
        """
        return pulumi.get(self, "renew_portable_data_disk")

