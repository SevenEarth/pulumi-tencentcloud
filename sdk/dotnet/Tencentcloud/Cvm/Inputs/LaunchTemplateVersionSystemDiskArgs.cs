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

    public sealed class LaunchTemplateVersionSystemDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the dedicated cluster to which the instance belongs.
        /// </summary>
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        /// <summary>
        /// System disk ID. System disks whose type is LOCAL_BASIC or LOCAL_SSD do not have an ID and do not support this parameter. It is only used as a response parameter for APIs such as DescribeInstances, and cannot be used as a request parameter for APIs such as RunInstances.
        /// </summary>
        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        /// <summary>
        /// System disk size; unit: GB; default value: 50 GB.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// The type of system disk. Default value: the type of hard disk currently in stock.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        public LaunchTemplateVersionSystemDiskArgs()
        {
        }
        public static new LaunchTemplateVersionSystemDiskArgs Empty => new LaunchTemplateVersionSystemDiskArgs();
    }
}
