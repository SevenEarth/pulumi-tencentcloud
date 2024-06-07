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

    public sealed class ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location to save the output object in COS. This parameter is valid and required when `Type` is COS.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("cosOutputStorage")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageCosOutputStorageArgs>? CosOutputStorage { get; set; }

        /// <summary>
        /// The AWS S3 bucket to save the output file. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("s3OutputStorage")]
        public Input<Inputs.ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageS3OutputStorageArgs>? S3OutputStorage { get; set; }

        /// <summary>
        /// The storage type for a media processing output file. Valid values: `COS`: Tencent Cloud COS `AWS-S3`: AWS S3. This type is only supported for AWS tasks, and the output bucket must be in the same region as the bucket of the source file.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageArgs()
        {
        }
        public static new ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageArgs Empty => new ScheduleActivityActivityParaAdaptiveDynamicStreamingTaskOutputStorageArgs();
    }
}
