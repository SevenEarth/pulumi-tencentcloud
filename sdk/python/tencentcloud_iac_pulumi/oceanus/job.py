# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobArgs', 'Job']

@pulumi.input_type
class JobArgs:
    def __init__(__self__, *,
                 cluster_type: pulumi.Input[int],
                 job_type: pulumi.Input[int],
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cu_mem: Optional[pulumi.Input[int]] = None,
                 flink_version: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Job resource.
        :param pulumi.Input[int] cluster_type: The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        :param pulumi.Input[int] job_type: The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        :param pulumi.Input[str] cluster_id: When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        :param pulumi.Input[int] cu_mem: Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        :param pulumi.Input[str] flink_version: The Flink version that the job runs.
        :param pulumi.Input[str] folder_id: The folder ID to which the job name belongs. The root directory is root.
        :param pulumi.Input[str] name: The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        :param pulumi.Input[str] remark: The remark information of the job. It can be set arbitrarily.
        :param pulumi.Input[str] work_space_id: The workspace SerialId.
        """
        pulumi.set(__self__, "cluster_type", cluster_type)
        pulumi.set(__self__, "job_type", job_type)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cu_mem is not None:
            pulumi.set(__self__, "cu_mem", cu_mem)
        if flink_version is not None:
            pulumi.set(__self__, "flink_version", flink_version)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Input[int]:
        """
        The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> pulumi.Input[int]:
        """
        The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        """
        return pulumi.get(self, "job_type")

    @job_type.setter
    def job_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "job_type", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="cuMem")
    def cu_mem(self) -> Optional[pulumi.Input[int]]:
        """
        Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        """
        return pulumi.get(self, "cu_mem")

    @cu_mem.setter
    def cu_mem(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cu_mem", value)

    @property
    @pulumi.getter(name="flinkVersion")
    def flink_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Flink version that the job runs.
        """
        return pulumi.get(self, "flink_version")

    @flink_version.setter
    def flink_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flink_version", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The folder ID to which the job name belongs. The root directory is root.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark information of the job. It can be set arbitrarily.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


@pulumi.input_type
class _JobState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[int]] = None,
                 cu_mem: Optional[pulumi.Input[int]] = None,
                 flink_version: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Job resources.
        :param pulumi.Input[str] cluster_id: When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        :param pulumi.Input[int] cluster_type: The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        :param pulumi.Input[int] cu_mem: Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        :param pulumi.Input[str] flink_version: The Flink version that the job runs.
        :param pulumi.Input[str] folder_id: The folder ID to which the job name belongs. The root directory is root.
        :param pulumi.Input[int] job_type: The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        :param pulumi.Input[str] name: The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        :param pulumi.Input[str] remark: The remark information of the job. It can be set arbitrarily.
        :param pulumi.Input[str] work_space_id: The workspace SerialId.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if cu_mem is not None:
            pulumi.set(__self__, "cu_mem", cu_mem)
        if flink_version is not None:
            pulumi.set(__self__, "flink_version", flink_version)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if job_type is not None:
            pulumi.set(__self__, "job_type", job_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remark is not None:
            pulumi.set(__self__, "remark", remark)
        if work_space_id is not None:
            pulumi.set(__self__, "work_space_id", work_space_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[pulumi.Input[int]]:
        """
        The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="cuMem")
    def cu_mem(self) -> Optional[pulumi.Input[int]]:
        """
        Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        """
        return pulumi.get(self, "cu_mem")

    @cu_mem.setter
    def cu_mem(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cu_mem", value)

    @property
    @pulumi.getter(name="flinkVersion")
    def flink_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Flink version that the job runs.
        """
        return pulumi.get(self, "flink_version")

    @flink_version.setter
    def flink_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flink_version", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The folder ID to which the job name belongs. The root directory is root.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional[pulumi.Input[int]]:
        """
        The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        """
        return pulumi.get(self, "job_type")

    @job_type.setter
    def job_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def remark(self) -> Optional[pulumi.Input[str]]:
        """
        The remark information of the job. It can be set arbitrarily.
        """
        return pulumi.get(self, "remark")

    @remark.setter
    def remark(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remark", value)

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        The workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

    @work_space_id.setter
    def work_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "work_space_id", value)


class Job(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[int]] = None,
                 cu_mem: Optional[pulumi.Input[int]] = None,
                 flink_version: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a oceanus job

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.Job("example",
            cluster_id="cluster-1kcd524h",
            cluster_type=2,
            cu_mem=4,
            flink_version="Flink-1.16",
            folder_id="folder-7ctl246z",
            job_type=1,
            remark="remark.",
            work_space_id="space-2idq8wbr")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        :param pulumi.Input[int] cluster_type: The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        :param pulumi.Input[int] cu_mem: Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        :param pulumi.Input[str] flink_version: The Flink version that the job runs.
        :param pulumi.Input[str] folder_id: The folder ID to which the job name belongs. The root directory is root.
        :param pulumi.Input[int] job_type: The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        :param pulumi.Input[str] name: The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        :param pulumi.Input[str] remark: The remark information of the job. It can be set arbitrarily.
        :param pulumi.Input[str] work_space_id: The workspace SerialId.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a oceanus job

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.oceanus.Job("example",
            cluster_id="cluster-1kcd524h",
            cluster_type=2,
            cu_mem=4,
            flink_version="Flink-1.16",
            folder_id="folder-7ctl246z",
            job_type=1,
            remark="remark.",
            work_space_id="space-2idq8wbr")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param JobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[int]] = None,
                 cu_mem: Optional[pulumi.Input[int]] = None,
                 flink_version: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 job_type: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remark: Optional[pulumi.Input[str]] = None,
                 work_space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobArgs.__new__(JobArgs)

            __props__.__dict__["cluster_id"] = cluster_id
            if cluster_type is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_type'")
            __props__.__dict__["cluster_type"] = cluster_type
            __props__.__dict__["cu_mem"] = cu_mem
            __props__.__dict__["flink_version"] = flink_version
            __props__.__dict__["folder_id"] = folder_id
            if job_type is None and not opts.urn:
                raise TypeError("Missing required property 'job_type'")
            __props__.__dict__["job_type"] = job_type
            __props__.__dict__["name"] = name
            __props__.__dict__["remark"] = remark
            __props__.__dict__["work_space_id"] = work_space_id
        super(Job, __self__).__init__(
            'tencentcloud:Oceanus/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_type: Optional[pulumi.Input[int]] = None,
            cu_mem: Optional[pulumi.Input[int]] = None,
            flink_version: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            job_type: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remark: Optional[pulumi.Input[str]] = None,
            work_space_id: Optional[pulumi.Input[str]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        :param pulumi.Input[int] cluster_type: The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        :param pulumi.Input[int] cu_mem: Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        :param pulumi.Input[str] flink_version: The Flink version that the job runs.
        :param pulumi.Input[str] folder_id: The folder ID to which the job name belongs. The root directory is root.
        :param pulumi.Input[int] job_type: The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        :param pulumi.Input[str] name: The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        :param pulumi.Input[str] remark: The remark information of the job. It can be set arbitrarily.
        :param pulumi.Input[str] work_space_id: The workspace SerialId.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobState.__new__(_JobState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["cluster_type"] = cluster_type
        __props__.__dict__["cu_mem"] = cu_mem
        __props__.__dict__["flink_version"] = flink_version
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["job_type"] = job_type
        __props__.__dict__["name"] = name
        __props__.__dict__["remark"] = remark
        __props__.__dict__["work_space_id"] = work_space_id
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[Optional[str]]:
        """
        When ClusterType=2, it is required to specify the ID of the exclusive cluster to which the job is submitted.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Output[int]:
        """
        The type of the cluster. 1 indicates shared cluster, and 2 indicates exclusive cluster.
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="cuMem")
    def cu_mem(self) -> pulumi.Output[Optional[int]]:
        """
        Set the memory specification of each CU, in GB. It supports 2, 4, 8, and 16 (which needs to apply for the whitelist before use). The default is 4, that is, 1 CU corresponds to 4 GB of running memory.
        """
        return pulumi.get(self, "cu_mem")

    @property
    @pulumi.getter(name="flinkVersion")
    def flink_version(self) -> pulumi.Output[Optional[str]]:
        """
        The Flink version that the job runs.
        """
        return pulumi.get(self, "flink_version")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[Optional[str]]:
        """
        The folder ID to which the job name belongs. The root directory is root.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> pulumi.Output[int]:
        """
        The type of the job. 1 indicates SQL job, and 2 indicates JAR job.
        """
        return pulumi.get(self, "job_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the job. It can be composed of Chinese, English, numbers, hyphens (-), underscores (_), and periods (.), and the length cannot exceed 50 characters. Note that the job name cannot be the same as an existing job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def remark(self) -> pulumi.Output[Optional[str]]:
        """
        The remark information of the job. It can be set arbitrarily.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="workSpaceId")
    def work_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        The workspace SerialId.
        """
        return pulumi.get(self, "work_space_id")

