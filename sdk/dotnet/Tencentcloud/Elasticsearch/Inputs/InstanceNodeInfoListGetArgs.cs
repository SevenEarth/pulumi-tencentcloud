// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Inputs
{

    public sealed class InstanceNodeInfoListGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Node disk size. Unit is GB, and default value is `100`.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Node disk type. Valid values are `CLOUD_SSD` and `CLOUD_PREMIUM`, `CLOUD_HSSD`. The default value is `CLOUD_SSD`.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Decides to encrypt this disk or not.
        /// </summary>
        [Input("encrypt")]
        public Input<bool>? Encrypt { get; set; }

        /// <summary>
        /// Number of nodes.
        /// </summary>
        [Input("nodeNum", required: true)]
        public Input<int> NodeNum { get; set; } = null!;

        /// <summary>
        /// Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// Node type. Valid values are `hotData`, `warmData` and `dedicatedMaster`. The default value is 'hotData`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public InstanceNodeInfoListGetArgs()
        {
        }
    }
}
