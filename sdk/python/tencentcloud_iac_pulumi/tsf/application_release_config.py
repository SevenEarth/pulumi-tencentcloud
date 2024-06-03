# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApplicationReleaseConfigArgs', 'ApplicationReleaseConfig']

@pulumi.input_type
class ApplicationReleaseConfigArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 release_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationReleaseConfig resource.
        :param pulumi.Input[str] config_id: Configuration ID.
        :param pulumi.Input[str] group_id: deployment group ID.
        :param pulumi.Input[str] release_desc: release description.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "group_id", group_id)
        if release_desc is not None:
            pulumi.set(__self__, "release_desc", release_desc)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        """
        Configuration ID.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        deployment group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> Optional[pulumi.Input[str]]:
        """
        release description.
        """
        return pulumi.get(self, "release_desc")

    @release_desc.setter
    def release_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_desc", value)


@pulumi.input_type
class _ApplicationReleaseConfigState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 config_name: Optional[pulumi.Input[str]] = None,
                 config_release_id: Optional[pulumi.Input[str]] = None,
                 config_version: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None,
                 release_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationReleaseConfig resources.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] cluster_id: cluster ID.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[str] config_id: Configuration ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_release_id: configuration item release ID.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] group_id: deployment group ID.
        :param pulumi.Input[str] group_name: deployment group name.
        :param pulumi.Input[str] namespace_id: Namespace ID.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] release_desc: release description.
        :param pulumi.Input[str] release_time: release time.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if config_name is not None:
            pulumi.set(__self__, "config_name", config_name)
        if config_release_id is not None:
            pulumi.set(__self__, "config_release_id", config_release_id)
        if config_version is not None:
            pulumi.set(__self__, "config_version", config_version)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if release_desc is not None:
            pulumi.set(__self__, "release_desc", release_desc)
        if release_time is not None:
            pulumi.set(__self__, "release_time", release_time)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application ID.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        Configuration ID.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item name.
        """
        return pulumi.get(self, "config_name")

    @config_name.setter
    def config_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_name", value)

    @property
    @pulumi.getter(name="configReleaseId")
    def config_release_id(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item release ID.
        """
        return pulumi.get(self, "config_release_id")

    @config_release_id.setter
    def config_release_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_release_id", value)

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> Optional[pulumi.Input[str]]:
        """
        configuration item version.
        """
        return pulumi.get(self, "config_version")

    @config_version.setter
    def config_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_version", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        deployment group ID.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[str]]:
        """
        deployment group name.
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        Namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        namespace name.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> Optional[pulumi.Input[str]]:
        """
        release description.
        """
        return pulumi.get(self, "release_desc")

    @release_desc.setter
    def release_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_desc", value)

    @property
    @pulumi.getter(name="releaseTime")
    def release_time(self) -> Optional[pulumi.Input[str]]:
        """
        release time.
        """
        return pulumi.get(self, "release_time")

    @release_time.setter
    def release_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "release_time", value)


class ApplicationReleaseConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a tsf application_release_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_release_config = tencentcloud.tsf.ApplicationReleaseConfig("applicationReleaseConfig",
            config_id="dcfg-nalqbqwv",
            group_id="group-yxmz72gv",
            release_desc="terraform-test")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf application_release_config can be imported using the configId#groupId#configReleaseId, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/applicationReleaseConfig:ApplicationReleaseConfig application_release_config dcfg-nalqbqwv#group-yxmz72gv#dcfgr-maeeq2ea
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: Configuration ID.
        :param pulumi.Input[str] group_id: deployment group ID.
        :param pulumi.Input[str] release_desc: release description.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationReleaseConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a tsf application_release_config

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        application_release_config = tencentcloud.tsf.ApplicationReleaseConfig("applicationReleaseConfig",
            config_id="dcfg-nalqbqwv",
            group_id="group-yxmz72gv",
            release_desc="terraform-test")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        tsf application_release_config can be imported using the configId#groupId#configReleaseId, e.g.

        ```sh
        $ pulumi import tencentcloud:Tsf/applicationReleaseConfig:ApplicationReleaseConfig application_release_config dcfg-nalqbqwv#group-yxmz72gv#dcfgr-maeeq2ea
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationReleaseConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationReleaseConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 release_desc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationReleaseConfigArgs.__new__(ApplicationReleaseConfigArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["release_desc"] = release_desc
            __props__.__dict__["application_id"] = None
            __props__.__dict__["cluster_id"] = None
            __props__.__dict__["cluster_name"] = None
            __props__.__dict__["config_name"] = None
            __props__.__dict__["config_release_id"] = None
            __props__.__dict__["config_version"] = None
            __props__.__dict__["group_name"] = None
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_name"] = None
            __props__.__dict__["release_time"] = None
        super(ApplicationReleaseConfig, __self__).__init__(
            'tencentcloud:Tsf/applicationReleaseConfig:ApplicationReleaseConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            config_name: Optional[pulumi.Input[str]] = None,
            config_release_id: Optional[pulumi.Input[str]] = None,
            config_version: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            release_desc: Optional[pulumi.Input[str]] = None,
            release_time: Optional[pulumi.Input[str]] = None) -> 'ApplicationReleaseConfig':
        """
        Get an existing ApplicationReleaseConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: Application ID.
        :param pulumi.Input[str] cluster_id: cluster ID.
        :param pulumi.Input[str] cluster_name: cluster name.
        :param pulumi.Input[str] config_id: Configuration ID.
        :param pulumi.Input[str] config_name: configuration item name.
        :param pulumi.Input[str] config_release_id: configuration item release ID.
        :param pulumi.Input[str] config_version: configuration item version.
        :param pulumi.Input[str] group_id: deployment group ID.
        :param pulumi.Input[str] group_name: deployment group name.
        :param pulumi.Input[str] namespace_id: Namespace ID.
        :param pulumi.Input[str] namespace_name: namespace name.
        :param pulumi.Input[str] release_desc: release description.
        :param pulumi.Input[str] release_time: release time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationReleaseConfigState.__new__(_ApplicationReleaseConfigState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["config_name"] = config_name
        __props__.__dict__["config_release_id"] = config_release_id
        __props__.__dict__["config_version"] = config_version
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["release_desc"] = release_desc
        __props__.__dict__["release_time"] = release_time
        return ApplicationReleaseConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        Application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        Configuration ID.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> pulumi.Output[str]:
        """
        configuration item name.
        """
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter(name="configReleaseId")
    def config_release_id(self) -> pulumi.Output[str]:
        """
        configuration item release ID.
        """
        return pulumi.get(self, "config_release_id")

    @property
    @pulumi.getter(name="configVersion")
    def config_version(self) -> pulumi.Output[str]:
        """
        configuration item version.
        """
        return pulumi.get(self, "config_version")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        deployment group ID.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        deployment group name.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        Namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        namespace name.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="releaseDesc")
    def release_desc(self) -> pulumi.Output[str]:
        """
        release description.
        """
        return pulumi.get(self, "release_desc")

    @property
    @pulumi.getter(name="releaseTime")
    def release_time(self) -> pulumi.Output[str]:
        """
        release time.
        """
        return pulumi.get(self, "release_time")

