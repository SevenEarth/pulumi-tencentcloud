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
    public sealed class EditMediaOperationTaskNotifyConfigAwsSqs
    {
        /// <summary>
        /// The key ID required to read from/write to the SQS queue.
        /// </summary>
        public readonly string? S3SecretId;
        /// <summary>
        /// The key required to read from/write to the SQS queue.
        /// </summary>
        public readonly string? S3SecretKey;
        /// <summary>
        /// The name of the SQS queue.
        /// </summary>
        public readonly string SqsQueueName;
        /// <summary>
        /// The region of the SQS queue.
        /// </summary>
        public readonly string SqsRegion;

        [OutputConstructor]
        private EditMediaOperationTaskNotifyConfigAwsSqs(
            string? s3SecretId,

            string? s3SecretKey,

            string sqsQueueName,

            string sqsRegion)
        {
            S3SecretId = s3SecretId;
            S3SecretKey = s3SecretKey;
            SqsQueueName = sqsQueueName;
            SqsRegion = sqsRegion;
        }
    }
}
