# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LogConfigArgs', 'LogConfig']

@pulumi.input_type
class LogConfigArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 environment_id: pulumi.Input[str],
                 input_type: pulumi.Input[str],
                 log_type: pulumi.Input[str],
                 logset_id: pulumi.Input[str],
                 topic_id: pulumi.Input[str],
                 workload_id: pulumi.Input[str],
                 beginning_regex: Optional[pulumi.Input[str]] = None,
                 file_pattern: Optional[pulumi.Input[str]] = None,
                 log_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogConfig resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] input_type: container_stdout or container_file.
        :param pulumi.Input[str] log_type: minimalist_log or multiline_log.
        :param pulumi.Input[str] logset_id: logset.
        :param pulumi.Input[str] topic_id: topic.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        :param pulumi.Input[str] beginning_regex: regex pattern.
        :param pulumi.Input[str] file_pattern: file name pattern if container_file.
        :param pulumi.Input[str] log_path: directory if container_file.
        :param pulumi.Input[str] name: appConfig name.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "input_type", input_type)
        pulumi.set(__self__, "log_type", log_type)
        pulumi.set(__self__, "logset_id", logset_id)
        pulumi.set(__self__, "topic_id", topic_id)
        pulumi.set(__self__, "workload_id", workload_id)
        if beginning_regex is not None:
            pulumi.set(__self__, "beginning_regex", beginning_regex)
        if file_pattern is not None:
            pulumi.set(__self__, "file_pattern", file_pattern)
        if log_path is not None:
            pulumi.set(__self__, "log_path", log_path)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="inputType")
    def input_type(self) -> pulumi.Input[str]:
        """
        container_stdout or container_file.
        """
        return pulumi.get(self, "input_type")

    @input_type.setter
    def input_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "input_type", value)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Input[str]:
        """
        minimalist_log or multiline_log.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter(name="logsetId")
    def logset_id(self) -> pulumi.Input[str]:
        """
        logset.
        """
        return pulumi.get(self, "logset_id")

    @logset_id.setter
    def logset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "logset_id", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Input[str]:
        """
        topic.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> pulumi.Input[str]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")

    @workload_id.setter
    def workload_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workload_id", value)

    @property
    @pulumi.getter(name="beginningRegex")
    def beginning_regex(self) -> Optional[pulumi.Input[str]]:
        """
        regex pattern.
        """
        return pulumi.get(self, "beginning_regex")

    @beginning_regex.setter
    def beginning_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "beginning_regex", value)

    @property
    @pulumi.getter(name="filePattern")
    def file_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        file name pattern if container_file.
        """
        return pulumi.get(self, "file_pattern")

    @file_pattern.setter
    def file_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_pattern", value)

    @property
    @pulumi.getter(name="logPath")
    def log_path(self) -> Optional[pulumi.Input[str]]:
        """
        directory if container_file.
        """
        return pulumi.get(self, "log_path")

    @log_path.setter
    def log_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        appConfig name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LogConfigState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 beginning_regex: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 file_pattern: Optional[pulumi.Input[str]] = None,
                 input_type: Optional[pulumi.Input[str]] = None,
                 log_path: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 logset_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogConfig resources.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[str] beginning_regex: regex pattern.
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] file_pattern: file name pattern if container_file.
        :param pulumi.Input[str] input_type: container_stdout or container_file.
        :param pulumi.Input[str] log_path: directory if container_file.
        :param pulumi.Input[str] log_type: minimalist_log or multiline_log.
        :param pulumi.Input[str] logset_id: logset.
        :param pulumi.Input[str] name: appConfig name.
        :param pulumi.Input[str] topic_id: topic.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if beginning_regex is not None:
            pulumi.set(__self__, "beginning_regex", beginning_regex)
        if environment_id is not None:
            pulumi.set(__self__, "environment_id", environment_id)
        if file_pattern is not None:
            pulumi.set(__self__, "file_pattern", file_pattern)
        if input_type is not None:
            pulumi.set(__self__, "input_type", input_type)
        if log_path is not None:
            pulumi.set(__self__, "log_path", log_path)
        if log_type is not None:
            pulumi.set(__self__, "log_type", log_type)
        if logset_id is not None:
            pulumi.set(__self__, "logset_id", logset_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if topic_id is not None:
            pulumi.set(__self__, "topic_id", topic_id)
        if workload_id is not None:
            pulumi.set(__self__, "workload_id", workload_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="beginningRegex")
    def beginning_regex(self) -> Optional[pulumi.Input[str]]:
        """
        regex pattern.
        """
        return pulumi.get(self, "beginning_regex")

    @beginning_regex.setter
    def beginning_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "beginning_regex", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> Optional[pulumi.Input[str]]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter(name="filePattern")
    def file_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        file name pattern if container_file.
        """
        return pulumi.get(self, "file_pattern")

    @file_pattern.setter
    def file_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_pattern", value)

    @property
    @pulumi.getter(name="inputType")
    def input_type(self) -> Optional[pulumi.Input[str]]:
        """
        container_stdout or container_file.
        """
        return pulumi.get(self, "input_type")

    @input_type.setter
    def input_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_type", value)

    @property
    @pulumi.getter(name="logPath")
    def log_path(self) -> Optional[pulumi.Input[str]]:
        """
        directory if container_file.
        """
        return pulumi.get(self, "log_path")

    @log_path.setter
    def log_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_path", value)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> Optional[pulumi.Input[str]]:
        """
        minimalist_log or multiline_log.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter(name="logsetId")
    def logset_id(self) -> Optional[pulumi.Input[str]]:
        """
        logset.
        """
        return pulumi.get(self, "logset_id")

    @logset_id.setter
    def logset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logset_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        appConfig name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        topic.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_id", value)

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> Optional[pulumi.Input[str]]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")

    @workload_id.setter
    def workload_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload_id", value)


class LogConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 beginning_regex: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 file_pattern: Optional[pulumi.Input[str]] = None,
                 input_type: Optional[pulumi.Input[str]] = None,
                 log_path: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 logset_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tem logConfig

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        log_config = tencentcloud.tem.LogConfig("logConfig",
            environment_id="en-o5edaepv",
            application_id="app-3j29aa2p",
            workload_id=resource["tencentcloud_tem_workload"]["workload"]["id"],
            logset_id="b5824781-8d5b-4029-a2f7-d03c37f72bdf",
            topic_id="5a85bb6d-8e41-4e04-b7bd-c05e04782f94",
            input_type="container_stdout",
            log_type="minimalist_log")
        ```

        ## Import

        tem logConfig can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tem/logConfig:LogConfig logConfig environmentId#applicationId#name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[str] beginning_regex: regex pattern.
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] file_pattern: file name pattern if container_file.
        :param pulumi.Input[str] input_type: container_stdout or container_file.
        :param pulumi.Input[str] log_path: directory if container_file.
        :param pulumi.Input[str] log_type: minimalist_log or multiline_log.
        :param pulumi.Input[str] logset_id: logset.
        :param pulumi.Input[str] name: appConfig name.
        :param pulumi.Input[str] topic_id: topic.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tem logConfig

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        log_config = tencentcloud.tem.LogConfig("logConfig",
            environment_id="en-o5edaepv",
            application_id="app-3j29aa2p",
            workload_id=resource["tencentcloud_tem_workload"]["workload"]["id"],
            logset_id="b5824781-8d5b-4029-a2f7-d03c37f72bdf",
            topic_id="5a85bb6d-8e41-4e04-b7bd-c05e04782f94",
            input_type="container_stdout",
            log_type="minimalist_log")
        ```

        ## Import

        tem logConfig can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tem/logConfig:LogConfig logConfig environmentId#applicationId#name
        ```

        :param str resource_name: The name of the resource.
        :param LogConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 beginning_regex: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 file_pattern: Optional[pulumi.Input[str]] = None,
                 input_type: Optional[pulumi.Input[str]] = None,
                 log_path: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 logset_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 workload_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogConfigArgs.__new__(LogConfigArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["beginning_regex"] = beginning_regex
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["file_pattern"] = file_pattern
            if input_type is None and not opts.urn:
                raise TypeError("Missing required property 'input_type'")
            __props__.__dict__["input_type"] = input_type
            __props__.__dict__["log_path"] = log_path
            if log_type is None and not opts.urn:
                raise TypeError("Missing required property 'log_type'")
            __props__.__dict__["log_type"] = log_type
            if logset_id is None and not opts.urn:
                raise TypeError("Missing required property 'logset_id'")
            __props__.__dict__["logset_id"] = logset_id
            __props__.__dict__["name"] = name
            if topic_id is None and not opts.urn:
                raise TypeError("Missing required property 'topic_id'")
            __props__.__dict__["topic_id"] = topic_id
            if workload_id is None and not opts.urn:
                raise TypeError("Missing required property 'workload_id'")
            __props__.__dict__["workload_id"] = workload_id
        super(LogConfig, __self__).__init__(
            'tencentcloud:Tem/logConfig:LogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            beginning_regex: Optional[pulumi.Input[str]] = None,
            environment_id: Optional[pulumi.Input[str]] = None,
            file_pattern: Optional[pulumi.Input[str]] = None,
            input_type: Optional[pulumi.Input[str]] = None,
            log_path: Optional[pulumi.Input[str]] = None,
            log_type: Optional[pulumi.Input[str]] = None,
            logset_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            topic_id: Optional[pulumi.Input[str]] = None,
            workload_id: Optional[pulumi.Input[str]] = None) -> 'LogConfig':
        """
        Get an existing LogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: application ID.
        :param pulumi.Input[str] beginning_regex: regex pattern.
        :param pulumi.Input[str] environment_id: environment ID.
        :param pulumi.Input[str] file_pattern: file name pattern if container_file.
        :param pulumi.Input[str] input_type: container_stdout or container_file.
        :param pulumi.Input[str] log_path: directory if container_file.
        :param pulumi.Input[str] log_type: minimalist_log or multiline_log.
        :param pulumi.Input[str] logset_id: logset.
        :param pulumi.Input[str] name: appConfig name.
        :param pulumi.Input[str] topic_id: topic.
        :param pulumi.Input[str] workload_id: application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogConfigState.__new__(_LogConfigState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["beginning_regex"] = beginning_regex
        __props__.__dict__["environment_id"] = environment_id
        __props__.__dict__["file_pattern"] = file_pattern
        __props__.__dict__["input_type"] = input_type
        __props__.__dict__["log_path"] = log_path
        __props__.__dict__["log_type"] = log_type
        __props__.__dict__["logset_id"] = logset_id
        __props__.__dict__["name"] = name
        __props__.__dict__["topic_id"] = topic_id
        __props__.__dict__["workload_id"] = workload_id
        return LogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="beginningRegex")
    def beginning_regex(self) -> pulumi.Output[Optional[str]]:
        """
        regex pattern.
        """
        return pulumi.get(self, "beginning_regex")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        """
        environment ID.
        """
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter(name="filePattern")
    def file_pattern(self) -> pulumi.Output[Optional[str]]:
        """
        file name pattern if container_file.
        """
        return pulumi.get(self, "file_pattern")

    @property
    @pulumi.getter(name="inputType")
    def input_type(self) -> pulumi.Output[str]:
        """
        container_stdout or container_file.
        """
        return pulumi.get(self, "input_type")

    @property
    @pulumi.getter(name="logPath")
    def log_path(self) -> pulumi.Output[Optional[str]]:
        """
        directory if container_file.
        """
        return pulumi.get(self, "log_path")

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Output[str]:
        """
        minimalist_log or multiline_log.
        """
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter(name="logsetId")
    def logset_id(self) -> pulumi.Output[str]:
        """
        logset.
        """
        return pulumi.get(self, "logset_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        appConfig name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Output[str]:
        """
        topic.
        """
        return pulumi.get(self, "topic_id")

    @property
    @pulumi.getter(name="workloadId")
    def workload_id(self) -> pulumi.Output[str]:
        """
        application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
        """
        return pulumi.get(self, "workload_id")

