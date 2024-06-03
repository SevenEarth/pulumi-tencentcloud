// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class JobLoadGeoRegionsLoadDistribution
    {
        /// <summary>
        /// Percentage.
        /// </summary>
        public readonly int? Percentage;
        /// <summary>
        /// Region.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// Regional ID.
        /// </summary>
        public readonly int RegionId;

        [OutputConstructor]
        private JobLoadGeoRegionsLoadDistribution(
            int? percentage,

            string? region,

            int regionId)
        {
            Percentage = percentage;
            Region = region;
            RegionId = regionId;
        }
    }
}
