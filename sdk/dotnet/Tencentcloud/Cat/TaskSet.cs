// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cat
{
    /// <summary>
    /// Provides a resource to create a cat task_set
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var taskSet = new Tencentcloud.Cat.TaskSet("taskSet", new()
    ///     {
    ///         BatchTasks = new Tencentcloud.Cat.Inputs.TaskSetBatchTasksArgs
    ///         {
    ///             Name = "demo",
    ///             TargetAddress = "http://www.baidu.com",
    ///         },
    ///         TaskType = 5,
    ///         Nodes = new[]
    ///         {
    ///             "12136",
    ///             "12137",
    ///             "12138",
    ///             "12141",
    ///             "12144",
    ///         },
    ///         Interval = 6,
    ///         Parameters = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["ipType"] = 0,
    ///             ["grabBag"] = 0,
    ///             ["filterIp"] = 0,
    ///             ["netIcmpOn"] = 1,
    ///             ["netIcmpActivex"] = 0,
    ///             ["netIcmpTimeout"] = 20,
    ///             ["netIcmpInterval"] = 0.5,
    ///             ["netIcmpNum"] = 20,
    ///             ["netIcmpSize"] = 32,
    ///             ["netIcmpDataCut"] = 1,
    ///             ["netDnsOn"] = 1,
    ///             ["netDnsTimeout"] = 5,
    ///             ["netDnsQuerymethod"] = 1,
    ///             ["netDnsNs"] = "",
    ///             ["netDigOn"] = 1,
    ///             ["netDnsServer"] = 2,
    ///             ["netTracertOn"] = 1,
    ///             ["netTracertTimeout"] = 60,
    ///             ["netTracertNum"] = 30,
    ///             ["whiteList"] = "",
    ///             ["blackList"] = "",
    ///             ["netIcmpActivexStr"] = "",
    ///         }),
    ///         TaskCategory = 1,
    ///         Cron = "* 0-6 * * *",
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// cat task_set can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cat/taskSet:TaskSet task_set taskSet_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cat/taskSet:TaskSet")]
    public partial class TaskSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Batch task name address.
        /// </summary>
        [Output("batchTasks")]
        public Output<Outputs.TaskSetBatchTasks> BatchTasks { get; private set; } = null!;

        /// <summary>
        /// Timer task cron expression.
        /// </summary>
        [Output("cron")]
        public Output<string?> Cron { get; private set; } = null!;

        /// <summary>
        /// Task interval minutes in (1,5,10,15,30,60,120,240).
        /// </summary>
        [Output("interval")]
        public Output<int> Interval { get; private set; } = null!;

        /// <summary>
        /// `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        /// </summary>
        [Output("nodeIpType")]
        public Output<int> NodeIpType { get; private set; } = null!;

        /// <summary>
        /// Task Nodes.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<string>> Nodes { get; private set; } = null!;

        /// <summary>
        /// The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        /// </summary>
        [Output("operate")]
        public Output<string?> Operate { get; private set; } = null!;

        /// <summary>
        /// tasks parameters.
        /// </summary>
        [Output("parameters")]
        public Output<string> Parameters { get; private set; } = null!;

        /// <summary>
        /// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Task category,1:PC,2:Mobile.
        /// </summary>
        [Output("taskCategory")]
        public Output<int> TaskCategory { get; private set; } = null!;

        /// <summary>
        /// Task Id.
        /// </summary>
        [Output("taskId")]
        public Output<string> TaskId { get; private set; } = null!;

        /// <summary>
        /// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        /// </summary>
        [Output("taskType")]
        public Output<int> TaskType { get; private set; } = null!;


        /// <summary>
        /// Create a TaskSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TaskSet(string name, TaskSetArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cat/taskSet:TaskSet", name, args ?? new TaskSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TaskSet(string name, Input<string> id, TaskSetState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cat/taskSet:TaskSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TaskSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TaskSet Get(string name, Input<string> id, TaskSetState? state = null, CustomResourceOptions? options = null)
        {
            return new TaskSet(name, id, state, options);
        }
    }

    public sealed class TaskSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch task name address.
        /// </summary>
        [Input("batchTasks", required: true)]
        public Input<Inputs.TaskSetBatchTasksArgs> BatchTasks { get; set; } = null!;

        /// <summary>
        /// Timer task cron expression.
        /// </summary>
        [Input("cron")]
        public Input<string>? Cron { get; set; }

        /// <summary>
        /// Task interval minutes in (1,5,10,15,30,60,120,240).
        /// </summary>
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        /// <summary>
        /// `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        /// </summary>
        [Input("nodeIpType")]
        public Input<int>? NodeIpType { get; set; }

        [Input("nodes", required: true)]
        private InputList<string>? _nodes;

        /// <summary>
        /// Task Nodes.
        /// </summary>
        public InputList<string> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<string>());
            set => _nodes = value;
        }

        /// <summary>
        /// The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        /// </summary>
        [Input("operate")]
        public Input<string>? Operate { get; set; }

        /// <summary>
        /// tasks parameters.
        /// </summary>
        [Input("parameters", required: true)]
        public Input<string> Parameters { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Task category,1:PC,2:Mobile.
        /// </summary>
        [Input("taskCategory", required: true)]
        public Input<int> TaskCategory { get; set; } = null!;

        /// <summary>
        /// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        /// </summary>
        [Input("taskType", required: true)]
        public Input<int> TaskType { get; set; } = null!;

        public TaskSetArgs()
        {
        }
        public static new TaskSetArgs Empty => new TaskSetArgs();
    }

    public sealed class TaskSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Batch task name address.
        /// </summary>
        [Input("batchTasks")]
        public Input<Inputs.TaskSetBatchTasksGetArgs>? BatchTasks { get; set; }

        /// <summary>
        /// Timer task cron expression.
        /// </summary>
        [Input("cron")]
        public Input<string>? Cron { get; set; }

        /// <summary>
        /// Task interval minutes in (1,5,10,15,30,60,120,240).
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// `0`-Unlimit ip type, `1`-IPv4, `2`-IPv6.
        /// </summary>
        [Input("nodeIpType")]
        public Input<int>? NodeIpType { get; set; }

        [Input("nodes")]
        private InputList<string>? _nodes;

        /// <summary>
        /// Task Nodes.
        /// </summary>
        public InputList<string> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<string>());
            set => _nodes = value;
        }

        /// <summary>
        /// The input is valid when the parameter is modified, `suspend`/`resume`, used to suspend/resume the dial test task.
        /// </summary>
        [Input("operate")]
        public Input<string>? Operate { get; set; }

        /// <summary>
        /// tasks parameters.
        /// </summary>
        [Input("parameters")]
        public Input<string>? Parameters { get; set; }

        /// <summary>
        /// Task status 1:TaskPending, 2:TaskRunning,3:TaskRunException,4:TaskSuspending 5:TaskSuspendException,6:TaskSuspendException,7:TaskSuspended,9:TaskDeleted.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Task category,1:PC,2:Mobile.
        /// </summary>
        [Input("taskCategory")]
        public Input<int>? TaskCategory { get; set; }

        /// <summary>
        /// Task Id.
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// Task Type 1:Page Performance, 2:File upload,3:File Download,4:Port performance 5:Audio and video.
        /// </summary>
        [Input("taskType")]
        public Input<int>? TaskType { get; set; }

        public TaskSetState()
        {
        }
        public static new TaskSetState Empty => new TaskSetState();
    }
}
