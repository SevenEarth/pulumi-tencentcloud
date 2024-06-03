# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncJobStartOperationArgs', 'SyncJobStartOperation']

@pulumi.input_type
class SyncJobStartOperationArgs:
    def __init__(__self__, *,
                 job_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a SyncJobStartOperation resource.
        :param pulumi.Input[str] job_id: Synchronization instance id (i.e. identifies a synchronization job).
        """
        pulumi.set(__self__, "job_id", job_id)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Input[str]:
        """
        Synchronization instance id (i.e. identifies a synchronization job).
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_id", value)


@pulumi.input_type
class _SyncJobStartOperationState:
    def __init__(__self__, *,
                 job_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyncJobStartOperation resources.
        :param pulumi.Input[str] job_id: Synchronization instance id (i.e. identifies a synchronization job).
        """
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        Synchronization instance id (i.e. identifies a synchronization job).
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)


class SyncJobStartOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dts sync_job_start_operation

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_job_start_operation = tencentcloud.dts.SyncJobStartOperation("syncJobStartOperation", job_id="sync-werwfs23")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_id: Synchronization instance id (i.e. identifies a synchronization job).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncJobStartOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dts sync_job_start_operation

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_job_start_operation = tencentcloud.dts.SyncJobStartOperation("syncJobStartOperation", job_id="sync-werwfs23")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param SyncJobStartOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncJobStartOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncJobStartOperationArgs.__new__(SyncJobStartOperationArgs)

            if job_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_id'")
            __props__.__dict__["job_id"] = job_id
        super(SyncJobStartOperation, __self__).__init__(
            'tencentcloud:Dts/syncJobStartOperation:SyncJobStartOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            job_id: Optional[pulumi.Input[str]] = None) -> 'SyncJobStartOperation':
        """
        Get an existing SyncJobStartOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_id: Synchronization instance id (i.e. identifies a synchronization job).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncJobStartOperationState.__new__(_SyncJobStartOperationState)

        __props__.__dict__["job_id"] = job_id
        return SyncJobStartOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        Synchronization instance id (i.e. identifies a synchronization job).
        """
        return pulumi.get(self, "job_id")

