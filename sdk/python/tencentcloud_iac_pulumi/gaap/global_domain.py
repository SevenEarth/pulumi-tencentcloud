# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GlobalDomainArgs', 'GlobalDomain']

@pulumi.input_type
class GlobalDomainArgs:
    def __init__(__self__, *,
                 default_value: pulumi.Input[str],
                 project_id: pulumi.Input[int],
                 alias: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a GlobalDomain resource.
        :param pulumi.Input[str] default_value: Domain name default entry.
        :param pulumi.Input[int] project_id: Domain Name Project ID.
        :param pulumi.Input[str] alias: alias.
        :param pulumi.Input[str] status: Global domain statue. Available values: open and close, default is open.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        """
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "project_id", project_id)
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> pulumi.Input[str]:
        """
        Domain name default entry.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        Domain Name Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        alias.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Global domain statue. Available values: open and close, default is open.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _GlobalDomainState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering GlobalDomain resources.
        :param pulumi.Input[str] alias: alias.
        :param pulumi.Input[str] default_value: Domain name default entry.
        :param pulumi.Input[int] project_id: Domain Name Project ID.
        :param pulumi.Input[str] status: Global domain statue. Available values: open and close, default is open.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if default_value is not None:
            pulumi.set(__self__, "default_value", default_value)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        alias.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name default entry.
        """
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Domain Name Project ID.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Global domain statue. Available values: open and close, default is open.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class GlobalDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provides a resource to create a gaap global domain

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        global_domain = tencentcloud.gaap.GlobalDomain("globalDomain",
            alias="demo",
            default_value="xxxxxx.com",
            project_id=0,
            tags={
                "key": "value",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        gaap global_domain can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Gaap/globalDomain:GlobalDomain global_domain ${projectId}#${domainId}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: alias.
        :param pulumi.Input[str] default_value: Domain name default entry.
        :param pulumi.Input[int] project_id: Domain Name Project ID.
        :param pulumi.Input[str] status: Global domain statue. Available values: open and close, default is open.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a gaap global domain

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        global_domain = tencentcloud.gaap.GlobalDomain("globalDomain",
            alias="demo",
            default_value="xxxxxx.com",
            project_id=0,
            tags={
                "key": "value",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        gaap global_domain can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Gaap/globalDomain:GlobalDomain global_domain ${projectId}#${domainId}
        ```

        :param str resource_name: The name of the resource.
        :param GlobalDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_value: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GlobalDomainArgs.__new__(GlobalDomainArgs)

            __props__.__dict__["alias"] = alias
            if default_value is None and not opts.urn:
                raise TypeError("Missing required property 'default_value'")
            __props__.__dict__["default_value"] = default_value
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
        super(GlobalDomain, __self__).__init__(
            'tencentcloud:Gaap/globalDomain:GlobalDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            default_value: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'GlobalDomain':
        """
        Get an existing GlobalDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: alias.
        :param pulumi.Input[str] default_value: Domain name default entry.
        :param pulumi.Input[int] project_id: Domain Name Project ID.
        :param pulumi.Input[str] status: Global domain statue. Available values: open and close, default is open.
        :param pulumi.Input[Mapping[str, Any]] tags: Instance tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalDomainState.__new__(_GlobalDomainState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["default_value"] = default_value
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return GlobalDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[Optional[str]]:
        """
        alias.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> pulumi.Output[str]:
        """
        Domain name default entry.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        Domain Name Project ID.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Global domain statue. Available values: open and close, default is open.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Instance tags.
        """
        return pulumi.get(self, "tags")

