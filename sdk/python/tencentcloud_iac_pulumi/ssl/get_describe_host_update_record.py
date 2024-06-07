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
    'GetDescribeHostUpdateRecordResult',
    'AwaitableGetDescribeHostUpdateRecordResult',
    'get_describe_host_update_record',
    'get_describe_host_update_record_output',
]

@pulumi.output_type
class GetDescribeHostUpdateRecordResult:
    """
    A collection of values returned by getDescribeHostUpdateRecord.
    """
    def __init__(__self__, certificate_id=None, deploy_record_lists=None, id=None, old_certificate_id=None, result_output_file=None):
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if deploy_record_lists and not isinstance(deploy_record_lists, list):
            raise TypeError("Expected argument 'deploy_record_lists' to be a list")
        pulumi.set(__self__, "deploy_record_lists", deploy_record_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if old_certificate_id and not isinstance(old_certificate_id, str):
            raise TypeError("Expected argument 'old_certificate_id' to be a str")
        pulumi.set(__self__, "old_certificate_id", old_certificate_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[str]:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="deployRecordLists")
    def deploy_record_lists(self) -> Sequence['outputs.GetDescribeHostUpdateRecordDeployRecordListResult']:
        """
        Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        """
        return pulumi.get(self, "deploy_record_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="oldCertificateId")
    def old_certificate_id(self) -> Optional[str]:
        return pulumi.get(self, "old_certificate_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGetDescribeHostUpdateRecordResult(GetDescribeHostUpdateRecordResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDescribeHostUpdateRecordResult(
            certificate_id=self.certificate_id,
            deploy_record_lists=self.deploy_record_lists,
            id=self.id,
            old_certificate_id=self.old_certificate_id,
            result_output_file=self.result_output_file)


def get_describe_host_update_record(certificate_id: Optional[str] = None,
                                    old_certificate_id: Optional[str] = None,
                                    result_output_file: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDescribeHostUpdateRecordResult:
    """
    Use this data source to query detailed information of ssl describe_host_update_record

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_update_record = tencentcloud.Ssl.get_describe_host_update_record(old_certificate_id="8u8DII0l")
    ```
    <!--End PulumiCodeChooser -->


    :param str certificate_id: New certificate ID.
    :param str old_certificate_id: Original certificate ID.
    :param str result_output_file: Used to save results.
    """
    __args__ = dict()
    __args__['certificateId'] = certificate_id
    __args__['oldCertificateId'] = old_certificate_id
    __args__['resultOutputFile'] = result_output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssl/getDescribeHostUpdateRecord:getDescribeHostUpdateRecord', __args__, opts=opts, typ=GetDescribeHostUpdateRecordResult).value

    return AwaitableGetDescribeHostUpdateRecordResult(
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        deploy_record_lists=pulumi.get(__ret__, 'deploy_record_lists'),
        id=pulumi.get(__ret__, 'id'),
        old_certificate_id=pulumi.get(__ret__, 'old_certificate_id'),
        result_output_file=pulumi.get(__ret__, 'result_output_file'))


@_utilities.lift_output_func(get_describe_host_update_record)
def get_describe_host_update_record_output(certificate_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           old_certificate_id: Optional[pulumi.Input[Optional[str]]] = None,
                                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDescribeHostUpdateRecordResult]:
    """
    Use this data source to query detailed information of ssl describe_host_update_record

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    describe_host_update_record = tencentcloud.Ssl.get_describe_host_update_record(old_certificate_id="8u8DII0l")
    ```
    <!--End PulumiCodeChooser -->


    :param str certificate_id: New certificate ID.
    :param str old_certificate_id: Original certificate ID.
    :param str result_output_file: Used to save results.
    """
    ...
