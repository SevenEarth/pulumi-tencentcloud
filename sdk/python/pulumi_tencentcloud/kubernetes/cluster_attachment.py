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

__all__ = ['ClusterAttachmentArgs', 'ClusterAttachment']

@pulumi.input_type
class ClusterAttachmentArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 hostname: Optional[pulumi.Input[str]] = None,
                 key_ids: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 unschedulable: Optional[pulumi.Input[int]] = None,
                 worker_config: Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']] = None,
                 worker_config_overrides: Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']] = None):
        """
        The set of arguments for constructing a ClusterAttachment resource.
        :param pulumi.Input[str] cluster_id: ID of the cluster.
        :param pulumi.Input[str] instance_id: ID of the CVM instance, this cvm will reinstall the system.
        :param pulumi.Input[str] hostname: The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        :param pulumi.Input[str] key_ids: The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        :param pulumi.Input[Mapping[str, Any]] labels: Labels of tke attachment exits CVM.
        :param pulumi.Input[str] password: Password to access, should be set if `key_ids` not set.
        :param pulumi.Input[int] unschedulable: Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        :param pulumi.Input['ClusterAttachmentWorkerConfigArgs'] worker_config: Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        :param pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs'] worker_config_overrides: Override variable worker_config, commonly used to attach existing instances.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if key_ids is not None:
            pulumi.set(__self__, "key_ids", key_ids)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if unschedulable is not None:
            pulumi.set(__self__, "unschedulable", unschedulable)
        if worker_config is not None:
            pulumi.set(__self__, "worker_config", worker_config)
        if worker_config_overrides is not None:
            pulumi.set(__self__, "worker_config_overrides", worker_config_overrides)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        ID of the CVM instance, this cvm will reinstall the system.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="keyIds")
    def key_ids(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        """
        return pulumi.get(self, "key_ids")

    @key_ids.setter
    def key_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_ids", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Labels of tke attachment exits CVM.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password to access, should be set if `key_ids` not set.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def unschedulable(self) -> Optional[pulumi.Input[int]]:
        """
        Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        """
        return pulumi.get(self, "unschedulable")

    @unschedulable.setter
    def unschedulable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unschedulable", value)

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']]:
        """
        Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config")

    @worker_config.setter
    def worker_config(self, value: Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']]):
        pulumi.set(self, "worker_config", value)

    @property
    @pulumi.getter(name="workerConfigOverrides")
    def worker_config_overrides(self) -> Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']]:
        """
        Override variable worker_config, commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config_overrides")

    @worker_config_overrides.setter
    def worker_config_overrides(self, value: Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']]):
        pulumi.set(self, "worker_config_overrides", value)


@pulumi.input_type
class _ClusterAttachmentState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_ids: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 unschedulable: Optional[pulumi.Input[int]] = None,
                 worker_config: Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']] = None,
                 worker_config_overrides: Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']] = None):
        """
        Input properties used for looking up and filtering ClusterAttachment resources.
        :param pulumi.Input[str] cluster_id: ID of the cluster.
        :param pulumi.Input[str] hostname: The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        :param pulumi.Input[str] instance_id: ID of the CVM instance, this cvm will reinstall the system.
        :param pulumi.Input[str] key_ids: The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        :param pulumi.Input[Mapping[str, Any]] labels: Labels of tke attachment exits CVM.
        :param pulumi.Input[str] password: Password to access, should be set if `key_ids` not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of security group IDs after attach to cluster.
        :param pulumi.Input[str] state: State of the node.
        :param pulumi.Input[int] unschedulable: Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        :param pulumi.Input['ClusterAttachmentWorkerConfigArgs'] worker_config: Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        :param pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs'] worker_config_overrides: Override variable worker_config, commonly used to attach existing instances.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if key_ids is not None:
            pulumi.set(__self__, "key_ids", key_ids)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if unschedulable is not None:
            pulumi.set(__self__, "unschedulable", unschedulable)
        if worker_config is not None:
            pulumi.set(__self__, "worker_config", worker_config)
        if worker_config_overrides is not None:
            pulumi.set(__self__, "worker_config_overrides", worker_config_overrides)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the CVM instance, this cvm will reinstall the system.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="keyIds")
    def key_ids(self) -> Optional[pulumi.Input[str]]:
        """
        The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        """
        return pulumi.get(self, "key_ids")

    @key_ids.setter
    def key_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_ids", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Labels of tke attachment exits CVM.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password to access, should be set if `key_ids` not set.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of security group IDs after attach to cluster.
        """
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the node.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def unschedulable(self) -> Optional[pulumi.Input[int]]:
        """
        Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        """
        return pulumi.get(self, "unschedulable")

    @unschedulable.setter
    def unschedulable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unschedulable", value)

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']]:
        """
        Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config")

    @worker_config.setter
    def worker_config(self, value: Optional[pulumi.Input['ClusterAttachmentWorkerConfigArgs']]):
        pulumi.set(self, "worker_config", value)

    @property
    @pulumi.getter(name="workerConfigOverrides")
    def worker_config_overrides(self) -> Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']]:
        """
        Override variable worker_config, commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config_overrides")

    @worker_config_overrides.setter
    def worker_config_overrides(self, value: Optional[pulumi.Input['ClusterAttachmentWorkerConfigOverridesArgs']]):
        pulumi.set(self, "worker_config_overrides", value)


class ClusterAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_ids: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 unschedulable: Optional[pulumi.Input[int]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigArgs']]] = None,
                 worker_config_overrides: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigOverridesArgs']]] = None,
                 __props__=None):
        """
        Provide a resource to attach an existing  cvm to kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-3"
        cluster_cidr = config.get("clusterCidr")
        if cluster_cidr is None:
            cluster_cidr = "172.16.0.0/16"
        default_instance_type = config.get("defaultInstanceType")
        if default_instance_type is None:
            default_instance_type = "S1.SMALL1"
        default_instance = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="centos")
        vpc = tencentcloud.Vpc.get_subnets(is_default=True,
            availability_zone=availability_zone)
        default_types = tencentcloud.Instance.get_types(filters=[tencentcloud.instance.GetTypesFilterArgs(
                name="instance-family",
                values=["SA2"],
            )],
            cpu_core_count=8,
            memory_size=16)
        foo = tencentcloud.instance.Instance("foo",
            instance_name="tf-auto-test-1-1",
            availability_zone=availability_zone,
            image_id=default_instance.images[0].image_id,
            instance_type=default_instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50)
        managed_cluster = tencentcloud.kubernetes.Cluster("managedCluster",
            vpc_id=vpc.instance_lists[0].vpc_id,
            cluster_cidr="10.1.0.0/16",
            cluster_max_pod_num=32,
            cluster_name="keep",
            cluster_desc="test cluster desc",
            cluster_max_service_num=32,
            worker_configs=[tencentcloud.kubernetes.ClusterWorkerConfigArgs(
                count=1,
                availability_zone=availability_zone,
                instance_type=default_instance_type,
                system_disk_type="CLOUD_SSD",
                system_disk_size=60,
                internet_charge_type="TRAFFIC_POSTPAID_BY_HOUR",
                internet_max_bandwidth_out=100,
                public_ip_assigned=True,
                subnet_id=vpc.instance_lists[0].subnet_id,
                data_disks=[tencentcloud.kubernetes.ClusterWorkerConfigDataDiskArgs(
                    disk_type="CLOUD_PREMIUM",
                    disk_size=50,
                )],
                enhanced_security_service=False,
                enhanced_monitor_service=False,
                user_data="dGVzdA==",
                password="ZZXXccvv1212",
            )],
            cluster_deploy_type="MANAGED_CLUSTER")
        test_attach = tencentcloud.kubernetes.ClusterAttachment("testAttach",
            cluster_id=managed_cluster.id,
            instance_id=foo.id,
            password="Lo4wbdit",
            labels={
                "test1": "test1",
                "test2": "test2",
            },
            worker_config_overrides=tencentcloud.kubernetes.ClusterAttachmentWorkerConfigOverridesArgs(
                desired_pod_num=8,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: ID of the cluster.
        :param pulumi.Input[str] hostname: The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        :param pulumi.Input[str] instance_id: ID of the CVM instance, this cvm will reinstall the system.
        :param pulumi.Input[str] key_ids: The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        :param pulumi.Input[Mapping[str, Any]] labels: Labels of tke attachment exits CVM.
        :param pulumi.Input[str] password: Password to access, should be set if `key_ids` not set.
        :param pulumi.Input[int] unschedulable: Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        :param pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigArgs']] worker_config: Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        :param pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigOverridesArgs']] worker_config_overrides: Override variable worker_config, commonly used to attach existing instances.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provide a resource to attach an existing  cvm to kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-3"
        cluster_cidr = config.get("clusterCidr")
        if cluster_cidr is None:
            cluster_cidr = "172.16.0.0/16"
        default_instance_type = config.get("defaultInstanceType")
        if default_instance_type is None:
            default_instance_type = "S1.SMALL1"
        default_instance = tencentcloud.Images.get_instance(image_types=["PUBLIC_IMAGE"],
            os_name="centos")
        vpc = tencentcloud.Vpc.get_subnets(is_default=True,
            availability_zone=availability_zone)
        default_types = tencentcloud.Instance.get_types(filters=[tencentcloud.instance.GetTypesFilterArgs(
                name="instance-family",
                values=["SA2"],
            )],
            cpu_core_count=8,
            memory_size=16)
        foo = tencentcloud.instance.Instance("foo",
            instance_name="tf-auto-test-1-1",
            availability_zone=availability_zone,
            image_id=default_instance.images[0].image_id,
            instance_type=default_instance_type,
            system_disk_type="CLOUD_PREMIUM",
            system_disk_size=50)
        managed_cluster = tencentcloud.kubernetes.Cluster("managedCluster",
            vpc_id=vpc.instance_lists[0].vpc_id,
            cluster_cidr="10.1.0.0/16",
            cluster_max_pod_num=32,
            cluster_name="keep",
            cluster_desc="test cluster desc",
            cluster_max_service_num=32,
            worker_configs=[tencentcloud.kubernetes.ClusterWorkerConfigArgs(
                count=1,
                availability_zone=availability_zone,
                instance_type=default_instance_type,
                system_disk_type="CLOUD_SSD",
                system_disk_size=60,
                internet_charge_type="TRAFFIC_POSTPAID_BY_HOUR",
                internet_max_bandwidth_out=100,
                public_ip_assigned=True,
                subnet_id=vpc.instance_lists[0].subnet_id,
                data_disks=[tencentcloud.kubernetes.ClusterWorkerConfigDataDiskArgs(
                    disk_type="CLOUD_PREMIUM",
                    disk_size=50,
                )],
                enhanced_security_service=False,
                enhanced_monitor_service=False,
                user_data="dGVzdA==",
                password="ZZXXccvv1212",
            )],
            cluster_deploy_type="MANAGED_CLUSTER")
        test_attach = tencentcloud.kubernetes.ClusterAttachment("testAttach",
            cluster_id=managed_cluster.id,
            instance_id=foo.id,
            password="Lo4wbdit",
            labels={
                "test1": "test1",
                "test2": "test2",
            },
            worker_config_overrides=tencentcloud.kubernetes.ClusterAttachmentWorkerConfigOverridesArgs(
                desired_pod_num=8,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param ClusterAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 key_ids: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 unschedulable: Optional[pulumi.Input[int]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigArgs']]] = None,
                 worker_config_overrides: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigOverridesArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterAttachmentArgs.__new__(ClusterAttachmentArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["hostname"] = hostname
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["key_ids"] = key_ids
            __props__.__dict__["labels"] = labels
            __props__.__dict__["password"] = password
            __props__.__dict__["unschedulable"] = unschedulable
            __props__.__dict__["worker_config"] = worker_config
            __props__.__dict__["worker_config_overrides"] = worker_config_overrides
            __props__.__dict__["security_groups"] = None
            __props__.__dict__["state"] = None
        super(ClusterAttachment, __self__).__init__(
            'tencentcloud:Kubernetes/clusterAttachment:ClusterAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            key_ids: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            password: Optional[pulumi.Input[str]] = None,
            security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            state: Optional[pulumi.Input[str]] = None,
            unschedulable: Optional[pulumi.Input[int]] = None,
            worker_config: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigArgs']]] = None,
            worker_config_overrides: Optional[pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigOverridesArgs']]] = None) -> 'ClusterAttachment':
        """
        Get an existing ClusterAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: ID of the cluster.
        :param pulumi.Input[str] hostname: The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        :param pulumi.Input[str] instance_id: ID of the CVM instance, this cvm will reinstall the system.
        :param pulumi.Input[str] key_ids: The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        :param pulumi.Input[Mapping[str, Any]] labels: Labels of tke attachment exits CVM.
        :param pulumi.Input[str] password: Password to access, should be set if `key_ids` not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: A list of security group IDs after attach to cluster.
        :param pulumi.Input[str] state: State of the node.
        :param pulumi.Input[int] unschedulable: Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        :param pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigArgs']] worker_config: Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        :param pulumi.Input[pulumi.InputType['ClusterAttachmentWorkerConfigOverridesArgs']] worker_config_overrides: Override variable worker_config, commonly used to attach existing instances.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterAttachmentState.__new__(_ClusterAttachmentState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["key_ids"] = key_ids
        __props__.__dict__["labels"] = labels
        __props__.__dict__["password"] = password
        __props__.__dict__["security_groups"] = security_groups
        __props__.__dict__["state"] = state
        __props__.__dict__["unschedulable"] = unschedulable
        __props__.__dict__["worker_config"] = worker_config
        __props__.__dict__["worker_config_overrides"] = worker_config_overrides
        return ClusterAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        ID of the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[Optional[str]]:
        """
        The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        ID of the CVM instance, this cvm will reinstall the system.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="keyIds")
    def key_ids(self) -> pulumi.Output[Optional[str]]:
        """
        The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
        """
        return pulumi.get(self, "key_ids")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Labels of tke attachment exits CVM.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Password to access, should be set if `key_ids` not set.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of security group IDs after attach to cluster.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the node.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def unschedulable(self) -> pulumi.Output[Optional[int]]:
        """
        Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        """
        return pulumi.get(self, "unschedulable")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> pulumi.Output[Optional['outputs.ClusterAttachmentWorkerConfig']]:
        """
        Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config")

    @property
    @pulumi.getter(name="workerConfigOverrides")
    def worker_config_overrides(self) -> pulumi.Output[Optional['outputs.ClusterAttachmentWorkerConfigOverrides']]:
        """
        Override variable worker_config, commonly used to attach existing instances.
        """
        return pulumi.get(self, "worker_config_overrides")

