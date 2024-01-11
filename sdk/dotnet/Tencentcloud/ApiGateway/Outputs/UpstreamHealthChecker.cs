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
    public sealed class UpstreamHealthChecker
    {
        /// <summary>
        /// Detect the requested path during active health checks. The default is&amp;#39;/&amp;#39;.
        /// </summary>
        public readonly string? ActiveCheckHttpPath;
        /// <summary>
        /// The time interval for active health checks is 5 seconds by default.
        /// </summary>
        public readonly int? ActiveCheckInterval;
        /// <summary>
        /// The detection request for active health check timed out in seconds. The default is 5 seconds.
        /// </summary>
        public readonly int? ActiveCheckTimeout;
        /// <summary>
        /// Identify whether active health checks are enabled.
        /// </summary>
        public readonly bool EnableActiveCheck;
        /// <summary>
        /// Identify whether passive health checks are enabled.
        /// </summary>
        public readonly bool EnablePassiveCheck;
        /// <summary>
        /// The HTTP status code that determines a successful request during a health check.
        /// </summary>
        public readonly string HealthyHttpStatus;
        /// <summary>
        /// HTTP continuous error threshold. 0 means HTTP checking is disabled. Value range: [0, 254].
        /// </summary>
        public readonly int HttpFailureThreshold;
        /// <summary>
        /// TCP continuous error threshold. 0 indicates disabling TCP checking. Value range: [0, 254].
        /// </summary>
        public readonly int TcpFailureThreshold;
        /// <summary>
        /// Continuous timeout threshold. 0 indicates disabling timeout checking. Value range: [0, 254].
        /// </summary>
        public readonly int TimeoutThreshold;
        /// <summary>
        /// The HTTP status code that determines a failed request during a health check.
        /// </summary>
        public readonly string UnhealthyHttpStatus;
        /// <summary>
        /// The automatic recovery time of abnormal node status, in seconds. When only passive checking is enabled, it must be set to a value&amp;gt;0, otherwise the passive exception node will not be able to recover. The default is 30 seconds.
        /// </summary>
        public readonly int? UnhealthyTimeout;

        [OutputConstructor]
        private UpstreamHealthChecker(
            string? activeCheckHttpPath,

            int? activeCheckInterval,

            int? activeCheckTimeout,

            bool enableActiveCheck,

            bool enablePassiveCheck,

            string healthyHttpStatus,

            int httpFailureThreshold,

            int tcpFailureThreshold,

            int timeoutThreshold,

            string unhealthyHttpStatus,

            int? unhealthyTimeout)
        {
            ActiveCheckHttpPath = activeCheckHttpPath;
            ActiveCheckInterval = activeCheckInterval;
            ActiveCheckTimeout = activeCheckTimeout;
            EnableActiveCheck = enableActiveCheck;
            EnablePassiveCheck = enablePassiveCheck;
            HealthyHttpStatus = healthyHttpStatus;
            HttpFailureThreshold = httpFailureThreshold;
            TcpFailureThreshold = tcpFailureThreshold;
            TimeoutThreshold = timeoutThreshold;
            UnhealthyHttpStatus = unhealthyHttpStatus;
            UnhealthyTimeout = unhealthyTimeout;
        }
    }
}