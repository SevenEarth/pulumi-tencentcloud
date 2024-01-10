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
    public sealed class ProcessLiveStreamOperationAiAnalysisTask
    {
        /// <summary>
        /// Video content analysis template ID.
        /// </summary>
        public readonly int Definition;
        /// <summary>
        /// An extended parameter, whose value is a stringfied JSON.Note: This parameter is for customers with special requirements. It needs to be customized offline.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? ExtendedParameter;

        [OutputConstructor]
        private ProcessLiveStreamOperationAiAnalysisTask(
            int definition,

            string? extendedParameter)
        {
            Definition = definition;
            ExtendedParameter = extendedParameter;
        }
    }
}
