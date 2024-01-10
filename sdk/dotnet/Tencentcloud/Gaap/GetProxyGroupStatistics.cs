// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Gaap
{
    public static class GetProxyGroupStatistics
    {
        /// <summary>
        /// Use this data source to query detailed information of gaap proxy group statistics
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
        ///         var proxyGroupStatistics = Output.Create(Tencentcloud.Gaap.GetProxyGroupStatistics.InvokeAsync(new Tencentcloud.Gaap.GetProxyGroupStatisticsArgs
        ///         {
        ///             EndTime = "2023-10-09 23:59:59",
        ///             Granularity = 300,
        ///             GroupId = "link-8lpyo88p",
        ///             MetricNames = 
        ///             {
        ///                 "InBandwidth",
        ///                 "OutBandwidth",
        ///                 "InFlow",
        ///                 "OutFlow",
        ///             },
        ///             StartTime = "2023-10-09 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProxyGroupStatisticsResult> InvokeAsync(GetProxyGroupStatisticsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProxyGroupStatisticsResult>("tencentcloud:Gaap/getProxyGroupStatistics:getProxyGroupStatistics", args ?? new GetProxyGroupStatisticsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of gaap proxy group statistics
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
        ///         var proxyGroupStatistics = Output.Create(Tencentcloud.Gaap.GetProxyGroupStatistics.InvokeAsync(new Tencentcloud.Gaap.GetProxyGroupStatisticsArgs
        ///         {
        ///             EndTime = "2023-10-09 23:59:59",
        ///             Granularity = 300,
        ///             GroupId = "link-8lpyo88p",
        ///             MetricNames = 
        ///             {
        ///                 "InBandwidth",
        ///                 "OutBandwidth",
        ///                 "InFlow",
        ///                 "OutFlow",
        ///             },
        ///             StartTime = "2023-10-09 00:00:00",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProxyGroupStatisticsResult> Invoke(GetProxyGroupStatisticsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProxyGroupStatisticsResult>("tencentcloud:Gaap/getProxyGroupStatistics:getProxyGroupStatistics", args ?? new GetProxyGroupStatisticsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProxyGroupStatisticsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// End Time.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 1 day, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 3600 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 86400 seconds.
        /// </summary>
        [Input("granularity", required: true)]
        public int Granularity { get; set; }

        /// <summary>
        /// Group Id.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        [Input("metricNames", required: true)]
        private List<string>? _metricNames;

        /// <summary>
        /// Metric Names. support, InBandwidth, OutBandwidth, Concurrent, InPackets, OutPackets.
        /// </summary>
        public List<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new List<string>());
            set => _metricNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start Time.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetProxyGroupStatisticsArgs()
        {
        }
    }

    public sealed class GetProxyGroupStatisticsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// End Time.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Monitoring granularity, currently supporting 60 300 3600 86400, in seconds.When the time range does not exceed 1 day, support a minimum granularity of 60 seconds;When the time range does not exceed 7 days, support a minimum granularity of 3600 seconds;When the time range does not exceed 30 days, the minimum granularity supported is 86400 seconds.
        /// </summary>
        [Input("granularity", required: true)]
        public Input<int> Granularity { get; set; } = null!;

        /// <summary>
        /// Group Id.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("metricNames", required: true)]
        private InputList<string>? _metricNames;

        /// <summary>
        /// Metric Names. support, InBandwidth, OutBandwidth, Concurrent, InPackets, OutPackets.
        /// </summary>
        public InputList<string> MetricNames
        {
            get => _metricNames ?? (_metricNames = new InputList<string>());
            set => _metricNames = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start Time.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetProxyGroupStatisticsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProxyGroupStatisticsResult
    {
        public readonly string EndTime;
        public readonly int Granularity;
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> MetricNames;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;
        /// <summary>
        /// proxy Group Statistics.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProxyGroupStatisticsStatisticsDataResult> StatisticsDatas;

        [OutputConstructor]
        private GetProxyGroupStatisticsResult(
            string endTime,

            int granularity,

            string groupId,

            string id,

            ImmutableArray<string> metricNames,

            string? resultOutputFile,

            string startTime,

            ImmutableArray<Outputs.GetProxyGroupStatisticsStatisticsDataResult> statisticsDatas)
        {
            EndTime = endTime;
            Granularity = granularity;
            GroupId = groupId;
            Id = id;
            MetricNames = metricNames;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            StatisticsDatas = statisticsDatas;
        }
    }
}
