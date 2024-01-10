// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdwpg.Inputs
{

    public sealed class InstanceResourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// resource count.
        /// </summary>
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        /// <summary>
        /// disk Information.
        /// </summary>
        [Input("diskSpec", required: true)]
        public Input<Inputs.InstanceResourceDiskSpecGetArgs> DiskSpec { get; set; } = null!;

        /// <summary>
        /// resource name.
        /// </summary>
        [Input("specName", required: true)]
        public Input<string> SpecName { get; set; } = null!;

        /// <summary>
        /// resource type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public InstanceResourceGetArgs()
        {
        }
    }
}
