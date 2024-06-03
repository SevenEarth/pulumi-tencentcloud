# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TriggerJobSavepointArgs', 'TriggerJobSavepoint']

@pulumi.input_type
class TriggerJobSavepointArgs:
    def __init__(__self__, *,
                 job_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerJobSavepoint resource.
        :param pulumi.Input[str] job_id: Job SerialId.
        :param pulumi.Input[str] description: Savepoint description.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        pulumi.set(__self__, "job_id", job_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Input[str]:
        """
        Job SerialId.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Savepoint description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


@pulumi.input_type
class _TriggerJobSavepointState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TriggerJobSavepoint resources.
        :param pulumi.Input[str] description: Savepoint description.
        :param pulumi.Input[str] job_id: Job SerialId.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Savepoint description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        Job SerialId.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


class TriggerJobSavepoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a oceanus trigger_job_savepoint

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.TriggerJobSavepoint("example",
            description="description.",
            job_id="cql-4xwincyn",
            work_space_id="space-2idq8wbr")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Savepoint description.
        :param pulumi.Input[str] job_id: Job SerialId.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerJobSavepointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a oceanus trigger_job_savepoint

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.TriggerJobSavepoint("example",
            description="description.",
            job_id="cql-4xwincyn",
            work_space_id="space-2idq8wbr")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param TriggerJobSavepointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerJobSavepointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TriggerJobSavepointArgs.__new__(TriggerJobSavepointArgs)

            __props__.__dict__["description"] = description
            if job_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_id'")
            __props__.__dict__["job_id"] = job_id
            __props__.__dict__["work_space_id"] = work_space_id
        super(TriggerJobSavepoint, __self__).__init__(
            'tencentcloud:Oceanus/triggerJobSavepoint:TriggerJobSavepoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            job_id: Optional[pulumi.Input[str]] = None,
            work_space_id: Optional[pulumi.Input[str]] = None) -> 'TriggerJobSavepoint':
        """
        Get an existing TriggerJobSavepoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Savepoint description.
        :param pulumi.Input[str] job_id: Job SerialId.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerJobSavepointState.__new__(_TriggerJobSavepointState)

        __props__.__dict__["description"] = description
        __props__.__dict__["job_id"] = job_id
        __props__.__dict__["work_space_id"] = work_space_id
        return TriggerJobSavepoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Savepoint description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        Job SerialId.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

