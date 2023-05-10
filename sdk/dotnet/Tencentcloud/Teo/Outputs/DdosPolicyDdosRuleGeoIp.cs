// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class DdosPolicyDdosRuleGeoIp
    {
        /// <summary>
        /// Region ID. See details in data source `security_policy_regions`.
        /// </summary>
        public readonly ImmutableArray<int> RegionIds;
        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private DdosPolicyDdosRuleGeoIp(
            ImmutableArray<int> regionIds,

            string? @switch)
        {
            RegionIds = regionIds;
            Switch = @switch;
        }
    }
}