// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css.Outputs
{

    [OutputType]
    public sealed class RecordTemplateAacParam
    {
        /// <summary>
        /// Classification of on-demand applications. This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? ClassId;
        /// <summary>
        /// Whether to enable recording in the current format. The default value is 0. 0: No, 1: Yes.
        /// </summary>
        public readonly int? Enable;
        /// <summary>
        /// Task flow. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Procedure;
        /// <summary>
        /// Recording interval. Unit: second, default: 1800. Value range: 30-7200. This parameter is invalid for HLS. When recording HLS, a file is generated from streaming to streaming.
        /// </summary>
        public readonly int? RecordInterval;
        /// <summary>
        /// Video storage strategy. Normal: standard storage. Cold: low frequency storage. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? StorageMode;
        /// <summary>
        /// Recording storage duration. Unit: second. Value range: 0 - 1500 days. 0: indicates permanent storage.
        /// </summary>
        public readonly int? StorageTime;
        /// <summary>
        /// Record file name.Special placeholders supported are: `StreamID`: Stream ID,`StartYear`: Start time - year,`StartMonth`: Start time - month,`StartDay`: Start time - day,`StartHour`: Start time - hour,`StartMinute`: Start time - minutes,`StartSecond`: Start time - seconds,`StartMillisecond`: Start time - milliseconds,`EndYear`: End time - year,`EndMonth`: End time - month,`EndDay`: End time - day,`EndHour`: End time - hour,`EndMinute`: End time - minutes,`EndSecond`: End time - seconds,`EndMillisecond`: End time - millisecondsIf the default recording file name is not set as ,`StreamID`,`StartYear`,`StartMonth`,`StartDay`,`StartHour`,`StartMinute`,`StartSecond`,`EndYear`,`EndMonth`,`EndDay`,`EndHour`,`EndMinute`,`EndSecond`.
        /// </summary>
        public readonly string? VodFileName;
        /// <summary>
        /// The ID of the on-demand sub app.
        /// </summary>
        public readonly int? VodSubAppId;

        [OutputConstructor]
        private RecordTemplateAacParam(
            int? classId,

            int? enable,

            string? procedure,

            int? recordInterval,

            string? storageMode,

            int? storageTime,

            string? vodFileName,

            int? vodSubAppId)
        {
            ClassId = classId;
            Enable = enable;
            Procedure = procedure;
            RecordInterval = recordInterval;
            StorageMode = storageMode;
            StorageTime = storageTime;
            VodFileName = vodFileName;
            VodSubAppId = vodSubAppId;
        }
    }
}
