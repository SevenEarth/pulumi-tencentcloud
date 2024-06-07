# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LiveTranscodeRuleAttachmentArgs', 'LiveTranscodeRuleAttachment']

@pulumi.input_type
class LiveTranscodeRuleAttachmentArgs:
    def __init__(__self__, *,
                 app_name: pulumi.Input[str],
                 domain_name: pulumi.Input[str],
                 stream_name: pulumi.Input[str],
                 template_id: pulumi.Input[int]):
        """
        The set of arguments for constructing a LiveTranscodeRuleAttachment resource.
        :param pulumi.Input[str] app_name: app name which you want to bind, can be empty string if not binding specific app name.
        :param pulumi.Input[str] domain_name: domain name hich you want to bind the transcode template.
        :param pulumi.Input[str] stream_name: stream name which you want to bind, can be empty string if not binding specific stream.
        :param pulumi.Input[int] template_id: template created by css_live_transcode_template.
        """
        pulumi.set(__self__, "app_name", app_name)
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "stream_name", stream_name)
        pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> pulumi.Input[str]:
        """
        app name which you want to bind, can be empty string if not binding specific app name.
        """
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        domain name hich you want to bind the transcode template.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Input[str]:
        """
        stream name which you want to bind, can be empty string if not binding specific stream.
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[int]:
        """
        template created by css_live_transcode_template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class _LiveTranscodeRuleAttachmentState:
    def __init__(__self__, *,
                 app_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LiveTranscodeRuleAttachment resources.
        :param pulumi.Input[str] app_name: app name which you want to bind, can be empty string if not binding specific app name.
        :param pulumi.Input[str] create_time: create time.
        :param pulumi.Input[str] domain_name: domain name hich you want to bind the transcode template.
        :param pulumi.Input[str] stream_name: stream name which you want to bind, can be empty string if not binding specific stream.
        :param pulumi.Input[int] template_id: template created by css_live_transcode_template.
        :param pulumi.Input[str] update_time: update time.
        """
        if app_name is not None:
            pulumi.set(__self__, "app_name", app_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if stream_name is not None:
            pulumi.set(__self__, "stream_name", stream_name)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> Optional[pulumi.Input[str]]:
        """
        app name which you want to bind, can be empty string if not binding specific app name.
        """
        return pulumi.get(self, "app_name")

    @app_name.setter
    def app_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        domain name hich you want to bind the transcode template.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> Optional[pulumi.Input[str]]:
        """
        stream name which you want to bind, can be empty string if not binding specific stream.
        """
        return pulumi.get(self, "stream_name")

    @stream_name.setter
    def stream_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_name", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[int]]:
        """
        template created by css_live_transcode_template.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        update time.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class LiveTranscodeRuleAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_name: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to create a css live_transcode_rule_attachment

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        task = tencentcloud.css.PullStreamTask("task",
            source_type="%s",
            source_urls=["%s"],
            domain_name="%s",
            app_name="%s",
            stream_name="%s",
            start_time="%s",
            end_time="%s",
            operator="%s",
            comment="This is a demo.")
        temp = tencentcloud.css.LiveTranscodeTemplate("temp",
            template_name="xxx",
            acodec="aac",
            video_bitrate=100,
            vcodec="origin",
            description="This_is_a_tf_test_temp.",
            need_video=1,
            need_audio=1)
        live_transcode_rule_attachment = tencentcloud.css.LiveTranscodeRuleAttachment("liveTranscodeRuleAttachment",
            domain_name=task.domain_name,
            app_name=task.app_name,
            stream_name=task.stream_name,
            template_id=temp.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        css live_transcode_rule_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Css/liveTranscodeRuleAttachment:LiveTranscodeRuleAttachment live_transcode_rule_attachment liveTranscodeRuleAttachment_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_name: app name which you want to bind, can be empty string if not binding specific app name.
        :param pulumi.Input[str] domain_name: domain name hich you want to bind the transcode template.
        :param pulumi.Input[str] stream_name: stream name which you want to bind, can be empty string if not binding specific stream.
        :param pulumi.Input[int] template_id: template created by css_live_transcode_template.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LiveTranscodeRuleAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a css live_transcode_rule_attachment

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        task = tencentcloud.css.PullStreamTask("task",
            source_type="%s",
            source_urls=["%s"],
            domain_name="%s",
            app_name="%s",
            stream_name="%s",
            start_time="%s",
            end_time="%s",
            operator="%s",
            comment="This is a demo.")
        temp = tencentcloud.css.LiveTranscodeTemplate("temp",
            template_name="xxx",
            acodec="aac",
            video_bitrate=100,
            vcodec="origin",
            description="This_is_a_tf_test_temp.",
            need_video=1,
            need_audio=1)
        live_transcode_rule_attachment = tencentcloud.css.LiveTranscodeRuleAttachment("liveTranscodeRuleAttachment",
            domain_name=task.domain_name,
            app_name=task.app_name,
            stream_name=task.stream_name,
            template_id=temp.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        css live_transcode_rule_attachment can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Css/liveTranscodeRuleAttachment:LiveTranscodeRuleAttachment live_transcode_rule_attachment liveTranscodeRuleAttachment_id
        ```

        :param str resource_name: The name of the resource.
        :param LiveTranscodeRuleAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LiveTranscodeRuleAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_name: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 stream_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LiveTranscodeRuleAttachmentArgs.__new__(LiveTranscodeRuleAttachmentArgs)

            if app_name is None and not opts.urn:
                raise TypeError("Missing required property 'app_name'")
            __props__.__dict__["app_name"] = app_name
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            if stream_name is None and not opts.urn:
                raise TypeError("Missing required property 'stream_name'")
            __props__.__dict__["stream_name"] = stream_name
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(LiveTranscodeRuleAttachment, __self__).__init__(
            'tencentcloud:Css/liveTranscodeRuleAttachment:LiveTranscodeRuleAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            stream_name: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'LiveTranscodeRuleAttachment':
        """
        Get an existing LiveTranscodeRuleAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_name: app name which you want to bind, can be empty string if not binding specific app name.
        :param pulumi.Input[str] create_time: create time.
        :param pulumi.Input[str] domain_name: domain name hich you want to bind the transcode template.
        :param pulumi.Input[str] stream_name: stream name which you want to bind, can be empty string if not binding specific stream.
        :param pulumi.Input[int] template_id: template created by css_live_transcode_template.
        :param pulumi.Input[str] update_time: update time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LiveTranscodeRuleAttachmentState.__new__(_LiveTranscodeRuleAttachmentState)

        __props__.__dict__["app_name"] = app_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["stream_name"] = stream_name
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["update_time"] = update_time
        return LiveTranscodeRuleAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appName")
    def app_name(self) -> pulumi.Output[str]:
        """
        app name which you want to bind, can be empty string if not binding specific app name.
        """
        return pulumi.get(self, "app_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        create time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        domain name hich you want to bind the transcode template.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="streamName")
    def stream_name(self) -> pulumi.Output[str]:
        """
        stream name which you want to bind, can be empty string if not binding specific stream.
        """
        return pulumi.get(self, "stream_name")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[int]:
        """
        template created by css_live_transcode_template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        update time.
        """
        return pulumi.get(self, "update_time")

