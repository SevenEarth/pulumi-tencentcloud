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

    public sealed class ScenarioLoadLoadSpecScriptOriginGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pressure testing time.
        /// </summary>
        [Input("durationSeconds", required: true)]
        public Input<int> DurationSeconds { get; set; } = null!;

        /// <summary>
        /// Number of machines.
        /// </summary>
        [Input("machineNumber", required: true)]
        public Input<int> MachineNumber { get; set; } = null!;

        /// <summary>
        /// Machine specification.
        /// </summary>
        [Input("machineSpecification", required: true)]
        public Input<string> MachineSpecification { get; set; } = null!;

        public ScenarioLoadLoadSpecScriptOriginGetArgs()
        {
        }
        public static new ScenarioLoadLoadSpecScriptOriginGetArgs Empty => new ScenarioLoadLoadSpecScriptOriginGetArgs();
    }
}
