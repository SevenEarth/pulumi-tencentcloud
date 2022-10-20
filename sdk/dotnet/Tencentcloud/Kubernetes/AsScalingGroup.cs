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
    /// Auto scaling group for kubernetes cluster (offlined).
    /// 
    /// &gt; **NOTE:**  This resource was offline no longer suppored.
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/asScalingGroup:AsScalingGroup")]
    public partial class AsScalingGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Output("autoScalingConfig")]
        public Output<Outputs.AsScalingGroupAutoScalingConfig> AutoScalingConfig { get; private set; } = null!;

        /// <summary>
        /// Auto scaling group parameters.
        /// </summary>
        [Output("autoScalingGroup")]
        public Output<Outputs.AsScalingGroupAutoScalingGroup> AutoScalingGroup { get; private set; } = null!;

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        [Output("extraArgs")]
        public Output<ImmutableArray<string>> ExtraArgs { get; private set; } = null!;

        /// <summary>
        /// Labels of kubernetes AS Group created nodes.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Output("unschedulable")]
        public Output<int?> Unschedulable { get; private set; } = null!;


        /// <summary>
        /// Create a AsScalingGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AsScalingGroup(string name, AsScalingGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/asScalingGroup:AsScalingGroup", name, args ?? new AsScalingGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AsScalingGroup(string name, Input<string> id, AsScalingGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/asScalingGroup:AsScalingGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AsScalingGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AsScalingGroup Get(string name, Input<string> id, AsScalingGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AsScalingGroup(name, id, state, options);
        }
    }

    public sealed class AsScalingGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Input("autoScalingConfig", required: true)]
        public Input<Inputs.AsScalingGroupAutoScalingConfigArgs> AutoScalingConfig { get; set; } = null!;

        /// <summary>
        /// Auto scaling group parameters.
        /// </summary>
        [Input("autoScalingGroup", required: true)]
        public Input<Inputs.AsScalingGroupAutoScalingGroupArgs> AutoScalingGroup { get; set; } = null!;

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes AS Group created nodes.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        public AsScalingGroupArgs()
        {
        }
    }

    public sealed class AsScalingGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto scaling config parameters.
        /// </summary>
        [Input("autoScalingConfig")]
        public Input<Inputs.AsScalingGroupAutoScalingConfigGetArgs>? AutoScalingConfig { get; set; }

        /// <summary>
        /// Auto scaling group parameters.
        /// </summary>
        [Input("autoScalingGroup")]
        public Input<Inputs.AsScalingGroupAutoScalingGroupGetArgs>? AutoScalingGroup { get; set; }

        /// <summary>
        /// ID of the cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Custom parameter information related to the node.
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// Labels of kubernetes AS Group created nodes.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
        /// </summary>
        [Input("unschedulable")]
        public Input<int>? Unschedulable { get; set; }

        public AsScalingGroupState()
        {
        }
    }
}
