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
    /// Provides a resource to create a vpc ipv6_eni_address
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/ipv6EniAddress:Ipv6EniAddress")]
    public partial class Ipv6EniAddress : Pulumi.CustomResource
    {
        /// <summary>
        /// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<Outputs.Ipv6EniAddressIpv6Address>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// ENI instance `ID`, in the form of `eni-m6dyj72l`.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// VPC `ID`, in the form of `vpc-m6dyj72l`.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv6EniAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv6EniAddress(string name, Ipv6EniAddressArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6EniAddress:Ipv6EniAddress", name, args ?? new Ipv6EniAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv6EniAddress(string name, Input<string> id, Ipv6EniAddressState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/ipv6EniAddress:Ipv6EniAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv6EniAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv6EniAddress Get(string name, Input<string> id, Ipv6EniAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv6EniAddress(name, id, state, options);
        }
    }

    public sealed class Ipv6EniAddressArgs : Pulumi.ResourceArgs
    {
        [Input("ipv6Addresses")]
        private InputList<Inputs.Ipv6EniAddressIpv6AddressArgs>? _ipv6Addresses;

        /// <summary>
        /// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        /// </summary>
        public InputList<Inputs.Ipv6EniAddressIpv6AddressArgs> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<Inputs.Ipv6EniAddressIpv6AddressArgs>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// ENI instance `ID`, in the form of `eni-m6dyj72l`.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        /// <summary>
        /// VPC `ID`, in the form of `vpc-m6dyj72l`.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public Ipv6EniAddressArgs()
        {
        }
    }

    public sealed class Ipv6EniAddressState : Pulumi.ResourceArgs
    {
        [Input("ipv6Addresses")]
        private InputList<Inputs.Ipv6EniAddressIpv6AddressGetArgs>? _ipv6Addresses;

        /// <summary>
        /// The specified `IPv6` address list, up to 10 can be specified at a time. Combined with the input parameter `Ipv6AddressCount` to calculate the quota. Mandatory one with Ipv6AddressCount.
        /// </summary>
        public InputList<Inputs.Ipv6EniAddressIpv6AddressGetArgs> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<Inputs.Ipv6EniAddressIpv6AddressGetArgs>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// ENI instance `ID`, in the form of `eni-m6dyj72l`.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// VPC `ID`, in the form of `vpc-m6dyj72l`.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public Ipv6EniAddressState()
        {
        }
    }
}
