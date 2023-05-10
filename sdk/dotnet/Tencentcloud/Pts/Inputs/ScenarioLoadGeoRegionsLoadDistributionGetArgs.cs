// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class ScenarioLoadGeoRegionsLoadDistributionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Percentage.
        /// </summary>
        [Input("percentage")]
        public Input<int>? Percentage { get; set; }

        /// <summary>
        /// Region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Regional ID.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<int> RegionId { get; set; } = null!;

        public ScenarioLoadGeoRegionsLoadDistributionGetArgs()
        {
        }
    }
}
