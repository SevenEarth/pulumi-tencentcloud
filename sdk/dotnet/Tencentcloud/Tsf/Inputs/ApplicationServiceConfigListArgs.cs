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

    public sealed class ApplicationServiceConfigListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Health check configuration.
        /// </summary>
        [Input("healthCheck")]
        public Input<Inputs.ApplicationServiceConfigListHealthCheckArgs>? HealthCheck { get; set; }

        /// <summary>
        /// Service name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("ports", required: true)]
        private InputList<Inputs.ApplicationServiceConfigListPortArgs>? _ports;

        /// <summary>
        /// List of port information.
        /// </summary>
        public InputList<Inputs.ApplicationServiceConfigListPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ApplicationServiceConfigListPortArgs>());
            set => _ports = value;
        }

        public ApplicationServiceConfigListArgs()
        {
        }
    }
}
