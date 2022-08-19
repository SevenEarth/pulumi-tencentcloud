// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Postgresql.Inputs
{

    public sealed class InstanceBackupPlanGetArgs : Pulumi.ResourceArgs
    {
        [Input("backupPeriods")]
        private InputList<string>? _backupPeriods;

        /// <summary>
        /// List of backup period per week, available values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. NOTE: At least specify two days.
        /// </summary>
        public InputList<string> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<string>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// Specify days of the retention.
        /// </summary>
        [Input("baseBackupRetentionPeriod")]
        public Input<int>? BaseBackupRetentionPeriod { get; set; }

        /// <summary>
        /// Specify latest backup start time, format `hh:mm:ss`.
        /// </summary>
        [Input("maxBackupStartTime")]
        public Input<string>? MaxBackupStartTime { get; set; }

        /// <summary>
        /// Specify earliest backup start time, format `hh:mm:ss`.
        /// </summary>
        [Input("minBackupStartTime")]
        public Input<string>? MinBackupStartTime { get; set; }

        public InstanceBackupPlanGetArgs()
        {
        }
    }
}
