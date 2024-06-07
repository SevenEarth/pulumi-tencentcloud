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

    public sealed class ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS S3 bucket.
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// The path of the AWS S3 object.
        /// </summary>
        [Input("s3Object", required: true)]
        public Input<string> S3Object { get; set; } = null!;

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

        public ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs()
        {
        }
        public static new ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs Empty => new ScheduleActivityActivityParaSnapshotByTimeOffsetTaskWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs();
    }
}
