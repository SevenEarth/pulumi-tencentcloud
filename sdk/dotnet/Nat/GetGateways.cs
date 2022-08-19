// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Nat
{
    public static class GetGateways
    {
        /// <summary>
        /// Use this data source to query detailed information of NAT gateways.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Tencentcloud.Nat.GetGateways.InvokeAsync(new Tencentcloud.Nat.GetGatewaysArgs
        ///         {
        ///             Id = "nat-xfaq1",
        ///             Name = "main",
        ///             VpcId = "vpc-xfqag",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGatewaysResult> InvokeAsync(GetGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGatewaysResult>("tencentcloud:Nat/getGateways:getGateways", args ?? new GetGatewaysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of NAT gateways.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foo = Output.Create(Tencentcloud.Nat.GetGateways.InvokeAsync(new Tencentcloud.Nat.GetGatewaysArgs
        ///         {
        ///             Id = "nat-xfaq1",
        ///             Name = "main",
        ///             VpcId = "vpc-xfqag",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGatewaysResult> Invoke(GetGatewaysInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGatewaysResult>("tencentcloud:Nat/getGateways:getGateways", args ?? new GetGatewaysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewaysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the NAT gateway.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetGatewaysArgs()
        {
        }
    }

    public sealed class GetGatewaysInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the NAT gateway.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// ID of the VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GetGatewaysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGatewaysResult
    {
        /// <summary>
        /// ID of the NAT gateway.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the NAT gateway.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Information list of the dedicated NATs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGatewaysNatResult> Nats;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// ID of the VPC.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private GetGatewaysResult(
            string? id,

            string? name,

            ImmutableArray<Outputs.GetGatewaysNatResult> nats,

            string? resultOutputFile,

            string? vpcId)
        {
            Id = id;
            Name = name;
            Nats = nats;
            ResultOutputFile = resultOutputFile;
            VpcId = vpcId;
        }
    }
}
