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
    /// Provides a resource to create a dts sync_job_resume_operation
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
    ///     var syncJobResumeOperation = new Tencentcloud.Dts.SyncJobResumeOperation("syncJobResumeOperation", new()
    ///     {
    ///         JobId = "sync-werwfs23",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation")]
    public partial class SyncJobResumeOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;


        /// <summary>
        /// Create a SyncJobResumeOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncJobResumeOperation(string name, SyncJobResumeOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation", name, args ?? new SyncJobResumeOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncJobResumeOperation(string name, Input<string> id, SyncJobResumeOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dts/syncJobResumeOperation:SyncJobResumeOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyncJobResumeOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncJobResumeOperation Get(string name, Input<string> id, SyncJobResumeOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncJobResumeOperation(name, id, state, options);
        }
    }

    public sealed class SyncJobResumeOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        public SyncJobResumeOperationArgs()
        {
        }
        public static new SyncJobResumeOperationArgs Empty => new SyncJobResumeOperationArgs();
    }

    public sealed class SyncJobResumeOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Synchronization instance id (i.e. identifies a synchronization job).
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        public SyncJobResumeOperationState()
        {
        }
        public static new SyncJobResumeOperationState Empty => new SyncJobResumeOperationState();
    }
}
