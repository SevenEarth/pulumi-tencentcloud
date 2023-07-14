// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BatchReportGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delivery bucket for task completion reports.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Whether to output the task completion report.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<string> Enabled { get; set; } = null!;

        /// <summary>
        /// Task completion report format information. Legal value: Report_CSV_V1.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// Prefix information for the task completion report. Length 0-256 bytes.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Task completion report the task information that needs to be recorded to determine whether to record the execution information of all operations or the information of failed operations. Legal values: AllTasks, FailedTasksOnly.
        /// </summary>
        [Input("reportScope", required: true)]
        public Input<string> ReportScope { get; set; } = null!;

        public BatchReportGetArgs()
        {
        }
    }
}
