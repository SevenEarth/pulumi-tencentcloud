# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EncryptAttributesConfigArgs', 'EncryptAttributesConfig']

@pulumi.input_type
class EncryptAttributesConfigArgs:
    def __init__(__self__, *,
                 encrypt_enabled: pulumi.Input[int],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a EncryptAttributesConfig resource.
        :param pulumi.Input[int] encrypt_enabled: whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        :param pulumi.Input[str] instance_id: instance id.
        """
        pulumi.set(__self__, "encrypt_enabled", encrypt_enabled)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="encryptEnabled")
    def encrypt_enabled(self) -> pulumi.Input[int]:
        """
        whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        """
        return pulumi.get(self, "encrypt_enabled")

    @encrypt_enabled.setter
    def encrypt_enabled(self, value: pulumi.Input[int]):
        pulumi.set(self, "encrypt_enabled", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _EncryptAttributesConfigState:
    def __init__(__self__, *,
                 encrypt_enabled: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EncryptAttributesConfig resources.
        :param pulumi.Input[int] encrypt_enabled: whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        :param pulumi.Input[str] instance_id: instance id.
        """
        if encrypt_enabled is not None:
            pulumi.set(__self__, "encrypt_enabled", encrypt_enabled)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="encryptEnabled")
    def encrypt_enabled(self) -> Optional[pulumi.Input[int]]:
        """
        whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        """
        return pulumi.get(self, "encrypt_enabled")

    @encrypt_enabled.setter
    def encrypt_enabled(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "encrypt_enabled", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class EncryptAttributesConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encrypt_enabled: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a dcdb encrypt_attributes_config

        > **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        internal = tencentcloud.Security.get_groups(name="default")
        vpc = tencentcloud.Vpc.get_instances(name="Default-VPC")
        subnet = tencentcloud.Vpc.get_subnets(vpc_id=vpc.instance_lists[0].vpc_id)
        vpc_id = subnet.instance_lists[0].vpc_id
        subnet_id = subnet.instance_lists[0].subnet_id
        sg_id = internal.security_groups[0].security_group_id
        prepaid_instance = tencentcloud.dcdb.DbInstance("prepaidInstance",
            instance_name="test_dcdb_db_post_instance",
            zones=[var["default_az"]],
            period=1,
            shard_memory=2,
            shard_storage=10,
            shard_node_count=2,
            shard_count=2,
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            db_version_id="8.0",
            resource_tags=[tencentcloud.dcdb.DbInstanceResourceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )],
            security_group_ids=[sg_id])
        hourdb_instance = tencentcloud.dcdb.HourdbInstance("hourdbInstance",
            instance_name="test_dcdb_db_hourdb_instance",
            zones=[var["default_az"]],
            shard_memory=2,
            shard_storage=10,
            shard_node_count=2,
            shard_count=2,
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            security_group_id=sg_id,
            db_version_id="8.0",
            resource_tags=[tencentcloud.dcdb.HourdbInstanceResourceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        prepaid_dcdb_id = prepaid_instance.id
        hourdb_dcdb_id = hourdb_instance.id
        # for postpaid instance
        config_hourdb = tencentcloud.dcdb.EncryptAttributesConfig("configHourdb",
            instance_id=hourdb_dcdb_id,
            encrypt_enabled=1)
        # for prepaid instance
        config_prepaid = tencentcloud.dcdb.EncryptAttributesConfig("configPrepaid",
            instance_id=prepaid_dcdb_id,
            encrypt_enabled=1)
        ```

        ## Import

        dcdb encrypt_attributes_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig encrypt_attributes_config encrypt_attributes_config_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] encrypt_enabled: whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        :param pulumi.Input[str] instance_id: instance id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EncryptAttributesConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a dcdb encrypt_attributes_config

        > **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        internal = tencentcloud.Security.get_groups(name="default")
        vpc = tencentcloud.Vpc.get_instances(name="Default-VPC")
        subnet = tencentcloud.Vpc.get_subnets(vpc_id=vpc.instance_lists[0].vpc_id)
        vpc_id = subnet.instance_lists[0].vpc_id
        subnet_id = subnet.instance_lists[0].subnet_id
        sg_id = internal.security_groups[0].security_group_id
        prepaid_instance = tencentcloud.dcdb.DbInstance("prepaidInstance",
            instance_name="test_dcdb_db_post_instance",
            zones=[var["default_az"]],
            period=1,
            shard_memory=2,
            shard_storage=10,
            shard_node_count=2,
            shard_count=2,
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            db_version_id="8.0",
            resource_tags=[tencentcloud.dcdb.DbInstanceResourceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )],
            security_group_ids=[sg_id])
        hourdb_instance = tencentcloud.dcdb.HourdbInstance("hourdbInstance",
            instance_name="test_dcdb_db_hourdb_instance",
            zones=[var["default_az"]],
            shard_memory=2,
            shard_storage=10,
            shard_node_count=2,
            shard_count=2,
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            security_group_id=sg_id,
            db_version_id="8.0",
            resource_tags=[tencentcloud.dcdb.HourdbInstanceResourceTagArgs(
                tag_key="aaa",
                tag_value="bbb",
            )])
        prepaid_dcdb_id = prepaid_instance.id
        hourdb_dcdb_id = hourdb_instance.id
        # for postpaid instance
        config_hourdb = tencentcloud.dcdb.EncryptAttributesConfig("configHourdb",
            instance_id=hourdb_dcdb_id,
            encrypt_enabled=1)
        # for prepaid instance
        config_prepaid = tencentcloud.dcdb.EncryptAttributesConfig("configPrepaid",
            instance_id=prepaid_dcdb_id,
            encrypt_enabled=1)
        ```

        ## Import

        dcdb encrypt_attributes_config can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig encrypt_attributes_config encrypt_attributes_config_id
        ```

        :param str resource_name: The name of the resource.
        :param EncryptAttributesConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EncryptAttributesConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encrypt_enabled: Optional[pulumi.Input[int]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = EncryptAttributesConfigArgs.__new__(EncryptAttributesConfigArgs)

            if encrypt_enabled is None and not opts.urn:
                raise TypeError("Missing required property 'encrypt_enabled'")
            __props__.__dict__["encrypt_enabled"] = encrypt_enabled
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(EncryptAttributesConfig, __self__).__init__(
            'tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            encrypt_enabled: Optional[pulumi.Input[int]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'EncryptAttributesConfig':
        """
        Get an existing EncryptAttributesConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] encrypt_enabled: whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        :param pulumi.Input[str] instance_id: instance id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EncryptAttributesConfigState.__new__(_EncryptAttributesConfigState)

        __props__.__dict__["encrypt_enabled"] = encrypt_enabled
        __props__.__dict__["instance_id"] = instance_id
        return EncryptAttributesConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="encryptEnabled")
    def encrypt_enabled(self) -> pulumi.Output[int]:
        """
        whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        """
        return pulumi.get(self, "encrypt_enabled")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        instance id.
        """
        return pulumi.get(self, "instance_id")

