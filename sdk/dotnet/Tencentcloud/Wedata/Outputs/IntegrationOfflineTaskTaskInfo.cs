// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata.Outputs
{

    [OutputType]
    public sealed class IntegrationOfflineTaskTaskInfo
    {
        /// <summary>
        /// User App Id.
        /// </summary>
        public readonly string? AppId;
        /// <summary>
        /// Task configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoConfig> Configs;
        /// <summary>
        /// Create time.
        /// </summary>
        public readonly string? CreateTime;
        /// <summary>
        /// Creator User ID.
        /// </summary>
        public readonly string? CreatorUin;
        /// <summary>
        /// Data proxy url.
        /// </summary>
        public readonly ImmutableArray<string> DataProxyUrls;
        /// <summary>
        /// Execute context.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoExecuteContext> ExecuteContexts;
        /// <summary>
        /// Executor group name.
        /// </summary>
        public readonly string? ExecutorGroupName;
        /// <summary>
        /// Executor resource ID.
        /// </summary>
        public readonly string? ExecutorId;
        /// <summary>
        /// Node extension configuration information.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoExtConfig> ExtConfigs;
        /// <summary>
        /// Whether the task been submitted.
        /// </summary>
        public readonly bool? HasVersion;
        /// <summary>
        /// InLong manager url.
        /// </summary>
        public readonly string? InLongManagerUrl;
        /// <summary>
        /// InLong manager version.
        /// </summary>
        public readonly string? InLongManagerVersion;
        /// <summary>
        /// InLong stream id.
        /// </summary>
        public readonly string? InLongStreamId;
        /// <summary>
        /// Incharge user.
        /// </summary>
        public readonly string? Incharge;
        /// <summary>
        /// Input datasource type.
        /// </summary>
        public readonly string? InputDatasourceType;
        /// <summary>
        /// Instance version.
        /// </summary>
        public readonly int? InstanceVersion;
        /// <summary>
        /// The last time the task was run.
        /// </summary>
        public readonly string? LastRunTime;
        /// <summary>
        /// Whether the task been locked.
        /// </summary>
        public readonly bool? Locked;
        /// <summary>
        /// User locked task.
        /// </summary>
        public readonly string? Locker;
        /// <summary>
        /// Node mapping.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMapping> Mappings;
        /// <summary>
        /// Number of reads.
        /// </summary>
        public readonly int? NumRecordsIn;
        /// <summary>
        /// Number of writes.
        /// </summary>
        public readonly int? NumRecordsOut;
        /// <summary>
        /// Times of restarts.
        /// </summary>
        public readonly int? NumRestarts;
        /// <summary>
        /// Offline task scheduling configuration.
        /// </summary>
        public readonly Outputs.IntegrationOfflineTaskTaskInfoOfflineTaskAddEntity? OfflineTaskAddEntity;
        /// <summary>
        /// Operator User ID.
        /// </summary>
        public readonly string? OperatorUin;
        /// <summary>
        /// Output datasource type.
        /// </summary>
        public readonly string? OutputDatasourceType;
        /// <summary>
        /// Owner User ID.
        /// </summary>
        public readonly string? OwnerUin;
        /// <summary>
        /// Reading stage, 0: full amount, 1: partial full amount, 2: all incremental.
        /// </summary>
        public readonly int? ReadPhase;
        /// <summary>
        /// Read latency.
        /// </summary>
        public readonly double? ReaderDelay;
        /// <summary>
        /// The amount of resources consumed by real-time task.
        /// </summary>
        public readonly double? RunningCu;
        /// <summary>
        /// Task scheduling id (job id such as oceanus or us).
        /// </summary>
        public readonly string? ScheduleTaskId;
        /// <summary>
        /// Task status 1. Not started | Task initialization, 2. Task starting, 3. Running, 4. Paused, 5. Task stopping, 6. Stopped, 7. Execution failed, 8. deleted, 9. Locked, 404. unknown status.
        /// </summary>
        public readonly int? Status;
        /// <summary>
        /// The time the task was stopped.
        /// </summary>
        public readonly string? StopTime;
        /// <summary>
        /// Whether the task version has been submitted for operation and maintenance.
        /// </summary>
        public readonly bool? Submit;
        /// <summary>
        /// Resource tiering status, 0: in progress, 1: successful, 2: failed.
        /// </summary>
        public readonly int? SwitchResource;
        /// <summary>
        /// Synchronization type: 1. Whole database synchronization, 2. Single table synchronization.
        /// </summary>
        public readonly int? SyncType;
        /// <summary>
        /// Task alarm regular.
        /// </summary>
        public readonly ImmutableArray<string> TaskAlarmRegularLists;
        /// <summary>
        /// Inlong Task Group ID.
        /// </summary>
        public readonly string? TaskGroupId;
        /// <summary>
        /// Task display mode, 0: canvas mode, 1: form mode.
        /// </summary>
        public readonly string? TaskMode;
        /// <summary>
        /// Update time.
        /// </summary>
        public readonly string? UpdateTime;
        /// <summary>
        /// The workflow id to which the task belongs.
        /// </summary>
        public readonly string? WorkflowId;

        [OutputConstructor]
        private IntegrationOfflineTaskTaskInfo(
            string? appId,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoConfig> configs,

            string? createTime,

            string? creatorUin,

            ImmutableArray<string> dataProxyUrls,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoExecuteContext> executeContexts,

            string? executorGroupName,

            string? executorId,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoExtConfig> extConfigs,

            bool? hasVersion,

            string? inLongManagerUrl,

            string? inLongManagerVersion,

            string? inLongStreamId,

            string? incharge,

            string? inputDatasourceType,

            int? instanceVersion,

            string? lastRunTime,

            bool? locked,

            string? locker,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMapping> mappings,

            int? numRecordsIn,

            int? numRecordsOut,

            int? numRestarts,

            Outputs.IntegrationOfflineTaskTaskInfoOfflineTaskAddEntity? offlineTaskAddEntity,

            string? operatorUin,

            string? outputDatasourceType,

            string? ownerUin,

            int? readPhase,

            double? readerDelay,

            double? runningCu,

            string? scheduleTaskId,

            int? status,

            string? stopTime,

            bool? submit,

            int? switchResource,

            int? syncType,

            ImmutableArray<string> taskAlarmRegularLists,

            string? taskGroupId,

            string? taskMode,

            string? updateTime,

            string? workflowId)
        {
            AppId = appId;
            Configs = configs;
            CreateTime = createTime;
            CreatorUin = creatorUin;
            DataProxyUrls = dataProxyUrls;
            ExecuteContexts = executeContexts;
            ExecutorGroupName = executorGroupName;
            ExecutorId = executorId;
            ExtConfigs = extConfigs;
            HasVersion = hasVersion;
            InLongManagerUrl = inLongManagerUrl;
            InLongManagerVersion = inLongManagerVersion;
            InLongStreamId = inLongStreamId;
            Incharge = incharge;
            InputDatasourceType = inputDatasourceType;
            InstanceVersion = instanceVersion;
            LastRunTime = lastRunTime;
            Locked = locked;
            Locker = locker;
            Mappings = mappings;
            NumRecordsIn = numRecordsIn;
            NumRecordsOut = numRecordsOut;
            NumRestarts = numRestarts;
            OfflineTaskAddEntity = offlineTaskAddEntity;
            OperatorUin = operatorUin;
            OutputDatasourceType = outputDatasourceType;
            OwnerUin = ownerUin;
            ReadPhase = readPhase;
            ReaderDelay = readerDelay;
            RunningCu = runningCu;
            ScheduleTaskId = scheduleTaskId;
            Status = status;
            StopTime = stopTime;
            Submit = submit;
            SwitchResource = switchResource;
            SyncType = syncType;
            TaskAlarmRegularLists = taskAlarmRegularLists;
            TaskGroupId = taskGroupId;
            TaskMode = taskMode;
            UpdateTime = updateTime;
            WorkflowId = workflowId;
        }
    }
}
