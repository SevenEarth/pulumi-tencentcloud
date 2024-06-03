// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Inputs
{

    public sealed class ProxyEndPointInstanceWeightArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Instance Weight.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public ProxyEndPointInstanceWeightArgs()
        {
        }
        public static new ProxyEndPointInstanceWeightArgs Empty => new ProxyEndPointInstanceWeightArgs();
    }
}
