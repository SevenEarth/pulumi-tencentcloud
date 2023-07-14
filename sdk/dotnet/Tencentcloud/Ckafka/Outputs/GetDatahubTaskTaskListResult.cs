// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetDatahubTaskTaskListResult
    {
        /// <summary>
        /// CreateTime.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Datahub Id.
        /// </summary>
        public readonly string DatahubId;
        /// <summary>
        /// ErrorMessage.
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// data resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceResult> SourceResources;
        /// <summary>
        /// Status, -1 failed to create, 0 to create, 1 to run, 2 to delete, 3 to deleted, 4 to delete failed, 5 to pause, 6 to pause, 7 to pause, 8 to resume, 9 to resume failed.
        /// </summary>
        public readonly int Status;
        /// <summary>
        /// StepList.
        /// </summary>
        public readonly ImmutableArray<string> StepLists;
        /// <summary>
        /// Target Resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatahubTaskTaskListTargetResourceResult> TargetResources;
        /// <summary>
        /// Task Current Step.
        /// </summary>
        public readonly string TaskCurrentStep;
        /// <summary>
        /// task ID.
        /// </summary>
        public readonly string TaskId;
        /// <summary>
        /// TaskName.
        /// </summary>
        public readonly string TaskName;
        /// <summary>
        /// Creation progress percentage.
        /// </summary>
        public readonly double TaskProgress;
        /// <summary>
        /// Task type, SOURCE|SINK.
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private GetDatahubTaskTaskListResult(
            string createTime,

            string datahubId,

            string errorMessage,

            ImmutableArray<Outputs.GetDatahubTaskTaskListSourceResourceResult> sourceResources,

            int status,

            ImmutableArray<string> stepLists,

            ImmutableArray<Outputs.GetDatahubTaskTaskListTargetResourceResult> targetResources,

            string taskCurrentStep,

            string taskId,

            string taskName,

            double taskProgress,

            string taskType)
        {
            CreateTime = createTime;
            DatahubId = datahubId;
            ErrorMessage = errorMessage;
            SourceResources = sourceResources;
            Status = status;
            StepLists = stepLists;
            TargetResources = targetResources;
            TaskCurrentStep = taskCurrentStep;
            TaskId = taskId;
            TaskName = taskName;
            TaskProgress = taskProgress;
            TaskType = taskType;
        }
    }
}
