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
    public sealed class ContentReviewTemplatePornConfigureImgReviewInfo
    {
        /// <summary>
        /// The score threshold for judging suspected violations. When the smart review reaches the score above, it is considered suspected violations. If it is not filled, the default is 90 points. Value range: 0~100.
        /// </summary>
        public readonly int? BlockConfidence;
        /// <summary>
        /// Terrorism image filter tag, if the review result contains the selected tag, the result will be returned, if the filter tag is empty, all the review results will be returned, the optional value is:guns, crowd, bloody, police, banners, militant, explosion, terrorists, scenario.
        /// </summary>
        public readonly ImmutableArray<string> LabelSets;
        /// <summary>
        /// The score threshold for judging whether manual review is required for violations. When the intelligent review reaches the score above, it is considered that manual review is required. If it is not filled, the default is 80 points. Value range: 0~100.
        /// </summary>
        public readonly int? ReviewConfidence;
        /// <summary>
        /// Terrorism image task switch, optional value:ON/OFF.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private ContentReviewTemplatePornConfigureImgReviewInfo(
            int? blockConfidence,

            ImmutableArray<string> labelSets,

            int? reviewConfidence,

            string @switch)
        {
            BlockConfidence = blockConfidence;
            LabelSets = labelSets;
            ReviewConfidence = reviewConfidence;
            Switch = @switch;
        }
    }
}
