// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetSnapshotFilesSnapshotFileSetResult
    {
        /// <summary>
        /// backup time.
        /// </summary>
        public readonly string BackupTime;
        /// <summary>
        /// InstanceId.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Uin of operator.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// snap shot file id.
        /// </summary>
        public readonly string SnapshotFileId;
        /// <summary>
        /// Snapshot Policy Id.
        /// </summary>
        public readonly string SnapshotPolicyId;

        [OutputConstructor]
        private GetSnapshotFilesSnapshotFileSetResult(
            string backupTime,

            string instanceId,

            string @operator,

            string snapshotFileId,

            string snapshotPolicyId)
        {
            BackupTime = backupTime;
            InstanceId = instanceId;
            Operator = @operator;
            SnapshotFileId = snapshotFileId;
            SnapshotPolicyId = snapshotPolicyId;
        }
    }
}
