// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse
{
    /// <summary>
    /// Provides a resource to create a tse cngw_gateway
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-4";
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///             AvailabilityZone = availabilityZone,
    ///             CidrBlock = "10.0.1.0/24",
    ///         });
    ///         var cngwGateway = new Tencentcloud.Tse.CngwGateway("cngwGateway", new Tencentcloud.Tse.CngwGatewayArgs
    ///         {
    ///             Description = "terraform test1",
    ///             EnableCls = true,
    ///             EngineRegion = "ap-guangzhou",
    ///             FeatureVersion = "STANDARD",
    ///             GatewayVersion = "2.5.1",
    ///             IngressClassName = "tse-nginx-ingress",
    ///             InternetMaxBandwidthOut = 0,
    ///             TradeType = 0,
    ///             Type = "kong",
    ///             NodeConfig = new Tencentcloud.Tse.Inputs.CngwGatewayNodeConfigArgs
    ///             {
    ///                 Number = 2,
    ///                 Specification = "1c2g",
    ///             },
    ///             VpcConfig = new Tencentcloud.Tse.Inputs.CngwGatewayVpcConfigArgs
    ///             {
    ///                 SubnetId = subnet.Id,
    ///                 VpcId = vpc.Id,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tse/cngwGateway:CngwGateway")]
    public partial class CngwGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// whether to enable CLS log. Default value: fasle.
        /// </summary>
        [Output("enableCls")]
        public Output<bool?> EnableCls { get; private set; } = null!;

        /// <summary>
        /// engine region of gateway.
        /// </summary>
        [Output("engineRegion")]
        public Output<string> EngineRegion { get; private set; } = null!;

        /// <summary>
        /// product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        /// </summary>
        [Output("featureVersion")]
        public Output<string> FeatureVersion { get; private set; } = null!;

        /// <summary>
        /// gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        /// </summary>
        [Output("gatewayVersion")]
        public Output<string> GatewayVersion { get; private set; } = null!;

        /// <summary>
        /// ingress class name.
        /// </summary>
        [Output("ingressClassName")]
        public Output<string> IngressClassName { get; private set; } = null!;

        /// <summary>
        /// Port information that the instance listens to.
        /// </summary>
        [Output("instancePorts")]
        public Output<ImmutableArray<Outputs.CngwGatewayInstancePort>> InstancePorts { get; private set; } = null!;

        /// <summary>
        /// internet configration.
        /// </summary>
        [Output("internetConfig")]
        public Output<Outputs.CngwGatewayInternetConfig?> InternetConfig { get; private set; } = null!;

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Output("internetMaxBandwidthOut")]
        public Output<int?> InternetMaxBandwidthOut { get; private set; } = null!;

        /// <summary>
        /// gateway name, supports up to 60 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// gateway node configration.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.CngwGatewayNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// Public IP address list.
        /// </summary>
        [Output("publicIpAddresses")]
        public Output<ImmutableArray<string>> PublicIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        /// </summary>
        [Output("tradeType")]
        public Output<int?> TradeType { get; private set; } = null!;

        /// <summary>
        /// gateway type,currently only supports kong.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// vpc information.
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.CngwGatewayVpcConfig> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a CngwGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CngwGateway(string name, CngwGatewayArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwGateway:CngwGateway", name, args ?? new CngwGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CngwGateway(string name, Input<string> id, CngwGatewayState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwGateway:CngwGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CngwGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CngwGateway Get(string name, Input<string> id, CngwGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new CngwGateway(name, id, state, options);
        }
    }

    public sealed class CngwGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to enable CLS log. Default value: fasle.
        /// </summary>
        [Input("enableCls")]
        public Input<bool>? EnableCls { get; set; }

        /// <summary>
        /// engine region of gateway.
        /// </summary>
        [Input("engineRegion")]
        public Input<string>? EngineRegion { get; set; }

        /// <summary>
        /// product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        /// </summary>
        [Input("featureVersion")]
        public Input<string>? FeatureVersion { get; set; }

        /// <summary>
        /// gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        /// </summary>
        [Input("gatewayVersion", required: true)]
        public Input<string> GatewayVersion { get; set; } = null!;

        /// <summary>
        /// ingress class name.
        /// </summary>
        [Input("ingressClassName")]
        public Input<string>? IngressClassName { get; set; }

        /// <summary>
        /// internet configration.
        /// </summary>
        [Input("internetConfig")]
        public Input<Inputs.CngwGatewayInternetConfigArgs>? InternetConfig { get; set; }

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// gateway name, supports up to 60 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// gateway node configration.
        /// </summary>
        [Input("nodeConfig", required: true)]
        public Input<Inputs.CngwGatewayNodeConfigArgs> NodeConfig { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        /// </summary>
        [Input("tradeType")]
        public Input<int>? TradeType { get; set; }

        /// <summary>
        /// gateway type,currently only supports kong.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// vpc information.
        /// </summary>
        [Input("vpcConfig", required: true)]
        public Input<Inputs.CngwGatewayVpcConfigArgs> VpcConfig { get; set; } = null!;

        public CngwGatewayArgs()
        {
        }
    }

    public sealed class CngwGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// description information, up to 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to enable CLS log. Default value: fasle.
        /// </summary>
        [Input("enableCls")]
        public Input<bool>? EnableCls { get; set; }

        /// <summary>
        /// engine region of gateway.
        /// </summary>
        [Input("engineRegion")]
        public Input<string>? EngineRegion { get; set; }

        /// <summary>
        /// product version. Reference value: `TRIAL`, `STANDARD`(default value), `PROFESSIONAL`.
        /// </summary>
        [Input("featureVersion")]
        public Input<string>? FeatureVersion { get; set; }

        /// <summary>
        /// gateway vwersion. Reference value: `2.4.1`, `2.5.1`.
        /// </summary>
        [Input("gatewayVersion")]
        public Input<string>? GatewayVersion { get; set; }

        /// <summary>
        /// ingress class name.
        /// </summary>
        [Input("ingressClassName")]
        public Input<string>? IngressClassName { get; set; }

        [Input("instancePorts")]
        private InputList<Inputs.CngwGatewayInstancePortGetArgs>? _instancePorts;

        /// <summary>
        /// Port information that the instance listens to.
        /// </summary>
        public InputList<Inputs.CngwGatewayInstancePortGetArgs> InstancePorts
        {
            get => _instancePorts ?? (_instancePorts = new InputList<Inputs.CngwGatewayInstancePortGetArgs>());
            set => _instancePorts = value;
        }

        /// <summary>
        /// internet configration.
        /// </summary>
        [Input("internetConfig")]
        public Input<Inputs.CngwGatewayInternetConfigGetArgs>? InternetConfig { get; set; }

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// gateway name, supports up to 60 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// gateway node configration.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.CngwGatewayNodeConfigGetArgs>? NodeConfig { get; set; }

        [Input("publicIpAddresses")]
        private InputList<string>? _publicIpAddresses;

        /// <summary>
        /// Public IP address list.
        /// </summary>
        public InputList<string> PublicIpAddresses
        {
            get => _publicIpAddresses ?? (_publicIpAddresses = new InputList<string>());
            set => _publicIpAddresses = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// trade type. Reference value: `0`: postpaid, `1`:Prepaid (Interface does not support the creation of prepaid instances yet).
        /// </summary>
        [Input("tradeType")]
        public Input<int>? TradeType { get; set; }

        /// <summary>
        /// gateway type,currently only supports kong.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// vpc information.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.CngwGatewayVpcConfigGetArgs>? VpcConfig { get; set; }

        public CngwGatewayState()
        {
        }
    }
}
