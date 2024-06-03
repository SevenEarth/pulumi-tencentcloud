// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    public static class GetPrice
    {
        /// <summary>
        /// Use this data source to query detailed information of mariadb price
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
        ///     var price = Tencentcloud.Mariadb.GetPrice.Invoke(new()
        ///     {
        ///         BuyCount = 1,
        ///         Memory = 2,
        ///         NodeCount = 2,
        ///         Paymode = "prepaid",
        ///         Period = 1,
        ///         Storage = 20,
        ///         Zone = "ap-guangzhou-3",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPriceResult> InvokeAsync(GetPriceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPriceResult>("tencentcloud:Mariadb/getPrice:getPrice", args ?? new GetPriceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mariadb price
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
        ///     var price = Tencentcloud.Mariadb.GetPrice.Invoke(new()
        ///     {
        ///         BuyCount = 1,
        ///         Memory = 2,
        ///         NodeCount = 2,
        ///         Paymode = "prepaid",
        ///         Period = 1,
        ///         Storage = 20,
        ///         Zone = "ap-guangzhou-3",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPriceResult> Invoke(GetPriceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPriceResult>("tencentcloud:Mariadb/getPrice:getPrice", args ?? new GetPriceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPriceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public string? AmountUnit { get; set; }

        /// <summary>
        /// The quantity you want to purchase is queried by default for the price of purchasing 1 instance.
        /// </summary>
        [Input("buyCount", required: true)]
        public int BuyCount { get; set; }

        /// <summary>
        /// Memory size in GB, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("memory", required: true)]
        public int Memory { get; set; }

        /// <summary>
        /// Number of instance nodes, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("nodeCount", required: true)]
        public int NodeCount { get; set; }

        /// <summary>
        /// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
        /// </summary>
        [Input("paymode")]
        public string? Paymode { get; set; }

        /// <summary>
        /// Purchase period in months.
        /// </summary>
        [Input("period")]
        public int? Period { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Storage capacity in GB. The maximum and minimum storage space can be obtained by querying instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("storage", required: true)]
        public int Storage { get; set; }

        /// <summary>
        /// AZ ID of the purchased instance.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetPriceArgs()
        {
        }
        public static new GetPriceArgs Empty => new GetPriceArgs();
    }

    public sealed class GetPriceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public Input<string>? AmountUnit { get; set; }

        /// <summary>
        /// The quantity you want to purchase is queried by default for the price of purchasing 1 instance.
        /// </summary>
        [Input("buyCount", required: true)]
        public Input<int> BuyCount { get; set; } = null!;

        /// <summary>
        /// Memory size in GB, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Number of instance nodes, which can be obtained by querying the instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("nodeCount", required: true)]
        public Input<int> NodeCount { get; set; } = null!;

        /// <summary>
        /// Billing type. Valid values: `postpaid` (pay-as-you-go), `prepaid` (monthly subscription).
        /// </summary>
        [Input("paymode")]
        public Input<string>? Paymode { get; set; }

        /// <summary>
        /// Purchase period in months.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Storage capacity in GB. The maximum and minimum storage space can be obtained by querying instance specification through the `DescribeDBInstanceSpecs` API.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// AZ ID of the purchased instance.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetPriceInvokeArgs()
        {
        }
        public static new GetPriceInvokeArgs Empty => new GetPriceInvokeArgs();
    }


    [OutputType]
    public sealed class GetPriceResult
    {
        public readonly string? AmountUnit;
        public readonly int BuyCount;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int Memory;
        public readonly int NodeCount;
        /// <summary>
        /// Original price * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int OriginalPrice;
        public readonly string? Paymode;
        public readonly int? Period;
        /// <summary>
        /// The actual price may be different from the original price due to discounts. * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int MariadbPrice;
        public readonly string? ResultOutputFile;
        public readonly int Storage;
        public readonly string Zone;

        [OutputConstructor]
        private GetPriceResult(
            string? amountUnit,

            int buyCount,

            string id,

            int memory,

            int nodeCount,

            int originalPrice,

            string? paymode,

            int? period,

            int price,

            string? resultOutputFile,

            int storage,

            string zone)
        {
            AmountUnit = amountUnit;
            BuyCount = buyCount;
            Id = id;
            Memory = memory;
            NodeCount = nodeCount;
            OriginalPrice = originalPrice;
            Paymode = paymode;
            Period = period;
            MariadbPrice = price;
            ResultOutputFile = resultOutputFile;
            Storage = storage;
            Zone = zone;
        }
    }
}
