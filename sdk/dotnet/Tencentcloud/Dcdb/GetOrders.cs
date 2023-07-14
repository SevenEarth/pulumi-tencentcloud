// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    public static class GetOrders
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb orders
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
        ///         var orders = Output.Create(Tencentcloud.Dcdb.GetOrders.InvokeAsync(new Tencentcloud.Dcdb.GetOrdersArgs
        ///         {
        ///             DealNames = 
        ///             {
        ///                 "2023061224903413767xxxx",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrdersResult> InvokeAsync(GetOrdersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrdersResult>("tencentcloud:Dcdb/getOrders:getOrders", args ?? new GetOrdersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb orders
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
        ///         var orders = Output.Create(Tencentcloud.Dcdb.GetOrders.InvokeAsync(new Tencentcloud.Dcdb.GetOrdersArgs
        ///         {
        ///             DealNames = 
        ///             {
        ///                 "2023061224903413767xxxx",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrdersResult> Invoke(GetOrdersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrdersResult>("tencentcloud:Dcdb/getOrders:getOrders", args ?? new GetOrdersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrdersArgs : Pulumi.InvokeArgs
    {
        [Input("dealNames", required: true)]
        private List<string>? _dealNames;

        /// <summary>
        /// List of long order numbers to be queried, which are returned for the APIs for creating, renewing, or scaling instances.
        /// </summary>
        public List<string> DealNames
        {
            get => _dealNames ?? (_dealNames = new List<string>());
            set => _dealNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetOrdersArgs()
        {
        }
    }

    public sealed class GetOrdersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dealNames", required: true)]
        private InputList<string>? _dealNames;

        /// <summary>
        /// List of long order numbers to be queried, which are returned for the APIs for creating, renewing, or scaling instances.
        /// </summary>
        public InputList<string> DealNames
        {
            get => _dealNames ?? (_dealNames = new InputList<string>());
            set => _dealNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetOrdersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetOrdersResult
    {
        public readonly ImmutableArray<string> DealNames;
        /// <summary>
        /// Order information list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrdersDealResult> Deals;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetOrdersResult(
            ImmutableArray<string> dealNames,

            ImmutableArray<Outputs.GetOrdersDealResult> deals,

            string id,

            string? resultOutputFile)
        {
            DealNames = dealNames;
            Deals = deals;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
