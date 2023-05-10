// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class GetBackupSummariesItemResult
    {
        /// <summary>
        /// The number of automatic data backups for this instance.
        /// </summary>
        public readonly int AutoBackupCount;
        /// <summary>
        /// The automatic data backup capacity of this instance.
        /// </summary>
        public readonly int AutoBackupVolume;
        /// <summary>
        /// The total backup (including data backup and log backup) of the instance occupies capacity.
        /// </summary>
        public readonly int BackupVolume;
        /// <summary>
        /// The number of log backups for this instance.
        /// </summary>
        public readonly int BinlogBackupCount;
        /// <summary>
        /// The capacity of the instance log backup.
        /// </summary>
        public readonly int BinlogBackupVolume;
        /// <summary>
        /// The total number of data backups (including automatic backups and manual backups) of the instance.
        /// </summary>
        public readonly int DataBackupCount;
        /// <summary>
        /// The total data backup capacity of this instance.
        /// </summary>
        public readonly int DataBackupVolume;
        /// <summary>
        /// Instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The number of manual data backups for this instance.
        /// </summary>
        public readonly int ManualBackupCount;
        /// <summary>
        /// The capacity of manual data backup for this instance.
        /// </summary>
        public readonly int ManualBackupVolume;

        [OutputConstructor]
        private GetBackupSummariesItemResult(
            int autoBackupCount,

            int autoBackupVolume,

            int backupVolume,

            int binlogBackupCount,

            int binlogBackupVolume,

            int dataBackupCount,

            int dataBackupVolume,

            string instanceId,

            int manualBackupCount,

            int manualBackupVolume)
        {
            AutoBackupCount = autoBackupCount;
            AutoBackupVolume = autoBackupVolume;
            BackupVolume = backupVolume;
            BinlogBackupCount = binlogBackupCount;
            BinlogBackupVolume = binlogBackupVolume;
            DataBackupCount = dataBackupCount;
            DataBackupVolume = dataBackupVolume;
            InstanceId = instanceId;
            ManualBackupCount = manualBackupCount;
            ManualBackupVolume = manualBackupVolume;
        }
    }
}
