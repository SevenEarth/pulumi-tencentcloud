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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetResult
    {
        /// <summary>
        /// The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetCosInputInfoResult> CosInputInfos;
        /// <summary>
        /// The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetS3InputInfoResult> S3InputInfos;
        /// <summary>
        /// The trigger type. Valid values:`CosFileUpload`: Tencent Cloud COS trigger.`AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetUrlInputInfoResult> UrlInputInfos;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetResult(
            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetCosInputInfoResult> cosInputInfos,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetS3InputInfoResult> s3InputInfos,

            string type,

            ImmutableArray<Outputs.GetSchedulesScheduleInfoSetActivityActivityParaTranscodeTaskHeadTailParameterTailSetUrlInputInfoResult> urlInputInfos)
        {
            CosInputInfos = cosInputInfos;
            S3InputInfos = s3InputInfos;
            Type = type;
            UrlInputInfos = urlInputInfos;
        }
    }
}
