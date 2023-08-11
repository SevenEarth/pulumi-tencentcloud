// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    public static class GetInstance
    {
        /// <summary>
        /// Use this data source to get information about a MySQL instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var mysql = Output.Create(Tencentcloud.Mysql.GetInstance.InvokeAsync(new Tencentcloud.Mysql.GetInstanceArgs
        ///         {
        ///             MysqlId = "cdb-fitq5t9h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("tencentcloud:Mysql/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a MySQL instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var mysql = Output.Create(Tencentcloud.Mysql.GetInstance.InvokeAsync(new Tencentcloud.Mysql.GetInstanceArgs
        ///         {
        ///             MysqlId = "cdb-fitq5t9h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("tencentcloud:Mysql/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pay type of instance, valid values are `PREPAID` and `POSTPAID`.
        /// </summary>
        [Input("chargeType")]
        public string? ChargeType { get; set; }

        /// <summary>
        /// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
        /// </summary>
        [Input("engineVersion")]
        public string? EngineVersion { get; set; }

        /// <summary>
        /// Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
        /// </summary>
        [Input("initFlag")]
        public int? InitFlag { get; set; }

        /// <summary>
        /// Name of mysql instance.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        /// <summary>
        /// Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
        /// </summary>
        [Input("instanceRole")]
        public string? InstanceRole { get; set; }

        /// <summary>
        /// Number of results returned for a single request. Default is `20`, and maximum is 2000.
        /// </summary>
        [Input("limit")]
        public int? Limit { get; set; }

        /// <summary>
        /// Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
        /// </summary>
        [Input("mysqlId")]
        public string? MysqlId { get; set; }

        /// <summary>
        /// Record offset. Default is 0.
        /// </summary>
        [Input("offset")]
        public int? Offset { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `charge_type` instead. Pay type of instance, `0`: prepay, `1`: postpaid.
        /// </summary>
        [Input("payType")]
        public int? PayType { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Security groups ID of instance.
        /// </summary>
        [Input("securityGroupId")]
        public string? SecurityGroupId { get; set; }

        /// <summary>
        /// Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
        /// </summary>
        [Input("status")]
        public int? Status { get; set; }

        /// <summary>
        /// Indicates whether to query disaster recovery instances.
        /// </summary>
        [Input("withDr")]
        public int? WithDr { get; set; }

        /// <summary>
        /// Indicates whether to query master instances.
        /// </summary>
        [Input("withMaster")]
        public int? WithMaster { get; set; }

        /// <summary>
        /// Indicates whether to query read-only instances.
        /// </summary>
        [Input("withRo")]
        public int? WithRo { get; set; }

        public GetInstanceArgs()
        {
        }
    }

    public sealed class GetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pay type of instance, valid values are `PREPAID` and `POSTPAID`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
        /// </summary>
        [Input("initFlag")]
        public Input<int>? InitFlag { get; set; }

        /// <summary>
        /// Name of mysql instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
        /// </summary>
        [Input("instanceRole")]
        public Input<string>? InstanceRole { get; set; }

        /// <summary>
        /// Number of results returned for a single request. Default is `20`, and maximum is 2000.
        /// </summary>
        [Input("limit")]
        public Input<int>? Limit { get; set; }

        /// <summary>
        /// Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
        /// </summary>
        [Input("mysqlId")]
        public Input<string>? MysqlId { get; set; }

        /// <summary>
        /// Record offset. Default is 0.
        /// </summary>
        [Input("offset")]
        public Input<int>? Offset { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.36.0. Please use `charge_type` instead. Pay type of instance, `0`: prepay, `1`: postpaid.
        /// </summary>
        [Input("payType")]
        public Input<int>? PayType { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Security groups ID of instance.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Indicates whether to query disaster recovery instances.
        /// </summary>
        [Input("withDr")]
        public Input<int>? WithDr { get; set; }

        /// <summary>
        /// Indicates whether to query master instances.
        /// </summary>
        [Input("withMaster")]
        public Input<int>? WithMaster { get; set; }

        /// <summary>
        /// Indicates whether to query read-only instances.
        /// </summary>
        [Input("withRo")]
        public Input<int>? WithRo { get; set; }

        public GetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        /// <summary>
        /// Pay type of instance.
        /// </summary>
        public readonly string? ChargeType;
        /// <summary>
        /// The version number of the database engine to use. Supported versions include `5.5`/`5.6`/`5.7`/`8.0`.
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Initialization mark. Available values: `0` - Uninitialized; `1` - Initialized.
        /// </summary>
        public readonly int? InitFlag;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceInstanceListResult> InstanceLists;
        /// <summary>
        /// Name of mysql instance.
        /// </summary>
        public readonly string? InstanceName;
        /// <summary>
        /// Instance type. Supported values include: `master` - master instance, `dr` - disaster recovery instance, and `ro` - read-only instance.
        /// </summary>
        public readonly string? InstanceRole;
        public readonly int? Limit;
        /// <summary>
        /// Instance ID, such as `cdb-c1nl9rpv`. It is identical to the instance ID displayed in the database console page.
        /// </summary>
        public readonly string? MysqlId;
        public readonly int? Offset;
        /// <summary>
        /// Pay type of instance, `0`: prepaid, `1`: postpaid.
        /// </summary>
        public readonly int? PayType;
        public readonly string? ResultOutputFile;
        public readonly string? SecurityGroupId;
        /// <summary>
        /// Instance status. Available values: `0` - Creating; `1` - Running; `4` - Isolating; `5` - Isolated.
        /// </summary>
        public readonly int? Status;
        public readonly int? WithDr;
        public readonly int? WithMaster;
        public readonly int? WithRo;

        [OutputConstructor]
        private GetInstanceResult(
            string? chargeType,

            string? engineVersion,

            string id,

            int? initFlag,

            ImmutableArray<Outputs.GetInstanceInstanceListResult> instanceLists,

            string? instanceName,

            string? instanceRole,

            int? limit,

            string? mysqlId,

            int? offset,

            int? payType,

            string? resultOutputFile,

            string? securityGroupId,

            int? status,

            int? withDr,

            int? withMaster,

            int? withRo)
        {
            ChargeType = chargeType;
            EngineVersion = engineVersion;
            Id = id;
            InitFlag = initFlag;
            InstanceLists = instanceLists;
            InstanceName = instanceName;
            InstanceRole = instanceRole;
            Limit = limit;
            MysqlId = mysqlId;
            Offset = offset;
            PayType = payType;
            ResultOutputFile = resultOutputFile;
            SecurityGroupId = securityGroupId;
            Status = status;
            WithDr = withDr;
            WithMaster = withMaster;
            WithRo = withRo;
        }
    }
}
