// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    public static class GetSlowLogs
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_logs
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
        ///         var slowLogs = Output.Create(Tencentcloud.Dbbrain.GetSlowLogs.InvokeAsync(new Tencentcloud.Dbbrain.GetSlowLogsArgs
        ///         {
        ///             EndTime = "%s",
        ///             InstanceId = "%s",
        ///             Md5 = "4961208426639258265",
        ///             Product = "mysql",
        ///             StartTime = "%s",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSlowLogsResult> InvokeAsync(GetSlowLogsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSlowLogsResult>("tencentcloud:Dbbrain/getSlowLogs:getSlowLogs", args ?? new GetSlowLogsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain slow_logs
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
        ///         var slowLogs = Output.Create(Tencentcloud.Dbbrain.GetSlowLogs.InvokeAsync(new Tencentcloud.Dbbrain.GetSlowLogsArgs
        ///         {
        ///             EndTime = "%s",
        ///             InstanceId = "%s",
        ///             Md5 = "4961208426639258265",
        ///             Product = "mysql",
        ///             StartTime = "%s",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSlowLogsResult> Invoke(GetSlowLogsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSlowLogsResult>("tencentcloud:Dbbrain/getSlowLogs:getSlowLogs", args ?? new GetSlowLogsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSlowLogsArgs : Pulumi.InvokeArgs
    {
        [Input("dbs")]
        private List<string>? _dbs;

        /// <summary>
        /// database list.
        /// </summary>
        public List<string> Dbs
        {
            get => _dbs ?? (_dbs = new List<string>());
            set => _dbs = value;
        }

        /// <summary>
        /// The deadline, such as 2019-09-11 10:13:14, the interval between the deadline and the start time is less than 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        /// <summary>
        /// instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("ips")]
        private List<string>? _ips;

        /// <summary>
        /// ip.
        /// </summary>
        public List<string> Ips
        {
            get => _ips ?? (_ips = new List<string>());
            set => _ips = value;
        }

        [Input("keys")]
        private List<string>? _keys;

        /// <summary>
        /// keywords.
        /// </summary>
        public List<string> Keys
        {
            get => _keys ?? (_keys = new List<string>());
            set => _keys = value;
        }

        /// <summary>
        /// md5 value of sql template.
        /// </summary>
        [Input("md5", required: true)]
        public string Md5 { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, such as 2019-09-10 12:13:14.
        /// </summary>
        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        [Input("times")]
        private List<int>? _times;

        /// <summary>
        /// Time-consuming interval, the left and right boundaries of the time-consuming interval correspond to the 0th element and the first element of the array respectively.
        /// </summary>
        public List<int> Times
        {
            get => _times ?? (_times = new List<int>());
            set => _times = value;
        }

        [Input("users")]
        private List<string>? _users;

        /// <summary>
        /// user.
        /// </summary>
        public List<string> Users
        {
            get => _users ?? (_users = new List<string>());
            set => _users = value;
        }

        public GetSlowLogsArgs()
        {
        }
    }

    public sealed class GetSlowLogsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dbs")]
        private InputList<string>? _dbs;

        /// <summary>
        /// database list.
        /// </summary>
        public InputList<string> Dbs
        {
            get => _dbs ?? (_dbs = new InputList<string>());
            set => _dbs = value;
        }

        /// <summary>
        /// The deadline, such as 2019-09-11 10:13:14, the interval between the deadline and the start time is less than 7 days.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        /// <summary>
        /// instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("ips")]
        private InputList<string>? _ips;

        /// <summary>
        /// ip.
        /// </summary>
        public InputList<string> Ips
        {
            get => _ips ?? (_ips = new InputList<string>());
            set => _ips = value;
        }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// keywords.
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// md5 value of sql template.
        /// </summary>
        [Input("md5", required: true)]
        public Input<string> Md5 { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time, such as 2019-09-10 12:13:14.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        [Input("times")]
        private InputList<int>? _times;

        /// <summary>
        /// Time-consuming interval, the left and right boundaries of the time-consuming interval correspond to the 0th element and the first element of the array respectively.
        /// </summary>
        public InputList<int> Times
        {
            get => _times ?? (_times = new InputList<int>());
            set => _times = value;
        }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// user.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public GetSlowLogsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSlowLogsResult
    {
        public readonly ImmutableArray<string> Dbs;
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<string> Ips;
        public readonly ImmutableArray<string> Keys;
        public readonly string Md5;
        public readonly string Product;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Slow log details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSlowLogsRowResult> Rows;
        public readonly string StartTime;
        public readonly ImmutableArray<int> Times;
        public readonly ImmutableArray<string> Users;

        [OutputConstructor]
        private GetSlowLogsResult(
            ImmutableArray<string> dbs,

            string endTime,

            string id,

            string instanceId,

            ImmutableArray<string> ips,

            ImmutableArray<string> keys,

            string md5,

            string product,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetSlowLogsRowResult> rows,

            string startTime,

            ImmutableArray<int> times,

            ImmutableArray<string> users)
        {
            Dbs = dbs;
            EndTime = endTime;
            Id = id;
            InstanceId = instanceId;
            Ips = ips;
            Keys = keys;
            Md5 = md5;
            Product = product;
            ResultOutputFile = resultOutputFile;
            Rows = rows;
            StartTime = startTime;
            Times = times;
            Users = users;
        }
    }
}
