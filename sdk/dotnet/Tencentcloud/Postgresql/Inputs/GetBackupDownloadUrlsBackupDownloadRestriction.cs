// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql.Inputs
{

    public sealed class GetBackupDownloadUrlsBackupDownloadRestrictionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether IP is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
        /// </summary>
        [Input("ipRestrictionEffect")]
        public string? IpRestrictionEffect { get; set; }

        [Input("ipSets")]
        private List<string>? _ipSets;

        /// <summary>
        /// Whether it is allowed to download IP list of the backup files.
        /// </summary>
        public List<string> IpSets
        {
            get => _ipSets ?? (_ipSets = new List<string>());
            set => _ipSets = value;
        }

        /// <summary>
        /// Type of the network restrictions for downloading backup files. Valid values: `NONE` (backups can be downloaded over both private and public networks), `INTRANET` (backups can only be downloaded over the private network), `CUSTOMIZE` (backups can be downloaded over specified VPCs or at specified IPs).
        /// </summary>
        [Input("restrictionType")]
        public string? RestrictionType { get; set; }

        [Input("vpcIdSets")]
        private List<string>? _vpcIdSets;

        /// <summary>
        /// Whether it is allowed to download the VPC ID list of the backup files.
        /// </summary>
        public List<string> VpcIdSets
        {
            get => _vpcIdSets ?? (_vpcIdSets = new List<string>());
            set => _vpcIdSets = value;
        }

        /// <summary>
        /// Whether VPC is allowed. Valid values: `ALLOW` (allow), `DENY` (deny).
        /// </summary>
        [Input("vpcRestrictionEffect")]
        public string? VpcRestrictionEffect { get; set; }

        public GetBackupDownloadUrlsBackupDownloadRestrictionArgs()
        {
        }
        public static new GetBackupDownloadUrlsBackupDownloadRestrictionArgs Empty => new GetBackupDownloadUrlsBackupDownloadRestrictionArgs();
    }
}
