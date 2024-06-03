// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateDataDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Dedicated Cluster(CDC) ID.
        /// </summary>
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        /// <summary>
        /// Whether the data disk is destroyed along with the instance, true or false.
        /// </summary>
        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        /// <summary>
        /// Data disk ID.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// The size of the data disk.
        /// </summary>
        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        /// <summary>
        /// The type of data disk.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Whether the data disk is encrypted, TRUE or FALSE.
        /// </summary>
        [Input("encrypt")]
        public Input<bool>? Encrypt { get; set; }

        /// <summary>
        /// The id of custom CMK.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Data disk snapshot ID.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// Cloud disk performance, MB/s.
        /// </summary>
        [Input("throughputPerformance")]
        public Input<int>? ThroughputPerformance { get; set; }

        public LaunchTemplateDataDiskArgs()
        {
        }
        public static new LaunchTemplateDataDiskArgs Empty => new LaunchTemplateDataDiskArgs();
    }
}
