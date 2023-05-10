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
    /// Provides a resource to create a vpc vpn_gateway_ssl_client_cert
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
    ///         var vpnGatewaySslClientCert = new Tencentcloud.Vpn.GatewaySslClientCert("vpnGatewaySslClientCert", new Tencentcloud.Vpn.GatewaySslClientCertArgs
    ///         {
    ///             SslVpnClientId = "vpnc-123456",
    ///             Switch = "off",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc vpn_gateway_ssl_client_cert can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert vpn_gateway_ssl_client_cert ssl_client_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert")]
    public partial class GatewaySslClientCert : Pulumi.CustomResource
    {
        /// <summary>
        /// SSL-VPN-CLIENT Instance ID.
        /// </summary>
        [Output("sslVpnClientId")]
        public Output<string> SslVpnClientId { get; private set; } = null!;

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Output("switch")]
        public Output<string?> Switch { get; private set; } = null!;


        /// <summary>
        /// Create a GatewaySslClientCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewaySslClientCert(string name, GatewaySslClientCertArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert", name, args ?? new GatewaySslClientCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewaySslClientCert(string name, Input<string> id, GatewaySslClientCertState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewaySslClientCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewaySslClientCert Get(string name, Input<string> id, GatewaySslClientCertState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewaySslClientCert(name, id, state, options);
        }
    }

    public sealed class GatewaySslClientCertArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSL-VPN-CLIENT Instance ID.
        /// </summary>
        [Input("sslVpnClientId", required: true)]
        public Input<string> SslVpnClientId { get; set; } = null!;

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public GatewaySslClientCertArgs()
        {
        }
    }

    public sealed class GatewaySslClientCertState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSL-VPN-CLIENT Instance ID.
        /// </summary>
        [Input("sslVpnClientId")]
        public Input<string>? SslVpnClientId { get; set; }

        /// <summary>
        /// `on`: Enable, `off`: Disable.
        /// </summary>
        [Input("switch")]
        public Input<string>? Switch { get; set; }

        public GatewaySslClientCertState()
        {
        }
    }
}
