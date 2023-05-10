# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EmailAddressArgs', 'EmailAddress']

@pulumi.input_type
class EmailAddressArgs:
    def __init__(__self__, *,
                 email_address: pulumi.Input[str],
                 email_sender_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EmailAddress resource.
        :param pulumi.Input[str] email_address: Your sender address. (You can create up to 10 sender addresses for each domain.).
        :param pulumi.Input[str] email_sender_name: Sender name.
        """
        pulumi.set(__self__, "email_address", email_address)
        if email_sender_name is not None:
            pulumi.set(__self__, "email_sender_name", email_sender_name)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> pulumi.Input[str]:
        """
        Your sender address. (You can create up to 10 sender addresses for each domain.).
        """
        return pulumi.get(self, "email_address")

    @email_address.setter
    def email_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "email_address", value)

    @property
    @pulumi.getter(name="emailSenderName")
    def email_sender_name(self) -> Optional[pulumi.Input[str]]:
        """
        Sender name.
        """
        return pulumi.get(self, "email_sender_name")

    @email_sender_name.setter
    def email_sender_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_sender_name", value)


@pulumi.input_type
class _EmailAddressState:
    def __init__(__self__, *,
                 email_address: Optional[pulumi.Input[str]] = None,
                 email_sender_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EmailAddress resources.
        :param pulumi.Input[str] email_address: Your sender address. (You can create up to 10 sender addresses for each domain.).
        :param pulumi.Input[str] email_sender_name: Sender name.
        """
        if email_address is not None:
            pulumi.set(__self__, "email_address", email_address)
        if email_sender_name is not None:
            pulumi.set(__self__, "email_sender_name", email_sender_name)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> Optional[pulumi.Input[str]]:
        """
        Your sender address. (You can create up to 10 sender addresses for each domain.).
        """
        return pulumi.get(self, "email_address")

    @email_address.setter
    def email_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_address", value)

    @property
    @pulumi.getter(name="emailSenderName")
    def email_sender_name(self) -> Optional[pulumi.Input[str]]:
        """
        Sender name.
        """
        return pulumi.get(self, "email_sender_name")

    @email_sender_name.setter
    def email_sender_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_sender_name", value)


class EmailAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email_address: Optional[pulumi.Input[str]] = None,
                 email_sender_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a ses email_address

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        email_address = tencentcloud.ses.EmailAddress("emailAddress",
            email_address="aaa@iac-tf.cloud",
            email_sender_name="aaa")
        ```

        ## Import

        ses email_address can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ses/emailAddress:EmailAddress email_address aaa@iac-tf.cloud
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email_address: Your sender address. (You can create up to 10 sender addresses for each domain.).
        :param pulumi.Input[str] email_sender_name: Sender name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EmailAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ses email_address

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        email_address = tencentcloud.ses.EmailAddress("emailAddress",
            email_address="aaa@iac-tf.cloud",
            email_sender_name="aaa")
        ```

        ## Import

        ses email_address can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Ses/emailAddress:EmailAddress email_address aaa@iac-tf.cloud
        ```

        :param str resource_name: The name of the resource.
        :param EmailAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmailAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email_address: Optional[pulumi.Input[str]] = None,
                 email_sender_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = EmailAddressArgs.__new__(EmailAddressArgs)

            if email_address is None and not opts.urn:
                raise TypeError("Missing required property 'email_address'")
            __props__.__dict__["email_address"] = email_address
            __props__.__dict__["email_sender_name"] = email_sender_name
        super(EmailAddress, __self__).__init__(
            'tencentcloud:Ses/emailAddress:EmailAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email_address: Optional[pulumi.Input[str]] = None,
            email_sender_name: Optional[pulumi.Input[str]] = None) -> 'EmailAddress':
        """
        Get an existing EmailAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email_address: Your sender address. (You can create up to 10 sender addresses for each domain.).
        :param pulumi.Input[str] email_sender_name: Sender name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EmailAddressState.__new__(_EmailAddressState)

        __props__.__dict__["email_address"] = email_address
        __props__.__dict__["email_sender_name"] = email_sender_name
        return EmailAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> pulumi.Output[str]:
        """
        Your sender address. (You can create up to 10 sender addresses for each domain.).
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter(name="emailSenderName")
    def email_sender_name(self) -> pulumi.Output[Optional[str]]:
        """
        Sender name.
        """
        return pulumi.get(self, "email_sender_name")

