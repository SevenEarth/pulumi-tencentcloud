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
    /// Provides a snapshot policy resource.
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
    ///     var snapshotPolicy = new Tencentcloud.Cbs.SnapshotPolicy("snapshotPolicy", new()
    ///     {
    ///         RepeatHours = new[]
    ///         {
    ///             1,
    ///         },
    ///         RepeatWeekdays = new[]
    ///         {
    ///             1,
    ///             4,
    ///         },
    ///         RetentionDays = 7,
    ///         SnapshotPolicyName = "mysnapshotpolicyname",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// CBS snapshot policy can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy snapshot_policy asp-jliex1tn
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy")]
    public partial class SnapshotPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        /// </summary>
        [Output("repeatHours")]
        public Output<ImmutableArray<int>> RepeatHours { get; private set; } = null!;

        /// <summary>
        /// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        /// </summary>
        [Output("repeatWeekdays")]
        public Output<ImmutableArray<int>> RepeatWeekdays { get; private set; } = null!;

        /// <summary>
        /// Retention days of the snapshot, and the default value is 7.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;

        /// <summary>
        /// Name of snapshot policy. The maximum length can not exceed 60 bytes.
        /// </summary>
        [Output("snapshotPolicyName")]
        public Output<string> SnapshotPolicyName { get; private set; } = null!;


        /// <summary>
        /// Create a SnapshotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnapshotPolicy(string name, SnapshotPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy", name, args ?? new SnapshotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnapshotPolicy(string name, Input<string> id, SnapshotPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnapshotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnapshotPolicy Get(string name, Input<string> id, SnapshotPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SnapshotPolicy(name, id, state, options);
        }
    }

    public sealed class SnapshotPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("repeatHours", required: true)]
        private InputList<int>? _repeatHours;

        /// <summary>
        /// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        /// </summary>
        public InputList<int> RepeatHours
        {
            get => _repeatHours ?? (_repeatHours = new InputList<int>());
            set => _repeatHours = value;
        }

        [Input("repeatWeekdays", required: true)]
        private InputList<int>? _repeatWeekdays;

        /// <summary>
        /// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        /// </summary>
        public InputList<int> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<int>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// Retention days of the snapshot, and the default value is 7.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// Name of snapshot policy. The maximum length can not exceed 60 bytes.
        /// </summary>
        [Input("snapshotPolicyName", required: true)]
        public Input<string> SnapshotPolicyName { get; set; } = null!;

        public SnapshotPolicyArgs()
        {
        }
        public static new SnapshotPolicyArgs Empty => new SnapshotPolicyArgs();
    }

    public sealed class SnapshotPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("repeatHours")]
        private InputList<int>? _repeatHours;

        /// <summary>
        /// Trigger times of periodic snapshot. Valid value ranges: (0~23). The 0 means 00:00, and so on.
        /// </summary>
        public InputList<int> RepeatHours
        {
            get => _repeatHours ?? (_repeatHours = new InputList<int>());
            set => _repeatHours = value;
        }

        [Input("repeatWeekdays")]
        private InputList<int>? _repeatWeekdays;

        /// <summary>
        /// Periodic snapshot is enabled. Valid values: [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.
        /// </summary>
        public InputList<int> RepeatWeekdays
        {
            get => _repeatWeekdays ?? (_repeatWeekdays = new InputList<int>());
            set => _repeatWeekdays = value;
        }

        /// <summary>
        /// Retention days of the snapshot, and the default value is 7.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        /// <summary>
        /// Name of snapshot policy. The maximum length can not exceed 60 bytes.
        /// </summary>
        [Input("snapshotPolicyName")]
        public Input<string>? SnapshotPolicyName { get; set; }

        public SnapshotPolicyState()
        {
        }
        public static new SnapshotPolicyState Empty => new SnapshotPolicyState();
    }
}
