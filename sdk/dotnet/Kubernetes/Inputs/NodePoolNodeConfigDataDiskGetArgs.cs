// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Inputs
{

    public sealed class NodePoolNodeConfigDataDiskGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicate whether to auto format and mount or not. Default is `false`.
        /// </summary>
        [Input("autoFormatAndMount")]
        public Input<bool>? AutoFormatAndMount { get; set; }

        /// <summary>
        /// The name of the device or partition to mount. NOTE: this argument doesn't support setting in node pool, or will leads to mount error.
        /// </summary>
        [Input("diskPartition")]
        public Input<string>? DiskPartition { get; set; }

        /// <summary>
        /// Volume of disk in GB. Default is `0`.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Types of disk. Valid value: `CLOUD_PREMIUM` and `CLOUD_SSD`.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// File system, e.g. `ext3/ext4/xfs`.
        /// </summary>
        [Input("fileSystem")]
        public Input<string>? FileSystem { get; set; }

        /// <summary>
        /// Mount target.
        /// </summary>
        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        public NodePoolNodeConfigDataDiskGetArgs()
        {
        }
    }
}
