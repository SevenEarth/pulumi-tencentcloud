// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateEnhancedServiceAutomationServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable TencentCloud Automation Tools(TAT), TRUE or FALSE.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public LaunchTemplateEnhancedServiceAutomationServiceArgs()
        {
        }
    }
}
