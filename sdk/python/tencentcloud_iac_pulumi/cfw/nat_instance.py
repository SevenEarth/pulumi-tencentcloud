# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NatInstanceArgs', 'NatInstance']

@pulumi.input_type
class NatInstanceArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[int],
                 width: pulumi.Input[int],
                 zone_sets: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cross_a_zone: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_gw_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 new_mode_items: Optional[pulumi.Input['NatInstanceNewModeItemsArgs']] = None):
        """
        The set of arguments for constructing a NatInstance resource.
        :param pulumi.Input[int] mode: Mode 1: access mode; 0: new mode.
        :param pulumi.Input[int] width: Bandwidth.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_sets: Zone list.
        :param pulumi.Input[int] cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        :param pulumi.Input[str] name: Firewall instance name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_gw_lists: A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        :param pulumi.Input['NatInstanceNewModeItemsArgs'] new_mode_items: New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        """
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "width", width)
        pulumi.set(__self__, "zone_sets", zone_sets)
        if cross_a_zone is not None:
            pulumi.set(__self__, "cross_a_zone", cross_a_zone)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat_gw_lists is not None:
            pulumi.set(__self__, "nat_gw_lists", nat_gw_lists)
        if new_mode_items is not None:
            pulumi.set(__self__, "new_mode_items", new_mode_items)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[int]:
        """
        Mode 1: access mode; 0: new mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: pulumi.Input[int]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def width(self) -> pulumi.Input[int]:
        """
        Bandwidth.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: pulumi.Input[int]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="zoneSets")
    def zone_sets(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Zone list.
        """
        return pulumi.get(self, "zone_sets")

    @zone_sets.setter
    def zone_sets(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "zone_sets", value)

    @property
    @pulumi.getter(name="crossAZone")
    def cross_a_zone(self) -> Optional[pulumi.Input[int]]:
        """
        Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        """
        return pulumi.get(self, "cross_a_zone")

    @cross_a_zone.setter
    def cross_a_zone(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cross_a_zone", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="natGwLists")
    def nat_gw_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        """
        return pulumi.get(self, "nat_gw_lists")

    @nat_gw_lists.setter
    def nat_gw_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nat_gw_lists", value)

    @property
    @pulumi.getter(name="newModeItems")
    def new_mode_items(self) -> Optional[pulumi.Input['NatInstanceNewModeItemsArgs']]:
        """
        New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        """
        return pulumi.get(self, "new_mode_items")

    @new_mode_items.setter
    def new_mode_items(self, value: Optional[pulumi.Input['NatInstanceNewModeItemsArgs']]):
        pulumi.set(self, "new_mode_items", value)


@pulumi.input_type
class _NatInstanceState:
    def __init__(__self__, *,
                 cross_a_zone: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_gw_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 new_mode_items: Optional[pulumi.Input['NatInstanceNewModeItemsArgs']] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 zone_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering NatInstance resources.
        :param pulumi.Input[int] cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        :param pulumi.Input[int] mode: Mode 1: access mode; 0: new mode.
        :param pulumi.Input[str] name: Firewall instance name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_gw_lists: A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        :param pulumi.Input['NatInstanceNewModeItemsArgs'] new_mode_items: New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        :param pulumi.Input[int] width: Bandwidth.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_sets: Zone list.
        """
        if cross_a_zone is not None:
            pulumi.set(__self__, "cross_a_zone", cross_a_zone)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nat_gw_lists is not None:
            pulumi.set(__self__, "nat_gw_lists", nat_gw_lists)
        if new_mode_items is not None:
            pulumi.set(__self__, "new_mode_items", new_mode_items)
        if width is not None:
            pulumi.set(__self__, "width", width)
        if zone_sets is not None:
            pulumi.set(__self__, "zone_sets", zone_sets)

    @property
    @pulumi.getter(name="crossAZone")
    def cross_a_zone(self) -> Optional[pulumi.Input[int]]:
        """
        Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        """
        return pulumi.get(self, "cross_a_zone")

    @cross_a_zone.setter
    def cross_a_zone(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cross_a_zone", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[int]]:
        """
        Mode 1: access mode; 0: new mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Firewall instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="natGwLists")
    def nat_gw_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        """
        return pulumi.get(self, "nat_gw_lists")

    @nat_gw_lists.setter
    def nat_gw_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nat_gw_lists", value)

    @property
    @pulumi.getter(name="newModeItems")
    def new_mode_items(self) -> Optional[pulumi.Input['NatInstanceNewModeItemsArgs']]:
        """
        New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        """
        return pulumi.get(self, "new_mode_items")

    @new_mode_items.setter
    def new_mode_items(self, value: Optional[pulumi.Input['NatInstanceNewModeItemsArgs']]):
        pulumi.set(self, "new_mode_items", value)

    @property
    @pulumi.getter
    def width(self) -> Optional[pulumi.Input[int]]:
        """
        Bandwidth.
        """
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "width", value)

    @property
    @pulumi.getter(name="zoneSets")
    def zone_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Zone list.
        """
        return pulumi.get(self, "zone_sets")

    @zone_sets.setter
    def zone_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "zone_sets", value)


class NatInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cross_a_zone: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_gw_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 new_mode_items: Optional[pulumi.Input[pulumi.InputType['NatInstanceNewModeItemsArgs']]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 zone_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a cfw nat_instance

        ## Example Usage

        ### If mode is 0

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.cfw.NatInstance("example",
            cross_a_zone=0,
            mode=0,
            new_mode_items=tencentcloud.cfw.NatInstanceNewModeItemsArgs(
                eips=["152.136.168.192"],
                vpc_lists=["vpc-5063ta4i"],
            ),
            width=20,
            zone_sets=["ap-guangzhou-7"])
        ```
        <!--End PulumiCodeChooser -->

        ### If mode is 1

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.cfw.NatInstance("example",
            cross_a_zone=0,
            mode=1,
            nat_gw_lists=["nat-9wwkz1kr"],
            width=20,
            zone_sets=[
                "ap-guangzhou-6",
                "ap-guangzhou-7",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cfw nat_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cfw/natInstance:NatInstance example cfwnat-54a21421
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        :param pulumi.Input[int] mode: Mode 1: access mode; 0: new mode.
        :param pulumi.Input[str] name: Firewall instance name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_gw_lists: A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        :param pulumi.Input[pulumi.InputType['NatInstanceNewModeItemsArgs']] new_mode_items: New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        :param pulumi.Input[int] width: Bandwidth.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_sets: Zone list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cfw nat_instance

        ## Example Usage

        ### If mode is 0

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.cfw.NatInstance("example",
            cross_a_zone=0,
            mode=0,
            new_mode_items=tencentcloud.cfw.NatInstanceNewModeItemsArgs(
                eips=["152.136.168.192"],
                vpc_lists=["vpc-5063ta4i"],
            ),
            width=20,
            zone_sets=["ap-guangzhou-7"])
        ```
        <!--End PulumiCodeChooser -->

        ### If mode is 1

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.cfw.NatInstance("example",
            cross_a_zone=0,
            mode=1,
            nat_gw_lists=["nat-9wwkz1kr"],
            width=20,
            zone_sets=[
                "ap-guangzhou-6",
                "ap-guangzhou-7",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cfw nat_instance can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cfw/natInstance:NatInstance example cfwnat-54a21421
        ```

        :param str resource_name: The name of the resource.
        :param NatInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cross_a_zone: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nat_gw_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 new_mode_items: Optional[pulumi.Input[pulumi.InputType['NatInstanceNewModeItemsArgs']]] = None,
                 width: Optional[pulumi.Input[int]] = None,
                 zone_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NatInstanceArgs.__new__(NatInstanceArgs)

            __props__.__dict__["cross_a_zone"] = cross_a_zone
            if mode is None and not opts.urn:
                raise TypeError("Missing required property 'mode'")
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            __props__.__dict__["nat_gw_lists"] = nat_gw_lists
            __props__.__dict__["new_mode_items"] = new_mode_items
            if width is None and not opts.urn:
                raise TypeError("Missing required property 'width'")
            __props__.__dict__["width"] = width
            if zone_sets is None and not opts.urn:
                raise TypeError("Missing required property 'zone_sets'")
            __props__.__dict__["zone_sets"] = zone_sets
        super(NatInstance, __self__).__init__(
            'tencentcloud:Cfw/natInstance:NatInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cross_a_zone: Optional[pulumi.Input[int]] = None,
            mode: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nat_gw_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            new_mode_items: Optional[pulumi.Input[pulumi.InputType['NatInstanceNewModeItemsArgs']]] = None,
            width: Optional[pulumi.Input[int]] = None,
            zone_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'NatInstance':
        """
        Get an existing NatInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cross_a_zone: Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        :param pulumi.Input[int] mode: Mode 1: access mode; 0: new mode.
        :param pulumi.Input[str] name: Firewall instance name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nat_gw_lists: A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        :param pulumi.Input[pulumi.InputType['NatInstanceNewModeItemsArgs']] new_mode_items: New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        :param pulumi.Input[int] width: Bandwidth.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zone_sets: Zone list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NatInstanceState.__new__(_NatInstanceState)

        __props__.__dict__["cross_a_zone"] = cross_a_zone
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["nat_gw_lists"] = nat_gw_lists
        __props__.__dict__["new_mode_items"] = new_mode_items
        __props__.__dict__["width"] = width
        __props__.__dict__["zone_sets"] = zone_sets
        return NatInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="crossAZone")
    def cross_a_zone(self) -> pulumi.Output[Optional[int]]:
        """
        Off-site disaster recovery 1: use off-site disaster recovery; 0: do not use off-site disaster recovery; if empty, the default is not to use off-site disaster recovery.
        """
        return pulumi.get(self, "cross_a_zone")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[int]:
        """
        Mode 1: access mode; 0: new mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Firewall instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natGwLists")
    def nat_gw_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of nat gateways connected to the access mode, at least one of NewModeItems and NatgwList is passed.
        """
        return pulumi.get(self, "nat_gw_lists")

    @property
    @pulumi.getter(name="newModeItems")
    def new_mode_items(self) -> pulumi.Output[Optional['outputs.NatInstanceNewModeItems']]:
        """
        New mode passing parameters are added, at least one of new_mode_items and nat_gw_list is passed.
        """
        return pulumi.get(self, "new_mode_items")

    @property
    @pulumi.getter
    def width(self) -> pulumi.Output[int]:
        """
        Bandwidth.
        """
        return pulumi.get(self, "width")

    @property
    @pulumi.getter(name="zoneSets")
    def zone_sets(self) -> pulumi.Output[Sequence[str]]:
        """
        Zone list.
        """
        return pulumi.get(self, "zone_sets")

