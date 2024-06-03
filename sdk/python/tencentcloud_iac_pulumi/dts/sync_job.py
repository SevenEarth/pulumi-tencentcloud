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

__all__ = ['SyncJobArgs', 'SyncJob']

@pulumi.input_type
class SyncJobArgs:
    def __init__(__self__, *,
                 dst_database_type: pulumi.Input[str],
                 dst_region: pulumi.Input[str],
                 pay_mode: pulumi.Input[str],
                 src_database_type: pulumi.Input[str],
                 src_region: pulumi.Input[str],
                 auto_renew: Optional[pulumi.Input[int]] = None,
                 existed_job_id: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]] = None):
        """
        The set of arguments for constructing a SyncJob resource.
        :param pulumi.Input[str] dst_database_type: destination database type.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] pay_mode: pay mode, optional value is PrePay or PostPay.
        :param pulumi.Input[str] src_database_type: source database type.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[int] auto_renew: auto renew.
        :param pulumi.Input[str] existed_job_id: existed job id.
        :param pulumi.Input[str] instance_class: instance class.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] specification: specification.
        :param pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]] tags: tags.
        """
        pulumi.set(__self__, "dst_database_type", dst_database_type)
        pulumi.set(__self__, "dst_region", dst_region)
        pulumi.set(__self__, "pay_mode", pay_mode)
        pulumi.set(__self__, "src_database_type", src_database_type)
        pulumi.set(__self__, "src_region", src_region)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if existed_job_id is not None:
            pulumi.set(__self__, "existed_job_id", existed_job_id)
        if instance_class is not None:
            pulumi.set(__self__, "instance_class", instance_class)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)
        if specification is not None:
            pulumi.set(__self__, "specification", specification)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> pulumi.Input[str]:
        """
        destination database type.
        """
        return pulumi.get(self, "dst_database_type")

    @dst_database_type.setter
    def dst_database_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "dst_database_type", value)

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> pulumi.Input[str]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @dst_region.setter
    def dst_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "dst_region", value)

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> pulumi.Input[str]:
        """
        pay mode, optional value is PrePay or PostPay.
        """
        return pulumi.get(self, "pay_mode")

    @pay_mode.setter
    def pay_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "pay_mode", value)

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> pulumi.Input[str]:
        """
        source database type.
        """
        return pulumi.get(self, "src_database_type")

    @src_database_type.setter
    def src_database_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "src_database_type", value)

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> pulumi.Input[str]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @src_region.setter
    def src_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "src_region", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[int]]:
        """
        auto renew.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="existedJobId")
    def existed_job_id(self) -> Optional[pulumi.Input[str]]:
        """
        existed job id.
        """
        return pulumi.get(self, "existed_job_id")

    @existed_job_id.setter
    def existed_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "existed_job_id", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> Optional[pulumi.Input[str]]:
        """
        instance class.
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[pulumi.Input[str]]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @job_name.setter
    def job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_name", value)

    @property
    @pulumi.getter
    def specification(self) -> Optional[pulumi.Input[str]]:
        """
        specification.
        """
        return pulumi.get(self, "specification")

    @specification.setter
    def specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "specification", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SyncJobState:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[int]] = None,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 existed_job_id: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 pay_mode: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering SyncJob resources.
        :param pulumi.Input[int] auto_renew: auto renew.
        :param pulumi.Input[str] dst_database_type: destination database type.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] existed_job_id: existed job id.
        :param pulumi.Input[str] instance_class: instance class.
        :param pulumi.Input[str] job_id: job id.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] pay_mode: pay mode, optional value is PrePay or PostPay.
        :param pulumi.Input[str] specification: specification.
        :param pulumi.Input[str] src_database_type: source database type.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]] tags: tags.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if dst_database_type is not None:
            pulumi.set(__self__, "dst_database_type", dst_database_type)
        if dst_region is not None:
            pulumi.set(__self__, "dst_region", dst_region)
        if existed_job_id is not None:
            pulumi.set(__self__, "existed_job_id", existed_job_id)
        if instance_class is not None:
            pulumi.set(__self__, "instance_class", instance_class)
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
        if job_name is not None:
            pulumi.set(__self__, "job_name", job_name)
        if pay_mode is not None:
            pulumi.set(__self__, "pay_mode", pay_mode)
        if specification is not None:
            pulumi.set(__self__, "specification", specification)
        if src_database_type is not None:
            pulumi.set(__self__, "src_database_type", src_database_type)
        if src_region is not None:
            pulumi.set(__self__, "src_region", src_region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[int]]:
        """
        auto renew.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> Optional[pulumi.Input[str]]:
        """
        destination database type.
        """
        return pulumi.get(self, "dst_database_type")

    @dst_database_type.setter
    def dst_database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst_database_type", value)

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> Optional[pulumi.Input[str]]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @dst_region.setter
    def dst_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dst_region", value)

    @property
    @pulumi.getter(name="existedJobId")
    def existed_job_id(self) -> Optional[pulumi.Input[str]]:
        """
        existed job id.
        """
        return pulumi.get(self, "existed_job_id")

    @existed_job_id.setter
    def existed_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "existed_job_id", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> Optional[pulumi.Input[str]]:
        """
        instance class.
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        """
        job id.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> Optional[pulumi.Input[str]]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @job_name.setter
    def job_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_name", value)

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> Optional[pulumi.Input[str]]:
        """
        pay mode, optional value is PrePay or PostPay.
        """
        return pulumi.get(self, "pay_mode")

    @pay_mode.setter
    def pay_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pay_mode", value)

    @property
    @pulumi.getter
    def specification(self) -> Optional[pulumi.Input[str]]:
        """
        specification.
        """
        return pulumi.get(self, "specification")

    @specification.setter
    def specification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "specification", value)

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> Optional[pulumi.Input[str]]:
        """
        source database type.
        """
        return pulumi.get(self, "src_database_type")

    @src_database_type.setter
    def src_database_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_database_type", value)

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> Optional[pulumi.Input[str]]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @src_region.setter
    def src_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "src_region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SyncJobTagArgs']]]]):
        pulumi.set(self, "tags", value)


class SyncJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[int]] = None,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 existed_job_id: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 pay_mode: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SyncJobTagArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a dts sync_job

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_job = tencentcloud.dts.SyncJob("syncJob",
            auto_renew=0,
            dst_database_type="cynosdbmysql",
            dst_region="ap-guangzhou",
            instance_class="micro",
            pay_mode="PostPay",
            src_database_type="mysql",
            src_region="ap-guangzhou",
            tags=[tencentcloud.dts.SyncJobTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_renew: auto renew.
        :param pulumi.Input[str] dst_database_type: destination database type.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] existed_job_id: existed job id.
        :param pulumi.Input[str] instance_class: instance class.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] pay_mode: pay mode, optional value is PrePay or PostPay.
        :param pulumi.Input[str] specification: specification.
        :param pulumi.Input[str] src_database_type: source database type.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SyncJobTagArgs']]]] tags: tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dts sync_job

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        sync_job = tencentcloud.dts.SyncJob("syncJob",
            auto_renew=0,
            dst_database_type="cynosdbmysql",
            dst_region="ap-guangzhou",
            instance_class="micro",
            pay_mode="PostPay",
            src_database_type="mysql",
            src_region="ap-guangzhou",
            tags=[tencentcloud.dts.SyncJobTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param SyncJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[int]] = None,
                 dst_database_type: Optional[pulumi.Input[str]] = None,
                 dst_region: Optional[pulumi.Input[str]] = None,
                 existed_job_id: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 pay_mode: Optional[pulumi.Input[str]] = None,
                 specification: Optional[pulumi.Input[str]] = None,
                 src_database_type: Optional[pulumi.Input[str]] = None,
                 src_region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SyncJobTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncJobArgs.__new__(SyncJobArgs)

            __props__.__dict__["auto_renew"] = auto_renew
            if dst_database_type is None and not opts.urn:
                raise TypeError("Missing required property 'dst_database_type'")
            __props__.__dict__["dst_database_type"] = dst_database_type
            if dst_region is None and not opts.urn:
                raise TypeError("Missing required property 'dst_region'")
            __props__.__dict__["dst_region"] = dst_region
            __props__.__dict__["existed_job_id"] = existed_job_id
            __props__.__dict__["instance_class"] = instance_class
            __props__.__dict__["job_name"] = job_name
            if pay_mode is None and not opts.urn:
                raise TypeError("Missing required property 'pay_mode'")
            __props__.__dict__["pay_mode"] = pay_mode
            __props__.__dict__["specification"] = specification
            if src_database_type is None and not opts.urn:
                raise TypeError("Missing required property 'src_database_type'")
            __props__.__dict__["src_database_type"] = src_database_type
            if src_region is None and not opts.urn:
                raise TypeError("Missing required property 'src_region'")
            __props__.__dict__["src_region"] = src_region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["job_id"] = None
        super(SyncJob, __self__).__init__(
            'tencentcloud:Dts/syncJob:SyncJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renew: Optional[pulumi.Input[int]] = None,
            dst_database_type: Optional[pulumi.Input[str]] = None,
            dst_region: Optional[pulumi.Input[str]] = None,
            existed_job_id: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            job_id: Optional[pulumi.Input[str]] = None,
            job_name: Optional[pulumi.Input[str]] = None,
            pay_mode: Optional[pulumi.Input[str]] = None,
            specification: Optional[pulumi.Input[str]] = None,
            src_database_type: Optional[pulumi.Input[str]] = None,
            src_region: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SyncJobTagArgs']]]]] = None) -> 'SyncJob':
        """
        Get an existing SyncJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_renew: auto renew.
        :param pulumi.Input[str] dst_database_type: destination database type.
        :param pulumi.Input[str] dst_region: destination region.
        :param pulumi.Input[str] existed_job_id: existed job id.
        :param pulumi.Input[str] instance_class: instance class.
        :param pulumi.Input[str] job_id: job id.
        :param pulumi.Input[str] job_name: job name.
        :param pulumi.Input[str] pay_mode: pay mode, optional value is PrePay or PostPay.
        :param pulumi.Input[str] specification: specification.
        :param pulumi.Input[str] src_database_type: source database type.
        :param pulumi.Input[str] src_region: source region.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SyncJobTagArgs']]]] tags: tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncJobState.__new__(_SyncJobState)

        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["dst_database_type"] = dst_database_type
        __props__.__dict__["dst_region"] = dst_region
        __props__.__dict__["existed_job_id"] = existed_job_id
        __props__.__dict__["instance_class"] = instance_class
        __props__.__dict__["job_id"] = job_id
        __props__.__dict__["job_name"] = job_name
        __props__.__dict__["pay_mode"] = pay_mode
        __props__.__dict__["specification"] = specification
        __props__.__dict__["src_database_type"] = src_database_type
        __props__.__dict__["src_region"] = src_region
        __props__.__dict__["tags"] = tags
        return SyncJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[int]:
        """
        auto renew.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter(name="dstDatabaseType")
    def dst_database_type(self) -> pulumi.Output[str]:
        """
        destination database type.
        """
        return pulumi.get(self, "dst_database_type")

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> pulumi.Output[str]:
        """
        destination region.
        """
        return pulumi.get(self, "dst_region")

    @property
    @pulumi.getter(name="existedJobId")
    def existed_job_id(self) -> pulumi.Output[str]:
        """
        existed job id.
        """
        return pulumi.get(self, "existed_job_id")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        """
        instance class.
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        job id.
        """
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="jobName")
    def job_name(self) -> pulumi.Output[str]:
        """
        job name.
        """
        return pulumi.get(self, "job_name")

    @property
    @pulumi.getter(name="payMode")
    def pay_mode(self) -> pulumi.Output[str]:
        """
        pay mode, optional value is PrePay or PostPay.
        """
        return pulumi.get(self, "pay_mode")

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Output[str]:
        """
        specification.
        """
        return pulumi.get(self, "specification")

    @property
    @pulumi.getter(name="srcDatabaseType")
    def src_database_type(self) -> pulumi.Output[str]:
        """
        source database type.
        """
        return pulumi.get(self, "src_database_type")

    @property
    @pulumi.getter(name="srcRegion")
    def src_region(self) -> pulumi.Output[str]:
        """
        source region.
        """
        return pulumi.get(self, "src_region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.SyncJobTag']]:
        """
        tags.
        """
        return pulumi.get(self, "tags")

