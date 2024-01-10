// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetLogstashInstanceOperationsOperationTaskProcessInfoResult
    {
        /// <summary>
        /// Completed quantity.
        /// </summary>
        public readonly int Completed;
        /// <summary>
        /// Remaining quantity.
        /// </summary>
        public readonly int Remain;
        /// <summary>
        /// Task type. 60: restart task 70: fragment migration task 80: node modification task.
        /// </summary>
        public readonly int TaskType;
        /// <summary>
        /// Total quantity.
        /// </summary>
        public readonly int Total;

        [OutputConstructor]
        private GetLogstashInstanceOperationsOperationTaskProcessInfoResult(
            int completed,

            int remain,

            int taskType,

            int total)
        {
            Completed = completed;
            Remain = remain;
            TaskType = taskType;
            Total = total;
        }
    }
}
