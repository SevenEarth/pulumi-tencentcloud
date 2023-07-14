// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dc
{
    public static class GetInternetAddressQuota
    {
        /// <summary>
        /// Use this data source to query detailed information of dc internet_address_quota
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
        ///         var internetAddressQuota = Output.Create(Tencentcloud.Dc.GetInternetAddressQuota.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInternetAddressQuotaResult> InvokeAsync(GetInternetAddressQuotaArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInternetAddressQuotaResult>("tencentcloud:Dc/getInternetAddressQuota:getInternetAddressQuota", args ?? new GetInternetAddressQuotaArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dc internet_address_quota
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
        ///         var internetAddressQuota = Output.Create(Tencentcloud.Dc.GetInternetAddressQuota.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInternetAddressQuotaResult> Invoke(GetInternetAddressQuotaInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInternetAddressQuotaResult>("tencentcloud:Dc/getInternetAddressQuota:getInternetAddressQuota", args ?? new GetInternetAddressQuotaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInternetAddressQuotaArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInternetAddressQuotaArgs()
        {
        }
    }

    public sealed class GetInternetAddressQuotaInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInternetAddressQuotaInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInternetAddressQuotaResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Number of used BGP type IPv4 Internet addresses.
        /// </summary>
        public readonly int Ipv4BgpNum;
        /// <summary>
        /// BGP type IPv4 Internet address quota.
        /// </summary>
        public readonly int Ipv4BgpQuota;
        /// <summary>
        /// The number of non-BGP Internet addresses used.
        /// </summary>
        public readonly int Ipv4OtherNum;
        /// <summary>
        /// Non-BGP type IPv4 Internet address quota.
        /// </summary>
        public readonly int Ipv4OtherQuota;
        /// <summary>
        /// The minimum prefix length allowed on the IPv6 Internet public network.
        /// </summary>
        public readonly int Ipv6PrefixLen;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInternetAddressQuotaResult(
            string id,

            int ipv4BgpNum,

            int ipv4BgpQuota,

            int ipv4OtherNum,

            int ipv4OtherQuota,

            int ipv6PrefixLen,

            string? resultOutputFile)
        {
            Id = id;
            Ipv4BgpNum = ipv4BgpNum;
            Ipv4BgpQuota = ipv4BgpQuota;
            Ipv4OtherNum = ipv4OtherNum;
            Ipv4OtherQuota = ipv4OtherQuota;
            Ipv6PrefixLen = ipv6PrefixLen;
            ResultOutputFile = resultOutputFile;
        }
    }
}
