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
    /// Provides a resource to create a cfs auto_snapshot_policy
    /// 
    /// ## Example Usage
    /// ### use day of week
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var autoSnapshotPolicy = new Tencentcloud.Cfs.AutoSnapshotPolicy("autoSnapshotPolicy", new Tencentcloud.Cfs.AutoSnapshotPolicyArgs
    ///         {
    ///             AliveDays = 7,
    ///             DayOfWeek = "1,2",
    ///             Hour = "2,3",
    ///             PolicyName = "policy_name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### use day of month
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var autoSnapshotPolicy = new Tencentcloud.Cfs.AutoSnapshotPolicy("autoSnapshotPolicy", new Tencentcloud.Cfs.AutoSnapshotPolicyArgs
    ///         {
    ///             AliveDays = 7,
    ///             DayOfMonth = "2,3,4",
    ///             Hour = "2,3",
    ///             PolicyName = "policy_name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### use interval days
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var autoSnapshotPolicy = new Tencentcloud.Cfs.AutoSnapshotPolicy("autoSnapshotPolicy", new Tencentcloud.Cfs.AutoSnapshotPolicyArgs
    ///         {
    ///             AliveDays = 7,
    ///             Hour = "2,3",
    ///             IntervalDays = 1,
    ///             PolicyName = "policy_name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cfs auto_snapshot_policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy auto_snapshot_policy auto_snapshot_policy_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy")]
    public partial class AutoSnapshotPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Snapshot retention period.
        /// </summary>
        [Output("aliveDays")]
        public Output<int?> AliveDays { get; private set; } = null!;

        /// <summary>
        /// The specific day (day 1 to day 31) of the month on which to create a snapshot.
        /// </summary>
        [Output("dayOfMonth")]
        public Output<string?> DayOfMonth { get; private set; } = null!;

        /// <summary>
        /// The day of the week on which to repeat the snapshot operation.
        /// </summary>
        [Output("dayOfWeek")]
        public Output<string?> DayOfWeek { get; private set; } = null!;

        /// <summary>
        /// The time point when to repeat the snapshot operation.
        /// </summary>
        [Output("hour")]
        public Output<string> Hour { get; private set; } = null!;

        /// <summary>
        /// The snapshot interval, in days.
        /// </summary>
        [Output("intervalDays")]
        public Output<int?> IntervalDays { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("policyName")]
        public Output<string?> PolicyName { get; private set; } = null!;


        /// <summary>
        /// Create a AutoSnapshotPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoSnapshotPolicy(string name, AutoSnapshotPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy", name, args ?? new AutoSnapshotPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoSnapshotPolicy(string name, Input<string> id, AutoSnapshotPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/autoSnapshotPolicy:AutoSnapshotPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoSnapshotPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoSnapshotPolicy Get(string name, Input<string> id, AutoSnapshotPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoSnapshotPolicy(name, id, state, options);
        }
    }

    public sealed class AutoSnapshotPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Snapshot retention period.
        /// </summary>
        [Input("aliveDays")]
        public Input<int>? AliveDays { get; set; }

        /// <summary>
        /// The specific day (day 1 to day 31) of the month on which to create a snapshot.
        /// </summary>
        [Input("dayOfMonth")]
        public Input<string>? DayOfMonth { get; set; }

        /// <summary>
        /// The day of the week on which to repeat the snapshot operation.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<string>? DayOfWeek { get; set; }

        /// <summary>
        /// The time point when to repeat the snapshot operation.
        /// </summary>
        [Input("hour", required: true)]
        public Input<string> Hour { get; set; } = null!;

        /// <summary>
        /// The snapshot interval, in days.
        /// </summary>
        [Input("intervalDays")]
        public Input<int>? IntervalDays { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public AutoSnapshotPolicyArgs()
        {
        }
    }

    public sealed class AutoSnapshotPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Snapshot retention period.
        /// </summary>
        [Input("aliveDays")]
        public Input<int>? AliveDays { get; set; }

        /// <summary>
        /// The specific day (day 1 to day 31) of the month on which to create a snapshot.
        /// </summary>
        [Input("dayOfMonth")]
        public Input<string>? DayOfMonth { get; set; }

        /// <summary>
        /// The day of the week on which to repeat the snapshot operation.
        /// </summary>
        [Input("dayOfWeek")]
        public Input<string>? DayOfWeek { get; set; }

        /// <summary>
        /// The time point when to repeat the snapshot operation.
        /// </summary>
        [Input("hour")]
        public Input<string>? Hour { get; set; }

        /// <summary>
        /// The snapshot interval, in days.
        /// </summary>
        [Input("intervalDays")]
        public Input<int>? IntervalDays { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public AutoSnapshotPolicyState()
        {
        }
    }
}
