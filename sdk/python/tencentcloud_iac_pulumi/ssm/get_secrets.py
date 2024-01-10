# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetSecretsResult',
    'AwaitableGetSecretsResult',
    'get_secrets',
    'get_secrets_output',
]

@pulumi.output_type
class GetSecretsResult:
    """
    A collection of values returned by getSecrets.
    """
    def __init__(__self__, id=None, order_type=None, product_name=None, result_output_file=None, secret_lists=None, secret_name=None, secret_type=None, state=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order_type and not isinstance(order_type, int):
            raise TypeError("Expected argument 'order_type' to be a int")
        pulumi.set(__self__, "order_type", order_type)
        if product_name and not isinstance(product_name, str):
            raise TypeError("Expected argument 'product_name' to be a str")
        pulumi.set(__self__, "product_name", product_name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if secret_lists and not isinstance(secret_lists, list):
            raise TypeError("Expected argument 'secret_lists' to be a list")
        pulumi.set(__self__, "secret_lists", secret_lists)
        if secret_name and not isinstance(secret_name, str):
            raise TypeError("Expected argument 'secret_name' to be a str")
        pulumi.set(__self__, "secret_name", secret_name)
        if secret_type and not isinstance(secret_type, int):
            raise TypeError("Expected argument 'secret_type' to be a int")
        pulumi.set(__self__, "secret_type", secret_type)
        if state and not isinstance(state, int):
            raise TypeError("Expected argument 'state' to be a int")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="orderType")
    def order_type(self) -> Optional[int]:
        return pulumi.get(self, "order_type")

    @property
    @pulumi.getter(name="productName")
    def product_name(self) -> Optional[str]:
        """
        Cloud product name, only effective when SecretType is 1, which means the credential type is cloud product credential.
        """
        return pulumi.get(self, "product_name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="secretLists")
    def secret_lists(self) -> Sequence['outputs.GetSecretsSecretListResult']:
        """
        A list of SSM secrets.
        """
        return pulumi.get(self, "secret_lists")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[str]:
        """
        Name of secret.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[int]:
        """
        0- User defined credentials; 1- Cloud product credentials; 2- SSH key pair credentials; 3- Cloud API key pair credentials.
        """
        return pulumi.get(self, "secret_type")

    @property
    @pulumi.getter
    def state(self) -> Optional[int]:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableGetSecretsResult(GetSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretsResult(
            id=self.id,
            order_type=self.order_type,
            product_name=self.product_name,
            result_output_file=self.result_output_file,
            secret_lists=self.secret_lists,
            secret_name=self.secret_name,
            secret_type=self.secret_type,
            state=self.state,
            tags=self.tags)


def get_secrets(order_type: Optional[int] = None,
                product_name: Optional[str] = None,
                result_output_file: Optional[str] = None,
                secret_name: Optional[str] = None,
                secret_type: Optional[int] = None,
                state: Optional[int] = None,
                tags: Optional[Mapping[str, Any]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretsResult:
    """
    Use this data source to query detailed information of SSM secret

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    example_secret = tencentcloud.ssm.Secret("exampleSecret",
        secret_name="tf_example",
        description="desc.",
        tags={
            "createdBy": "terraform",
        })
    example_secrets = example_secret.secret_name.apply(lambda secret_name: tencentcloud.Ssm.get_secrets_output(secret_name=secret_name,
        state=1))
    ```
    ### OR you can filter by tags

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Ssm.get_secrets(secret_name=tencentcloud_ssm_secret["example"]["secret_name"],
        state=1,
        tags={
            "createdBy": "terraform",
        })
    ```


    :param int order_type: The order to sort the create time of secret. `0` - desc, `1` - asc. Default value is `0`.
    :param str product_name: This parameter only takes effect when the SecretType parameter value is 1. When the SecretType value is 1, if the Product Name value is empty, it means to query all types of cloud product credentials. If the Product Name value is MySQL, it means to query MySQL database credentials. If the Product Name value is Tdsql mysql, it means to query Tdsql (MySQL version) credentials.
    :param str result_output_file: Used to save results.
    :param str secret_name: Secret name used to filter result.
    :param int secret_type: 0- represents user-defined credentials, defaults to 0. 1- represents the user's cloud product credentials. 2- represents SSH key pair credentials. 3- represents cloud API key pair credentials.
    :param int state: Filter by state of secret. `0` - all secrets are queried, `1` - only Enabled secrets are queried, `2` - only Disabled secrets are queried, `3` - only PendingDelete secrets are queried.
    :param Mapping[str, Any] tags: Tags to filter secret.
    """
    __args__ = dict()
    __args__['orderType'] = order_type
    __args__['productName'] = product_name
    __args__['resultOutputFile'] = result_output_file
    __args__['secretName'] = secret_name
    __args__['secretType'] = secret_type
    __args__['state'] = state
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ssm/getSecrets:getSecrets', __args__, opts=opts, typ=GetSecretsResult).value

    return AwaitableGetSecretsResult(
        id=__ret__.id,
        order_type=__ret__.order_type,
        product_name=__ret__.product_name,
        result_output_file=__ret__.result_output_file,
        secret_lists=__ret__.secret_lists,
        secret_name=__ret__.secret_name,
        secret_type=__ret__.secret_type,
        state=__ret__.state,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_secrets)
def get_secrets_output(order_type: Optional[pulumi.Input[Optional[int]]] = None,
                       product_name: Optional[pulumi.Input[Optional[str]]] = None,
                       result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       secret_name: Optional[pulumi.Input[Optional[str]]] = None,
                       secret_type: Optional[pulumi.Input[Optional[int]]] = None,
                       state: Optional[pulumi.Input[Optional[int]]] = None,
                       tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretsResult]:
    """
    Use this data source to query detailed information of SSM secret

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud
    import tencentcloud_iac_pulumi as tencentcloud

    example_secret = tencentcloud.ssm.Secret("exampleSecret",
        secret_name="tf_example",
        description="desc.",
        tags={
            "createdBy": "terraform",
        })
    example_secrets = example_secret.secret_name.apply(lambda secret_name: tencentcloud.Ssm.get_secrets_output(secret_name=secret_name,
        state=1))
    ```
    ### OR you can filter by tags

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Ssm.get_secrets(secret_name=tencentcloud_ssm_secret["example"]["secret_name"],
        state=1,
        tags={
            "createdBy": "terraform",
        })
    ```


    :param int order_type: The order to sort the create time of secret. `0` - desc, `1` - asc. Default value is `0`.
    :param str product_name: This parameter only takes effect when the SecretType parameter value is 1. When the SecretType value is 1, if the Product Name value is empty, it means to query all types of cloud product credentials. If the Product Name value is MySQL, it means to query MySQL database credentials. If the Product Name value is Tdsql mysql, it means to query Tdsql (MySQL version) credentials.
    :param str result_output_file: Used to save results.
    :param str secret_name: Secret name used to filter result.
    :param int secret_type: 0- represents user-defined credentials, defaults to 0. 1- represents the user's cloud product credentials. 2- represents SSH key pair credentials. 3- represents cloud API key pair credentials.
    :param int state: Filter by state of secret. `0` - all secrets are queried, `1` - only Enabled secrets are queried, `2` - only Disabled secrets are queried, `3` - only PendingDelete secrets are queried.
    :param Mapping[str, Any] tags: Tags to filter secret.
    """
    ...
