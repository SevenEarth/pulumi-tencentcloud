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

__all__ = ['SnapshotPolicyArgs', 'SnapshotPolicy']

@pulumi.input_type
class SnapshotPolicyArgs:
    def __init__(__self__, *,
                 backup_type: pulumi.Input[str],
                 cos_bucket: pulumi.Input[str],
                 cos_region: pulumi.Input[str],
                 create_new_cos: pulumi.Input[bool],
                 keep_time: pulumi.Input[int],
                 snapshot_policy_name: pulumi.Input[str],
                 backup_policies: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]] = None):
        """
        The set of arguments for constructing a SnapshotPolicy resource.
        :param pulumi.Input[str] backup_type: Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        :param pulumi.Input[str] cos_bucket: cos bucket.
        :param pulumi.Input[str] cos_region: The region where the cos bucket is located.
        :param pulumi.Input[bool] create_new_cos: Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[int] keep_time: The retention time supports 1 to 365 days.
        :param pulumi.Input[str] snapshot_policy_name: Snapshot policy name.
        :param pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]] backup_policies: Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        """
        pulumi.set(__self__, "backup_type", backup_type)
        pulumi.set(__self__, "cos_bucket", cos_bucket)
        pulumi.set(__self__, "cos_region", cos_region)
        pulumi.set(__self__, "create_new_cos", create_new_cos)
        pulumi.set(__self__, "keep_time", keep_time)
        pulumi.set(__self__, "snapshot_policy_name", snapshot_policy_name)
        if backup_policies is not None:
            pulumi.set(__self__, "backup_policies", backup_policies)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Input[str]:
        """
        Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="cosBucket")
    def cos_bucket(self) -> pulumi.Input[str]:
        """
        cos bucket.
        """
        return pulumi.get(self, "cos_bucket")

    @cos_bucket.setter
    def cos_bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "cos_bucket", value)

    @property
    @pulumi.getter(name="cosRegion")
    def cos_region(self) -> pulumi.Input[str]:
        """
        The region where the cos bucket is located.
        """
        return pulumi.get(self, "cos_region")

    @cos_region.setter
    def cos_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "cos_region", value)

    @property
    @pulumi.getter(name="createNewCos")
    def create_new_cos(self) -> pulumi.Input[bool]:
        """
        Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "create_new_cos")

    @create_new_cos.setter
    def create_new_cos(self, value: pulumi.Input[bool]):
        pulumi.set(self, "create_new_cos", value)

    @property
    @pulumi.getter(name="keepTime")
    def keep_time(self) -> pulumi.Input[int]:
        """
        The retention time supports 1 to 365 days.
        """
        return pulumi.get(self, "keep_time")

    @keep_time.setter
    def keep_time(self, value: pulumi.Input[int]):
        pulumi.set(self, "keep_time", value)

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> pulumi.Input[str]:
        """
        Snapshot policy name.
        """
        return pulumi.get(self, "snapshot_policy_name")

    @snapshot_policy_name.setter
    def snapshot_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "snapshot_policy_name", value)

    @property
    @pulumi.getter(name="backupPolicies")
    def backup_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]]:
        """
        Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "backup_policies")

    @backup_policies.setter
    def backup_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]]):
        pulumi.set(self, "backup_policies", value)


@pulumi.input_type
class _SnapshotPolicyState:
    def __init__(__self__, *,
                 backup_policies: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cos_bucket: Optional[pulumi.Input[str]] = None,
                 cos_region: Optional[pulumi.Input[str]] = None,
                 create_new_cos: Optional[pulumi.Input[bool]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 keep_time: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnapshotPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]] backup_policies: Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[str] backup_type: Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        :param pulumi.Input[str] cos_bucket: cos bucket.
        :param pulumi.Input[str] cos_region: The region where the cos bucket is located.
        :param pulumi.Input[bool] create_new_cos: Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[str] create_time: Creation time.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[bool] enable: Enabled state, True-enabled, False-disabled, the default is True.
        :param pulumi.Input[int] keep_time: The retention time supports 1 to 365 days.
        :param pulumi.Input[str] snapshot_policy_id: Snapshot policy Id.
        :param pulumi.Input[str] snapshot_policy_name: Snapshot policy name.
        """
        if backup_policies is not None:
            pulumi.set(__self__, "backup_policies", backup_policies)
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if cos_bucket is not None:
            pulumi.set(__self__, "cos_bucket", cos_bucket)
        if cos_region is not None:
            pulumi.set(__self__, "cos_region", cos_region)
        if create_new_cos is not None:
            pulumi.set(__self__, "create_new_cos", create_new_cos)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if keep_time is not None:
            pulumi.set(__self__, "keep_time", keep_time)
        if snapshot_policy_id is not None:
            pulumi.set(__self__, "snapshot_policy_id", snapshot_policy_id)
        if snapshot_policy_name is not None:
            pulumi.set(__self__, "snapshot_policy_name", snapshot_policy_name)

    @property
    @pulumi.getter(name="backupPolicies")
    def backup_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]]:
        """
        Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "backup_policies")

    @backup_policies.setter
    def backup_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotPolicyBackupPolicyArgs']]]]):
        pulumi.set(self, "backup_policies", value)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="cosBucket")
    def cos_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        cos bucket.
        """
        return pulumi.get(self, "cos_bucket")

    @cos_bucket.setter
    def cos_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cos_bucket", value)

    @property
    @pulumi.getter(name="cosRegion")
    def cos_region(self) -> Optional[pulumi.Input[str]]:
        """
        The region where the cos bucket is located.
        """
        return pulumi.get(self, "cos_region")

    @cos_region.setter
    def cos_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cos_region", value)

    @property
    @pulumi.getter(name="createNewCos")
    def create_new_cos(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "create_new_cos")

    @create_new_cos.setter
    def create_new_cos(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_new_cos", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time.Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabled state, True-enabled, False-disabled, the default is True.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="keepTime")
    def keep_time(self) -> Optional[pulumi.Input[int]]:
        """
        The retention time supports 1 to 365 days.
        """
        return pulumi.get(self, "keep_time")

    @keep_time.setter
    def keep_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keep_time", value)

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        Snapshot policy Id.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @snapshot_policy_id.setter
    def snapshot_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_policy_id", value)

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        Snapshot policy name.
        """
        return pulumi.get(self, "snapshot_policy_name")

    @snapshot_policy_name.setter
    def snapshot_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_policy_name", value)


class SnapshotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotPolicyBackupPolicyArgs']]]]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cos_bucket: Optional[pulumi.Input[str]] = None,
                 cos_region: Optional[pulumi.Input[str]] = None,
                 create_new_cos: Optional[pulumi.Input[bool]] = None,
                 keep_time: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a vpc snapshot_policy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_bucket = tencentcloud.cos.Bucket("exampleBucket",
            bucket="tf-example-1308919341",
            acl="private")
        example_snapshot_policy = tencentcloud.vpc.SnapshotPolicy("exampleSnapshotPolicy",
            snapshot_policy_name="tf-example",
            backup_type="time",
            cos_bucket=example_bucket.bucket,
            cos_region="ap-guangzhou",
            create_new_cos=False,
            keep_time=2,
            backup_policies=[
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="monday",
                    backup_time="00:00:00",
                ),
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="tuesday",
                    backup_time="01:00:00",
                ),
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="wednesday",
                    backup_time="02:00:00",
                ),
            ])
        ```

        ## Import

        vpc snapshot_policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/snapshotPolicy:SnapshotPolicy snapshot_policy snapshot_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotPolicyBackupPolicyArgs']]]] backup_policies: Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[str] backup_type: Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        :param pulumi.Input[str] cos_bucket: cos bucket.
        :param pulumi.Input[str] cos_region: The region where the cos bucket is located.
        :param pulumi.Input[bool] create_new_cos: Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[int] keep_time: The retention time supports 1 to 365 days.
        :param pulumi.Input[str] snapshot_policy_name: Snapshot policy name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a vpc snapshot_policy

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        example_bucket = tencentcloud.cos.Bucket("exampleBucket",
            bucket="tf-example-1308919341",
            acl="private")
        example_snapshot_policy = tencentcloud.vpc.SnapshotPolicy("exampleSnapshotPolicy",
            snapshot_policy_name="tf-example",
            backup_type="time",
            cos_bucket=example_bucket.bucket,
            cos_region="ap-guangzhou",
            create_new_cos=False,
            keep_time=2,
            backup_policies=[
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="monday",
                    backup_time="00:00:00",
                ),
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="tuesday",
                    backup_time="01:00:00",
                ),
                tencentcloud.vpc.SnapshotPolicyBackupPolicyArgs(
                    backup_day="wednesday",
                    backup_time="02:00:00",
                ),
            ])
        ```

        ## Import

        vpc snapshot_policy can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Vpc/snapshotPolicy:SnapshotPolicy snapshot_policy snapshot_policy_id
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotPolicyBackupPolicyArgs']]]]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 cos_bucket: Optional[pulumi.Input[str]] = None,
                 cos_region: Optional[pulumi.Input[str]] = None,
                 create_new_cos: Optional[pulumi.Input[bool]] = None,
                 keep_time: Optional[pulumi.Input[int]] = None,
                 snapshot_policy_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = SnapshotPolicyArgs.__new__(SnapshotPolicyArgs)

            __props__.__dict__["backup_policies"] = backup_policies
            if backup_type is None and not opts.urn:
                raise TypeError("Missing required property 'backup_type'")
            __props__.__dict__["backup_type"] = backup_type
            if cos_bucket is None and not opts.urn:
                raise TypeError("Missing required property 'cos_bucket'")
            __props__.__dict__["cos_bucket"] = cos_bucket
            if cos_region is None and not opts.urn:
                raise TypeError("Missing required property 'cos_region'")
            __props__.__dict__["cos_region"] = cos_region
            if create_new_cos is None and not opts.urn:
                raise TypeError("Missing required property 'create_new_cos'")
            __props__.__dict__["create_new_cos"] = create_new_cos
            if keep_time is None and not opts.urn:
                raise TypeError("Missing required property 'keep_time'")
            __props__.__dict__["keep_time"] = keep_time
            if snapshot_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'snapshot_policy_name'")
            __props__.__dict__["snapshot_policy_name"] = snapshot_policy_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["enable"] = None
            __props__.__dict__["snapshot_policy_id"] = None
        super(SnapshotPolicy, __self__).__init__(
            'tencentcloud:Vpc/snapshotPolicy:SnapshotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotPolicyBackupPolicyArgs']]]]] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            cos_bucket: Optional[pulumi.Input[str]] = None,
            cos_region: Optional[pulumi.Input[str]] = None,
            create_new_cos: Optional[pulumi.Input[bool]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            keep_time: Optional[pulumi.Input[int]] = None,
            snapshot_policy_id: Optional[pulumi.Input[str]] = None,
            snapshot_policy_name: Optional[pulumi.Input[str]] = None) -> 'SnapshotPolicy':
        """
        Get an existing SnapshotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotPolicyBackupPolicyArgs']]]] backup_policies: Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[str] backup_type: Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        :param pulumi.Input[str] cos_bucket: cos bucket.
        :param pulumi.Input[str] cos_region: The region where the cos bucket is located.
        :param pulumi.Input[bool] create_new_cos: Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[str] create_time: Creation time.Note: This field may return null, indicating that no valid value can be obtained.
        :param pulumi.Input[bool] enable: Enabled state, True-enabled, False-disabled, the default is True.
        :param pulumi.Input[int] keep_time: The retention time supports 1 to 365 days.
        :param pulumi.Input[str] snapshot_policy_id: Snapshot policy Id.
        :param pulumi.Input[str] snapshot_policy_name: Snapshot policy name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotPolicyState.__new__(_SnapshotPolicyState)

        __props__.__dict__["backup_policies"] = backup_policies
        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["cos_bucket"] = cos_bucket
        __props__.__dict__["cos_region"] = cos_region
        __props__.__dict__["create_new_cos"] = create_new_cos
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["enable"] = enable
        __props__.__dict__["keep_time"] = keep_time
        __props__.__dict__["snapshot_policy_id"] = snapshot_policy_id
        __props__.__dict__["snapshot_policy_name"] = snapshot_policy_name
        return SnapshotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPolicies")
    def backup_policies(self) -> pulumi.Output[Optional[Sequence['outputs.SnapshotPolicyBackupPolicy']]]:
        """
        Time backup strategy. Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "backup_policies")

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        Backup strategy type, `operate`: operate backup, `time`: schedule backup.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="cosBucket")
    def cos_bucket(self) -> pulumi.Output[str]:
        """
        cos bucket.
        """
        return pulumi.get(self, "cos_bucket")

    @property
    @pulumi.getter(name="cosRegion")
    def cos_region(self) -> pulumi.Output[str]:
        """
        The region where the cos bucket is located.
        """
        return pulumi.get(self, "cos_region")

    @property
    @pulumi.getter(name="createNewCos")
    def create_new_cos(self) -> pulumi.Output[bool]:
        """
        Whether to create a new cos bucket, the default is False.Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "create_new_cos")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time.Note: This field may return null, indicating that no valid value can be obtained.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[bool]:
        """
        Enabled state, True-enabled, False-disabled, the default is True.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="keepTime")
    def keep_time(self) -> pulumi.Output[int]:
        """
        The retention time supports 1 to 365 days.
        """
        return pulumi.get(self, "keep_time")

    @property
    @pulumi.getter(name="snapshotPolicyId")
    def snapshot_policy_id(self) -> pulumi.Output[str]:
        """
        Snapshot policy Id.
        """
        return pulumi.get(self, "snapshot_policy_id")

    @property
    @pulumi.getter(name="snapshotPolicyName")
    def snapshot_policy_name(self) -> pulumi.Output[str]:
        """
        Snapshot policy name.
        """
        return pulumi.get(self, "snapshot_policy_name")

