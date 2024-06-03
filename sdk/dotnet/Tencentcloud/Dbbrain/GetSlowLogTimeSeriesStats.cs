// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    public static class GetSlowLogTimeSeriesStats
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_time_series_stats
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
        ///     var test = Tencentcloud.Dbbrain.GetSlowLogTimeSeriesStats.Invoke(new()
        ///     {
        ///         EndTime = "%s",
        ///         InstanceId = "%s",
        ///         Product = "mysql",
        ///         StartTime = "%s",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSlowLogTimeSeriesStatsResult> InvokeAsync(GetSlowLogTimeSeriesStatsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSlowLogTimeSeriesStatsResult>("tencentcloud:Dbbrain/getSlowLogTimeSeriesStats:getSlowLogTimeSeriesStats", args ?? new GetSlowLogTimeSeriesStatsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_log_time_series_stats
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
        ///     var test = Tencentcloud.Dbbrain.GetSlowLogTimeSeriesStats.Invoke(new()
        ///     {
        ///         EndTime = "%s",
        ///         InstanceId = "%s",
        ///         Product = "mysql",
        ///         StartTime = "%s",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSlowLogTimeSeriesStatsResult> Invoke(GetSlowLogTimeSeriesStatsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSlowLogTimeSeriesStatsResult>("tencentcloud:Dbbrain/getSlowLogTimeSeriesStats:getSlowLogTimeSeriesStats", args ?? new GetSlowLogTimeSeriesStatsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowLogTimeSeriesStatsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time, such as `2019-09-10 12:13:14`, the interval between the end time and the start time can be up to 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public string? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, such as `2019-09-10 12:13:14`.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public GetSlowLogTimeSeriesStatsArgs()
        {
        }
        public static new GetSlowLogTimeSeriesStatsArgs Empty => new GetSlowLogTimeSeriesStatsArgs();
    }

    public sealed class GetSlowLogTimeSeriesStatsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// End time, such as `2019-09-10 12:13:14`, the interval between the end time and the start time can be up to 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, such as `2019-09-10 12:13:14`.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public GetSlowLogTimeSeriesStatsInvokeArgs()
        {
        }
        public static new GetSlowLogTimeSeriesStatsInvokeArgs Empty => new GetSlowLogTimeSeriesStatsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSlowLogTimeSeriesStatsResult
    {
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// The unit time interval between bars, in seconds.
        /// </summary>
        public readonly int Period;
        public readonly string? Product;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Instan1ce cpu utilization monitoring data within a unit time interval.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowLogTimeSeriesStatsSeriesDataResult> SeriesDatas;
        public readonly string StartTime;
        /// <summary>
        /// Statistics on the number of slow logs in a unit time interval.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowLogTimeSeriesStatsTimeSeriesResult> TimeSeries;

        [OutputConstructor]
        private GetSlowLogTimeSeriesStatsResult(
            string endTime,

            string id,

            string instanceId,

            int period,

            string? product,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSlowLogTimeSeriesStatsSeriesDataResult> seriesDatas,

            string startTime,

            ImmutableArray<Outputs.GetSlowLogTimeSeriesStatsTimeSeriesResult> timeSeries)
        {
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Period = period;
            Product = product;
            ResultOutputFile = resultOutputFile;
            SeriesDatas = seriesDatas;
            StartTime = startTime;
            TimeSeries = timeSeries;
        }
    }
}
