// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cfw
{
    /// <summary>
    /// Provides a resource to create a cfw edge_firewall_switch
    /// 
    /// ## Example Usage
    /// ### If not set subnet_id
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEdgeFwSwitches = Output.Create(Tencentcloud.Cfw.GetEdgeFwSwitches.InvokeAsync());
    ///         var exampleEdgeFirewallSwitch = new Tencentcloud.Cfw.EdgeFirewallSwitch("exampleEdgeFirewallSwitch", new Tencentcloud.Cfw.EdgeFirewallSwitchArgs
    ///         {
    ///             PublicIp = exampleEdgeFwSwitches.Apply(exampleEdgeFwSwitches =&gt; exampleEdgeFwSwitches.Datas?[0]?.PublicIp),
    ///             SwitchMode = 1,
    ///             Enable = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### If set subnet id
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleEdgeFwSwitches = Output.Create(Tencentcloud.Cfw.GetEdgeFwSwitches.InvokeAsync());
    ///         var exampleEdgeFirewallSwitch = new Tencentcloud.Cfw.EdgeFirewallSwitch("exampleEdgeFirewallSwitch", new Tencentcloud.Cfw.EdgeFirewallSwitchArgs
    ///         {
    ///             PublicIp = exampleEdgeFwSwitches.Apply(exampleEdgeFwSwitches =&gt; exampleEdgeFwSwitches.Datas?[0]?.PublicIp),
    ///             SubnetId = "subnet-id",
    ///             SwitchMode = 1,
    ///             Enable = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cfw/edgeFirewallSwitch:EdgeFirewallSwitch")]
    public partial class EdgeFirewallSwitch : Pulumi.CustomResource
    {
        /// <summary>
        /// Switch, 0: off, 1: on.
        /// </summary>
        [Output("enable")]
        public Output<int> Enable { get; private set; } = null!;

        /// <summary>
        /// Public Ip.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// The first EIP switch in the vpc is turned on, and you need to specify a subnet to create a private connection. If `switch_mode` is 1 and `enable` is 1, this field is required.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// 0: bypass; 1: serial.
        /// </summary>
        [Output("switchMode")]
        public Output<int> SwitchMode { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeFirewallSwitch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeFirewallSwitch(string name, EdgeFirewallSwitchArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfw/edgeFirewallSwitch:EdgeFirewallSwitch", name, args ?? new EdgeFirewallSwitchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeFirewallSwitch(string name, Input<string> id, EdgeFirewallSwitchState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfw/edgeFirewallSwitch:EdgeFirewallSwitch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeFirewallSwitch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeFirewallSwitch Get(string name, Input<string> id, EdgeFirewallSwitchState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeFirewallSwitch(name, id, state, options);
        }
    }

    public sealed class EdgeFirewallSwitchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Switch, 0: off, 1: on.
        /// </summary>
        [Input("enable", required: true)]
        public Input<int> Enable { get; set; } = null!;

        /// <summary>
        /// Public Ip.
        /// </summary>
        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        /// <summary>
        /// The first EIP switch in the vpc is turned on, and you need to specify a subnet to create a private connection. If `switch_mode` is 1 and `enable` is 1, this field is required.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// 0: bypass; 1: serial.
        /// </summary>
        [Input("switchMode", required: true)]
        public Input<int> SwitchMode { get; set; } = null!;

        public EdgeFirewallSwitchArgs()
        {
        }
    }

    public sealed class EdgeFirewallSwitchState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Switch, 0: off, 1: on.
        /// </summary>
        [Input("enable")]
        public Input<int>? Enable { get; set; }

        /// <summary>
        /// Public Ip.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// The first EIP switch in the vpc is turned on, and you need to specify a subnet to create a private connection. If `switch_mode` is 1 and `enable` is 1, this field is required.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// 0: bypass; 1: serial.
        /// </summary>
        [Input("switchMode")]
        public Input<int>? SwitchMode { get; set; }

        public EdgeFirewallSwitchState()
        {
        }
    }
}
