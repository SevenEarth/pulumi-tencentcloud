# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CcSessionArgs', 'CcSession']

@pulumi.input_type
class CcSessionArgs:
    def __init__(__self__, *,
                 category: pulumi.Input[str],
                 domain: pulumi.Input[str],
                 edition: pulumi.Input[str],
                 end_mat: pulumi.Input[str],
                 end_offset: pulumi.Input[str],
                 key_or_start_mat: pulumi.Input[str],
                 session_name: pulumi.Input[str],
                 source: pulumi.Input[str],
                 start_offset: pulumi.Input[str]):
        """
        The set of arguments for constructing a CcSession resource.
        :param pulumi.Input[str] category: Session match pattern, Optional patterns are match, location.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] end_mat: Session end identifier, when Category is match.
        :param pulumi.Input[str] end_offset: End offset position, when Category is location.
        :param pulumi.Input[str] key_or_start_mat: Session identifier.
        :param pulumi.Input[str] session_name: Session Name.
        :param pulumi.Input[str] source: Session matching position, Optional locations are get, post, header, cookie.
        :param pulumi.Input[str] start_offset: Starting offset position, when Category is location.
        """
        pulumi.set(__self__, "category", category)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "edition", edition)
        pulumi.set(__self__, "end_mat", end_mat)
        pulumi.set(__self__, "end_offset", end_offset)
        pulumi.set(__self__, "key_or_start_mat", key_or_start_mat)
        pulumi.set(__self__, "session_name", session_name)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "start_offset", start_offset)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Input[str]:
        """
        Session match pattern, Optional patterns are match, location.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: pulumi.Input[str]):
        pulumi.set(self, "category", value)

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
    def edition(self) -> pulumi.Input[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: pulumi.Input[str]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="endMat")
    def end_mat(self) -> pulumi.Input[str]:
        """
        Session end identifier, when Category is match.
        """
        return pulumi.get(self, "end_mat")

    @end_mat.setter
    def end_mat(self, value: pulumi.Input[str]):
        pulumi.set(self, "end_mat", value)

    @property
    @pulumi.getter(name="endOffset")
    def end_offset(self) -> pulumi.Input[str]:
        """
        End offset position, when Category is location.
        """
        return pulumi.get(self, "end_offset")

    @end_offset.setter
    def end_offset(self, value: pulumi.Input[str]):
        pulumi.set(self, "end_offset", value)

    @property
    @pulumi.getter(name="keyOrStartMat")
    def key_or_start_mat(self) -> pulumi.Input[str]:
        """
        Session identifier.
        """
        return pulumi.get(self, "key_or_start_mat")

    @key_or_start_mat.setter
    def key_or_start_mat(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_or_start_mat", value)

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> pulumi.Input[str]:
        """
        Session Name.
        """
        return pulumi.get(self, "session_name")

    @session_name.setter
    def session_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "session_name", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        Session matching position, Optional locations are get, post, header, cookie.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="startOffset")
    def start_offset(self) -> pulumi.Input[str]:
        """
        Starting offset position, when Category is location.
        """
        return pulumi.get(self, "start_offset")

    @start_offset.setter
    def start_offset(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_offset", value)


@pulumi.input_type
class _CcSessionState:
    def __init__(__self__, *,
                 category: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 end_mat: Optional[pulumi.Input[str]] = None,
                 end_offset: Optional[pulumi.Input[str]] = None,
                 key_or_start_mat: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[int]] = None,
                 session_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 start_offset: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CcSession resources.
        :param pulumi.Input[str] category: Session match pattern, Optional patterns are match, location.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] end_mat: Session end identifier, when Category is match.
        :param pulumi.Input[str] end_offset: End offset position, when Category is location.
        :param pulumi.Input[str] key_or_start_mat: Session identifier.
        :param pulumi.Input[int] session_id: Session ID.
        :param pulumi.Input[str] session_name: Session Name.
        :param pulumi.Input[str] source: Session matching position, Optional locations are get, post, header, cookie.
        :param pulumi.Input[str] start_offset: Starting offset position, when Category is location.
        """
        if category is not None:
            pulumi.set(__self__, "category", category)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if end_mat is not None:
            pulumi.set(__self__, "end_mat", end_mat)
        if end_offset is not None:
            pulumi.set(__self__, "end_offset", end_offset)
        if key_or_start_mat is not None:
            pulumi.set(__self__, "key_or_start_mat", key_or_start_mat)
        if session_id is not None:
            pulumi.set(__self__, "session_id", session_id)
        if session_name is not None:
            pulumi.set(__self__, "session_name", session_name)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if start_offset is not None:
            pulumi.set(__self__, "start_offset", start_offset)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        Session match pattern, Optional patterns are match, location.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

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
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="endMat")
    def end_mat(self) -> Optional[pulumi.Input[str]]:
        """
        Session end identifier, when Category is match.
        """
        return pulumi.get(self, "end_mat")

    @end_mat.setter
    def end_mat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_mat", value)

    @property
    @pulumi.getter(name="endOffset")
    def end_offset(self) -> Optional[pulumi.Input[str]]:
        """
        End offset position, when Category is location.
        """
        return pulumi.get(self, "end_offset")

    @end_offset.setter
    def end_offset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_offset", value)

    @property
    @pulumi.getter(name="keyOrStartMat")
    def key_or_start_mat(self) -> Optional[pulumi.Input[str]]:
        """
        Session identifier.
        """
        return pulumi.get(self, "key_or_start_mat")

    @key_or_start_mat.setter
    def key_or_start_mat(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_or_start_mat", value)

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> Optional[pulumi.Input[int]]:
        """
        Session ID.
        """
        return pulumi.get(self, "session_id")

    @session_id.setter
    def session_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_id", value)

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> Optional[pulumi.Input[str]]:
        """
        Session Name.
        """
        return pulumi.get(self, "session_name")

    @session_name.setter
    def session_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Session matching position, Optional locations are get, post, header, cookie.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="startOffset")
    def start_offset(self) -> Optional[pulumi.Input[str]]:
        """
        Starting offset position, when Category is location.
        """
        return pulumi.get(self, "start_offset")

    @start_offset.setter
    def start_offset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_offset", value)


class CcSession(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 end_mat: Optional[pulumi.Input[str]] = None,
                 end_offset: Optional[pulumi.Input[str]] = None,
                 key_or_start_mat: Optional[pulumi.Input[str]] = None,
                 session_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 start_offset: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a waf cc_session

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.waf.CcSession("example",
            category="match",
            domain="www.demo.com",
            edition="sparta-waf",
            end_mat="&",
            end_offset="-1",
            key_or_start_mat="key_a=123",
            session_name="terraformDemo",
            source="get",
            start_offset="-1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        waf cc_session can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Waf/ccSession:CcSession example www.demo.com#sparta-waf#2000000253
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: Session match pattern, Optional patterns are match, location.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] end_mat: Session end identifier, when Category is match.
        :param pulumi.Input[str] end_offset: End offset position, when Category is location.
        :param pulumi.Input[str] key_or_start_mat: Session identifier.
        :param pulumi.Input[str] session_name: Session Name.
        :param pulumi.Input[str] source: Session matching position, Optional locations are get, post, header, cookie.
        :param pulumi.Input[str] start_offset: Starting offset position, when Category is location.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CcSessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a waf cc_session

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.waf.CcSession("example",
            category="match",
            domain="www.demo.com",
            edition="sparta-waf",
            end_mat="&",
            end_offset="-1",
            key_or_start_mat="key_a=123",
            session_name="terraformDemo",
            source="get",
            start_offset="-1")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        waf cc_session can be imported using the id, e.g.

        ```sh
        $ pulumi import tencentcloud:Waf/ccSession:CcSession example www.demo.com#sparta-waf#2000000253
        ```

        :param str resource_name: The name of the resource.
        :param CcSessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CcSessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 end_mat: Optional[pulumi.Input[str]] = None,
                 end_offset: Optional[pulumi.Input[str]] = None,
                 key_or_start_mat: Optional[pulumi.Input[str]] = None,
                 session_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 start_offset: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CcSessionArgs.__new__(CcSessionArgs)

            if category is None and not opts.urn:
                raise TypeError("Missing required property 'category'")
            __props__.__dict__["category"] = category
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if edition is None and not opts.urn:
                raise TypeError("Missing required property 'edition'")
            __props__.__dict__["edition"] = edition
            if end_mat is None and not opts.urn:
                raise TypeError("Missing required property 'end_mat'")
            __props__.__dict__["end_mat"] = end_mat
            if end_offset is None and not opts.urn:
                raise TypeError("Missing required property 'end_offset'")
            __props__.__dict__["end_offset"] = end_offset
            if key_or_start_mat is None and not opts.urn:
                raise TypeError("Missing required property 'key_or_start_mat'")
            __props__.__dict__["key_or_start_mat"] = key_or_start_mat
            if session_name is None and not opts.urn:
                raise TypeError("Missing required property 'session_name'")
            __props__.__dict__["session_name"] = session_name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            if start_offset is None and not opts.urn:
                raise TypeError("Missing required property 'start_offset'")
            __props__.__dict__["start_offset"] = start_offset
            __props__.__dict__["session_id"] = None
        super(CcSession, __self__).__init__(
            'tencentcloud:Waf/ccSession:CcSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            category: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            end_mat: Optional[pulumi.Input[str]] = None,
            end_offset: Optional[pulumi.Input[str]] = None,
            key_or_start_mat: Optional[pulumi.Input[str]] = None,
            session_id: Optional[pulumi.Input[int]] = None,
            session_name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            start_offset: Optional[pulumi.Input[str]] = None) -> 'CcSession':
        """
        Get an existing CcSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: Session match pattern, Optional patterns are match, location.
        :param pulumi.Input[str] domain: Domain.
        :param pulumi.Input[str] edition: Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        :param pulumi.Input[str] end_mat: Session end identifier, when Category is match.
        :param pulumi.Input[str] end_offset: End offset position, when Category is location.
        :param pulumi.Input[str] key_or_start_mat: Session identifier.
        :param pulumi.Input[int] session_id: Session ID.
        :param pulumi.Input[str] session_name: Session Name.
        :param pulumi.Input[str] source: Session matching position, Optional locations are get, post, header, cookie.
        :param pulumi.Input[str] start_offset: Starting offset position, when Category is location.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CcSessionState.__new__(_CcSessionState)

        __props__.__dict__["category"] = category
        __props__.__dict__["domain"] = domain
        __props__.__dict__["edition"] = edition
        __props__.__dict__["end_mat"] = end_mat
        __props__.__dict__["end_offset"] = end_offset
        __props__.__dict__["key_or_start_mat"] = key_or_start_mat
        __props__.__dict__["session_id"] = session_id
        __props__.__dict__["session_name"] = session_name
        __props__.__dict__["source"] = source
        __props__.__dict__["start_offset"] = start_offset
        return CcSession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        """
        Session match pattern, Optional patterns are match, location.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        Domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="endMat")
    def end_mat(self) -> pulumi.Output[str]:
        """
        Session end identifier, when Category is match.
        """
        return pulumi.get(self, "end_mat")

    @property
    @pulumi.getter(name="endOffset")
    def end_offset(self) -> pulumi.Output[str]:
        """
        End offset position, when Category is location.
        """
        return pulumi.get(self, "end_offset")

    @property
    @pulumi.getter(name="keyOrStartMat")
    def key_or_start_mat(self) -> pulumi.Output[str]:
        """
        Session identifier.
        """
        return pulumi.get(self, "key_or_start_mat")

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> pulumi.Output[int]:
        """
        Session ID.
        """
        return pulumi.get(self, "session_id")

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> pulumi.Output[str]:
        """
        Session Name.
        """
        return pulumi.get(self, "session_name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Session matching position, Optional locations are get, post, header, cookie.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="startOffset")
    def start_offset(self) -> pulumi.Output[str]:
        """
        Starting offset position, when Category is location.
        """
        return pulumi.get(self, "start_offset")

