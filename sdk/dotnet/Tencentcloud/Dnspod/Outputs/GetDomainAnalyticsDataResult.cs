// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod.Outputs
{

    [OutputType]
    public sealed class GetDomainAnalyticsDataResult
    {
        /// <summary>
        /// For daily statistics, it is the statistical date.
        /// </summary>
        public readonly string DateKey;
        /// <summary>
        /// For hourly statistics, it is the hour of the current time (0-23), for example, when HourKey is 23, the statistical period is the resolution volume from 22:00 to 23:00. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int HourKey;
        /// <summary>
        /// Subtotal of resolution volume for the current statistical dimension.
        /// </summary>
        public readonly int Num;

        [OutputConstructor]
        private GetDomainAnalyticsDataResult(
            string dateKey,

            int hourKey,

            int num)
        {
            DateKey = dateKey;
            HourKey = hourKey;
            Num = num;
        }
    }
}