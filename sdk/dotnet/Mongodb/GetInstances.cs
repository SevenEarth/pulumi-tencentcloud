// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mongodb
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of Mongodb instances.
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
        ///         var mongodb = Output.Create(Tencentcloud.Mongodb.GetInstances.InvokeAsync(new Tencentcloud.Mongodb.GetInstancesArgs
        ///         {
        ///             ClusterType = "REPLSET",
        ///             InstanceId = "cmgo-l6lwdsel",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:Mongodb/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of Mongodb instances.
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
        ///         var mongodb = Output.Create(Tencentcloud.Mongodb.GetInstances.InvokeAsync(new Tencentcloud.Mongodb.GetInstancesArgs
        ///         {
        ///             ClusterType = "REPLSET",
        ///             InstanceId = "cmgo-l6lwdsel",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:Mongodb/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of Mongodb cluster, and available values include replica set cluster(expressed with `REPLSET`), sharding cluster(expressed with `SHARD`).
        /// </summary>
        [Input("clusterType")]
        public string? ClusterType { get; set; }

        /// <summary>
        /// ID of the Mongodb instance to be queried.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// Name prefix of the Mongodb instance.
        /// </summary>
        [Input("instanceNamePrefix")]
        public string? InstanceNamePrefix { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// Tags of the Mongodb instance to be queried.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Type of Mongodb cluster, and available values include replica set cluster(expressed with `REPLSET`), sharding cluster(expressed with `SHARD`).
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// ID of the Mongodb instance to be queried.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name prefix of the Mongodb instance.
        /// </summary>
        [Input("instanceNamePrefix")]
        public Input<string>? InstanceNamePrefix { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the Mongodb instance to be queried.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// Type of Mongodb cluster.
        /// </summary>
        public readonly string? ClusterType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the Mongodb instance.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// A list of instances. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceListResult> InstanceLists;
        public readonly string? InstanceNamePrefix;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Tags of the Mongodb instance.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetInstancesResult(
            string? clusterType,

            string id,

            string? instanceId,

            ImmutableArray<Outputs.GetInstancesInstanceListResult> instanceLists,

            string? instanceNamePrefix,

            string? resultOutputFile,

            ImmutableDictionary<string, object>? tags)
        {
            ClusterType = clusterType;
            Id = id;
            InstanceId = instanceId;
            InstanceLists = instanceLists;
            InstanceNamePrefix = instanceNamePrefix;
            ResultOutputFile = resultOutputFile;
            Tags = tags;
        }
    }
}
