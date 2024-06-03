// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class GetProductEventDimensionInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance dimension name, eg: `deviceWanIp` for internet ip.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Instance dimension value, eg: `119.119.119.119` for internet ip.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GetProductEventDimensionInputArgs()
        {
        }
        public static new GetProductEventDimensionInputArgs Empty => new GetProductEventDimensionInputArgs();
    }
}
