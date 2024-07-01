// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class RealtimeLogDeliveryLogFormat
    {
        /// <summary>
        /// A string to be added before each log delivery batch. Each log delivery batch may contain multiple log records.
        /// </summary>
        public readonly string? BatchPrefix;
        /// <summary>
        /// A string to append after each log delivery batch.
        /// </summary>
        public readonly string? BatchSuffix;
        /// <summary>
        /// In a single log record, a string is inserted between fields as a separator. The possible values are: `	`: tab character; `,`: comma; `;`: semicolon.
        /// </summary>
        public readonly string? FieldDelimiter;
        /// <summary>
        /// The default output format type for log delivery. The possible values are: `json`: Use the default log output format JSON Lines. The fields in a single log are presented as key-value pairs; `csv`: Use the default log output format csv. Only field values are presented in a single log, without field names.
        /// </summary>
        public readonly string FormatType;
        /// <summary>
        /// The string inserted between log records as a separator. The possible values are: `
        /// `: newline character; `	`: tab character; `,`: comma.
        /// </summary>
        public readonly string? RecordDelimiter;
        /// <summary>
        /// A string to prepend to each log record.
        /// </summary>
        public readonly string? RecordPrefix;
        /// <summary>
        /// A string to append to each log record.
        /// 
        /// The `s3` object supports the following:
        /// </summary>
        public readonly string? RecordSuffix;

        [OutputConstructor]
        private RealtimeLogDeliveryLogFormat(
            string? batchPrefix,

            string? batchSuffix,

            string? fieldDelimiter,

            string formatType,

            string? recordDelimiter,

            string? recordPrefix,

            string? recordSuffix)
        {
            BatchPrefix = batchPrefix;
            BatchSuffix = batchSuffix;
            FieldDelimiter = fieldDelimiter;
            FormatType = formatType;
            RecordDelimiter = recordDelimiter;
            RecordPrefix = recordPrefix;
            RecordSuffix = recordSuffix;
        }
    }
}
