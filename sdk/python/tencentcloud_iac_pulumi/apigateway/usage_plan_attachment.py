# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UsagePlanAttachmentArgs', 'UsagePlanAttachment']

@pulumi.input_type
class UsagePlanAttachmentArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 usage_plan_id: pulumi.Input[str],
                 access_key_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 bind_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UsagePlanAttachment resource.
        :param pulumi.Input[str] environment: The environment to be bound. Valid values: `test`, `prepub`, `release`.
        :param pulumi.Input[str] service_id: ID of the service.
        :param pulumi.Input[str] usage_plan_id: ID of the usage plan.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_key_ids: Array of key IDs to be bound.
        :param pulumi.Input[str] api_id: ID of the API. This parameter will be required when `bind_type` is `API`.
        :param pulumi.Input[str] bind_type: Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "service_id", service_id)
        pulumi.set(__self__, "usage_plan_id", usage_plan_id)
        if access_key_ids is not None:
            pulumi.set(__self__, "access_key_ids", access_key_ids)
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if bind_type is not None:
            pulumi.set(__self__, "bind_type", bind_type)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        The environment to be bound. Valid values: `test`, `prepub`, `release`.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        ID of the service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> pulumi.Input[str]:
        """
        ID of the usage plan.
        """
        return pulumi.get(self, "usage_plan_id")

    @usage_plan_id.setter
    def usage_plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "usage_plan_id", value)

    @property
    @pulumi.getter(name="accessKeyIds")
    def access_key_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of key IDs to be bound.
        """
        return pulumi.get(self, "access_key_ids")

    @access_key_ids.setter
    def access_key_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "access_key_ids", value)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the API. This parameter will be required when `bind_type` is `API`.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="bindType")
    def bind_type(self) -> Optional[pulumi.Input[str]]:
        """
        Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        """
        return pulumi.get(self, "bind_type")

    @bind_type.setter
    def bind_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_type", value)


@pulumi.input_type
class _UsagePlanAttachmentState:
    def __init__(__self__, *,
                 access_key_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 bind_type: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UsagePlanAttachment resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_key_ids: Array of key IDs to be bound.
        :param pulumi.Input[str] api_id: ID of the API. This parameter will be required when `bind_type` is `API`.
        :param pulumi.Input[str] bind_type: Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        :param pulumi.Input[str] environment: The environment to be bound. Valid values: `test`, `prepub`, `release`.
        :param pulumi.Input[str] service_id: ID of the service.
        :param pulumi.Input[str] usage_plan_id: ID of the usage plan.
        """
        if access_key_ids is not None:
            pulumi.set(__self__, "access_key_ids", access_key_ids)
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if bind_type is not None:
            pulumi.set(__self__, "bind_type", bind_type)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if usage_plan_id is not None:
            pulumi.set(__self__, "usage_plan_id", usage_plan_id)

    @property
    @pulumi.getter(name="accessKeyIds")
    def access_key_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of key IDs to be bound.
        """
        return pulumi.get(self, "access_key_ids")

    @access_key_ids.setter
    def access_key_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "access_key_ids", value)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the API. This parameter will be required when `bind_type` is `API`.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="bindType")
    def bind_type(self) -> Optional[pulumi.Input[str]]:
        """
        Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        """
        return pulumi.get(self, "bind_type")

    @bind_type.setter
    def bind_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bind_type", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The environment to be bound. Valid values: `test`, `prepub`, `release`.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the usage plan.
        """
        return pulumi.get(self, "usage_plan_id")

    @usage_plan_id.setter
    def usage_plan_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage_plan_id", value)


class UsagePlanAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 bind_type: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to attach API gateway usage plan to service.

        > **NOTE:** If the `auth_type` parameter of API is not `SECRET`, it cannot be bound access key.

        ## Example Usage
        ### Normal creation

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_usage_plan = tencentcloud.api_gateway.UsagePlan("exampleUsagePlan",
            usage_plan_name="tf_example",
            usage_plan_desc="desc.",
            max_request_num=100,
            max_request_num_pre_sec=10)
        example_service = tencentcloud.api_gateway.Service("exampleService",
            service_name="tf_example",
            protocol="http&https",
            service_desc="desc.",
            net_types=[
                "INNER",
                "OUTER",
            ],
            ip_version="IPv4")
        example_api = tencentcloud.api_gateway.Api("exampleApi",
            service_id=example_service.id,
            api_name="tf_example",
            api_desc="my hello api update",
            auth_type="SECRET",
            protocol="HTTP",
            enable_cors=True,
            request_config_path="/user/info",
            request_config_method="POST",
            request_parameters=[tencentcloud.api.gateway.ApiRequestParameterArgs(
                name="email",
                position="QUERY",
                type="string",
                desc="desc.",
                default_value="test@qq.com",
                required=True,
            )],
            service_config_type="HTTP",
            service_config_timeout=10,
            service_config_url="http://www.tencent.com",
            service_config_path="/user",
            service_config_method="POST",
            response_type="XML",
            response_success_example="<note>success</note>",
            response_fail_example="<note>fail</note>",
            response_error_codes=[tencentcloud.api.gateway.ApiResponseErrorCodeArgs(
                code=500,
                msg="system error",
                desc="system error code",
                converted_code=5000,
                need_convert=True,
            )])
        example_usage_plan_attachment = tencentcloud.api_gateway.UsagePlanAttachment("exampleUsagePlanAttachment",
            usage_plan_id=example_usage_plan.id,
            service_id=example_service.id,
            environment="release",
            bind_type="API",
            api_id=example_api.id)
        ```
        ### Bind the key to a usage plan

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_api_key = tencentcloud.api_gateway.ApiKey("exampleApiKey",
            secret_name="tf_example",
            status="on")
        example_usage_plan_attachment = tencentcloud.api_gateway.UsagePlanAttachment("exampleUsagePlanAttachment",
            usage_plan_id=tencentcloud_api_gateway_usage_plan["example"]["id"],
            service_id=tencentcloud_api_gateway_service["example"]["id"],
            environment="release",
            bind_type="API",
            api_id=tencentcloud_api_gateway_api["example"]["id"],
            access_key_ids=[example_api_key.id])
        ```

        ## Import

        API gateway usage plan attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment attach_service usagePlan-pe7fbdgn#service-kuqd6xqk#release#API#api-p8gtanvy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_key_ids: Array of key IDs to be bound.
        :param pulumi.Input[str] api_id: ID of the API. This parameter will be required when `bind_type` is `API`.
        :param pulumi.Input[str] bind_type: Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        :param pulumi.Input[str] environment: The environment to be bound. Valid values: `test`, `prepub`, `release`.
        :param pulumi.Input[str] service_id: ID of the service.
        :param pulumi.Input[str] usage_plan_id: ID of the usage plan.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UsagePlanAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to attach API gateway usage plan to service.

        > **NOTE:** If the `auth_type` parameter of API is not `SECRET`, it cannot be bound access key.

        ## Example Usage
        ### Normal creation

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_usage_plan = tencentcloud.api_gateway.UsagePlan("exampleUsagePlan",
            usage_plan_name="tf_example",
            usage_plan_desc="desc.",
            max_request_num=100,
            max_request_num_pre_sec=10)
        example_service = tencentcloud.api_gateway.Service("exampleService",
            service_name="tf_example",
            protocol="http&https",
            service_desc="desc.",
            net_types=[
                "INNER",
                "OUTER",
            ],
            ip_version="IPv4")
        example_api = tencentcloud.api_gateway.Api("exampleApi",
            service_id=example_service.id,
            api_name="tf_example",
            api_desc="my hello api update",
            auth_type="SECRET",
            protocol="HTTP",
            enable_cors=True,
            request_config_path="/user/info",
            request_config_method="POST",
            request_parameters=[tencentcloud.api.gateway.ApiRequestParameterArgs(
                name="email",
                position="QUERY",
                type="string",
                desc="desc.",
                default_value="test@qq.com",
                required=True,
            )],
            service_config_type="HTTP",
            service_config_timeout=10,
            service_config_url="http://www.tencent.com",
            service_config_path="/user",
            service_config_method="POST",
            response_type="XML",
            response_success_example="<note>success</note>",
            response_fail_example="<note>fail</note>",
            response_error_codes=[tencentcloud.api.gateway.ApiResponseErrorCodeArgs(
                code=500,
                msg="system error",
                desc="system error code",
                converted_code=5000,
                need_convert=True,
            )])
        example_usage_plan_attachment = tencentcloud.api_gateway.UsagePlanAttachment("exampleUsagePlanAttachment",
            usage_plan_id=example_usage_plan.id,
            service_id=example_service.id,
            environment="release",
            bind_type="API",
            api_id=example_api.id)
        ```
        ### Bind the key to a usage plan

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_api_key = tencentcloud.api_gateway.ApiKey("exampleApiKey",
            secret_name="tf_example",
            status="on")
        example_usage_plan_attachment = tencentcloud.api_gateway.UsagePlanAttachment("exampleUsagePlanAttachment",
            usage_plan_id=tencentcloud_api_gateway_usage_plan["example"]["id"],
            service_id=tencentcloud_api_gateway_service["example"]["id"],
            environment="release",
            bind_type="API",
            api_id=tencentcloud_api_gateway_api["example"]["id"],
            access_key_ids=[example_api_key.id])
        ```

        ## Import

        API gateway usage plan attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment attach_service usagePlan-pe7fbdgn#service-kuqd6xqk#release#API#api-p8gtanvy
        ```

        :param str resource_name: The name of the resource.
        :param UsagePlanAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsagePlanAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 bind_type: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 usage_plan_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = UsagePlanAttachmentArgs.__new__(UsagePlanAttachmentArgs)

            __props__.__dict__["access_key_ids"] = access_key_ids
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["bind_type"] = bind_type
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            if usage_plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'usage_plan_id'")
            __props__.__dict__["usage_plan_id"] = usage_plan_id
        super(UsagePlanAttachment, __self__).__init__(
            'tencentcloud:ApiGateway/usagePlanAttachment:UsagePlanAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            bind_type: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            usage_plan_id: Optional[pulumi.Input[str]] = None) -> 'UsagePlanAttachment':
        """
        Get an existing UsagePlanAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] access_key_ids: Array of key IDs to be bound.
        :param pulumi.Input[str] api_id: ID of the API. This parameter will be required when `bind_type` is `API`.
        :param pulumi.Input[str] bind_type: Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        :param pulumi.Input[str] environment: The environment to be bound. Valid values: `test`, `prepub`, `release`.
        :param pulumi.Input[str] service_id: ID of the service.
        :param pulumi.Input[str] usage_plan_id: ID of the usage plan.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UsagePlanAttachmentState.__new__(_UsagePlanAttachmentState)

        __props__.__dict__["access_key_ids"] = access_key_ids
        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["bind_type"] = bind_type
        __props__.__dict__["environment"] = environment
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["usage_plan_id"] = usage_plan_id
        return UsagePlanAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKeyIds")
    def access_key_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Array of key IDs to be bound.
        """
        return pulumi.get(self, "access_key_ids")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the API. This parameter will be required when `bind_type` is `API`.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="bindType")
    def bind_type(self) -> pulumi.Output[Optional[str]]:
        """
        Binding type. Valid values: `API`, `SERVICE`. Default value is `SERVICE`.
        """
        return pulumi.get(self, "bind_type")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The environment to be bound. Valid values: `test`, `prepub`, `release`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        ID of the service.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="usagePlanId")
    def usage_plan_id(self) -> pulumi.Output[str]:
        """
        ID of the usage plan.
        """
        return pulumi.get(self, "usage_plan_id")

