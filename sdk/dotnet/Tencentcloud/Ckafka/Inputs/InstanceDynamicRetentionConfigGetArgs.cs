// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class InstanceDynamicRetentionConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum retention time, in minutes.
        /// </summary>
        [Input("bottomRetention")]
        public Input<int>? BottomRetention { get; set; }

        /// <summary>
        /// Disk quota threshold (in percentage) for triggering the message retention time change event.
        /// </summary>
        [Input("diskQuotaPercentage")]
        public Input<int>? DiskQuotaPercentage { get; set; }

        /// <summary>
        /// Whether the dynamic message retention time configuration is enabled. 0: disabled; 1: enabled.
        /// </summary>
        [Input("enable")]
        public Input<int>? Enable { get; set; }

        /// <summary>
        /// Percentage by which the message retention time is shortened each time.
        /// </summary>
        [Input("stepForwardPercentage")]
        public Input<int>? StepForwardPercentage { get; set; }

        public InstanceDynamicRetentionConfigGetArgs()
        {
        }
        public static new InstanceDynamicRetentionConfigGetArgs Empty => new InstanceDynamicRetentionConfigGetArgs();
    }
}
