// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// start time.
        /// </summary>
        [Input("startAt", required: true)]
        public Input<string> StartAt { get; set; } = null!;

        /// <summary>
        /// target replica number.
        /// </summary>
        [Input("targetReplicas", required: true)]
        public Input<int> TargetReplicas { get; set; } = null!;

        public ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleGetArgs()
        {
        }
        public static new ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleGetArgs Empty => new ScaleRuleAutoscalerCronHorizontalAutoscalerScheduleGetArgs();
    }
}
