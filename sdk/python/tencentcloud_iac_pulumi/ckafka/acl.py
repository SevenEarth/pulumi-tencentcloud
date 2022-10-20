# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AclArgs', 'Acl']

@pulumi.input_type
class AclArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 operation_type: pulumi.Input[str],
                 resource_name: pulumi.Input[str],
                 host: Optional[pulumi.Input[str]] = None,
                 permission_type: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Acl resource.
        :param pulumi.Input[str] instance_id: ID of the ckafka instance.
        :param pulumi.Input[str] operation_type: ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        :param pulumi.Input[str] resource_name: ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        :param pulumi.Input[str] host: IP address allowed to access. The default value is `*`, which means that any host can access.
        :param pulumi.Input[str] permission_type: ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        :param pulumi.Input[str] principal: User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        :param pulumi.Input[str] resource_type: ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "operation_type", operation_type)
        pulumi.set(__self__, "resource_name", resource_name)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if permission_type is not None:
            pulumi.set(__self__, "permission_type", permission_type)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the ckafka instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="operationType")
    def operation_type(self) -> pulumi.Input[str]:
        """
        ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        """
        return pulumi.get(self, "operation_type")

    @operation_type.setter
    def operation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "operation_type", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Input[str]:
        """
        ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        IP address allowed to access. The default value is `*`, which means that any host can access.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="permissionType")
    def permission_type(self) -> Optional[pulumi.Input[str]]:
        """
        ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "permission_type")

    @permission_type.setter
    def permission_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_type", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


@pulumi.input_type
class _AclState:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 operation_type: Optional[pulumi.Input[str]] = None,
                 permission_type: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 resource_name: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Acl resources.
        :param pulumi.Input[str] host: IP address allowed to access. The default value is `*`, which means that any host can access.
        :param pulumi.Input[str] instance_id: ID of the ckafka instance.
        :param pulumi.Input[str] operation_type: ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        :param pulumi.Input[str] permission_type: ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        :param pulumi.Input[str] principal: User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        :param pulumi.Input[str] resource_name: ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        :param pulumi.Input[str] resource_type: ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        if host is not None:
            pulumi.set(__self__, "host", host)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if operation_type is not None:
            pulumi.set(__self__, "operation_type", operation_type)
        if permission_type is not None:
            pulumi.set(__self__, "permission_type", permission_type)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        IP address allowed to access. The default value is `*`, which means that any host can access.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the ckafka instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="operationType")
    def operation_type(self) -> Optional[pulumi.Input[str]]:
        """
        ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        """
        return pulumi.get(self, "operation_type")

    @operation_type.setter
    def operation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operation_type", value)

    @property
    @pulumi.getter(name="permissionType")
    def permission_type(self) -> Optional[pulumi.Input[str]]:
        """
        ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "permission_type")

    @permission_type.setter
    def permission_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_type", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[str]]:
        """
        ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


class Acl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 operation_type: Optional[pulumi.Input[str]] = None,
                 permission_type: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a Ckafka Acl.

        ## Example Usage

        Ckafka Acl

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.ckafka.Acl("foo",
            instance_id="ckafka-f9ife4zz",
            resource_type="TOPIC",
            resource_name_="topic-tf-test",
            operation_type="WRITE",
            permission_type="ALLOW",
            host="*",
            principal=tencentcloud_ckafka_user["foo"]["account_name"])
        ```

        ## Import

        Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_name, e.g.

        ```sh
         $ pulumi import tencentcloud:Ckafka/acl:Acl foo ckafka-f9ife4zz#ALLOW#test#*#WRITE#TOPIC#topic-tf-test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: IP address allowed to access. The default value is `*`, which means that any host can access.
        :param pulumi.Input[str] instance_id: ID of the ckafka instance.
        :param pulumi.Input[str] operation_type: ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        :param pulumi.Input[str] permission_type: ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        :param pulumi.Input[str] principal: User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        :param pulumi.Input[str] resource_name_: ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        :param pulumi.Input[str] resource_type: ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a Ckafka Acl.

        ## Example Usage

        Ckafka Acl

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.ckafka.Acl("foo",
            instance_id="ckafka-f9ife4zz",
            resource_type="TOPIC",
            resource_name_="topic-tf-test",
            operation_type="WRITE",
            permission_type="ALLOW",
            host="*",
            principal=tencentcloud_ckafka_user["foo"]["account_name"])
        ```

        ## Import

        Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_name, e.g.

        ```sh
         $ pulumi import tencentcloud:Ckafka/acl:Acl foo ckafka-f9ife4zz#ALLOW#test#*#WRITE#TOPIC#topic-tf-test
        ```

        :param str resource_name: The name of the resource.
        :param AclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 operation_type: Optional[pulumi.Input[str]] = None,
                 permission_type: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = AclArgs.__new__(AclArgs)

            __props__.__dict__["host"] = host
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if operation_type is None and not opts.urn:
                raise TypeError("Missing required property 'operation_type'")
            __props__.__dict__["operation_type"] = operation_type
            __props__.__dict__["permission_type"] = permission_type
            __props__.__dict__["principal"] = principal
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__.__dict__["resource_name"] = resource_name_
            __props__.__dict__["resource_type"] = resource_type
        super(Acl, __self__).__init__(
            'tencentcloud:Ckafka/acl:Acl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            host: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            operation_type: Optional[pulumi.Input[str]] = None,
            permission_type: Optional[pulumi.Input[str]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            resource_name_: Optional[pulumi.Input[str]] = None,
            resource_type: Optional[pulumi.Input[str]] = None) -> 'Acl':
        """
        Get an existing Acl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host: IP address allowed to access. The default value is `*`, which means that any host can access.
        :param pulumi.Input[str] instance_id: ID of the ckafka instance.
        :param pulumi.Input[str] operation_type: ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        :param pulumi.Input[str] permission_type: ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        :param pulumi.Input[str] principal: User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        :param pulumi.Input[str] resource_name_: ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        :param pulumi.Input[str] resource_type: ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AclState.__new__(_AclState)

        __props__.__dict__["host"] = host
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["operation_type"] = operation_type
        __props__.__dict__["permission_type"] = permission_type
        __props__.__dict__["principal"] = principal
        __props__.__dict__["resource_name"] = resource_name_
        __props__.__dict__["resource_type"] = resource_type
        return Acl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        IP address allowed to access. The default value is `*`, which means that any host can access.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the ckafka instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="operationType")
    def operation_type(self) -> pulumi.Output[str]:
        """
        ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
        """
        return pulumi.get(self, "operation_type")

    @property
    @pulumi.getter(name="permissionType")
    def permission_type(self) -> pulumi.Output[Optional[str]]:
        """
        ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "permission_type")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[Optional[str]]:
        """
        User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Output[str]:
        """
        ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[str]]:
        """
        ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
        """
        return pulumi.get(self, "resource_type")

