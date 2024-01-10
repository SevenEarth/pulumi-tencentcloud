// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ses.Outputs
{

    [OutputType]
    public sealed class BatchSendEmailCycleParam
    {
        /// <summary>
        /// Start time of the task.
        /// </summary>
        public readonly string BeginTime;
        /// <summary>
        /// Task recurrence in hours.
        /// </summary>
        public readonly int IntervalTime;
        /// <summary>
        /// Specifies whether to end the cycle. This parameter is used to update the task. Valid values: 0: No; 1: Yes.
        /// </summary>
        public readonly int? TermCycle;

        [OutputConstructor]
        private BatchSendEmailCycleParam(
            string beginTime,

            int intervalTime,

            int? termCycle)
        {
            BeginTime = beginTime;
            IntervalTime = intervalTime;
            TermCycle = termCycle;
        }
    }
}
