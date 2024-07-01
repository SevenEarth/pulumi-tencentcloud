// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class RealtimeLogDeliveryLogFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A string to be added before each log delivery batch. Each log delivery batch may contain multiple log records.
        /// </summary>
        [Input("batchPrefix")]
        public Input<string>? BatchPrefix { get; set; }

        /// <summary>
        /// A string to append after each log delivery batch.
        /// </summary>
        [Input("batchSuffix")]
        public Input<string>? BatchSuffix { get; set; }

        /// <summary>
        /// In a single log record, a string is inserted between fields as a separator. The possible values are: `	`: tab character; `,`: comma; `;`: semicolon.
        /// </summary>
        [Input("fieldDelimiter")]
        public Input<string>? FieldDelimiter { get; set; }

        /// <summary>
        /// The default output format type for log delivery. The possible values are: `json`: Use the default log output format JSON Lines. The fields in a single log are presented as key-value pairs; `csv`: Use the default log output format csv. Only field values are presented in a single log, without field names.
        /// </summary>
        [Input("formatType", required: true)]
        public Input<string> FormatType { get; set; } = null!;

        /// <summary>
        /// The string inserted between log records as a separator. The possible values are: `
        /// `: newline character; `	`: tab character; `,`: comma.
        /// </summary>
        [Input("recordDelimiter")]
        public Input<string>? RecordDelimiter { get; set; }

        /// <summary>
        /// A string to prepend to each log record.
        /// </summary>
        [Input("recordPrefix")]
        public Input<string>? RecordPrefix { get; set; }

        /// <summary>
        /// A string to append to each log record.
        /// 
        /// The `s3` object supports the following:
        /// </summary>
        [Input("recordSuffix")]
        public Input<string>? RecordSuffix { get; set; }

        public RealtimeLogDeliveryLogFormatArgs()
        {
        }
        public static new RealtimeLogDeliveryLogFormatArgs Empty => new RealtimeLogDeliveryLogFormatArgs();
    }
}
