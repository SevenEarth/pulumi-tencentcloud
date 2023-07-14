// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ccn.Outputs
{

    [OutputType]
    public sealed class InstancesRejectAttachInstance
    {
        /// <summary>
        /// Description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Attachment Instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Instance Region.
        /// </summary>
        public readonly string InstanceRegion;
        /// <summary>
        /// InstanceType: `VPC`, `DIRECTCONNECT`, `BMVPC`, `VPNGW`.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// ID of the routing table associated with the instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? RouteTableId;

        [OutputConstructor]
        private InstancesRejectAttachInstance(
            string? description,

            string instanceId,

            string instanceRegion,

            string? instanceType,

            string? routeTableId)
        {
            Description = description;
            InstanceId = instanceId;
            InstanceRegion = instanceRegion;
            InstanceType = instanceType;
            RouteTableId = routeTableId;
        }
    }
}
