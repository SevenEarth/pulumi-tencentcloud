// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class DeployVmGroupHealthCheckSettingsReadinessProbe
    {
        /// <summary>
        /// The health check method. HTTP indicates checking through an HTTP interface, CMD indicates checking through executing a command, and TCP indicates checking through establishing a TCP connection. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ActionType;
        /// <summary>
        /// The command to be executed for command check. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// The number of consecutive successful health checks required for the backend container to transition from success to failure. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? FailureThreshold;
        /// <summary>
        /// The time to delay the start of the container health check. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? InitialDelaySeconds;
        /// <summary>
        /// The request path for HTTP health checks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The time interval for performing health checks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? PeriodSeconds;
        /// <summary>
        /// The port used for health checks, ranging from 1 to 65535. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The protocol used for HTTP health checks. HTTP and HTTPS are supported. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? Scheme;
        /// <summary>
        /// The number of consecutive successful health checks required for the backend container to transition from failure to success. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? SuccessThreshold;
        /// <summary>
        /// The maximum timeout period for each health check response. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int? TimeoutSeconds;
        /// <summary>
        /// The type of readiness probe. TSF_DEFAULT represents the default readiness probe of TSF, while K8S_NATIVE represents the native readiness probe of Kubernetes. If this field is not specified, the native readiness probe of Kubernetes is used by default. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DeployVmGroupHealthCheckSettingsReadinessProbe(
            string actionType,

            ImmutableArray<string> commands,

            int? failureThreshold,

            int? initialDelaySeconds,

            string? path,

            int? periodSeconds,

            int? port,

            string? scheme,

            int? successThreshold,

            int? timeoutSeconds,

            string? type)
        {
            ActionType = actionType;
            Commands = commands;
            FailureThreshold = failureThreshold;
            InitialDelaySeconds = initialDelaySeconds;
            Path = path;
            PeriodSeconds = periodSeconds;
            Port = port;
            Scheme = scheme;
            SuccessThreshold = successThreshold;
            TimeoutSeconds = timeoutSeconds;
            Type = type;
        }
    }
}
