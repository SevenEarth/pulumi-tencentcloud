# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPublishSubscribesResult',
    'AwaitableGetPublishSubscribesResult',
    'get_publish_subscribes',
    'get_publish_subscribes_output',
]

@pulumi.output_type
class GetPublishSubscribesResult:
    """
    A collection of values returned by getPublishSubscribes.
    """
    def __init__(__self__, id=None, instance_id=None, pub_or_sub_instance_id=None, pub_or_sub_instance_ip=None, publish_database=None, publish_subscribe_id=None, publish_subscribe_lists=None, publish_subscribe_name=None, result_output_file=None, subscribe_database=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if pub_or_sub_instance_id and not isinstance(pub_or_sub_instance_id, str):
            raise TypeError("Expected argument 'pub_or_sub_instance_id' to be a str")
        pulumi.set(__self__, "pub_or_sub_instance_id", pub_or_sub_instance_id)
        if pub_or_sub_instance_ip and not isinstance(pub_or_sub_instance_ip, str):
            raise TypeError("Expected argument 'pub_or_sub_instance_ip' to be a str")
        pulumi.set(__self__, "pub_or_sub_instance_ip", pub_or_sub_instance_ip)
        if publish_database and not isinstance(publish_database, str):
            raise TypeError("Expected argument 'publish_database' to be a str")
        pulumi.set(__self__, "publish_database", publish_database)
        if publish_subscribe_id and not isinstance(publish_subscribe_id, int):
            raise TypeError("Expected argument 'publish_subscribe_id' to be a int")
        pulumi.set(__self__, "publish_subscribe_id", publish_subscribe_id)
        if publish_subscribe_lists and not isinstance(publish_subscribe_lists, list):
            raise TypeError("Expected argument 'publish_subscribe_lists' to be a list")
        pulumi.set(__self__, "publish_subscribe_lists", publish_subscribe_lists)
        if publish_subscribe_name and not isinstance(publish_subscribe_name, str):
            raise TypeError("Expected argument 'publish_subscribe_name' to be a str")
        pulumi.set(__self__, "publish_subscribe_name", publish_subscribe_name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if subscribe_database and not isinstance(subscribe_database, str):
            raise TypeError("Expected argument 'subscribe_database' to be a str")
        pulumi.set(__self__, "subscribe_database", subscribe_database)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="pubOrSubInstanceId")
    def pub_or_sub_instance_id(self) -> Optional[str]:
        return pulumi.get(self, "pub_or_sub_instance_id")

    @property
    @pulumi.getter(name="pubOrSubInstanceIp")
    def pub_or_sub_instance_ip(self) -> Optional[str]:
        return pulumi.get(self, "pub_or_sub_instance_ip")

    @property
    @pulumi.getter(name="publishDatabase")
    def publish_database(self) -> Optional[str]:
        """
        Name of the publish SQL Server instance.
        """
        return pulumi.get(self, "publish_database")

    @property
    @pulumi.getter(name="publishSubscribeId")
    def publish_subscribe_id(self) -> Optional[int]:
        """
        The id of the Publish and Subscribe.
        """
        return pulumi.get(self, "publish_subscribe_id")

    @property
    @pulumi.getter(name="publishSubscribeLists")
    def publish_subscribe_lists(self) -> Sequence['outputs.GetPublishSubscribesPublishSubscribeListResult']:
        """
        Publish and subscribe list. Each element contains the following attributes.
        """
        return pulumi.get(self, "publish_subscribe_lists")

    @property
    @pulumi.getter(name="publishSubscribeName")
    def publish_subscribe_name(self) -> Optional[str]:
        """
        The name of the Publish and Subscribe.
        """
        return pulumi.get(self, "publish_subscribe_name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="subscribeDatabase")
    def subscribe_database(self) -> Optional[str]:
        """
        Name of the subscribe SQL Server instance.
        """
        return pulumi.get(self, "subscribe_database")


class AwaitableGetPublishSubscribesResult(GetPublishSubscribesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublishSubscribesResult(
            id=self.id,
            instance_id=self.instance_id,
            pub_or_sub_instance_id=self.pub_or_sub_instance_id,
            pub_or_sub_instance_ip=self.pub_or_sub_instance_ip,
            publish_database=self.publish_database,
            publish_subscribe_id=self.publish_subscribe_id,
            publish_subscribe_lists=self.publish_subscribe_lists,
            publish_subscribe_name=self.publish_subscribe_name,
            result_output_file=self.result_output_file,
            subscribe_database=self.subscribe_database)


def get_publish_subscribes(instance_id: Optional[str] = None,
                           pub_or_sub_instance_id: Optional[str] = None,
                           pub_or_sub_instance_ip: Optional[str] = None,
                           publish_database: Optional[str] = None,
                           publish_subscribe_id: Optional[int] = None,
                           publish_subscribe_name: Optional[str] = None,
                           result_output_file: Optional[str] = None,
                           subscribe_database: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublishSubscribesResult:
    """
    Use this data source to query Publish Subscribe resources for the specific SQL Server instance.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    zones = tencentcloud.Availability.get_zones_by_product(product="sqlserver")
    vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
    subnet = tencentcloud.subnet.Instance("subnet",
        availability_zone=zones.zones[4].name,
        vpc_id=vpc.id,
        cidr_block="10.0.0.0/16",
        is_multicast=False)
    security_group = tencentcloud.security.Group("securityGroup", description="desc.")
    example_pub_general_cloud_instance = tencentcloud.sqlserver.GeneralCloudInstance("examplePubGeneralCloudInstance",
        zone=zones.zones[4].name,
        memory=4,
        storage=100,
        cpu=2,
        machine_type="CLOUD_HSSD",
        instance_charge_type="POSTPAID",
        project_id=0,
        subnet_id=subnet.id,
        vpc_id=vpc.id,
        db_version="2008R2",
        security_group_lists=[security_group.id],
        weeklies=[
            1,
            2,
            3,
            5,
            6,
            7,
        ],
        start_time="00:00",
        span=6,
        resource_tags=[tencentcloud.sqlserver.GeneralCloudInstanceResourceTagArgs(
            tag_key="test",
            tag_value="test",
        )],
        collation="Chinese_PRC_CI_AS",
        time_zone="China Standard Time")
    example_sub_general_cloud_instance = tencentcloud.sqlserver.GeneralCloudInstance("exampleSubGeneralCloudInstance",
        zone=zones.zones[4].name,
        memory=4,
        storage=100,
        cpu=2,
        machine_type="CLOUD_HSSD",
        instance_charge_type="POSTPAID",
        project_id=0,
        subnet_id=subnet.id,
        vpc_id=vpc.id,
        db_version="2008R2",
        security_group_lists=[security_group.id],
        weeklies=[
            1,
            2,
            3,
            5,
            6,
            7,
        ],
        start_time="00:00",
        span=6,
        resource_tags=[tencentcloud.sqlserver.GeneralCloudInstanceResourceTagArgs(
            tag_key="test",
            tag_value="test",
        )],
        collation="Chinese_PRC_CI_AS",
        time_zone="China Standard Time")
    example_pub_db = tencentcloud.sqlserver.Db("examplePubDb",
        instance_id=example_pub_general_cloud_instance.id,
        charset="Chinese_PRC_BIN",
        remark="test-remark")
    example_sub_db = tencentcloud.sqlserver.Db("exampleSubDb",
        instance_id=example_sub_general_cloud_instance.id,
        charset="Chinese_PRC_BIN",
        remark="test-remark")
    example_publish_subscribe = tencentcloud.sqlserver.PublishSubscribe("examplePublishSubscribe",
        publish_instance_id=example_pub_general_cloud_instance.id,
        subscribe_instance_id=example_sub_general_cloud_instance.id,
        publish_subscribe_name="example",
        delete_subscribe_db=False,
        database_tuples=[tencentcloud.sqlserver.PublishSubscribeDatabaseTupleArgs(
            publish_database=example_pub_db.name,
            subscribe_database=example_sub_db.name,
        )])
    example_publish_subscribes = tencentcloud.Sqlserver.get_publish_subscribes_output(instance_id=example_publish_subscribe.publish_instance_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: ID of the SQL Server instance.
    :param str pub_or_sub_instance_id: The subscribe/publish instance ID. It is related to whether the `instance_id` is a publish instance or a subscribe instance. when `instance_id` is a publish instance, this field is filtered according to the subscribe instance ID; when `instance_id` is a subscribe instance, this field is filtering according to the publish instance ID.
    :param str pub_or_sub_instance_ip: The intranet IP of the subscribe/publish instance. It is related to whether the `instance_id` is a publish instance or a subscribe instance. when `instance_id` is a publish instance, this field is filtered according to the intranet IP of the subscribe instance; when `instance_id` is a subscribe instance, this field is based on the publish instance intranet IP filter.
    :param str publish_database: Name of publish database.
    :param int publish_subscribe_id: The id of the Publish and Subscribe.
    :param str publish_subscribe_name: The name of the Publish and Subscribe.
    :param str result_output_file: Used to store results.
    :param str subscribe_database: Name of subscribe database.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['pubOrSubInstanceId'] = pub_or_sub_instance_id
    __args__['pubOrSubInstanceIp'] = pub_or_sub_instance_ip
    __args__['publishDatabase'] = publish_database
    __args__['publishSubscribeId'] = publish_subscribe_id
    __args__['publishSubscribeName'] = publish_subscribe_name
    __args__['resultOutputFile'] = result_output_file
    __args__['subscribeDatabase'] = subscribe_database
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Sqlserver/getPublishSubscribes:getPublishSubscribes', __args__, opts=opts, typ=GetPublishSubscribesResult).value

    return AwaitableGetPublishSubscribesResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        pub_or_sub_instance_id=pulumi.get(__ret__, 'pub_or_sub_instance_id'),
        pub_or_sub_instance_ip=pulumi.get(__ret__, 'pub_or_sub_instance_ip'),
        publish_database=pulumi.get(__ret__, 'publish_database'),
        publish_subscribe_id=pulumi.get(__ret__, 'publish_subscribe_id'),
        publish_subscribe_lists=pulumi.get(__ret__, 'publish_subscribe_lists'),
        publish_subscribe_name=pulumi.get(__ret__, 'publish_subscribe_name'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        subscribe_database=pulumi.get(__ret__, 'subscribe_database'))


@_utilities.lift_output_func(get_publish_subscribes)
def get_publish_subscribes_output(instance_id: Optional[pulumi.Input[str]] = None,
                                  pub_or_sub_instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  pub_or_sub_instance_ip: Optional[pulumi.Input[Optional[str]]] = None,
                                  publish_database: Optional[pulumi.Input[Optional[str]]] = None,
                                  publish_subscribe_id: Optional[pulumi.Input[Optional[int]]] = None,
                                  publish_subscribe_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  subscribe_database: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPublishSubscribesResult]:
    """
    Use this data source to query Publish Subscribe resources for the specific SQL Server instance.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    zones = tencentcloud.Availability.get_zones_by_product(product="sqlserver")
    vpc = tencentcloud.vpc.Instance("vpc", cidr_block="10.0.0.0/16")
    subnet = tencentcloud.subnet.Instance("subnet",
        availability_zone=zones.zones[4].name,
        vpc_id=vpc.id,
        cidr_block="10.0.0.0/16",
        is_multicast=False)
    security_group = tencentcloud.security.Group("securityGroup", description="desc.")
    example_pub_general_cloud_instance = tencentcloud.sqlserver.GeneralCloudInstance("examplePubGeneralCloudInstance",
        zone=zones.zones[4].name,
        memory=4,
        storage=100,
        cpu=2,
        machine_type="CLOUD_HSSD",
        instance_charge_type="POSTPAID",
        project_id=0,
        subnet_id=subnet.id,
        vpc_id=vpc.id,
        db_version="2008R2",
        security_group_lists=[security_group.id],
        weeklies=[
            1,
            2,
            3,
            5,
            6,
            7,
        ],
        start_time="00:00",
        span=6,
        resource_tags=[tencentcloud.sqlserver.GeneralCloudInstanceResourceTagArgs(
            tag_key="test",
            tag_value="test",
        )],
        collation="Chinese_PRC_CI_AS",
        time_zone="China Standard Time")
    example_sub_general_cloud_instance = tencentcloud.sqlserver.GeneralCloudInstance("exampleSubGeneralCloudInstance",
        zone=zones.zones[4].name,
        memory=4,
        storage=100,
        cpu=2,
        machine_type="CLOUD_HSSD",
        instance_charge_type="POSTPAID",
        project_id=0,
        subnet_id=subnet.id,
        vpc_id=vpc.id,
        db_version="2008R2",
        security_group_lists=[security_group.id],
        weeklies=[
            1,
            2,
            3,
            5,
            6,
            7,
        ],
        start_time="00:00",
        span=6,
        resource_tags=[tencentcloud.sqlserver.GeneralCloudInstanceResourceTagArgs(
            tag_key="test",
            tag_value="test",
        )],
        collation="Chinese_PRC_CI_AS",
        time_zone="China Standard Time")
    example_pub_db = tencentcloud.sqlserver.Db("examplePubDb",
        instance_id=example_pub_general_cloud_instance.id,
        charset="Chinese_PRC_BIN",
        remark="test-remark")
    example_sub_db = tencentcloud.sqlserver.Db("exampleSubDb",
        instance_id=example_sub_general_cloud_instance.id,
        charset="Chinese_PRC_BIN",
        remark="test-remark")
    example_publish_subscribe = tencentcloud.sqlserver.PublishSubscribe("examplePublishSubscribe",
        publish_instance_id=example_pub_general_cloud_instance.id,
        subscribe_instance_id=example_sub_general_cloud_instance.id,
        publish_subscribe_name="example",
        delete_subscribe_db=False,
        database_tuples=[tencentcloud.sqlserver.PublishSubscribeDatabaseTupleArgs(
            publish_database=example_pub_db.name,
            subscribe_database=example_sub_db.name,
        )])
    example_publish_subscribes = tencentcloud.Sqlserver.get_publish_subscribes_output(instance_id=example_publish_subscribe.publish_instance_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str instance_id: ID of the SQL Server instance.
    :param str pub_or_sub_instance_id: The subscribe/publish instance ID. It is related to whether the `instance_id` is a publish instance or a subscribe instance. when `instance_id` is a publish instance, this field is filtered according to the subscribe instance ID; when `instance_id` is a subscribe instance, this field is filtering according to the publish instance ID.
    :param str pub_or_sub_instance_ip: The intranet IP of the subscribe/publish instance. It is related to whether the `instance_id` is a publish instance or a subscribe instance. when `instance_id` is a publish instance, this field is filtered according to the intranet IP of the subscribe instance; when `instance_id` is a subscribe instance, this field is based on the publish instance intranet IP filter.
    :param str publish_database: Name of publish database.
    :param int publish_subscribe_id: The id of the Publish and Subscribe.
    :param str publish_subscribe_name: The name of the Publish and Subscribe.
    :param str result_output_file: Used to store results.
    :param str subscribe_database: Name of subscribe database.
    """
    ...
