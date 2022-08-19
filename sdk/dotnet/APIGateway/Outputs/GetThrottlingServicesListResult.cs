// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetThrottlingServicesListResult
    {
        /// <summary>
        /// A list of Throttling policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetThrottlingServicesListEnvironmentResult> Environments;
        /// <summary>
        /// Service ID for query.
        /// </summary>
        public readonly string ServiceId;

        [OutputConstructor]
        private GetThrottlingServicesListResult(
            ImmutableArray<Outputs.GetThrottlingServicesListEnvironmentResult> environments,

            string serviceId)
        {
            Environments = environments;
            ServiceId = serviceId;
        }
    }
}
