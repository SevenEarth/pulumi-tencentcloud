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
    public sealed class ApiTargetServicesHealthCheckConf
    {
        /// <summary>
        /// Threshold percentage.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int? ErrorThresholdPercentage;
        /// <summary>
        /// Whether to initiate a health check.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly bool? IsHealthCheck;
        /// <summary>
        /// Health check threshold.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int? RequestVolumeThreshold;
        /// <summary>
        /// Window size.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int? SleepWindowInMilliseconds;

        [OutputConstructor]
        private ApiTargetServicesHealthCheckConf(
            int? errorThresholdPercentage,

            bool? isHealthCheck,

            int? requestVolumeThreshold,

            int? sleepWindowInMilliseconds)
        {
            ErrorThresholdPercentage = errorThresholdPercentage;
            IsHealthCheck = isHealthCheck;
            RequestVolumeThreshold = requestVolumeThreshold;
            SleepWindowInMilliseconds = sleepWindowInMilliseconds;
        }
    }
}
