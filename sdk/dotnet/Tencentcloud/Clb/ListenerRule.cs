// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb
{
    /// <summary>
    /// Provides a resource to create a CLB listener rule.
    /// 
    /// &gt; **NOTE:** This resource only be applied to the HTTP or HTTPS listeners.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Clb.ListenerRule("foo", new Tencentcloud.Clb.ListenerRuleArgs
    ///         {
    ///             CertificateCaId = "VfqO4zkB",
    ///             CertificateId = "VjANRdz8",
    ///             CertificateSslMode = "MUTUAL",
    ///             ClbId = "lb-k2zjp9lv",
    ///             Domain = "foo.net",
    ///             HealthCheckHealthNum = 3,
    ///             HealthCheckHttpCode = 2,
    ///             HealthCheckHttpDomain = "Default Domain",
    ///             HealthCheckHttpMethod = "GET",
    ///             HealthCheckHttpPath = "Default Path",
    ///             HealthCheckIntervalTime = 5,
    ///             HealthCheckSwitch = true,
    ///             HealthCheckUnhealthNum = 3,
    ///             ListenerId = "lbl-hh141sn9",
    ///             Scheduler = "WRR",
    ///             SessionExpireTime = 30,
    ///             Url = "/bar",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CLB listener rule can be imported using the id (version &gt;= 1.47.0), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Clb/listenerRule:ListenerRule foo lb-7a0t6zqb#lbl-hh141sn9#loc-agg236ys
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clb/listenerRule:ListenerRule")]
    public partial class ListenerRule : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Output("certificateCaId")]
        public Output<string?> CertificateCaId { get; private set; } = null!;

        /// <summary>
        /// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Output("certificateId")]
        public Output<string?> CertificateId { get; private set; } = null!;

        /// <summary>
        /// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Output("certificateSslMode")]
        public Output<string?> CertificateSslMode { get; private set; } = null!;

        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Output("clbId")]
        public Output<string> ClbId { get; private set; } = null!;

        /// <summary>
        /// Domain name of the listener rule.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`. The default is `HTTP`.
        /// </summary>
        [Output("forwardType")]
        public Output<string> ForwardType { get; private set; } = null!;

        /// <summary>
        /// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Output("healthCheckHealthNum")]
        public Output<int> HealthCheckHealthNum { get; private set; } = null!;

        /// <summary>
        /// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
        /// </summary>
        [Output("healthCheckHttpCode")]
        public Output<int> HealthCheckHttpCode { get; private set; } = null!;

        /// <summary>
        /// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Output("healthCheckHttpDomain")]
        public Output<string> HealthCheckHttpDomain { get; private set; } = null!;

        /// <summary>
        /// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
        /// </summary>
        [Output("healthCheckHttpMethod")]
        public Output<string> HealthCheckHttpMethod { get; private set; } = null!;

        /// <summary>
        /// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Output("healthCheckHttpPath")]
        public Output<string> HealthCheckHttpPath { get; private set; } = null!;

        /// <summary>
        /// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Output("healthCheckIntervalTime")]
        public Output<int> HealthCheckIntervalTime { get; private set; } = null!;

        /// <summary>
        /// Indicates whether health check is enabled.
        /// </summary>
        [Output("healthCheckSwitch")]
        public Output<bool> HealthCheckSwitch { get; private set; } = null!;

        /// <summary>
        /// Time out of health check. The value range is 2-60.
        /// </summary>
        [Output("healthCheckTimeOut")]
        public Output<int> HealthCheckTimeOut { get; private set; } = null!;

        /// <summary>
        /// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
        /// </summary>
        [Output("healthCheckType")]
        public Output<string> HealthCheckType { get; private set; } = null!;

        /// <summary>
        /// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Output("healthCheckUnhealthNum")]
        public Output<int> HealthCheckUnhealthNum { get; private set; } = null!;

        /// <summary>
        /// Indicate to apply HTTP2.0 protocol or not.
        /// </summary>
        [Output("http2Switch")]
        public Output<bool> Http2Switch { get; private set; } = null!;

        /// <summary>
        /// ID of CLB listener.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names.
        /// </summary>
        [Output("quic")]
        public Output<bool> Quic { get; private set; } = null!;

        /// <summary>
        /// ID of this CLB listener rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Output("scheduler")]
        public Output<string?> Scheduler { get; private set; } = null!;

        /// <summary>
        /// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Output("sessionExpireTime")]
        public Output<int?> SessionExpireTime { get; private set; } = null!;

        /// <summary>
        /// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
        /// </summary>
        [Output("targetType")]
        public Output<string?> TargetType { get; private set; } = null!;

        /// <summary>
        /// Url of the listener rule.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ListenerRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ListenerRule(string name, ListenerRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/listenerRule:ListenerRule", name, args ?? new ListenerRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ListenerRule(string name, Input<string> id, ListenerRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/listenerRule:ListenerRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ListenerRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ListenerRule Get(string name, Input<string> id, ListenerRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ListenerRule(name, id, state, options);
        }
    }

    public sealed class ListenerRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateCaId")]
        public Input<string>? CertificateCaId { get; set; }

        /// <summary>
        /// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateSslMode")]
        public Input<string>? CertificateSslMode { get; set; }

        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Input("clbId", required: true)]
        public Input<string> ClbId { get; set; } = null!;

        /// <summary>
        /// Domain name of the listener rule.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`. The default is `HTTP`.
        /// </summary>
        [Input("forwardType")]
        public Input<string>? ForwardType { get; set; }

        /// <summary>
        /// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckHealthNum")]
        public Input<int>? HealthCheckHealthNum { get; set; }

        /// <summary>
        /// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
        /// </summary>
        [Input("healthCheckHttpCode")]
        public Input<int>? HealthCheckHttpCode { get; set; }

        /// <summary>
        /// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Input("healthCheckHttpDomain")]
        public Input<string>? HealthCheckHttpDomain { get; set; }

        /// <summary>
        /// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
        /// </summary>
        [Input("healthCheckHttpMethod")]
        public Input<string>? HealthCheckHttpMethod { get; set; }

        /// <summary>
        /// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Input("healthCheckHttpPath")]
        public Input<string>? HealthCheckHttpPath { get; set; }

        /// <summary>
        /// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckIntervalTime")]
        public Input<int>? HealthCheckIntervalTime { get; set; }

        /// <summary>
        /// Indicates whether health check is enabled.
        /// </summary>
        [Input("healthCheckSwitch")]
        public Input<bool>? HealthCheckSwitch { get; set; }

        /// <summary>
        /// Time out of health check. The value range is 2-60.
        /// </summary>
        [Input("healthCheckTimeOut")]
        public Input<int>? HealthCheckTimeOut { get; set; }

        /// <summary>
        /// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
        /// </summary>
        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckUnhealthNum")]
        public Input<int>? HealthCheckUnhealthNum { get; set; }

        /// <summary>
        /// Indicate to apply HTTP2.0 protocol or not.
        /// </summary>
        [Input("http2Switch")]
        public Input<bool>? Http2Switch { get; set; }

        /// <summary>
        /// ID of CLB listener.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names.
        /// </summary>
        [Input("quic")]
        public Input<bool>? Quic { get; set; }

        /// <summary>
        /// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("sessionExpireTime")]
        public Input<int>? SessionExpireTime { get; set; }

        /// <summary>
        /// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// Url of the listener rule.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public ListenerRuleArgs()
        {
        }
    }

    public sealed class ListenerRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the client certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateCaId")]
        public Input<string>? CertificateCaId { get; set; }

        /// <summary>
        /// ID of the server certificate. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Type of certificate. Valid values: `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of HTTPS protocol.
        /// </summary>
        [Input("certificateSslMode")]
        public Input<string>? CertificateSslMode { get; set; }

        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Input("clbId")]
        public Input<string>? ClbId { get; set; }

        /// <summary>
        /// Domain name of the listener rule.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Forwarding protocol between the CLB instance and real server. Valid values: `HTTP`, `HTTPS`, `TRPC`. The default is `HTTP`.
        /// </summary>
        [Input("forwardType")]
        public Input<string>? ForwardType { get; set; }

        /// <summary>
        /// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is [2-10]. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckHealthNum")]
        public Input<int>? HealthCheckHealthNum { get; set; }

        /// <summary>
        /// HTTP Status Code. The default is 31. Valid value ranges: [1~31]. `1 means the return value '1xx' is health. `2` means the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.
        /// </summary>
        [Input("healthCheckHttpCode")]
        public Input<int>? HealthCheckHttpCode { get; set; }

        /// <summary>
        /// Domain name of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Input("healthCheckHttpDomain")]
        public Input<string>? HealthCheckHttpDomain { get; set; }

        /// <summary>
        /// Methods of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol. The default is `HEAD`, the available value are `HEAD` and `GET`.
        /// </summary>
        [Input("healthCheckHttpMethod")]
        public Input<string>? HealthCheckHttpMethod { get; set; }

        /// <summary>
        /// Path of health check. NOTES: Only supports listeners of `HTTP` and `HTTPS` protocol.
        /// </summary>
        [Input("healthCheckHttpPath")]
        public Input<string>? HealthCheckHttpPath { get; set; }

        /// <summary>
        /// Interval time of health check. Valid value ranges: (2~300) sec. and the default is `5` sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckIntervalTime")]
        public Input<int>? HealthCheckIntervalTime { get; set; }

        /// <summary>
        /// Indicates whether health check is enabled.
        /// </summary>
        [Input("healthCheckSwitch")]
        public Input<bool>? HealthCheckSwitch { get; set; }

        /// <summary>
        /// Time out of health check. The value range is 2-60.
        /// </summary>
        [Input("healthCheckTimeOut")]
        public Input<int>? HealthCheckTimeOut { get; set; }

        /// <summary>
        /// Type of health check. Valid value is `CUSTOM`, `TCP`, `HTTP`.
        /// </summary>
        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is [2-10].  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("healthCheckUnhealthNum")]
        public Input<int>? HealthCheckUnhealthNum { get; set; }

        /// <summary>
        /// Indicate to apply HTTP2.0 protocol or not.
        /// </summary>
        [Input("http2Switch")]
        public Input<bool>? Http2Switch { get; set; }

        /// <summary>
        /// ID of CLB listener.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// Whether to enable QUIC. Note: QUIC can be enabled only for HTTPS domain names.
        /// </summary>
        [Input("quic")]
        public Input<bool>? Quic { get; set; }

        /// <summary>
        /// ID of this CLB listener rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// Scheduling method of the CLB listener rules. Valid values: `WRR`, `IP HASH`, `LEAST_CONN`. The default is `WRR`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as `WRR`, and not available when listener protocol is `TCP_SSL`.  NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in `tencentcloud.Clb.ListenerRule`.
        /// </summary>
        [Input("sessionExpireTime")]
        public Input<int>? SessionExpireTime { get; set; }

        /// <summary>
        /// Backend target type. Valid values: `NODE`, `TARGETGROUP`. `NODE` means to bind ordinary nodes, `TARGETGROUP` means to bind target group.
        /// </summary>
        [Input("targetType")]
        public Input<string>? TargetType { get; set; }

        /// <summary>
        /// Url of the listener rule.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ListenerRuleState()
        {
        }
    }
}
