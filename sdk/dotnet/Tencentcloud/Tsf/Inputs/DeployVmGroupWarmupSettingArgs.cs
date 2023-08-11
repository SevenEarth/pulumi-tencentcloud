// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class DeployVmGroupWarmupSettingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Preheating curvature, with a value between 1 and 5.
        /// </summary>
        [Input("curvature")]
        public Input<int>? Curvature { get; set; }

        /// <summary>
        /// Whether to enable preheating.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether to enable preheating protection. If protection is enabled and more than 50% of nodes are in preheating state, preheating will be aborted.
        /// </summary>
        [Input("enabledProtection")]
        public Input<bool>? EnabledProtection { get; set; }

        /// <summary>
        /// warmup time.
        /// </summary>
        [Input("warmupTime")]
        public Input<int>? WarmupTime { get; set; }

        public DeployVmGroupWarmupSettingArgs()
        {
        }
    }
}