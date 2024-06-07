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
    public static class GetListenerStatistics
    {
        /// <summary>
        /// Use this data source to query detailed information of gaap listener statistics
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
        ///     var listenerStatistics = Tencentcloud.Gaap.GetListenerStatistics.Invoke(new()
        ///     {
        ///         EndTime = "2023-10-19 23:59:59",
        ///         Granularity = 300,
        ///         ListenerId = "listener-xxxxxx",
        ///         MetricNames = new[]
        ///         {
        ///             "InBandwidth",
        ///             "OutBandwidth",
        ///             "InPackets",
        ///             "OutPackets",
        ///             "Concurrent",
        ///         },
        ///         StartTime = "2023-10-19 00:00:00",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetListenerStatisticsResult> InvokeAsync(GetListenerStatisticsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetListenerStatisticsResult>("tencentcloud:Gaap/getListenerStatistics:getListenerStatistics", args ?? new GetListenerStatisticsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of gaap listener statistics
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
        ///     var listenerStatistics = Tencentcloud.Gaap.GetListenerStatistics.Invoke(new()
        ///     {
        ///         EndTime = "2023-10-19 23:59:59",
        ///         Granularity = 300,
        ///         ListenerId = "listener-xxxxxx",
        ///         MetricNames = new[]
        ///         {
        ///             "InBandwidth",
        ///             "OutBandwidth",
        ///             "InPackets",
        ///             "OutPackets",
        ///             "Concurrent",
        ///         },
        ///         StartTime = "2023-10-19 00:00:00",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetListenerStatisticsResult> Invoke(GetListenerStatisticsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetListenerStatisticsResult>("tencentcloud:Gaap/getListenerStatistics:getListenerStatistics", args ?? new GetListenerStatisticsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetListenerStatisticsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End Time.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Monitoring granularity, currently supporting 300 3600 86400, in seconds.The query time range does not exceed 1 day and supports a minimum granularity of 300 seconds;The query interval should not exceed 7 days and support a minimum granularity of 3600 seconds;The query interval exceeds 7 days and supports a minimum granularity of 86400 seconds.
        /// </summary>
        [Input("granularity", required: true)]
        public int Granularity { get; set; }

        /// <summary>
        /// Listener Id.
        /// </summary>
        [Input("listenerId", required: true)]
        public string ListenerId { get; set; } = null!;

        [Input("metricNames", required: true)]
        private List<string>? _metricNames;

        /// <summary>
        /// List of statistical indicator names. Supporting: InBandwidth, OutBandwidth, Concurrent, InPackets, OutPackets.
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

        public GetListenerStatisticsArgs()
        {
        }
        public static new GetListenerStatisticsArgs Empty => new GetListenerStatisticsArgs();
    }

    public sealed class GetListenerStatisticsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End Time.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Monitoring granularity, currently supporting 300 3600 86400, in seconds.The query time range does not exceed 1 day and supports a minimum granularity of 300 seconds;The query interval should not exceed 7 days and support a minimum granularity of 3600 seconds;The query interval exceeds 7 days and supports a minimum granularity of 86400 seconds.
        /// </summary>
        [Input("granularity", required: true)]
        public Input<int> Granularity { get; set; } = null!;

        /// <summary>
        /// Listener Id.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        [Input("metricNames", required: true)]
        private InputList<string>? _metricNames;

        /// <summary>
        /// List of statistical indicator names. Supporting: InBandwidth, OutBandwidth, Concurrent, InPackets, OutPackets.
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

        public GetListenerStatisticsInvokeArgs()
        {
        }
        public static new GetListenerStatisticsInvokeArgs Empty => new GetListenerStatisticsInvokeArgs();
    }


    [OutputType]
    public sealed class GetListenerStatisticsResult
    {
        public readonly string EndTime;
        public readonly int Granularity;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ListenerId;
        public readonly ImmutableArray<string> MetricNames;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;
        /// <summary>
        /// Channel Group Statistics.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetListenerStatisticsStatisticsDataResult> StatisticsDatas;

        [OutputConstructor]
        private GetListenerStatisticsResult(
            string endTime,

            int granularity,

            string id,

            string listenerId,

            ImmutableArray<string> metricNames,

            string? resultOutputFile,

            string startTime,

            ImmutableArray<Outputs.GetListenerStatisticsStatisticsDataResult> statisticsDatas)
        {
            EndTime = endTime;
            Granularity = granularity;
            Id = id;
            ListenerId = listenerId;
            MetricNames = metricNames;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            StatisticsDatas = statisticsDatas;
        }
    }
}
