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

    public sealed class NetworkAclQuintupleNetworkAclQuintupleSetIngressGetArgs : Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationCidr")]
        public Input<string>? DestinationCidr { get; set; }

        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        [Input("networkAclDirection")]
        public Input<string>? NetworkAclDirection { get; set; }

        [Input("networkAclQuintupleEntryId")]
        public Input<string>? NetworkAclQuintupleEntryId { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("sourceCidr")]
        public Input<string>? SourceCidr { get; set; }

        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        public NetworkAclQuintupleNetworkAclQuintupleSetIngressGetArgs()
        {
        }
    }
}
