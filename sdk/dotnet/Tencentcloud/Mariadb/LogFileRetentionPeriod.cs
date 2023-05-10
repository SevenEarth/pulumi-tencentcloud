// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    /// <summary>
    /// Provides a resource to create a mariadb log_file_retention_period
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
    ///         var logFileRetentionPeriod = new Tencentcloud.Mariadb.LogFileRetentionPeriod("logFileRetentionPeriod", new Tencentcloud.Mariadb.LogFileRetentionPeriodArgs
    ///         {
    ///             Days = 8,
    ///             InstanceId = "tdsql-4pzs5b67",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mariadb log_file_retention_period can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod log_file_retention_period tdsql-4pzs5b67
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod")]
    public partial class LogFileRetentionPeriod : Pulumi.CustomResource
    {
        /// <summary>
        /// The number of days to save, cannot exceed 30.
        /// </summary>
        [Output("days")]
        public Output<int> Days { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a LogFileRetentionPeriod resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogFileRetentionPeriod(string name, LogFileRetentionPeriodArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod", name, args ?? new LogFileRetentionPeriodArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogFileRetentionPeriod(string name, Input<string> id, LogFileRetentionPeriodState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogFileRetentionPeriod resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogFileRetentionPeriod Get(string name, Input<string> id, LogFileRetentionPeriodState? state = null, CustomResourceOptions? options = null)
        {
            return new LogFileRetentionPeriod(name, id, state, options);
        }
    }

    public sealed class LogFileRetentionPeriodArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to save, cannot exceed 30.
        /// </summary>
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public LogFileRetentionPeriodArgs()
        {
        }
    }

    public sealed class LogFileRetentionPeriodState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to save, cannot exceed 30.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public LogFileRetentionPeriodState()
        {
        }
    }
}
