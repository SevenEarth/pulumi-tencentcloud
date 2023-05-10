// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ccn.Outputs
{

    [OutputType]
    public sealed class GetCrossBorderRegionBandwidthLimitsCcnBandwidthSetResult
    {
        /// <summary>
        /// ccn id.
        /// </summary>
        public readonly string CcnId;
        /// <summary>
        /// bandwidth limit of cross region.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCrossBorderRegionBandwidthLimitsCcnBandwidthSetCcnRegionBandwidthLimitResult> CcnRegionBandwidthLimits;
        /// <summary>
        /// create time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// expired time.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// `POSTPAID` or `PREPAID`.
        /// </summary>
        public readonly string InstanceChargeType;
        /// <summary>
        /// if cross region.
        /// </summary>
        public readonly bool IsCrossBorder;
        /// <summary>
        /// `true` means locked.
        /// </summary>
        public readonly bool IsSecurityLock;
        /// <summary>
        /// market id.
        /// </summary>
        public readonly string MarketId;
        /// <summary>
        /// Id of RegionFlowControl.
        /// </summary>
        public readonly string RegionFlowControlId;
        /// <summary>
        /// renew flag.
        /// </summary>
        public readonly string RenewFlag;
        /// <summary>
        /// update time.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// user account id.
        /// </summary>
        public readonly string UserAccountId;

        [OutputConstructor]
        private GetCrossBorderRegionBandwidthLimitsCcnBandwidthSetResult(
            string ccnId,

            ImmutableArray<Outputs.GetCrossBorderRegionBandwidthLimitsCcnBandwidthSetCcnRegionBandwidthLimitResult> ccnRegionBandwidthLimits,

            string createdTime,

            string expiredTime,

            string instanceChargeType,

            bool isCrossBorder,

            bool isSecurityLock,

            string marketId,

            string regionFlowControlId,

            string renewFlag,

            string updateTime,

            string userAccountId)
        {
            CcnId = ccnId;
            CcnRegionBandwidthLimits = ccnRegionBandwidthLimits;
            CreatedTime = createdTime;
            ExpiredTime = expiredTime;
            InstanceChargeType = instanceChargeType;
            IsCrossBorder = isCrossBorder;
            IsSecurityLock = isSecurityLock;
            MarketId = marketId;
            RegionFlowControlId = regionFlowControlId;
            RenewFlag = renewFlag;
            UpdateTime = updateTime;
            UserAccountId = userAccountId;
        }
    }
}
