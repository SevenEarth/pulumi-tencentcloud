// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Gaap.Outputs
{

    [OutputType]
    public sealed class GetProxyDetailProxyDetailAccessRegionInfoSupportFeatureResult
    {
        /// <summary>
        /// A list of network types supported by the access area, with normal indicating support for regular BGP, cn2 indicating premium BGP, triple indicating three networks, and secure_EIP represents a custom secure EIP.
        /// </summary>
        public readonly ImmutableArray<string> NetworkTypes;

        [OutputConstructor]
        private GetProxyDetailProxyDetailAccessRegionInfoSupportFeatureResult(ImmutableArray<string> networkTypes)
        {
            NetworkTypes = networkTypes;
        }
    }
}
