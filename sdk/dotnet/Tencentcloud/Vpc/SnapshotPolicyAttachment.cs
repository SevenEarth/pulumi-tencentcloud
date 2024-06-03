// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provides a resource to create a vpc snapshot_policy_attachment
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
    ///     var exampleBucket = new Tencentcloud.Cos.Bucket("exampleBucket", new()
    ///     {
    ///         CosBucket = "tf-example-1308919341",
    ///         Acl = "private",
    ///     });
    /// 
    ///     var exampleSnapshotPolicy = new Tencentcloud.Vpc.SnapshotPolicy("exampleSnapshotPolicy", new()
    ///     {
    ///         SnapshotPolicyName = "tf-example",
    ///         BackupType = "time",
    ///         CosBucket = exampleBucket.CosBucket,
    ///         CosRegion = "ap-guangzhou",
    ///         CreateNewCos = false,
    ///         KeepTime = 2,
    ///         BackupPolicies = new[]
    ///         {
    ///             new Tencentcloud.Vpc.Inputs.SnapshotPolicyBackupPolicyArgs
    ///             {
    ///                 BackupDay = "monday",
    ///                 BackupTime = "00:00:00",
    ///             },
    ///             new Tencentcloud.Vpc.Inputs.SnapshotPolicyBackupPolicyArgs
    ///             {
    ///                 BackupDay = "tuesday",
    ///                 BackupTime = "01:00:00",
    ///             },
    ///             new Tencentcloud.Vpc.Inputs.SnapshotPolicyBackupPolicyArgs
    ///             {
    ///                 BackupDay = "wednesday",
    ///                 BackupTime = "02:00:00",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleGroup = new Tencentcloud.Security.Group("exampleGroup", new()
    ///     {
    ///         Description = "desc.",
    ///     });
    /// 
    ///     var attachment = new Tencentcloud.Vpc.SnapshotPolicyAttachment("attachment", new()
    ///     {
    ///         SnapshotPolicyId = exampleSnapshotPolicy.Id,
    ///         Instances = new[]
    ///         {
    ///             new Tencentcloud.Vpc.Inputs.SnapshotPolicyAttachmentInstanceArgs
    ///             {
    ///                 InstanceType = "securitygroup",
    ///                 InstanceId = exampleGroup.Id,
    ///                 InstanceName = "tf-example",
    ///                 InstanceRegion = "ap-guangzhou",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// vpc snapshot_policy_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Vpc/snapshotPolicyAttachment:SnapshotPolicyAttachment snapshot_policy_attachment snapshot_policy_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/snapshotPolicyAttachment:SnapshotPolicyAttachment")]
    public partial class SnapshotPolicyAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Associated instance information.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<Outputs.SnapshotPolicyAttachmentInstance>> Instances { get; private set; } = null!;

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Output("snapshotPolicyId")]
        public Output<string> SnapshotPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotPolicyAttachment(string name, SnapshotPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/snapshotPolicyAttachment:SnapshotPolicyAttachment", name, args ?? new SnapshotPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotPolicyAttachment(string name, Input<string> id, SnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/snapshotPolicyAttachment:SnapshotPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotPolicyAttachment Get(string name, Input<string> id, SnapshotPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class SnapshotPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("instances", required: true)]
        private InputList<Inputs.SnapshotPolicyAttachmentInstanceArgs>? _instances;

        /// <summary>
        /// Associated instance information.
        /// </summary>
        public InputList<Inputs.SnapshotPolicyAttachmentInstanceArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.SnapshotPolicyAttachmentInstanceArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Input("snapshotPolicyId", required: true)]
        public Input<string> SnapshotPolicyId { get; set; } = null!;

        public SnapshotPolicyAttachmentArgs()
        {
        }
        public static new SnapshotPolicyAttachmentArgs Empty => new SnapshotPolicyAttachmentArgs();
    }

    public sealed class SnapshotPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        [Input("instances")]
        private InputList<Inputs.SnapshotPolicyAttachmentInstanceGetArgs>? _instances;

        /// <summary>
        /// Associated instance information.
        /// </summary>
        public InputList<Inputs.SnapshotPolicyAttachmentInstanceGetArgs> Instances
        {
            get => _instances ?? (_instances = new InputList<Inputs.SnapshotPolicyAttachmentInstanceGetArgs>());
            set => _instances = value;
        }

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Input("snapshotPolicyId")]
        public Input<string>? SnapshotPolicyId { get; set; }

        public SnapshotPolicyAttachmentState()
        {
        }
        public static new SnapshotPolicyAttachmentState Empty => new SnapshotPolicyAttachmentState();
    }
}
