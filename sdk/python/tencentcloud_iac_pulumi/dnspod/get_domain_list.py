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

__all__ = [
    'GetDomainListResult',
    'AwaitableGetDomainListResult',
    'get_domain_list',
    'get_domain_list_output',
]

@pulumi.output_type
class GetDomainListResult:
    """
    A collection of values returned by getDomainList.
    """
    def __init__(__self__, domain_lists=None, group_ids=None, id=None, keyword=None, packages=None, project_id=None, record_count_begin=None, record_count_end=None, remark=None, result_output_file=None, sort_field=None, sort_type=None, statuses=None, tags=None, type=None, updated_at_begin=None, updated_at_end=None):
        if domain_lists and not isinstance(domain_lists, list):
            raise TypeError("Expected argument 'domain_lists' to be a list")
        pulumi.set(__self__, "domain_lists", domain_lists)
        if group_ids and not isinstance(group_ids, list):
            raise TypeError("Expected argument 'group_ids' to be a list")
        pulumi.set(__self__, "group_ids", group_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keyword and not isinstance(keyword, str):
            raise TypeError("Expected argument 'keyword' to be a str")
        pulumi.set(__self__, "keyword", keyword)
        if packages and not isinstance(packages, list):
            raise TypeError("Expected argument 'packages' to be a list")
        pulumi.set(__self__, "packages", packages)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if record_count_begin and not isinstance(record_count_begin, int):
            raise TypeError("Expected argument 'record_count_begin' to be a int")
        pulumi.set(__self__, "record_count_begin", record_count_begin)
        if record_count_end and not isinstance(record_count_end, int):
            raise TypeError("Expected argument 'record_count_end' to be a int")
        pulumi.set(__self__, "record_count_end", record_count_end)
        if remark and not isinstance(remark, str):
            raise TypeError("Expected argument 'remark' to be a str")
        pulumi.set(__self__, "remark", remark)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sort_field and not isinstance(sort_field, str):
            raise TypeError("Expected argument 'sort_field' to be a str")
        pulumi.set(__self__, "sort_field", sort_field)
        if sort_type and not isinstance(sort_type, str):
            raise TypeError("Expected argument 'sort_type' to be a str")
        pulumi.set(__self__, "sort_type", sort_type)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at_begin and not isinstance(updated_at_begin, str):
            raise TypeError("Expected argument 'updated_at_begin' to be a str")
        pulumi.set(__self__, "updated_at_begin", updated_at_begin)
        if updated_at_end and not isinstance(updated_at_end, str):
            raise TypeError("Expected argument 'updated_at_end' to be a str")
        pulumi.set(__self__, "updated_at_end", updated_at_end)

    @property
    @pulumi.getter(name="domainLists")
    def domain_lists(self) -> Sequence['outputs.GetDomainListDomainListResult']:
        """
        Domain list.
        """
        return pulumi.get(self, "domain_lists")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[Sequence[int]]:
        """
        Group Id the domain belongs to.
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keyword(self) -> Optional[str]:
        return pulumi.get(self, "keyword")

    @property
    @pulumi.getter
    def packages(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "packages")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="recordCountBegin")
    def record_count_begin(self) -> Optional[int]:
        return pulumi.get(self, "record_count_begin")

    @property
    @pulumi.getter(name="recordCountEnd")
    def record_count_end(self) -> Optional[int]:
        return pulumi.get(self, "record_count_end")

    @property
    @pulumi.getter
    def remark(self) -> Optional[str]:
        """
        Domain remark description.
        """
        return pulumi.get(self, "remark")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="sortField")
    def sort_field(self) -> Optional[str]:
        return pulumi.get(self, "sort_field")

    @property
    @pulumi.getter(name="sortType")
    def sort_type(self) -> Optional[str]:
        return pulumi.get(self, "sort_type")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        Domain status, normal: ENABLE, paused: PAUSE, banned: SPAM.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GetDomainListTagResult']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAtBegin")
    def updated_at_begin(self) -> Optional[str]:
        return pulumi.get(self, "updated_at_begin")

    @property
    @pulumi.getter(name="updatedAtEnd")
    def updated_at_end(self) -> Optional[str]:
        return pulumi.get(self, "updated_at_end")


class AwaitableGetDomainListResult(GetDomainListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainListResult(
            domain_lists=self.domain_lists,
            group_ids=self.group_ids,
            id=self.id,
            keyword=self.keyword,
            packages=self.packages,
            project_id=self.project_id,
            record_count_begin=self.record_count_begin,
            record_count_end=self.record_count_end,
            remark=self.remark,
            result_output_file=self.result_output_file,
            sort_field=self.sort_field,
            sort_type=self.sort_type,
            statuses=self.statuses,
            tags=self.tags,
            type=self.type,
            updated_at_begin=self.updated_at_begin,
            updated_at_end=self.updated_at_end)


def get_domain_list(group_ids: Optional[Sequence[int]] = None,
                    keyword: Optional[str] = None,
                    packages: Optional[Sequence[str]] = None,
                    project_id: Optional[int] = None,
                    record_count_begin: Optional[int] = None,
                    record_count_end: Optional[int] = None,
                    remark: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    sort_field: Optional[str] = None,
                    sort_type: Optional[str] = None,
                    statuses: Optional[Sequence[str]] = None,
                    tags: Optional[Sequence[pulumi.InputType['GetDomainListTagArgs']]] = None,
                    type: Optional[str] = None,
                    updated_at_begin: Optional[str] = None,
                    updated_at_end: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainListResult:
    """
    Use this data source to query detailed information of dnspod domain_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_list = tencentcloud.Dnspod.get_domain_list(group_ids=[1],
        keyword="",
        packages=[""],
        project_id=-1,
        record_count_begin=0,
        record_count_end=100,
        remark="",
        sort_field="UPDATED_ON",
        sort_type="DESC",
        statuses=["PAUSE"],
        tags=[tencentcloud.dnspod.GetDomainListTagArgs(
            tag_key="created_by",
            tag_values=["terraform"],
        )],
        type="ALL",
        updated_at_begin="2021-05-01 03:00:00",
        updated_at_end="2024-05-10 20:00:00")
    ```


    :param Sequence[int] group_ids: Get domain names based on domain group id, which can be obtained through the GroupId field in DescribeDomain or DescribeDomainList interface.
    :param str keyword: Get domain names based on keywords.
    :param Sequence[str] packages: Get domain names based on the package, which can be obtained through the Grade field in DescribeDomain or DescribeDomainList interface.
    :param int project_id: Project ID.
    :param int record_count_begin: The start point of the domain name&amp;#39;s record count query range.
    :param int record_count_end: The end point of the domain name&amp;#39;s record count query range.
    :param str remark: Get domain names based on remark information.
    :param str result_output_file: Used to save results.
    :param str sort_field: Sorting field. Available values are NAME, STATUS, RECORDS, GRADE, UPDATED_ON. NAME: Domain name STATUS: Domain status RECORDS: Number of records GRADE: Package level UPDATED_ON: Update time.
    :param str sort_type: Sorting type, ascending: ASC, descending: DESC.
    :param Sequence[str] statuses: Get domain names based on domain status. Available values are ENABLE, LOCK, PAUSE, SPAM. ENABLE: Normal LOCK: Locked PAUSE: Paused SPAM: Banned.
    :param Sequence[pulumi.InputType['GetDomainListTagArgs']] tags: Tag description list.
    :param str type: Get domain names based on domain group type. Available values are ALL, MINE, SHARE, RECENT. ALL: All MINE: My domain names SHARE: Domain names shared with me RECENT: Recently operated domain names.
    :param str updated_at_begin: The start time of the domain name&amp;#39;s update time to be obtained, such as &amp;#39;2021-05-01 03:00:00&amp;#39;.
    :param str updated_at_end: The end time of the domain name&amp;#39;s update time to be obtained, such as &amp;#39;2021-05-10 20:00:00&amp;#39;.
    """
    __args__ = dict()
    __args__['groupIds'] = group_ids
    __args__['keyword'] = keyword
    __args__['packages'] = packages
    __args__['projectId'] = project_id
    __args__['recordCountBegin'] = record_count_begin
    __args__['recordCountEnd'] = record_count_end
    __args__['remark'] = remark
    __args__['resultOutputFile'] = result_output_file
    __args__['sortField'] = sort_field
    __args__['sortType'] = sort_type
    __args__['statuses'] = statuses
    __args__['tags'] = tags
    __args__['type'] = type
    __args__['updatedAtBegin'] = updated_at_begin
    __args__['updatedAtEnd'] = updated_at_end
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Dnspod/getDomainList:getDomainList', __args__, opts=opts, typ=GetDomainListResult).value

    return AwaitableGetDomainListResult(
        domain_lists=__ret__.domain_lists,
        group_ids=__ret__.group_ids,
        id=__ret__.id,
        keyword=__ret__.keyword,
        packages=__ret__.packages,
        project_id=__ret__.project_id,
        record_count_begin=__ret__.record_count_begin,
        record_count_end=__ret__.record_count_end,
        remark=__ret__.remark,
        result_output_file=__ret__.result_output_file,
        sort_field=__ret__.sort_field,
        sort_type=__ret__.sort_type,
        statuses=__ret__.statuses,
        tags=__ret__.tags,
        type=__ret__.type,
        updated_at_begin=__ret__.updated_at_begin,
        updated_at_end=__ret__.updated_at_end)


@_utilities.lift_output_func(get_domain_list)
def get_domain_list_output(group_ids: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                           keyword: Optional[pulumi.Input[Optional[str]]] = None,
                           packages: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           project_id: Optional[pulumi.Input[Optional[int]]] = None,
                           record_count_begin: Optional[pulumi.Input[Optional[int]]] = None,
                           record_count_end: Optional[pulumi.Input[Optional[int]]] = None,
                           remark: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           sort_field: Optional[pulumi.Input[Optional[str]]] = None,
                           sort_type: Optional[pulumi.Input[Optional[str]]] = None,
                           statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDomainListTagArgs']]]]] = None,
                           type: Optional[pulumi.Input[str]] = None,
                           updated_at_begin: Optional[pulumi.Input[Optional[str]]] = None,
                           updated_at_end: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainListResult]:
    """
    Use this data source to query detailed information of dnspod domain_list

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    domain_list = tencentcloud.Dnspod.get_domain_list(group_ids=[1],
        keyword="",
        packages=[""],
        project_id=-1,
        record_count_begin=0,
        record_count_end=100,
        remark="",
        sort_field="UPDATED_ON",
        sort_type="DESC",
        statuses=["PAUSE"],
        tags=[tencentcloud.dnspod.GetDomainListTagArgs(
            tag_key="created_by",
            tag_values=["terraform"],
        )],
        type="ALL",
        updated_at_begin="2021-05-01 03:00:00",
        updated_at_end="2024-05-10 20:00:00")
    ```


    :param Sequence[int] group_ids: Get domain names based on domain group id, which can be obtained through the GroupId field in DescribeDomain or DescribeDomainList interface.
    :param str keyword: Get domain names based on keywords.
    :param Sequence[str] packages: Get domain names based on the package, which can be obtained through the Grade field in DescribeDomain or DescribeDomainList interface.
    :param int project_id: Project ID.
    :param int record_count_begin: The start point of the domain name&amp;#39;s record count query range.
    :param int record_count_end: The end point of the domain name&amp;#39;s record count query range.
    :param str remark: Get domain names based on remark information.
    :param str result_output_file: Used to save results.
    :param str sort_field: Sorting field. Available values are NAME, STATUS, RECORDS, GRADE, UPDATED_ON. NAME: Domain name STATUS: Domain status RECORDS: Number of records GRADE: Package level UPDATED_ON: Update time.
    :param str sort_type: Sorting type, ascending: ASC, descending: DESC.
    :param Sequence[str] statuses: Get domain names based on domain status. Available values are ENABLE, LOCK, PAUSE, SPAM. ENABLE: Normal LOCK: Locked PAUSE: Paused SPAM: Banned.
    :param Sequence[pulumi.InputType['GetDomainListTagArgs']] tags: Tag description list.
    :param str type: Get domain names based on domain group type. Available values are ALL, MINE, SHARE, RECENT. ALL: All MINE: My domain names SHARE: Domain names shared with me RECENT: Recently operated domain names.
    :param str updated_at_begin: The start time of the domain name&amp;#39;s update time to be obtained, such as &amp;#39;2021-05-01 03:00:00&amp;#39;.
    :param str updated_at_end: The end time of the domain name&amp;#39;s update time to be obtained, such as &amp;#39;2021-05-10 20:00:00&amp;#39;.
    """
    ...
