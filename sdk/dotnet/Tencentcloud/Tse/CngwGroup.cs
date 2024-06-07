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
    /// Provides a resource to create a tse cngw_group
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
    ///     var config = new Config();
    ///     var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-4";
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         AvailabilityZone = availabilityZone,
    ///         CidrBlock = "10.0.1.0/24",
    ///     });
    /// 
    ///     var cngwGateway = new Tencentcloud.Tse.CngwGateway("cngwGateway", new()
    ///     {
    ///         Description = "terraform test1",
    ///         EnableCls = true,
    ///         EngineRegion = "ap-guangzhou",
    ///         FeatureVersion = "STANDARD",
    ///         GatewayVersion = "2.5.1",
    ///         IngressClassName = "tse-nginx-ingress",
    ///         InternetMaxBandwidthOut = 0,
    ///         TradeType = 0,
    ///         Type = "kong",
    ///         NodeConfig = new Tencentcloud.Tse.Inputs.CngwGatewayNodeConfigArgs
    ///         {
    ///             Number = 2,
    ///             Specification = "1c2g",
    ///         },
    ///         VpcConfig = new Tencentcloud.Tse.Inputs.CngwGatewayVpcConfigArgs
    ///         {
    ///             SubnetId = subnet.Id,
    ///             VpcId = vpc.Id,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    ///     var cngwGroup = new Tencentcloud.Tse.CngwGroup("cngwGroup", new()
    ///     {
    ///         Description = "terraform desc",
    ///         GatewayId = cngwGateway.Id,
    ///         SubnetId = subnet.Id,
    ///         NodeConfig = new Tencentcloud.Tse.Inputs.CngwGroupNodeConfigArgs
    ///         {
    ///             Number = 2,
    ///             Specification = "1c2g",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tse/cngwGroup:CngwGroup")]
    public partial class CngwGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// description information of group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// gateway IDonly postpaid gateway supported.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// gateway group id.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// internet configration.
        /// </summary>
        [Output("internetConfig")]
        public Output<Outputs.CngwGroupInternetConfig?> InternetConfig { get; private set; } = null!;

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Output("internetMaxBandwidthOut")]
        public Output<int?> InternetMaxBandwidthOut { get; private set; } = null!;

        /// <summary>
        /// gateway group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// group node configration.
        /// </summary>
        [Output("nodeConfig")]
        public Output<Outputs.CngwGroupNodeConfig> NodeConfig { get; private set; } = null!;

        /// <summary>
        /// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a CngwGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CngwGroup(string name, CngwGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwGroup:CngwGroup", name, args ?? new CngwGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CngwGroup(string name, Input<string> id, CngwGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwGroup:CngwGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CngwGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CngwGroup Get(string name, Input<string> id, CngwGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new CngwGroup(name, id, state, options);
        }
    }

    public sealed class CngwGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description information of group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// gateway IDonly postpaid gateway supported.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// internet configration.
        /// </summary>
        [Input("internetConfig")]
        public Input<Inputs.CngwGroupInternetConfigArgs>? InternetConfig { get; set; }

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// gateway group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// group node configration.
        /// </summary>
        [Input("nodeConfig", required: true)]
        public Input<Inputs.CngwGroupNodeConfigArgs> NodeConfig { get; set; } = null!;

        /// <summary>
        /// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public CngwGroupArgs()
        {
        }
        public static new CngwGroupArgs Empty => new CngwGroupArgs();
    }

    public sealed class CngwGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description information of group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// gateway IDonly postpaid gateway supported.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// gateway group id.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// internet configration.
        /// </summary>
        [Input("internetConfig")]
        public Input<Inputs.CngwGroupInternetConfigGetArgs>? InternetConfig { get; set; }

        /// <summary>
        /// public network outbound traffic bandwidth,[1,2048]Mbps.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        /// <summary>
        /// gateway group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// group node configration.
        /// </summary>
        [Input("nodeConfig")]
        public Input<Inputs.CngwGroupNodeConfigGetArgs>? NodeConfig { get; set; }

        /// <summary>
        /// subnet ID. Assign an IP address to the engine in the VPC subnet. Reference value:- subnet-ahde9me9.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public CngwGroupState()
        {
        }
        public static new CngwGroupState Empty => new CngwGroupState();
    }
}
