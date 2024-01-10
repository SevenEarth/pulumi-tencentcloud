// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class DefaultAlarmThresholdDefaultAlarmConfig
    {
        /// <summary>
        /// Alarm threshold, in Mbps, with a value of&amp;gt;=0; When used as an input parameter, setting 0 will delete the alarm threshold configuration;.
        /// </summary>
        public readonly int? AlarmThreshold;
        /// <summary>
        /// Alarm threshold type, value [1 (incoming traffic alarm threshold) 2 (attack cleaning traffic alarm threshold)].
        /// </summary>
        public readonly int? AlarmType;

        [OutputConstructor]
        private DefaultAlarmThresholdDefaultAlarmConfig(
            int? alarmThreshold,

            int? alarmType)
        {
            AlarmThreshold = alarmThreshold;
            AlarmType = alarmType;
        }
    }
}
