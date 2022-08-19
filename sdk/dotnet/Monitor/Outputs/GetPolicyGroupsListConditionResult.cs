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
    public sealed class GetPolicyGroupsListConditionResult
    {
        /// <summary>
        /// Alarm sending cycle per second. `&lt;0` does not fire, `0` only fires once, and `&gt;0` fires every triggerTime second.
        /// </summary>
        public readonly int AlarmNotifyPeriod;
        /// <summary>
        /// Alarm sending convergence type. `0` continuous alarm, `1` index alarm.
        /// </summary>
        public readonly int AlarmNotifyType;
        /// <summary>
        /// Compare type, `1` means more than, `2`  means greater than or equal, `3` means less than, `4` means less than or equal to, `5` means equal, `6` means not equal, `7` means days rose, `8` means days fell, `9` means weeks rose, `10` means weeks fell, `11` means period rise, `12` means period fell.
        /// </summary>
        public readonly int CalcType;
        /// <summary>
        /// Threshold value.
        /// </summary>
        public readonly string CalcValue;
        /// <summary>
        /// How long does the triggering rule last (per second).
        /// </summary>
        public readonly int ContinueTime;
        /// <summary>
        /// The ID of this metric.
        /// </summary>
        public readonly int MetricId;
        /// <summary>
        /// The name of this metric.
        /// </summary>
        public readonly string MetricShowName;
        /// <summary>
        /// The unit of this metric.
        /// </summary>
        public readonly string MetricUnit;
        /// <summary>
        /// Data aggregation cycle (unit second).
        /// </summary>
        public readonly int Period;
        /// <summary>
        /// Threshold rule ID.
        /// </summary>
        public readonly int RuleId;

        [OutputConstructor]
        private GetPolicyGroupsListConditionResult(
            int alarmNotifyPeriod,

            int alarmNotifyType,

            int calcType,

            string calcValue,

            int continueTime,

            int metricId,

            string metricShowName,

            string metricUnit,

            int period,

            int ruleId)
        {
            AlarmNotifyPeriod = alarmNotifyPeriod;
            AlarmNotifyType = alarmNotifyType;
            CalcType = calcType;
            CalcValue = calcValue;
            ContinueTime = continueTime;
            MetricId = metricId;
            MetricShowName = metricShowName;
            MetricUnit = metricUnit;
            Period = period;
            RuleId = ruleId;
        }
    }
}
