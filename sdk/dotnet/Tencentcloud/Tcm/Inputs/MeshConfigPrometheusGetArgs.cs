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

    public sealed class MeshConfigPrometheusGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom prometheus.
        /// </summary>
        [Input("customProm")]
        public Input<Inputs.MeshConfigPrometheusCustomPromGetArgs>? CustomProm { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Vpc id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public MeshConfigPrometheusGetArgs()
        {
        }
        public static new MeshConfigPrometheusGetArgs Empty => new MeshConfigPrometheusGetArgs();
    }
}
