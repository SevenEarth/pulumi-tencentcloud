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

__all__ = ['MeshArgs', 'Mesh']

@pulumi.input_type
class MeshArgs:
    def __init__(__self__, *,
                 config: pulumi.Input['MeshConfigArgs'],
                 display_name: pulumi.Input[str],
                 mesh_version: pulumi.Input[str],
                 type: pulumi.Input[str],
                 mesh_id: Optional[pulumi.Input[str]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]] = None):
        """
        The set of arguments for constructing a Mesh resource.
        :param pulumi.Input['MeshConfigArgs'] config: Mesh configuration.
        :param pulumi.Input[str] display_name: Mesh name.
        :param pulumi.Input[str] mesh_version: Mesh version.
        :param pulumi.Input[str] type: Mesh type.
        :param pulumi.Input[str] mesh_id: Mesh ID.
        :param pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]] tag_lists: A list of associated tags.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "mesh_version", mesh_version)
        pulumi.set(__self__, "type", type)
        if mesh_id is not None:
            pulumi.set(__self__, "mesh_id", mesh_id)
        if tag_lists is not None:
            pulumi.set(__self__, "tag_lists", tag_lists)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['MeshConfigArgs']:
        """
        Mesh configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['MeshConfigArgs']):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Mesh name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="meshVersion")
    def mesh_version(self) -> pulumi.Input[str]:
        """
        Mesh version.
        """
        return pulumi.get(self, "mesh_version")

    @mesh_version.setter
    def mesh_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "mesh_version", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Mesh type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="meshId")
    def mesh_id(self) -> Optional[pulumi.Input[str]]:
        """
        Mesh ID.
        """
        return pulumi.get(self, "mesh_id")

    @mesh_id.setter
    def mesh_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mesh_id", value)

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]]:
        """
        A list of associated tags.
        """
        return pulumi.get(self, "tag_lists")

    @tag_lists.setter
    def tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]]):
        pulumi.set(self, "tag_lists", value)


@pulumi.input_type
class _MeshState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input['MeshConfigArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mesh_id: Optional[pulumi.Input[str]] = None,
                 mesh_version: Optional[pulumi.Input[str]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Mesh resources.
        :param pulumi.Input['MeshConfigArgs'] config: Mesh configuration.
        :param pulumi.Input[str] display_name: Mesh name.
        :param pulumi.Input[str] mesh_id: Mesh ID.
        :param pulumi.Input[str] mesh_version: Mesh version.
        :param pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]] tag_lists: A list of associated tags.
        :param pulumi.Input[str] type: Mesh type.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if mesh_id is not None:
            pulumi.set(__self__, "mesh_id", mesh_id)
        if mesh_version is not None:
            pulumi.set(__self__, "mesh_version", mesh_version)
        if tag_lists is not None:
            pulumi.set(__self__, "tag_lists", tag_lists)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['MeshConfigArgs']]:
        """
        Mesh configuration.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['MeshConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Mesh name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="meshId")
    def mesh_id(self) -> Optional[pulumi.Input[str]]:
        """
        Mesh ID.
        """
        return pulumi.get(self, "mesh_id")

    @mesh_id.setter
    def mesh_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mesh_id", value)

    @property
    @pulumi.getter(name="meshVersion")
    def mesh_version(self) -> Optional[pulumi.Input[str]]:
        """
        Mesh version.
        """
        return pulumi.get(self, "mesh_version")

    @mesh_version.setter
    def mesh_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mesh_version", value)

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]]:
        """
        A list of associated tags.
        """
        return pulumi.get(self, "tag_lists")

    @tag_lists.setter
    def tag_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MeshTagListArgs']]]]):
        pulumi.set(self, "tag_lists", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Mesh type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Mesh(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['MeshConfigArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mesh_id: Optional[pulumi.Input[str]] = None,
                 mesh_version: Optional[pulumi.Input[str]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MeshTagListArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tcm mesh

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        basic = tencentcloud.tcm.Mesh("basic",
            config=tencentcloud.tcm.MeshConfigArgs(
                istio=tencentcloud.tcm.MeshConfigIstioArgs(
                    disable_http_retry=True,
                    disable_policy_checks=True,
                    enable_pilot_http=True,
                    outbound_traffic_policy="ALLOW_ANY",
                    smart_dns=tencentcloud.tcm.MeshConfigIstioSmartDnsArgs(
                        istio_meta_dns_auto_allocate=True,
                        istio_meta_dns_capture=True,
                    ),
                ),
            ),
            display_name="test_mesh",
            mesh_version="1.12.5",
            tag_lists=[tencentcloud.tcm.MeshTagListArgs(
                key="key",
                passthrough=True,
                value="value",
            )],
            type="HOSTED")
        ```

        ## Import

        tcm mesh can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcm/mesh:Mesh mesh mesh_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MeshConfigArgs']] config: Mesh configuration.
        :param pulumi.Input[str] display_name: Mesh name.
        :param pulumi.Input[str] mesh_id: Mesh ID.
        :param pulumi.Input[str] mesh_version: Mesh version.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MeshTagListArgs']]]] tag_lists: A list of associated tags.
        :param pulumi.Input[str] type: Mesh type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MeshArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tcm mesh

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        basic = tencentcloud.tcm.Mesh("basic",
            config=tencentcloud.tcm.MeshConfigArgs(
                istio=tencentcloud.tcm.MeshConfigIstioArgs(
                    disable_http_retry=True,
                    disable_policy_checks=True,
                    enable_pilot_http=True,
                    outbound_traffic_policy="ALLOW_ANY",
                    smart_dns=tencentcloud.tcm.MeshConfigIstioSmartDnsArgs(
                        istio_meta_dns_auto_allocate=True,
                        istio_meta_dns_capture=True,
                    ),
                ),
            ),
            display_name="test_mesh",
            mesh_version="1.12.5",
            tag_lists=[tencentcloud.tcm.MeshTagListArgs(
                key="key",
                passthrough=True,
                value="value",
            )],
            type="HOSTED")
        ```

        ## Import

        tcm mesh can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcm/mesh:Mesh mesh mesh_id
        ```

        :param str resource_name: The name of the resource.
        :param MeshArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MeshArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['MeshConfigArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mesh_id: Optional[pulumi.Input[str]] = None,
                 mesh_version: Optional[pulumi.Input[str]] = None,
                 tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MeshTagListArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = MeshArgs.__new__(MeshArgs)

            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["mesh_id"] = mesh_id
            if mesh_version is None and not opts.urn:
                raise TypeError("Missing required property 'mesh_version'")
            __props__.__dict__["mesh_version"] = mesh_version
            __props__.__dict__["tag_lists"] = tag_lists
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Mesh, __self__).__init__(
            'tencentcloud:Tcm/mesh:Mesh',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['MeshConfigArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            mesh_id: Optional[pulumi.Input[str]] = None,
            mesh_version: Optional[pulumi.Input[str]] = None,
            tag_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MeshTagListArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Mesh':
        """
        Get an existing Mesh resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MeshConfigArgs']] config: Mesh configuration.
        :param pulumi.Input[str] display_name: Mesh name.
        :param pulumi.Input[str] mesh_id: Mesh ID.
        :param pulumi.Input[str] mesh_version: Mesh version.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MeshTagListArgs']]]] tag_lists: A list of associated tags.
        :param pulumi.Input[str] type: Mesh type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MeshState.__new__(_MeshState)

        __props__.__dict__["config"] = config
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["mesh_id"] = mesh_id
        __props__.__dict__["mesh_version"] = mesh_version
        __props__.__dict__["tag_lists"] = tag_lists
        __props__.__dict__["type"] = type
        return Mesh(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.MeshConfig']:
        """
        Mesh configuration.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Mesh name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="meshId")
    def mesh_id(self) -> pulumi.Output[str]:
        """
        Mesh ID.
        """
        return pulumi.get(self, "mesh_id")

    @property
    @pulumi.getter(name="meshVersion")
    def mesh_version(self) -> pulumi.Output[str]:
        """
        Mesh version.
        """
        return pulumi.get(self, "mesh_version")

    @property
    @pulumi.getter(name="tagLists")
    def tag_lists(self) -> pulumi.Output[Optional[Sequence['outputs.MeshTagList']]]:
        """
        A list of associated tags.
        """
        return pulumi.get(self, "tag_lists")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Mesh type.
        """
        return pulumi.get(self, "type")
