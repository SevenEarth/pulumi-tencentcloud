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
    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplateImageContentS3InputInfo
    {
        /// <summary>
        /// The AWS S3 bucket.
        /// </summary>
        public readonly string S3Bucket;
        /// <summary>
        /// The path of the AWS S3 object.
        /// </summary>
        public readonly string S3Object;
        /// <summary>
        /// The region of the AWS S3 bucket.
        /// </summary>
        public readonly string S3Region;
        /// <summary>
        /// The key ID required to access the AWS S3 object.
        /// </summary>
        public readonly string? S3SecretId;
        /// <summary>
        /// The key required to access the AWS S3 object.
        /// </summary>
        public readonly string? S3SecretKey;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplateImageContentS3InputInfo(
            string s3Bucket,

            string s3Object,

            string s3Region,

            string? s3SecretId,

            string? s3SecretKey)
        {
            S3Bucket = s3Bucket;
            S3Object = s3Object;
            S3Region = s3Region;
            S3SecretId = s3SecretId;
            S3SecretKey = s3SecretKey;
        }
    }
}
