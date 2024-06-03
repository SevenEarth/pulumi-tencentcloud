// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a monitor tmp_tke_basic_config
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig")]
    public partial class TmpTkeBasicConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Type of cluster.
        /// </summary>
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        /// <summary>
        /// Full configuration in yaml format.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// config type, `service_monitors`, `pod_monitors`, `raw_jobs`.
        /// </summary>
        [Output("configType")]
        public Output<string> ConfigType { get; private set; } = null!;

        /// <summary>
        /// ID of instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Configure the name of the metric to keep on.
        /// </summary>
        [Output("metricsNames")]
        public Output<ImmutableArray<string>> MetricsNames { get; private set; } = null!;

        /// <summary>
        /// Name. The naming rule is: namespace/name. If you don&amp;#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeBasicConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeBasicConfig(string name, TmpTkeBasicConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig", name, args ?? new TmpTkeBasicConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeBasicConfig(string name, Input<string> id, TmpTkeBasicConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeBasicConfig:TmpTkeBasicConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeBasicConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeBasicConfig Get(string name, Input<string> id, TmpTkeBasicConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeBasicConfig(name, id, state, options);
        }
    }

    public sealed class TmpTkeBasicConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Type of cluster.
        /// </summary>
        [Input("clusterType", required: true)]
        public Input<string> ClusterType { get; set; } = null!;

        /// <summary>
        /// ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("metricsNames", required: true)]
        private InputList<string>? _metricsNames;

        /// <summary>
        /// Configure the name of the metric to keep on.
        /// </summary>
        public InputList<string> MetricsNames
        {
            get => _metricsNames ?? (_metricsNames = new InputList<string>());
            set => _metricsNames = value;
        }

        /// <summary>
        /// Name. The naming rule is: namespace/name. If you don&amp;#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TmpTkeBasicConfigArgs()
        {
        }
        public static new TmpTkeBasicConfigArgs Empty => new TmpTkeBasicConfigArgs();
    }

    public sealed class TmpTkeBasicConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Type of cluster.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// Full configuration in yaml format.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// config type, `service_monitors`, `pod_monitors`, `raw_jobs`.
        /// </summary>
        [Input("configType")]
        public Input<string>? ConfigType { get; set; }

        /// <summary>
        /// ID of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("metricsNames")]
        private InputList<string>? _metricsNames;

        /// <summary>
        /// Configure the name of the metric to keep on.
        /// </summary>
        public InputList<string> MetricsNames
        {
            get => _metricsNames ?? (_metricsNames = new InputList<string>());
            set => _metricsNames = value;
        }

        /// <summary>
        /// Name. The naming rule is: namespace/name. If you don&amp;#39;t have any namespace, use the default namespace: kube-system, otherwise use the specified one.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TmpTkeBasicConfigState()
        {
        }
        public static new TmpTkeBasicConfigState Empty => new TmpTkeBasicConfigState();
    }
}
