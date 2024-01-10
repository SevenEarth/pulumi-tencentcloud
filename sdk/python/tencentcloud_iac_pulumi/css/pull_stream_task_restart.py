# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PullStreamTaskRestartArgs', 'PullStreamTaskRestart']

@pulumi.input_type
class PullStreamTaskRestartArgs:
    def __init__(__self__, *,
                 operator: pulumi.Input[str],
                 task_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a PullStreamTaskRestart resource.
        :param pulumi.Input[str] operator: Task operator.
        :param pulumi.Input[str] task_id: Task Id.
        """
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "task_id", task_id)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[str]:
        """
        Task operator.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Input[str]:
        """
        Task Id.
        """
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "task_id", value)


@pulumi.input_type
class _PullStreamTaskRestartState:
    def __init__(__self__, *,
                 operator: Optional[pulumi.Input[str]] = None,
                 task_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PullStreamTaskRestart resources.
        :param pulumi.Input[str] operator: Task operator.
        :param pulumi.Input[str] task_id: Task Id.
        """
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        Task operator.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> Optional[pulumi.Input[str]]:
        """
        Task Id.
        """
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_id", value)


class PullStreamTaskRestart(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 task_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a css restart_push_task

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        restart_push_task = tencentcloud.css.PullStreamTaskRestart("restartPushTask",
            operator="tf-test",
            task_id="3573")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] operator: Task operator.
        :param pulumi.Input[str] task_id: Task Id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PullStreamTaskRestartArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a css restart_push_task

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        restart_push_task = tencentcloud.css.PullStreamTaskRestart("restartPushTask",
            operator="tf-test",
            task_id="3573")
        ```

        :param str resource_name: The name of the resource.
        :param PullStreamTaskRestartArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PullStreamTaskRestartArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 task_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PullStreamTaskRestartArgs.__new__(PullStreamTaskRestartArgs)

            if operator is None and not opts.urn:
                raise TypeError("Missing required property 'operator'")
            __props__.__dict__["operator"] = operator
            if task_id is None and not opts.urn:
                raise TypeError("Missing required property 'task_id'")
            __props__.__dict__["task_id"] = task_id
        super(PullStreamTaskRestart, __self__).__init__(
            'tencentcloud:Css/pullStreamTaskRestart:PullStreamTaskRestart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            operator: Optional[pulumi.Input[str]] = None,
            task_id: Optional[pulumi.Input[str]] = None) -> 'PullStreamTaskRestart':
        """
        Get an existing PullStreamTaskRestart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] operator: Task operator.
        :param pulumi.Input[str] task_id: Task Id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PullStreamTaskRestartState.__new__(_PullStreamTaskRestartState)

        __props__.__dict__["operator"] = operator
        __props__.__dict__["task_id"] = task_id
        return PullStreamTaskRestart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Output[str]:
        """
        Task operator.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Output[str]:
        """
        Task Id.
        """
        return pulumi.get(self, "task_id")

