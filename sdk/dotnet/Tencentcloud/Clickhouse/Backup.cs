// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse
{
    /// <summary>
    /// Provides a resource to open clickhouse backup
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
    ///         var backup = new Tencentcloud.Clickhouse.Backup("backup", new Tencentcloud.Clickhouse.BackupArgs
    ///         {
    ///             CosBucketName = "xxxxxx",
    ///             InstanceId = "cdwch-xxxxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// clickhouse backup can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Clickhouse/backup:Backup backup instance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clickhouse/backup:Backup")]
    public partial class Backup : Pulumi.CustomResource
    {
        /// <summary>
        /// COS bucket name.
        /// </summary>
        [Output("cosBucketName")]
        public Output<string> CosBucketName { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/backup:Backup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// COS bucket name.
        /// </summary>
        [Input("cosBucketName", required: true)]
        public Input<string> CosBucketName { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public BackupArgs()
        {
        }
    }

    public sealed class BackupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// COS bucket name.
        /// </summary>
        [Input("cosBucketName")]
        public Input<string>? CosBucketName { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public BackupState()
        {
        }
    }
}