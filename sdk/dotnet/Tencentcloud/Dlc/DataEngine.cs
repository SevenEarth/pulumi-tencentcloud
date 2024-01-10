// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc
{
    /// <summary>
    /// Provides a resource to create a dlc data_engine
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
    ///         var dataEngine = new Tencentcloud.Dlc.DataEngine("dataEngine", new Tencentcloud.Dlc.DataEngineArgs
    ///         {
    ///             AutoResume = false,
    ///             AutoSuspend = false,
    ///             CidrBlock = "10.255.0.0/16",
    ///             ClusterType = "spark_cu",
    ///             CrontabResumeSuspend = 0,
    ///             DataEngineName = "testSpark",
    ///             DefaultDataEngine = false,
    ///             EngineExecType = "BATCH",
    ///             EngineType = "spark",
    ///             MaxClusters = 1,
    ///             Message = "test spark1",
    ///             MinClusters = 1,
    ///             Mode = 1,
    ///             PayMode = 0,
    ///             Size = 16,
    ///             TimeSpan = 1,
    ///             TimeUnit = "h",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dlc data_engine can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dlc/dataEngine:DataEngine data_engine data_engine_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/dataEngine:DataEngine")]
    public partial class DataEngine : Pulumi.CustomResource
    {
        /// <summary>
        /// Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
        /// </summary>
        [Output("autoRenew")]
        public Output<int?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically start the cluster, prepay not support.
        /// </summary>
        [Output("autoResume")]
        public Output<bool> AutoResume { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically suspend the cluster, prepay not support.
        /// </summary>
        [Output("autoSuspend")]
        public Output<bool?> AutoSuspend { get; private set; } = null!;

        /// <summary>
        /// Cluster automatic suspension time, default 10 minutes.
        /// </summary>
        [Output("autoSuspendTime")]
        public Output<int> AutoSuspendTime { get; private set; } = null!;

        /// <summary>
        /// Engine VPC network segment, just like 192.0.2.1/24.
        /// </summary>
        [Output("cidrBlock")]
        public Output<string?> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Engine cluster type, only support: spark_cu/presto_cu.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
        /// </summary>
        [Output("crontabResumeSuspend")]
        public Output<int?> CrontabResumeSuspend { get; private set; } = null!;

        /// <summary>
        /// Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
        /// </summary>
        [Output("crontabResumeSuspendStrategy")]
        public Output<Outputs.DataEngineCrontabResumeSuspendStrategy> CrontabResumeSuspendStrategy { get; private set; } = null!;

        /// <summary>
        /// Cluster advanced configuration.
        /// </summary>
        [Output("dataEngineConfigPairs")]
        public Output<ImmutableArray<Outputs.DataEngineDataEngineConfigPair>> DataEngineConfigPairs { get; private set; } = null!;

        /// <summary>
        /// Engine name.
        /// </summary>
        [Output("dataEngineName")]
        public Output<string> DataEngineName { get; private set; } = null!;

        /// <summary>
        /// Whether it is the default virtual cluster.
        /// </summary>
        [Output("defaultDataEngine")]
        public Output<bool?> DefaultDataEngine { get; private set; } = null!;

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster elastic limit.
        /// </summary>
        [Output("elasticLimit")]
        public Output<int?> ElasticLimit { get; private set; } = null!;

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
        /// </summary>
        [Output("elasticSwitch")]
        public Output<bool?> ElasticSwitch { get; private set; } = null!;

        /// <summary>
        /// Engine exec type, only support SQL(default) or BATCH.
        /// </summary>
        [Output("engineExecType")]
        public Output<string?> EngineExecType { get; private set; } = null!;

        /// <summary>
        /// Engine type, only support: spark/presto.
        /// </summary>
        [Output("engineType")]
        public Output<string> EngineType { get; private set; } = null!;

        /// <summary>
        /// Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
        /// </summary>
        [Output("imageVersionName")]
        public Output<string> ImageVersionName { get; private set; } = null!;

        /// <summary>
        /// Primary cluster name, specified when creating a disaster recovery cluster.
        /// </summary>
        [Output("mainClusterName")]
        public Output<string?> MainClusterName { get; private set; } = null!;

        /// <summary>
        /// Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Output("maxClusters")]
        public Output<int?> MaxClusters { get; private set; } = null!;

        /// <summary>
        /// Maximum number of concurrent tasks in a single cluster, default 5.
        /// </summary>
        [Output("maxConcurrency")]
        public Output<int> MaxConcurrency { get; private set; } = null!;

        /// <summary>
        /// Engine description information.
        /// </summary>
        [Output("message")]
        public Output<string?> Message { get; private set; } = null!;

        /// <summary>
        /// Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Output("minClusters")]
        public Output<int?> MinClusters { get; private set; } = null!;

        /// <summary>
        /// Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
        /// </summary>
        [Output("mode")]
        public Output<int> Mode { get; private set; } = null!;

        /// <summary>
        /// Engine pay mode type, only support 0: postPay, 1: prePay(default).
        /// </summary>
        [Output("payMode")]
        public Output<int?> PayMode { get; private set; } = null!;

        /// <summary>
        /// Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// For spark Batch ExecType, cluster session resource configuration template.
        /// </summary>
        [Output("sessionResourceTemplate")]
        public Output<Outputs.DataEngineSessionResourceTemplate> SessionResourceTemplate { get; private set; } = null!;

        /// <summary>
        /// Cluster size. Required when updating.
        /// </summary>
        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;

        /// <summary>
        /// Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
        /// </summary>
        [Output("timeSpan")]
        public Output<int?> TimeSpan { get; private set; } = null!;

        /// <summary>
        /// Engine TimeUnit, prePay: use m(default), postPay: use h.
        /// </summary>
        [Output("timeUnit")]
        public Output<string?> TimeUnit { get; private set; } = null!;

        /// <summary>
        /// Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
        /// </summary>
        [Output("tolerableQueueTime")]
        public Output<int?> TolerableQueueTime { get; private set; } = null!;


        /// <summary>
        /// Create a DataEngine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataEngine(string name, DataEngineArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/dataEngine:DataEngine", name, args ?? new DataEngineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataEngine(string name, Input<string> id, DataEngineState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/dataEngine:DataEngine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataEngine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataEngine Get(string name, Input<string> id, DataEngineState? state = null, CustomResourceOptions? options = null)
        {
            return new DataEngine(name, id, state, options);
        }
    }

    public sealed class DataEngineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
        /// </summary>
        [Input("autoRenew")]
        public Input<int>? AutoRenew { get; set; }

        /// <summary>
        /// Whether to automatically start the cluster, prepay not support.
        /// </summary>
        [Input("autoResume", required: true)]
        public Input<bool> AutoResume { get; set; } = null!;

        /// <summary>
        /// Whether to automatically suspend the cluster, prepay not support.
        /// </summary>
        [Input("autoSuspend")]
        public Input<bool>? AutoSuspend { get; set; }

        /// <summary>
        /// Cluster automatic suspension time, default 10 minutes.
        /// </summary>
        [Input("autoSuspendTime")]
        public Input<int>? AutoSuspendTime { get; set; }

        /// <summary>
        /// Engine VPC network segment, just like 192.0.2.1/24.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Engine cluster type, only support: spark_cu/presto_cu.
        /// </summary>
        [Input("clusterType", required: true)]
        public Input<string> ClusterType { get; set; } = null!;

        /// <summary>
        /// Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
        /// </summary>
        [Input("crontabResumeSuspend")]
        public Input<int>? CrontabResumeSuspend { get; set; }

        /// <summary>
        /// Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
        /// </summary>
        [Input("crontabResumeSuspendStrategy")]
        public Input<Inputs.DataEngineCrontabResumeSuspendStrategyArgs>? CrontabResumeSuspendStrategy { get; set; }

        [Input("dataEngineConfigPairs")]
        private InputList<Inputs.DataEngineDataEngineConfigPairArgs>? _dataEngineConfigPairs;

        /// <summary>
        /// Cluster advanced configuration.
        /// </summary>
        public InputList<Inputs.DataEngineDataEngineConfigPairArgs> DataEngineConfigPairs
        {
            get => _dataEngineConfigPairs ?? (_dataEngineConfigPairs = new InputList<Inputs.DataEngineDataEngineConfigPairArgs>());
            set => _dataEngineConfigPairs = value;
        }

        /// <summary>
        /// Engine name.
        /// </summary>
        [Input("dataEngineName", required: true)]
        public Input<string> DataEngineName { get; set; } = null!;

        /// <summary>
        /// Whether it is the default virtual cluster.
        /// </summary>
        [Input("defaultDataEngine")]
        public Input<bool>? DefaultDataEngine { get; set; }

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster elastic limit.
        /// </summary>
        [Input("elasticLimit")]
        public Input<int>? ElasticLimit { get; set; }

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
        /// </summary>
        [Input("elasticSwitch")]
        public Input<bool>? ElasticSwitch { get; set; }

        /// <summary>
        /// Engine exec type, only support SQL(default) or BATCH.
        /// </summary>
        [Input("engineExecType")]
        public Input<string>? EngineExecType { get; set; }

        /// <summary>
        /// Engine type, only support: spark/presto.
        /// </summary>
        [Input("engineType", required: true)]
        public Input<string> EngineType { get; set; } = null!;

        /// <summary>
        /// Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
        /// </summary>
        [Input("imageVersionName")]
        public Input<string>? ImageVersionName { get; set; }

        /// <summary>
        /// Primary cluster name, specified when creating a disaster recovery cluster.
        /// </summary>
        [Input("mainClusterName")]
        public Input<string>? MainClusterName { get; set; }

        /// <summary>
        /// Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Input("maxClusters")]
        public Input<int>? MaxClusters { get; set; }

        /// <summary>
        /// Maximum number of concurrent tasks in a single cluster, default 5.
        /// </summary>
        [Input("maxConcurrency")]
        public Input<int>? MaxConcurrency { get; set; }

        /// <summary>
        /// Engine description information.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Input("minClusters")]
        public Input<int>? MinClusters { get; set; }

        /// <summary>
        /// Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
        /// </summary>
        [Input("mode", required: true)]
        public Input<int> Mode { get; set; } = null!;

        /// <summary>
        /// Engine pay mode type, only support 0: postPay, 1: prePay(default).
        /// </summary>
        [Input("payMode")]
        public Input<int>? PayMode { get; set; }

        /// <summary>
        /// Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// For spark Batch ExecType, cluster session resource configuration template.
        /// </summary>
        [Input("sessionResourceTemplate")]
        public Input<Inputs.DataEngineSessionResourceTemplateArgs>? SessionResourceTemplate { get; set; }

        /// <summary>
        /// Cluster size. Required when updating.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
        /// </summary>
        [Input("timeSpan")]
        public Input<int>? TimeSpan { get; set; }

        /// <summary>
        /// Engine TimeUnit, prePay: use m(default), postPay: use h.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        /// <summary>
        /// Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
        /// </summary>
        [Input("tolerableQueueTime")]
        public Input<int>? TolerableQueueTime { get; set; }

        public DataEngineArgs()
        {
        }
    }

    public sealed class DataEngineState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
        /// </summary>
        [Input("autoRenew")]
        public Input<int>? AutoRenew { get; set; }

        /// <summary>
        /// Whether to automatically start the cluster, prepay not support.
        /// </summary>
        [Input("autoResume")]
        public Input<bool>? AutoResume { get; set; }

        /// <summary>
        /// Whether to automatically suspend the cluster, prepay not support.
        /// </summary>
        [Input("autoSuspend")]
        public Input<bool>? AutoSuspend { get; set; }

        /// <summary>
        /// Cluster automatic suspension time, default 10 minutes.
        /// </summary>
        [Input("autoSuspendTime")]
        public Input<int>? AutoSuspendTime { get; set; }

        /// <summary>
        /// Engine VPC network segment, just like 192.0.2.1/24.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Engine cluster type, only support: spark_cu/presto_cu.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
        /// </summary>
        [Input("crontabResumeSuspend")]
        public Input<int>? CrontabResumeSuspend { get; set; }

        /// <summary>
        /// Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
        /// </summary>
        [Input("crontabResumeSuspendStrategy")]
        public Input<Inputs.DataEngineCrontabResumeSuspendStrategyGetArgs>? CrontabResumeSuspendStrategy { get; set; }

        [Input("dataEngineConfigPairs")]
        private InputList<Inputs.DataEngineDataEngineConfigPairGetArgs>? _dataEngineConfigPairs;

        /// <summary>
        /// Cluster advanced configuration.
        /// </summary>
        public InputList<Inputs.DataEngineDataEngineConfigPairGetArgs> DataEngineConfigPairs
        {
            get => _dataEngineConfigPairs ?? (_dataEngineConfigPairs = new InputList<Inputs.DataEngineDataEngineConfigPairGetArgs>());
            set => _dataEngineConfigPairs = value;
        }

        /// <summary>
        /// Engine name.
        /// </summary>
        [Input("dataEngineName")]
        public Input<string>? DataEngineName { get; set; }

        /// <summary>
        /// Whether it is the default virtual cluster.
        /// </summary>
        [Input("defaultDataEngine")]
        public Input<bool>? DefaultDataEngine { get; set; }

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster elastic limit.
        /// </summary>
        [Input("elasticLimit")]
        public Input<int>? ElasticLimit { get; set; }

        /// <summary>
        /// For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
        /// </summary>
        [Input("elasticSwitch")]
        public Input<bool>? ElasticSwitch { get; set; }

        /// <summary>
        /// Engine exec type, only support SQL(default) or BATCH.
        /// </summary>
        [Input("engineExecType")]
        public Input<string>? EngineExecType { get; set; }

        /// <summary>
        /// Engine type, only support: spark/presto.
        /// </summary>
        [Input("engineType")]
        public Input<string>? EngineType { get; set; }

        /// <summary>
        /// Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
        /// </summary>
        [Input("imageVersionName")]
        public Input<string>? ImageVersionName { get; set; }

        /// <summary>
        /// Primary cluster name, specified when creating a disaster recovery cluster.
        /// </summary>
        [Input("mainClusterName")]
        public Input<string>? MainClusterName { get; set; }

        /// <summary>
        /// Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Input("maxClusters")]
        public Input<int>? MaxClusters { get; set; }

        /// <summary>
        /// Maximum number of concurrent tasks in a single cluster, default 5.
        /// </summary>
        [Input("maxConcurrency")]
        public Input<int>? MaxConcurrency { get; set; }

        /// <summary>
        /// Engine description information.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
        /// </summary>
        [Input("minClusters")]
        public Input<int>? MinClusters { get; set; }

        /// <summary>
        /// Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
        /// </summary>
        [Input("mode")]
        public Input<int>? Mode { get; set; }

        /// <summary>
        /// Engine pay mode type, only support 0: postPay, 1: prePay(default).
        /// </summary>
        [Input("payMode")]
        public Input<int>? PayMode { get; set; }

        /// <summary>
        /// Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// For spark Batch ExecType, cluster session resource configuration template.
        /// </summary>
        [Input("sessionResourceTemplate")]
        public Input<Inputs.DataEngineSessionResourceTemplateGetArgs>? SessionResourceTemplate { get; set; }

        /// <summary>
        /// Cluster size. Required when updating.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
        /// </summary>
        [Input("timeSpan")]
        public Input<int>? TimeSpan { get; set; }

        /// <summary>
        /// Engine TimeUnit, prePay: use m(default), postPay: use h.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        /// <summary>
        /// Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
        /// </summary>
        [Input("tolerableQueueTime")]
        public Input<int>? TolerableQueueTime { get; set; }

        public DataEngineState()
        {
        }
    }
}
