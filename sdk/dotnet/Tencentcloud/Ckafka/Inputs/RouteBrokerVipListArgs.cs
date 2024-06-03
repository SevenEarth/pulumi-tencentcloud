// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class RouteBrokerVipListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Virtual IP.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Virtual port.
        /// </summary>
        [Input("vport")]
        public Input<string>? Vport { get; set; }

        public RouteBrokerVipListArgs()
        {
        }
        public static new RouteBrokerVipListArgs Empty => new RouteBrokerVipListArgs();
    }
}
