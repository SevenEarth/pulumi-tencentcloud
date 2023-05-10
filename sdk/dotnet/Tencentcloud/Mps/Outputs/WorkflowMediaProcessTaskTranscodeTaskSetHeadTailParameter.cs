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
    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameter
    {
        /// <summary>
        /// Title list.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSet> HeadSets;
        /// <summary>
        /// Ending List.
        /// </summary>
        public readonly ImmutableArray<Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSet> TailSets;

        [OutputConstructor]
        private WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameter(
            ImmutableArray<Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSet> headSets,

            ImmutableArray<Outputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSet> tailSets)
        {
            HeadSets = headSets;
            TailSets = tailSets;
        }
    }
}
