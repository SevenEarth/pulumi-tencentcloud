# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 enable_url_group: pulumi.Input[int],
                 instance_id: pulumi.Input[str],
                 rate: pulumi.Input[str],
                 type: pulumi.Input[str],
                 desc: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[int] enable_url_group: Whether to enable aggregation.
        :param pulumi.Input[str] instance_id: Business system ID.
        :param pulumi.Input[str] rate: Project sampling rate (greater than or equal to 0).
        :param pulumi.Input[str] type: Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        :param pulumi.Input[str] desc: Description of the created project (optional and up to 1,000 characters).
        :param pulumi.Input[str] name: Name of the created project (required and up to 200 characters).
        :param pulumi.Input[str] repo: Repository address of the project (optional and up to 256 characters).
        :param pulumi.Input[str] url: Webpage address of the project (optional and up to 256 characters).
        """
        pulumi.set(__self__, "enable_url_group", enable_url_group)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "rate", rate)
        pulumi.set(__self__, "type", type)
        if desc is not None:
            pulumi.set(__self__, "desc", desc)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="enableUrlGroup")
    def enable_url_group(self) -> pulumi.Input[int]:
        """
        Whether to enable aggregation.
        """
        return pulumi.get(self, "enable_url_group")

    @enable_url_group.setter
    def enable_url_group(self, value: pulumi.Input[int]):
        pulumi.set(self, "enable_url_group", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Business system ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Input[str]:
        """
        Project sampling rate (greater than or equal to 0).
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: pulumi.Input[str]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def desc(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the created project (optional and up to 1,000 characters).
        """
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the created project (required and up to 200 characters).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input[str]]:
        """
        Repository address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Webpage address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 creator: Optional[pulumi.Input[str]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 enable_url_group: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_key: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 is_star: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_status: Optional[pulumi.Input[int]] = None,
                 rate: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[str] create_time: Creata Time.
        :param pulumi.Input[str] creator: Creator ID.
        :param pulumi.Input[str] desc: Description of the created project (optional and up to 1,000 characters).
        :param pulumi.Input[int] enable_url_group: Whether to enable aggregation.
        :param pulumi.Input[str] instance_id: Business system ID.
        :param pulumi.Input[str] instance_key: Instance key.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[int] is_star: Starred status. `1`: yes; `0`: no.
        :param pulumi.Input[str] key: Unique project key (12 characters).
        :param pulumi.Input[str] name: Name of the created project (required and up to 200 characters).
        :param pulumi.Input[int] project_status: Project status (`1`: Creating; `2`: Running; `3`: Abnormal; `4`: Restarting; `5`: Stopping; `6`: Stopped; `7`: Terminating; `8`: Terminated).
        :param pulumi.Input[str] rate: Project sampling rate (greater than or equal to 0).
        :param pulumi.Input[str] repo: Repository address of the project (optional and up to 256 characters).
        :param pulumi.Input[str] type: Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        :param pulumi.Input[str] url: Webpage address of the project (optional and up to 256 characters).
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if creator is not None:
            pulumi.set(__self__, "creator", creator)
        if desc is not None:
            pulumi.set(__self__, "desc", desc)
        if enable_url_group is not None:
            pulumi.set(__self__, "enable_url_group", enable_url_group)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_key is not None:
            pulumi.set(__self__, "instance_key", instance_key)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if is_star is not None:
            pulumi.set(__self__, "is_star", is_star)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_status is not None:
            pulumi.set(__self__, "project_status", project_status)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creata Time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def creator(self) -> Optional[pulumi.Input[str]]:
        """
        Creator ID.
        """
        return pulumi.get(self, "creator")

    @creator.setter
    def creator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator", value)

    @property
    @pulumi.getter
    def desc(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the created project (optional and up to 1,000 characters).
        """
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter(name="enableUrlGroup")
    def enable_url_group(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to enable aggregation.
        """
        return pulumi.get(self, "enable_url_group")

    @enable_url_group.setter
    def enable_url_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "enable_url_group", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Business system ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceKey")
    def instance_key(self) -> Optional[pulumi.Input[str]]:
        """
        Instance key.
        """
        return pulumi.get(self, "instance_key")

    @instance_key.setter
    def instance_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_key", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Instance name.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="isStar")
    def is_star(self) -> Optional[pulumi.Input[int]]:
        """
        Starred status. `1`: yes; `0`: no.
        """
        return pulumi.get(self, "is_star")

    @is_star.setter
    def is_star(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "is_star", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Unique project key (12 characters).
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the created project (required and up to 200 characters).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectStatus")
    def project_status(self) -> Optional[pulumi.Input[int]]:
        """
        Project status (`1`: Creating; `2`: Running; `3`: Abnormal; `4`: Restarting; `5`: Stopping; `6`: Stopped; `7`: Terminating; `8`: Terminated).
        """
        return pulumi.get(self, "project_status")

    @project_status.setter
    def project_status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_status", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[str]]:
        """
        Project sampling rate (greater than or equal to 0).
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input[str]]:
        """
        Repository address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Webpage address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 enable_url_group: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a rum project

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        project = tencentcloud.rum.Project("project",
            desc="projectDesc-1",
            enable_url_group=0,
            instance_id="rum-pasZKEI3RLgakj",
            rate="100",
            repo="",
            type="web",
            url="iac-tf.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        rum project can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Rum/project:Project project project_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] desc: Description of the created project (optional and up to 1,000 characters).
        :param pulumi.Input[int] enable_url_group: Whether to enable aggregation.
        :param pulumi.Input[str] instance_id: Business system ID.
        :param pulumi.Input[str] name: Name of the created project (required and up to 200 characters).
        :param pulumi.Input[str] rate: Project sampling rate (greater than or equal to 0).
        :param pulumi.Input[str] repo: Repository address of the project (optional and up to 256 characters).
        :param pulumi.Input[str] type: Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        :param pulumi.Input[str] url: Webpage address of the project (optional and up to 256 characters).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a rum project

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        project = tencentcloud.rum.Project("project",
            desc="projectDesc-1",
            enable_url_group=0,
            instance_id="rum-pasZKEI3RLgakj",
            rate="100",
            repo="",
            type="web",
            url="iac-tf.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        rum project can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Rum/project:Project project project_id
        ```

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 enable_url_group: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[str]] = None,
                 repo: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["desc"] = desc
            if enable_url_group is None and not opts.urn:
                raise TypeError("Missing required property 'enable_url_group'")
            __props__.__dict__["enable_url_group"] = enable_url_group
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            if rate is None and not opts.urn:
                raise TypeError("Missing required property 'rate'")
            __props__.__dict__["rate"] = rate
            __props__.__dict__["repo"] = repo
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["url"] = url
            __props__.__dict__["create_time"] = None
            __props__.__dict__["creator"] = None
            __props__.__dict__["instance_key"] = None
            __props__.__dict__["instance_name"] = None
            __props__.__dict__["is_star"] = None
            __props__.__dict__["key"] = None
            __props__.__dict__["project_status"] = None
        super(Project, __self__).__init__(
            'tencentcloud:Rum/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            creator: Optional[pulumi.Input[str]] = None,
            desc: Optional[pulumi.Input[str]] = None,
            enable_url_group: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_key: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            is_star: Optional[pulumi.Input[int]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_status: Optional[pulumi.Input[int]] = None,
            rate: Optional[pulumi.Input[str]] = None,
            repo: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creata Time.
        :param pulumi.Input[str] creator: Creator ID.
        :param pulumi.Input[str] desc: Description of the created project (optional and up to 1,000 characters).
        :param pulumi.Input[int] enable_url_group: Whether to enable aggregation.
        :param pulumi.Input[str] instance_id: Business system ID.
        :param pulumi.Input[str] instance_key: Instance key.
        :param pulumi.Input[str] instance_name: Instance name.
        :param pulumi.Input[int] is_star: Starred status. `1`: yes; `0`: no.
        :param pulumi.Input[str] key: Unique project key (12 characters).
        :param pulumi.Input[str] name: Name of the created project (required and up to 200 characters).
        :param pulumi.Input[int] project_status: Project status (`1`: Creating; `2`: Running; `3`: Abnormal; `4`: Restarting; `5`: Stopping; `6`: Stopped; `7`: Terminating; `8`: Terminated).
        :param pulumi.Input[str] rate: Project sampling rate (greater than or equal to 0).
        :param pulumi.Input[str] repo: Repository address of the project (optional and up to 256 characters).
        :param pulumi.Input[str] type: Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        :param pulumi.Input[str] url: Webpage address of the project (optional and up to 256 characters).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["creator"] = creator
        __props__.__dict__["desc"] = desc
        __props__.__dict__["enable_url_group"] = enable_url_group
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_key"] = instance_key
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["is_star"] = is_star
        __props__.__dict__["key"] = key
        __props__.__dict__["name"] = name
        __props__.__dict__["project_status"] = project_status
        __props__.__dict__["rate"] = rate
        __props__.__dict__["repo"] = repo
        __props__.__dict__["type"] = type
        __props__.__dict__["url"] = url
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creata Time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def creator(self) -> pulumi.Output[str]:
        """
        Creator ID.
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter
    def desc(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the created project (optional and up to 1,000 characters).
        """
        return pulumi.get(self, "desc")

    @property
    @pulumi.getter(name="enableUrlGroup")
    def enable_url_group(self) -> pulumi.Output[int]:
        """
        Whether to enable aggregation.
        """
        return pulumi.get(self, "enable_url_group")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Business system ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceKey")
    def instance_key(self) -> pulumi.Output[str]:
        """
        Instance key.
        """
        return pulumi.get(self, "instance_key")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="isStar")
    def is_star(self) -> pulumi.Output[int]:
        """
        Starred status. `1`: yes; `0`: no.
        """
        return pulumi.get(self, "is_star")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Unique project key (12 characters).
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the created project (required and up to 200 characters).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectStatus")
    def project_status(self) -> pulumi.Output[int]:
        """
        Project status (`1`: Creating; `2`: Running; `3`: Abnormal; `4`: Restarting; `5`: Stopping; `6`: Stopped; `7`: Terminating; `8`: Terminated).
        """
        return pulumi.get(self, "project_status")

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Output[str]:
        """
        Project sampling rate (greater than or equal to 0).
        """
        return pulumi.get(self, "rate")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional[str]]:
        """
        Repository address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "repo")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Project type (valid values: `web`, `mp`, `android`, `ios`, `node`, `hippy`, `weex`, `viola`, `rn`).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        Webpage address of the project (optional and up to 256 characters).
        """
        return pulumi.get(self, "url")

