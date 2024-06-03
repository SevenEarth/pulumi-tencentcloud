// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb.Inputs
{

    public sealed class TargetGroupTargetGroupInstanceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The internal ip of target group instance.
        /// </summary>
        [Input("bindIp", required: true)]
        public Input<string> BindIp { get; set; } = null!;

        /// <summary>
        /// The new port of target group instance.
        /// </summary>
        [Input("newPort")]
        public Input<int>? NewPort { get; set; }

        /// <summary>
        /// The port of target group instance.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The weight of target group instance.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TargetGroupTargetGroupInstanceGetArgs()
        {
        }
        public static new TargetGroupTargetGroupInstanceGetArgs Empty => new TargetGroupTargetGroupInstanceGetArgs();
    }
}
