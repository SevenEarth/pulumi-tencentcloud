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
    public sealed class GetClusterResultContentOperationInfoInitResult
    {
        /// <summary>
        /// Reason for not displaying. Note: This field may return null, indicating no valid value.
        /// </summary>
        public readonly string DisabledReason;
        /// <summary>
        /// The availability of the button (whether it is clickable) may return null indicating that the information is not available.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Whether to display the button. Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly bool Supported;

        [OutputConstructor]
        private GetClusterResultContentOperationInfoInitResult(
            string disabledReason,

            bool enabled,

            bool supported)
        {
            DisabledReason = disabledReason;
            Enabled = enabled;
            Supported = supported;
        }
    }
}
