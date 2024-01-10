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
    public sealed class GetAlarmConditionsTemplateTemplateGroupListConditionResult
    {
        /// <summary>
        /// Alarm notification frequency.
        /// </summary>
        public readonly int AlarmNotifyPeriod;
        /// <summary>
        /// Predefined repeated notification strategy (0- alarm only once, 1- exponential alarm, 2- connection alarm).
        /// </summary>
        public readonly int AlarmNotifyType;
        /// <summary>
        /// Detection method.
        /// </summary>
        public readonly string CalcType;
        /// <summary>
        /// Detection value.
        /// </summary>
        public readonly string CalcValue;
        /// <summary>
        /// Duration in seconds.
        /// </summary>
        public readonly string ContinueTime;
        /// <summary>
        /// Whether it is an advanced indicator, 0: No; 1: Yes.
        /// </summary>
        public readonly int IsAdvanced;
        /// <summary>
        /// Whether to activate advanced indicators, 0: No; 1: Yes.
        /// </summary>
        public readonly int IsOpen;
        /// <summary>
        /// Indicator display name (external).
        /// </summary>
        public readonly string MetricDisplayName;
        /// <summary>
        /// Indicator ID.
        /// </summary>
        public readonly int MetricId;
        /// <summary>
        /// Cycle.
        /// </summary>
        public readonly int Period;
        /// <summary>
        /// Product ID.
        /// </summary>
        public readonly string ProductId;
        /// <summary>
        /// Rule ID.
        /// </summary>
        public readonly int RuleId;
        /// <summary>
        /// Indicator unit.
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private GetAlarmConditionsTemplateTemplateGroupListConditionResult(
            int alarmNotifyPeriod,

            int alarmNotifyType,

            string calcType,

            string calcValue,

            string continueTime,

            int isAdvanced,

            int isOpen,

            string metricDisplayName,

            int metricId,

            int period,

            string productId,

            int ruleId,

            string unit)
        {
            AlarmNotifyPeriod = alarmNotifyPeriod;
            AlarmNotifyType = alarmNotifyType;
            CalcType = calcType;
            CalcValue = calcValue;
            ContinueTime = continueTime;
            IsAdvanced = isAdvanced;
            IsOpen = isOpen;
            MetricDisplayName = metricDisplayName;
            MetricId = metricId;
            Period = period;
            ProductId = productId;
            RuleId = ruleId;
            Unit = unit;
        }
    }
}
