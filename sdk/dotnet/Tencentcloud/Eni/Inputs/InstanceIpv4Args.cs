// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eni.Inputs
{

    public sealed class InstanceIpv4Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the IP, maximum length 25.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Intranet IP.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Indicates whether the IP is primary.
        /// </summary>
        [Input("primary", required: true)]
        public Input<bool> Primary { get; set; } = null!;

        public InstanceIpv4Args()
        {
        }
        public static new InstanceIpv4Args Empty => new InstanceIpv4Args();
    }
}
