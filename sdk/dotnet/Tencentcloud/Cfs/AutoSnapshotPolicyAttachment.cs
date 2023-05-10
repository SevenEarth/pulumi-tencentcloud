// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cfs
{
    /// <summary>
    /// Provides a resource to create a cfs auto_snapshot_policy_attachment
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
    ///         var autoSnapshotPolicyAttachment = new Tencentcloud.Cfs.AutoSnapshotPolicyAttachment("autoSnapshotPolicyAttachment", new Tencentcloud.Cfs.AutoSnapshotPolicyAttachmentArgs
    ///         {
    ///             AutoSnapshotPolicyId = "asp-basic",
    ///             FileSystemIds = "cfs-4xzkct19,cfs-iobiaxtj",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cfs auto_snapshot_policy_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cfs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment auto_snapshot_policy_attachment auto_snapshot_policy_id#file_system_ids
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cfs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment")]
    public partial class AutoSnapshotPolicyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the snapshot to be unbound.
        /// </summary>
        [Output("autoSnapshotPolicyId")]
        public Output<string> AutoSnapshotPolicyId { get; private set; } = null!;

        /// <summary>
        /// List of IDs of the file systems to be unbound, separated by comma.
        /// </summary>
        [Output("fileSystemIds")]
        public Output<string> FileSystemIds { get; private set; } = null!;


        /// <summary>
        /// Create a AutoSnapshotPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoSnapshotPolicyAttachment(string name, AutoSnapshotPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment", name, args ?? new AutoSnapshotPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoSnapshotPolicyAttachment(string name, Input<string> id, AutoSnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoSnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoSnapshotPolicyAttachment Get(string name, Input<string> id, AutoSnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoSnapshotPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class AutoSnapshotPolicyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the snapshot to be unbound.
        /// </summary>
        [Input("autoSnapshotPolicyId", required: true)]
        public Input<string> AutoSnapshotPolicyId { get; set; } = null!;

        /// <summary>
        /// List of IDs of the file systems to be unbound, separated by comma.
        /// </summary>
        [Input("fileSystemIds", required: true)]
        public Input<string> FileSystemIds { get; set; } = null!;

        public AutoSnapshotPolicyAttachmentArgs()
        {
        }
    }

    public sealed class AutoSnapshotPolicyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the snapshot to be unbound.
        /// </summary>
        [Input("autoSnapshotPolicyId")]
        public Input<string>? AutoSnapshotPolicyId { get; set; }

        /// <summary>
        /// List of IDs of the file systems to be unbound, separated by comma.
        /// </summary>
        [Input("fileSystemIds")]
        public Input<string>? FileSystemIds { get; set; }

        public AutoSnapshotPolicyAttachmentState()
        {
        }
    }
}
