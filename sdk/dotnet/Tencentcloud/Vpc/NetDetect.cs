// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provides a resource to create a vpc net_detect
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// vpc net_detect can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/netDetect:NetDetect net_detect net_detect_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/netDetect:NetDetect")]
    public partial class NetDetect : Pulumi.CustomResource
    {
        /// <summary>
        /// An array of probe destination IPv4 addresses. Up to two.
        /// </summary>
        [Output("detectDestinationIps")]
        public Output<ImmutableArray<string>> DetectDestinationIps { get; private set; } = null!;

        /// <summary>
        /// Network probe description.
        /// </summary>
        [Output("netDetectDescription")]
        public Output<string?> NetDetectDescription { get; private set; } = null!;

        /// <summary>
        /// Network probe name, the maximum length cannot exceed 60 bytes.
        /// </summary>
        [Output("netDetectName")]
        public Output<string> NetDetectName { get; private set; } = null!;

        /// <summary>
        /// The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        /// </summary>
        [Output("nextHopDestination")]
        public Output<string?> NextHopDestination { get; private set; } = null!;

        /// <summary>
        /// The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        /// </summary>
        [Output("nextHopType")]
        public Output<string?> NextHopType { get; private set; } = null!;

        /// <summary>
        /// Subnet instance ID. Such as:subnet-12345678.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// `VPC` instance `ID`. Such as:`vpc-12345678`.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a NetDetect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetDetect(string name, NetDetectArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/netDetect:NetDetect", name, args ?? new NetDetectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetDetect(string name, Input<string> id, NetDetectState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/netDetect:NetDetect", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetDetect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetDetect Get(string name, Input<string> id, NetDetectState? state = null, CustomResourceOptions? options = null)
        {
            return new NetDetect(name, id, state, options);
        }
    }

    public sealed class NetDetectArgs : Pulumi.ResourceArgs
    {
        [Input("detectDestinationIps", required: true)]
        private InputList<string>? _detectDestinationIps;

        /// <summary>
        /// An array of probe destination IPv4 addresses. Up to two.
        /// </summary>
        public InputList<string> DetectDestinationIps
        {
            get => _detectDestinationIps ?? (_detectDestinationIps = new InputList<string>());
            set => _detectDestinationIps = value;
        }

        /// <summary>
        /// Network probe description.
        /// </summary>
        [Input("netDetectDescription")]
        public Input<string>? NetDetectDescription { get; set; }

        /// <summary>
        /// Network probe name, the maximum length cannot exceed 60 bytes.
        /// </summary>
        [Input("netDetectName", required: true)]
        public Input<string> NetDetectName { get; set; } = null!;

        /// <summary>
        /// The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        /// </summary>
        [Input("nextHopDestination")]
        public Input<string>? NextHopDestination { get; set; }

        /// <summary>
        /// The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// Subnet instance ID. Such as:subnet-12345678.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// `VPC` instance `ID`. Such as:`vpc-12345678`.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public NetDetectArgs()
        {
        }
    }

    public sealed class NetDetectState : Pulumi.ResourceArgs
    {
        [Input("detectDestinationIps")]
        private InputList<string>? _detectDestinationIps;

        /// <summary>
        /// An array of probe destination IPv4 addresses. Up to two.
        /// </summary>
        public InputList<string> DetectDestinationIps
        {
            get => _detectDestinationIps ?? (_detectDestinationIps = new InputList<string>());
            set => _detectDestinationIps = value;
        }

        /// <summary>
        /// Network probe description.
        /// </summary>
        [Input("netDetectDescription")]
        public Input<string>? NetDetectDescription { get; set; }

        /// <summary>
        /// Network probe name, the maximum length cannot exceed 60 bytes.
        /// </summary>
        [Input("netDetectName")]
        public Input<string>? NetDetectName { get; set; }

        /// <summary>
        /// The destination gateway of the next hop, the value is related to the next hop type. If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678; If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678; If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 10.0.0.12; If the next hop type is CCN, and the value is the cloud network ID, such as: ccn-12345678; If the next hop type is NONEXTHOP, and the specified network probe is a network probe without a next hop.
        /// </summary>
        [Input("nextHopDestination")]
        public Input<string>? NextHopDestination { get; set; }

        /// <summary>
        /// The next hop type, currently we support the following types: `VPN`: VPN gateway; `DIRECTCONNECT`: private line gateway; `PEERCONNECTION`: peer connection; `NAT`: NAT gateway; `NORMAL_CVM`: normal cloud server; `CCN`: cloud networking gateway; `NONEXTHOP`: no next hop.
        /// </summary>
        [Input("nextHopType")]
        public Input<string>? NextHopType { get; set; }

        /// <summary>
        /// Subnet instance ID. Such as:subnet-12345678.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// `VPC` instance `ID`. Such as:`vpc-12345678`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public NetDetectState()
        {
        }
    }
}
