// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    public static class GetUserTask
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql user_task
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var userTask = Output.Create(Tencentcloud.Mysql.GetUserTask.InvokeAsync(new Tencentcloud.Mysql.GetUserTaskArgs
        ///         {
        ///             AsyncRequestId = "f2fe828c-773af816-0a08f542-94bb2a9c",
        ///             InstanceId = "cdb-fitq5t9h",
        ///             StartTimeBegin = "2017-12-31 10:40:01",
        ///             StartTimeEnd = "2017-12-31 10:40:01",
        ///             TaskStatuses = 
        ///             {
        ///                 "2",
        ///             },
        ///             TaskTypes = 
        ///             {
        ///                 "5",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserTaskResult> InvokeAsync(GetUserTaskArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserTaskResult>("tencentcloud:Mysql/getUserTask:getUserTask", args ?? new GetUserTaskArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql user_task
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var userTask = Output.Create(Tencentcloud.Mysql.GetUserTask.InvokeAsync(new Tencentcloud.Mysql.GetUserTaskArgs
        ///         {
        ///             AsyncRequestId = "f2fe828c-773af816-0a08f542-94bb2a9c",
        ///             InstanceId = "cdb-fitq5t9h",
        ///             StartTimeBegin = "2017-12-31 10:40:01",
        ///             StartTimeEnd = "2017-12-31 10:40:01",
        ///             TaskStatuses = 
        ///             {
        ///                 "2",
        ///             },
        ///             TaskTypes = 
        ///             {
        ///                 "5",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserTaskResult> Invoke(GetUserTaskInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserTaskResult>("tencentcloud:Mysql/getUserTask:getUserTask", args ?? new GetUserTaskInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserTaskArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asynchronous task request ID, the AsyncRequestId returned by executing cloud database-related operations.
        /// </summary>
        [Input("asyncRequestId")]
        public string? AsyncRequestId { get; set; }

        /// <summary>
        /// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the value of the field InstanceId in the output parameter.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// The start time of the first task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
        /// </summary>
        [Input("startTimeBegin")]
        public string? StartTimeBegin { get; set; }

        /// <summary>
        /// The start time of the last task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
        /// </summary>
        [Input("startTimeEnd")]
        public string? StartTimeEnd { get; set; }

        [Input("taskStatuses")]
        private List<string>? _taskStatuses;

        /// <summary>
        /// Task status. If no value is passed, all task statuses will be queried. Supported values include: `UNDEFINED` - undefined; `INITIAL` - initialization; `RUNNING` - running; `SUCCEED` - the execution was successful; `FAILED` - execution failed; `KILLED` - terminated; `REMOVED` - removed; `PAUSED` - Paused.
        /// </summary>
        public List<string> TaskStatuses
        {
            get => _taskStatuses ?? (_taskStatuses = new List<string>());
            set => _taskStatuses = value;
        }

        [Input("taskTypes")]
        private List<string>? _taskTypes;

        /// <summary>
        /// Task type. If no value is passed, all task types will be queried. Supported values include: `ROLLBACK` - database rollback; `SQL OPERATION` - SQL operation; `IMPORT DATA` - data import; `MODIFY PARAM` - parameter setting; `INITIAL` - initialize the cloud database instance; `REBOOT` - restarts the cloud database instance; `OPEN GTID` - open the cloud database instance GTID; `UPGRADE RO` - read-only instance upgrade; `BATCH ROLLBACK` - database batch rollback; `UPGRADE MASTER` - master upgrade; `DROP TABLES` - delete cloud database tables; `SWITCH DR TO MASTER` - The disaster recovery instance.
        /// </summary>
        public List<string> TaskTypes
        {
            get => _taskTypes ?? (_taskTypes = new List<string>());
            set => _taskTypes = value;
        }

        public GetUserTaskArgs()
        {
        }
    }

    public sealed class GetUserTaskInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Asynchronous task request ID, the AsyncRequestId returned by executing cloud database-related operations.
        /// </summary>
        [Input("asyncRequestId")]
        public Input<string>? AsyncRequestId { get; set; }

        /// <summary>
        /// Instance ID, the format is: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page, and you can use the [query instance list] (https://cloud.tencent.com/document/api/236/15872) interface Gets the value of the field InstanceId in the output parameter.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// The start time of the first task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
        /// </summary>
        [Input("startTimeBegin")]
        public Input<string>? StartTimeBegin { get; set; }

        /// <summary>
        /// The start time of the last task, used for range query, the time format is as follows: 2017-12-31 10:40:01.
        /// </summary>
        [Input("startTimeEnd")]
        public Input<string>? StartTimeEnd { get; set; }

        [Input("taskStatuses")]
        private InputList<string>? _taskStatuses;

        /// <summary>
        /// Task status. If no value is passed, all task statuses will be queried. Supported values include: `UNDEFINED` - undefined; `INITIAL` - initialization; `RUNNING` - running; `SUCCEED` - the execution was successful; `FAILED` - execution failed; `KILLED` - terminated; `REMOVED` - removed; `PAUSED` - Paused.
        /// </summary>
        public InputList<string> TaskStatuses
        {
            get => _taskStatuses ?? (_taskStatuses = new InputList<string>());
            set => _taskStatuses = value;
        }

        [Input("taskTypes")]
        private InputList<string>? _taskTypes;

        /// <summary>
        /// Task type. If no value is passed, all task types will be queried. Supported values include: `ROLLBACK` - database rollback; `SQL OPERATION` - SQL operation; `IMPORT DATA` - data import; `MODIFY PARAM` - parameter setting; `INITIAL` - initialize the cloud database instance; `REBOOT` - restarts the cloud database instance; `OPEN GTID` - open the cloud database instance GTID; `UPGRADE RO` - read-only instance upgrade; `BATCH ROLLBACK` - database batch rollback; `UPGRADE MASTER` - master upgrade; `DROP TABLES` - delete cloud database tables; `SWITCH DR TO MASTER` - The disaster recovery instance.
        /// </summary>
        public InputList<string> TaskTypes
        {
            get => _taskTypes ?? (_taskTypes = new InputList<string>());
            set => _taskTypes = value;
        }

        public GetUserTaskInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserTaskResult
    {
        /// <summary>
        /// The request ID of the asynchronous task.
        /// </summary>
        public readonly string? AsyncRequestId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        /// <summary>
        /// The returned instance task information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserTaskItemResult> Items;
        public readonly string? ResultOutputFile;
        public readonly string? StartTimeBegin;
        public readonly string? StartTimeEnd;
        /// <summary>
        /// Instance task status, possible values include:UNDEFINED - undefined;INITIAL - initialization;RUNNING - running;SUCCEED - the execution was successful;FAILED - execution failed;KILLED - terminated;REMOVED - removed;PAUSED - Paused.WAITING - waiting (cancellable).
        /// </summary>
        public readonly ImmutableArray<string> TaskStatuses;
        public readonly ImmutableArray<string> TaskTypes;

        [OutputConstructor]
        private GetUserTaskResult(
            string? asyncRequestId,

            string id,

            string? instanceId,

            ImmutableArray<Outputs.GetUserTaskItemResult> items,

            string? resultOutputFile,

            string? startTimeBegin,

            string? startTimeEnd,

            ImmutableArray<string> taskStatuses,

            ImmutableArray<string> taskTypes)
        {
            AsyncRequestId = asyncRequestId;
            Id = id;
            InstanceId = instanceId;
            Items = items;
            ResultOutputFile = resultOutputFile;
            StartTimeBegin = startTimeBegin;
            StartTimeEnd = startTimeEnd;
            TaskStatuses = taskStatuses;
            TaskTypes = taskTypes;
        }
    }
}
