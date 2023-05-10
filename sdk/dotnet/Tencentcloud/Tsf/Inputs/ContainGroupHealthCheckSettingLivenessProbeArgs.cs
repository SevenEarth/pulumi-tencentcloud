// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class ContainGroupHealthCheckSettingLivenessProbeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// health check method. HTTP: check by HTTP interface; CMD: check by executing command; TCP: check by establishing TCP connection.
        /// </summary>
        [Input("actionType")]
        public Input<string>? ActionType { get; set; }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Execute command check mode, the command to execute.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// Indicates the number of consecutive health check successes for the backend container from success to failure.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// The time for the container to delay starting the health check.
        /// </summary>
        [Input("initialDelaySeconds")]
        public Input<int>? InitialDelaySeconds { get; set; }

        /// <summary>
        /// The request path of the HTTP health check interface.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The interval at which health checks are performed.
        /// </summary>
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        /// <summary>
        /// service port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The inspection protocol used by the HTTP health check method. HTTP and HTTPS are supported.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// Indicates the number of consecutive health check successes for the backend container from failure to success.
        /// </summary>
        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        /// <summary>
        /// The maximum timeout for each health check response.
        /// </summary>
        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        /// <summary>
        /// TSF_DEFAULT: tsf default readiness probe. K8S_NATIVE: k8s native probe. If not filled, it defaults to k8s native probe.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContainGroupHealthCheckSettingLivenessProbeArgs()
        {
        }
    }
}
