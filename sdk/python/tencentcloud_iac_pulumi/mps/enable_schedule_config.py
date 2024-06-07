# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EnableScheduleConfigArgs', 'EnableScheduleConfig']

@pulumi.input_type
class EnableScheduleConfigArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 schedule_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a EnableScheduleConfig resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] schedule_id: The scheme ID.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "schedule_id", schedule_id)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> pulumi.Input[int]:
        """
        The scheme ID.
        """
        return pulumi.get(self, "schedule_id")

    @schedule_id.setter
    def schedule_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "schedule_id", value)


@pulumi.input_type
class _EnableScheduleConfigState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 schedule_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering EnableScheduleConfig resources.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] schedule_id: The scheme ID.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if schedule_id is not None:
            pulumi.set(__self__, "schedule_id", schedule_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> Optional[pulumi.Input[int]]:
        """
        The scheme ID.
        """
        return pulumi.get(self, "schedule_id")

    @schedule_id.setter
    def schedule_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "schedule_id", value)


class EnableScheduleConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 schedule_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mps enable_schedule_config

        ## Example Usage

        ### Enable the mps schedule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        object = tencentcloud.Cos.get_bucket_object(bucket=f"keep-bucket-{local['app_id']}",
            key="/mps-test/test.mov")
        output = tencentcloud.cos.Bucket("output",
            bucket=f"tf-bucket-mps-schedule-config-output1-{local['app_id']}",
            force_clean=True,
            acl="public-read")
        example = tencentcloud.mps.Schedule("example",
            schedule_name="tf_test_mps_schedule_config",
            trigger=tencentcloud.mps.ScheduleTriggerArgs(
                type="CosFileUpload",
                cos_file_upload_trigger=tencentcloud.mps.ScheduleTriggerCosFileUploadTriggerArgs(
                    bucket=object.bucket,
                    region="%s",
                    dir="/upload/",
                    formats=[
                        "flv",
                        "mov",
                    ],
                ),
            ),
            activities=[
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="input",
                    reardrive_indices=[
                        1,
                        2,
                    ],
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[3],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[
                        6,
                        7,
                    ],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[
                        4,
                        5,
                    ],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[8],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[9],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="output",
                ),
            ],
            output_storage=tencentcloud.mps.ScheduleOutputStorageArgs(
                type="COS",
                cos_output_storage=tencentcloud.mps.ScheduleOutputStorageCosOutputStorageArgs(
                    bucket=output.bucket,
                    region="%s",
                ),
            ),
            output_dir="output/")
        config = tencentcloud.mps.EnableScheduleConfig("config",
            schedule_id=example.id,
            enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ### Disable the mps schedule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.mps.EnableScheduleConfig("config",
            schedule_id=tencentcloud_mps_schedule["example"]["id"],
            enabled=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        mps enable_schedule_config can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mps/enableScheduleConfig:EnableScheduleConfig enable_schedule_config enable_schedule_config_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] schedule_id: The scheme ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnableScheduleConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mps enable_schedule_config

        ## Example Usage

        ### Enable the mps schedule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        object = tencentcloud.Cos.get_bucket_object(bucket=f"keep-bucket-{local['app_id']}",
            key="/mps-test/test.mov")
        output = tencentcloud.cos.Bucket("output",
            bucket=f"tf-bucket-mps-schedule-config-output1-{local['app_id']}",
            force_clean=True,
            acl="public-read")
        example = tencentcloud.mps.Schedule("example",
            schedule_name="tf_test_mps_schedule_config",
            trigger=tencentcloud.mps.ScheduleTriggerArgs(
                type="CosFileUpload",
                cos_file_upload_trigger=tencentcloud.mps.ScheduleTriggerCosFileUploadTriggerArgs(
                    bucket=object.bucket,
                    region="%s",
                    dir="/upload/",
                    formats=[
                        "flv",
                        "mov",
                    ],
                ),
            ),
            activities=[
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="input",
                    reardrive_indices=[
                        1,
                        2,
                    ],
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[3],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[
                        6,
                        7,
                    ],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[
                        4,
                        5,
                    ],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[8],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[9],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="action-trans",
                    reardrive_indices=[10],
                    activity_para=tencentcloud.mps.ScheduleActivityActivityParaArgs(
                        transcode_task=tencentcloud.mps.ScheduleActivityActivityParaTranscodeTaskArgs(
                            definition=10,
                        ),
                    ),
                ),
                tencentcloud.mps.ScheduleActivityArgs(
                    activity_type="output",
                ),
            ],
            output_storage=tencentcloud.mps.ScheduleOutputStorageArgs(
                type="COS",
                cos_output_storage=tencentcloud.mps.ScheduleOutputStorageCosOutputStorageArgs(
                    bucket=output.bucket,
                    region="%s",
                ),
            ),
            output_dir="output/")
        config = tencentcloud.mps.EnableScheduleConfig("config",
            schedule_id=example.id,
            enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ### Disable the mps schedule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.mps.EnableScheduleConfig("config",
            schedule_id=tencentcloud_mps_schedule["example"]["id"],
            enabled=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        mps enable_schedule_config can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mps/enableScheduleConfig:EnableScheduleConfig enable_schedule_config enable_schedule_config_id
        ```

        :param str resource_name: The name of the resource.
        :param EnableScheduleConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnableScheduleConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 schedule_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnableScheduleConfigArgs.__new__(EnableScheduleConfigArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if schedule_id is None and not opts.urn:
                raise TypeError("Missing required property 'schedule_id'")
            __props__.__dict__["schedule_id"] = schedule_id
        super(EnableScheduleConfig, __self__).__init__(
            'tencentcloud:Mps/enableScheduleConfig:EnableScheduleConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            schedule_id: Optional[pulumi.Input[int]] = None) -> 'EnableScheduleConfig':
        """
        Get an existing EnableScheduleConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: true: enable; false: disable.
        :param pulumi.Input[int] schedule_id: The scheme ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnableScheduleConfigState.__new__(_EnableScheduleConfigState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["schedule_id"] = schedule_id
        return EnableScheduleConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        true: enable; false: disable.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="scheduleId")
    def schedule_id(self) -> pulumi.Output[int]:
        """
        The scheme ID.
        """
        return pulumi.get(self, "schedule_id")

