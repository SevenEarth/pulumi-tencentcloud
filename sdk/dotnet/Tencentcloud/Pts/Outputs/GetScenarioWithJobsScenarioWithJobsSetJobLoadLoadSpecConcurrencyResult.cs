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
    public sealed class GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyResult
    {
        /// <summary>
        /// The waiting period for a graceful shutdown.
        /// </summary>
        public readonly int GracefulStopSeconds;
        /// <summary>
        /// The iteration count of the load test.
        /// </summary>
        public readonly int IterationCount;
        /// <summary>
        /// The maximum RPS.
        /// </summary>
        public readonly int MaxRequestsPerSecond;
        /// <summary>
        /// The recource count of the load test.
        /// </summary>
        public readonly int Resources;
        /// <summary>
        /// The configuration for the multi-stage load test.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyStageResult> Stages;

        [OutputConstructor]
        private GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyResult(
            int gracefulStopSeconds,

            int iterationCount,

            int maxRequestsPerSecond,

            int resources,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetJobLoadLoadSpecConcurrencyStageResult> stages)
        {
            GracefulStopSeconds = gracefulStopSeconds;
            IterationCount = iterationCount;
            MaxRequestsPerSecond = maxRequestsPerSecond;
            Resources = resources;
            Stages = stages;
        }
    }
}
