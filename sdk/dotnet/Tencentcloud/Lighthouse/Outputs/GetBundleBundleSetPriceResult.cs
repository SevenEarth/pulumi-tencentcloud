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
    public sealed class GetBundleBundleSetPriceResult
    {
        /// <summary>
        /// Instance price.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBundleBundleSetPriceInstancePriceResult> InstancePrices;

        [OutputConstructor]
        private GetBundleBundleSetPriceResult(ImmutableArray<Outputs.GetBundleBundleSetPriceInstancePriceResult> instancePrices)
        {
            InstancePrices = instancePrices;
        }
    }
}