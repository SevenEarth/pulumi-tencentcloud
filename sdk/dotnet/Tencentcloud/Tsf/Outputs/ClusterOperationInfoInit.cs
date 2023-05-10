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
    public sealed class ClusterOperationInfoInit
    {
        /// <summary>
        /// Reason for not showing.
        /// </summary>
        public readonly string? DisabledReason;
        /// <summary>
        /// Is the button clickable.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// whether to show the button.
        /// </summary>
        public readonly bool? Supported;

        [OutputConstructor]
        private ClusterOperationInfoInit(
            string? disabledReason,

            bool? enabled,

            bool? supported)
        {
            DisabledReason = disabledReason;
            Enabled = enabled;
            Supported = supported;
        }
    }
}
