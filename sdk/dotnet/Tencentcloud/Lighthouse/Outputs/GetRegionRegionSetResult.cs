// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Lighthouse.Outputs
{

    [OutputType]
    public sealed class GetRegionRegionSetResult
    {
        /// <summary>
        /// Whether the region is in the Chinese mainland.
        /// </summary>
        public readonly bool IsChinaMainland;
        /// <summary>
        /// Region name.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Region description.
        /// </summary>
        public readonly string RegionName;
        /// <summary>
        /// Region availability status.
        /// </summary>
        public readonly string RegionState;

        [OutputConstructor]
        private GetRegionRegionSetResult(
            bool isChinaMainland,

            string region,

            string regionName,

            string regionState)
        {
            IsChinaMainland = isChinaMainland;
            Region = region;
            RegionName = regionName;
            RegionState = regionState;
        }
    }
}
