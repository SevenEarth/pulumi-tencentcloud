# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 namespace_name: pulumi.Input[str],
                 brief_desc: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] namespace_name: Name of the TCR namespace.
        :param pulumi.Input[str] brief_desc: Brief description of the repository. Valid length is [1~100].
        :param pulumi.Input[str] description: Description of the repository. Valid length is [1~1000].
        :param pulumi.Input[str] name: Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "namespace_name", namespace_name)
        if brief_desc is not None:
            pulumi.set(__self__, "brief_desc", brief_desc)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Input[str]:
        """
        Name of the TCR namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="briefDesc")
    def brief_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Brief description of the repository. Valid length is [1~100].
        """
        return pulumi.get(self, "brief_desc")

    @brief_desc.setter
    def brief_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "brief_desc", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the repository. Valid length is [1~1000].
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 brief_desc: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] brief_desc: Brief description of the repository. Valid length is [1~100].
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description of the repository. Valid length is [1~1000].
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[bool] is_public: Indicate the repository is public or not.
        :param pulumi.Input[str] name: Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        :param pulumi.Input[str] namespace_name: Name of the TCR namespace.
        :param pulumi.Input[str] update_time: Last updated time.
        :param pulumi.Input[str] url: URL of the repository.
        """
        if brief_desc is not None:
            pulumi.set(__self__, "brief_desc", brief_desc)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace_name is not None:
            pulumi.set(__self__, "namespace_name", namespace_name)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="briefDesc")
    def brief_desc(self) -> Optional[pulumi.Input[str]]:
        """
        Brief description of the repository. Valid length is [1~100].
        """
        return pulumi.get(self, "brief_desc")

    @brief_desc.setter
    def brief_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "brief_desc", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the repository. Valid length is [1~1000].
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate the repository is public or not.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the TCR namespace.
        """
        return pulumi.get(self, "namespace_name")

    @namespace_name.setter
    def namespace_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_name", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Last updated time.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the repository.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 brief_desc: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create tcr repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        test = tencentcloud.Tcr.get_instances(name="test")
        foo = tencentcloud.tcr.Repository("foo",
            instance_id=test.instance_lists[0].id,
            namespace_name="exampleNamespace")
        ```

        ## Import

        tcr repository can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/repository:Repository foo cls-cda1iex1#namespace#repository
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] brief_desc: Brief description of the repository. Valid length is [1~100].
        :param pulumi.Input[str] description: Description of the repository. Valid length is [1~1000].
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[str] name: Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        :param pulumi.Input[str] namespace_name: Name of the TCR namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create tcr repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        test = tencentcloud.Tcr.get_instances(name="test")
        foo = tencentcloud.tcr.Repository("foo",
            instance_id=test.instance_lists[0].id,
            namespace_name="exampleNamespace")
        ```

        ## Import

        tcr repository can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Tcr/repository:Repository foo cls-cda1iex1#namespace#repository
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 brief_desc: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["brief_desc"] = brief_desc
            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            if namespace_name is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_name'")
            __props__.__dict__["namespace_name"] = namespace_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["is_public"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["url"] = None
        super(Repository, __self__).__init__(
            'tencentcloud:Tcr/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            brief_desc: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_name: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] brief_desc: Brief description of the repository. Valid length is [1~100].
        :param pulumi.Input[str] create_time: Create time.
        :param pulumi.Input[str] description: Description of the repository. Valid length is [1~1000].
        :param pulumi.Input[str] instance_id: ID of the TCR instance.
        :param pulumi.Input[bool] is_public: Indicate the repository is public or not.
        :param pulumi.Input[str] name: Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        :param pulumi.Input[str] namespace_name: Name of the TCR namespace.
        :param pulumi.Input[str] update_time: Last updated time.
        :param pulumi.Input[str] url: URL of the repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["brief_desc"] = brief_desc
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_public"] = is_public
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace_name"] = namespace_name
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["url"] = url
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="briefDesc")
    def brief_desc(self) -> pulumi.Output[Optional[str]]:
        """
        Brief description of the repository. Valid length is [1~100].
        """
        return pulumi.get(self, "brief_desc")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the repository. Valid length is [1~1000].
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the TCR instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[bool]:
        """
        Indicate the repository is public or not.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceName")
    def namespace_name(self) -> pulumi.Output[str]:
        """
        Name of the TCR namespace.
        """
        return pulumi.get(self, "namespace_name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last updated time.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the repository.
        """
        return pulumi.get(self, "url")

