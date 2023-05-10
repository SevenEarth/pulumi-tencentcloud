// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class ScenarioLoadLoadSpecConcurrencyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Wait time for graceful termination of the task.
        /// </summary>
        [Input("gracefulStopSeconds")]
        public Input<int>? GracefulStopSeconds { get; set; }

        /// <summary>
        /// Number of runs.
        /// </summary>
        [Input("iterationCount")]
        public Input<int>? IterationCount { get; set; }

        /// <summary>
        /// Maximum RPS.
        /// </summary>
        [Input("maxRequestsPerSecond")]
        public Input<int>? MaxRequestsPerSecond { get; set; }

        [Input("stages")]
        private InputList<Inputs.ScenarioLoadLoadSpecConcurrencyStageGetArgs>? _stages;

        /// <summary>
        /// Multi-phase configuration array.
        /// </summary>
        public InputList<Inputs.ScenarioLoadLoadSpecConcurrencyStageGetArgs> Stages
        {
            get => _stages ?? (_stages = new InputList<Inputs.ScenarioLoadLoadSpecConcurrencyStageGetArgs>());
            set => _stages = value;
        }

        public ScenarioLoadLoadSpecConcurrencyGetArgs()
        {
        }
    }
}