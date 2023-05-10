// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ccn.Inputs
{

    public sealed class InstancesAcceptAttachInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Attachment Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Instance Region.
        /// </summary>
        [Input("instanceRegion", required: true)]
        public Input<string> InstanceRegion { get; set; } = null!;

        /// <summary>
        /// InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        public InstancesAcceptAttachInstanceArgs()
        {
        }
    }
}
