# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobCopyArgs', 'JobCopy']

@pulumi.input_type
class JobCopyArgs:
    def __init__(__self__, *,
                 source_id: pulumi.Input[str],
                 target_cluster_id: pulumi.Input[str],
                 job_type: Optional[pulumi.Input[int]] = None,
                 source_name: Optional[pulumi.Input[str]] = None,
                 target_folder_id: Optional[pulumi.Input[str]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobCopy resource.
        :param pulumi.Input[str] source_id: The serial ID of the job to be copied.
        :param pulumi.Input[str] target_cluster_id: The cluster serial ID of the target cluster.
        :param pulumi.Input[int] job_type: The type of the source job.
        :param pulumi.Input[str] source_name: The name of the job to be copied.
        :param pulumi.Input[str] target_folder_id: The directory ID of the new job.
        :param pulumi.Input[str] target_name: The name of the new job.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        pulumi.set(__self__, "source_id", source_id)
        pulumi.set(__self__, "target_cluster_id", target_cluster_id)
        if job_type is not None:
            pulumi.set(__self__, "job_type", job_type)
        if source_name is not None:
            pulumi.set(__self__, "source_name", source_name)
        if target_folder_id is not None:
            pulumi.set(__self__, "target_folder_id", target_folder_id)
        if target_name is not None:
            pulumi.set(__self__, "target_name", target_name)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Input[str]:
        """
        The serial ID of the job to be copied.
        """
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter(name="targetClusterId")
    def target_cluster_id(self) -> pulumi.Input[str]:
        """
        The cluster serial ID of the target cluster.
        """
        return pulumi.get(self, "target_cluster_id")

    @target_cluster_id.setter
    def target_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_cluster_id", value)

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional[pulumi.Input[int]]:
        """
        The type of the source job.
        """
        return pulumi.get(self, "job_type")

    @job_type.setter
    def job_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_type", value)

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job to be copied.
        """
        return pulumi.get(self, "source_name")

    @source_name.setter
    def source_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_name", value)

    @property
    @pulumi.getter(name="targetFolderId")
    def target_folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The directory ID of the new job.
        """
        return pulumi.get(self, "target_folder_id")

    @target_folder_id.setter
    def target_folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_folder_id", value)

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new job.
        """
        return pulumi.get(self, "target_name")

    @target_name.setter
    def target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_name", value)

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
class _JobCopyState:
    def __init__(__self__, *,
                 job_id: Optional[pulumi.Input[str]] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 source_name: Optional[pulumi.Input[str]] = None,
                 target_cluster_id: Optional[pulumi.Input[str]] = None,
                 target_folder_id: Optional[pulumi.Input[str]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobCopy resources.
        :param pulumi.Input[str] job_id: Copy Job ID.
        :param pulumi.Input[int] job_type: The type of the source job.
        :param pulumi.Input[str] source_id: The serial ID of the job to be copied.
        :param pulumi.Input[str] source_name: The name of the job to be copied.
        :param pulumi.Input[str] target_cluster_id: The cluster serial ID of the target cluster.
        :param pulumi.Input[str] target_folder_id: The directory ID of the new job.
        :param pulumi.Input[str] target_name: The name of the new job.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
        if job_type is not None:
            pulumi.set(__self__, "job_type", job_type)
        if source_id is not None:
            pulumi.set(__self__, "source_id", source_id)
        if source_name is not None:
            pulumi.set(__self__, "source_name", source_name)
        if target_cluster_id is not None:
            pulumi.set(__self__, "target_cluster_id", target_cluster_id)
        if target_folder_id is not None:
            pulumi.set(__self__, "target_folder_id", target_folder_id)
        if target_name is not None:
            pulumi.set(__self__, "target_name", target_name)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        Copy Job ID.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional[pulumi.Input[int]]:
        """
        The type of the source job.
        """
        return pulumi.get(self, "job_type")

    @job_type.setter
    def job_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_type", value)

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> Optional[pulumi.Input[str]]:
        """
        The serial ID of the job to be copied.
        """
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_id", value)

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job to be copied.
        """
        return pulumi.get(self, "source_name")

    @source_name.setter
    def source_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_name", value)

    @property
    @pulumi.getter(name="targetClusterId")
    def target_cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster serial ID of the target cluster.
        """
        return pulumi.get(self, "target_cluster_id")

    @target_cluster_id.setter
    def target_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_cluster_id", value)

    @property
    @pulumi.getter(name="targetFolderId")
    def target_folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The directory ID of the new job.
        """
        return pulumi.get(self, "target_folder_id")

    @target_folder_id.setter
    def target_folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_folder_id", value)

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the new job.
        """
        return pulumi.get(self, "target_name")

    @target_name.setter
    def target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_name", value)

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


class JobCopy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 source_name: Optional[pulumi.Input[str]] = None,
                 target_cluster_id: Optional[pulumi.Input[str]] = None,
                 target_folder_id: Optional[pulumi.Input[str]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a oceanus job_copy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.JobCopy("example",
            job_type=2,
            source_id="cql-0nob2hx8",
            source_name="keep_jar",
            target_cluster_id="cluster-1kcd524h",
            target_folder_id="folder-7ctl246z",
            target_name="tf_copy_example",
            work_space_id="space-2idq8wbr")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] job_type: The type of the source job.
        :param pulumi.Input[str] source_id: The serial ID of the job to be copied.
        :param pulumi.Input[str] source_name: The name of the job to be copied.
        :param pulumi.Input[str] target_cluster_id: The cluster serial ID of the target cluster.
        :param pulumi.Input[str] target_folder_id: The directory ID of the new job.
        :param pulumi.Input[str] target_name: The name of the new job.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobCopyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a oceanus job_copy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.JobCopy("example",
            job_type=2,
            source_id="cql-0nob2hx8",
            source_name="keep_jar",
            target_cluster_id="cluster-1kcd524h",
            target_folder_id="folder-7ctl246z",
            target_name="tf_copy_example",
            work_space_id="space-2idq8wbr")
        ```

        :param str resource_name: The name of the resource.
        :param JobCopyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobCopyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 source_id: Optional[pulumi.Input[str]] = None,
                 source_name: Optional[pulumi.Input[str]] = None,
                 target_cluster_id: Optional[pulumi.Input[str]] = None,
                 target_folder_id: Optional[pulumi.Input[str]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = JobCopyArgs.__new__(JobCopyArgs)

            __props__.__dict__["job_type"] = job_type
            if source_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_id'")
            __props__.__dict__["source_id"] = source_id
            __props__.__dict__["source_name"] = source_name
            if target_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_cluster_id'")
            __props__.__dict__["target_cluster_id"] = target_cluster_id
            __props__.__dict__["target_folder_id"] = target_folder_id
            __props__.__dict__["target_name"] = target_name
            __props__.__dict__["work_space_id"] = work_space_id
            __props__.__dict__["job_id"] = None
        super(JobCopy, __self__).__init__(
            'tencentcloud:Oceanus/jobCopy:JobCopy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            job_id: Optional[pulumi.Input[str]] = None,
            job_type: Optional[pulumi.Input[int]] = None,
            source_id: Optional[pulumi.Input[str]] = None,
            source_name: Optional[pulumi.Input[str]] = None,
            target_cluster_id: Optional[pulumi.Input[str]] = None,
            target_folder_id: Optional[pulumi.Input[str]] = None,
            target_name: Optional[pulumi.Input[str]] = None,
            work_space_id: Optional[pulumi.Input[str]] = None) -> 'JobCopy':
        """
        Get an existing JobCopy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_id: Copy Job ID.
        :param pulumi.Input[int] job_type: The type of the source job.
        :param pulumi.Input[str] source_id: The serial ID of the job to be copied.
        :param pulumi.Input[str] source_name: The name of the job to be copied.
        :param pulumi.Input[str] target_cluster_id: The cluster serial ID of the target cluster.
        :param pulumi.Input[str] target_folder_id: The directory ID of the new job.
        :param pulumi.Input[str] target_name: The name of the new job.
        :param pulumi.Input[str] work_space_id: Workspace SerialId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobCopyState.__new__(_JobCopyState)

        __props__.__dict__["job_id"] = job_id
        __props__.__dict__["job_type"] = job_type
        __props__.__dict__["source_id"] = source_id
        __props__.__dict__["source_name"] = source_name
        __props__.__dict__["target_cluster_id"] = target_cluster_id
        __props__.__dict__["target_folder_id"] = target_folder_id
        __props__.__dict__["target_name"] = target_name
        __props__.__dict__["work_space_id"] = work_space_id
        return JobCopy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        Copy Job ID.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> pulumi.Output[Optional[int]]:
        """
        The type of the source job.
        """
        return pulumi.get(self, "job_type")

    @property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Output[str]:
        """
        The serial ID of the job to be copied.
        """
        return pulumi.get(self, "source_id")

    @property
    @pulumi.getter(name="sourceName")
    def source_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the job to be copied.
        """
        return pulumi.get(self, "source_name")

    @property
    @pulumi.getter(name="targetClusterId")
    def target_cluster_id(self) -> pulumi.Output[str]:
        """
        The cluster serial ID of the target cluster.
        """
        return pulumi.get(self, "target_cluster_id")

    @property
    @pulumi.getter(name="targetFolderId")
    def target_folder_id(self) -> pulumi.Output[Optional[str]]:
        """
        The directory ID of the new job.
        """
        return pulumi.get(self, "target_folder_id")

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the new job.
        """
        return pulumi.get(self, "target_name")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        Workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

