// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql.Outputs
{

    [OutputType]
    public sealed class GetBackupDownloadUrlsBackupDownloadRestrictionResult
    {
        /// <summary>
        /// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
        /// </summary>
        public readonly string? IpRestrictionEffect;
        /// <summary>
        /// Whether it is allowed to download IP list of the backup files.
        /// </summary>
        public readonly ImmutableArray<string> IpSets;
        /// <summary>
        /// Type of the network restrictions for downloading backup files. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
        /// </summary>
        public readonly string? RestrictionType;
        /// <summary>
        /// Whether it is allowed to download the VPC ID list of the backup files.
        /// </summary>
        public readonly ImmutableArray<string> VpcIdSets;
        /// <summary>
        /// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
        /// </summary>
        public readonly string? VpcRestrictionEffect;

        [OutputConstructor]
        private GetBackupDownloadUrlsBackupDownloadRestrictionResult(
            string? ipRestrictionEffect,

            ImmutableArray<string> ipSets,

            string? restrictionType,

            ImmutableArray<string> vpcIdSets,

            string? vpcRestrictionEffect)
        {
            IpRestrictionEffect = ipRestrictionEffect;
            IpSets = ipSets;
            RestrictionType = restrictionType;
            VpcIdSets = vpcIdSets;
            VpcRestrictionEffect = vpcRestrictionEffect;
        }
    }
}
