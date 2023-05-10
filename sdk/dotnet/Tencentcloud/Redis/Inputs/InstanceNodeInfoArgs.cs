// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis.Inputs
{

    public sealed class InstanceNodeInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the master or replica node.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Indicates whether the node is master.
        /// </summary>
        [Input("master")]
        public Input<bool>? Master { get; set; }

        /// <summary>
        /// ID of the availability zone of the master or replica node.
        /// </summary>
        [Input("zoneId")]
        public Input<int>? ZoneId { get; set; }

        public InstanceNodeInfoArgs()
        {
        }
    }
}
