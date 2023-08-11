# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceTdeArgs', 'InstanceTde']

@pulumi.input_type
class InstanceTdeArgs:
    def __init__(__self__, *,
                 certificate_attribution: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 quote_uin: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceTde resource.
        :param pulumi.Input[str] certificate_attribution: Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] quote_uin: Other referenced main account IDs, required when CertificateAttribute is others.
        """
        pulumi.set(__self__, "certificate_attribution", certificate_attribution)
        pulumi.set(__self__, "instance_id", instance_id)
        if quote_uin is not None:
            pulumi.set(__self__, "quote_uin", quote_uin)

    @property
    @pulumi.getter(name="certificateAttribution")
    def certificate_attribution(self) -> pulumi.Input[str]:
        """
        Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        """
        return pulumi.get(self, "certificate_attribution")

    @certificate_attribution.setter
    def certificate_attribution(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_attribution", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="quoteUin")
    def quote_uin(self) -> Optional[pulumi.Input[str]]:
        """
        Other referenced main account IDs, required when CertificateAttribute is others.
        """
        return pulumi.get(self, "quote_uin")

    @quote_uin.setter
    def quote_uin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quote_uin", value)


@pulumi.input_type
class _InstanceTdeState:
    def __init__(__self__, *,
                 certificate_attribution: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 quote_uin: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceTde resources.
        :param pulumi.Input[str] certificate_attribution: Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] quote_uin: Other referenced main account IDs, required when CertificateAttribute is others.
        """
        if certificate_attribution is not None:
            pulumi.set(__self__, "certificate_attribution", certificate_attribution)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if quote_uin is not None:
            pulumi.set(__self__, "quote_uin", quote_uin)

    @property
    @pulumi.getter(name="certificateAttribution")
    def certificate_attribution(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        """
        return pulumi.get(self, "certificate_attribution")

    @certificate_attribution.setter
    def certificate_attribution(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_attribution", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="quoteUin")
    def quote_uin(self) -> Optional[pulumi.Input[str]]:
        """
        Other referenced main account IDs, required when CertificateAttribute is others.
        """
        return pulumi.get(self, "quote_uin")

    @quote_uin.setter
    def quote_uin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "quote_uin", value)


class InstanceTde(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_attribution: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 quote_uin: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to create a sqlserver instance_tde

        ## Example Usage

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
        example_basic_instance = tencentcloud.sqlserver.BasicInstance("exampleBasicInstance",
            availability_zone=zones.zones[4].name,
            charge_type="POSTPAID_BY_HOUR",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            project_id=0,
            memory=4,
            storage=100,
            cpu=2,
            machine_type="CLOUD_PREMIUM",
            maintenance_week_sets=[
                1,
                2,
                3,
            ],
            maintenance_start_time="09:00",
            maintenance_time_span=3,
            security_groups=[security_group.id],
            tags={
                "test": "test",
            })
        example_instance_tde = tencentcloud.sqlserver.InstanceTde("exampleInstanceTde",
            instance_id=example_basic_instance.id,
            certificate_attribution="self")
        ```

        ## Import

        sqlserver instance_tde can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/instanceTde:InstanceTde example mssql-farjz9tz
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_attribution: Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] quote_uin: Other referenced main account IDs, required when CertificateAttribute is others.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceTdeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a sqlserver instance_tde

        ## Example Usage

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
        example_basic_instance = tencentcloud.sqlserver.BasicInstance("exampleBasicInstance",
            availability_zone=zones.zones[4].name,
            charge_type="POSTPAID_BY_HOUR",
            vpc_id=vpc.id,
            subnet_id=subnet.id,
            project_id=0,
            memory=4,
            storage=100,
            cpu=2,
            machine_type="CLOUD_PREMIUM",
            maintenance_week_sets=[
                1,
                2,
                3,
            ],
            maintenance_start_time="09:00",
            maintenance_time_span=3,
            security_groups=[security_group.id],
            tags={
                "test": "test",
            })
        example_instance_tde = tencentcloud.sqlserver.InstanceTde("exampleInstanceTde",
            instance_id=example_basic_instance.id,
            certificate_attribution="self")
        ```

        ## Import

        sqlserver instance_tde can be imported using the id, e.g.

        ```sh
         $ pulumi import tencentcloud:Sqlserver/instanceTde:InstanceTde example mssql-farjz9tz
        ```

        :param str resource_name: The name of the resource.
        :param InstanceTdeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceTdeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_attribution: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 quote_uin: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceTdeArgs.__new__(InstanceTdeArgs)

            if certificate_attribution is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_attribution'")
            __props__.__dict__["certificate_attribution"] = certificate_attribution
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["quote_uin"] = quote_uin
        super(InstanceTde, __self__).__init__(
            'tencentcloud:Sqlserver/instanceTde:InstanceTde',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_attribution: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            quote_uin: Optional[pulumi.Input[str]] = None) -> 'InstanceTde':
        """
        Get an existing InstanceTde resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_attribution: Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[str] quote_uin: Other referenced main account IDs, required when CertificateAttribute is others.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceTdeState.__new__(_InstanceTdeState)

        __props__.__dict__["certificate_attribution"] = certificate_attribution
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["quote_uin"] = quote_uin
        return InstanceTde(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateAttribution")
    def certificate_attribution(self) -> pulumi.Output[str]:
        """
        Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
        """
        return pulumi.get(self, "certificate_attribution")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="quoteUin")
    def quote_uin(self) -> pulumi.Output[Optional[str]]:
        """
        Other referenced main account IDs, required when CertificateAttribute is others.
        """
        return pulumi.get(self, "quote_uin")

