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

    public sealed class ContentReviewTemplateTerrorismConfigureOcrReviewInfoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The score threshold for judging suspected violations. When the smart review reaches the score above, it is considered suspected violations. If it is not filled, the default is 100 points. Value range: 0~100.
        /// </summary>
        [Input("blockConfidence")]
        public Input<int>? BlockConfidence { get; set; }

        /// <summary>
        /// The score threshold for judging whether manual review is required for violations. When the intelligent review reaches the score above, it is considered that manual review is required. If it is not filled, the default is 75 points. Value range: 0~100.
        /// </summary>
        [Input("reviewConfidence")]
        public Input<int>? ReviewConfidence { get; set; }

        /// <summary>
        /// User-defined ocr text review task switch, optional value:ON/OFF.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public ContentReviewTemplateTerrorismConfigureOcrReviewInfoGetArgs()
        {
        }
        public static new ContentReviewTemplateTerrorismConfigureOcrReviewInfoGetArgs Empty => new ContentReviewTemplateTerrorismConfigureOcrReviewInfoGetArgs();
    }
}
