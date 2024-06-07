// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb.Inputs
{

    public sealed class InstanceStandbyInstanceListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the ID of standby instance.
        /// </summary>
        [Input("standbyInstanceId")]
        public Input<string>? StandbyInstanceId { get; set; }

        /// <summary>
        /// Indicates the region of standby instance.
        /// </summary>
        [Input("standbyInstanceRegion")]
        public Input<string>? StandbyInstanceRegion { get; set; }

        public InstanceStandbyInstanceListGetArgs()
        {
        }
        public static new InstanceStandbyInstanceListGetArgs Empty => new InstanceStandbyInstanceListGetArgs();
    }
}
