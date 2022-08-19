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
    public sealed class GetThrottlingApisListResult
    {
        /// <summary>
        /// List of throttling policies bound to API.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetThrottlingApisListApiEnvironmentStrategyResult> ApiEnvironmentStrategies;
        /// <summary>
        /// Unique service ID of API.
        /// </summary>
        public readonly string ServiceId;

        [OutputConstructor]
        private GetThrottlingApisListResult(
            ImmutableArray<Outputs.GetThrottlingApisListApiEnvironmentStrategyResult> apiEnvironmentStrategies,

            string serviceId)
        {
            ApiEnvironmentStrategies = apiEnvironmentStrategies;
            ServiceId = serviceId;
        }
    }
}
