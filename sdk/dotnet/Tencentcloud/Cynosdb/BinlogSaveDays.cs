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
    /// Provides a resource to create a cynosdb binlog_save_days
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
    ///         var binlogSaveDays = new Tencentcloud.Cynosdb.BinlogSaveDays("binlogSaveDays", new Tencentcloud.Cynosdb.BinlogSaveDaysArgs
    ///         {
    ///             BinlogSaveDays = 7,
    ///             ClusterId = "cynosdbmysql-123",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cynosdb binlog_save_days can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays binlog_save_days binlog_save_days_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays")]
    public partial class BinlogSaveDays : Pulumi.CustomResource
    {
        /// <summary>
        /// Binlog retention days.
        /// </summary>
        [Output("binlogSaveDays")]
        public Output<int> CynosdbBinlogSaveDays { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;


        /// <summary>
        /// Create a BinlogSaveDays resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BinlogSaveDays(string name, BinlogSaveDaysArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays", name, args ?? new BinlogSaveDaysArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BinlogSaveDays(string name, Input<string> id, BinlogSaveDaysState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/binlogSaveDays:BinlogSaveDays", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BinlogSaveDays resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BinlogSaveDays Get(string name, Input<string> id, BinlogSaveDaysState? state = null, CustomResourceOptions? options = null)
        {
            return new BinlogSaveDays(name, id, state, options);
        }
    }

    public sealed class BinlogSaveDaysArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Binlog retention days.
        /// </summary>
        [Input("binlogSaveDays", required: true)]
        public Input<int> CynosdbBinlogSaveDays { get; set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        public BinlogSaveDaysArgs()
        {
        }
    }

    public sealed class BinlogSaveDaysState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Binlog retention days.
        /// </summary>
        [Input("binlogSaveDays")]
        public Input<int>? CynosdbBinlogSaveDays { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        public BinlogSaveDaysState()
        {
        }
    }
}
