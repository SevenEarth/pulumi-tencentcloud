// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus
{
    public static class GetCheckSavepoint
    {
        /// <summary>
        /// Use this data source to query detailed information of oceanus check_savepoint
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Tencentcloud.Oceanus.GetCheckSavepoint.Invoke(new()
        ///     {
        ///         JobId = "cql-314rw6w0",
        ///         RecordType = 1,
        ///         SavepointPath = "cosn://52xkpymp-12345/12345/10000/cql-12345/2/flink-savepoints/savepoint-000000-12334",
        ///         SerialId = "svp-52xkpymp",
        ///         WorkSpaceId = "space-2idq8wbr",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCheckSavepointResult> InvokeAsync(GetCheckSavepointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCheckSavepointResult>("tencentcloud:Oceanus/getCheckSavepoint:getCheckSavepoint", args ?? new GetCheckSavepointArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of oceanus check_savepoint
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Tencentcloud.Oceanus.GetCheckSavepoint.Invoke(new()
        ///     {
        ///         JobId = "cql-314rw6w0",
        ///         RecordType = 1,
        ///         SavepointPath = "cosn://52xkpymp-12345/12345/10000/cql-12345/2/flink-savepoints/savepoint-000000-12334",
        ///         SerialId = "svp-52xkpymp",
        ///         WorkSpaceId = "space-2idq8wbr",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCheckSavepointResult> Invoke(GetCheckSavepointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCheckSavepointResult>("tencentcloud:Oceanus/getCheckSavepoint:getCheckSavepoint", args ?? new GetCheckSavepointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCheckSavepointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Job id.
        /// </summary>
        [Input("jobId", required: true)]
        public string JobId { get; set; } = null!;

        /// <summary>
        /// Snapshot type. 1:savepoint; 2:checkpoint; 3:cancelWithSavepoint.
        /// </summary>
        [Input("recordType", required: true)]
        public int RecordType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Snapshot path, currently only supports COS path.
        /// </summary>
        [Input("savepointPath", required: true)]
        public string SavepointPath { get; set; } = null!;

        /// <summary>
        /// Snapshot resource ID.
        /// </summary>
        [Input("serialId", required: true)]
        public string SerialId { get; set; } = null!;

        /// <summary>
        /// Workspace ID.
        /// </summary>
        [Input("workSpaceId", required: true)]
        public string WorkSpaceId { get; set; } = null!;

        public GetCheckSavepointArgs()
        {
        }
        public static new GetCheckSavepointArgs Empty => new GetCheckSavepointArgs();
    }

    public sealed class GetCheckSavepointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Job id.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// Snapshot type. 1:savepoint; 2:checkpoint; 3:cancelWithSavepoint.
        /// </summary>
        [Input("recordType", required: true)]
        public Input<int> RecordType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Snapshot path, currently only supports COS path.
        /// </summary>
        [Input("savepointPath", required: true)]
        public Input<string> SavepointPath { get; set; } = null!;

        /// <summary>
        /// Snapshot resource ID.
        /// </summary>
        [Input("serialId", required: true)]
        public Input<string> SerialId { get; set; } = null!;

        /// <summary>
        /// Workspace ID.
        /// </summary>
        [Input("workSpaceId", required: true)]
        public Input<string> WorkSpaceId { get; set; } = null!;

        public GetCheckSavepointInvokeArgs()
        {
        }
        public static new GetCheckSavepointInvokeArgs Empty => new GetCheckSavepointInvokeArgs();
    }


    [OutputType]
    public sealed class GetCheckSavepointResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string JobId;
        public readonly int RecordType;
        public readonly string? ResultOutputFile;
        public readonly string SavepointPath;
        /// <summary>
        /// 1=available, 2=unavailable.
        /// </summary>
        public readonly int SavepointStatus;
        public readonly string SerialId;
        public readonly string WorkSpaceId;

        [OutputConstructor]
        private GetCheckSavepointResult(
            string id,

            string jobId,

            int recordType,

            string? resultOutputFile,

            string savepointPath,

            int savepointStatus,

            string serialId,

            string workSpaceId)
        {
            Id = id;
            JobId = jobId;
            RecordType = recordType;
            ResultOutputFile = resultOutputFile;
            SavepointPath = savepointPath;
            SavepointStatus = savepointStatus;
            SerialId = serialId;
            WorkSpaceId = workSpaceId;
        }
    }
}
