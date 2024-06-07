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
from ._inputs import *

__all__ = ['UpdateCertificateInstanceOperationArgs', 'UpdateCertificateInstanceOperation']

@pulumi.input_type
class UpdateCertificateInstanceOperationArgs:
    def __init__(__self__, *,
                 old_certificate_id: pulumi.Input[str],
                 resource_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 allow_download: Optional[pulumi.Input[bool]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_private_key: Optional[pulumi.Input[str]] = None,
                 certificate_public_key: Optional[pulumi.Input[str]] = None,
                 expiring_notification_switch: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repeatable: Optional[pulumi.Input[bool]] = None,
                 resource_types_regions: Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]] = None):
        """
        The set of arguments for constructing a UpdateCertificateInstanceOperation resource.
        :param pulumi.Input[str] old_certificate_id: Update the original certificate ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        :param pulumi.Input[bool] allow_download: Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[str] certificate_id: Update new certificate ID.
        :param pulumi.Input[str] certificate_private_key: Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[str] certificate_public_key: Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[int] expiring_notification_switch: Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        :param pulumi.Input[int] project_id: Project ID, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[bool] repeatable: Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]] resource_types_regions: List of regions where cloud resources need to be deploye.
        """
        pulumi.set(__self__, "old_certificate_id", old_certificate_id)
        pulumi.set(__self__, "resource_types", resource_types)
        if allow_download is not None:
            pulumi.set(__self__, "allow_download", allow_download)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_private_key is not None:
            pulumi.set(__self__, "certificate_private_key", certificate_private_key)
        if certificate_public_key is not None:
            pulumi.set(__self__, "certificate_public_key", certificate_public_key)
        if expiring_notification_switch is not None:
            pulumi.set(__self__, "expiring_notification_switch", expiring_notification_switch)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repeatable is not None:
            pulumi.set(__self__, "repeatable", repeatable)
        if resource_types_regions is not None:
            pulumi.set(__self__, "resource_types_regions", resource_types_regions)

    @property
    @pulumi.getter(name="oldCertificateId")
    def old_certificate_id(self) -> pulumi.Input[str]:
        """
        Update the original certificate ID.
        """
        return pulumi.get(self, "old_certificate_id")

    @old_certificate_id.setter
    def old_certificate_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "old_certificate_id", value)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        """
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "resource_types", value)

    @property
    @pulumi.getter(name="allowDownload")
    def allow_download(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "allow_download")

    @allow_download.setter
    def allow_download(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_download", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        Update new certificate ID.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificatePrivateKey")
    def certificate_private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_private_key")

    @certificate_private_key.setter
    def certificate_private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_private_key", value)

    @property
    @pulumi.getter(name="certificatePublicKey")
    def certificate_public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_public_key")

    @certificate_public_key.setter
    def certificate_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_public_key", value)

    @property
    @pulumi.getter(name="expiringNotificationSwitch")
    def expiring_notification_switch(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        """
        return pulumi.get(self, "expiring_notification_switch")

    @expiring_notification_switch.setter
    def expiring_notification_switch(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiring_notification_switch", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def repeatable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "repeatable")

    @repeatable.setter
    def repeatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "repeatable", value)

    @property
    @pulumi.getter(name="resourceTypesRegions")
    def resource_types_regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]:
        """
        List of regions where cloud resources need to be deploye.
        """
        return pulumi.get(self, "resource_types_regions")

    @resource_types_regions.setter
    def resource_types_regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]):
        pulumi.set(self, "resource_types_regions", value)


@pulumi.input_type
class _UpdateCertificateInstanceOperationState:
    def __init__(__self__, *,
                 allow_download: Optional[pulumi.Input[bool]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_private_key: Optional[pulumi.Input[str]] = None,
                 certificate_public_key: Optional[pulumi.Input[str]] = None,
                 expiring_notification_switch: Optional[pulumi.Input[int]] = None,
                 old_certificate_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repeatable: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_types_regions: Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]] = None):
        """
        Input properties used for looking up and filtering UpdateCertificateInstanceOperation resources.
        :param pulumi.Input[bool] allow_download: Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[str] certificate_id: Update new certificate ID.
        :param pulumi.Input[str] certificate_private_key: Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[str] certificate_public_key: Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[int] expiring_notification_switch: Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        :param pulumi.Input[str] old_certificate_id: Update the original certificate ID.
        :param pulumi.Input[int] project_id: Project ID, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[bool] repeatable: Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        :param pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]] resource_types_regions: List of regions where cloud resources need to be deploye.
        """
        if allow_download is not None:
            pulumi.set(__self__, "allow_download", allow_download)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_private_key is not None:
            pulumi.set(__self__, "certificate_private_key", certificate_private_key)
        if certificate_public_key is not None:
            pulumi.set(__self__, "certificate_public_key", certificate_public_key)
        if expiring_notification_switch is not None:
            pulumi.set(__self__, "expiring_notification_switch", expiring_notification_switch)
        if old_certificate_id is not None:
            pulumi.set(__self__, "old_certificate_id", old_certificate_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repeatable is not None:
            pulumi.set(__self__, "repeatable", repeatable)
        if resource_types is not None:
            pulumi.set(__self__, "resource_types", resource_types)
        if resource_types_regions is not None:
            pulumi.set(__self__, "resource_types_regions", resource_types_regions)

    @property
    @pulumi.getter(name="allowDownload")
    def allow_download(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "allow_download")

    @allow_download.setter
    def allow_download(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_download", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        Update new certificate ID.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificatePrivateKey")
    def certificate_private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_private_key")

    @certificate_private_key.setter
    def certificate_private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_private_key", value)

    @property
    @pulumi.getter(name="certificatePublicKey")
    def certificate_public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_public_key")

    @certificate_public_key.setter
    def certificate_public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_public_key", value)

    @property
    @pulumi.getter(name="expiringNotificationSwitch")
    def expiring_notification_switch(self) -> Optional[pulumi.Input[int]]:
        """
        Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        """
        return pulumi.get(self, "expiring_notification_switch")

    @expiring_notification_switch.setter
    def expiring_notification_switch(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiring_notification_switch", value)

    @property
    @pulumi.getter(name="oldCertificateId")
    def old_certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        Update the original certificate ID.
        """
        return pulumi.get(self, "old_certificate_id")

    @old_certificate_id.setter
    def old_certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "old_certificate_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[int]]:
        """
        Project ID, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def repeatable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "repeatable")

    @repeatable.setter
    def repeatable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "repeatable", value)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        """
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_types", value)

    @property
    @pulumi.getter(name="resourceTypesRegions")
    def resource_types_regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]:
        """
        List of regions where cloud resources need to be deploye.
        """
        return pulumi.get(self, "resource_types_regions")

    @resource_types_regions.setter
    def resource_types_regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]):
        pulumi.set(self, "resource_types_regions", value)


class UpdateCertificateInstanceOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_download: Optional[pulumi.Input[bool]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_private_key: Optional[pulumi.Input[str]] = None,
                 certificate_public_key: Optional[pulumi.Input[str]] = None,
                 expiring_notification_switch: Optional[pulumi.Input[int]] = None,
                 old_certificate_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repeatable: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_types_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a ssl update_certificate_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        update_certificate_instance = tencentcloud.ssl.UpdateCertificateInstanceOperation("updateCertificateInstance",
            certificate_id="8x1eUSSl",
            old_certificate_id="8xNdi2ig",
            resource_types=["cdn"])
        ```
        <!--End PulumiCodeChooser -->

        ### Upload certificate

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        update_certificate_instance = tencentcloud.ssl.UpdateCertificateInstanceOperation("updateCertificateInstance",
            old_certificate_id="xxx",
            certificate_public_key=(lambda path: open(path).read())("xxx.crt"),
            certificate_private_key=(lambda path: open(path).read())("xxx.key"),
            repeatable=True,
            resource_types=["cdn"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_download: Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[str] certificate_id: Update new certificate ID.
        :param pulumi.Input[str] certificate_private_key: Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[str] certificate_public_key: Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[int] expiring_notification_switch: Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        :param pulumi.Input[str] old_certificate_id: Update the original certificate ID.
        :param pulumi.Input[int] project_id: Project ID, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[bool] repeatable: Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]] resource_types_regions: List of regions where cloud resources need to be deploye.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UpdateCertificateInstanceOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a ssl update_certificate_instance

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        update_certificate_instance = tencentcloud.ssl.UpdateCertificateInstanceOperation("updateCertificateInstance",
            certificate_id="8x1eUSSl",
            old_certificate_id="8xNdi2ig",
            resource_types=["cdn"])
        ```
        <!--End PulumiCodeChooser -->

        ### Upload certificate

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        update_certificate_instance = tencentcloud.ssl.UpdateCertificateInstanceOperation("updateCertificateInstance",
            old_certificate_id="xxx",
            certificate_public_key=(lambda path: open(path).read())("xxx.crt"),
            certificate_private_key=(lambda path: open(path).read())("xxx.key"),
            repeatable=True,
            resource_types=["cdn"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param UpdateCertificateInstanceOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UpdateCertificateInstanceOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_download: Optional[pulumi.Input[bool]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_private_key: Optional[pulumi.Input[str]] = None,
                 certificate_public_key: Optional[pulumi.Input[str]] = None,
                 expiring_notification_switch: Optional[pulumi.Input[int]] = None,
                 old_certificate_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 repeatable: Optional[pulumi.Input[bool]] = None,
                 resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_types_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UpdateCertificateInstanceOperationArgs.__new__(UpdateCertificateInstanceOperationArgs)

            __props__.__dict__["allow_download"] = allow_download
            __props__.__dict__["certificate_id"] = certificate_id
            __props__.__dict__["certificate_private_key"] = None if certificate_private_key is None else pulumi.Output.secret(certificate_private_key)
            __props__.__dict__["certificate_public_key"] = None if certificate_public_key is None else pulumi.Output.secret(certificate_public_key)
            __props__.__dict__["expiring_notification_switch"] = expiring_notification_switch
            if old_certificate_id is None and not opts.urn:
                raise TypeError("Missing required property 'old_certificate_id'")
            __props__.__dict__["old_certificate_id"] = old_certificate_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repeatable"] = repeatable
            if resource_types is None and not opts.urn:
                raise TypeError("Missing required property 'resource_types'")
            __props__.__dict__["resource_types"] = resource_types
            __props__.__dict__["resource_types_regions"] = resource_types_regions
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["certificatePrivateKey", "certificatePublicKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(UpdateCertificateInstanceOperation, __self__).__init__(
            'tencentcloud:Ssl/updateCertificateInstanceOperation:UpdateCertificateInstanceOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_download: Optional[pulumi.Input[bool]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            certificate_private_key: Optional[pulumi.Input[str]] = None,
            certificate_public_key: Optional[pulumi.Input[str]] = None,
            expiring_notification_switch: Optional[pulumi.Input[int]] = None,
            old_certificate_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            repeatable: Optional[pulumi.Input[bool]] = None,
            resource_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_types_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]]] = None) -> 'UpdateCertificateInstanceOperation':
        """
        Get an existing UpdateCertificateInstanceOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_download: Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[str] certificate_id: Update new certificate ID.
        :param pulumi.Input[str] certificate_private_key: Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[str] certificate_public_key: Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        :param pulumi.Input[int] expiring_notification_switch: Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        :param pulumi.Input[str] old_certificate_id: Update the original certificate ID.
        :param pulumi.Input[int] project_id: Project ID, if you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[bool] repeatable: Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_types: The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UpdateCertificateInstanceOperationResourceTypesRegionArgs']]]] resource_types_regions: List of regions where cloud resources need to be deploye.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UpdateCertificateInstanceOperationState.__new__(_UpdateCertificateInstanceOperationState)

        __props__.__dict__["allow_download"] = allow_download
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["certificate_private_key"] = certificate_private_key
        __props__.__dict__["certificate_public_key"] = certificate_public_key
        __props__.__dict__["expiring_notification_switch"] = expiring_notification_switch
        __props__.__dict__["old_certificate_id"] = old_certificate_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repeatable"] = repeatable
        __props__.__dict__["resource_types"] = resource_types
        __props__.__dict__["resource_types_regions"] = resource_types_regions
        return UpdateCertificateInstanceOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowDownload")
    def allow_download(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to allow downloading, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "allow_download")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[Optional[str]]:
        """
        Update new certificate ID.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificatePrivateKey")
    def certificate_private_key(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate private key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_private_key")

    @property
    @pulumi.getter(name="certificatePublicKey")
    def certificate_public_key(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate public key. If you upload the certificate public key, CertificateId does not need to be passed.
        """
        return pulumi.get(self, "certificate_public_key")

    @property
    @pulumi.getter(name="expiringNotificationSwitch")
    def expiring_notification_switch(self) -> pulumi.Output[Optional[int]]:
        """
        Whether to ignore expiration reminders for old certificates 0: Do not ignore notifications. 1: Ignore the notification and ignore the OldCertificateId expiration reminder.
        """
        return pulumi.get(self, "expiring_notification_switch")

    @property
    @pulumi.getter(name="oldCertificateId")
    def old_certificate_id(self) -> pulumi.Output[str]:
        """
        Update the original certificate ID.
        """
        return pulumi.get(self, "old_certificate_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[int]]:
        """
        Project ID, if you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def repeatable(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the same certificate is allowed to be uploaded repeatedly. If you choose to upload the certificate, you can configure this parameter.
        """
        return pulumi.get(self, "repeatable")

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Output[Sequence[str]]:
        """
        The resource type that needs to be deployed. The parameter value is optional: clb, cdn, waf, live, ddos, teo, apigateway, vod, tke, tcb.
        """
        return pulumi.get(self, "resource_types")

    @property
    @pulumi.getter(name="resourceTypesRegions")
    def resource_types_regions(self) -> pulumi.Output[Optional[Sequence['outputs.UpdateCertificateInstanceOperationResourceTypesRegion']]]:
        """
        List of regions where cloud resources need to be deploye.
        """
        return pulumi.get(self, "resource_types_regions")

