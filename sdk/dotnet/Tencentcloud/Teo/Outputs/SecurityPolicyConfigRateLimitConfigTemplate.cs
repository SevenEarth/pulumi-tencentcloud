// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyConfigRateLimitConfigTemplate
    {
        /// <summary>
        /// Detail of the template.
        /// </summary>
        public readonly Outputs.SecurityPolicyConfigRateLimitConfigTemplateDetail? Detail;
        /// <summary>
        /// Template Name. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private SecurityPolicyConfigRateLimitConfigTemplate(
            Outputs.SecurityPolicyConfigRateLimitConfigTemplateDetail? detail,

            string? mode)
        {
            Detail = detail;
            Mode = mode;
        }
    }
}
