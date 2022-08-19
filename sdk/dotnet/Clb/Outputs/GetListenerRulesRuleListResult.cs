// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class GetListenerRulesRuleListResult
    {
        /// <summary>
        /// ID of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.
        /// </summary>
        public readonly string CertificateCaId;
        /// <summary>
        /// ID of the server certificate. NOTES: Only supports listeners of 'HTTPS'  and 'TCP_SSL' protocol.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// Type of SSL Mode, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'.NOTES: Only supports listeners of 'HTTPS'  and 'TCP_SSL' protocol.
        /// </summary>
        public readonly string CertificateSslMode;
        /// <summary>
        /// ID of the CLB to be queried.
        /// </summary>
        public readonly string ClbId;
        /// <summary>
        /// Domain name of the forwarding rule to be queried.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// Health threshold of health check, and the default is `3`. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.
        /// </summary>
        public readonly int HealthCheckHealthNum;
        /// <summary>
        /// HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value 4xx is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
        /// </summary>
        public readonly int HealthCheckHttpCode;
        /// <summary>
        /// Domain name of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.
        /// </summary>
        public readonly string HealthCheckHttpDomain;
        /// <summary>
        /// Methods of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.
        /// </summary>
        public readonly string HealthCheckHttpMethod;
        /// <summary>
        /// Path of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.
        /// </summary>
        public readonly string HealthCheckHttpPath;
        /// <summary>
        /// Interval time of health check. The value range is 5-300 sec, and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.
        /// </summary>
        public readonly int HealthCheckIntervalTime;
        /// <summary>
        /// Indicates whether health check is enabled.
        /// </summary>
        public readonly bool HealthCheckSwitch;
        /// <summary>
        /// Unhealth threshold of health check, and the default is `3`. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.
        /// </summary>
        public readonly int HealthCheckUnhealthNum;
        /// <summary>
        /// Indicate to set HTTP2 protocol or not.
        /// </summary>
        public readonly bool Http2Switch;
        /// <summary>
        /// ID of the CLB listener to be queried.
        /// </summary>
        public readonly string ListenerId;
        /// <summary>
        /// ID of the forwarding rule to be queried.
        /// </summary>
        public readonly string RuleId;
        /// <summary>
        /// Scheduling method of the forwarding rule of thr CLB listener, and available values include `WRR`, `IP HASH` and `LEAST_CONN`. The default is `WRR`.
        /// </summary>
        public readonly string Scheduler;
        /// <summary>
        /// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.
        /// </summary>
        public readonly int SessionExpireTime;
        /// <summary>
        /// Url of the forwarding rule to be queried.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private GetListenerRulesRuleListResult(
            string certificateCaId,

            string certificateId,

            string certificateSslMode,

            string clbId,

            string? domain,

            int healthCheckHealthNum,

            int healthCheckHttpCode,

            string healthCheckHttpDomain,

            string healthCheckHttpMethod,

            string healthCheckHttpPath,

            int healthCheckIntervalTime,

            bool healthCheckSwitch,

            int healthCheckUnhealthNum,

            bool http2Switch,

            string listenerId,

            string ruleId,

            string scheduler,

            int sessionExpireTime,

            string? url)
        {
            CertificateCaId = certificateCaId;
            CertificateId = certificateId;
            CertificateSslMode = certificateSslMode;
            ClbId = clbId;
            Domain = domain;
            HealthCheckHealthNum = healthCheckHealthNum;
            HealthCheckHttpCode = healthCheckHttpCode;
            HealthCheckHttpDomain = healthCheckHttpDomain;
            HealthCheckHttpMethod = healthCheckHttpMethod;
            HealthCheckHttpPath = healthCheckHttpPath;
            HealthCheckIntervalTime = healthCheckIntervalTime;
            HealthCheckSwitch = healthCheckSwitch;
            HealthCheckUnhealthNum = healthCheckUnhealthNum;
            Http2Switch = http2Switch;
            ListenerId = listenerId;
            RuleId = ruleId;
            Scheduler = scheduler;
            SessionExpireTime = sessionExpireTime;
            Url = url;
        }
    }
}
