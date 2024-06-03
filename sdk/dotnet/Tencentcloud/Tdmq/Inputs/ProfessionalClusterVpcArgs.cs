// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq.Inputs
{

    public sealed class ProfessionalClusterVpcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of Subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Id of VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ProfessionalClusterVpcArgs()
        {
        }
        public static new ProfessionalClusterVpcArgs Empty => new ProfessionalClusterVpcArgs();
    }
}
