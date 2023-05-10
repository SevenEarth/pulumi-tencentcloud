// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateVersionInstanceChargePrepaid
    {
        /// <summary>
        /// Subscription period; unit: month; valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
        /// </summary>
        public readonly int Period;
        /// <summary>
        /// Auto renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically DISABLE_NOTIFY_AND_MANUAL_RENEW: neither notify upon expiration nor renew automatically &amp;lt;br&amp;gt;&amp;lt;br&amp;gt;Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient.
        /// </summary>
        public readonly string? RenewFlag;

        [OutputConstructor]
        private LaunchTemplateVersionInstanceChargePrepaid(
            int period,

            string? renewFlag)
        {
            Period = period;
            RenewFlag = renewFlag;
        }
    }
}
