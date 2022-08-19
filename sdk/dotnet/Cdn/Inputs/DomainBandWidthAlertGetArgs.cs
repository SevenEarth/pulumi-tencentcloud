// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainBandWidthAlertGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert percentage.
        /// </summary>
        [Input("alertPercentage")]
        public Input<int>? AlertPercentage { get; set; }

        /// <summary>
        /// Switch alert.
        /// </summary>
        [Input("alertSwitch")]
        public Input<string>? AlertSwitch { get; set; }

        /// <summary>
        /// threshold of bps.
        /// </summary>
        [Input("bpsThreshold")]
        public Input<int>? BpsThreshold { get; set; }

        /// <summary>
        /// Counter measure.
        /// </summary>
        [Input("counterMeasure")]
        public Input<string>? CounterMeasure { get; set; }

        [Input("lastTriggerTime")]
        public Input<string>? LastTriggerTime { get; set; }

        [Input("lastTriggerTimeOverseas")]
        public Input<string>? LastTriggerTimeOverseas { get; set; }

        /// <summary>
        /// Metric.
        /// </summary>
        [Input("metric")]
        public Input<string>? Metric { get; set; }

        /// <summary>
        /// Specify statistic item configuration.
        /// </summary>
        [Input("statisticItem")]
        public Input<Inputs.DomainBandWidthAlertStatisticItemGetArgs>? StatisticItem { get; set; }

        /// <summary>
        /// Configuration switch, available values: `on`, `off` (default).
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public DomainBandWidthAlertGetArgs()
        {
        }
    }
}
