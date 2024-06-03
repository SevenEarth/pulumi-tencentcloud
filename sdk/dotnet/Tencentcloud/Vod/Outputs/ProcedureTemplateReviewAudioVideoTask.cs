// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Outputs
{

    [OutputType]
    public sealed class ProcedureTemplateReviewAudioVideoTask
    {
        /// <summary>
        /// Review template.
        /// </summary>
        public readonly string? Definition;
        /// <summary>
        /// The type of moderated content. Valid values:
        /// </summary>
        public readonly ImmutableArray<string> ReviewContents;

        [OutputConstructor]
        private ProcedureTemplateReviewAudioVideoTask(
            string? definition,

            ImmutableArray<string> reviewContents)
        {
            Definition = definition;
            ReviewContents = reviewContents;
        }
    }
}
