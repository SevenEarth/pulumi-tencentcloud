// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb
{
    public static class GetPlatformProducts
    {
        /// <summary>
        /// Use this data source to query detailed information of eb platform_products
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
        ///     var platformProducts = Tencentcloud.Eb.GetPlatformProducts.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetPlatformProductsResult> InvokeAsync(GetPlatformProductsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlatformProductsResult>("tencentcloud:Eb/getPlatformProducts:getPlatformProducts", args ?? new GetPlatformProductsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of eb platform_products
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
        ///     var platformProducts = Tencentcloud.Eb.GetPlatformProducts.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetPlatformProductsResult> Invoke(GetPlatformProductsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlatformProductsResult>("tencentcloud:Eb/getPlatformProducts:getPlatformProducts", args ?? new GetPlatformProductsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlatformProductsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetPlatformProductsArgs()
        {
        }
        public static new GetPlatformProductsArgs Empty => new GetPlatformProductsArgs();
    }

    public sealed class GetPlatformProductsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetPlatformProductsInvokeArgs()
        {
        }
        public static new GetPlatformProductsInvokeArgs Empty => new GetPlatformProductsInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlatformProductsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Platform product list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPlatformProductsPlatformProductResult> EbPlatformProducts;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetPlatformProductsResult(
            string id,

            ImmutableArray<Outputs.GetPlatformProductsPlatformProductResult> platformProducts,

            string? resultOutputFile)
        {
            Id = id;
            EbPlatformProducts = platformProducts;
            ResultOutputFile = resultOutputFile;
        }
    }
}
