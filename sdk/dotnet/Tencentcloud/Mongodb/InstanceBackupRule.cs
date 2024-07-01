// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb
{
    /// <summary>
    /// Provides a resource to create mongodb backup rule
    /// 
    /// ## Import
    /// 
    /// mongodb backup_rule can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule backup_rule ${instanceId}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule")]
    public partial class InstanceBackupRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set automatic backup method. Valid values:
        /// - 0: Logical backup;
        /// - 1: Physical backup;
        /// - 3: Snapshot backup (supported only in cloud disk version).
        /// </summary>
        [Output("backupMethod")]
        public Output<int> BackupMethod { get; private set; } = null!;

        /// <summary>
        /// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
        /// </summary>
        [Output("backupRetentionPeriod")]
        public Output<int> BackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that backup starts at 02:00.
        /// </summary>
        [Output("backupTime")]
        public Output<int> BackupTime { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceBackupRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceBackupRule(string name, InstanceBackupRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule", name, args ?? new InstanceBackupRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceBackupRule(string name, Input<string> id, InstanceBackupRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mongodb/instanceBackupRule:InstanceBackupRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceBackupRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceBackupRule Get(string name, Input<string> id, InstanceBackupRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceBackupRule(name, id, state, options);
        }
    }

    public sealed class InstanceBackupRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set automatic backup method. Valid values:
        /// - 0: Logical backup;
        /// - 1: Physical backup;
        /// - 3: Snapshot backup (supported only in cloud disk version).
        /// </summary>
        [Input("backupMethod", required: true)]
        public Input<int> BackupMethod { get; set; } = null!;

        /// <summary>
        /// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that backup starts at 02:00.
        /// </summary>
        [Input("backupTime", required: true)]
        public Input<int> BackupTime { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public InstanceBackupRuleArgs()
        {
        }
        public static new InstanceBackupRuleArgs Empty => new InstanceBackupRuleArgs();
    }

    public sealed class InstanceBackupRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set automatic backup method. Valid values:
        /// - 0: Logical backup;
        /// - 1: Physical backup;
        /// - 3: Snapshot backup (supported only in cloud disk version).
        /// </summary>
        [Input("backupMethod")]
        public Input<int>? BackupMethod { get; set; }

        /// <summary>
        /// Specify the number of days to save backup data. The default is 7 days, and the support settings are 7, 30, 90, 180, 365.
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// Set the start time for automatic backup. The value range is: [0,23]. For example, setting this parameter to 2 means that backup starts at 02:00.
        /// </summary>
        [Input("backupTime")]
        public Input<int>? BackupTime { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public InstanceBackupRuleState()
        {
        }
        public static new InstanceBackupRuleState Empty => new InstanceBackupRuleState();
    }
}
