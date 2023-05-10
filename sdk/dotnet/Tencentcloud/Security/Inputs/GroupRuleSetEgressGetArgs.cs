// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Security.Inputs
{

    public sealed class GroupRuleSetEgressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Specify Group ID of Address template like `ipmg-xxxxxxxx`, conflict with `source_security_id` and `cidr_block`.
        /// </summary>
        [Input("addressTemplateGroup")]
        public Input<string>? AddressTemplateGroup { get; set; }

        /// <summary>
        /// Specify Address template ID like `ipm-xxxxxxxx`, conflict with `source_security_id` and `cidr_block`.
        /// </summary>
        [Input("addressTemplateId")]
        public Input<string>? AddressTemplateId { get; set; }

        /// <summary>
        /// An IP address network or CIDR segment. NOTE: `cidr_block`, `ipv6_cidr_block`, `source_security_id` and `address_template_*` are exclusive and cannot be set in the same time.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Description of the security group rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An IPV6 address network or CIDR segment, and conflict with `source_security_id` and `address_template_*`.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and conflicts with `service_template_*`.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `service_template_*`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Specify Group ID of Protocol template ID like `ppmg-xxxxxxxx`, conflict with `cidr_block` and `port`.
        /// </summary>
        [Input("serviceTemplateGroup")]
        public Input<string>? ServiceTemplateGroup { get; set; }

        /// <summary>
        /// Specify Protocol template ID like `ppm-xxxxxxxx`, conflict with `cidr_block` and `port`.
        /// </summary>
        [Input("serviceTemplateId")]
        public Input<string>? ServiceTemplateId { get; set; }

        /// <summary>
        /// ID of the nested security group, and conflicts with `cidr_block` and `address_template_*`.
        /// </summary>
        [Input("sourceSecurityId")]
        public Input<string>? SourceSecurityId { get; set; }

        public GroupRuleSetEgressGetArgs()
        {
        }
    }
}