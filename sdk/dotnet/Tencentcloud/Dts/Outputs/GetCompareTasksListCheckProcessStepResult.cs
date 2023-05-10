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
    public sealed class GetCompareTasksListCheckProcessStepResult
    {
        public readonly ImmutableArray<Outputs.GetCompareTasksListCheckProcessStepErrorResult> Errors;
        /// <summary>
        /// progress info.
        /// </summary>
        public readonly int? Percent;
        public readonly string? StartTime;
        /// <summary>
        /// compare task status, optional value is created/readyRun/running/success/stopping/failed/canceled.
        /// </summary>
        public readonly string? Status;
        public readonly string? StepId;
        public readonly string? StepMessage;
        public readonly string? StepName;
        public readonly int? StepNo;
        public readonly ImmutableArray<Outputs.GetCompareTasksListCheckProcessStepWarningResult> Warnings;

        [OutputConstructor]
        private GetCompareTasksListCheckProcessStepResult(
            ImmutableArray<Outputs.GetCompareTasksListCheckProcessStepErrorResult> errors,

            int? percent,

            string? startTime,

            string? status,

            string? stepId,

            string? stepMessage,

            string? stepName,

            int? stepNo,

            ImmutableArray<Outputs.GetCompareTasksListCheckProcessStepWarningResult> warnings)
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
