// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    public static class GetDatabaseObjects
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb database_objects
        /// </summary>
        public static Task<GetDatabaseObjectsResult> InvokeAsync(GetDatabaseObjectsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseObjectsResult>("tencentcloud:Dcdb/getDatabaseObjects:getDatabaseObjects", args ?? new GetDatabaseObjectsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb database_objects
        /// </summary>
        public static Output<GetDatabaseObjectsResult> Invoke(GetDatabaseObjectsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseObjectsResult>("tencentcloud:Dcdb/getDatabaseObjects:getDatabaseObjects", args ?? new GetDatabaseObjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseObjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database name, obtained through the DescribeDatabases api.
        /// </summary>
        [Input("dbName", required: true)]
        public string DbName { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDatabaseObjectsArgs()
        {
        }
        public static new GetDatabaseObjectsArgs Empty => new GetDatabaseObjectsArgs();
    }

    public sealed class GetDatabaseObjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database name, obtained through the DescribeDatabases api.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDatabaseObjectsInvokeArgs()
        {
        }
        public static new GetDatabaseObjectsInvokeArgs Empty => new GetDatabaseObjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseObjectsResult
    {
        public readonly string DbName;
        /// <summary>
        /// Function list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseObjectsFuncResult> Funcs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Procedure list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseObjectsProcResult> Procs;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Table list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseObjectsTableResult> Tables;
        /// <summary>
        /// View list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseObjectsViewResult> Views;

        [OutputConstructor]
        private GetDatabaseObjectsResult(
            string dbName,

            ImmutableArray<Outputs.GetDatabaseObjectsFuncResult> funcs,

            string id,

            string instanceId,

            ImmutableArray<Outputs.GetDatabaseObjectsProcResult> procs,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetDatabaseObjectsTableResult> tables,

            ImmutableArray<Outputs.GetDatabaseObjectsViewResult> views)
        {
            DbName = dbName;
            Funcs = funcs;
            Id = id;
            InstanceId = instanceId;
            Procs = procs;
            ResultOutputFile = resultOutputFile;
            Tables = tables;
            Views = views;
        }
    }
}
