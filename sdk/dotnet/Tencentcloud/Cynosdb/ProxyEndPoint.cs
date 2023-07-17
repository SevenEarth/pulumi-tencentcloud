// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    /// <summary>
    /// Provides a resource to create a cynosdb proxy_end_point
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
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.108",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Open connection pool
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             ConnectionPoolTimeOut = 30,
    ///             ConnectionPoolType = "SessionConnectionPool",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             OpenConnectionPool = "yes",
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.108",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Close connection pool
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             OpenConnectionPool = "no",
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.108",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             ConsistencyTimeOut = 30,
    ///             ConsistencyType = "global",
    ///             FailOver = "yes",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             OpenConnectionPool = "no",
    ///             RwType = "READWRITE",
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.108",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-rikr6z4o",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             OpenConnectionPool = "no",
    ///             RwType = "READONLY",
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.108",
    ///             Vport = 3306,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Comprehensive parameter examples
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var proxyEndPoint = new Tencentcloud.Cynosdb.ProxyEndPoint("proxyEndPoint", new Tencentcloud.Cynosdb.ProxyEndPointArgs
    ///         {
    ///             AccessMode = "nearby",
    ///             AutoAddRo = "yes",
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             ConnectionPoolTimeOut = 30,
    ///             ConnectionPoolType = "SessionConnectionPool",
    ///             ConsistencyTimeOut = 30,
    ///             ConsistencyType = "global",
    ///             Description = "desc value",
    ///             FailOver = "yes",
    ///             InstanceWeights = 
    ///             {
    ///                 new Tencentcloud.Cynosdb.Inputs.ProxyEndPointInstanceWeightArgs
    ///                 {
    ///                     InstanceId = "cynosdbmysql-ins-afqx1hy0",
    ///                     Weight = 1,
    ///                 },
    ///             },
    ///             OpenConnectionPool = "yes",
    ///             RwType = "READWRITE",
    ///             SecurityGroupIds = 
    ///             {
    ///                 "sg-7kpsbxdb",
    ///             },
    ///             TransSplit = true,
    ///             UniqueSubnetId = "subnet-dwj7ipnc",
    ///             UniqueVpcId = "vpc-4owdpnwr",
    ///             Vip = "172.16.112.118",
    ///             Vport = 3306,
    ///             WeightMode = "system",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/proxyEndPoint:ProxyEndPoint")]
    public partial class ProxyEndPoint : Pulumi.CustomResource
    {
        /// <summary>
        /// Connection mode: nearby, balance.
        /// </summary>
        [Output("accessMode")]
        public Output<string> AccessMode { get; private set; } = null!;

        /// <summary>
        /// Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
        /// </summary>
        [Output("autoAddRo")]
        public Output<string> AutoAddRo { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Connection pool threshold: unit (second).
        /// </summary>
        [Output("connectionPoolTimeOut")]
        public Output<int> ConnectionPoolTimeOut { get; private set; } = null!;

        /// <summary>
        /// Connection pool type: SessionConnectionPool (session level Connection pool).
        /// </summary>
        [Output("connectionPoolType")]
        public Output<string> ConnectionPoolType { get; private set; } = null!;

        /// <summary>
        /// Consistency timeout.
        /// </summary>
        [Output("consistencyTimeOut")]
        public Output<int> ConsistencyTimeOut { get; private set; } = null!;

        /// <summary>
        /// Consistency type: event, global, session.
        /// </summary>
        [Output("consistencyType")]
        public Output<string> ConsistencyType { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enable Failover. yes or no.
        /// </summary>
        [Output("failOver")]
        public Output<string> FailOver { get; private set; } = null!;

        /// <summary>
        /// Instance Group ID.
        /// </summary>
        [Output("instanceGroupId")]
        public Output<string> InstanceGroupId { get; private set; } = null!;

        /// <summary>
        /// Instance Weight.
        /// </summary>
        [Output("instanceWeights")]
        public Output<ImmutableArray<Outputs.ProxyEndPointInstanceWeight>> InstanceWeights { get; private set; } = null!;

        /// <summary>
        /// Whether to enable Connection pool, yes - enable, no - do not enable.
        /// </summary>
        [Output("openConnectionPool")]
        public Output<string> OpenConnectionPool { get; private set; } = null!;

        /// <summary>
        /// Proxy Group ID.
        /// </summary>
        [Output("proxyGroupId")]
        public Output<string> ProxyGroupId { get; private set; } = null!;

        /// <summary>
        /// Read and write attributes: READWRITE, READONLY.
        /// </summary>
        [Output("rwType")]
        public Output<string> RwType { get; private set; } = null!;

        /// <summary>
        /// Security Group ID Array.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Transaction splitting.
        /// </summary>
        [Output("transSplit")]
        public Output<bool> TransSplit { get; private set; } = null!;

        /// <summary>
        /// The private network subnet ID is consistent with the cluster subnet ID by default.
        /// </summary>
        [Output("uniqueSubnetId")]
        public Output<string> UniqueSubnetId { get; private set; } = null!;

        /// <summary>
        /// Private network ID, which is consistent with the cluster private network ID by default.
        /// </summary>
        [Output("uniqueVpcId")]
        public Output<string> UniqueVpcId { get; private set; } = null!;

        /// <summary>
        /// VIP Information.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// Port Information.
        /// </summary>
        [Output("vport")]
        public Output<int> Vport { get; private set; } = null!;

        /// <summary>
        /// Weight mode: system system allocation, custom customization.
        /// </summary>
        [Output("weightMode")]
        public Output<string> WeightMode { get; private set; } = null!;


        /// <summary>
        /// Create a ProxyEndPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProxyEndPoint(string name, ProxyEndPointArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/proxyEndPoint:ProxyEndPoint", name, args ?? new ProxyEndPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProxyEndPoint(string name, Input<string> id, ProxyEndPointState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/proxyEndPoint:ProxyEndPoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProxyEndPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProxyEndPoint Get(string name, Input<string> id, ProxyEndPointState? state = null, CustomResourceOptions? options = null)
        {
            return new ProxyEndPoint(name, id, state, options);
        }
    }

    public sealed class ProxyEndPointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection mode: nearby, balance.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

        /// <summary>
        /// Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
        /// </summary>
        [Input("autoAddRo")]
        public Input<string>? AutoAddRo { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Connection pool threshold: unit (second).
        /// </summary>
        [Input("connectionPoolTimeOut")]
        public Input<int>? ConnectionPoolTimeOut { get; set; }

        /// <summary>
        /// Connection pool type: SessionConnectionPool (session level Connection pool).
        /// </summary>
        [Input("connectionPoolType")]
        public Input<string>? ConnectionPoolType { get; set; }

        /// <summary>
        /// Consistency timeout.
        /// </summary>
        [Input("consistencyTimeOut")]
        public Input<int>? ConsistencyTimeOut { get; set; }

        /// <summary>
        /// Consistency type: event, global, session.
        /// </summary>
        [Input("consistencyType")]
        public Input<string>? ConsistencyType { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable Failover. yes or no.
        /// </summary>
        [Input("failOver")]
        public Input<string>? FailOver { get; set; }

        [Input("instanceWeights")]
        private InputList<Inputs.ProxyEndPointInstanceWeightArgs>? _instanceWeights;

        /// <summary>
        /// Instance Weight.
        /// </summary>
        public InputList<Inputs.ProxyEndPointInstanceWeightArgs> InstanceWeights
        {
            get => _instanceWeights ?? (_instanceWeights = new InputList<Inputs.ProxyEndPointInstanceWeightArgs>());
            set => _instanceWeights = value;
        }

        /// <summary>
        /// Whether to enable Connection pool, yes - enable, no - do not enable.
        /// </summary>
        [Input("openConnectionPool")]
        public Input<string>? OpenConnectionPool { get; set; }

        /// <summary>
        /// Read and write attributes: READWRITE, READONLY.
        /// </summary>
        [Input("rwType")]
        public Input<string>? RwType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security Group ID Array.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Transaction splitting.
        /// </summary>
        [Input("transSplit")]
        public Input<bool>? TransSplit { get; set; }

        /// <summary>
        /// The private network subnet ID is consistent with the cluster subnet ID by default.
        /// </summary>
        [Input("uniqueSubnetId", required: true)]
        public Input<string> UniqueSubnetId { get; set; } = null!;

        /// <summary>
        /// Private network ID, which is consistent with the cluster private network ID by default.
        /// </summary>
        [Input("uniqueVpcId", required: true)]
        public Input<string> UniqueVpcId { get; set; } = null!;

        /// <summary>
        /// VIP Information.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Port Information.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        /// <summary>
        /// Weight mode: system system allocation, custom customization.
        /// </summary>
        [Input("weightMode")]
        public Input<string>? WeightMode { get; set; }

        public ProxyEndPointArgs()
        {
        }
    }

    public sealed class ProxyEndPointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection mode: nearby, balance.
        /// </summary>
        [Input("accessMode")]
        public Input<string>? AccessMode { get; set; }

        /// <summary>
        /// Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
        /// </summary>
        [Input("autoAddRo")]
        public Input<string>? AutoAddRo { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Connection pool threshold: unit (second).
        /// </summary>
        [Input("connectionPoolTimeOut")]
        public Input<int>? ConnectionPoolTimeOut { get; set; }

        /// <summary>
        /// Connection pool type: SessionConnectionPool (session level Connection pool).
        /// </summary>
        [Input("connectionPoolType")]
        public Input<string>? ConnectionPoolType { get; set; }

        /// <summary>
        /// Consistency timeout.
        /// </summary>
        [Input("consistencyTimeOut")]
        public Input<int>? ConsistencyTimeOut { get; set; }

        /// <summary>
        /// Consistency type: event, global, session.
        /// </summary>
        [Input("consistencyType")]
        public Input<string>? ConsistencyType { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enable Failover. yes or no.
        /// </summary>
        [Input("failOver")]
        public Input<string>? FailOver { get; set; }

        /// <summary>
        /// Instance Group ID.
        /// </summary>
        [Input("instanceGroupId")]
        public Input<string>? InstanceGroupId { get; set; }

        [Input("instanceWeights")]
        private InputList<Inputs.ProxyEndPointInstanceWeightGetArgs>? _instanceWeights;

        /// <summary>
        /// Instance Weight.
        /// </summary>
        public InputList<Inputs.ProxyEndPointInstanceWeightGetArgs> InstanceWeights
        {
            get => _instanceWeights ?? (_instanceWeights = new InputList<Inputs.ProxyEndPointInstanceWeightGetArgs>());
            set => _instanceWeights = value;
        }

        /// <summary>
        /// Whether to enable Connection pool, yes - enable, no - do not enable.
        /// </summary>
        [Input("openConnectionPool")]
        public Input<string>? OpenConnectionPool { get; set; }

        /// <summary>
        /// Proxy Group ID.
        /// </summary>
        [Input("proxyGroupId")]
        public Input<string>? ProxyGroupId { get; set; }

        /// <summary>
        /// Read and write attributes: READWRITE, READONLY.
        /// </summary>
        [Input("rwType")]
        public Input<string>? RwType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// Security Group ID Array.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Transaction splitting.
        /// </summary>
        [Input("transSplit")]
        public Input<bool>? TransSplit { get; set; }

        /// <summary>
        /// The private network subnet ID is consistent with the cluster subnet ID by default.
        /// </summary>
        [Input("uniqueSubnetId")]
        public Input<string>? UniqueSubnetId { get; set; }

        /// <summary>
        /// Private network ID, which is consistent with the cluster private network ID by default.
        /// </summary>
        [Input("uniqueVpcId")]
        public Input<string>? UniqueVpcId { get; set; }

        /// <summary>
        /// VIP Information.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// Port Information.
        /// </summary>
        [Input("vport")]
        public Input<int>? Vport { get; set; }

        /// <summary>
        /// Weight mode: system system allocation, custom customization.
        /// </summary>
        [Input("weightMode")]
        public Input<string>? WeightMode { get; set; }

        public ProxyEndPointState()
        {
        }
    }
}
