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

    public sealed class JobLoadGeoRegionsLoadDistributionGetArgs : Pulumi.ResourceArgs
    {
        [Input("percentage")]
        public Input<int>? Percentage { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("regionId", required: true)]
        public Input<int> RegionId { get; set; } = null!;

        public JobLoadGeoRegionsLoadDistributionGetArgs()
        {
        }
    }
}
