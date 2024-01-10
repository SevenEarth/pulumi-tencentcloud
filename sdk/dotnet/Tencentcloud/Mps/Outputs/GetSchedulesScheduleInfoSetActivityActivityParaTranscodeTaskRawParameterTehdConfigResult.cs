// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskRawParameterTehdConfigResult
    {
        /// <summary>
        /// Maximum bitrate, which is valid when `Type` is `TESHD`. If this parameter is left empty or 0 is entered, there will be no upper limit for bitrate.
        /// </summary>
        public readonly int MaxVideoBitrate;
        /// <summary>
        /// The trigger type. Valid values:`CosFileUpload`: Tencent Cloud COS trigger.`AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskRawParameterTehdConfigResult(
            int maxVideoBitrate,

            string type)
        {
            MaxVideoBitrate = maxVideoBitrate;
            Type = type;
        }
    }
}
