// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class DdosPolicyDdosRulePacketFilterPacketFilterGetArgs : Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("depth")]
        public Input<int>? Depth { get; set; }

        [Input("depth2")]
        public Input<int>? Depth2 { get; set; }

        [Input("dportEnd")]
        public Input<int>? DportEnd { get; set; }

        [Input("dportStart")]
        public Input<int>? DportStart { get; set; }

        [Input("isNot")]
        public Input<int>? IsNot { get; set; }

        [Input("isNot2")]
        public Input<int>? IsNot2 { get; set; }

        [Input("matchBegin")]
        public Input<string>? MatchBegin { get; set; }

        [Input("matchBegin2")]
        public Input<string>? MatchBegin2 { get; set; }

        [Input("matchLogic")]
        public Input<string>? MatchLogic { get; set; }

        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        [Input("matchType2")]
        public Input<string>? MatchType2 { get; set; }

        [Input("offset")]
        public Input<int>? Offset { get; set; }

        [Input("offset2")]
        public Input<int>? Offset2 { get; set; }

        [Input("packetMax")]
        public Input<int>? PacketMax { get; set; }

        [Input("packetMin")]
        public Input<int>? PacketMin { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("sportEnd")]
        public Input<int>? SportEnd { get; set; }

        [Input("sportStart")]
        public Input<int>? SportStart { get; set; }

        [Input("str")]
        public Input<string>? Str { get; set; }

        [Input("str2")]
        public Input<string>? Str2 { get; set; }

        public DdosPolicyDdosRulePacketFilterPacketFilterGetArgs()
        {
        }
    }
}
