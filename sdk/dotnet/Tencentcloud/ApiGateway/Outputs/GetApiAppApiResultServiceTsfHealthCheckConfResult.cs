// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetApiAppApiResultServiceTsfHealthCheckConfResult
    {
        /// <summary>
        /// Threshold percentage.
        /// </summary>
        public readonly int ErrorThresholdPercentage;
        /// <summary>
        /// Whether to enable health check.
        /// </summary>
        public readonly bool IsHealthCheck;
        /// <summary>
        /// Health check threshold.
        /// </summary>
        public readonly int RequestVolumeThreshold;
        /// <summary>
        /// Window size.
        /// </summary>
        public readonly int SleepWindowInMilliseconds;

        [OutputConstructor]
        private GetApiAppApiResultServiceTsfHealthCheckConfResult(
            int errorThresholdPercentage,

            bool isHealthCheck,

            int requestVolumeThreshold,

            int sleepWindowInMilliseconds)
        {
            ErrorThresholdPercentage = errorThresholdPercentage;
            IsHealthCheck = isHealthCheck;
            RequestVolumeThreshold = requestVolumeThreshold;
            SleepWindowInMilliseconds = sleepWindowInMilliseconds;
        }
    }
}
