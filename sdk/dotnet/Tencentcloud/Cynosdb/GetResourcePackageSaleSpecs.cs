// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    public static class GetResourcePackageSaleSpecs
    {
        /// <summary>
        /// Use this data source to query detailed information of cynosdb resource_package_sale_specs
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
        ///     var resourcePackageSaleSpecs = Tencentcloud.Cynosdb.GetResourcePackageSaleSpecs.Invoke(new()
        ///     {
        ///         InstanceType = "cynosdb-serverless",
        ///         PackageRegion = "china",
        ///         PackageType = "CCU",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetResourcePackageSaleSpecsResult> InvokeAsync(GetResourcePackageSaleSpecsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourcePackageSaleSpecsResult>("tencentcloud:Cynosdb/getResourcePackageSaleSpecs:getResourcePackageSaleSpecs", args ?? new GetResourcePackageSaleSpecsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cynosdb resource_package_sale_specs
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
        ///     var resourcePackageSaleSpecs = Tencentcloud.Cynosdb.GetResourcePackageSaleSpecs.Invoke(new()
        ///     {
        ///         InstanceType = "cynosdb-serverless",
        ///         PackageRegion = "china",
        ///         PackageType = "CCU",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetResourcePackageSaleSpecsResult> Invoke(GetResourcePackageSaleSpecsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourcePackageSaleSpecsResult>("tencentcloud:Cynosdb/getResourcePackageSaleSpecs:getResourcePackageSaleSpecs", args ?? new GetResourcePackageSaleSpecsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcePackageSaleSpecsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance Type. Value range: cynosdb-serverless, cynosdb, cdb.
        /// </summary>
        [Input("instanceType", required: true)]
        public string InstanceType { get; set; } = null!;

        /// <summary>
        /// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
        /// </summary>
        [Input("packageRegion", required: true)]
        public string PackageRegion { get; set; } = null!;

        /// <summary>
        /// Resource package type CCU - Computing resource package DISK - Storage resource package.
        /// </summary>
        [Input("packageType", required: true)]
        public string PackageType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetResourcePackageSaleSpecsArgs()
        {
        }
        public static new GetResourcePackageSaleSpecsArgs Empty => new GetResourcePackageSaleSpecsArgs();
    }

    public sealed class GetResourcePackageSaleSpecsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Instance Type. Value range: cynosdb-serverless, cynosdb, cdb.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
        /// </summary>
        [Input("packageRegion", required: true)]
        public Input<string> PackageRegion { get; set; } = null!;

        /// <summary>
        /// Resource package type CCU - Computing resource package DISK - Storage resource package.
        /// </summary>
        [Input("packageType", required: true)]
        public Input<string> PackageType { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetResourcePackageSaleSpecsInvokeArgs()
        {
        }
        public static new GetResourcePackageSaleSpecsInvokeArgs Empty => new GetResourcePackageSaleSpecsInvokeArgs();
    }


    [OutputType]
    public sealed class GetResourcePackageSaleSpecsResult
    {
        /// <summary>
        /// Resource package details note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcePackageSaleSpecsDetailResult> Details;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceType;
        /// <summary>
        /// Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string PackageRegion;
        /// <summary>
        /// Resource package type CCU - Compute resource package DISK - Store resource package Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string PackageType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetResourcePackageSaleSpecsResult(
            ImmutableArray<Outputs.GetResourcePackageSaleSpecsDetailResult> details,

            string id,

            string instanceType,

            string packageRegion,

            string packageType,

            string? resultOutputFile)
        {
            Details = details;
            Id = id;
            InstanceType = instanceType;
            PackageRegion = packageRegion;
            PackageType = packageType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
