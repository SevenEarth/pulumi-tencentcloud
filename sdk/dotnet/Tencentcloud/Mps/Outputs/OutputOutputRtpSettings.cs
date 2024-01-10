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
    public sealed class OutputOutputRtpSettings
    {
        /// <summary>
        /// The target address of the relay can be filled in 1~2.
        /// </summary>
        public readonly ImmutableArray<Outputs.OutputOutputRtpSettingsDestination> Destinations;
        /// <summary>
        /// You can only fill in none.
        /// </summary>
        public readonly string Fec;
        /// <summary>
        /// Idle timeout, unit ms.
        /// </summary>
        public readonly int IdleTimeout;

        [OutputConstructor]
        private OutputOutputRtpSettings(
            ImmutableArray<Outputs.OutputOutputRtpSettingsDestination> destinations,

            string fec,

            int idleTimeout)
        {
            Destinations = destinations;
            Fec = fec;
            IdleTimeout = idleTimeout;
        }
    }
}
