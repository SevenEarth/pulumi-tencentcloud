// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    public static class GetInstanceTaskList
    {
        /// <summary>
        /// Use this data source to query detailed information of redis instance_task_list
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
        ///         var instanceTaskList = Output.Create(Tencentcloud.Redis.GetInstanceTaskList.InvokeAsync(new Tencentcloud.Redis.GetInstanceTaskListArgs
        ///         {
        ///             BeginTime = "2021-12-30 00:00:00",
        ///             EndTime = "2021-12-30 00:00:00",
        ///             InstanceId = "crs-c1nl9rpv",
        ///             InstanceName = "",
        ///             OperateUins = 
        ///             {
        ///                 "",
        ///             },
        ///             ProjectIds = 
        ///             {
        ///                 "",
        ///             },
        ///             Results = 
        ///             {
        ///                 "",
        ///             },
        ///             TaskStatuses = 
        ///             {
        ///                 "",
        ///             },
        ///             TaskTypes = 
        ///             {
        ///                 "",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceTaskListResult> InvokeAsync(GetInstanceTaskListArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTaskListResult>("tencentcloud:Redis/getInstanceTaskList:getInstanceTaskList", args ?? new GetInstanceTaskListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of redis instance_task_list
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
        ///         var instanceTaskList = Output.Create(Tencentcloud.Redis.GetInstanceTaskList.InvokeAsync(new Tencentcloud.Redis.GetInstanceTaskListArgs
        ///         {
        ///             BeginTime = "2021-12-30 00:00:00",
        ///             EndTime = "2021-12-30 00:00:00",
        ///             InstanceId = "crs-c1nl9rpv",
        ///             InstanceName = "",
        ///             OperateUins = 
        ///             {
        ///                 "",
        ///             },
        ///             ProjectIds = 
        ///             {
        ///                 "",
        ///             },
        ///             Results = 
        ///             {
        ///                 "",
        ///             },
        ///             TaskStatuses = 
        ///             {
        ///                 "",
        ///             },
        ///             TaskTypes = 
        ///             {
        ///                 "",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceTaskListResult> Invoke(GetInstanceTaskListInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceTaskListResult>("tencentcloud:Redis/getInstanceTaskList:getInstanceTaskList", args ?? new GetInstanceTaskListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceTaskListArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Start time.
        /// </summary>
        [Input("beginTime")]
        public string? BeginTime { get; set; }

        /// <summary>
        /// Termination time.
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        [Input("operateUins")]
        private List<string>? _operateUins;

        /// <summary>
        /// Operator Uin.
        /// </summary>
        public List<string> OperateUins
        {
            get => _operateUins ?? (_operateUins = new List<string>());
            set => _operateUins = value;
        }

        [Input("projectIds")]
        private List<int>? _projectIds;

        /// <summary>
        /// Project Id.
        /// </summary>
        public List<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new List<int>());
            set => _projectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("results")]
        private List<int>? _results;

        /// <summary>
        /// Task status.
        /// </summary>
        public List<int> Results
        {
            get => _results ?? (_results = new List<int>());
            set => _results = value;
        }

        [Input("taskStatuses")]
        private List<int>? _taskStatuses;

        /// <summary>
        /// Task status.
        /// </summary>
        public List<int> TaskStatuses
        {
            get => _taskStatuses ?? (_taskStatuses = new List<int>());
            set => _taskStatuses = value;
        }

        [Input("taskTypes")]
        private List<string>? _taskTypes;

        /// <summary>
        /// Task type.
        /// </summary>
        public List<string> TaskTypes
        {
            get => _taskTypes ?? (_taskTypes = new List<string>());
            set => _taskTypes = value;
        }

        public GetInstanceTaskListArgs()
        {
        }
    }

    public sealed class GetInstanceTaskListInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Start time.
        /// </summary>
        [Input("beginTime")]
        public Input<string>? BeginTime { get; set; }

        /// <summary>
        /// Termination time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("operateUins")]
        private InputList<string>? _operateUins;

        /// <summary>
        /// Operator Uin.
        /// </summary>
        public InputList<string> OperateUins
        {
            get => _operateUins ?? (_operateUins = new InputList<string>());
            set => _operateUins = value;
        }

        [Input("projectIds")]
        private InputList<int>? _projectIds;

        /// <summary>
        /// Project Id.
        /// </summary>
        public InputList<int> ProjectIds
        {
            get => _projectIds ?? (_projectIds = new InputList<int>());
            set => _projectIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("results")]
        private InputList<int>? _results;

        /// <summary>
        /// Task status.
        /// </summary>
        public InputList<int> Results
        {
            get => _results ?? (_results = new InputList<int>());
            set => _results = value;
        }

        [Input("taskStatuses")]
        private InputList<int>? _taskStatuses;

        /// <summary>
        /// Task status.
        /// </summary>
        public InputList<int> TaskStatuses
        {
            get => _taskStatuses ?? (_taskStatuses = new InputList<int>());
            set => _taskStatuses = value;
        }

        [Input("taskTypes")]
        private InputList<string>? _taskTypes;

        /// <summary>
        /// Task type.
        /// </summary>
        public InputList<string> TaskTypes
        {
            get => _taskTypes ?? (_taskTypes = new InputList<string>());
            set => _taskTypes = value;
        }

        public GetInstanceTaskListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceTaskListResult
    {
        public readonly string? BeginTime;
        /// <summary>
        /// The end time.
        /// </summary>
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The name of instance.
        /// </summary>
        public readonly string? InstanceName;
        public readonly ImmutableArray<string> OperateUins;
        public readonly ImmutableArray<int> ProjectIds;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Task status.
        /// </summary>
        public readonly ImmutableArray<int> Results;
        public readonly ImmutableArray<int> TaskStatuses;
        public readonly ImmutableArray<string> TaskTypes;
        /// <summary>
        /// Task details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTaskListTaskResult> Tasks;

        [OutputConstructor]
        private GetInstanceTaskListResult(
            string? beginTime,

            string? endTime,

            string id,

            string? instanceId,

            string? instanceName,

            ImmutableArray<string> operateUins,

            ImmutableArray<int> projectIds,

            string? resultOutputFile,

            ImmutableArray<int> results,

            ImmutableArray<int> taskStatuses,

            ImmutableArray<string> taskTypes,

            ImmutableArray<Outputs.GetInstanceTaskListTaskResult> tasks)
        {
            BeginTime = beginTime;
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            OperateUins = operateUins;
            ProjectIds = projectIds;
            ResultOutputFile = resultOutputFile;
            Results = results;
            TaskStatuses = taskStatuses;
            TaskTypes = taskTypes;
            Tasks = tasks;
        }
    }
}