// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetBandwidthPackageQuotaQuotaSetResult
    {
        /// <summary>
        /// current amount.
        /// </summary>
        public readonly int QuotaCurrent;
        /// <summary>
        /// Quota type.
        /// </summary>
        public readonly string QuotaId;
        /// <summary>
        /// quota amount.
        /// </summary>
        public readonly int QuotaLimit;

        [OutputConstructor]
        private GetBandwidthPackageQuotaQuotaSetResult(
            int quotaCurrent,

            string quotaId,

            int quotaLimit)
        {
            QuotaCurrent = quotaCurrent;
            QuotaId = quotaId;
            QuotaLimit = quotaLimit;
        }
    }
}
