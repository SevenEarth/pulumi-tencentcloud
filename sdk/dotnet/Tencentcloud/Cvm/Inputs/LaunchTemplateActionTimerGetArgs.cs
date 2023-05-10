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

    public sealed class LaunchTemplateActionTimerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Execution time.
        /// </summary>
        [Input("actionTime")]
        public Input<string>? ActionTime { get; set; }

        /// <summary>
        /// Extended data.
        /// </summary>
        [Input("externals")]
        public Input<Inputs.LaunchTemplateActionTimerExternalsGetArgs>? Externals { get; set; }

        /// <summary>
        /// Timer name.
        /// </summary>
        [Input("timerAction")]
        public Input<string>? TimerAction { get; set; }

        public LaunchTemplateActionTimerGetArgs()
        {
        }
    }
}