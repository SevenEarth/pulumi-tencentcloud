# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGetParametersForImportResult',
    'AwaitableGetGetParametersForImportResult',
    'get_get_parameters_for_import',
    'get_get_parameters_for_import_output',
]

@pulumi.output_type
class GetGetParametersForImportResult:
    """
    A collection of values returned by getGetParametersForImport.
    """
    def __init__(__self__, id=None, import_token=None, key_id=None, parameters_valid_to=None, public_key=None, result_output_file=None, wrapping_algorithm=None, wrapping_key_spec=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if import_token and not isinstance(import_token, str):
            raise TypeError("Expected argument 'import_token' to be a str")
        pulumi.set(__self__, "import_token", import_token)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if parameters_valid_to and not isinstance(parameters_valid_to, int):
            raise TypeError("Expected argument 'parameters_valid_to' to be a int")
        pulumi.set(__self__, "parameters_valid_to", parameters_valid_to)
        if public_key and not isinstance(public_key, str):
            raise TypeError("Expected argument 'public_key' to be a str")
        pulumi.set(__self__, "public_key", public_key)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if wrapping_algorithm and not isinstance(wrapping_algorithm, str):
            raise TypeError("Expected argument 'wrapping_algorithm' to be a str")
        pulumi.set(__self__, "wrapping_algorithm", wrapping_algorithm)
        if wrapping_key_spec and not isinstance(wrapping_key_spec, str):
            raise TypeError("Expected argument 'wrapping_key_spec' to be a str")
        pulumi.set(__self__, "wrapping_key_spec", wrapping_key_spec)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="importToken")
    def import_token(self) -> str:
        """
        The token required for importing key material is used as the parameter of ImportKeyMaterial.
        """
        return pulumi.get(self, "import_token")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="parametersValidTo")
    def parameters_valid_to(self) -> int:
        """
        The validity period of the exported token and public key cannot be imported after this period, and you need to call GetParametersForImport again to obtain it.
        """
        return pulumi.get(self, "parameters_valid_to")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        Base64-encoded public key content.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="wrappingAlgorithm")
    def wrapping_algorithm(self) -> str:
        return pulumi.get(self, "wrapping_algorithm")

    @property
    @pulumi.getter(name="wrappingKeySpec")
    def wrapping_key_spec(self) -> str:
        return pulumi.get(self, "wrapping_key_spec")


class AwaitableGetGetParametersForImportResult(GetGetParametersForImportResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGetParametersForImportResult(
            id=self.id,
            import_token=self.import_token,
            key_id=self.key_id,
            parameters_valid_to=self.parameters_valid_to,
            public_key=self.public_key,
            result_output_file=self.result_output_file,
            wrapping_algorithm=self.wrapping_algorithm,
            wrapping_key_spec=self.wrapping_key_spec)


def get_get_parameters_for_import(key_id: Optional[str] = None,
                                  result_output_file: Optional[str] = None,
                                  wrapping_algorithm: Optional[str] = None,
                                  wrapping_key_spec: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGetParametersForImportResult:
    """
    Use this data source to query detailed information of kms get_parameters_for_import

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Kms.get_get_parameters_for_import(key_id="786aea8c-4aec-11ee-b601-525400281a45",
        wrapping_algorithm="RSAES_OAEP_SHA_1",
        wrapping_key_spec="RSA_2048")
    ```


    :param str key_id: CMK unique identifier.
    :param str result_output_file: Used to save results.
    :param str wrapping_algorithm: Specifies the algorithm for encrypting key material, currently supports RSAES_PKCS1_V1_5, RSAES_OAEP_SHA_1, RSAES_OAEP_SHA_256.
    :param str wrapping_key_spec: Specifies the type of encryption key material, currently only supports RSA_2048.
    """
    __args__ = dict()
    __args__['keyId'] = key_id
    __args__['resultOutputFile'] = result_output_file
    __args__['wrappingAlgorithm'] = wrapping_algorithm
    __args__['wrappingKeySpec'] = wrapping_key_spec
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Kms/getGetParametersForImport:getGetParametersForImport', __args__, opts=opts, typ=GetGetParametersForImportResult).value

    return AwaitableGetGetParametersForImportResult(
        id=__ret__.id,
        import_token=__ret__.import_token,
        key_id=__ret__.key_id,
        parameters_valid_to=__ret__.parameters_valid_to,
        public_key=__ret__.public_key,
        result_output_file=__ret__.result_output_file,
        wrapping_algorithm=__ret__.wrapping_algorithm,
        wrapping_key_spec=__ret__.wrapping_key_spec)


@_utilities.lift_output_func(get_get_parameters_for_import)
def get_get_parameters_for_import_output(key_id: Optional[pulumi.Input[str]] = None,
                                         result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         wrapping_algorithm: Optional[pulumi.Input[str]] = None,
                                         wrapping_key_spec: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGetParametersForImportResult]:
    """
    Use this data source to query detailed information of kms get_parameters_for_import

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tencentcloud as tencentcloud

    example = tencentcloud.Kms.get_get_parameters_for_import(key_id="786aea8c-4aec-11ee-b601-525400281a45",
        wrapping_algorithm="RSAES_OAEP_SHA_1",
        wrapping_key_spec="RSA_2048")
    ```


    :param str key_id: CMK unique identifier.
    :param str result_output_file: Used to save results.
    :param str wrapping_algorithm: Specifies the algorithm for encrypting key material, currently supports RSAES_PKCS1_V1_5, RSAES_OAEP_SHA_1, RSAES_OAEP_SHA_256.
    :param str wrapping_key_spec: Specifies the type of encryption key material, currently only supports RSA_2048.
    """
    ...
