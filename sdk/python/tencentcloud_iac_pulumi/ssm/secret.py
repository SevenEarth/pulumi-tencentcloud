# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretArgs', 'Secret']

@pulumi.input_type
class SecretArgs:
    def __init__(__self__, *,
                 secret_name: pulumi.Input[str],
                 additional_config: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 secret_type: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Secret resource.
        :param pulumi.Input[str] secret_name: Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[str] additional_config: Additional config for specific secret types in JSON string format.
        :param pulumi.Input[str] description: Description of secret. The maximum is 2048 bytes.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable secret. Default value is `true`.
        :param pulumi.Input[str] kms_key_id: KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        :param pulumi.Input[int] recovery_window_in_days: Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        :param pulumi.Input[int] secret_type: Type of secret. `0`: user-defined secret. `4`: redis secret.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of secret.
        """
        pulumi.set(__self__, "secret_name", secret_name)
        if additional_config is not None:
            pulumi.set(__self__, "additional_config", additional_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if recovery_window_in_days is not None:
            pulumi.set(__self__, "recovery_window_in_days", recovery_window_in_days)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Input[str]:
        """
        Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="additionalConfig")
    def additional_config(self) -> Optional[pulumi.Input[str]]:
        """
        Additional config for specific secret types in JSON string format.
        """
        return pulumi.get(self, "additional_config")

    @additional_config.setter
    def additional_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of secret. The maximum is 2048 bytes.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable secret. Default value is `true`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="recoveryWindowInDays")
    def recovery_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        """
        return pulumi.get(self, "recovery_window_in_days")

    @recovery_window_in_days.setter
    def recovery_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "recovery_window_in_days", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[int]]:
        """
        Type of secret. `0`: user-defined secret. `4`: redis secret.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secret_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of secret.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecretState:
    def __init__(__self__, *,
                 additional_config: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Secret resources.
        :param pulumi.Input[str] additional_config: Additional config for specific secret types in JSON string format.
        :param pulumi.Input[str] description: Description of secret. The maximum is 2048 bytes.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable secret. Default value is `true`.
        :param pulumi.Input[str] kms_key_id: KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        :param pulumi.Input[int] recovery_window_in_days: Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        :param pulumi.Input[str] secret_name: Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[int] secret_type: Type of secret. `0`: user-defined secret. `4`: redis secret.
        :param pulumi.Input[str] status: Status of secret.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of secret.
        """
        if additional_config is not None:
            pulumi.set(__self__, "additional_config", additional_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if recovery_window_in_days is not None:
            pulumi.set(__self__, "recovery_window_in_days", recovery_window_in_days)
        if secret_name is not None:
            pulumi.set(__self__, "secret_name", secret_name)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="additionalConfig")
    def additional_config(self) -> Optional[pulumi.Input[str]]:
        """
        Additional config for specific secret types in JSON string format.
        """
        return pulumi.get(self, "additional_config")

    @additional_config.setter
    def additional_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of secret. The maximum is 2048 bytes.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specify whether to enable secret. Default value is `true`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="recoveryWindowInDays")
    def recovery_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        """
        return pulumi.get(self, "recovery_window_in_days")

    @recovery_window_in_days.setter
    def recovery_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "recovery_window_in_days", value)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[int]]:
        """
        Type of secret. `0`: user-defined secret. `4`: redis secret.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secret_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of secret.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Tags of secret.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)


class Secret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_config: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Provide a resource to create a SSM secret.

        ## Example Usage
        ### Create user defined secret

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.ssm.Secret("foo",
            description="user defined secret",
            is_enabled=True,
            recovery_window_in_days=0,
            secret_name="test",
            tags={
                "test-tag": "test",
            })
        ```
        ### Create redis secret

        ```python
        import pulumi
        import json
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        instance = tencentcloud.Redis.get_instances(zone="ap-guangzhou-6")
        secret = tencentcloud.ssm.Secret("secret",
            secret_name="for-redis-test",
            description="redis secret",
            is_enabled=False,
            secret_type=4,
            additional_config=json.dumps({
                "Region": "ap-guangzhou",
                "Privilege": "r",
                "InstanceId": instance.instance_lists[0].redis_id,
                "ReadonlyPolicy": ["master"],
                "Remark": "for tf test",
            }),
            tags={
                "test-tag": "test",
            },
            recovery_window_in_days=0)
        ```

        ## Import

        SSM secret can be imported using the secretName, e.g.

        ```sh
         $ pulumi import tencentcloud:Ssm/secret:Secret foo test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] additional_config: Additional config for specific secret types in JSON string format.
        :param pulumi.Input[str] description: Description of secret. The maximum is 2048 bytes.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable secret. Default value is `true`.
        :param pulumi.Input[str] kms_key_id: KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        :param pulumi.Input[int] recovery_window_in_days: Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        :param pulumi.Input[str] secret_name: Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[int] secret_type: Type of secret. `0`: user-defined secret. `4`: redis secret.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of secret.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to create a SSM secret.

        ## Example Usage
        ### Create user defined secret

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.ssm.Secret("foo",
            description="user defined secret",
            is_enabled=True,
            recovery_window_in_days=0,
            secret_name="test",
            tags={
                "test-tag": "test",
            })
        ```
        ### Create redis secret

        ```python
        import pulumi
        import json
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        instance = tencentcloud.Redis.get_instances(zone="ap-guangzhou-6")
        secret = tencentcloud.ssm.Secret("secret",
            secret_name="for-redis-test",
            description="redis secret",
            is_enabled=False,
            secret_type=4,
            additional_config=json.dumps({
                "Region": "ap-guangzhou",
                "Privilege": "r",
                "InstanceId": instance.instance_lists[0].redis_id,
                "ReadonlyPolicy": ["master"],
                "Remark": "for tf test",
            }),
            tags={
                "test-tag": "test",
            },
            recovery_window_in_days=0)
        ```

        ## Import

        SSM secret can be imported using the secretName, e.g.

        ```sh
         $ pulumi import tencentcloud:Ssm/secret:Secret foo test
        ```

        :param str resource_name: The name of the resource.
        :param SecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_config: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 recovery_window_in_days: Optional[pulumi.Input[int]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[int]] = None,
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
            __props__ = SecretArgs.__new__(SecretArgs)

            __props__.__dict__["additional_config"] = additional_config
            __props__.__dict__["description"] = description
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["kms_key_id"] = kms_key_id
            __props__.__dict__["recovery_window_in_days"] = recovery_window_in_days
            if secret_name is None and not opts.urn:
                raise TypeError("Missing required property 'secret_name'")
            __props__.__dict__["secret_name"] = secret_name
            __props__.__dict__["secret_type"] = secret_type
            __props__.__dict__["tags"] = tags
            __props__.__dict__["status"] = None
        super(Secret, __self__).__init__(
            'tencentcloud:Ssm/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_config: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            recovery_window_in_days: Optional[pulumi.Input[int]] = None,
            secret_name: Optional[pulumi.Input[str]] = None,
            secret_type: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] additional_config: Additional config for specific secret types in JSON string format.
        :param pulumi.Input[str] description: Description of secret. The maximum is 2048 bytes.
        :param pulumi.Input[bool] is_enabled: Specify whether to enable secret. Default value is `true`.
        :param pulumi.Input[str] kms_key_id: KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        :param pulumi.Input[int] recovery_window_in_days: Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        :param pulumi.Input[str] secret_name: Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        :param pulumi.Input[int] secret_type: Type of secret. `0`: user-defined secret. `4`: redis secret.
        :param pulumi.Input[str] status: Status of secret.
        :param pulumi.Input[Mapping[str, Any]] tags: Tags of secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretState.__new__(_SecretState)

        __props__.__dict__["additional_config"] = additional_config
        __props__.__dict__["description"] = description
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["kms_key_id"] = kms_key_id
        __props__.__dict__["recovery_window_in_days"] = recovery_window_in_days
        __props__.__dict__["secret_name"] = secret_name
        __props__.__dict__["secret_type"] = secret_type
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalConfig")
    def additional_config(self) -> pulumi.Output[Optional[str]]:
        """
        Additional config for specific secret types in JSON string format.
        """
        return pulumi.get(self, "additional_config")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of secret. The maximum is 2048 bytes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specify whether to enable secret. Default value is `true`.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        """
        KMS keyId used to encrypt secret. If it is empty, it means that the CMK created by SSM for you by default is used for encryption. You can also specify the KMS CMK created by yourself in the same region for encryption.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="recoveryWindowInDays")
    def recovery_window_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        Specify the scheduled deletion date. Default value is `0` that means to delete immediately. 1-30 means the number of days reserved, completely deleted after this date.
        """
        return pulumi.get(self, "recovery_window_in_days")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Output[str]:
        """
        Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> pulumi.Output[int]:
        """
        Type of secret. `0`: user-defined secret. `4`: redis secret.
        """
        return pulumi.get(self, "secret_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of secret.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Tags of secret.
        """
        return pulumi.get(self, "tags")

