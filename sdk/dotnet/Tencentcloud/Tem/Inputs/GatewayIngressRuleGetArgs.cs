// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class GatewayIngressRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// host name.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// rule payload.
        /// </summary>
        [Input("http", required: true)]
        public Input<Inputs.GatewayIngressRuleHttpGetArgs> Http { get; set; } = null!;

        /// <summary>
        /// protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public GatewayIngressRuleGetArgs()
        {
        }
    }
}
