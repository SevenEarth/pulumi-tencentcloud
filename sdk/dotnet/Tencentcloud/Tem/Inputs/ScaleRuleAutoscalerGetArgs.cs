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

    public sealed class ScaleRuleAutoscalerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name.
        /// </summary>
        [Input("autoscalerName", required: true)]
        public Input<string> AutoscalerName { get; set; } = null!;

        [Input("cronHorizontalAutoscalers")]
        private InputList<Inputs.ScaleRuleAutoscalerCronHorizontalAutoscalerGetArgs>? _cronHorizontalAutoscalers;

        /// <summary>
        /// scaler based on cron configuration.
        /// </summary>
        public InputList<Inputs.ScaleRuleAutoscalerCronHorizontalAutoscalerGetArgs> CronHorizontalAutoscalers
        {
            get => _cronHorizontalAutoscalers ?? (_cronHorizontalAutoscalers = new InputList<Inputs.ScaleRuleAutoscalerCronHorizontalAutoscalerGetArgs>());
            set => _cronHorizontalAutoscalers = value;
        }

        /// <summary>
        /// description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// enable AutoScaler.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("horizontalAutoscalers")]
        private InputList<Inputs.ScaleRuleAutoscalerHorizontalAutoscalerGetArgs>? _horizontalAutoscalers;

        /// <summary>
        /// scaler based on metrics.
        /// </summary>
        public InputList<Inputs.ScaleRuleAutoscalerHorizontalAutoscalerGetArgs> HorizontalAutoscalers
        {
            get => _horizontalAutoscalers ?? (_horizontalAutoscalers = new InputList<Inputs.ScaleRuleAutoscalerHorizontalAutoscalerGetArgs>());
            set => _horizontalAutoscalers = value;
        }

        /// <summary>
        /// maximal replica number.
        /// </summary>
        [Input("maxReplicas", required: true)]
        public Input<int> MaxReplicas { get; set; } = null!;

        /// <summary>
        /// minimal replica number.
        /// </summary>
        [Input("minReplicas", required: true)]
        public Input<int> MinReplicas { get; set; } = null!;

        public ScaleRuleAutoscalerGetArgs()
        {
        }
        public static new ScaleRuleAutoscalerGetArgs Empty => new ScaleRuleAutoscalerGetArgs();
    }
}
