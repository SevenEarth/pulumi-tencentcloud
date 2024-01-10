// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc.Outputs
{

    [OutputType]
    public sealed class GetDescribeUserInfoUserInfoDataPolicyInfoResult
    {
        /// <summary>
        /// Policy set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeUserInfoUserInfoDataPolicyInfoPolicySetResult> PolicySets;
        /// <summary>
        /// Total count.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetDescribeUserInfoUserInfoDataPolicyInfoResult(
            ImmutableArray<Outputs.GetDescribeUserInfoUserInfoDataPolicyInfoPolicySetResult> policySets,

            int totalCount)
        {
            PolicySets = policySets;
            TotalCount = totalCount;
        }
    }
}
