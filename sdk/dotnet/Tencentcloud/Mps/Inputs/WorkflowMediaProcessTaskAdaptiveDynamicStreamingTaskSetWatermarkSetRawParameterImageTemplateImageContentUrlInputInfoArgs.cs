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

    public sealed class WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs()
        {
        }
        public static new WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs Empty => new WorkflowMediaProcessTaskAdaptiveDynamicStreamingTaskSetWatermarkSetRawParameterImageTemplateImageContentUrlInputInfoArgs();
    }
}
