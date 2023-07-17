// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class GetInstanceByCertIdCertSetResult
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        public readonly string CertId;
        /// <summary>
        /// List of CLB instances associated with certificate. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceByCertIdCertSetLoadBalancerResult> LoadBalancers;

        [OutputConstructor]
        private GetInstanceByCertIdCertSetResult(
            string certId,

            ImmutableArray<Outputs.GetInstanceByCertIdCertSetLoadBalancerResult> loadBalancers)
        {
            CertId = certId;
            LoadBalancers = loadBalancers;
        }
    }
}
