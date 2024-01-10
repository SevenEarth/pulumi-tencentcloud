// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod
{
    public static class GetDomainAnalytics
    {
        /// <summary>
        /// Use this data source to query detailed information of dnspod domain_analytics
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
        ///         var domainAnalytics = Output.Create(Tencentcloud.Dnspod.GetDomainAnalytics.InvokeAsync(new Tencentcloud.Dnspod.GetDomainAnalyticsArgs
        ///         {
        ///             DnsFormat = "HOUR",
        ///             Domain = "dnspod.cn",
        ///             EndDate = "2023-10-12",
        ///             StartDate = "2023-10-07",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDomainAnalyticsResult> InvokeAsync(GetDomainAnalyticsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainAnalyticsResult>("tencentcloud:Dnspod/getDomainAnalytics:getDomainAnalytics", args ?? new GetDomainAnalyticsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dnspod domain_analytics
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
        ///         var domainAnalytics = Output.Create(Tencentcloud.Dnspod.GetDomainAnalytics.InvokeAsync(new Tencentcloud.Dnspod.GetDomainAnalyticsArgs
        ///         {
        ///             DnsFormat = "HOUR",
        ///             Domain = "dnspod.cn",
        ///             EndDate = "2023-10-12",
        ///             StartDate = "2023-10-07",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDomainAnalyticsResult> Invoke(GetDomainAnalyticsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainAnalyticsResult>("tencentcloud:Dnspod/getDomainAnalytics:getDomainAnalytics", args ?? new GetDomainAnalyticsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainAnalyticsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
        /// </summary>
        [Input("dnsFormat")]
        public string? DnsFormat { get; set; }

        /// <summary>
        /// The domain name to query for resolution volume.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
        /// </summary>
        [Input("domainId")]
        public int? DomainId { get; set; }

        /// <summary>
        /// The end date of the query, format: YYYY-MM-DD.
        /// </summary>
        [Input("endDate", required: true)]
        public string EndDate { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// The start date of the query, format: YYYY-MM-DD.
        /// </summary>
        [Input("startDate", required: true)]
        public string StartDate { get; set; } = null!;

        public GetDomainAnalyticsArgs()
        {
        }
    }

    public sealed class GetDomainAnalyticsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
        /// </summary>
        [Input("dnsFormat")]
        public Input<string>? DnsFormat { get; set; }

        /// <summary>
        /// The domain name to query for resolution volume.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
        /// </summary>
        [Input("domainId")]
        public Input<int>? DomainId { get; set; }

        /// <summary>
        /// The end date of the query, format: YYYY-MM-DD.
        /// </summary>
        [Input("endDate", required: true)]
        public Input<string> EndDate { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// The start date of the query, format: YYYY-MM-DD.
        /// </summary>
        [Input("startDate", required: true)]
        public Input<string> StartDate { get; set; } = null!;

        public GetDomainAnalyticsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainAnalyticsResult
    {
        /// <summary>
        /// Domain alias resolution volume statistics information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainAnalyticsAliasDataResult> AliasDatas;
        /// <summary>
        /// Subtotal of resolution volume for the current statistical dimension.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainAnalyticsDataResult> Datas;
        /// <summary>
        /// DATE: Statistics by day dimension HOUR: Statistics by hour dimension.
        /// </summary>
        public readonly string? DnsFormat;
        /// <summary>
        /// The domain name currently being queried.
        /// </summary>
        public readonly string Domain;
        public readonly int? DomainId;
        /// <summary>
        /// End time of the current statistical period.
        /// </summary>
        public readonly string EndDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Domain resolution volume statistics query information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainAnalyticsInfoResult> Infos;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Start time of the current statistical period.
        /// </summary>
        public readonly string StartDate;

        [OutputConstructor]
        private GetDomainAnalyticsResult(
            ImmutableArray<Outputs.GetDomainAnalyticsAliasDataResult> aliasDatas,

            ImmutableArray<Outputs.GetDomainAnalyticsDataResult> datas,

            string? dnsFormat,

            string domain,

            int? domainId,

            string endDate,

            string id,

            ImmutableArray<Outputs.GetDomainAnalyticsInfoResult> infos,

            string? resultOutputFile,

            string startDate)
        {
            AliasDatas = aliasDatas;
            Datas = datas;
            DnsFormat = dnsFormat;
            Domain = domain;
            DomainId = domainId;
            EndDate = endDate;
            Id = id;
            Infos = infos;
            ResultOutputFile = resultOutputFile;
            StartDate = startDate;
        }
    }
}
