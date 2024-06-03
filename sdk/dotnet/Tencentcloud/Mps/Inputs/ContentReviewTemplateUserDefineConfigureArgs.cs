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

    public sealed class ContentReviewTemplateUserDefineConfigureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-defined asr text review control parameters.
        /// </summary>
        [Input("asrReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureAsrReviewInfoArgs>? AsrReviewInfo { get; set; }

        /// <summary>
        /// User-defined face review control parameters.
        /// </summary>
        [Input("faceReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureFaceReviewInfoArgs>? FaceReviewInfo { get; set; }

        /// <summary>
        /// User-defined ocr text review control parameters.
        /// </summary>
        [Input("ocrReviewInfo")]
        public Input<Inputs.ContentReviewTemplateUserDefineConfigureOcrReviewInfoArgs>? OcrReviewInfo { get; set; }

        public ContentReviewTemplateUserDefineConfigureArgs()
        {
        }
        public static new ContentReviewTemplateUserDefineConfigureArgs Empty => new ContentReviewTemplateUserDefineConfigureArgs();
    }
}
