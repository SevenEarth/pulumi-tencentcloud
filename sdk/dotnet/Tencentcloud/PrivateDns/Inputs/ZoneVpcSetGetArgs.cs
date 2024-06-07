// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.PrivateDns.Inputs
{

    public sealed class ZoneVpcSetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPC REGION.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("uniqVpcId", required: true)]
        public Input<string> UniqVpcId { get; set; } = null!;

        public ZoneVpcSetGetArgs()
        {
        }
        public static new ZoneVpcSetGetArgs Empty => new ZoneVpcSetGetArgs();
    }
}
