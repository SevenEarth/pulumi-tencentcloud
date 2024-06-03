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
    'GetDescribeIndexListResult',
    'AwaitableGetDescribeIndexListResult',
    'get_describe_index_list',
    'get_describe_index_list_output',
]

@pulumi.output_type
class GetDescribeIndexListResult:
    """
    A collection of values returned by getDescribeIndexList.
    """
    def __init__(__self__, id=None, index_meta_fields=None, index_name=None, index_status_lists=None, index_type=None, instance_id=None, order=None, order_by=None, password=None, result_output_file=None, username=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if index_meta_fields and not isinstance(index_meta_fields, list):
            raise TypeError("Expected argument 'index_meta_fields' to be a list")
        pulumi.set(__self__, "index_meta_fields", index_meta_fields)
        if index_name and not isinstance(index_name, str):
            raise TypeError("Expected argument 'index_name' to be a str")
        pulumi.set(__self__, "index_name", index_name)
        if index_status_lists and not isinstance(index_status_lists, list):
            raise TypeError("Expected argument 'index_status_lists' to be a list")
        pulumi.set(__self__, "index_status_lists", index_status_lists)
        if index_type and not isinstance(index_type, str):
            raise TypeError("Expected argument 'index_type' to be a str")
        pulumi.set(__self__, "index_type", index_type)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if order and not isinstance(order, str):
            raise TypeError("Expected argument 'order' to be a str")
        pulumi.set(__self__, "order", order)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="indexMetaFields")
    def index_meta_fields(self) -> Sequence['outputs.GetDescribeIndexListIndexMetaFieldResult']:
        """
        Index metadata field.
        """
        return pulumi.get(self, "index_meta_fields")

    @property
    @pulumi.getter(name="indexName")
    def index_name(self) -> Optional[str]:
        """
        Index name.
        """
        return pulumi.get(self, "index_name")

    @property
    @pulumi.getter(name="indexStatusLists")
    def index_status_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "index_status_lists")

    @property
    @pulumi.getter(name="indexType")
    def index_type(self) -> str:
        """
        Index type.
        """
        return pulumi.get(self, "index_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")


class AwaitableGetDescribeIndexListResult(GetDescribeIndexListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeIndexListResult(
            id=self.id,
            index_meta_fields=self.index_meta_fields,
            index_name=self.index_name,
            index_status_lists=self.index_status_lists,
            index_type=self.index_type,
            instance_id=self.instance_id,
            order=self.order,
            order_by=self.order_by,
            password=self.password,
            result_output_file=self.result_output_file,
            username=self.username)


def get_describe_index_list(index_name: Optional[str] = None,
                            index_status_lists: Optional[Sequence[str]] = None,
                            index_type: Optional[str] = None,
                            instance_id: Optional[str] = None,
                            order: Optional[str] = None,
                            order_by: Optional[str] = None,
                            password: Optional[str] = None,
                            result_output_file: Optional[str] = None,
                            username: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeIndexListResult:
    """
    Use this data source to query detailed information of elasticsearch index list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_index_list = tencentcloud.Elasticsearch.get_describe_index_list(index_type="normal",
        instance_id="es-nni6pm4s")
    ```
    <!--End PulumiCodeChooser -->


    :param str index_name: Index name. If you fill in the blanks, get all indexes.
    :param Sequence[str] index_status_lists: Index status list.
    :param str index_type: Index type. `auto`: Autonomous index; `normal`: General index.
    :param str instance_id: ES cluster id.
    :param str order: Sort order, which supports asc and desc. The default is desc data format asc,desc.
    :param str order_by: Sort field. Support index name: IndexName, index storage: IndexStorage, index creation time: IndexCreateTime.
    :param str password: Cluster access password.
    :param str result_output_file: Used to save results.
    :param str username: Cluster access user name.
    """
    __args__ = dict()
    __args__['indexName'] = index_name
    __args__['indexStatusLists'] = index_status_lists
    __args__['indexType'] = index_type
    __args__['instanceId'] = instance_id
    __args__['order'] = order
    __args__['orderBy'] = order_by
    __args__['password'] = password
    __args__['resultOutputFile'] = result_output_file
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Elasticsearch/getDescribeIndexList:getDescribeIndexList', __args__, opts=opts, typ=GetDescribeIndexListResult).value

    return AwaitableGetDescribeIndexListResult(
        id=pulumi.get(__ret__, 'id'),
        index_meta_fields=pulumi.get(__ret__, 'index_meta_fields'),
        index_name=pulumi.get(__ret__, 'index_name'),
        index_status_lists=pulumi.get(__ret__, 'index_status_lists'),
        index_type=pulumi.get(__ret__, 'index_type'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        order=pulumi.get(__ret__, 'order'),
        order_by=pulumi.get(__ret__, 'order_by'),
        password=pulumi.get(__ret__, 'password'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        username=pulumi.get(__ret__, 'username'))


@_utilities.lift_output_func(get_describe_index_list)
def get_describe_index_list_output(index_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   index_status_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                   index_type: Optional[pulumi.Input[str]] = None,
                                   instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   order: Optional[pulumi.Input[Optional[str]]] = None,
                                   order_by: Optional[pulumi.Input[Optional[str]]] = None,
                                   password: Optional[pulumi.Input[Optional[str]]] = None,
                                   result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   username: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeIndexListResult]:
    """
    Use this data source to query detailed information of elasticsearch index list

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_index_list = tencentcloud.Elasticsearch.get_describe_index_list(index_type="normal",
        instance_id="es-nni6pm4s")
    ```
    <!--End PulumiCodeChooser -->


    :param str index_name: Index name. If you fill in the blanks, get all indexes.
    :param Sequence[str] index_status_lists: Index status list.
    :param str index_type: Index type. `auto`: Autonomous index; `normal`: General index.
    :param str instance_id: ES cluster id.
    :param str order: Sort order, which supports asc and desc. The default is desc data format asc,desc.
    :param str order_by: Sort field. Support index name: IndexName, index storage: IndexStorage, index creation time: IndexCreateTime.
    :param str password: Cluster access password.
    :param str result_output_file: Used to save results.
    :param str username: Cluster access user name.
    """
    ...
