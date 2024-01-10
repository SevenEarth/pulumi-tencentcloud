# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WebShellArgs', 'WebShell']

@pulumi.input_type
class WebShellArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 status: pulumi.Input[int]):
        """
        The set of arguments for constructing a WebShell resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] status: Webshell status, 1: open; 0: closed; 2: log.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[int]:
        """
        Webshell status, 1: open; 0: closed; 2: log.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[int]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _WebShellState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering WebShell resources.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] status: Webshell status, 1: open; 0: closed; 2: log.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        Webshell status, 1: open; 0: closed; 2: log.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)


class WebShell(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a waf web_shell

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.waf.WebShell("example",
            domain="demo.waf.com",
            status=0)
        ```

        ## Import

        waf web_shell can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Waf/webShell:WebShell example demo.waf.com
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] status: Webshell status, 1: open; 0: closed; 2: log.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebShellArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a waf web_shell

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.waf.WebShell("example",
            domain="demo.waf.com",
            status=0)
        ```

        ## Import

        waf web_shell can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Waf/webShell:WebShell example demo.waf.com
        ```

        :param str resource_name: The name of the resource.
        :param WebShellArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebShellArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
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
            __props__ = WebShellArgs.__new__(WebShellArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
        super(WebShell, __self__).__init__(
            'tencentcloud:Waf/webShell:WebShell',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None) -> 'WebShell':
        """
        Get an existing WebShell resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[int] status: Webshell status, 1: open; 0: closed; 2: log.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebShellState.__new__(_WebShellState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["status"] = status
        return WebShell(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        Webshell status, 1: open; 0: closed; 2: log.
        """
        return pulumi.get(self, "status")

