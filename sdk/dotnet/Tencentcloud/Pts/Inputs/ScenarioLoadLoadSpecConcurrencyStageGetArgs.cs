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

    public sealed class ScenarioLoadLoadSpecConcurrencyStageGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pressure time.
        /// </summary>
        [Input("durationSeconds")]
        public Input<int>? DurationSeconds { get; set; }

        /// <summary>
        /// Number of virtual users.
        /// </summary>
        [Input("targetVirtualUsers")]
        public Input<int>? TargetVirtualUsers { get; set; }

        public ScenarioLoadLoadSpecConcurrencyStageGetArgs()
        {
        }
    }
}