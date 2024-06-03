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

__all__ = ['ContentReviewTemplateArgs', 'ContentReviewTemplate']

@pulumi.input_type
class ContentReviewTemplateArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 political_configure: Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']] = None,
                 porn_configure: Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']] = None,
                 prohibited_configure: Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']] = None,
                 terrorism_configure: Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']] = None,
                 user_define_configure: Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']] = None):
        """
        The set of arguments for constructing a ContentReviewTemplate resource.
        :param pulumi.Input[str] comment: Content review template description information, length limit: 256 characters.
        :param pulumi.Input[str] name: Content review template name, length limit: 64 characters.
        :param pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs'] political_configure: Political control parameters.
        :param pulumi.Input['ContentReviewTemplatePornConfigureArgs'] porn_configure: Control parameters for porn image.
        :param pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs'] prohibited_configure: Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        :param pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs'] terrorism_configure: Control parameters for unsafe information.
        :param pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs'] user_define_configure: User-Defined Content Moderation Control Parameters.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if political_configure is not None:
            pulumi.set(__self__, "political_configure", political_configure)
        if porn_configure is not None:
            pulumi.set(__self__, "porn_configure", porn_configure)
        if prohibited_configure is not None:
            pulumi.set(__self__, "prohibited_configure", prohibited_configure)
        if terrorism_configure is not None:
            pulumi.set(__self__, "terrorism_configure", terrorism_configure)
        if user_define_configure is not None:
            pulumi.set(__self__, "user_define_configure", user_define_configure)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Content review template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Content review template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="politicalConfigure")
    def political_configure(self) -> Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']]:
        """
        Political control parameters.
        """
        return pulumi.get(self, "political_configure")

    @political_configure.setter
    def political_configure(self, value: Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']]):
        pulumi.set(self, "political_configure", value)

    @property
    @pulumi.getter(name="pornConfigure")
    def porn_configure(self) -> Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']]:
        """
        Control parameters for porn image.
        """
        return pulumi.get(self, "porn_configure")

    @porn_configure.setter
    def porn_configure(self, value: Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']]):
        pulumi.set(self, "porn_configure", value)

    @property
    @pulumi.getter(name="prohibitedConfigure")
    def prohibited_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']]:
        """
        Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        """
        return pulumi.get(self, "prohibited_configure")

    @prohibited_configure.setter
    def prohibited_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']]):
        pulumi.set(self, "prohibited_configure", value)

    @property
    @pulumi.getter(name="terrorismConfigure")
    def terrorism_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']]:
        """
        Control parameters for unsafe information.
        """
        return pulumi.get(self, "terrorism_configure")

    @terrorism_configure.setter
    def terrorism_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']]):
        pulumi.set(self, "terrorism_configure", value)

    @property
    @pulumi.getter(name="userDefineConfigure")
    def user_define_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']]:
        """
        User-Defined Content Moderation Control Parameters.
        """
        return pulumi.get(self, "user_define_configure")

    @user_define_configure.setter
    def user_define_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']]):
        pulumi.set(self, "user_define_configure", value)


@pulumi.input_type
class _ContentReviewTemplateState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 political_configure: Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']] = None,
                 porn_configure: Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']] = None,
                 prohibited_configure: Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']] = None,
                 terrorism_configure: Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']] = None,
                 user_define_configure: Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']] = None):
        """
        Input properties used for looking up and filtering ContentReviewTemplate resources.
        :param pulumi.Input[str] comment: Content review template description information, length limit: 256 characters.
        :param pulumi.Input[str] name: Content review template name, length limit: 64 characters.
        :param pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs'] political_configure: Political control parameters.
        :param pulumi.Input['ContentReviewTemplatePornConfigureArgs'] porn_configure: Control parameters for porn image.
        :param pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs'] prohibited_configure: Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        :param pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs'] terrorism_configure: Control parameters for unsafe information.
        :param pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs'] user_define_configure: User-Defined Content Moderation Control Parameters.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if political_configure is not None:
            pulumi.set(__self__, "political_configure", political_configure)
        if porn_configure is not None:
            pulumi.set(__self__, "porn_configure", porn_configure)
        if prohibited_configure is not None:
            pulumi.set(__self__, "prohibited_configure", prohibited_configure)
        if terrorism_configure is not None:
            pulumi.set(__self__, "terrorism_configure", terrorism_configure)
        if user_define_configure is not None:
            pulumi.set(__self__, "user_define_configure", user_define_configure)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Content review template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Content review template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="politicalConfigure")
    def political_configure(self) -> Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']]:
        """
        Political control parameters.
        """
        return pulumi.get(self, "political_configure")

    @political_configure.setter
    def political_configure(self, value: Optional[pulumi.Input['ContentReviewTemplatePoliticalConfigureArgs']]):
        pulumi.set(self, "political_configure", value)

    @property
    @pulumi.getter(name="pornConfigure")
    def porn_configure(self) -> Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']]:
        """
        Control parameters for porn image.
        """
        return pulumi.get(self, "porn_configure")

    @porn_configure.setter
    def porn_configure(self, value: Optional[pulumi.Input['ContentReviewTemplatePornConfigureArgs']]):
        pulumi.set(self, "porn_configure", value)

    @property
    @pulumi.getter(name="prohibitedConfigure")
    def prohibited_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']]:
        """
        Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        """
        return pulumi.get(self, "prohibited_configure")

    @prohibited_configure.setter
    def prohibited_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateProhibitedConfigureArgs']]):
        pulumi.set(self, "prohibited_configure", value)

    @property
    @pulumi.getter(name="terrorismConfigure")
    def terrorism_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']]:
        """
        Control parameters for unsafe information.
        """
        return pulumi.get(self, "terrorism_configure")

    @terrorism_configure.setter
    def terrorism_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateTerrorismConfigureArgs']]):
        pulumi.set(self, "terrorism_configure", value)

    @property
    @pulumi.getter(name="userDefineConfigure")
    def user_define_configure(self) -> Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']]:
        """
        User-Defined Content Moderation Control Parameters.
        """
        return pulumi.get(self, "user_define_configure")

    @user_define_configure.setter
    def user_define_configure(self, value: Optional[pulumi.Input['ContentReviewTemplateUserDefineConfigureArgs']]):
        pulumi.set(self, "user_define_configure", value)


class ContentReviewTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 political_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePoliticalConfigureArgs']]] = None,
                 porn_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePornConfigureArgs']]] = None,
                 prohibited_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateProhibitedConfigureArgs']]] = None,
                 terrorism_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateTerrorismConfigureArgs']]] = None,
                 user_define_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateUserDefineConfigureArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a mps content_review_template

        ## Import

        mps content_review_template can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate content_review_template definition
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Content review template description information, length limit: 256 characters.
        :param pulumi.Input[str] name: Content review template name, length limit: 64 characters.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplatePoliticalConfigureArgs']] political_configure: Political control parameters.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplatePornConfigureArgs']] porn_configure: Control parameters for porn image.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateProhibitedConfigureArgs']] prohibited_configure: Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateTerrorismConfigureArgs']] terrorism_configure: Control parameters for unsafe information.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateUserDefineConfigureArgs']] user_define_configure: User-Defined Content Moderation Control Parameters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContentReviewTemplateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a mps content_review_template

        ## Import

        mps content_review_template can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate content_review_template definition
        ```

        :param str resource_name: The name of the resource.
        :param ContentReviewTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContentReviewTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 political_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePoliticalConfigureArgs']]] = None,
                 porn_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePornConfigureArgs']]] = None,
                 prohibited_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateProhibitedConfigureArgs']]] = None,
                 terrorism_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateTerrorismConfigureArgs']]] = None,
                 user_define_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateUserDefineConfigureArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContentReviewTemplateArgs.__new__(ContentReviewTemplateArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["name"] = name
            __props__.__dict__["political_configure"] = political_configure
            __props__.__dict__["porn_configure"] = porn_configure
            __props__.__dict__["prohibited_configure"] = prohibited_configure
            __props__.__dict__["terrorism_configure"] = terrorism_configure
            __props__.__dict__["user_define_configure"] = user_define_configure
        super(ContentReviewTemplate, __self__).__init__(
            'tencentcloud:Mps/contentReviewTemplate:ContentReviewTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            political_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePoliticalConfigureArgs']]] = None,
            porn_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplatePornConfigureArgs']]] = None,
            prohibited_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateProhibitedConfigureArgs']]] = None,
            terrorism_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateTerrorismConfigureArgs']]] = None,
            user_define_configure: Optional[pulumi.Input[pulumi.InputType['ContentReviewTemplateUserDefineConfigureArgs']]] = None) -> 'ContentReviewTemplate':
        """
        Get an existing ContentReviewTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Content review template description information, length limit: 256 characters.
        :param pulumi.Input[str] name: Content review template name, length limit: 64 characters.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplatePoliticalConfigureArgs']] political_configure: Political control parameters.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplatePornConfigureArgs']] porn_configure: Control parameters for porn image.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateProhibitedConfigureArgs']] prohibited_configure: Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateTerrorismConfigureArgs']] terrorism_configure: Control parameters for unsafe information.
        :param pulumi.Input[pulumi.InputType['ContentReviewTemplateUserDefineConfigureArgs']] user_define_configure: User-Defined Content Moderation Control Parameters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContentReviewTemplateState.__new__(_ContentReviewTemplateState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["name"] = name
        __props__.__dict__["political_configure"] = political_configure
        __props__.__dict__["porn_configure"] = porn_configure
        __props__.__dict__["prohibited_configure"] = prohibited_configure
        __props__.__dict__["terrorism_configure"] = terrorism_configure
        __props__.__dict__["user_define_configure"] = user_define_configure
        return ContentReviewTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Content review template description information, length limit: 256 characters.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Content review template name, length limit: 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="politicalConfigure")
    def political_configure(self) -> pulumi.Output[Optional['outputs.ContentReviewTemplatePoliticalConfigure']]:
        """
        Political control parameters.
        """
        return pulumi.get(self, "political_configure")

    @property
    @pulumi.getter(name="pornConfigure")
    def porn_configure(self) -> pulumi.Output[Optional['outputs.ContentReviewTemplatePornConfigure']]:
        """
        Control parameters for porn image.
        """
        return pulumi.get(self, "porn_configure")

    @property
    @pulumi.getter(name="prohibitedConfigure")
    def prohibited_configure(self) -> pulumi.Output[Optional['outputs.ContentReviewTemplateProhibitedConfigure']]:
        """
        Prohibited control parameters. Prohibited content includes:abuse, drug-related violations.Note: this parameter is not yet supported.
        """
        return pulumi.get(self, "prohibited_configure")

    @property
    @pulumi.getter(name="terrorismConfigure")
    def terrorism_configure(self) -> pulumi.Output[Optional['outputs.ContentReviewTemplateTerrorismConfigure']]:
        """
        Control parameters for unsafe information.
        """
        return pulumi.get(self, "terrorism_configure")

    @property
    @pulumi.getter(name="userDefineConfigure")
    def user_define_configure(self) -> pulumi.Output[Optional['outputs.ContentReviewTemplateUserDefineConfigure']]:
        """
        User-Defined Content Moderation Control Parameters.
        """
        return pulumi.get(self, "user_define_configure")

