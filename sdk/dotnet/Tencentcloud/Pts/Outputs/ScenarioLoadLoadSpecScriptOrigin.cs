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
    public sealed class ScenarioLoadLoadSpecScriptOrigin
    {
        /// <summary>
        /// Pressure testing time.
        /// </summary>
        public readonly int DurationSeconds;
        /// <summary>
        /// Number of machines.
        /// </summary>
        public readonly int MachineNumber;
        /// <summary>
        /// Machine specification.
        /// </summary>
        public readonly string MachineSpecification;

        [OutputConstructor]
        private ScenarioLoadLoadSpecScriptOrigin(
            int durationSeconds,

            int machineNumber,

            string machineSpecification)
        {
            DurationSeconds = durationSeconds;
            MachineNumber = machineNumber;
            MachineSpecification = machineSpecification;
        }
    }
}
