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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
    'get_repository_output',
]

@pulumi.output_type
class GetRepositoryResult:
    """
    A collection of values returned by getRepository.
    """
    def __init__(__self__, id=None, repository_type=None, result_output_file=None, results=None, search_word=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository_type and not isinstance(repository_type, str):
            raise TypeError("Expected argument 'repository_type' to be a str")
        pulumi.set(__self__, "repository_type", repository_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if search_word and not isinstance(search_word, str):
            raise TypeError("Expected argument 'search_word' to be a str")
        pulumi.set(__self__, "search_word", search_word)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="repositoryType")
    def repository_type(self) -> Optional[str]:
        """
        Repository type (default Repository: default, private Repository: private).
        """
        return pulumi.get(self, "repository_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetRepositoryResultResult']:
        """
        A list of Repository information that meets the query criteria.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="searchWord")
    def search_word(self) -> Optional[str]:
        return pulumi.get(self, "search_word")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            id=self.id,
            repository_type=self.repository_type,
            result_output_file=self.result_output_file,
            results=self.results,
            search_word=self.search_word)


def get_repository(repository_type: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   search_word: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    Use this data source to query detailed information of tsf repository

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    repository = tencentcloud.Tsf.get_repository(repository_type="default",
        search_word="test")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository_type: Repository type (default Repository: default, private Repository: private).
    :param str result_output_file: Used to save results.
    :param str search_word: Query keywords (search by Repository name).
    """
    __args__ = dict()
    __args__['repositoryType'] = repository_type
    __args__['resultOutputFile'] = result_output_file
    __args__['searchWord'] = search_word
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tsf/getRepository:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        id=pulumi.get(__ret__, 'id'),
        repository_type=pulumi.get(__ret__, 'repository_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'),
        search_word=pulumi.get(__ret__, 'search_word'))


@_utilities.lift_output_func(get_repository)
def get_repository_output(repository_type: Optional[pulumi.Input[Optional[str]]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          search_word: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryResult]:
    """
    Use this data source to query detailed information of tsf repository

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    repository = tencentcloud.Tsf.get_repository(repository_type="default",
        search_word="test")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository_type: Repository type (default Repository: default, private Repository: private).
    :param str result_output_file: Used to save results.
    :param str search_word: Query keywords (search by Repository name).
    """
    ...
