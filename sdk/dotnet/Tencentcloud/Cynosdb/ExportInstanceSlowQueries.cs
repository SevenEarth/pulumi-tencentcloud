// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    /// <summary>
    /// Provides a resource to create a cynosdb export_instance_slow_queries
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exportInstanceSlowQueries = new Tencentcloud.Cynosdb.ExportInstanceSlowQueries("exportInstanceSlowQueries", new Tencentcloud.Cynosdb.ExportInstanceSlowQueriesArgs
    ///         {
    ///             Database = "db1",
    ///             EndTime = "2022-01-01 14:00:00",
    ///             FileType = "csv",
    ///             Host = "10.10.10.10",
    ///             InstanceId = "cynosdbmysql-ins-123",
    ///             StartTime = "2022-01-01 12:00:00",
    ///             Username = "root",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/exportInstanceSlowQueries:ExportInstanceSlowQueries")]
    public partial class ExportInstanceSlowQueries : Pulumi.CustomResource
    {
        /// <summary>
        /// Database name.
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// Latest transaction start time.
        /// </summary>
        [Output("endTime")]
        public Output<string?> EndTime { get; private set; } = null!;

        /// <summary>
        /// Slow query export content.
        /// </summary>
        [Output("fileContent")]
        public Output<string> FileContent { get; private set; } = null!;

        /// <summary>
        /// File type, optional values: csv, original.
        /// </summary>
        [Output("fileType")]
        public Output<string?> FileType { get; private set; } = null!;

        /// <summary>
        /// Client host.
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Earliest transaction start time.
        /// </summary>
        [Output("startTime")]
        public Output<string?> StartTime { get; private set; } = null!;

        /// <summary>
        /// user name.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ExportInstanceSlowQueries resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExportInstanceSlowQueries(string name, ExportInstanceSlowQueriesArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/exportInstanceSlowQueries:ExportInstanceSlowQueries", name, args ?? new ExportInstanceSlowQueriesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExportInstanceSlowQueries(string name, Input<string> id, ExportInstanceSlowQueriesState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/exportInstanceSlowQueries:ExportInstanceSlowQueries", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExportInstanceSlowQueries resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExportInstanceSlowQueries Get(string name, Input<string> id, ExportInstanceSlowQueriesState? state = null, CustomResourceOptions? options = null)
        {
            return new ExportInstanceSlowQueries(name, id, state, options);
        }
    }

    public sealed class ExportInstanceSlowQueriesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database name.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Latest transaction start time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// File type, optional values: csv, original.
        /// </summary>
        [Input("fileType")]
        public Input<string>? FileType { get; set; }

        /// <summary>
        /// Client host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Earliest transaction start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// user name.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ExportInstanceSlowQueriesArgs()
        {
        }
    }

    public sealed class ExportInstanceSlowQueriesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database name.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Latest transaction start time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Slow query export content.
        /// </summary>
        [Input("fileContent")]
        public Input<string>? FileContent { get; set; }

        /// <summary>
        /// File type, optional values: csv, original.
        /// </summary>
        [Input("fileType")]
        public Input<string>? FileType { get; set; }

        /// <summary>
        /// Client host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Earliest transaction start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// user name.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ExportInstanceSlowQueriesState()
        {
        }
    }
}
