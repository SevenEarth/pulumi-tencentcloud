// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetCompareTasksListCompareProcessStepResult
    {
        /// <summary>
        /// errors info.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCompareTasksListCompareProcessStepErrorResult> Errors;
        /// <summary>
        /// progress info.
        /// </summary>
        public readonly int? Percent;
        /// <summary>
        /// start time.
        /// </summary>
        public readonly string? StartTime;
        /// <summary>
        /// compare task status, optional value is created/readyRun/running/success/stopping/failed/canceled.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// step id.
        /// </summary>
        public readonly string? StepId;
        /// <summary>
        /// step message.
        /// </summary>
        public readonly string? StepMessage;
        /// <summary>
        /// step name.
        /// </summary>
        public readonly string? StepName;
        /// <summary>
        /// step number.
        /// </summary>
        public readonly int? StepNo;
        /// <summary>
        /// warnings info.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCompareTasksListCompareProcessStepWarningResult> Warnings;

        [OutputConstructor]
        private GetCompareTasksListCompareProcessStepResult(
            ImmutableArray<Outputs.GetCompareTasksListCompareProcessStepErrorResult> errors,

            int? percent,

            string? startTime,

            string? status,

            string? stepId,

            string? stepMessage,

            string? stepName,

            int? stepNo,

            ImmutableArray<Outputs.GetCompareTasksListCompareProcessStepWarningResult> warnings)
        {
            Errors = errors;
            Percent = percent;
            StartTime = startTime;
            Status = status;
            StepId = stepId;
            StepMessage = stepMessage;
            StepName = stepName;
            StepNo = stepNo;
            Warnings = warnings;
        }
    }
}
