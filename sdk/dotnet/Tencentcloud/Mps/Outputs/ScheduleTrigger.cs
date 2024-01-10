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
    public sealed class ScheduleTrigger
    {
        /// <summary>
        /// The AWS S3 trigger. This parameter is valid and required if `Type` is `AwsS3FileUpload`.Note: Currently, the key for the AWS S3 bucket, the trigger SQS queue, and the callback SQS queue must be the same.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleTriggerAwsS3FileUploadTrigger? AwsS3FileUploadTrigger;
        /// <summary>
        /// This parameter is required and valid when `Type` is `CosFileUpload`, indicating the COS trigger rule.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.ScheduleTriggerCosFileUploadTrigger? CosFileUploadTrigger;
        /// <summary>
        /// The trigger type. Valid values: `CosFileUpload`: Tencent Cloud COS trigger. `AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ScheduleTrigger(
            Outputs.ScheduleTriggerAwsS3FileUploadTrigger? awsS3FileUploadTrigger,

            Outputs.ScheduleTriggerCosFileUploadTrigger? cosFileUploadTrigger,

            string type)
        {
            AwsS3FileUploadTrigger = awsS3FileUploadTrigger;
            CosFileUploadTrigger = cosFileUploadTrigger;
            Type = type;
        }
    }
}
