# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ExternalKeyArgs', 'ExternalKey']

@pulumi.input_type
class ExternalKeyArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_material_base64: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 valid_to: Optional[pulumi.Input[int]] = None,
                 wrapping_algorithm: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExternalKey resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[str] key_material_base64: The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        :param pulumi.Input[int] valid_to: This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        :param pulumi.Input[str] wrapping_algorithm: The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        pulumi.set(__self__, "alias", alias)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_archived is not None:
            pulumi.set(__self__, "is_archived", is_archived)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if key_material_base64 is not None:
            pulumi.set(__self__, "key_material_base64", key_material_base64)
        if pending_delete_window_in_days is not None:
            pulumi.set(__self__, "pending_delete_window_in_days", pending_delete_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if valid_to is not None:
            pulumi.set(__self__, "valid_to", valid_to)
        if wrapping_algorithm is not None:
            pulumi.set(__self__, "wrapping_algorithm", wrapping_algorithm)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of CMK. The maximum is 1024 bytes.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isArchived")
    def is_archived(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_archived")

    @is_archived.setter
    def is_archived(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_archived", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="keyMaterialBase64")
    def key_material_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        """
        return pulumi.get(self, "key_material_base64")

    @key_material_base64.setter
    def key_material_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_material_base64", value)

    @property
    @pulumi.getter(name="pendingDeleteWindowInDays")
    def pending_delete_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        """
        return pulumi.get(self, "pending_delete_window_in_days")

    @pending_delete_window_in_days.setter
    def pending_delete_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_delete_window_in_days", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of CMK.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="validTo")
    def valid_to(self) -> Optional[pulumi.Input[int]]:
        """
        This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        """
        return pulumi.get(self, "valid_to")

    @valid_to.setter
    def valid_to(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "valid_to", value)

    @property
    @pulumi.getter(name="wrappingAlgorithm")
    def wrapping_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        return pulumi.get(self, "wrapping_algorithm")

    @wrapping_algorithm.setter
    def wrapping_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_algorithm", value)


@pulumi.input_type
class _ExternalKeyState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_material_base64: Optional[pulumi.Input[str]] = None,
                 key_state: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 valid_to: Optional[pulumi.Input[int]] = None,
                 wrapping_algorithm: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExternalKey resources.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[str] key_material_base64: The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        :param pulumi.Input[str] key_state: State of CMK.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        :param pulumi.Input[int] valid_to: This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        :param pulumi.Input[str] wrapping_algorithm: The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_archived is not None:
            pulumi.set(__self__, "is_archived", is_archived)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if key_material_base64 is not None:
            pulumi.set(__self__, "key_material_base64", key_material_base64)
        if key_state is not None:
            pulumi.set(__self__, "key_state", key_state)
        if pending_delete_window_in_days is not None:
            pulumi.set(__self__, "pending_delete_window_in_days", pending_delete_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if valid_to is not None:
            pulumi.set(__self__, "valid_to", valid_to)
        if wrapping_algorithm is not None:
            pulumi.set(__self__, "wrapping_algorithm", wrapping_algorithm)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of CMK. The maximum is 1024 bytes.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isArchived")
    def is_archived(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_archived")

    @is_archived.setter
    def is_archived(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_archived", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="keyMaterialBase64")
    def key_material_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        """
        return pulumi.get(self, "key_material_base64")

    @key_material_base64.setter
    def key_material_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_material_base64", value)

    @property
    @pulumi.getter(name="keyState")
    def key_state(self) -> Optional[pulumi.Input[str]]:
        """
        State of CMK.
        """
        return pulumi.get(self, "key_state")

    @key_state.setter
    def key_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_state", value)

    @property
    @pulumi.getter(name="pendingDeleteWindowInDays")
    def pending_delete_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        """
        return pulumi.get(self, "pending_delete_window_in_days")

    @pending_delete_window_in_days.setter
    def pending_delete_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_delete_window_in_days", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of CMK.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="validTo")
    def valid_to(self) -> Optional[pulumi.Input[int]]:
        """
        This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        """
        return pulumi.get(self, "valid_to")

    @valid_to.setter
    def valid_to(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "valid_to", value)

    @property
    @pulumi.getter(name="wrappingAlgorithm")
    def wrapping_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        return pulumi.get(self, "wrapping_algorithm")

    @wrapping_algorithm.setter
    def wrapping_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_algorithm", value)


class ExternalKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_material_base64: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 valid_to: Optional[pulumi.Input[int]] = None,
                 wrapping_algorithm: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provide a resource to create a KMS external key.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.kms.ExternalKey("foo",
            alias="test",
            description="describe key test message.",
            is_enabled=True,
            key_material_base64="MTIzMTIzMTIzMTIzMTIzQQ==",
            valid_to=2147443200,
            wrapping_algorithm="RSAES_PKCS1_V1_5")
        ```

        ## Import

        KMS external keys can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Kms/externalKey:ExternalKey foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[str] key_material_base64: The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        :param pulumi.Input[int] valid_to: This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        :param pulumi.Input[str] wrapping_algorithm: The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExternalKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a KMS external key.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.kms.ExternalKey("foo",
            alias="test",
            description="describe key test message.",
            is_enabled=True,
            key_material_base64="MTIzMTIzMTIzMTIzMTIzQQ==",
            valid_to=2147443200,
            wrapping_algorithm="RSAES_PKCS1_V1_5")
        ```

        ## Import

        KMS external keys can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Kms/externalKey:ExternalKey foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
        ```

        :param str resource_name: The name of the resource.
        :param ExternalKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExternalKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_material_base64: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 valid_to: Optional[pulumi.Input[int]] = None,
                 wrapping_algorithm: Optional[pulumi.Input[str]] = None,
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
            __props__ = ExternalKeyArgs.__new__(ExternalKeyArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["description"] = description
            __props__.__dict__["is_archived"] = is_archived
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["key_material_base64"] = key_material_base64
            __props__.__dict__["pending_delete_window_in_days"] = pending_delete_window_in_days
            __props__.__dict__["tags"] = tags
            __props__.__dict__["valid_to"] = valid_to
            __props__.__dict__["wrapping_algorithm"] = wrapping_algorithm
            __props__.__dict__["key_state"] = None
        super(ExternalKey, __self__).__init__(
            'tencentcloud:Kms/externalKey:ExternalKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_archived: Optional[pulumi.Input[bool]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            key_material_base64: Optional[pulumi.Input[str]] = None,
            key_state: Optional[pulumi.Input[str]] = None,
            pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            valid_to: Optional[pulumi.Input[int]] = None,
            wrapping_algorithm: Optional[pulumi.Input[str]] = None) -> 'ExternalKey':
        """
        Get an existing ExternalKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[str] key_material_base64: The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        :param pulumi.Input[str] key_state: State of CMK.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        :param pulumi.Input[int] valid_to: This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        :param pulumi.Input[str] wrapping_algorithm: The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExternalKeyState.__new__(_ExternalKeyState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["description"] = description
        __props__.__dict__["is_archived"] = is_archived
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["key_material_base64"] = key_material_base64
        __props__.__dict__["key_state"] = key_state
        __props__.__dict__["pending_delete_window_in_days"] = pending_delete_window_in_days
        __props__.__dict__["tags"] = tags
        __props__.__dict__["valid_to"] = valid_to
        __props__.__dict__["wrapping_algorithm"] = wrapping_algorithm
        return ExternalKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of CMK. The maximum is 1024 bytes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isArchived")
    def is_archived(self) -> pulumi.Output[Optional[bool]]:
        """
        Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_archived")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="keyMaterialBase64")
    def key_material_base64(self) -> pulumi.Output[Optional[str]]:
        """
        The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        """
        return pulumi.get(self, "key_material_base64")

    @property
    @pulumi.getter(name="keyState")
    def key_state(self) -> pulumi.Output[str]:
        """
        State of CMK.
        """
        return pulumi.get(self, "key_state")

    @property
    @pulumi.getter(name="pendingDeleteWindowInDays")
    def pending_delete_window_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        """
        return pulumi.get(self, "pending_delete_window_in_days")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tags of CMK.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="validTo")
    def valid_to(self) -> pulumi.Output[Optional[int]]:
        """
        This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        """
        return pulumi.get(self, "valid_to")

    @property
    @pulumi.getter(name="wrappingAlgorithm")
    def wrapping_algorithm(self) -> pulumi.Output[Optional[str]]:
        """
        The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        """
        return pulumi.get(self, "wrapping_algorithm")

