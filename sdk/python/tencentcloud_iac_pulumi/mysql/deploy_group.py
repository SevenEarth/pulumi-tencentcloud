# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeployGroupArgs', 'DeployGroup']

@pulumi.input_type
class DeployGroupArgs:
    def __init__(__self__, *,
                 deploy_group_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 dev_classes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 limit_num: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a DeployGroup resource.
        :param pulumi.Input[str] deploy_group_name: The name of deploy group. the maximum length cannot exceed 60 characters.
        :param pulumi.Input[str] description: The description of deploy group. the maximum length cannot exceed 200 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dev_classes: The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        :param pulumi.Input[int] limit_num: The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        pulumi.set(__self__, "deploy_group_name", deploy_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dev_classes is not None:
            pulumi.set(__self__, "dev_classes", dev_classes)
        if limit_num is not None:
            pulumi.set(__self__, "limit_num", limit_num)

    @property
    @pulumi.getter(name="deployGroupName")
    def deploy_group_name(self) -> pulumi.Input[str]:
        """
        The name of deploy group. the maximum length cannot exceed 60 characters.
        """
        return pulumi.get(self, "deploy_group_name")

    @deploy_group_name.setter
    def deploy_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "deploy_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of deploy group. the maximum length cannot exceed 200 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="devClasses")
    def dev_classes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        """
        return pulumi.get(self, "dev_classes")

    @dev_classes.setter
    def dev_classes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dev_classes", value)

    @property
    @pulumi.getter(name="limitNum")
    def limit_num(self) -> Optional[pulumi.Input[int]]:
        """
        The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        return pulumi.get(self, "limit_num")

    @limit_num.setter
    def limit_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "limit_num", value)


@pulumi.input_type
class _DeployGroupState:
    def __init__(__self__, *,
                 deploy_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dev_classes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 limit_num: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering DeployGroup resources.
        :param pulumi.Input[str] deploy_group_name: The name of deploy group. the maximum length cannot exceed 60 characters.
        :param pulumi.Input[str] description: The description of deploy group. the maximum length cannot exceed 200 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dev_classes: The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        :param pulumi.Input[int] limit_num: The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        if deploy_group_name is not None:
            pulumi.set(__self__, "deploy_group_name", deploy_group_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dev_classes is not None:
            pulumi.set(__self__, "dev_classes", dev_classes)
        if limit_num is not None:
            pulumi.set(__self__, "limit_num", limit_num)

    @property
    @pulumi.getter(name="deployGroupName")
    def deploy_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of deploy group. the maximum length cannot exceed 60 characters.
        """
        return pulumi.get(self, "deploy_group_name")

    @deploy_group_name.setter
    def deploy_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_group_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of deploy group. the maximum length cannot exceed 200 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="devClasses")
    def dev_classes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        """
        return pulumi.get(self, "dev_classes")

    @dev_classes.setter
    def dev_classes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dev_classes", value)

    @property
    @pulumi.getter(name="limitNum")
    def limit_num(self) -> Optional[pulumi.Input[int]]:
        """
        The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        return pulumi.get(self, "limit_num")

    @limit_num.setter
    def limit_num(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "limit_num", value)


class DeployGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dev_classes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 limit_num: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a mysql deploy_group

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        deploy_group = tencentcloud.mysql.DeployGroup("deployGroup",
            deploy_group_name="terrform-deploy",
            description="deploy test",
            dev_classes=["TS85"],
            limit_num=1)
        ```

        ## Import

        mysql deploy_group can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mysql/deployGroup:DeployGroup deploy_group deploy_group_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deploy_group_name: The name of deploy group. the maximum length cannot exceed 60 characters.
        :param pulumi.Input[str] description: The description of deploy group. the maximum length cannot exceed 200 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dev_classes: The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        :param pulumi.Input[int] limit_num: The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeployGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mysql deploy_group

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        deploy_group = tencentcloud.mysql.DeployGroup("deployGroup",
            deploy_group_name="terrform-deploy",
            description="deploy test",
            dev_classes=["TS85"],
            limit_num=1)
        ```

        ## Import

        mysql deploy_group can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Mysql/deployGroup:DeployGroup deploy_group deploy_group_id
        ```

        :param str resource_name: The name of the resource.
        :param DeployGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeployGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_group_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dev_classes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 limit_num: Optional[pulumi.Input[int]] = None,
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
            __props__ = DeployGroupArgs.__new__(DeployGroupArgs)

            if deploy_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'deploy_group_name'")
            __props__.__dict__["deploy_group_name"] = deploy_group_name
            __props__.__dict__["description"] = description
            __props__.__dict__["dev_classes"] = dev_classes
            __props__.__dict__["limit_num"] = limit_num
        super(DeployGroup, __self__).__init__(
            'tencentcloud:Mysql/deployGroup:DeployGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deploy_group_name: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dev_classes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            limit_num: Optional[pulumi.Input[int]] = None) -> 'DeployGroup':
        """
        Get an existing DeployGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deploy_group_name: The name of deploy group. the maximum length cannot exceed 60 characters.
        :param pulumi.Input[str] description: The description of deploy group. the maximum length cannot exceed 200 characters.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dev_classes: The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        :param pulumi.Input[int] limit_num: The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeployGroupState.__new__(_DeployGroupState)

        __props__.__dict__["deploy_group_name"] = deploy_group_name
        __props__.__dict__["description"] = description
        __props__.__dict__["dev_classes"] = dev_classes
        __props__.__dict__["limit_num"] = limit_num
        return DeployGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deployGroupName")
    def deploy_group_name(self) -> pulumi.Output[str]:
        """
        The name of deploy group. the maximum length cannot exceed 60 characters.
        """
        return pulumi.get(self, "deploy_group_name")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of deploy group. the maximum length cannot exceed 200 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="devClasses")
    def dev_classes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The device class of deploy group. optional value is SH12+SH02, TS85, etc.
        """
        return pulumi.get(self, "dev_classes")

    @property
    @pulumi.getter(name="limitNum")
    def limit_num(self) -> pulumi.Output[Optional[int]]:
        """
        The limit on the number of instances on the same physical machine in deploy group affinity policy 1.
        """
        return pulumi.get(self, "limit_num")

