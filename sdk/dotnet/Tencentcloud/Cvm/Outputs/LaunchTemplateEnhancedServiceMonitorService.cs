// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateEnhancedServiceMonitorService
    {
        /// <summary>
        /// Whether to enable cloud monitor service, TRUE or FALSE.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private LaunchTemplateEnhancedServiceMonitorService(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
