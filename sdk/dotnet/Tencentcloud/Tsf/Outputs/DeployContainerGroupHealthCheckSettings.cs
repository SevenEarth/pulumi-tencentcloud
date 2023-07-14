// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class DeployContainerGroupHealthCheckSettings
    {
        /// <summary>
        /// Liveness probe. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.DeployContainerGroupHealthCheckSettingsLivenessProbe? LivenessProbe;
        /// <summary>
        /// Readiness health check. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly Outputs.DeployContainerGroupHealthCheckSettingsReadinessProbe? ReadinessProbe;

        [OutputConstructor]
        private DeployContainerGroupHealthCheckSettings(
            Outputs.DeployContainerGroupHealthCheckSettingsLivenessProbe? livenessProbe,

            Outputs.DeployContainerGroupHealthCheckSettingsReadinessProbe? readinessProbe)
        {
            LivenessProbe = livenessProbe;
            ReadinessProbe = readinessProbe;
        }
    }
}
