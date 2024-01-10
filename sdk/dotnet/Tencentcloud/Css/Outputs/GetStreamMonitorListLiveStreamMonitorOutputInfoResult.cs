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
    public sealed class GetStreamMonitorListLiveStreamMonitorOutputInfoResult
    {
        /// <summary>
        /// The playback path for the monitoring task.It should be within 32 bytes and can only contain letters, numbers, `-`, `_`, and `.` characters.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly string OutputApp;
        /// <summary>
        /// The playback domain for the monitoring task.It should be within 128 bytes and can only be filled with an enabled playback domain.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly string OutputDomain;
        /// <summary>
        /// The height of the output stream in pixels for the monitoring task. The range is [1, 1080]. It is recommended to be at least 100 pixels.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly int OutputStreamHeight;
        /// <summary>
        /// The name of the output stream for the monitoring task.If not specified, the system will generate a name automatically.The name should be within 256 bytes and can only contain letters, numbers, `-`, `_`, and `.` characters.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly string OutputStreamName;
        /// <summary>
        /// The width of the output stream in pixels for the monitoring task. The range is [1, 1920]. It is recommended to be at least 100 pixels.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly int OutputStreamWidth;

        [OutputConstructor]
        private GetStreamMonitorListLiveStreamMonitorOutputInfoResult(
            string outputApp,

            string outputDomain,

            int outputStreamHeight,

            string outputStreamName,

            int outputStreamWidth)
        {
            OutputApp = outputApp;
            OutputDomain = outputDomain;
            OutputStreamHeight = outputStreamHeight;
            OutputStreamName = outputStreamName;
            OutputStreamWidth = outputStreamWidth;
        }
    }
}
