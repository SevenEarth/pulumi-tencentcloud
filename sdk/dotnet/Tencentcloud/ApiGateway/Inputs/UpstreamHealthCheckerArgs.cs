// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Inputs
{

    public sealed class UpstreamHealthCheckerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Detect the requested path during active health checks. The default is&amp;#39;/&amp;#39;.
        /// </summary>
        [Input("activeCheckHttpPath")]
        public Input<string>? ActiveCheckHttpPath { get; set; }

        /// <summary>
        /// The time interval for active health checks is 5 seconds by default.
        /// </summary>
        [Input("activeCheckInterval")]
        public Input<int>? ActiveCheckInterval { get; set; }

        /// <summary>
        /// The detection request for active health check timed out in seconds. The default is 5 seconds.
        /// </summary>
        [Input("activeCheckTimeout")]
        public Input<int>? ActiveCheckTimeout { get; set; }

        /// <summary>
        /// Identify whether active health checks are enabled.
        /// </summary>
        [Input("enableActiveCheck", required: true)]
        public Input<bool> EnableActiveCheck { get; set; } = null!;

        /// <summary>
        /// Identify whether passive health checks are enabled.
        /// </summary>
        [Input("enablePassiveCheck", required: true)]
        public Input<bool> EnablePassiveCheck { get; set; } = null!;

        /// <summary>
        /// The HTTP status code that determines a successful request during a health check.
        /// </summary>
        [Input("healthyHttpStatus", required: true)]
        public Input<string> HealthyHttpStatus { get; set; } = null!;

        /// <summary>
        /// HTTP continuous error threshold. 0 means HTTP checking is disabled. Value range: [0, 254].
        /// </summary>
        [Input("httpFailureThreshold", required: true)]
        public Input<int> HttpFailureThreshold { get; set; } = null!;

        /// <summary>
        /// TCP continuous error threshold. 0 indicates disabling TCP checking. Value range: [0, 254].
        /// </summary>
        [Input("tcpFailureThreshold", required: true)]
        public Input<int> TcpFailureThreshold { get; set; } = null!;

        /// <summary>
        /// Continuous timeout threshold. 0 indicates disabling timeout checking. Value range: [0, 254].
        /// </summary>
        [Input("timeoutThreshold", required: true)]
        public Input<int> TimeoutThreshold { get; set; } = null!;

        /// <summary>
        /// The HTTP status code that determines a failed request during a health check.
        /// </summary>
        [Input("unhealthyHttpStatus", required: true)]
        public Input<string> UnhealthyHttpStatus { get; set; } = null!;

        /// <summary>
        /// The automatic recovery time of abnormal node status, in seconds. When only passive checking is enabled, it must be set to a value&amp;gt;0, otherwise the passive exception node will not be able to recover. The default is 30 seconds.
        /// 
        /// The `k8s_service` object supports the following:
        /// </summary>
        [Input("unhealthyTimeout")]
        public Input<int>? UnhealthyTimeout { get; set; }

        public UpstreamHealthCheckerArgs()
        {
        }
        public static new UpstreamHealthCheckerArgs Empty => new UpstreamHealthCheckerArgs();
    }
}
