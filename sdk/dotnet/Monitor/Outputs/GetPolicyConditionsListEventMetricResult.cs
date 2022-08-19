// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetPolicyConditionsListEventMetricResult
    {
        /// <summary>
        /// The ID of this event metric.
        /// </summary>
        public readonly int EventId;
        /// <summary>
        /// The name of this event metric.
        /// </summary>
        public readonly string EventShowName;
        /// <summary>
        /// Whether to recover.
        /// </summary>
        public readonly bool NeedRecovered;

        [OutputConstructor]
        private GetPolicyConditionsListEventMetricResult(
            int eventId,

            string eventShowName,

            bool needRecovered)
        {
            EventId = eventId;
            EventShowName = eventShowName;
            NeedRecovered = needRecovered;
        }
    }
}
