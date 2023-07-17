// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetRouteConflictsRouteConflictSetConflictSetResult
    {
        /// <summary>
        /// create time.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// destination cidr block.
        /// </summary>
        public readonly string DestinationCidrBlock;
        /// <summary>
        /// Destination of Ipv6 Cidr Block.
        /// </summary>
        public readonly string DestinationIpv6CidrBlock;
        /// <summary>
        /// if enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// next hop id.
        /// </summary>
        public readonly string GatewayId;
        /// <summary>
        /// next gateway type.
        /// </summary>
        public readonly string GatewayType;
        /// <summary>
        /// if published To ccn.
        /// </summary>
        public readonly bool PublishedToVbc;
        /// <summary>
        /// route description.
        /// </summary>
        public readonly string RouteDescription;
        /// <summary>
        /// route id.
        /// </summary>
        public readonly int RouteId;
        /// <summary>
        /// unique policy id.
        /// </summary>
        public readonly string RouteItemId;
        /// <summary>
        /// Routing table instance ID, for example:rtb-azd4dt1c.
        /// </summary>
        public readonly string RouteTableId;
        /// <summary>
        /// routr type.
        /// </summary>
        public readonly string RouteType;

        [OutputConstructor]
        private GetRouteConflictsRouteConflictSetConflictSetResult(
            string createdTime,

            string destinationCidrBlock,

            string destinationIpv6CidrBlock,

            bool enabled,

            string gatewayId,

            string gatewayType,

            bool publishedToVbc,

            string routeDescription,

            int routeId,

            string routeItemId,

            string routeTableId,

            string routeType)
        {
            CreatedTime = createdTime;
            DestinationCidrBlock = destinationCidrBlock;
            DestinationIpv6CidrBlock = destinationIpv6CidrBlock;
            Enabled = enabled;
            GatewayId = gatewayId;
            GatewayType = gatewayType;
            PublishedToVbc = publishedToVbc;
            RouteDescription = routeDescription;
            RouteId = routeId;
            RouteItemId = routeItemId;
            RouteTableId = routeTableId;
            RouteType = routeType;
        }
    }
}
