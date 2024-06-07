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

__all__ = ['DiagnoseArgs', 'Diagnose']

@pulumi.input_type
class DiagnoseArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 cron_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Diagnose resource.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[str] cron_time: Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if cron_time is not None:
            pulumi.set(__self__, "cron_time", cron_time)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="cronTime")
    def cron_time(self) -> Optional[pulumi.Input[str]]:
        """
        Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        """
        return pulumi.get(self, "cron_time")

    @cron_time.setter
    def cron_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_time", value)


@pulumi.input_type
class _DiagnoseState:
    def __init__(__self__, *,
                 cron_time: Optional[pulumi.Input[str]] = None,
                 diagnose_job_metas: Optional[pulumi.Input[Sequence[pulumi.Input['DiagnoseDiagnoseJobMetaArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 max_count: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Diagnose resources.
        :param pulumi.Input[str] cron_time: Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        :param pulumi.Input[Sequence[pulumi.Input['DiagnoseDiagnoseJobMetaArgs']]] diagnose_job_metas: Diagnostic items and meta-information of intelligent operation and maintenance.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[int] max_count: The maximum number of manual triggers per day for intelligent operation and maintenance staff.
        """
        if cron_time is not None:
            pulumi.set(__self__, "cron_time", cron_time)
        if diagnose_job_metas is not None:
            pulumi.set(__self__, "diagnose_job_metas", diagnose_job_metas)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)

    @property
    @pulumi.getter(name="cronTime")
    def cron_time(self) -> Optional[pulumi.Input[str]]:
        """
        Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        """
        return pulumi.get(self, "cron_time")

    @cron_time.setter
    def cron_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron_time", value)

    @property
    @pulumi.getter(name="diagnoseJobMetas")
    def diagnose_job_metas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DiagnoseDiagnoseJobMetaArgs']]]]:
        """
        Diagnostic items and meta-information of intelligent operation and maintenance.
        """
        return pulumi.get(self, "diagnose_job_metas")

    @diagnose_job_metas.setter
    def diagnose_job_metas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DiagnoseDiagnoseJobMetaArgs']]]]):
        pulumi.set(self, "diagnose_job_metas", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of manual triggers per day for intelligent operation and maintenance staff.
        """
        return pulumi.get(self, "max_count")

    @max_count.setter
    def max_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_count", value)


class Diagnose(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a elasticsearch diagnose

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        diagnose = tencentcloud.elasticsearch.Diagnose("diagnose",
            cron_time="15:00:00",
            instance_id="es-xxxxxx")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        es diagnose can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Elasticsearch/diagnose:Diagnose diagnose diagnose_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_time: Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        :param pulumi.Input[str] instance_id: Instance id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DiagnoseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a elasticsearch diagnose

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        diagnose = tencentcloud.elasticsearch.Diagnose("diagnose",
            cron_time="15:00:00",
            instance_id="es-xxxxxx")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        es diagnose can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Elasticsearch/diagnose:Diagnose diagnose diagnose_id
        ```

        :param str resource_name: The name of the resource.
        :param DiagnoseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DiagnoseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DiagnoseArgs.__new__(DiagnoseArgs)

            __props__.__dict__["cron_time"] = cron_time
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["diagnose_job_metas"] = None
            __props__.__dict__["max_count"] = None
        super(Diagnose, __self__).__init__(
            'tencentcloud:Elasticsearch/diagnose:Diagnose',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_time: Optional[pulumi.Input[str]] = None,
            diagnose_job_metas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DiagnoseDiagnoseJobMetaArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            max_count: Optional[pulumi.Input[int]] = None) -> 'Diagnose':
        """
        Get an existing Diagnose resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_time: Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DiagnoseDiagnoseJobMetaArgs']]]] diagnose_job_metas: Diagnostic items and meta-information of intelligent operation and maintenance.
        :param pulumi.Input[str] instance_id: Instance id.
        :param pulumi.Input[int] max_count: The maximum number of manual triggers per day for intelligent operation and maintenance staff.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DiagnoseState.__new__(_DiagnoseState)

        __props__.__dict__["cron_time"] = cron_time
        __props__.__dict__["diagnose_job_metas"] = diagnose_job_metas
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_count"] = max_count
        return Diagnose(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronTime")
    def cron_time(self) -> pulumi.Output[Optional[str]]:
        """
        Intelligent operation and maintenance staff regularly patrol the inspection time every day, the time format is HH:00:00, such as 15:00:00.
        """
        return pulumi.get(self, "cron_time")

    @property
    @pulumi.getter(name="diagnoseJobMetas")
    def diagnose_job_metas(self) -> pulumi.Output[Sequence['outputs.DiagnoseDiagnoseJobMeta']]:
        """
        Diagnostic items and meta-information of intelligent operation and maintenance.
        """
        return pulumi.get(self, "diagnose_job_metas")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> pulumi.Output[int]:
        """
        The maximum number of manual triggers per day for intelligent operation and maintenance staff.
        """
        return pulumi.get(self, "max_count")

