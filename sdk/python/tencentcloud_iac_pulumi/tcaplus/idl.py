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

__all__ = ['IdlArgs', 'Idl']

@pulumi.input_type
class IdlArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 file_content: pulumi.Input[str],
                 file_ext_type: pulumi.Input[str],
                 file_name: pulumi.Input[str],
                 file_type: pulumi.Input[str],
                 tablegroup_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Idl resource.
        :param pulumi.Input[str] cluster_id: ID of the TcaplusDB cluster to which the table group belongs.
        :param pulumi.Input[str] file_content: IDL file content of the TcaplusDB table.
        :param pulumi.Input[str] file_ext_type: File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        :param pulumi.Input[str] file_name: Name of the IDL file.
        :param pulumi.Input[str] file_type: Type of the IDL file. Valid values are PROTO and TDR.
        :param pulumi.Input[str] tablegroup_id: ID of the table group to which the IDL file belongs.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "file_content", file_content)
        pulumi.set(__self__, "file_ext_type", file_ext_type)
        pulumi.set(__self__, "file_name", file_name)
        pulumi.set(__self__, "file_type", file_type)
        pulumi.set(__self__, "tablegroup_id", tablegroup_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        ID of the TcaplusDB cluster to which the table group belongs.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Input[str]:
        """
        IDL file content of the TcaplusDB table.
        """
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_content", value)

    @property
    @pulumi.getter(name="fileExtType")
    def file_ext_type(self) -> pulumi.Input[str]:
        """
        File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        """
        return pulumi.get(self, "file_ext_type")

    @file_ext_type.setter
    def file_ext_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_ext_type", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Input[str]:
        """
        Name of the IDL file.
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_name", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> pulumi.Input[str]:
        """
        Type of the IDL file. Valid values are PROTO and TDR.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter(name="tablegroupId")
    def tablegroup_id(self) -> pulumi.Input[str]:
        """
        ID of the table group to which the IDL file belongs.
        """
        return pulumi.get(self, "tablegroup_id")

    @tablegroup_id.setter
    def tablegroup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tablegroup_id", value)


@pulumi.input_type
class _IdlState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 file_ext_type: Optional[pulumi.Input[str]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 table_infos: Optional[pulumi.Input[Sequence[pulumi.Input['IdlTableInfoArgs']]]] = None,
                 tablegroup_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Idl resources.
        :param pulumi.Input[str] cluster_id: ID of the TcaplusDB cluster to which the table group belongs.
        :param pulumi.Input[str] file_content: IDL file content of the TcaplusDB table.
        :param pulumi.Input[str] file_ext_type: File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        :param pulumi.Input[str] file_name: Name of the IDL file.
        :param pulumi.Input[str] file_type: Type of the IDL file. Valid values are PROTO and TDR.
        :param pulumi.Input[Sequence[pulumi.Input['IdlTableInfoArgs']]] table_infos: Table info of the IDL.
        :param pulumi.Input[str] tablegroup_id: ID of the table group to which the IDL file belongs.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if file_content is not None:
            pulumi.set(__self__, "file_content", file_content)
        if file_ext_type is not None:
            pulumi.set(__self__, "file_ext_type", file_ext_type)
        if file_name is not None:
            pulumi.set(__self__, "file_name", file_name)
        if file_type is not None:
            pulumi.set(__self__, "file_type", file_type)
        if table_infos is not None:
            pulumi.set(__self__, "table_infos", table_infos)
        if tablegroup_id is not None:
            pulumi.set(__self__, "tablegroup_id", tablegroup_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the TcaplusDB cluster to which the table group belongs.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> Optional[pulumi.Input[str]]:
        """
        IDL file content of the TcaplusDB table.
        """
        return pulumi.get(self, "file_content")

    @file_content.setter
    def file_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_content", value)

    @property
    @pulumi.getter(name="fileExtType")
    def file_ext_type(self) -> Optional[pulumi.Input[str]]:
        """
        File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        """
        return pulumi.get(self, "file_ext_type")

    @file_ext_type.setter
    def file_ext_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_ext_type", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDL file.
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_name", value)

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the IDL file. Valid values are PROTO and TDR.
        """
        return pulumi.get(self, "file_type")

    @file_type.setter
    def file_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_type", value)

    @property
    @pulumi.getter(name="tableInfos")
    def table_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdlTableInfoArgs']]]]:
        """
        Table info of the IDL.
        """
        return pulumi.get(self, "table_infos")

    @table_infos.setter
    def table_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdlTableInfoArgs']]]]):
        pulumi.set(self, "table_infos", value)

    @property
    @pulumi.getter(name="tablegroupId")
    def tablegroup_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the table group to which the IDL file belongs.
        """
        return pulumi.get(self, "tablegroup_id")

    @tablegroup_id.setter
    def tablegroup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tablegroup_id", value)


class Idl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 file_ext_type: Optional[pulumi.Input[str]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 tablegroup_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Use this resource to create TcaplusDB IDL file.

        ## Example Usage

        ### Create a tcaplus database idl file

        The file will be with a specified cluster and tablegroup.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-3"
        vpc = tencentcloud.Vpc.get_subnets(is_default=True,
            availability_zone=availability_zone)
        vpc_id = vpc.instance_lists[0].vpc_id
        subnet_id = vpc.instance_lists[0].subnet_id
        example_cluster = tencentcloud.tcaplus.Cluster("exampleCluster",
            idl_type="PROTO",
            cluster_name="tf_example_tcaplus_cluster",
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            password="your_pw_123111",
            old_password_expire_last=3600)
        example_tablegroup = tencentcloud.tcaplus.Tablegroup("exampleTablegroup",
            cluster_id=example_cluster.id,
            tablegroup_name="tf_example_group_name")
        main = tencentcloud.tcaplus.Idl("main",
            cluster_id=example_cluster.id,
            tablegroup_id=example_tablegroup.id,
            file_name="tf_example_tcaplus_idl",
            file_type="PROTO",
            file_ext_type="proto",
            file_content=\"\"\"    syntax = "proto2";
            package myTcaplusTable;
            import "tcaplusservice.optionv1.proto";
            message tb_online {
                option(tcaplusservice.tcaplus_primary_key) = "uin,name,region";
                required int64 uin = 1;
                required string name = 2;
                required int32 region = 3;
                required int32 gamesvrid = 4;
                optional int32 logintime = 5 [default = 1];
                repeated int64 lockid = 6 [packed = true];
                optional bool is_available = 7 [default = false];
                optional pay_info pay = 8;
            }

            message pay_info {
                required int64 pay_id = 1;
                optional uint64 total_money = 2;
                optional uint64 pay_times = 3;
                optional pay_auth_info auth = 4;
                message pay_auth_info {
                    required string pay_keys = 1;
                    optional int64 update_time = 2;
                }
            }
        \"\"\")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: ID of the TcaplusDB cluster to which the table group belongs.
        :param pulumi.Input[str] file_content: IDL file content of the TcaplusDB table.
        :param pulumi.Input[str] file_ext_type: File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        :param pulumi.Input[str] file_name: Name of the IDL file.
        :param pulumi.Input[str] file_type: Type of the IDL file. Valid values are PROTO and TDR.
        :param pulumi.Input[str] tablegroup_id: ID of the table group to which the IDL file belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to create TcaplusDB IDL file.

        ## Example Usage

        ### Create a tcaplus database idl file

        The file will be with a specified cluster and tablegroup.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tencentcloud as tencentcloud
        import tencentcloud_iac_pulumi as tencentcloud

        config = pulumi.Config()
        availability_zone = config.get("availabilityZone")
        if availability_zone is None:
            availability_zone = "ap-guangzhou-3"
        vpc = tencentcloud.Vpc.get_subnets(is_default=True,
            availability_zone=availability_zone)
        vpc_id = vpc.instance_lists[0].vpc_id
        subnet_id = vpc.instance_lists[0].subnet_id
        example_cluster = tencentcloud.tcaplus.Cluster("exampleCluster",
            idl_type="PROTO",
            cluster_name="tf_example_tcaplus_cluster",
            vpc_id=vpc_id,
            subnet_id=subnet_id,
            password="your_pw_123111",
            old_password_expire_last=3600)
        example_tablegroup = tencentcloud.tcaplus.Tablegroup("exampleTablegroup",
            cluster_id=example_cluster.id,
            tablegroup_name="tf_example_group_name")
        main = tencentcloud.tcaplus.Idl("main",
            cluster_id=example_cluster.id,
            tablegroup_id=example_tablegroup.id,
            file_name="tf_example_tcaplus_idl",
            file_type="PROTO",
            file_ext_type="proto",
            file_content=\"\"\"    syntax = "proto2";
            package myTcaplusTable;
            import "tcaplusservice.optionv1.proto";
            message tb_online {
                option(tcaplusservice.tcaplus_primary_key) = "uin,name,region";
                required int64 uin = 1;
                required string name = 2;
                required int32 region = 3;
                required int32 gamesvrid = 4;
                optional int32 logintime = 5 [default = 1];
                repeated int64 lockid = 6 [packed = true];
                optional bool is_available = 7 [default = false];
                optional pay_info pay = 8;
            }

            message pay_info {
                required int64 pay_id = 1;
                optional uint64 total_money = 2;
                optional uint64 pay_times = 3;
                optional pay_auth_info auth = 4;
                message pay_auth_info {
                    required string pay_keys = 1;
                    optional int64 update_time = 2;
                }
            }
        \"\"\")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param IdlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 file_content: Optional[pulumi.Input[str]] = None,
                 file_ext_type: Optional[pulumi.Input[str]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 file_type: Optional[pulumi.Input[str]] = None,
                 tablegroup_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdlArgs.__new__(IdlArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if file_content is None and not opts.urn:
                raise TypeError("Missing required property 'file_content'")
            __props__.__dict__["file_content"] = file_content
            if file_ext_type is None and not opts.urn:
                raise TypeError("Missing required property 'file_ext_type'")
            __props__.__dict__["file_ext_type"] = file_ext_type
            if file_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_name'")
            __props__.__dict__["file_name"] = file_name
            if file_type is None and not opts.urn:
                raise TypeError("Missing required property 'file_type'")
            __props__.__dict__["file_type"] = file_type
            if tablegroup_id is None and not opts.urn:
                raise TypeError("Missing required property 'tablegroup_id'")
            __props__.__dict__["tablegroup_id"] = tablegroup_id
            __props__.__dict__["table_infos"] = None
        super(Idl, __self__).__init__(
            'tencentcloud:Tcaplus/idl:Idl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            file_content: Optional[pulumi.Input[str]] = None,
            file_ext_type: Optional[pulumi.Input[str]] = None,
            file_name: Optional[pulumi.Input[str]] = None,
            file_type: Optional[pulumi.Input[str]] = None,
            table_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdlTableInfoArgs']]]]] = None,
            tablegroup_id: Optional[pulumi.Input[str]] = None) -> 'Idl':
        """
        Get an existing Idl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: ID of the TcaplusDB cluster to which the table group belongs.
        :param pulumi.Input[str] file_content: IDL file content of the TcaplusDB table.
        :param pulumi.Input[str] file_ext_type: File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        :param pulumi.Input[str] file_name: Name of the IDL file.
        :param pulumi.Input[str] file_type: Type of the IDL file. Valid values are PROTO and TDR.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IdlTableInfoArgs']]]] table_infos: Table info of the IDL.
        :param pulumi.Input[str] tablegroup_id: ID of the table group to which the IDL file belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdlState.__new__(_IdlState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["file_content"] = file_content
        __props__.__dict__["file_ext_type"] = file_ext_type
        __props__.__dict__["file_name"] = file_name
        __props__.__dict__["file_type"] = file_type
        __props__.__dict__["table_infos"] = table_infos
        __props__.__dict__["tablegroup_id"] = tablegroup_id
        return Idl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        ID of the TcaplusDB cluster to which the table group belongs.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="fileContent")
    def file_content(self) -> pulumi.Output[str]:
        """
        IDL file content of the TcaplusDB table.
        """
        return pulumi.get(self, "file_content")

    @property
    @pulumi.getter(name="fileExtType")
    def file_ext_type(self) -> pulumi.Output[str]:
        """
        File ext type of the IDL file. If `file_type` is `PROTO`, `file_ext_type` must be 'proto'; If `file_type` is `TDR`, `file_ext_type` must be 'xml'.
        """
        return pulumi.get(self, "file_ext_type")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[str]:
        """
        Name of the IDL file.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter(name="fileType")
    def file_type(self) -> pulumi.Output[str]:
        """
        Type of the IDL file. Valid values are PROTO and TDR.
        """
        return pulumi.get(self, "file_type")

    @property
    @pulumi.getter(name="tableInfos")
    def table_infos(self) -> pulumi.Output[Sequence['outputs.IdlTableInfo']]:
        """
        Table info of the IDL.
        """
        return pulumi.get(self, "table_infos")

    @property
    @pulumi.getter(name="tablegroupId")
    def tablegroup_id(self) -> pulumi.Output[str]:
        """
        ID of the table group to which the IDL file belongs.
        """
        return pulumi.get(self, "tablegroup_id")

