# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BandwidthPackageArgs', 'BandwidthPackage']

@pulumi.input_type
class BandwidthPackageArgs:
    def __init__(__self__, *,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth: Optional[pulumi.Input[int]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a BandwidthPackage resource.
        :param pulumi.Input[str] bandwidth_package_name: Bandwidth package name.
        :param pulumi.Input[str] charge_type: Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        :param pulumi.Input[int] internet_max_bandwidth: Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        :param pulumi.Input[str] network_type: Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if bandwidth_package_name is not None:
            pulumi.set(__self__, "bandwidth_package_name", bandwidth_package_name)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if internet_max_bandwidth is not None:
            pulumi.set(__self__, "internet_max_bandwidth", internet_max_bandwidth)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth package name.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @bandwidth_package_name.setter
    def bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_name", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="internetMaxBandwidth")
    def internet_max_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        """
        return pulumi.get(self, "internet_max_bandwidth")

    @internet_max_bandwidth.setter
    def internet_max_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

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
class _BandwidthPackageState:
    def __init__(__self__, *,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth: Optional[pulumi.Input[int]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering BandwidthPackage resources.
        :param pulumi.Input[str] bandwidth_package_name: Bandwidth package name.
        :param pulumi.Input[str] charge_type: Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        :param pulumi.Input[int] internet_max_bandwidth: Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        :param pulumi.Input[str] network_type: Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if bandwidth_package_name is not None:
            pulumi.set(__self__, "bandwidth_package_name", bandwidth_package_name)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if internet_max_bandwidth is not None:
            pulumi.set(__self__, "internet_max_bandwidth", internet_max_bandwidth)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth package name.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @bandwidth_package_name.setter
    def bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_name", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="internetMaxBandwidth")
    def internet_max_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        """
        return pulumi.get(self, "internet_max_bandwidth")

    @internet_max_bandwidth.setter
    def internet_max_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internet_max_bandwidth", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

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


class BandwidthPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth: Optional[pulumi.Input[int]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc bandwidth_package

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        bandwidth_package = tencentcloud.vpc.BandwidthPackage("bandwidthPackage",
            bandwidth_package_name="test-001",
            charge_type="TOP5_POSTPAID_BY_MONTH",
            network_type="BGP",
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        vpc bandwidth_package can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/bandwidthPackage:BandwidthPackage bandwidth_package bandwidthPackage_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_name: Bandwidth package name.
        :param pulumi.Input[str] charge_type: Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        :param pulumi.Input[int] internet_max_bandwidth: Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        :param pulumi.Input[str] network_type: Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BandwidthPackageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc bandwidth_package

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        bandwidth_package = tencentcloud.vpc.BandwidthPackage("bandwidthPackage",
            bandwidth_package_name="test-001",
            charge_type="TOP5_POSTPAID_BY_MONTH",
            network_type="BGP",
            tags={
                "createdBy": "terraform",
            })
        ```

        ## Import

        vpc bandwidth_package can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/bandwidthPackage:BandwidthPackage bandwidth_package bandwidthPackage_id
        ```

        :param str resource_name: The name of the resource.
        :param BandwidthPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 internet_max_bandwidth: Optional[pulumi.Input[int]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = BandwidthPackageArgs.__new__(BandwidthPackageArgs)

            __props__.__dict__["bandwidth_package_name"] = bandwidth_package_name
            __props__.__dict__["charge_type"] = charge_type
            __props__.__dict__["internet_max_bandwidth"] = internet_max_bandwidth
            __props__.__dict__["network_type"] = network_type
            __props__.__dict__["tags"] = tags
        super(BandwidthPackage, __self__).__init__(
            'tencentcloud:Vpc/bandwidthPackage:BandwidthPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth_package_name: Optional[pulumi.Input[str]] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            internet_max_bandwidth: Optional[pulumi.Input[int]] = None,
            network_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'BandwidthPackage':
        """
        Get an existing BandwidthPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bandwidth_package_name: Bandwidth package name.
        :param pulumi.Input[str] charge_type: Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        :param pulumi.Input[int] internet_max_bandwidth: Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        :param pulumi.Input[str] network_type: Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPackageState.__new__(_BandwidthPackageState)

        __props__.__dict__["bandwidth_package_name"] = bandwidth_package_name
        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["internet_max_bandwidth"] = internet_max_bandwidth
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["tags"] = tags
        return BandwidthPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> pulumi.Output[Optional[str]]:
        """
        Bandwidth package name.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[Optional[str]]:
        """
        Bandwidth package billing type, default: TOP5_POSTPAID_BY_MONTH, optional value:- `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid- `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid- `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction).
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="internetMaxBandwidth")
    def internet_max_bandwidth(self) -> pulumi.Output[Optional[int]]:
        """
        Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
        """
        return pulumi.get(self, "internet_max_bandwidth")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[Optional[str]]:
        """
        Bandwidth packet type, default:BGP, optional:- `BGP`: common BGP shared bandwidth package- `HIGH_QUALITY_BGP`: Quality BGP Shared Bandwidth Package.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")
