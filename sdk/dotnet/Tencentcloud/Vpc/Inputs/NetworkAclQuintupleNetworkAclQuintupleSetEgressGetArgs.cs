// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Inputs
{

    public sealed class NetworkAclQuintupleNetworkAclQuintupleSetEgressGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action, ACCEPT or DROP.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// Creation time, used as an output parameter of DescribeNetworkAclQuintupleEntries.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination CIDR.
        /// </summary>
        [Input("destinationCidr")]
        public Input<string>? DestinationCidr { get; set; }

        /// <summary>
        /// Destination port (all, single port, range). When Protocol is ALL or ICMP, Port cannot be specified.
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Direction, INGRESS or EGRESS, is used as an output parameter of DescribeNetworkAclQuintupleEntries.
        /// </summary>
        [Input("networkAclDirection")]
        public Input<string>? NetworkAclDirection { get; set; }

        /// <summary>
        /// Unique ID of a network ACL entry.
        /// </summary>
        [Input("networkAclQuintupleEntryId")]
        public Input<string>? NetworkAclQuintupleEntryId { get; set; }

        /// <summary>
        /// Priority, starting from 1.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Protocol, value: TCP,UDP, ICMP, ALL.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Source CIDR.
        /// </summary>
        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        /// <summary>
        /// Source port (all, single port, range). When Protocol is ALL or ICMP, Port cannot be specified.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        public NetworkAclQuintupleNetworkAclQuintupleSetEgressGetArgs()
        {
        }
        public static new NetworkAclQuintupleNetworkAclQuintupleSetEgressGetArgs Empty => new NetworkAclQuintupleNetworkAclQuintupleSetEgressGetArgs();
    }
}
