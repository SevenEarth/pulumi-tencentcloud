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

    public sealed class ScheduleTriggerAwsS3FileUploadTriggerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS SQS queue. This parameter is required if `NotifyType` is `AWS-SQS`.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("awsSqs")]
        public Input<Inputs.ScheduleTriggerAwsS3FileUploadTriggerAwsSqsGetArgs>? AwsSqs { get; set; }

        /// <summary>
        /// Input path directory bound to a workflow, such as `/movie/201907/`. If this parameter is left empty, the `/` root directory will be used.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        [Input("formats")]
        private InputList<string>? _formats;

        /// <summary>
        /// Format list of files that can trigger a workflow, such as [mp4, flv, mov]. If this parameter is left empty, files in all formats can trigger the workflow.
        /// </summary>
        public InputList<string> Formats
        {
            get => _formats ?? (_formats = new InputList<string>());
            set => _formats = value;
        }

        /// <summary>
        /// The AWS S3 bucket.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// The region of the AWS S3 bucket.
        /// </summary>
        [Input("s3Region", required: true)]
        public Input<string> S3Region { get; set; } = null!;

        /// <summary>
        /// The key ID required to access the AWS S3 object.
        /// </summary>
        [Input("s3SecretId")]
        public Input<string>? S3SecretId { get; set; }

        /// <summary>
        /// The key required to access the AWS S3 object.
        /// </summary>
        [Input("s3SecretKey")]
        public Input<string>? S3SecretKey { get; set; }

        public ScheduleTriggerAwsS3FileUploadTriggerGetArgs()
        {
        }
        public static new ScheduleTriggerAwsS3FileUploadTriggerGetArgs Empty => new ScheduleTriggerAwsS3FileUploadTriggerGetArgs();
    }
}
