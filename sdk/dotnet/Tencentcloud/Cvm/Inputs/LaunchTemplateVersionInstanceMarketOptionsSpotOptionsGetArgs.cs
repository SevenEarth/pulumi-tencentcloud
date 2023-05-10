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

    public sealed class LaunchTemplateVersionInstanceMarketOptionsSpotOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bidding price.
        /// </summary>
        [Input("maxPrice", required: true)]
        public Input<string> MaxPrice { get; set; } = null!;

        /// <summary>
        /// Bidding request type. Currently only one-time is supported.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        public LaunchTemplateVersionInstanceMarketOptionsSpotOptionsGetArgs()
        {
        }
    }
}
