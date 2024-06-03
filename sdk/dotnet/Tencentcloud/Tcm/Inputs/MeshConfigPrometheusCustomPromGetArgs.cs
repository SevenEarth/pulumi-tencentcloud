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

    public sealed class MeshConfigPrometheusCustomPromGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication type of the prometheus.
        /// </summary>
        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        /// <summary>
        /// Whether it is public address, default false.
        /// </summary>
        [Input("isPublicAddr")]
        public Input<bool>? IsPublicAddr { get; set; }

        /// <summary>
        /// Password of the prometheus, used in basic authentication type.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Url of the prometheus.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Username of the prometheus, used in basic authentication type.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Vpc id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public MeshConfigPrometheusCustomPromGetArgs()
        {
        }
        public static new MeshConfigPrometheusCustomPromGetArgs Empty => new MeshConfigPrometheusCustomPromGetArgs();
    }
}
