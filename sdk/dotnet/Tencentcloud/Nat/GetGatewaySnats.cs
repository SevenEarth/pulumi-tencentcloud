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
    public static class GetGatewaySnats
    {
        /// <summary>
        /// Use this data source to query detailed information of VPN gateways.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var snat = Tencentcloud.Nat.GetGatewaySnats.Invoke(new()
        ///     {
        ///         NatGatewayId = tencentcloud_nat_gateway.My_nat.Id,
        ///         SubnetId = tencentcloud_nat_gateway_snat.My_subnet.Id,
        ///         PublicIpAddrs = new[]
        ///         {
        ///             "50.29.23.234",
        ///         },
        ///         Description = "snat demo",
        ///         ResultOutputFile = "./snat.txt",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetGatewaySnatsResult> InvokeAsync(GetGatewaySnatsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewaySnatsResult>("tencentcloud:Nat/getGatewaySnats:getGatewaySnats", args ?? new GetGatewaySnatsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of VPN gateways.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var snat = Tencentcloud.Nat.GetGatewaySnats.Invoke(new()
        ///     {
        ///         NatGatewayId = tencentcloud_nat_gateway.My_nat.Id,
        ///         SubnetId = tencentcloud_nat_gateway_snat.My_subnet.Id,
        ///         PublicIpAddrs = new[]
        ///         {
        ///             "50.29.23.234",
        ///         },
        ///         Description = "snat demo",
        ///         ResultOutputFile = "./snat.txt",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetGatewaySnatsResult> Invoke(GetGatewaySnatsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewaySnatsResult>("tencentcloud:Nat/getGatewaySnats:getGatewaySnats", args ?? new GetGatewaySnatsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewaySnatsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// NAT gateway ID.
        /// </summary>
        [Input("natGatewayId", required: true)]
        public string NatGatewayId { get; set; } = null!;

        [Input("publicIpAddrs")]
        private List<string>? _publicIpAddrs;

        /// <summary>
        /// Elastic IP address pool.
        /// </summary>
        public List<string> PublicIpAddrs
        {
            get => _publicIpAddrs ?? (_publicIpAddrs = new List<string>());
            set => _publicIpAddrs = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Subnet instance ID.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        public GetGatewaySnatsArgs()
        {
        }
        public static new GetGatewaySnatsArgs Empty => new GetGatewaySnatsArgs();
    }

    public sealed class GetGatewaySnatsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// NAT gateway ID.
        /// </summary>
        [Input("natGatewayId", required: true)]
        public Input<string> NatGatewayId { get; set; } = null!;

        [Input("publicIpAddrs")]
        private InputList<string>? _publicIpAddrs;

        /// <summary>
        /// Elastic IP address pool.
        /// </summary>
        public InputList<string> PublicIpAddrs
        {
            get => _publicIpAddrs ?? (_publicIpAddrs = new InputList<string>());
            set => _publicIpAddrs = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Subnet instance ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public GetGatewaySnatsInvokeArgs()
        {
        }
        public static new GetGatewaySnatsInvokeArgs Empty => new GetGatewaySnatsInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewaySnatsResult
    {
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string NatGatewayId;
        public readonly ImmutableArray<string> PublicIpAddrs;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Information list of the nat gateway snat.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewaySnatsSnatListResult> SnatLists;
        public readonly string? SubnetId;

        [OutputConstructor]
        private GetGatewaySnatsResult(
            string? description,

            string id,

            string? instanceId,

            string natGatewayId,

            ImmutableArray<string> publicIpAddrs,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetGatewaySnatsSnatListResult> snatLists,

            string? subnetId)
        {
            Description = description;
            Id = id;
            InstanceId = instanceId;
            NatGatewayId = natGatewayId;
            PublicIpAddrs = publicIpAddrs;
            ResultOutputFile = resultOutputFile;
            SnatLists = snatLists;
            SubnetId = subnetId;
        }
    }
}
