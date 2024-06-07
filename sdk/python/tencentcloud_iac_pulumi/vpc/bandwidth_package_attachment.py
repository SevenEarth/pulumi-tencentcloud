# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BandwidthPackageAttachmentArgs', 'BandwidthPackageAttachment']

@pulumi.input_type
class BandwidthPackageAttachmentArgs:
    def __init__(__self__, *,
                 bandwidth_package_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str],
                 network_type: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BandwidthPackageAttachment resource.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth package unique ID, in the form of `bwp-xxxx`.
        :param pulumi.Input[str] resource_id: The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        :param pulumi.Input[str] network_type: Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        :param pulumi.Input[str] protocol: Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        :param pulumi.Input[str] resource_type: Resource types, including `Address`, `LoadBalance`.
        """
        pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        pulumi.set(__self__, "resource_id", resource_id)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Input[str]:
        """
        Bandwidth package unique ID, in the form of `bwp-xxxx`.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource types, including `Address`, `LoadBalance`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


@pulumi.input_type
class _BandwidthPackageAttachmentState:
    def __init__(__self__, *,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BandwidthPackageAttachment resources.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth package unique ID, in the form of `bwp-xxxx`.
        :param pulumi.Input[str] network_type: Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        :param pulumi.Input[str] protocol: Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        :param pulumi.Input[str] resource_id: The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        :param pulumi.Input[str] resource_type: Resource types, including `Address`, `LoadBalance`.
        """
        if bandwidth_package_id is not None:
            pulumi.set(__self__, "bandwidth_package_id", bandwidth_package_id)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth package unique ID, in the form of `bwp-xxxx`.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @bandwidth_package_id.setter
    def bandwidth_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_id", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Resource types, including `Address`, `LoadBalance`.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


class BandwidthPackageAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc bandwidth_package_attachment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth package unique ID, in the form of `bwp-xxxx`.
        :param pulumi.Input[str] network_type: Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        :param pulumi.Input[str] protocol: Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        :param pulumi.Input[str] resource_id: The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        :param pulumi.Input[str] resource_type: Resource types, including `Address`, `LoadBalance`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BandwidthPackageAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc bandwidth_package_attachment

        :param str resource_name: The name of the resource.
        :param BandwidthPackageAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPackageAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BandwidthPackageAttachmentArgs.__new__(BandwidthPackageAttachmentArgs)

            if bandwidth_package_id is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth_package_id'")
            __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
            __props__.__dict__["network_type"] = network_type
            __props__.__dict__["protocol"] = protocol
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            __props__.__dict__["resource_type"] = resource_type
        super(BandwidthPackageAttachment, __self__).__init__(
            'tencentcloud:Vpc/bandwidthPackageAttachment:BandwidthPackageAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None) -> 'BandwidthPackageAttachment':
        """
        Get an existing BandwidthPackageAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_id: Bandwidth package unique ID, in the form of `bwp-xxxx`.
        :param pulumi.Input[str] network_type: Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        :param pulumi.Input[str] protocol: Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        :param pulumi.Input[str] resource_id: The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        :param pulumi.Input[str] resource_type: Resource types, including `Address`, `LoadBalance`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPackageAttachmentState.__new__(_BandwidthPackageAttachmentState)

        __props__.__dict__["bandwidth_package_id"] = bandwidth_package_id
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["resource_type"] = resource_type
        return BandwidthPackageAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandwidthPackageId")
    def bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        Bandwidth package unique ID, in the form of `bwp-xxxx`.
        """
        return pulumi.get(self, "bandwidth_package_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[Optional[str]]:
        """
        Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        """
        Resource types, including `Address`, `LoadBalance`.
        """
        return pulumi.get(self, "resource_type")

