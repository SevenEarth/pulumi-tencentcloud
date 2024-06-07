// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Nat
{
    /// <summary>
    /// Provides a resource to create a NAT gateway.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a traditional NAT gateway.
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
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var eipExample1 = new Tencentcloud.Eip.Instance("eipExample1");
    /// 
    ///     var eipExample2 = new Tencentcloud.Eip.Instance("eipExample2");
    /// 
    ///     var example = new Tencentcloud.Nat.Gateway("example", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         Bandwidth = 100,
    ///         MaxConcurrent = 1000000,
    ///         AssignedEipSets = new[]
    ///         {
    ///             eipExample1.PublicIp,
    ///             eipExample2.PublicIp,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "tf_tag_key", "tf_tag_value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create a standard NAT gateway.
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
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var eipExample1 = new Tencentcloud.Eip.Instance("eipExample1");
    /// 
    ///     var eipExample2 = new Tencentcloud.Eip.Instance("eipExample2");
    /// 
    ///     var example = new Tencentcloud.Nat.Gateway("example", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         AssignedEipSets = new[]
    ///         {
    ///             eipExample1.PublicIp,
    ///             eipExample2.PublicIp,
    ///         },
    ///         NatProductVersion = 2,
    ///         Tags = 
    ///         {
    ///             { "tf_tag_key", "tf_tag_value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// NAT gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Nat/gateway:Gateway foo nat-1asg3t63
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Nat/gateway:Gateway")]
    public partial class Gateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
        /// </summary>
        [Output("assignedEipSets")]
        public Output<ImmutableArray<string>> AssignedEipSets { get; private set; } = null!;

        /// <summary>
        /// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
        /// </summary>
        [Output("bandwidth")]
        public Output<int?> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Create time of the NAT gateway.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
        /// </summary>
        [Output("maxConcurrent")]
        public Output<int?> MaxConcurrent { get; private set; } = null!;

        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// 1: traditional NAT, 2: standard NAT, default value is 1.
        /// </summary>
        [Output("natProductVersion")]
        public Output<int> NatProductVersion { get; private set; } = null!;

        /// <summary>
        /// Subnet of NAT.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The available tags within this NAT gateway.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// ID of the vpc.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The availability zone, such as `ap-guangzhou-3`.
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
            : base("tencentcloud:Nat/gateway:Gateway", name, args ?? new GatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Gateway(string name, Input<string> id, GatewayState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Nat/gateway:Gateway", name, state, MakeResourceOptions(options, id))
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

    public sealed class GatewayArgs : global::Pulumi.ResourceArgs
    {
        [Input("assignedEipSets", required: true)]
        private InputList<string>? _assignedEipSets;

        /// <summary>
        /// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
        /// </summary>
        public InputList<string> AssignedEipSets
        {
            get => _assignedEipSets ?? (_assignedEipSets = new InputList<string>());
            set => _assignedEipSets = value;
        }

        /// <summary>
        /// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
        /// </summary>
        [Input("maxConcurrent")]
        public Input<int>? MaxConcurrent { get; set; }

        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// 1: traditional NAT, 2: standard NAT, default value is 1.
        /// </summary>
        [Input("natProductVersion")]
        public Input<int>? NatProductVersion { get; set; }

        /// <summary>
        /// Subnet of NAT.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The available tags within this NAT gateway.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the vpc.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        /// <summary>
        /// The availability zone, such as `ap-guangzhou-3`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GatewayArgs()
        {
        }
        public static new GatewayArgs Empty => new GatewayArgs();
    }

    public sealed class GatewayState : global::Pulumi.ResourceArgs
    {
        [Input("assignedEipSets")]
        private InputList<string>? _assignedEipSets;

        /// <summary>
        /// EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
        /// </summary>
        public InputList<string> AssignedEipSets
        {
            get => _assignedEipSets ?? (_assignedEipSets = new InputList<string>());
            set => _assignedEipSets = value;
        }

        /// <summary>
        /// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Create time of the NAT gateway.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
        /// </summary>
        [Input("maxConcurrent")]
        public Input<int>? MaxConcurrent { get; set; }

        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// 1: traditional NAT, 2: standard NAT, default value is 1.
        /// </summary>
        [Input("natProductVersion")]
        public Input<int>? NatProductVersion { get; set; }

        /// <summary>
        /// Subnet of NAT.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The available tags within this NAT gateway.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of the vpc.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The availability zone, such as `ap-guangzhou-3`.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GatewayState()
        {
        }
        public static new GatewayState Empty => new GatewayState();
    }
}
