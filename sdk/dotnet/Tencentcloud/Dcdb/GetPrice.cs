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
    public static class GetPrice
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb price
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
        ///         var price = Output.Create(Tencentcloud.Dcdb.GetPrice.InvokeAsync(new Tencentcloud.Dcdb.GetPriceArgs
        ///         {
        ///             InstanceCount = 1,
        ///             Zone = @var.Default_az,
        ///             Period = 1,
        ///             ShardNodeCount = 2,
        ///             ShardMemory = 2,
        ///             ShardStorage = 10,
        ///             ShardCount = 2,
        ///             Paymode = "postpaid",
        ///             AmountUnit = "pent",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPriceResult> InvokeAsync(GetPriceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPriceResult>("tencentcloud:Dcdb/getPrice:getPrice", args ?? new GetPriceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb price
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
        ///         var price = Output.Create(Tencentcloud.Dcdb.GetPrice.InvokeAsync(new Tencentcloud.Dcdb.GetPriceArgs
        ///         {
        ///             InstanceCount = 1,
        ///             Zone = @var.Default_az,
        ///             Period = 1,
        ///             ShardNodeCount = 2,
        ///             ShardMemory = 2,
        ///             ShardStorage = 10,
        ///             ShardCount = 2,
        ///             Paymode = "postpaid",
        ///             AmountUnit = "pent",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPriceResult> Invoke(GetPriceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPriceResult>("tencentcloud:Dcdb/getPrice:getPrice", args ?? new GetPriceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPriceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `pent` (cent), `microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public string? AmountUnit { get; set; }

        /// <summary>
        /// The count of instances wants to buy.
        /// </summary>
        [Input("instanceCount", required: true)]
        public int InstanceCount { get; set; }

        /// <summary>
        /// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
        /// </summary>
        [Input("paymode")]
        public string? Paymode { get; set; }

        /// <summary>
        /// Purchase period in months.
        /// </summary>
        [Input("period", required: true)]
        public int Period { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Number of instance shards.
        /// </summary>
        [Input("shardCount", required: true)]
        public int ShardCount { get; set; }

        /// <summary>
        /// Shard memory size in GB.
        /// </summary>
        [Input("shardMemory", required: true)]
        public int ShardMemory { get; set; }

        /// <summary>
        /// Number of instance shard nodes.
        /// </summary>
        [Input("shardNodeCount", required: true)]
        public int ShardNodeCount { get; set; }

        /// <summary>
        /// Shard storage capacity in GB.
        /// </summary>
        [Input("shardStorage", required: true)]
        public int ShardStorage { get; set; }

        /// <summary>
        /// AZ ID of the purchased instance.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetPriceArgs()
        {
        }
    }

    public sealed class GetPriceInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `pent` (cent), `microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public Input<string>? AmountUnit { get; set; }

        /// <summary>
        /// The count of instances wants to buy.
        /// </summary>
        [Input("instanceCount", required: true)]
        public Input<int> InstanceCount { get; set; } = null!;

        /// <summary>
        /// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
        /// </summary>
        [Input("paymode")]
        public Input<string>? Paymode { get; set; }

        /// <summary>
        /// Purchase period in months.
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Number of instance shards.
        /// </summary>
        [Input("shardCount", required: true)]
        public Input<int> ShardCount { get; set; } = null!;

        /// <summary>
        /// Shard memory size in GB.
        /// </summary>
        [Input("shardMemory", required: true)]
        public Input<int> ShardMemory { get; set; } = null!;

        /// <summary>
        /// Number of instance shard nodes.
        /// </summary>
        [Input("shardNodeCount", required: true)]
        public Input<int> ShardNodeCount { get; set; } = null!;

        /// <summary>
        /// Shard storage capacity in GB.
        /// </summary>
        [Input("shardStorage", required: true)]
        public Input<int> ShardStorage { get; set; } = null!;

        /// <summary>
        /// AZ ID of the purchased instance.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetPriceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPriceResult
    {
        public readonly string? AmountUnit;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int InstanceCount;
        /// <summary>
        /// Original price. Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int OriginalPrice;
        public readonly string? Paymode;
        public readonly int Period;
        /// <summary>
        /// The actual price may be different from the original price due to discounts. Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int DcdbPrice;
        public readonly string? ResultOutputFile;
        public readonly int ShardCount;
        public readonly int ShardMemory;
        public readonly int ShardNodeCount;
        public readonly int ShardStorage;
        public readonly string Zone;

        [OutputConstructor]
        private GetPriceResult(
            string? amountUnit,

            string id,

            int instanceCount,

            int originalPrice,

            string? paymode,

            int period,

            int price,

            string? resultOutputFile,

            int shardCount,

            int shardMemory,

            int shardNodeCount,

            int shardStorage,

            string zone)
        {
            AmountUnit = amountUnit;
            Id = id;
            InstanceCount = instanceCount;
            OriginalPrice = originalPrice;
            Paymode = paymode;
            Period = period;
            DcdbPrice = price;
            ResultOutputFile = resultOutputFile;
            ShardCount = shardCount;
            ShardMemory = shardMemory;
            ShardNodeCount = shardNodeCount;
            ShardStorage = shardStorage;
            Zone = zone;
        }
    }
}
