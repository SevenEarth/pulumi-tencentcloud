// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Inputs
{

    public sealed class HourdbInstanceResourceTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// tag key.
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// tag value.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

        public HourdbInstanceResourceTagArgs()
        {
        }
        public static new HourdbInstanceResourceTagArgs Empty => new HourdbInstanceResourceTagArgs();
    }
}
