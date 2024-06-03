// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    public static class GetMonitorReport
    {
        /// <summary>
        /// Use this data source to query detailed information of css monitor_report
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
        ///     var monitorReport = Tencentcloud.Css.GetMonitorReport.Invoke(new()
        ///     {
        ///         MonitorId = "0e8a12b5-df2a-4a1b-aa98-97d5610aa142",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetMonitorReportResult> InvokeAsync(GetMonitorReportArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitorReportResult>("tencentcloud:Css/getMonitorReport:getMonitorReport", args ?? new GetMonitorReportArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of css monitor_report
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
        ///     var monitorReport = Tencentcloud.Css.GetMonitorReport.Invoke(new()
        ///     {
        ///         MonitorId = "0e8a12b5-df2a-4a1b-aa98-97d5610aa142",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetMonitorReportResult> Invoke(GetMonitorReportInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorReportResult>("tencentcloud:Css/getMonitorReport:getMonitorReport", args ?? new GetMonitorReportInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitorReportArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor ID.
        /// </summary>
        [Input("monitorId", required: true)]
        public string MonitorId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetMonitorReportArgs()
        {
        }
        public static new GetMonitorReportArgs Empty => new GetMonitorReportArgs();
    }

    public sealed class GetMonitorReportInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor ID.
        /// </summary>
        [Input("monitorId", required: true)]
        public Input<string> MonitorId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetMonitorReportInvokeArgs()
        {
        }
        public static new GetMonitorReportInvokeArgs Empty => new GetMonitorReportInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitorReportResult
    {
        /// <summary>
        /// The information about the media diagnostic result.Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitorReportDiagnoseResultResult> DiagnoseResults;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MonitorId;
        /// <summary>
        /// The information about the media processing result.Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMonitorReportMpsResultResult> MpsResults;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetMonitorReportResult(
            ImmutableArray<Outputs.GetMonitorReportDiagnoseResultResult> diagnoseResults,

            string id,

            string monitorId,

            ImmutableArray<Outputs.GetMonitorReportMpsResultResult> mpsResults,

            string? resultOutputFile)
        {
            DiagnoseResults = diagnoseResults;
            Id = id;
            MonitorId = monitorId;
            MpsResults = mpsResults;
            ResultOutputFile = resultOutputFile;
        }
    }
}
