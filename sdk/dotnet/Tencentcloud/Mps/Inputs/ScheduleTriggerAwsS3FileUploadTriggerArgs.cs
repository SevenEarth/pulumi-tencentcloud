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

    public sealed class ScheduleTriggerAwsS3FileUploadTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SQS queue of the AWS S3 bucket.Note: The queue must be in the same region as the bucket.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("awsSqs")]
        public Input<Inputs.ScheduleTriggerAwsS3FileUploadTriggerAwsSqsArgs>? AwsSqs { get; set; }

        /// <summary>
        /// The bucket directory bound. It must be an absolute path that starts and ends with `/`, such as `/movie/201907/`. If you do not specify this, the root directory will be bound.	.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        [Input("formats")]
        private InputList<string>? _formats;

        /// <summary>
        /// The file formats that will trigger the scheme, such as [mp4, flv, mov]. If you do not specify this, the upload of files in any format will trigger the scheme.	.
        /// </summary>
        public InputList<string> Formats
        {
            get => _formats ?? (_formats = new InputList<string>());
            set => _formats = value;
        }

        /// <summary>
        /// The AWS S3 bucket bound to the scheme.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// The region of the AWS S3 bucket.
        /// </summary>
        [Input("s3Region", required: true)]
        public Input<string> S3Region { get; set; } = null!;

        /// <summary>
        /// The key ID of the AWS S3 bucket.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("s3SecretId")]
        public Input<string>? S3SecretId { get; set; }

        /// <summary>
        /// The key of the AWS S3 bucket.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("s3SecretKey")]
        public Input<string>? S3SecretKey { get; set; }

        public ScheduleTriggerAwsS3FileUploadTriggerArgs()
        {
        }
    }
}
