// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpn
{
    /// <summary>
    /// ## Example Usage
    /// ### VPC SSL VPN gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myCgw = new Tencentcloud.Vpn.Gateway("myCgw", new Tencentcloud.Vpn.GatewayArgs
    ///         {
    ///             Bandwidth = 5,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///             Type = "SSL",
    ///             VpcId = "vpc-86v957zb",
    ///             Zone = "ap-guangzhou-3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### CCN IPEC VPN gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myCgw = new Tencentcloud.Vpn.Gateway("myCgw", new Tencentcloud.Vpn.GatewayArgs
    ///         {
    ///             Bandwidth = 5,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///             Type = "CCN",
    ///             Zone = "ap-guangzhou-3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### CCN SSL VPN gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myCgw = new Tencentcloud.Vpn.Gateway("myCgw", new Tencentcloud.Vpn.GatewayArgs
    ///         {
    ///             Bandwidth = 5,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///             Type = "SSL_CCN",
    ///             Zone = "ap-guangzhou-3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### POSTPAID_BY_HOUR VPN gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myCgw = new Tencentcloud.Vpn.Gateway("myCgw", new Tencentcloud.Vpn.GatewayArgs
    ///         {
    ///             Bandwidth = 5,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///             VpcId = "vpc-dk8zmwuf",
    ///             Zone = "ap-guangzhou-3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### PREPAID VPN gateway
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myCgw = new Tencentcloud.Vpn.Gateway("myCgw", new Tencentcloud.Vpn.GatewayArgs
    ///         {
    ///             Bandwidth = 5,
    ///             ChargeType = "PREPAID",
    ///             PrepaidPeriod = 1,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///             VpcId = "vpc-dk8zmwuf",
    ///             Zone = "ap-guangzhou-3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPN gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpn/gateway:Gateway foo vpngw-8ccsnclt
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpn/gateway:Gateway")]
    public partial class Gateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is unsupported.
        /// </summary>
        [Output("bandwidth")]
        public Output<int?> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Output("cdcId")]
        public Output<string> CdcId { get; private set; } = null!;

        /// <summary>
        /// Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Create time of the VPN gateway.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Expired time of the VPN gateway when charge type is `PREPAID`.
        /// </summary>
        [Output("expiredTime")]
        public Output<string> ExpiredTime { get; private set; } = null!;

        /// <summary>
        /// Indicates whether ip address is blocked.
        /// </summary>
        [Output("isAddressBlocked")]
        public Output<bool> IsAddressBlocked { get; private set; } = null!;

        /// <summary>
        /// Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter is only required for SSL VPN gateways.
        /// </summary>
        [Output("maxConnection")]
        public Output<int> MaxConnection { get; private set; } = null!;

        /// <summary>
        /// Name of the VPN gateway. The length of character is limited to 1-60.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The plan of new purchase. Valid value: `PREPAID_TO_POSTPAID`.
        /// </summary>
        [Output("newPurchasePlan")]
        public Output<string> NewPurchasePlan { get; private set; } = null!;

        /// <summary>
        /// Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid period. This para can only be changed on `IPSEC` vpn gateway.
        /// </summary>
        [Output("prepaidPeriod")]
        public Output<int?> PrepaidPeriod { get; private set; } = null!;

        /// <summary>
        /// Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_AUTO_RENEW`, `NOTIFY_AND_MANUAL_RENEW`.
        /// </summary>
        [Output("prepaidRenewFlag")]
        public Output<string?> PrepaidRenewFlag { get; private set; } = null!;

        /// <summary>
        /// Public IP of the VPN gateway.
        /// </summary>
        [Output("publicIpAddress")]
        public Output<string> PublicIpAddress { get; private set; } = null!;

        /// <summary>
        /// Restrict state of gateway. Valid value: `PRETECIVELY_ISOLATED`, `NORMAL`.
        /// </summary>
        [Output("restrictState")]
        public Output<string> RestrictState { get; private set; } = null!;

        /// <summary>
        /// State of the VPN gateway. Valid value: `PENDING`, `DELETING`, `AVAILABLE`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of gateway instance, Default is `IPSEC`. Valid value: `IPSEC`, `SSL`, `CCN` and `SSL_CCN`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// ID of the VPC. Required if vpn gateway is not in `CCN` or `SSL_CCN` type, and doesn't make sense for `CCN` or `SSL_CCN` vpn gateway.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Zone of the VPN gateway.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Gateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Gateway(string name, GatewayArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Gateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Gateway Get(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Gateway(name, id, state, options);
        }
    }

    public sealed class GatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is unsupported.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        /// <summary>
        /// Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter is only required for SSL VPN gateways.
        /// </summary>
        [Input("maxConnection")]
        public Input<int>? MaxConnection { get; set; }

        /// <summary>
        /// Name of the VPN gateway. The length of character is limited to 1-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid period. This para can only be changed on `IPSEC` vpn gateway.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_AUTO_RENEW`, `NOTIFY_AND_MANUAL_RENEW`.
        /// </summary>
        [Input("prepaidRenewFlag")]
        public Input<string>? PrepaidRenewFlag { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of gateway instance, Default is `IPSEC`. Valid value: `IPSEC`, `SSL`, `CCN` and `SSL_CCN`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// ID of the VPC. Required if vpn gateway is not in `CCN` or `SSL_CCN` type, and doesn't make sense for `CCN` or `SSL_CCN` vpn gateway.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Zone of the VPN gateway.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GatewayArgs()
        {
        }
    }

    public sealed class GatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is unsupported.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        /// <summary>
        /// Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Create time of the VPN gateway.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Expired time of the VPN gateway when charge type is `PREPAID`.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        /// <summary>
        /// Indicates whether ip address is blocked.
        /// </summary>
        [Input("isAddressBlocked")]
        public Input<bool>? IsAddressBlocked { get; set; }

        /// <summary>
        /// Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter is only required for SSL VPN gateways.
        /// </summary>
        [Input("maxConnection")]
        public Input<int>? MaxConnection { get; set; }

        /// <summary>
        /// Name of the VPN gateway. The length of character is limited to 1-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The plan of new purchase. Valid value: `PREPAID_TO_POSTPAID`.
        /// </summary>
        [Input("newPurchasePlan")]
        public Input<string>? NewPurchasePlan { get; set; }

        /// <summary>
        /// Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid period. This para can only be changed on `IPSEC` vpn gateway.
        /// </summary>
        [Input("prepaidPeriod")]
        public Input<int>? PrepaidPeriod { get; set; }

        /// <summary>
        /// Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_AUTO_RENEW`, `NOTIFY_AND_MANUAL_RENEW`.
        /// </summary>
        [Input("prepaidRenewFlag")]
        public Input<string>? PrepaidRenewFlag { get; set; }

        /// <summary>
        /// Public IP of the VPN gateway.
        /// </summary>
        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        /// <summary>
        /// Restrict state of gateway. Valid value: `PRETECIVELY_ISOLATED`, `NORMAL`.
        /// </summary>
        [Input("restrictState")]
        public Input<string>? RestrictState { get; set; }

        /// <summary>
        /// State of the VPN gateway. Valid value: `PENDING`, `DELETING`, `AVAILABLE`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A list of tags used to associate different resources.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of gateway instance, Default is `IPSEC`. Valid value: `IPSEC`, `SSL`, `CCN` and `SSL_CCN`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// ID of the VPC. Required if vpn gateway is not in `CCN` or `SSL_CCN` type, and doesn't make sense for `CCN` or `SSL_CCN` vpn gateway.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Zone of the VPN gateway.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GatewayState()
        {
        }
    }
}
