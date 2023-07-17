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

    public sealed class ModifyInstanceDiskTypeDataDiskGetArgs : Pulumi.ResourceArgs
    {
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        [Input("deleteWithInstance")]
        public Input<bool>? DeleteWithInstance { get; set; }

        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        [Input("diskSize", required: true)]
        public Input<int> DiskSize { get; set; } = null!;

        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("encrypt")]
        public Input<bool>? Encrypt { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("throughputPerformance")]
        public Input<int>? ThroughputPerformance { get; set; }

        public ModifyInstanceDiskTypeDataDiskGetArgs()
        {
        }
    }
}
