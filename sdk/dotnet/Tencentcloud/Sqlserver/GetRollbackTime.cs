// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    public static class GetRollbackTime
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver rollback_time
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetRollbackTime.InvokeAsync(new Tencentcloud.Sqlserver.GetRollbackTimeArgs
        ///         {
        ///             Dbs = 
        ///             {
        ///                 "keep_pubsub_db",
        ///             },
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRollbackTimeResult> InvokeAsync(GetRollbackTimeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRollbackTimeResult>("tencentcloud:Sqlserver/getRollbackTime:getRollbackTime", args ?? new GetRollbackTimeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver rollback_time
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
        ///         var example = Output.Create(Tencentcloud.Sqlserver.GetRollbackTime.InvokeAsync(new Tencentcloud.Sqlserver.GetRollbackTimeArgs
        ///         {
        ///             Dbs = 
        ///             {
        ///                 "keep_pubsub_db",
        ///             },
        ///             InstanceId = "mssql-qelbzgwf",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRollbackTimeResult> Invoke(GetRollbackTimeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRollbackTimeResult>("tencentcloud:Sqlserver/getRollbackTime:getRollbackTime", args ?? new GetRollbackTimeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRollbackTimeArgs : Pulumi.InvokeArgs
    {
        [Input("dbs", required: true)]
        private List<string>? _dbs;

        /// <summary>
        /// List of databases to be queried.
        /// </summary>
        public List<string> Dbs
        {
            get => _dbs ?? (_dbs = new List<string>());
            set => _dbs = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRollbackTimeArgs()
        {
        }
    }

    public sealed class GetRollbackTimeInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dbs", required: true)]
        private InputList<string>? _dbs;

        /// <summary>
        /// List of databases to be queried.
        /// </summary>
        public InputList<string> Dbs
        {
            get => _dbs ?? (_dbs = new InputList<string>());
            set => _dbs = value;
        }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRollbackTimeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRollbackTimeResult
    {
        public readonly ImmutableArray<string> Dbs;
        /// <summary>
        /// Information of time range available for database rollback.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRollbackTimeDetailResult> Details;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRollbackTimeResult(
            ImmutableArray<string> dbs,

            ImmutableArray<Outputs.GetRollbackTimeDetailResult> details,

            string id,

            string instanceId,

            string? resultOutputFile)
        {
            Dbs = dbs;
            Details = details;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
