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
    public sealed class SecurityPolicyConfigWafConfigAiRule
    {
        /// <summary>
        /// Valid values:- `smart_status_close`: disabled.- `smart_status_open`: blocked.- `smart_status_observe`: observed.
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private SecurityPolicyConfigWafConfigAiRule(string? mode)
        {
            Mode = mode;
        }
    }
}
