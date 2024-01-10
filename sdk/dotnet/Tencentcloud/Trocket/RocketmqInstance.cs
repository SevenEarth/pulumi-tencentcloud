// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Trocket
{
    /// <summary>
    /// Provides a resource to create a rocketmq 5.x instance
    /// 
    /// &gt; **NOTE:** It only support create postpaid rocketmq 5.x instance.
    /// 
    /// ## Example Usage
    /// ### Basic Instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rocketmqInstance = new Tencentcloud.Trocket.RocketmqInstance("rocketmqInstance", new Tencentcloud.Trocket.RocketmqInstanceArgs
    ///         {
    ///             InstanceType = "EXPERIMENT",
    ///             Remark = "remark",
    ///             SkuCode = "experiment_500",
    ///             SubnetId = "subnet-xxxxxx",
    ///             Tags = 
    ///             {
    ///                 { "tag_key", "rocketmq" },
    ///                 { "tag_value", "5.x" },
    ///             },
    ///             VpcId = "vpc-xxxxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Enable Public Instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rocketmqInstancePublic = new Tencentcloud.Trocket.RocketmqInstance("rocketmqInstancePublic", new Tencentcloud.Trocket.RocketmqInstanceArgs
    ///         {
    ///             Bandwidth = 1,
    ///             EnablePublic = true,
    ///             InstanceType = "EXPERIMENT",
    ///             Remark = "remark",
    ///             SkuCode = "experiment_500",
    ///             SubnetId = "subnet-xxxxxx",
    ///             Tags = 
    ///             {
    ///                 { "tag_key", "rocketmq" },
    ///                 { "tag_value", "5.x" },
    ///             },
    ///             VpcId = "vpc-xxxxxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// trocket rocketmq_instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Trocket/rocketmqInstance:RocketmqInstance rocketmq_instance rocketmq_instance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Trocket/rocketmqInstance:RocketmqInstance")]
    public partial class RocketmqInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        /// </summary>
        [Output("enablePublic")]
        public Output<bool> EnablePublic { get; private set; } = null!;

        /// <summary>
        /// Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// Public network access whitelist.
        /// </summary>
        [Output("ipRules")]
        public Output<ImmutableArray<Outputs.RocketmqInstanceIpRule>> IpRules { get; private set; } = null!;

        /// <summary>
        /// Message retention time in hours.
        /// </summary>
        [Output("messageRetention")]
        public Output<int> MessageRetention { get; private set; } = null!;

        /// <summary>
        /// Instance name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Public network access address.
        /// </summary>
        [Output("publicEndPoint")]
        public Output<string> PublicEndPoint { get; private set; } = null!;

        /// <summary>
        /// Remark.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_4k, basic_6k.
        /// </summary>
        [Output("skuCode")]
        public Output<string> SkuCode { get; private set; } = null!;

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>> Tags { get; private set; } = null!;

        /// <summary>
        /// VPC access address.
        /// </summary>
        [Output("vpcEndPoint")]
        public Output<string> VpcEndPoint { get; private set; } = null!;

        /// <summary>
        /// VPC id.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a RocketmqInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RocketmqInstance(string name, RocketmqInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Trocket/rocketmqInstance:RocketmqInstance", name, args ?? new RocketmqInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RocketmqInstance(string name, Input<string> id, RocketmqInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Trocket/rocketmqInstance:RocketmqInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RocketmqInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RocketmqInstance Get(string name, Input<string> id, RocketmqInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new RocketmqInstance(name, id, state, options);
        }
    }

    public sealed class RocketmqInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        /// </summary>
        [Input("enablePublic")]
        public Input<bool>? EnablePublic { get; set; }

        /// <summary>
        /// Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("ipRules")]
        private InputList<Inputs.RocketmqInstanceIpRuleArgs>? _ipRules;

        /// <summary>
        /// Public network access whitelist.
        /// </summary>
        public InputList<Inputs.RocketmqInstanceIpRuleArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.RocketmqInstanceIpRuleArgs>());
            set => _ipRules = value;
        }

        /// <summary>
        /// Message retention time in hours.
        /// </summary>
        [Input("messageRetention")]
        public Input<int>? MessageRetention { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Remark.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_4k, basic_6k.
        /// </summary>
        [Input("skuCode", required: true)]
        public Input<string> SkuCode { get; set; } = null!;

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

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
        /// VPC id.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public RocketmqInstanceArgs()
        {
        }
    }

    public sealed class RocketmqInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Public network bandwidth. `bandwidth` must be greater than zero when `enable_public` equal true.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Whether to enable the public network. Must set `bandwidth` when `enable_public` equal true.
        /// </summary>
        [Input("enablePublic")]
        public Input<bool>? EnablePublic { get; set; }

        /// <summary>
        /// Instance type. Valid values: `EXPERIMENT`, `BASIC`, `PRO`, `PLATINUM`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("ipRules")]
        private InputList<Inputs.RocketmqInstanceIpRuleGetArgs>? _ipRules;

        /// <summary>
        /// Public network access whitelist.
        /// </summary>
        public InputList<Inputs.RocketmqInstanceIpRuleGetArgs> IpRules
        {
            get => _ipRules ?? (_ipRules = new InputList<Inputs.RocketmqInstanceIpRuleGetArgs>());
            set => _ipRules = value;
        }

        /// <summary>
        /// Message retention time in hours.
        /// </summary>
        [Input("messageRetention")]
        public Input<int>? MessageRetention { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Public network access address.
        /// </summary>
        [Input("publicEndPoint")]
        public Input<string>? PublicEndPoint { get; set; }

        /// <summary>
        /// Remark.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        /// <summary>
        /// SKU code. Available specifications are as follows: experiment_500, basic_1k, basic_2k, basic_4k, basic_6k.
        /// </summary>
        [Input("skuCode")]
        public Input<string>? SkuCode { get; set; }

        /// <summary>
        /// Subnet id.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

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
        /// VPC access address.
        /// </summary>
        [Input("vpcEndPoint")]
        public Input<string>? VpcEndPoint { get; set; }

        /// <summary>
        /// VPC id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public RocketmqInstanceState()
        {
        }
    }
}
