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

__all__ = ['ParamTemplateArgs', 'ParamTemplate']

@pulumi.input_type
class ParamTemplateArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 params_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]] = None,
                 product_type: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ParamTemplate resource.
        :param pulumi.Input[str] description: Parameter template description.
        :param pulumi.Input[str] name: Parameter template name.
        :param pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]] params_overrides: Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        :param pulumi.Input[int] product_type: Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        :param pulumi.Input[str] template_id: Specify which existed template import from.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if params_overrides is not None:
            pulumi.set(__self__, "params_overrides", params_overrides)
        if product_type is not None:
            pulumi.set(__self__, "product_type", product_type)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paramsOverrides")
    def params_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]]:
        """
        Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        """
        return pulumi.get(self, "params_overrides")

    @params_overrides.setter
    def params_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]]):
        pulumi.set(self, "params_overrides", value)

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> Optional[pulumi.Input[int]]:
        """
        Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        """
        return pulumi.get(self, "product_type")

    @product_type.setter
    def product_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "product_type", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specify which existed template import from.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class _ParamTemplateState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_details: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamDetailArgs']]]] = None,
                 params_overrides: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]] = None,
                 product_type: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ParamTemplate resources.
        :param pulumi.Input[str] description: Parameter template description.
        :param pulumi.Input[str] name: Parameter template name.
        :param pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamDetailArgs']]] param_details: Readonly full parameter list details.
        :param pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]] params_overrides: Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        :param pulumi.Input[int] product_type: Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        :param pulumi.Input[str] template_id: Specify which existed template import from.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if param_details is not None:
            pulumi.set(__self__, "param_details", param_details)
        if params_overrides is not None:
            pulumi.set(__self__, "params_overrides", params_overrides)
        if product_type is not None:
            pulumi.set(__self__, "product_type", product_type)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paramDetails")
    def param_details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamDetailArgs']]]]:
        """
        Readonly full parameter list details.
        """
        return pulumi.get(self, "param_details")

    @param_details.setter
    def param_details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamDetailArgs']]]]):
        pulumi.set(self, "param_details", value)

    @property
    @pulumi.getter(name="paramsOverrides")
    def params_overrides(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]]:
        """
        Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        """
        return pulumi.get(self, "params_overrides")

    @params_overrides.setter
    def params_overrides(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParamTemplateParamsOverrideArgs']]]]):
        pulumi.set(self, "params_overrides", value)

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> Optional[pulumi.Input[int]]:
        """
        Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        """
        return pulumi.get(self, "product_type")

    @product_type.setter
    def product_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "product_type", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specify which existed template import from.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


class ParamTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 params_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamsOverrideArgs']]]]] = None,
                 product_type: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a redis parameter template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        param_template = tencentcloud.redis.ParamTemplate("paramTemplate",
            description="This is an example redis param template.",
            params_overrides=[tencentcloud.redis.ParamTemplateParamsOverrideArgs(
                key="timeout",
                value="7200",
            )],
            product_type=6)
        ```

        Copy from another template

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        param_template = tencentcloud.redis.ParamTemplate("paramTemplate",
            description="This is an copied redis param template from xxx.",
            params_overrides=[tencentcloud.redis.ParamTemplateParamsOverrideArgs(
                key="timeout",
                value="7200",
            )],
            template_id="xxx")
        ```

        ## Import

        redis param_template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Redis/paramTemplate:ParamTemplate param_template param_template_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Parameter template description.
        :param pulumi.Input[str] name: Parameter template name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamsOverrideArgs']]]] params_overrides: Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        :param pulumi.Input[int] product_type: Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        :param pulumi.Input[str] template_id: Specify which existed template import from.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ParamTemplateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a redis parameter template

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        param_template = tencentcloud.redis.ParamTemplate("paramTemplate",
            description="This is an example redis param template.",
            params_overrides=[tencentcloud.redis.ParamTemplateParamsOverrideArgs(
                key="timeout",
                value="7200",
            )],
            product_type=6)
        ```

        Copy from another template

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        param_template = tencentcloud.redis.ParamTemplate("paramTemplate",
            description="This is an copied redis param template from xxx.",
            params_overrides=[tencentcloud.redis.ParamTemplateParamsOverrideArgs(
                key="timeout",
                value="7200",
            )],
            template_id="xxx")
        ```

        ## Import

        redis param_template can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Redis/paramTemplate:ParamTemplate param_template param_template_id
        ```

        :param str resource_name: The name of the resource.
        :param ParamTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParamTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 params_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamsOverrideArgs']]]]] = None,
                 product_type: Optional[pulumi.Input[int]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ParamTemplateArgs.__new__(ParamTemplateArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["params_overrides"] = params_overrides
            __props__.__dict__["product_type"] = product_type
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["param_details"] = None
        super(ParamTemplate, __self__).__init__(
            'tencentcloud:Redis/paramTemplate:ParamTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            param_details: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamDetailArgs']]]]] = None,
            params_overrides: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamsOverrideArgs']]]]] = None,
            product_type: Optional[pulumi.Input[int]] = None,
            template_id: Optional[pulumi.Input[str]] = None) -> 'ParamTemplate':
        """
        Get an existing ParamTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Parameter template description.
        :param pulumi.Input[str] name: Parameter template name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamDetailArgs']]]] param_details: Readonly full parameter list details.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParamTemplateParamsOverrideArgs']]]] params_overrides: Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        :param pulumi.Input[int] product_type: Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        :param pulumi.Input[str] template_id: Specify which existed template import from.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ParamTemplateState.__new__(_ParamTemplateState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["param_details"] = param_details
        __props__.__dict__["params_overrides"] = params_overrides
        __props__.__dict__["product_type"] = product_type
        __props__.__dict__["template_id"] = template_id
        return ParamTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Parameter template description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Parameter template name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paramDetails")
    def param_details(self) -> pulumi.Output[Sequence['outputs.ParamTemplateParamDetail']]:
        """
        Readonly full parameter list details.
        """
        return pulumi.get(self, "param_details")

    @property
    @pulumi.getter(name="paramsOverrides")
    def params_overrides(self) -> pulumi.Output[Optional[Sequence['outputs.ParamTemplateParamsOverride']]]:
        """
        Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        """
        return pulumi.get(self, "params_overrides")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> pulumi.Output[Optional[int]]:
        """
        Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        """
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specify which existed template import from.
        """
        return pulumi.get(self, "template_id")

