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
    public sealed class GetModifyInstanceBundleModifyBundleSetResult
    {
        /// <summary>
        /// Package information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetModifyInstanceBundleModifyBundleSetBundleResult> Bundles;
        /// <summary>
        /// Change the status of the package. Value:
        /// - SOLD_OUT: the package is sold out;
        /// - AVAILABLE: support package changes;
        /// - UNAVAILABLE: package changes are not supported for the time being.
        /// </summary>
        public readonly string ModifyBundleState;
        /// <summary>
        /// Change the price difference to be made up after the instance package.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetModifyInstanceBundleModifyBundleSetModifyPriceResult> ModifyPrices;
        /// <summary>
        /// Package change reason information is not supported. When the package status is changed to `AVAILABLE`, the information is empty.
        /// </summary>
        public readonly string NotSupportModifyMessage;

        [OutputConstructor]
        private GetModifyInstanceBundleModifyBundleSetResult(
            ImmutableArray<Outputs.GetModifyInstanceBundleModifyBundleSetBundleResult> bundles,

            string modifyBundleState,

            ImmutableArray<Outputs.GetModifyInstanceBundleModifyBundleSetModifyPriceResult> modifyPrices,

            string notSupportModifyMessage)
        {
            Bundles = bundles;
            ModifyBundleState = modifyBundleState;
            ModifyPrices = modifyPrices;
            NotSupportModifyMessage = notSupportModifyMessage;
        }
    }
}