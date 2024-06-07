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

    public sealed class DeployContainerGroupHealthCheckSettingsLivenessProbeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The health check method. HTTP: checks through an HTTP interface; CMD: checks by executing a command; TCP: checks by establishing a TCP connection. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// The command to be executed for command health checks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// The number of consecutive successful health checks required for the backend container to transition from success to failure. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        /// <summary>
        /// The time delay for the container to start the health check. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("initialDelaySeconds")]
        public Input<int>? InitialDelaySeconds { get; set; }

        /// <summary>
        /// The request path for HTTP health checks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The time interval for performing health checks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        /// <summary>
        /// The port used for health checks, ranging from 1 to 65535. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol used for HTTP health checks. HTTP and HTTPS are supported. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// The number of consecutive successful health checks required for the backend container to transition from failure to success. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        /// <summary>
        /// The maximum timeout period for each health check response. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        /// <summary>
        /// The type of readiness probe. TSF_DEFAULT represents the default readiness probe of TSF, while K8S_NATIVE represents the native readiness probe of Kubernetes. If this field is not specified, the native readiness probe of Kubernetes is used by default. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DeployContainerGroupHealthCheckSettingsLivenessProbeArgs()
        {
        }
        public static new DeployContainerGroupHealthCheckSettingsLivenessProbeArgs Empty => new DeployContainerGroupHealthCheckSettingsLivenessProbeArgs();
    }
}
