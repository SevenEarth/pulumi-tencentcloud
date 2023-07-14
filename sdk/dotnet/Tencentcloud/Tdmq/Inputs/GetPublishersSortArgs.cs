// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq.Inputs
{

    public sealed class GetPublishersSortInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// sorter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Ascending ASC, descending DESC.
        /// </summary>
        [Input("order", required: true)]
        public Input<string> Order { get; set; } = null!;

        public GetPublishersSortInputArgs()
        {
        }
    }
}
