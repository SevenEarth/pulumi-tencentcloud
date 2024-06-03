// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cbs
{
    /// <summary>
    /// Provides a resource to rollback cbs disk backup.
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
    ///     var operation = new Tencentcloud.Cbs.DiskBackupRollbackOperation("operation", new()
    ///     {
    ///         DiskBackupId = "dbp-xxx",
    ///         DiskId = "disk-xxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation")]
    public partial class DiskBackupRollbackOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud disk backup point ID.
        /// </summary>
        [Output("diskBackupId")]
        public Output<string> DiskBackupId { get; private set; } = null!;

        /// <summary>
        /// Cloud disk backup point original cloud disk ID.
        /// </summary>
        [Output("diskId")]
        public Output<string> DiskId { get; private set; } = null!;

        /// <summary>
        /// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
        /// </summary>
        [Output("isRollbackCompleted")]
        public Output<bool> IsRollbackCompleted { get; private set; } = null!;


        /// <summary>
        /// Create a DiskBackupRollbackOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskBackupRollbackOperation(string name, DiskBackupRollbackOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation", name, args ?? new DiskBackupRollbackOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DiskBackupRollbackOperation(string name, Input<string> id, DiskBackupRollbackOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskBackupRollbackOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskBackupRollbackOperation Get(string name, Input<string> id, DiskBackupRollbackOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskBackupRollbackOperation(name, id, state, options);
        }
    }

    public sealed class DiskBackupRollbackOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud disk backup point ID.
        /// </summary>
        [Input("diskBackupId", required: true)]
        public Input<string> DiskBackupId { get; set; } = null!;

        /// <summary>
        /// Cloud disk backup point original cloud disk ID.
        /// </summary>
        [Input("diskId", required: true)]
        public Input<string> DiskId { get; set; } = null!;

        public DiskBackupRollbackOperationArgs()
        {
        }
        public static new DiskBackupRollbackOperationArgs Empty => new DiskBackupRollbackOperationArgs();
    }

    public sealed class DiskBackupRollbackOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud disk backup point ID.
        /// </summary>
        [Input("diskBackupId")]
        public Input<string>? DiskBackupId { get; set; }

        /// <summary>
        /// Cloud disk backup point original cloud disk ID.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
        /// </summary>
        [Input("isRollbackCompleted")]
        public Input<bool>? IsRollbackCompleted { get; set; }

        public DiskBackupRollbackOperationState()
        {
        }
        public static new DiskBackupRollbackOperationState Empty => new DiskBackupRollbackOperationState();
    }
}
