// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class CngwServiceRateLimitLimitDetailQpsThreshold
    {
        /// <summary>
        /// the max threshold.
        /// </summary>
        public readonly int Max;
        /// <summary>
        /// qps threshold unit.Reference value:`second`, `minute`, `hour`, `day`, `month`, `year`.
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private CngwServiceRateLimitLimitDetailQpsThreshold(
            int max,

            string unit)
        {
            Max = max;
            Unit = unit;
        }
    }
}
