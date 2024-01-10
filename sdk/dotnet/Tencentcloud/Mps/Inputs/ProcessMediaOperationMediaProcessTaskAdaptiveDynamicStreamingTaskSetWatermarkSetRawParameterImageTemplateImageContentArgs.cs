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

    public sealed class ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The information of the COS object to process. This parameter is valid and required when `Type` is `COS`.
        /// </summary>
        [Input("cosInputInfo")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentCosInputInfoArgs>? CosInputInfo { get; set; }

        /// <summary>
        /// The information of the AWS S3 object processed. This parameter is required if `Type` is `AWS-S3`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("s3InputInfo")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentS3InputInfoArgs>? S3InputInfo { get; set; }

        /// <summary>
        /// The input type. Valid values:`COS`: A COS bucket address. `URL`: A URL. `AWS-S3`: An AWS S3 bucket address. Currently, this type is only supported for transcoding tasks.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The URL of the object to process. This parameter is valid and required when `Type` is `URL`.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("urlInputInfo")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs>? UrlInputInfo { get; set; }

        public ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentArgs()
        {
        }
    }
}
