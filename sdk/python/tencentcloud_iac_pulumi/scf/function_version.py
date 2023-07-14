# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FunctionVersionArgs', 'FunctionVersion']

@pulumi.input_type
class FunctionVersionArgs:
    def __init__(__self__, *,
                 function_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FunctionVersion resource.
        :param pulumi.Input[str] function_name: Name of the released function.
        :param pulumi.Input[str] description: Function description.
        :param pulumi.Input[str] namespace: Function namespace.
        """
        pulumi.set(__self__, "function_name", function_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        """
        Name of the released function.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Function description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _FunctionVersionState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 function_version: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FunctionVersion resources.
        :param pulumi.Input[str] description: Function description.
        :param pulumi.Input[str] function_name: Name of the released function.
        :param pulumi.Input[str] function_version: Version of the released function.
        :param pulumi.Input[str] namespace: Function namespace.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if function_name is not None:
            pulumi.set(__self__, "function_name", function_name)
        if function_version is not None:
            pulumi.set(__self__, "function_version", function_version)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Function description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the released function.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the released function.
        """
        return pulumi.get(self, "function_version")

    @function_version.setter
    def function_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_version", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


class FunctionVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a scf function_version

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        function_version = tencentcloud.scf.FunctionVersion("functionVersion",
            description="for-terraform-test",
            function_name="keep-1676351130",
            namespace="default")
        ```

        ## Import

        scf function_version can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Scf/functionVersion:FunctionVersion function_version functionName#namespace#functionVersion
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Function description.
        :param pulumi.Input[str] function_name: Name of the released function.
        :param pulumi.Input[str] namespace: Function namespace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a scf function_version

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        function_version = tencentcloud.scf.FunctionVersion("functionVersion",
            description="for-terraform-test",
            function_name="keep-1676351130",
            namespace="default")
        ```

        ## Import

        scf function_version can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Scf/functionVersion:FunctionVersion function_version functionName#namespace#functionVersion
        ```

        :param str resource_name: The name of the resource.
        :param FunctionVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
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
            __props__ = FunctionVersionArgs.__new__(FunctionVersionArgs)

            __props__.__dict__["description"] = description
            if function_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_name'")
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["function_version"] = None
        super(FunctionVersion, __self__).__init__(
            'tencentcloud:Scf/functionVersion:FunctionVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            function_version: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None) -> 'FunctionVersion':
        """
        Get an existing FunctionVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Function description.
        :param pulumi.Input[str] function_name: Name of the released function.
        :param pulumi.Input[str] function_version: Version of the released function.
        :param pulumi.Input[str] namespace: Function namespace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionVersionState.__new__(_FunctionVersionState)

        __props__.__dict__["description"] = description
        __props__.__dict__["function_name"] = function_name
        __props__.__dict__["function_version"] = function_version
        __props__.__dict__["namespace"] = namespace
        return FunctionVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Function description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        Name of the released function.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="functionVersion")
    def function_version(self) -> pulumi.Output[str]:
        """
        Version of the released function.
        """
        return pulumi.get(self, "function_version")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Function namespace.
        """
        return pulumi.get(self, "namespace")

