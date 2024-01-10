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

__all__ = ['RunJobArgs', 'RunJob']

@pulumi.input_type
class RunJobArgs:
    def __init__(__self__, *,
                 run_job_descriptions: pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]],
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RunJob resource.
        :param pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]] run_job_descriptions: The description information for batch job startup.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        pulumi.set(__self__, "run_job_descriptions", run_job_descriptions)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="runJobDescriptions")
    def run_job_descriptions(self) -> pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]]:
        """
        The description information for batch job startup.
        """
        return pulumi.get(self, "run_job_descriptions")

    @run_job_descriptions.setter
    def run_job_descriptions(self, value: pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]]):
        pulumi.set(self, "run_job_descriptions", value)

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
class _RunJobState:
    def __init__(__self__, *,
                 run_job_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RunJob resources.
        :param pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]] run_job_descriptions: The description information for batch job startup.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        if run_job_descriptions is not None:
            pulumi.set(__self__, "run_job_descriptions", run_job_descriptions)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="runJobDescriptions")
    def run_job_descriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]]]:
        """
        The description information for batch job startup.
        """
        return pulumi.get(self, "run_job_descriptions")

    @run_job_descriptions.setter
    def run_job_descriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RunJobRunJobDescriptionArgs']]]]):
        pulumi.set(self, "run_job_descriptions", value)

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


class RunJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 run_job_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunJobRunJobDescriptionArgs']]]]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a oceanus run_job

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.RunJob("example",
            run_job_descriptions=[tencentcloud.oceanus.RunJobRunJobDescriptionArgs(
                job_config_version=10,
                job_id="cql-4xwincyn",
                run_type=1,
                start_mode="LATEST",
                use_old_system_connector=False,
            )],
            work_space_id="space-2idq8wbr")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunJobRunJobDescriptionArgs']]]] run_job_descriptions: The description information for batch job startup.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RunJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a oceanus run_job

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.RunJob("example",
            run_job_descriptions=[tencentcloud.oceanus.RunJobRunJobDescriptionArgs(
                job_config_version=10,
                job_id="cql-4xwincyn",
                run_type=1,
                start_mode="LATEST",
                use_old_system_connector=False,
            )],
            work_space_id="space-2idq8wbr")
        ```

        :param str resource_name: The name of the resource.
        :param RunJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RunJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 run_job_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunJobRunJobDescriptionArgs']]]]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RunJobArgs.__new__(RunJobArgs)

            if run_job_descriptions is None and not opts.urn:
                raise TypeError("Missing required property 'run_job_descriptions'")
            __props__.__dict__["run_job_descriptions"] = run_job_descriptions
            __props__.__dict__["work_space_id"] = work_space_id
        super(RunJob, __self__).__init__(
            'tencentcloud:Oceanus/runJob:RunJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            run_job_descriptions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunJobRunJobDescriptionArgs']]]]] = None,
            work_space_id: Optional[pulumi.Input[str]] = None) -> 'RunJob':
        """
        Get an existing RunJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunJobRunJobDescriptionArgs']]]] run_job_descriptions: The description information for batch job startup.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RunJobState.__new__(_RunJobState)

        __props__.__dict__["run_job_descriptions"] = run_job_descriptions
        __props__.__dict__["work_space_id"] = work_space_id
        return RunJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="runJobDescriptions")
    def run_job_descriptions(self) -> pulumi.Output[Sequence['outputs.RunJobRunJobDescription']]:
        """
        The description information for batch job startup.
        """
        return pulumi.get(self, "run_job_descriptions")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

