// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf task
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var task = new Tencentcloud.Tsf.Task("task", new Tencentcloud.Tsf.TaskArgs
    ///         {
    ///             AdvanceSettings = new Tencentcloud.Tsf.Inputs.TaskAdvanceSettingsArgs
    ///             {
    ///                 SubTaskConcurrency = 2,
    ///             },
    ///             ExecuteType = "unicast",
    ///             GroupId = "group-y8pnmoga",
    ///             RetryCount = 0,
    ///             RetryInterval = 0,
    ///             SuccessOperator = "GTE",
    ///             SuccessRatio = "100",
    ///             TaskArgument = "a=c",
    ///             TaskContent = "/test",
    ///             TaskName = "terraform-test",
    ///             TaskRule = new Tencentcloud.Tsf.Inputs.TaskTaskRuleArgs
    ///             {
    ///                 Expression = "0 * 1 * * ? ",
    ///                 RuleType = "Cron",
    ///             },
    ///             TaskType = "java",
    ///             TimeOut = 60000,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tsf task can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tsf/task:Task task task-y37eqq95
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/task:Task")]
    public partial class Task : Pulumi.CustomResource
    {
        /// <summary>
        /// advanced settings.
        /// </summary>
        [Output("advanceSettings")]
        public Output<Outputs.TaskAdvanceSettings> AdvanceSettings { get; private set; } = null!;

        /// <summary>
        /// ID of the workflow to which it belongs.
        /// </summary>
        [Output("belongFlowIds")]
        public Output<ImmutableArray<string>> BelongFlowIds { get; private set; } = null!;

        /// <summary>
        /// execution type, unicast/broadcast.
        /// </summary>
        [Output("executeType")]
        public Output<string> ExecuteType { get; private set; } = null!;

        /// <summary>
        /// deployment group ID.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// Program id list.
        /// </summary>
        [Output("programIdLists")]
        public Output<ImmutableArray<string>> ProgramIdLists { get; private set; } = null!;

        /// <summary>
        /// number of retries, 0 &amp;amp;lt;= RetryCount&amp;amp;lt;= 10.
        /// </summary>
        [Output("retryCount")]
        public Output<int> RetryCount { get; private set; } = null!;

        /// <summary>
        /// retry interval, 0 &amp;amp;lt;= RetryInterval &amp;amp;lt;= 600000, time unit ms.
        /// </summary>
        [Output("retryInterval")]
        public Output<int> RetryInterval { get; private set; } = null!;

        /// <summary>
        /// Fragmentation parameters.
        /// </summary>
        [Output("shardArguments")]
        public Output<ImmutableArray<Outputs.TaskShardArgument>> ShardArguments { get; private set; } = null!;

        /// <summary>
        /// number of shards.
        /// </summary>
        [Output("shardCount")]
        public Output<int> ShardCount { get; private set; } = null!;

        /// <summary>
        /// the operator to judge the success of the task.
        /// </summary>
        [Output("successOperator")]
        public Output<string> SuccessOperator { get; private set; } = null!;

        /// <summary>
        /// The threshold for judging the success rate of the task, such as 100.
        /// </summary>
        [Output("successRatio")]
        public Output<string> SuccessRatio { get; private set; } = null!;

        /// <summary>
        /// task parameters, the length limit is 10000 characters.
        /// </summary>
        [Output("taskArgument")]
        public Output<string> TaskArgument { get; private set; } = null!;

        /// <summary>
        /// task content, length limit 65536 bytes.
        /// </summary>
        [Output("taskContent")]
        public Output<string> TaskContent { get; private set; } = null!;

        /// <summary>
        /// task ID.
        /// </summary>
        [Output("taskId")]
        public Output<string> TaskId { get; private set; } = null!;

        /// <summary>
        /// task history ID.
        /// </summary>
        [Output("taskLogId")]
        public Output<string> TaskLogId { get; private set; } = null!;

        /// <summary>
        /// task name, task length 64 characters.
        /// </summary>
        [Output("taskName")]
        public Output<string> TaskName { get; private set; } = null!;

        /// <summary>
        /// trigger rule.
        /// </summary>
        [Output("taskRule")]
        public Output<Outputs.TaskTaskRule> TaskRule { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the task, ENABLED/DISABLED.
        /// </summary>
        [Output("taskState")]
        public Output<string> TsfTaskState { get; private set; } = null!;

        /// <summary>
        /// task type, java.
        /// </summary>
        [Output("taskType")]
        public Output<string> TaskType { get; private set; } = null!;

        /// <summary>
        /// task timeout, time unit ms.
        /// </summary>
        [Output("timeOut")]
        public Output<int> TimeOut { get; private set; } = null!;

        /// <summary>
        /// trigger type.
        /// </summary>
        [Output("triggerType")]
        public Output<string> TriggerType { get; private set; } = null!;


        /// <summary>
        /// Create a Task resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Task(string name, TaskArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/task:Task", name, args ?? new TaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Task(string name, Input<string> id, TaskState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/task:Task", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Task resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Task Get(string name, Input<string> id, TaskState? state = null, CustomResourceOptions? options = null)
        {
            return new Task(name, id, state, options);
        }
    }

    public sealed class TaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// advanced settings.
        /// </summary>
        [Input("advanceSettings")]
        public Input<Inputs.TaskAdvanceSettingsArgs>? AdvanceSettings { get; set; }

        /// <summary>
        /// execution type, unicast/broadcast.
        /// </summary>
        [Input("executeType", required: true)]
        public Input<string> ExecuteType { get; set; } = null!;

        /// <summary>
        /// deployment group ID.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        /// <summary>
        /// number of retries, 0 &amp;amp;lt;= RetryCount&amp;amp;lt;= 10.
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// retry interval, 0 &amp;amp;lt;= RetryInterval &amp;amp;lt;= 600000, time unit ms.
        /// </summary>
        [Input("retryInterval")]
        public Input<int>? RetryInterval { get; set; }

        [Input("shardArguments")]
        private InputList<Inputs.TaskShardArgumentArgs>? _shardArguments;

        /// <summary>
        /// Fragmentation parameters.
        /// </summary>
        public InputList<Inputs.TaskShardArgumentArgs> ShardArguments
        {
            get => _shardArguments ?? (_shardArguments = new InputList<Inputs.TaskShardArgumentArgs>());
            set => _shardArguments = value;
        }

        /// <summary>
        /// number of shards.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// the operator to judge the success of the task.
        /// </summary>
        [Input("successOperator")]
        public Input<string>? SuccessOperator { get; set; }

        /// <summary>
        /// The threshold for judging the success rate of the task, such as 100.
        /// </summary>
        [Input("successRatio")]
        public Input<string>? SuccessRatio { get; set; }

        /// <summary>
        /// task parameters, the length limit is 10000 characters.
        /// </summary>
        [Input("taskArgument")]
        public Input<string>? TaskArgument { get; set; }

        /// <summary>
        /// task content, length limit 65536 bytes.
        /// </summary>
        [Input("taskContent", required: true)]
        public Input<string> TaskContent { get; set; } = null!;

        /// <summary>
        /// task name, task length 64 characters.
        /// </summary>
        [Input("taskName", required: true)]
        public Input<string> TaskName { get; set; } = null!;

        /// <summary>
        /// trigger rule.
        /// </summary>
        [Input("taskRule")]
        public Input<Inputs.TaskTaskRuleArgs>? TaskRule { get; set; }

        /// <summary>
        /// task type, java.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<string> TaskType { get; set; } = null!;

        /// <summary>
        /// task timeout, time unit ms.
        /// </summary>
        [Input("timeOut", required: true)]
        public Input<int> TimeOut { get; set; } = null!;

        public TaskArgs()
        {
        }
    }

    public sealed class TaskState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// advanced settings.
        /// </summary>
        [Input("advanceSettings")]
        public Input<Inputs.TaskAdvanceSettingsGetArgs>? AdvanceSettings { get; set; }

        [Input("belongFlowIds")]
        private InputList<string>? _belongFlowIds;

        /// <summary>
        /// ID of the workflow to which it belongs.
        /// </summary>
        public InputList<string> BelongFlowIds
        {
            get => _belongFlowIds ?? (_belongFlowIds = new InputList<string>());
            set => _belongFlowIds = value;
        }

        /// <summary>
        /// execution type, unicast/broadcast.
        /// </summary>
        [Input("executeType")]
        public Input<string>? ExecuteType { get; set; }

        /// <summary>
        /// deployment group ID.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("programIdLists")]
        private InputList<string>? _programIdLists;

        /// <summary>
        /// Program id list.
        /// </summary>
        public InputList<string> ProgramIdLists
        {
            get => _programIdLists ?? (_programIdLists = new InputList<string>());
            set => _programIdLists = value;
        }

        /// <summary>
        /// number of retries, 0 &amp;amp;lt;= RetryCount&amp;amp;lt;= 10.
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// retry interval, 0 &amp;amp;lt;= RetryInterval &amp;amp;lt;= 600000, time unit ms.
        /// </summary>
        [Input("retryInterval")]
        public Input<int>? RetryInterval { get; set; }

        [Input("shardArguments")]
        private InputList<Inputs.TaskShardArgumentGetArgs>? _shardArguments;

        /// <summary>
        /// Fragmentation parameters.
        /// </summary>
        public InputList<Inputs.TaskShardArgumentGetArgs> ShardArguments
        {
            get => _shardArguments ?? (_shardArguments = new InputList<Inputs.TaskShardArgumentGetArgs>());
            set => _shardArguments = value;
        }

        /// <summary>
        /// number of shards.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// the operator to judge the success of the task.
        /// </summary>
        [Input("successOperator")]
        public Input<string>? SuccessOperator { get; set; }

        /// <summary>
        /// The threshold for judging the success rate of the task, such as 100.
        /// </summary>
        [Input("successRatio")]
        public Input<string>? SuccessRatio { get; set; }

        /// <summary>
        /// task parameters, the length limit is 10000 characters.
        /// </summary>
        [Input("taskArgument")]
        public Input<string>? TaskArgument { get; set; }

        /// <summary>
        /// task content, length limit 65536 bytes.
        /// </summary>
        [Input("taskContent")]
        public Input<string>? TaskContent { get; set; }

        /// <summary>
        /// task ID.
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// task history ID.
        /// </summary>
        [Input("taskLogId")]
        public Input<string>? TaskLogId { get; set; }

        /// <summary>
        /// task name, task length 64 characters.
        /// </summary>
        [Input("taskName")]
        public Input<string>? TaskName { get; set; }

        /// <summary>
        /// trigger rule.
        /// </summary>
        [Input("taskRule")]
        public Input<Inputs.TaskTaskRuleGetArgs>? TaskRule { get; set; }

        /// <summary>
        /// Whether to enable the task, ENABLED/DISABLED.
        /// </summary>
        [Input("taskState")]
        public Input<string>? TsfTaskState { get; set; }

        /// <summary>
        /// task type, java.
        /// </summary>
        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// task timeout, time unit ms.
        /// </summary>
        [Input("timeOut")]
        public Input<int>? TimeOut { get; set; }

        /// <summary>
        /// trigger type.
        /// </summary>
        [Input("triggerType")]
        public Input<string>? TriggerType { get; set; }

        public TaskState()
        {
        }
    }
}
