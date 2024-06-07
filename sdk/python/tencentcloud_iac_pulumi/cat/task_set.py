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

__all__ = ['TaskSetArgs', 'TaskSet']

@pulumi.input_type
class TaskSetArgs:
    def __init__(__self__, *,
                 batch_tasks: pulumi.Input['TaskSetBatchTasksArgs'],
                 interval: pulumi.Input[int],
                 nodes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 parameters: pulumi.Input[str],
                 task_category: pulumi.Input[int],
                 task_type: pulumi.Input[int],
                 cron: Optional[pulumi.Input[str]] = None,
                 node_ip_type: Optional[pulumi.Input[int]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a TaskSet resource.
        :param pulumi.Input['TaskSetBatchTasksArgs'] batch_tasks: Batch task name address.
        :param pulumi.Input[int] interval: Task interval minutes in (1,5,10,15,30,60,120,240).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nodes: Task Nodes.
        :param pulumi.Input[str] parameters: tasks parameters.
        :param pulumi.Input[int] task_category: Task category,1:PC,2:Mobile.
        :param pulumi.Input[int] task_type: Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        :param pulumi.Input[str] cron: Timer task cron expression.
        :param pulumi.Input[int] node_ip_type: `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        :param pulumi.Input[str] operate: The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        """
        pulumi.set(__self__, "batch_tasks", batch_tasks)
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "nodes", nodes)
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "task_category", task_category)
        pulumi.set(__self__, "task_type", task_type)
        if cron is not None:
            pulumi.set(__self__, "cron", cron)
        if node_ip_type is not None:
            pulumi.set(__self__, "node_ip_type", node_ip_type)
        if operate is not None:
            pulumi.set(__self__, "operate", operate)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="batchTasks")
    def batch_tasks(self) -> pulumi.Input['TaskSetBatchTasksArgs']:
        """
        Batch task name address.
        """
        return pulumi.get(self, "batch_tasks")

    @batch_tasks.setter
    def batch_tasks(self, value: pulumi.Input['TaskSetBatchTasksArgs']):
        pulumi.set(self, "batch_tasks", value)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[int]:
        """
        Task interval minutes in (1,5,10,15,30,60,120,240).
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Task Nodes.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Input[str]:
        """
        tasks parameters.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="taskCategory")
    def task_category(self) -> pulumi.Input[int]:
        """
        Task category,1:PC,2:Mobile.
        """
        return pulumi.get(self, "task_category")

    @task_category.setter
    def task_category(self, value: pulumi.Input[int]):
        pulumi.set(self, "task_category", value)

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> pulumi.Input[int]:
        """
        Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        return pulumi.get(self, "task_type")

    @task_type.setter
    def task_type(self, value: pulumi.Input[int]):
        pulumi.set(self, "task_type", value)

    @property
    @pulumi.getter
    def cron(self) -> Optional[pulumi.Input[str]]:
        """
        Timer task cron expression.
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter(name="nodeIpType")
    def node_ip_type(self) -> Optional[pulumi.Input[int]]:
        """
        `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        """
        return pulumi.get(self, "node_ip_type")

    @node_ip_type.setter
    def node_ip_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_ip_type", value)

    @property
    @pulumi.getter
    def operate(self) -> Optional[pulumi.Input[str]]:
        """
        The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        """
        return pulumi.get(self, "operate")

    @operate.setter
    def operate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operate", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TaskSetState:
    def __init__(__self__, *,
                 batch_tasks: Optional[pulumi.Input['TaskSetBatchTasksArgs']] = None,
                 cron: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 node_ip_type: Optional[pulumi.Input[int]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 task_category: Optional[pulumi.Input[int]] = None,
                 task_id: Optional[pulumi.Input[str]] = None,
                 task_type: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering TaskSet resources.
        :param pulumi.Input['TaskSetBatchTasksArgs'] batch_tasks: Batch task name address.
        :param pulumi.Input[str] cron: Timer task cron expression.
        :param pulumi.Input[int] interval: Task interval minutes in (1,5,10,15,30,60,120,240).
        :param pulumi.Input[int] node_ip_type: `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nodes: Task Nodes.
        :param pulumi.Input[str] operate: The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        :param pulumi.Input[str] parameters: tasks parameters.
        :param pulumi.Input[int] status: Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] task_category: Task category,1:PC,2:Mobile.
        :param pulumi.Input[str] task_id: Task Id.
        :param pulumi.Input[int] task_type: Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        if batch_tasks is not None:
            pulumi.set(__self__, "batch_tasks", batch_tasks)
        if cron is not None:
            pulumi.set(__self__, "cron", cron)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if node_ip_type is not None:
            pulumi.set(__self__, "node_ip_type", node_ip_type)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if operate is not None:
            pulumi.set(__self__, "operate", operate)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if task_category is not None:
            pulumi.set(__self__, "task_category", task_category)
        if task_id is not None:
            pulumi.set(__self__, "task_id", task_id)
        if task_type is not None:
            pulumi.set(__self__, "task_type", task_type)

    @property
    @pulumi.getter(name="batchTasks")
    def batch_tasks(self) -> Optional[pulumi.Input['TaskSetBatchTasksArgs']]:
        """
        Batch task name address.
        """
        return pulumi.get(self, "batch_tasks")

    @batch_tasks.setter
    def batch_tasks(self, value: Optional[pulumi.Input['TaskSetBatchTasksArgs']]):
        pulumi.set(self, "batch_tasks", value)

    @property
    @pulumi.getter
    def cron(self) -> Optional[pulumi.Input[str]]:
        """
        Timer task cron expression.
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Task interval minutes in (1,5,10,15,30,60,120,240).
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="nodeIpType")
    def node_ip_type(self) -> Optional[pulumi.Input[int]]:
        """
        `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        """
        return pulumi.get(self, "node_ip_type")

    @node_ip_type.setter
    def node_ip_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "node_ip_type", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Task Nodes.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def operate(self) -> Optional[pulumi.Input[str]]:
        """
        The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        """
        return pulumi.get(self, "operate")

    @operate.setter
    def operate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operate", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[str]]:
        """
        tasks parameters.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="taskCategory")
    def task_category(self) -> Optional[pulumi.Input[int]]:
        """
        Task category,1:PC,2:Mobile.
        """
        return pulumi.get(self, "task_category")

    @task_category.setter
    def task_category(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "task_category", value)

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

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> Optional[pulumi.Input[int]]:
        """
        Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        return pulumi.get(self, "task_type")

    @task_type.setter
    def task_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "task_type", value)


class TaskSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_tasks: Optional[pulumi.Input[pulumi.InputType['TaskSetBatchTasksArgs']]] = None,
                 cron: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 node_ip_type: Optional[pulumi.Input[int]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 task_category: Optional[pulumi.Input[int]] = None,
                 task_type: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a cat task_set

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import tencentcloud_iac_pulumi as tencentcloud

        task_set = tencentcloud.cat.TaskSet("taskSet",
            batch_tasks=tencentcloud.cat.TaskSetBatchTasksArgs(
                name="demo",
                target_address="http://www.baidu.com",
            ),
            task_type=5,
            nodes=[
                "12136",
                "12137",
                "12138",
                "12141",
                "12144",
            ],
            interval=6,
            parameters=json.dumps({
                "ipType": 0,
                "grabBag": 0,
                "filterIp": 0,
                "netIcmpOn": 1,
                "netIcmpActivex": 0,
                "netIcmpTimeout": 20,
                "netIcmpInterval": 0.5,
                "netIcmpNum": 20,
                "netIcmpSize": 32,
                "netIcmpDataCut": 1,
                "netDnsOn": 1,
                "netDnsTimeout": 5,
                "netDnsQuerymethod": 1,
                "netDnsNs": "",
                "netDigOn": 1,
                "netDnsServer": 2,
                "netTracertOn": 1,
                "netTracertTimeout": 60,
                "netTracertNum": 30,
                "whiteList": "",
                "blackList": "",
                "netIcmpActivexStr": "",
            }),
            task_category=1,
            cron="* 0-6 * * *",
            tags={
                "createdBy": "terraform",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cat task_set can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cat/taskSet:TaskSet task_set taskSet_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TaskSetBatchTasksArgs']] batch_tasks: Batch task name address.
        :param pulumi.Input[str] cron: Timer task cron expression.
        :param pulumi.Input[int] interval: Task interval minutes in (1,5,10,15,30,60,120,240).
        :param pulumi.Input[int] node_ip_type: `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nodes: Task Nodes.
        :param pulumi.Input[str] operate: The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        :param pulumi.Input[str] parameters: tasks parameters.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] task_category: Task category,1:PC,2:Mobile.
        :param pulumi.Input[int] task_type: Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a cat task_set

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import tencentcloud_iac_pulumi as tencentcloud

        task_set = tencentcloud.cat.TaskSet("taskSet",
            batch_tasks=tencentcloud.cat.TaskSetBatchTasksArgs(
                name="demo",
                target_address="http://www.baidu.com",
            ),
            task_type=5,
            nodes=[
                "12136",
                "12137",
                "12138",
                "12141",
                "12144",
            ],
            interval=6,
            parameters=json.dumps({
                "ipType": 0,
                "grabBag": 0,
                "filterIp": 0,
                "netIcmpOn": 1,
                "netIcmpActivex": 0,
                "netIcmpTimeout": 20,
                "netIcmpInterval": 0.5,
                "netIcmpNum": 20,
                "netIcmpSize": 32,
                "netIcmpDataCut": 1,
                "netDnsOn": 1,
                "netDnsTimeout": 5,
                "netDnsQuerymethod": 1,
                "netDnsNs": "",
                "netDigOn": 1,
                "netDnsServer": 2,
                "netTracertOn": 1,
                "netTracertTimeout": 60,
                "netTracertNum": 30,
                "whiteList": "",
                "blackList": "",
                "netIcmpActivexStr": "",
            }),
            task_category=1,
            cron="* 0-6 * * *",
            tags={
                "createdBy": "terraform",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        cat task_set can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Cat/taskSet:TaskSet task_set taskSet_id
        ```

        :param str resource_name: The name of the resource.
        :param TaskSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 batch_tasks: Optional[pulumi.Input[pulumi.InputType['TaskSetBatchTasksArgs']]] = None,
                 cron: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 node_ip_type: Optional[pulumi.Input[int]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 operate: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 task_category: Optional[pulumi.Input[int]] = None,
                 task_type: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TaskSetArgs.__new__(TaskSetArgs)

            if batch_tasks is None and not opts.urn:
                raise TypeError("Missing required property 'batch_tasks'")
            __props__.__dict__["batch_tasks"] = batch_tasks
            __props__.__dict__["cron"] = cron
            if interval is None and not opts.urn:
                raise TypeError("Missing required property 'interval'")
            __props__.__dict__["interval"] = interval
            __props__.__dict__["node_ip_type"] = node_ip_type
            if nodes is None and not opts.urn:
                raise TypeError("Missing required property 'nodes'")
            __props__.__dict__["nodes"] = nodes
            __props__.__dict__["operate"] = operate
            if parameters is None and not opts.urn:
                raise TypeError("Missing required property 'parameters'")
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["tags"] = tags
            if task_category is None and not opts.urn:
                raise TypeError("Missing required property 'task_category'")
            __props__.__dict__["task_category"] = task_category
            if task_type is None and not opts.urn:
                raise TypeError("Missing required property 'task_type'")
            __props__.__dict__["task_type"] = task_type
            __props__.__dict__["status"] = None
            __props__.__dict__["task_id"] = None
        super(TaskSet, __self__).__init__(
            'tencentcloud:Cat/taskSet:TaskSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            batch_tasks: Optional[pulumi.Input[pulumi.InputType['TaskSetBatchTasksArgs']]] = None,
            cron: Optional[pulumi.Input[str]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            node_ip_type: Optional[pulumi.Input[int]] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            operate: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            task_category: Optional[pulumi.Input[int]] = None,
            task_id: Optional[pulumi.Input[str]] = None,
            task_type: Optional[pulumi.Input[int]] = None) -> 'TaskSet':
        """
        Get an existing TaskSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TaskSetBatchTasksArgs']] batch_tasks: Batch task name address.
        :param pulumi.Input[str] cron: Timer task cron expression.
        :param pulumi.Input[int] interval: Task interval minutes in (1,5,10,15,30,60,120,240).
        :param pulumi.Input[int] node_ip_type: `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] nodes: Task Nodes.
        :param pulumi.Input[str] operate: The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        :param pulumi.Input[str] parameters: tasks parameters.
        :param pulumi.Input[int] status: Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        :param pulumi.Input[Mapping[str, Any]] tags: Tag description list.
        :param pulumi.Input[int] task_category: Task category,1:PC,2:Mobile.
        :param pulumi.Input[str] task_id: Task Id.
        :param pulumi.Input[int] task_type: Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TaskSetState.__new__(_TaskSetState)

        __props__.__dict__["batch_tasks"] = batch_tasks
        __props__.__dict__["cron"] = cron
        __props__.__dict__["interval"] = interval
        __props__.__dict__["node_ip_type"] = node_ip_type
        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["operate"] = operate
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["task_category"] = task_category
        __props__.__dict__["task_id"] = task_id
        __props__.__dict__["task_type"] = task_type
        return TaskSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="batchTasks")
    def batch_tasks(self) -> pulumi.Output['outputs.TaskSetBatchTasks']:
        """
        Batch task name address.
        """
        return pulumi.get(self, "batch_tasks")

    @property
    @pulumi.getter
    def cron(self) -> pulumi.Output[Optional[str]]:
        """
        Timer task cron expression.
        """
        return pulumi.get(self, "cron")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[int]:
        """
        Task interval minutes in (1,5,10,15,30,60,120,240).
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="nodeIpType")
    def node_ip_type(self) -> pulumi.Output[int]:
        """
        `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        """
        return pulumi.get(self, "node_ip_type")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Sequence[str]]:
        """
        Task Nodes.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def operate(self) -> pulumi.Output[Optional[str]]:
        """
        The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        """
        return pulumi.get(self, "operate")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[str]:
        """
        tasks parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tag description list.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskCategory")
    def task_category(self) -> pulumi.Output[int]:
        """
        Task category,1:PC,2:Mobile.
        """
        return pulumi.get(self, "task_category")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Output[str]:
        """
        Task Id.
        """
        return pulumi.get(self, "task_id")

    @property
    @pulumi.getter(name="taskType")
    def task_type(self) -> pulumi.Output[int]:
        """
        Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        """
        return pulumi.get(self, "task_type")

