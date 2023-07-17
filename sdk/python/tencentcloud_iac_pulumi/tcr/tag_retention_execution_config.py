# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TagRetentionExecutionConfigArgs', 'TagRetentionExecutionConfig']

@pulumi.input_type
class TagRetentionExecutionConfigArgs:
    def __init__(__self__, *,
                 registry_id: pulumi.Input[str],
                 retention_id: pulumi.Input[int],
                 dry_run: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TagRetentionExecutionConfig resource.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[int] retention_id: retention id.
        :param pulumi.Input[bool] dry_run: Whether to simulate execution, the default value is false, that is, non-simulation execution.
        """
        pulumi.set(__self__, "registry_id", registry_id)
        pulumi.set(__self__, "retention_id", retention_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Input[str]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> pulumi.Input[int]:
        """
        retention id.
        """
        return pulumi.get(self, "retention_id")

    @retention_id.setter
    def retention_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "retention_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to simulate execution, the default value is false, that is, non-simulation execution.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)


@pulumi.input_type
class _TagRetentionExecutionConfigState:
    def __init__(__self__, *,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 execution_id: Optional[pulumi.Input[int]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering TagRetentionExecutionConfig resources.
        :param pulumi.Input[bool] dry_run: Whether to simulate execution, the default value is false, that is, non-simulation execution.
        :param pulumi.Input[int] execution_id: execution id.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[int] retention_id: retention id.
        """
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if execution_id is not None:
            pulumi.set(__self__, "execution_id", execution_id)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if retention_id is not None:
            pulumi.set(__self__, "retention_id", retention_id)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to simulate execution, the default value is false, that is, non-simulation execution.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> Optional[pulumi.Input[int]]:
        """
        execution id.
        """
        return pulumi.get(self, "execution_id")

    @execution_id.setter
    def execution_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "execution_id", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[str]]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> Optional[pulumi.Input[int]]:
        """
        retention id.
        """
        return pulumi.get(self, "retention_id")

    @retention_id.setter
    def retention_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_id", value)


class TagRetentionExecutionConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a tcr tag_retention_execution_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_ns = tencentcloud.tcr.Namespace("myNs",
            instance_id=tencentcloud_tcr_instance["mytcr_retention"]["id"],
            is_public=True,
            is_auto_scan=True,
            is_prevent_vul=True,
            severity="medium",
            cve_whitelist_items=[tencentcloud.tcr.NamespaceCveWhitelistItemArgs(
                cve_id="cve-xxxxx",
            )])
        my_rule = tencentcloud.tcr.TagRetentionRule("myRule",
            registry_id=tencentcloud_tcr_instance["mytcr_retention"]["id"],
            namespace_name=my_ns.name,
            retention_rule=tencentcloud.tcr.TagRetentionRuleRetentionRuleArgs(
                key="nDaysSinceLastPush",
                value=2,
            ),
            cron_setting="manual",
            disabled=True)
        tag_retention_execution_config = tencentcloud.tcr.TagRetentionExecutionConfig("tagRetentionExecutionConfig",
            registry_id=my_rule.registry_id,
            retention_id=my_rule.retention_id,
            dry_run=False)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: Whether to simulate execution, the default value is false, that is, non-simulation execution.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[int] retention_id: retention id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagRetentionExecutionConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tcr tag_retention_execution_config

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        my_ns = tencentcloud.tcr.Namespace("myNs",
            instance_id=tencentcloud_tcr_instance["mytcr_retention"]["id"],
            is_public=True,
            is_auto_scan=True,
            is_prevent_vul=True,
            severity="medium",
            cve_whitelist_items=[tencentcloud.tcr.NamespaceCveWhitelistItemArgs(
                cve_id="cve-xxxxx",
            )])
        my_rule = tencentcloud.tcr.TagRetentionRule("myRule",
            registry_id=tencentcloud_tcr_instance["mytcr_retention"]["id"],
            namespace_name=my_ns.name,
            retention_rule=tencentcloud.tcr.TagRetentionRuleRetentionRuleArgs(
                key="nDaysSinceLastPush",
                value=2,
            ),
            cron_setting="manual",
            disabled=True)
        tag_retention_execution_config = tencentcloud.tcr.TagRetentionExecutionConfig("tagRetentionExecutionConfig",
            registry_id=my_rule.registry_id,
            retention_id=my_rule.retention_id,
            dry_run=False)
        ```

        :param str resource_name: The name of the resource.
        :param TagRetentionExecutionConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagRetentionExecutionConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 registry_id: Optional[pulumi.Input[str]] = None,
                 retention_id: Optional[pulumi.Input[int]] = None,
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
            __props__ = TagRetentionExecutionConfigArgs.__new__(TagRetentionExecutionConfigArgs)

            __props__.__dict__["dry_run"] = dry_run
            if registry_id is None and not opts.urn:
                raise TypeError("Missing required property 'registry_id'")
            __props__.__dict__["registry_id"] = registry_id
            if retention_id is None and not opts.urn:
                raise TypeError("Missing required property 'retention_id'")
            __props__.__dict__["retention_id"] = retention_id
            __props__.__dict__["execution_id"] = None
        super(TagRetentionExecutionConfig, __self__).__init__(
            'tencentcloud:Tcr/tagRetentionExecutionConfig:TagRetentionExecutionConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            execution_id: Optional[pulumi.Input[int]] = None,
            registry_id: Optional[pulumi.Input[str]] = None,
            retention_id: Optional[pulumi.Input[int]] = None) -> 'TagRetentionExecutionConfig':
        """
        Get an existing TagRetentionExecutionConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: Whether to simulate execution, the default value is false, that is, non-simulation execution.
        :param pulumi.Input[int] execution_id: execution id.
        :param pulumi.Input[str] registry_id: instance id.
        :param pulumi.Input[int] retention_id: retention id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagRetentionExecutionConfigState.__new__(_TagRetentionExecutionConfigState)

        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["execution_id"] = execution_id
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["retention_id"] = retention_id
        return TagRetentionExecutionConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to simulate execution, the default value is false, that is, non-simulation execution.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Output[int]:
        """
        execution id.
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[str]:
        """
        instance id.
        """
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> pulumi.Output[int]:
        """
        retention id.
        """
        return pulumi.get(self, "retention_id")

