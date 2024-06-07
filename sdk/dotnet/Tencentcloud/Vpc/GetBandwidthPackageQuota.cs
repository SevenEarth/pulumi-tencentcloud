// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    public static class GetBandwidthPackageQuota
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc bandwidth_package_quota
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
        ///     var bandwidthPackageQuota = Tencentcloud.Vpc.GetBandwidthPackageQuota.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBandwidthPackageQuotaResult> InvokeAsync(GetBandwidthPackageQuotaArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBandwidthPackageQuotaResult>("tencentcloud:Vpc/getBandwidthPackageQuota:getBandwidthPackageQuota", args ?? new GetBandwidthPackageQuotaArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc bandwidth_package_quota
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
        ///     var bandwidthPackageQuota = Tencentcloud.Vpc.GetBandwidthPackageQuota.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBandwidthPackageQuotaResult> Invoke(GetBandwidthPackageQuotaInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBandwidthPackageQuotaResult>("tencentcloud:Vpc/getBandwidthPackageQuota:getBandwidthPackageQuota", args ?? new GetBandwidthPackageQuotaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBandwidthPackageQuotaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBandwidthPackageQuotaArgs()
        {
        }
        public static new GetBandwidthPackageQuotaArgs Empty => new GetBandwidthPackageQuotaArgs();
    }

    public sealed class GetBandwidthPackageQuotaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBandwidthPackageQuotaInvokeArgs()
        {
        }
        public static new GetBandwidthPackageQuotaInvokeArgs Empty => new GetBandwidthPackageQuotaInvokeArgs();
    }


    [OutputType]
    public sealed class GetBandwidthPackageQuotaResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Bandwidth Package Quota Details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBandwidthPackageQuotaQuotaSetResult> QuotaSets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBandwidthPackageQuotaResult(
            string id,

            ImmutableArray<Outputs.GetBandwidthPackageQuotaQuotaSetResult> quotaSets,

            string? resultOutputFile)
        {
            Id = id;
            QuotaSets = quotaSets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
