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
    public sealed class ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSet
    {
        /// <summary>
        /// The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetCosInputInfo? CosInputInfo;
        /// <summary>
        /// The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetS3InputInfo? S3InputInfo;
        /// <summary>
        /// The input type. Valid values: `COS`: A COS bucket address.  `URL`: A URL.  `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetUrlInputInfo? UrlInputInfo;

        [OutputConstructor]
        private ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSet(
            Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetCosInputInfo? cosInputInfo,

            Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetS3InputInfo? s3InputInfo,

            string type,

            Outputs.ScheduleActivityActivityParaTranscodeTaskHeadTailParameterHeadSetUrlInputInfo? urlInputInfo)
        {
            CosInputInfo = cosInputInfo;
            S3InputInfo = s3InputInfo;
            Type = type;
            UrlInputInfo = urlInputInfo;
        }
    }
}
