// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts
{
    /// <summary>
    /// Provides a resource to create a dts sync_job_pause_operation
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var syncJobPauseOperation = new Tencentcloud.Dts.SyncJobPauseOperation("syncJobPauseOperation", new()
    ///     {
    ///         JobId = "sync-werwfs23",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dts/syncJobPauseOperation:SyncJobPauseOperation")]
    public partial class SyncJobPauseOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;


        /// <summary>
        /// Create a SyncJobPauseOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncJobPauseOperation(string name, SyncJobPauseOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobPauseOperation:SyncJobPauseOperation", name, args ?? new SyncJobPauseOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncJobPauseOperation(string name, Input<string> id, SyncJobPauseOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobPauseOperation:SyncJobPauseOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyncJobPauseOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncJobPauseOperation Get(string name, Input<string> id, SyncJobPauseOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncJobPauseOperation(name, id, state, options);
        }
    }

    public sealed class SyncJobPauseOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        public SyncJobPauseOperationArgs()
        {
        }
        public static new SyncJobPauseOperationArgs Empty => new SyncJobPauseOperationArgs();
    }

    public sealed class SyncJobPauseOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        public SyncJobPauseOperationState()
        {
        }
        public static new SyncJobPauseOperationState Empty => new SyncJobPauseOperationState();
    }
}
