# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EventBusArgs', 'EventBus']

@pulumi.input_type
class EventBusArgs:
    def __init__(__self__, *,
                 event_bus_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enable_store: Optional[pulumi.Input[bool]] = None,
                 save_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a EventBus resource.
        :param pulumi.Input[str] event_bus_name: Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        :param pulumi.Input[str] description: Event set description, unlimited character type, description within 200 characters.
        :param pulumi.Input[bool] enable_store: Whether the EB storage is enabled.
        :param pulumi.Input[int] save_days: EB storage duration.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "event_bus_name", event_bus_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_store is not None:
            pulumi.set(__self__, "enable_store", enable_store)
        if save_days is not None:
            pulumi.set(__self__, "save_days", save_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> pulumi.Input[str]:
        """
        Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Event set description, unlimited character type, description within 200 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableStore")
    def enable_store(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the EB storage is enabled.
        """
        return pulumi.get(self, "enable_store")

    @enable_store.setter
    def enable_store(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_store", value)

    @property
    @pulumi.getter(name="saveDays")
    def save_days(self) -> Optional[pulumi.Input[int]]:
        """
        EB storage duration.
        """
        return pulumi.get(self, "save_days")

    @save_days.setter
    def save_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "save_days", value)

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
class _EventBusState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_store: Optional[pulumi.Input[bool]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 save_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering EventBus resources.
        :param pulumi.Input[str] description: Event set description, unlimited character type, description within 200 characters.
        :param pulumi.Input[bool] enable_store: Whether the EB storage is enabled.
        :param pulumi.Input[str] event_bus_name: Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        :param pulumi.Input[int] save_days: EB storage duration.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_store is not None:
            pulumi.set(__self__, "enable_store", enable_store)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)
        if save_days is not None:
            pulumi.set(__self__, "save_days", save_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Event set description, unlimited character type, description within 200 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableStore")
    def enable_store(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the EB storage is enabled.
        """
        return pulumi.get(self, "enable_store")

    @enable_store.setter
    def enable_store(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_store", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        """
        Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        """
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter(name="saveDays")
    def save_days(self) -> Optional[pulumi.Input[int]]:
        """
        EB storage duration.
        """
        return pulumi.get(self, "save_days")

    @save_days.setter
    def save_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "save_days", value)

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


class EventBus(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_store: Optional[pulumi.Input[bool]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 save_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a eb event_bus

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.eb.EventBus("foo",
            description="event bus desc",
            enable_store=False,
            event_bus_name="tf-event_bus",
            save_days=1,
            tags={
                "createdBy": "terraform",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        eb event_bus can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Eb/eventBus:EventBus event_bus event_bus_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Event set description, unlimited character type, description within 200 characters.
        :param pulumi.Input[bool] enable_store: Whether the EB storage is enabled.
        :param pulumi.Input[str] event_bus_name: Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        :param pulumi.Input[int] save_days: EB storage duration.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventBusArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a eb event_bus

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.eb.EventBus("foo",
            description="event bus desc",
            enable_store=False,
            event_bus_name="tf-event_bus",
            save_days=1,
            tags={
                "createdBy": "terraform",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        eb event_bus can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Eb/eventBus:EventBus event_bus event_bus_id
        ```

        :param str resource_name: The name of the resource.
        :param EventBusArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventBusArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_store: Optional[pulumi.Input[bool]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 save_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventBusArgs.__new__(EventBusArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["enable_store"] = enable_store
            if event_bus_name is None and not opts.urn:
                raise TypeError("Missing required property 'event_bus_name'")
            __props__.__dict__["event_bus_name"] = event_bus_name
            __props__.__dict__["save_days"] = save_days
            __props__.__dict__["tags"] = tags
        super(EventBus, __self__).__init__(
            'tencentcloud:Eb/eventBus:EventBus',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_store: Optional[pulumi.Input[bool]] = None,
            event_bus_name: Optional[pulumi.Input[str]] = None,
            save_days: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'EventBus':
        """
        Get an existing EventBus resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Event set description, unlimited character type, description within 200 characters.
        :param pulumi.Input[bool] enable_store: Whether the EB storage is enabled.
        :param pulumi.Input[str] event_bus_name: Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        :param pulumi.Input[int] save_days: EB storage duration.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventBusState.__new__(_EventBusState)

        __props__.__dict__["description"] = description
        __props__.__dict__["enable_store"] = enable_store
        __props__.__dict__["event_bus_name"] = event_bus_name
        __props__.__dict__["save_days"] = save_days
        __props__.__dict__["tags"] = tags
        return EventBus(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Event set description, unlimited character type, description within 200 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableStore")
    def enable_store(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the EB storage is enabled.
        """
        return pulumi.get(self, "enable_store")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> pulumi.Output[str]:
        """
        Event set name, which can only contain letters, numbers, underscores, hyphens, starts with a letter and ends with a number or letter, 2~60 characters.
        """
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter(name="saveDays")
    def save_days(self) -> pulumi.Output[Optional[int]]:
        """
        EB storage duration.
        """
        return pulumi.get(self, "save_days")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

