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
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
    'get_application_output',
]

@pulumi.output_type
class GetApplicationResult:
    """
    A collection of values returned by getApplication.
    """
    def __init__(__self__, application_id_lists=None, application_resource_type_lists=None, application_type=None, id=None, microservice_type=None, result_output_file=None, results=None):
        if application_id_lists and not isinstance(application_id_lists, list):
            raise TypeError("Expected argument 'application_id_lists' to be a list")
        pulumi.set(__self__, "application_id_lists", application_id_lists)
        if application_resource_type_lists and not isinstance(application_resource_type_lists, list):
            raise TypeError("Expected argument 'application_resource_type_lists' to be a list")
        pulumi.set(__self__, "application_resource_type_lists", application_resource_type_lists)
        if application_type and not isinstance(application_type, str):
            raise TypeError("Expected argument 'application_type' to be a str")
        pulumi.set(__self__, "application_type", application_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if microservice_type and not isinstance(microservice_type, str):
            raise TypeError("Expected argument 'microservice_type' to be a str")
        pulumi.set(__self__, "microservice_type", microservice_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="applicationIdLists")
    def application_id_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "application_id_lists")

    @property
    @pulumi.getter(name="applicationResourceTypeLists")
    def application_resource_type_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "application_resource_type_lists")

    @property
    @pulumi.getter(name="applicationType")
    def application_type(self) -> Optional[str]:
        """
        The type of the application.
        """
        return pulumi.get(self, "application_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="microserviceType")
    def microservice_type(self) -> Optional[str]:
        """
        The microservice type of the application.
        """
        return pulumi.get(self, "microservice_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetApplicationResultResult']:
        """
        The application paging list information.
        """
        return pulumi.get(self, "results")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            application_id_lists=self.application_id_lists,
            application_resource_type_lists=self.application_resource_type_lists,
            application_type=self.application_type,
            id=self.id,
            microservice_type=self.microservice_type,
            result_output_file=self.result_output_file,
            results=self.results)


def get_application(application_id_lists: Optional[Sequence[str]] = None,
                    application_resource_type_lists: Optional[Sequence[str]] = None,
                    application_type: Optional[str] = None,
                    microservice_type: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    Use this data source to query detailed information of tsf application

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    application = tencentcloud.Tsf.get_application(application_id_lists=["application-a24x29xv"],
        application_type="V",
        microservice_type="N")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] application_id_lists: Id list.
    :param Sequence[str] application_resource_type_lists: An array of application resource types.
    :param str application_type: The application type. V OR C, V means VM, C means container.
    :param str microservice_type: The microservice type of the application.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['applicationIdLists'] = application_id_lists
    __args__['applicationResourceTypeLists'] = application_resource_type_lists
    __args__['applicationType'] = application_type
    __args__['microserviceType'] = microservice_type
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Tsf/getApplication:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        application_id_lists=pulumi.get(__ret__, 'application_id_lists'),
        application_resource_type_lists=pulumi.get(__ret__, 'application_resource_type_lists'),
        application_type=pulumi.get(__ret__, 'application_type'),
        id=pulumi.get(__ret__, 'id'),
        microservice_type=pulumi.get(__ret__, 'microservice_type'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'),
        results=pulumi.get(__ret__, 'results'))


@_utilities.lift_output_func(get_application)
def get_application_output(application_id_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           application_resource_type_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           application_type: Optional[pulumi.Input[Optional[str]]] = None,
                           microservice_type: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationResult]:
    """
    Use this data source to query detailed information of tsf application

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    application = tencentcloud.Tsf.get_application(application_id_lists=["application-a24x29xv"],
        application_type="V",
        microservice_type="N")
    ```
    <!--End PulumiCodeChooser -->


    :param Sequence[str] application_id_lists: Id list.
    :param Sequence[str] application_resource_type_lists: An array of application resource types.
    :param str application_type: The application type. V OR C, V means VM, C means container.
    :param str microservice_type: The microservice type of the application.
    :param str result_output_file: Used to save results.
    """
    ...
