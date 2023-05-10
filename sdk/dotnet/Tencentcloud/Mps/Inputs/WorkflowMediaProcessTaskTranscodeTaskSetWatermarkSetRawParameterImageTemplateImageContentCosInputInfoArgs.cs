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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplateImageContentCosInputInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the COS Bucket where the media processing object file is located.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Input path for media processing object files.
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        /// <summary>
        /// The park to which the COS Bucket where the media processing target file resides belongs.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public WorkflowMediaProcessTaskTranscodeTaskSetWatermarkSetRawParameterImageTemplateImageContentCosInputInfoArgs()
        {
        }
    }
}
