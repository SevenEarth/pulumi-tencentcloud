// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sidecar inject configuration.
        /// </summary>
        [Input("inject")]
        public Input<Inputs.MeshConfigInjectArgs>? Inject { get; set; }

        /// <summary>
        /// Istio configuration.
        /// </summary>
        [Input("istio")]
        public Input<Inputs.MeshConfigIstioArgs>? Istio { get; set; }

        /// <summary>
        /// Prometheus configuration.
        /// </summary>
        [Input("prometheus")]
        public Input<Inputs.MeshConfigPrometheusArgs>? Prometheus { get; set; }

        /// <summary>
        /// Default sidecar requests and limits.
        /// </summary>
        [Input("sidecarResources")]
        public Input<Inputs.MeshConfigSidecarResourcesArgs>? SidecarResources { get; set; }

        /// <summary>
        /// Tracing config.
        /// </summary>
        [Input("tracing")]
        public Input<Inputs.MeshConfigTracingArgs>? Tracing { get; set; }

        public MeshConfigArgs()
        {
        }
        public static new MeshConfigArgs Empty => new MeshConfigArgs();
    }
}
