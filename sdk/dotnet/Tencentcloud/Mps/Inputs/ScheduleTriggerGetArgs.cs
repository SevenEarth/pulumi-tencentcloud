// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleTriggerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS S3 trigger. This parameter is valid and required if `Type` is `AwsS3FileUpload`.Note: Currently, the key for the AWS S3 bucket, the trigger SQS queue, and the callback SQS queue must be the same.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("awsS3FileUploadTrigger")]
        public Input<Inputs.ScheduleTriggerAwsS3FileUploadTriggerGetArgs>? AwsS3FileUploadTrigger { get; set; }

        /// <summary>
        /// This parameter is required and valid when `Type` is `CosFileUpload`, indicating the COS trigger rule.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("cosFileUploadTrigger")]
        public Input<Inputs.ScheduleTriggerCosFileUploadTriggerGetArgs>? CosFileUploadTrigger { get; set; }

        /// <summary>
        /// The trigger type. Valid values: `CosFileUpload`: Tencent Cloud COS trigger. `AwsS3FileUpload`: AWS S3 trigger. Currently, this type is only supported for transcoding tasks and schemes (not supported for workflows).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ScheduleTriggerGetArgs()
        {
        }
    }
}
