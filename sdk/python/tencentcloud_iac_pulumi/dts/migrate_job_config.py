# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MigrateJobConfigArgs', 'MigrateJobConfig']

@pulumi.input_type
class MigrateJobConfigArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 job_id: pulumi.Input[str],
                 complete_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MigrateJobConfig resource.
        :param pulumi.Input[str] action: The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        :param pulumi.Input[str] job_id: job id.
        :param pulumi.Input[str] complete_mode: complete mode, optional value is waitForSync or immediately.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "job_id", job_id)
        if complete_mode is not None:
            pulumi.set(__self__, "complete_mode", complete_mode)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Input[str]:
        """
        job id.
        """
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="completeMode")
    def complete_mode(self) -> Optional[pulumi.Input[str]]:
        """
        complete mode, optional value is waitForSync or immediately.
        """
        return pulumi.get(self, "complete_mode")

    @complete_mode.setter
    def complete_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "complete_mode", value)


@pulumi.input_type
class _MigrateJobConfigState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 complete_mode: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MigrateJobConfig resources.
        :param pulumi.Input[str] action: The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        :param pulumi.Input[str] complete_mode: complete mode, optional value is waitForSync or immediately.
        :param pulumi.Input[str] job_id: job id.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if complete_mode is not None:
            pulumi.set(__self__, "complete_mode", complete_mode)
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="completeMode")
    def complete_mode(self) -> Optional[pulumi.Input[str]]:
        """
        complete mode, optional value is waitForSync or immediately.
        """
        return pulumi.get(self, "complete_mode")

    @complete_mode.setter
    def complete_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "complete_mode", value)

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


class MigrateJobConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 complete_mode: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dts migrate_job_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        service = tencentcloud.dts.MigrateService("service",
            src_database_type="mysql",
            dst_database_type="cynosdbmysql",
            src_region="ap-guangzhou",
            dst_region="ap-guangzhou",
            instance_class="small",
            job_name="tf_test_xxx",
            tags=[tencentcloud.dts.MigrateServiceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        job = tencentcloud.dts.MigrateJob("job",
            service_id=service.id,
            run_mode="immediate",
            migrate_option=tencentcloud.dts.MigrateJobMigrateOptionArgs(
                database_table=tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableArgs(
                    object_mode="partial",
                    databases=[tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableDatabaseArgs(
                        db_name="tf_ci_test",
                        db_mode="partial",
                        table_mode="partial",
                        tables=[tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableDatabaseTableArgs(
                            table_name="test",
                            new_table_name="test_xxx",
                            table_edit_mode="rename",
                        )],
                    )],
                ),
            ),
            src_info=tencentcloud.dts.MigrateJobSrcInfoArgs(
                region="ap-guangzhou",
                access_type="cdb",
                database_type="mysql",
                node_type="simple",
                infos=[tencentcloud.dts.MigrateJobSrcInfoInfoArgs(
                    user="root",
                    password="xxx",
                    instance_id="id",
                )],
            ),
            dst_info=tencentcloud.dts.MigrateJobDstInfoArgs(
                region="ap-guangzhou",
                access_type="cdb",
                database_type="cynosdbmysql",
                node_type="simple",
                infos=[tencentcloud.dts.MigrateJobDstInfoInfoArgs(
                    user="user",
                    password="xxx",
                    instance_id="id",
                )],
            ),
            auto_retry_time_range_minutes=0)
        start = tencentcloud.dts.MigrateJobStartOperation("start", job_id=job.id)
        # pause the migration job
        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=start.id,
            action="pause")
        ```
        ### Continue the a migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="continue")
        ```
        ### Complete a migration job when the status is readyComplete

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="continue")
        ```
        ### Stop a running migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="stop")
        ```
        ### Isolate a stopped/canceled migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="isolate")
        ```
        ### Recover a isolated migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="recover")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        :param pulumi.Input[str] complete_mode: complete mode, optional value is waitForSync or immediately.
        :param pulumi.Input[str] job_id: job id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MigrateJobConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dts migrate_job_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        service = tencentcloud.dts.MigrateService("service",
            src_database_type="mysql",
            dst_database_type="cynosdbmysql",
            src_region="ap-guangzhou",
            dst_region="ap-guangzhou",
            instance_class="small",
            job_name="tf_test_xxx",
            tags=[tencentcloud.dts.MigrateServiceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        job = tencentcloud.dts.MigrateJob("job",
            service_id=service.id,
            run_mode="immediate",
            migrate_option=tencentcloud.dts.MigrateJobMigrateOptionArgs(
                database_table=tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableArgs(
                    object_mode="partial",
                    databases=[tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableDatabaseArgs(
                        db_name="tf_ci_test",
                        db_mode="partial",
                        table_mode="partial",
                        tables=[tencentcloud.dts.MigrateJobMigrateOptionDatabaseTableDatabaseTableArgs(
                            table_name="test",
                            new_table_name="test_xxx",
                            table_edit_mode="rename",
                        )],
                    )],
                ),
            ),
            src_info=tencentcloud.dts.MigrateJobSrcInfoArgs(
                region="ap-guangzhou",
                access_type="cdb",
                database_type="mysql",
                node_type="simple",
                infos=[tencentcloud.dts.MigrateJobSrcInfoInfoArgs(
                    user="root",
                    password="xxx",
                    instance_id="id",
                )],
            ),
            dst_info=tencentcloud.dts.MigrateJobDstInfoArgs(
                region="ap-guangzhou",
                access_type="cdb",
                database_type="cynosdbmysql",
                node_type="simple",
                infos=[tencentcloud.dts.MigrateJobDstInfoInfoArgs(
                    user="user",
                    password="xxx",
                    instance_id="id",
                )],
            ),
            auto_retry_time_range_minutes=0)
        start = tencentcloud.dts.MigrateJobStartOperation("start", job_id=job.id)
        # pause the migration job
        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=start.id,
            action="pause")
        ```
        ### Continue the a migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="continue")
        ```
        ### Complete a migration job when the status is readyComplete

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="continue")
        ```
        ### Stop a running migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="stop")
        ```
        ### Isolate a stopped/canceled migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="isolate")
        ```
        ### Recover a isolated migration job

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        config = tencentcloud.dts.MigrateJobConfig("config",
            job_id=tencentcloud_dts_migrate_job_start_operation["start"]["id"],
            action="recover")
        ```

        :param str resource_name: The name of the resource.
        :param MigrateJobConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MigrateJobConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 complete_mode: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = MigrateJobConfigArgs.__new__(MigrateJobConfigArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["complete_mode"] = complete_mode
            if job_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_id'")
            __props__.__dict__["job_id"] = job_id
        super(MigrateJobConfig, __self__).__init__(
            'tencentcloud:Dts/migrateJobConfig:MigrateJobConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            complete_mode: Optional[pulumi.Input[str]] = None,
            job_id: Optional[pulumi.Input[str]] = None) -> 'MigrateJobConfig':
        """
        Get an existing MigrateJobConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        :param pulumi.Input[str] complete_mode: complete mode, optional value is waitForSync or immediately.
        :param pulumi.Input[str] job_id: job id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MigrateJobConfigState.__new__(_MigrateJobConfigState)

        __props__.__dict__["action"] = action
        __props__.__dict__["complete_mode"] = complete_mode
        __props__.__dict__["job_id"] = job_id
        return MigrateJobConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The operation want to perform. Valid values are: `pause`, `continue`, `complete`, `recover`,`stop`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="completeMode")
    def complete_mode(self) -> pulumi.Output[Optional[str]]:
        """
        complete mode, optional value is waitForSync or immediately.
        """
        return pulumi.get(self, "complete_mode")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        """
        job id.
        """
        return pulumi.get(self, "job_id")

