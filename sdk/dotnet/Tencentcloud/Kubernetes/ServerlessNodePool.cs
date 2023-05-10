// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes
{
    /// <summary>
    /// Provide a resource to create serverless node pool of cluster.
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
    ///         var exampleServerlessNodePool = new Tencentcloud.Kubernetes.ServerlessNodePool("exampleServerlessNodePool", new Tencentcloud.Kubernetes.ServerlessNodePoolArgs
    ///         {
    ///             ClusterId = tencentcloud_kubernetes_cluster.Example.Id,
    ///             ServerlessNodes = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.ServerlessNodePoolServerlessNodeArgs
    ///                 {
    ///                     DisplayName = "serverless_node1",
    ///                     SubnetId = "subnet-xxx",
    ///                 },
    ///                 new Tencentcloud.Kubernetes.Inputs.ServerlessNodePoolServerlessNodeArgs
    ///                 {
    ///                     DisplayName = "serverless_node2",
    ///                     SubnetId = "subnet-xxx",
    ///                 },
    ///             },
    ///             SecurityGroupIds = 
    ///             {
    ///                 "sg-xxx",
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "example1", "test1" },
    ///                 { "example2", "test2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// serverless node pool can be imported, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool test cls-xxx#np-xxx
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool")]
    public partial class ServerlessNodePool : Pulumi.CustomResource
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// life state of serverless node pool.
        /// </summary>
        [Output("lifeState")]
        public Output<string> LifeState { get; private set; } = null!;

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        [Output("serverlessNodes")]
        public Output<ImmutableArray<Outputs.ServerlessNodePoolServerlessNode>> ServerlessNodes { get; private set; } = null!;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        [Output("taints")]
        public Output<ImmutableArray<Outputs.ServerlessNodePoolTaint>> Taints { get; private set; } = null!;


        /// <summary>
        /// Create a ServerlessNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerlessNodePool(string name, ServerlessNodePoolArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, args ?? new ServerlessNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerlessNodePool(string name, Input<string> id, ServerlessNodePoolState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerlessNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerlessNodePool Get(string name, Input<string> id, ServerlessNodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerlessNodePool(name, id, state, options);
        }
    }

    public sealed class ServerlessNodePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("serverlessNodes", required: true)]
        private InputList<Inputs.ServerlessNodePoolServerlessNodeArgs>? _serverlessNodes;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolServerlessNodeArgs> ServerlessNodes
        {
            get => _serverlessNodes ?? (_serverlessNodes = new InputList<Inputs.ServerlessNodePoolServerlessNodeArgs>());
            set => _serverlessNodes = value;
        }

        [Input("taints")]
        private InputList<Inputs.ServerlessNodePoolTaintArgs>? _taints;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolTaintArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.ServerlessNodePoolTaintArgs>());
            set => _taints = value;
        }

        public ServerlessNodePoolArgs()
        {
        }
    }

    public sealed class ServerlessNodePoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// cluster id of serverless node pool.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// labels of serverless node.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// life state of serverless node pool.
        /// </summary>
        [Input("lifeState")]
        public Input<string>? LifeState { get; set; }

        /// <summary>
        /// serverless node pool name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups of serverless node pool.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("serverlessNodes")]
        private InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs>? _serverlessNodes;

        /// <summary>
        /// node list of serverless node pool.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs> ServerlessNodes
        {
            get => _serverlessNodes ?? (_serverlessNodes = new InputList<Inputs.ServerlessNodePoolServerlessNodeGetArgs>());
            set => _serverlessNodes = value;
        }

        [Input("taints")]
        private InputList<Inputs.ServerlessNodePoolTaintGetArgs>? _taints;

        /// <summary>
        /// taints of serverless node.
        /// </summary>
        public InputList<Inputs.ServerlessNodePoolTaintGetArgs> Taints
        {
            get => _taints ?? (_taints = new InputList<Inputs.ServerlessNodePoolTaintGetArgs>());
            set => _taints = value;
        }

        public ServerlessNodePoolState()
        {
        }
    }
}
