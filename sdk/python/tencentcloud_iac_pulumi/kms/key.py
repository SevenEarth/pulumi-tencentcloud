# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_rotation_enabled: Optional[pulumi.Input[bool]] = None,
                 key_usage: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] key_rotation_enabled: Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        :param pulumi.Input[str] key_usage: Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        """
        pulumi.set(__self__, "alias", alias)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_archived is not None:
            pulumi.set(__self__, "is_archived", is_archived)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if key_rotation_enabled is not None:
            pulumi.set(__self__, "key_rotation_enabled", key_rotation_enabled)
        if key_usage is not None:
            pulumi.set(__self__, "key_usage", key_usage)
        if pending_delete_window_in_days is not None:
            pulumi.set(__self__, "pending_delete_window_in_days", pending_delete_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="keyRotationEnabled")
    def key_rotation_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        """
        return pulumi.get(self, "key_rotation_enabled")

    @key_rotation_enabled.setter
    def key_rotation_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "key_rotation_enabled", value)

    @property
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> Optional[pulumi.Input[str]]:
        """
        Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        """
        return pulumi.get(self, "key_usage")

    @key_usage.setter
    def key_usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_usage", value)

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


@pulumi.input_type
class _KeyState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_rotation_enabled: Optional[pulumi.Input[bool]] = None,
                 key_state: Optional[pulumi.Input[str]] = None,
                 key_usage: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Key resources.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] key_rotation_enabled: Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        :param pulumi.Input[str] key_state: State of CMK.
        :param pulumi.Input[str] key_usage: Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_archived is not None:
            pulumi.set(__self__, "is_archived", is_archived)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if key_rotation_enabled is not None:
            pulumi.set(__self__, "key_rotation_enabled", key_rotation_enabled)
        if key_state is not None:
            pulumi.set(__self__, "key_state", key_state)
        if key_usage is not None:
            pulumi.set(__self__, "key_usage", key_usage)
        if pending_delete_window_in_days is not None:
            pulumi.set(__self__, "pending_delete_window_in_days", pending_delete_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="keyRotationEnabled")
    def key_rotation_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        """
        return pulumi.get(self, "key_rotation_enabled")

    @key_rotation_enabled.setter
    def key_rotation_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "key_rotation_enabled", value)

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
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> Optional[pulumi.Input[str]]:
        """
        Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        """
        return pulumi.get(self, "key_usage")

    @key_usage.setter
    def key_usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_usage", value)

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


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_archived: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 key_rotation_enabled: Optional[pulumi.Input[bool]] = None,
                 key_usage: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provide a resource to create a KMS key.

        ## Example Usage
        ### Create and enable a instance.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.kms.Key("example",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=True,
            key_rotation_enabled=False,
            tags={
                "createdBy": "terraform",
            })
        ```
        ### Specify the Key Usage as an asymmetry method.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example2 = tencentcloud.kms.Key("example2",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=False,
            key_usage="ASYMMETRIC_DECRYPT_RSA_2048")
        ```
        ### Disable the kms key instance.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example3 = tencentcloud.kms.Key("example3",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=False,
            key_rotation_enabled=False,
            tags={
                "test-tag": "unit-test",
            })
        ```

        ## Import

        KMS keys can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Kms/key:Key foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] key_rotation_enabled: Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        :param pulumi.Input[str] key_usage: Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a KMS key.

        ## Example Usage
        ### Create and enable a instance.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example = tencentcloud.kms.Key("example",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=True,
            key_rotation_enabled=False,
            tags={
                "createdBy": "terraform",
            })
        ```
        ### Specify the Key Usage as an asymmetry method.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example2 = tencentcloud.kms.Key("example2",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=False,
            key_usage="ASYMMETRIC_DECRYPT_RSA_2048")
        ```
        ### Disable the kms key instance.

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example3 = tencentcloud.kms.Key("example3",
            alias="tf-example-kms-key",
            description="example of kms key",
            is_enabled=False,
            key_rotation_enabled=False,
            tags={
                "test-tag": "unit-test",
            })
        ```

        ## Import

        KMS keys can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Kms/key:Key foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
        ```

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
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
                 key_rotation_enabled: Optional[pulumi.Input[bool]] = None,
                 key_usage: Optional[pulumi.Input[str]] = None,
                 pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
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
            __props__ = KeyArgs.__new__(KeyArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["description"] = description
            __props__.__dict__["is_archived"] = is_archived
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["key_rotation_enabled"] = key_rotation_enabled
            __props__.__dict__["key_usage"] = key_usage
            __props__.__dict__["pending_delete_window_in_days"] = pending_delete_window_in_days
            __props__.__dict__["tags"] = tags
            __props__.__dict__["key_state"] = None
        super(Key, __self__).__init__(
            'tencentcloud:Kms/key:Key',
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
            key_rotation_enabled: Optional[pulumi.Input[bool]] = None,
            key_state: Optional[pulumi.Input[str]] = None,
            key_usage: Optional[pulumi.Input[str]] = None,
            pending_delete_window_in_days: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] description: Description of CMK. The maximum is 1024 bytes.
        :param pulumi.Input[bool] is_archived: Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        :param pulumi.Input[bool] key_rotation_enabled: Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        :param pulumi.Input[str] key_state: State of CMK.
        :param pulumi.Input[str] key_usage: Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        :param pulumi.Input[int] pending_delete_window_in_days: Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of CMK.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyState.__new__(_KeyState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["description"] = description
        __props__.__dict__["is_archived"] = is_archived
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["key_rotation_enabled"] = key_rotation_enabled
        __props__.__dict__["key_state"] = key_state
        __props__.__dict__["key_usage"] = key_usage
        __props__.__dict__["pending_delete_window_in_days"] = pending_delete_window_in_days
        __props__.__dict__["tags"] = tags
        return Key(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="keyRotationEnabled")
    def key_rotation_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specify whether to enable key rotation, valid when key_usage is `ENCRYPT_DECRYPT`. Default value is `false`.
        """
        return pulumi.get(self, "key_rotation_enabled")

    @property
    @pulumi.getter(name="keyState")
    def key_state(self) -> pulumi.Output[str]:
        """
        State of CMK.
        """
        return pulumi.get(self, "key_state")

    @property
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> pulumi.Output[Optional[str]]:
        """
        Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        """
        return pulumi.get(self, "key_usage")

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

