# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PeerConnectAcceptOperationArgs', 'PeerConnectAcceptOperation']

@pulumi.input_type
class PeerConnectAcceptOperationArgs:
    def __init__(__self__, *,
                 peering_connection_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a PeerConnectAcceptOperation resource.
        :param pulumi.Input[str] peering_connection_id: Peer connection unique ID.
        """
        pulumi.set(__self__, "peering_connection_id", peering_connection_id)

    @property
    @pulumi.getter(name="peeringConnectionId")
    def peering_connection_id(self) -> pulumi.Input[str]:
        """
        Peer connection unique ID.
        """
        return pulumi.get(self, "peering_connection_id")

    @peering_connection_id.setter
    def peering_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peering_connection_id", value)


@pulumi.input_type
class _PeerConnectAcceptOperationState:
    def __init__(__self__, *,
                 peering_connection_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PeerConnectAcceptOperation resources.
        :param pulumi.Input[str] peering_connection_id: Peer connection unique ID.
        """
        if peering_connection_id is not None:
            pulumi.set(__self__, "peering_connection_id", peering_connection_id)

    @property
    @pulumi.getter(name="peeringConnectionId")
    def peering_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        Peer connection unique ID.
        """
        return pulumi.get(self, "peering_connection_id")

    @peering_connection_id.setter
    def peering_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_connection_id", value)


class PeerConnectAcceptOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc peer_connect_accept_operation

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        peer_connect_accept_operation = tencentcloud.vpc.PeerConnectAcceptOperation("peerConnectAcceptOperation", peering_connection_id="pcx-abced")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peering_connection_id: Peer connection unique ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PeerConnectAcceptOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc peer_connect_accept_operation

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        peer_connect_accept_operation = tencentcloud.vpc.PeerConnectAcceptOperation("peerConnectAcceptOperation", peering_connection_id="pcx-abced")
        ```

        :param str resource_name: The name of the resource.
        :param PeerConnectAcceptOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeerConnectAcceptOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peering_connection_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PeerConnectAcceptOperationArgs.__new__(PeerConnectAcceptOperationArgs)

            if peering_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'peering_connection_id'")
            __props__.__dict__["peering_connection_id"] = peering_connection_id
        super(PeerConnectAcceptOperation, __self__).__init__(
            'tencentcloud:Vpc/peerConnectAcceptOperation:PeerConnectAcceptOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            peering_connection_id: Optional[pulumi.Input[str]] = None) -> 'PeerConnectAcceptOperation':
        """
        Get an existing PeerConnectAcceptOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peering_connection_id: Peer connection unique ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PeerConnectAcceptOperationState.__new__(_PeerConnectAcceptOperationState)

        __props__.__dict__["peering_connection_id"] = peering_connection_id
        return PeerConnectAcceptOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="peeringConnectionId")
    def peering_connection_id(self) -> pulumi.Output[str]:
        """
        Peer connection unique ID.
        """
        return pulumi.get(self, "peering_connection_id")

