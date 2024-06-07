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

    public sealed class LaunchTemplateEnhancedServiceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable TencentCloud Automation Tools(TAT).
        /// </summary>
        [Input("automationService")]
        public Input<Inputs.LaunchTemplateEnhancedServiceAutomationServiceGetArgs>? AutomationService { get; set; }

        /// <summary>
        /// Enable cloud monitor service.
        /// </summary>
        [Input("monitorService")]
        public Input<Inputs.LaunchTemplateEnhancedServiceMonitorServiceGetArgs>? MonitorService { get; set; }

        /// <summary>
        /// Enable cloud security service.
        /// </summary>
        [Input("securityService")]
        public Input<Inputs.LaunchTemplateEnhancedServiceSecurityServiceGetArgs>? SecurityService { get; set; }

        public LaunchTemplateEnhancedServiceGetArgs()
        {
        }
        public static new LaunchTemplateEnhancedServiceGetArgs Empty => new LaunchTemplateEnhancedServiceGetArgs();
    }
}
