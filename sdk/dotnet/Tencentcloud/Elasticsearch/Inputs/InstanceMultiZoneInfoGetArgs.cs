// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Inputs
{

    public sealed class InstanceMultiZoneInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Availability zone.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The ID of a VPC subnetwork.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public InstanceMultiZoneInfoGetArgs()
        {
        }
        public static new InstanceMultiZoneInfoGetArgs Empty => new InstanceMultiZoneInfoGetArgs();
    }
}
