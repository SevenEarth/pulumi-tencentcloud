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
    public static class GetRenewalPrice
    {
        /// <summary>
        /// Use this data source to query detailed information of mariadb renewal_price
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
        ///     var renewalPrice = Tencentcloud.Mariadb.GetRenewalPrice.Invoke(new()
        ///     {
        ///         InstanceId = "tdsql-9vqvls95",
        ///         Period = 2,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetRenewalPriceResult> InvokeAsync(GetRenewalPriceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRenewalPriceResult>("tencentcloud:Mariadb/getRenewalPrice:getRenewalPrice", args ?? new GetRenewalPriceArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mariadb renewal_price
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
        ///     var renewalPrice = Tencentcloud.Mariadb.GetRenewalPrice.Invoke(new()
        ///     {
        ///         InstanceId = "tdsql-9vqvls95",
        ///         Period = 2,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetRenewalPriceResult> Invoke(GetRenewalPriceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRenewalPriceResult>("tencentcloud:Mariadb/getRenewalPrice:getRenewalPrice", args ?? new GetRenewalPriceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRenewalPriceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public string? AmountUnit { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Renewal duration, default: 1 month.
        /// </summary>
        [Input("period")]
        public int? Period { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRenewalPriceArgs()
        {
        }
        public static new GetRenewalPriceArgs Empty => new GetRenewalPriceArgs();
    }

    public sealed class GetRenewalPriceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Price unit. Valid values: `* pent` (cent), `* microPent` (microcent).
        /// </summary>
        [Input("amountUnit")]
        public Input<string>? AmountUnit { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Renewal duration, default: 1 month.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRenewalPriceInvokeArgs()
        {
        }
        public static new GetRenewalPriceInvokeArgs Empty => new GetRenewalPriceInvokeArgs();
    }


    [OutputType]
    public sealed class GetRenewalPriceResult
    {
        public readonly string? AmountUnit;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Original price * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int OriginalPrice;
        public readonly int? Period;
        /// <summary>
        /// The actual price may be different from the original price due to discounts. * Unit: Cent (default). If the request parameter contains `AmountUnit`, see `AmountUnit` description. * Currency: CNY (Chinese site), USD (international site).
        /// </summary>
        public readonly int Price;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetRenewalPriceResult(
            string? amountUnit,

            string id,

            string instanceId,

            int originalPrice,

            int? period,

            int price,

            string? resultOutputFile)
        {
            AmountUnit = amountUnit;
            Id = id;
            InstanceId = instanceId;
            OriginalPrice = originalPrice;
            Period = period;
            Price = price;
            ResultOutputFile = resultOutputFile;
        }
    }
}
