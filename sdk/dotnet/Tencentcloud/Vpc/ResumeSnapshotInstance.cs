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
    /// Provides a resource to create a vpc resume_snapshot_instance
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic example
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
    ///     var resumeSnapshotInstance = new Tencentcloud.Vpc.ResumeSnapshotInstance("resumeSnapshotInstance", new()
    ///     {
    ///         InstanceId = "ntrgm89v",
    ///         SnapshotFileId = "ssfile-emtabuwu2z",
    ///         SnapshotPolicyId = "sspolicy-1t6cobbv",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Complete example
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleSnapshotFiles = Tencentcloud.Vpc.GetSnapshotFiles.Invoke(new()
    ///     {
    ///         BusinessType = "securitygroup",
    ///         InstanceId = "sg-902tl7t7",
    ///         StartDate = "2022-10-10 00:00:00",
    ///         EndDate = "2023-10-30 00:00:00",
    ///     });
    /// 
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
    ///     var exampleResumeSnapshotInstance = new Tencentcloud.Vpc.ResumeSnapshotInstance("exampleResumeSnapshotInstance", new()
    ///     {
    ///         SnapshotPolicyId = exampleSnapshotPolicy.Id,
    ///         SnapshotFileId = exampleSnapshotFiles.Apply(getSnapshotFilesResult =&gt; getSnapshotFilesResult.SnapshotFileSets[0]?.SnapshotFileId),
    ///         InstanceId = "policy-1t6cob",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance")]
    public partial class ResumeSnapshotInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Snapshot file Id.
        /// </summary>
        [Output("snapshotFileId")]
        public Output<string> SnapshotFileId { get; private set; } = null!;

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Output("snapshotPolicyId")]
        public Output<string> SnapshotPolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a ResumeSnapshotInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResumeSnapshotInstance(string name, ResumeSnapshotInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance", name, args ?? new ResumeSnapshotInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResumeSnapshotInstance(string name, Input<string> id, ResumeSnapshotInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResumeSnapshotInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResumeSnapshotInstance Get(string name, Input<string> id, ResumeSnapshotInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ResumeSnapshotInstance(name, id, state, options);
        }
    }

    public sealed class ResumeSnapshotInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Snapshot file Id.
        /// </summary>
        [Input("snapshotFileId", required: true)]
        public Input<string> SnapshotFileId { get; set; } = null!;

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Input("snapshotPolicyId", required: true)]
        public Input<string> SnapshotPolicyId { get; set; } = null!;

        public ResumeSnapshotInstanceArgs()
        {
        }
        public static new ResumeSnapshotInstanceArgs Empty => new ResumeSnapshotInstanceArgs();
    }

    public sealed class ResumeSnapshotInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Snapshot file Id.
        /// </summary>
        [Input("snapshotFileId")]
        public Input<string>? SnapshotFileId { get; set; }

        /// <summary>
        /// Snapshot policy Id.
        /// </summary>
        [Input("snapshotPolicyId")]
        public Input<string>? SnapshotPolicyId { get; set; }

        public ResumeSnapshotInstanceState()
        {
        }
        public static new ResumeSnapshotInstanceState Empty => new ResumeSnapshotInstanceState();
    }
}
