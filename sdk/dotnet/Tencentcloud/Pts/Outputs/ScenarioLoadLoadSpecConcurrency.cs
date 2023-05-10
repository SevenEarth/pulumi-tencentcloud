// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class ScenarioLoadLoadSpecConcurrency
    {
        /// <summary>
        /// Wait time for graceful termination of the task.
        /// </summary>
        public readonly int? GracefulStopSeconds;
        /// <summary>
        /// Number of runs.
        /// </summary>
        public readonly int? IterationCount;
        /// <summary>
        /// Maximum RPS.
        /// </summary>
        public readonly int? MaxRequestsPerSecond;
        /// <summary>
        /// Multi-phase configuration array.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScenarioLoadLoadSpecConcurrencyStage> Stages;

        [OutputConstructor]
        private ScenarioLoadLoadSpecConcurrency(
            int? gracefulStopSeconds,

            int? iterationCount,

            int? maxRequestsPerSecond,

            ImmutableArray<Outputs.ScenarioLoadLoadSpecConcurrencyStage> stages)
        {
            GracefulStopSeconds = gracefulStopSeconds;
            IterationCount = iterationCount;
            MaxRequestsPerSecond = maxRequestsPerSecond;
            Stages = stages;
        }
    }
}
