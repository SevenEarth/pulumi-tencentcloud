// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmConditionsTemplateTemplateGroupListEventConditionResult
    {
        /// <summary>
        /// Alarm notification frequency.
        /// </summary>
        public readonly string AlarmNotifyPeriod;
        /// <summary>
        /// Predefined repeated notification strategy (0- alarm only once, 1- exponential alarm, 2- connection alarm).
        /// </summary>
        public readonly string AlarmNotifyType;
        /// <summary>
        /// Event Display Name (External).
        /// </summary>
        public readonly string EventDisplayName;
        /// <summary>
        /// Event ID.
        /// </summary>
        public readonly string EventId;
        /// <summary>
        /// Rule ID.
        /// </summary>
        public readonly string RuleId;

        [OutputConstructor]
        private GetAlarmConditionsTemplateTemplateGroupListEventConditionResult(
            string alarmNotifyPeriod,

            string alarmNotifyType,

            string eventDisplayName,

            string eventId,

            string ruleId)
        {
            AlarmNotifyPeriod = alarmNotifyPeriod;
            AlarmNotifyType = alarmNotifyType;
            EventDisplayName = eventDisplayName;
            EventId = eventId;
            RuleId = ruleId;
        }
    }
}
