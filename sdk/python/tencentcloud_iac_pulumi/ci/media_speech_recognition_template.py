# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MediaSpeechRecognitionTemplateArgs', 'MediaSpeechRecognitionTemplate']

@pulumi.input_type
class MediaSpeechRecognitionTemplateArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 speech_recognition: pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs'],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MediaSpeechRecognitionTemplate resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs'] speech_recognition: audio configuration.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "speech_recognition", speech_recognition)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="speechRecognition")
    def speech_recognition(self) -> pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']:
        """
        audio configuration.
        """
        return pulumi.get(self, "speech_recognition")

    @speech_recognition.setter
    def speech_recognition(self, value: pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']):
        pulumi.set(self, "speech_recognition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _MediaSpeechRecognitionTemplateState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 speech_recognition: Optional[pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']] = None):
        """
        Input properties used for looking up and filtering MediaSpeechRecognitionTemplate resources.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs'] speech_recognition: audio configuration.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if speech_recognition is not None:
            pulumi.set(__self__, "speech_recognition", speech_recognition)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="speechRecognition")
    def speech_recognition(self) -> Optional[pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']]:
        """
        audio configuration.
        """
        return pulumi.get(self, "speech_recognition")

    @speech_recognition.setter
    def speech_recognition(self, value: Optional[pulumi.Input['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']]):
        pulumi.set(self, "speech_recognition", value)


class MediaSpeechRecognitionTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 speech_recognition: Optional[pulumi.Input[pulumi.InputType['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to create a ci media_speech_recognition_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        media_speech_recognition_template = tencentcloud.ci.MediaSpeechRecognitionTemplate("mediaSpeechRecognitionTemplate",
            bucket="terraform-ci-1308919341",
            speech_recognition=tencentcloud.ci.MediaSpeechRecognitionTemplateSpeechRecognitionArgs(
                channel_num="1",
                convert_num_mode="0",
                engine_model_type="16k_zh",
                filter_dirty="0",
                filter_modal="1",
                filter_punc="0",
                output_file_type="txt",
                res_text_format="1",
                speaker_diarization="1",
                speaker_number="0",
            ))
        ```

        ## Import

        ci media_speech_recognition_template can be imported using the bucket#templateId, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate media_speech_recognition_template terraform-ci-xxxxxx#t1d794430f2f1f4350b11e905ce2c6167e
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input[pulumi.InputType['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']] speech_recognition: audio configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MediaSpeechRecognitionTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ci media_speech_recognition_template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        media_speech_recognition_template = tencentcloud.ci.MediaSpeechRecognitionTemplate("mediaSpeechRecognitionTemplate",
            bucket="terraform-ci-1308919341",
            speech_recognition=tencentcloud.ci.MediaSpeechRecognitionTemplateSpeechRecognitionArgs(
                channel_num="1",
                convert_num_mode="0",
                engine_model_type="16k_zh",
                filter_dirty="0",
                filter_modal="1",
                filter_punc="0",
                output_file_type="txt",
                res_text_format="1",
                speaker_diarization="1",
                speaker_number="0",
            ))
        ```

        ## Import

        ci media_speech_recognition_template can be imported using the bucket#templateId, e.g.

        ```sh
         $ pulumi import tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate media_speech_recognition_template terraform-ci-xxxxxx#t1d794430f2f1f4350b11e905ce2c6167e
        ```

        :param str resource_name: The name of the resource.
        :param MediaSpeechRecognitionTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MediaSpeechRecognitionTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 speech_recognition: Optional[pulumi.Input[pulumi.InputType['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']]] = None,
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
            __props__ = MediaSpeechRecognitionTemplateArgs.__new__(MediaSpeechRecognitionTemplateArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["name"] = name
            if speech_recognition is None and not opts.urn:
                raise TypeError("Missing required property 'speech_recognition'")
            __props__.__dict__["speech_recognition"] = speech_recognition
        super(MediaSpeechRecognitionTemplate, __self__).__init__(
            'tencentcloud:Ci/mediaSpeechRecognitionTemplate:MediaSpeechRecognitionTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            speech_recognition: Optional[pulumi.Input[pulumi.InputType['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']]] = None) -> 'MediaSpeechRecognitionTemplate':
        """
        Get an existing MediaSpeechRecognitionTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: bucket name.
        :param pulumi.Input[str] name: The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        :param pulumi.Input[pulumi.InputType['MediaSpeechRecognitionTemplateSpeechRecognitionArgs']] speech_recognition: audio configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MediaSpeechRecognitionTemplateState.__new__(_MediaSpeechRecognitionTemplateState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["name"] = name
        __props__.__dict__["speech_recognition"] = speech_recognition
        return MediaSpeechRecognitionTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        bucket name.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The template name only supports `Chinese`, `English`, `numbers`, `_`, `-` and `*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="speechRecognition")
    def speech_recognition(self) -> pulumi.Output['outputs.MediaSpeechRecognitionTemplateSpeechRecognition']:
        """
        audio configuration.
        """
        return pulumi.get(self, "speech_recognition")

