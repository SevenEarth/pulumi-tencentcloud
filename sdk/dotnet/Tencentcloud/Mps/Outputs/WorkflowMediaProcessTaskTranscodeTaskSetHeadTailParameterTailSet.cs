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
    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSet
    {
        /// <summary>
        /// Valid when Type is COS, this item is required, indicating media processing COS object information.
        /// </summary>
        public readonly Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetCosInputInfo? CosInputInfo;
        /// <summary>
        /// Enter the type of source object, which supports COS and URL.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Valid when Type is URL, this item is required, indicating media processing URL object information.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetUrlInputInfo? UrlInputInfo;

        [OutputConstructor]
        private WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSet(
            Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetCosInputInfo? cosInputInfo,

            string type,

            Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetUrlInputInfo? urlInputInfo)
        {
            CosInputInfo = cosInputInfo;
            Type = type;
            UrlInputInfo = urlInputInfo;
        }
    }
}
