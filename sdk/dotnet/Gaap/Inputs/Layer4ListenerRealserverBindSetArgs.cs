// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap.Inputs
{

    public sealed class Layer4ListenerRealserverBindSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the GAAP realserver.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// IP of the GAAP realserver.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Port of the GAAP realserver.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Scheduling weight, default value is `1`. The range of values is [1,100].
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public Layer4ListenerRealserverBindSetArgs()
        {
        }
    }
}
