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
    /// Provides a resource to create a VPN gateway route.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var route = new Tencentcloud.Vpn.GatewayRoute("route", new()
    ///     {
    ///         DestinationCidrBlock = "10.0.0.0/16",
    ///         InstanceId = "vpnx-5b5dmao3",
    ///         InstanceType = "VPNCONN",
    ///         Priority = 100,
    ///         Status = "DISABLE",
    ///         VpnGatewayId = "vpngw-ak9sjem2",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// VPN gateway route can be imported using the id, the id format must be '{vpn_gateway_id}#{route_id}', e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Vpn/gatewayRoute:GatewayRoute route1 vpngw-ak9sjem2#vpngw-8ccsnclt
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpn/gatewayRoute:GatewayRoute")]
    public partial class GatewayRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Destination IDC IP range.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Instance ID of the next hop.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Priority. Valid values: 0 and 100.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Route ID.
        /// </summary>
        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        /// <summary>
        /// Status. Valid values: ENABLE and DISABLE.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Route type. Default value: Static.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// VPN gateway ID.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string> VpnGatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayRoute(string name, GatewayRouteArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gatewayRoute:GatewayRoute", name, args ?? new GatewayRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayRoute(string name, Input<string> id, GatewayRouteState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gatewayRoute:GatewayRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewayRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayRoute Get(string name, Input<string> id, GatewayRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewayRoute(name, id, state, options);
        }
    }

    public sealed class GatewayRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination IDC IP range.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// Instance ID of the next hop.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Priority. Valid values: 0 and 100.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        /// <summary>
        /// Status. Valid values: ENABLE and DISABLE.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// VPN gateway ID.
        /// </summary>
        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public GatewayRouteArgs()
        {
        }
        public static new GatewayRouteArgs Empty => new GatewayRouteArgs();
    }

    public sealed class GatewayRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Destination IDC IP range.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// Instance ID of the next hop.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Next hop type (type of the associated instance). Valid values: VPNCONN (VPN tunnel) and CCN (CCN instance).
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Priority. Valid values: 0 and 100.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Route ID.
        /// </summary>
        [Input("routeId")]
        public Input<string>? RouteId { get; set; }

        /// <summary>
        /// Status. Valid values: ENABLE and DISABLE.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Route type. Default value: Static.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// VPN gateway ID.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public GatewayRouteState()
        {
        }
        public static new GatewayRouteState Empty => new GatewayRouteState();
    }
}
