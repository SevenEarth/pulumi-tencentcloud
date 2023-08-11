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
    /// Provides a resource to create a monitor grafanaPlugin
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
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-6";
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///             AvailabilityZone = availabilityZone,
    ///             CidrBlock = "10.0.1.0/24",
    ///         });
    ///         var foo = new Tencentcloud.Monitor.GrafanaInstance("foo", new Tencentcloud.Monitor.GrafanaInstanceArgs
    ///         {
    ///             InstanceName = "test-grafana",
    ///             VpcId = vpc.Id,
    ///             SubnetIds = 
    ///             {
    ///                 subnet.Id,
    ///             },
    ///             GrafanaInitPassword = "1234567890",
    ///             EnableInternet = false,
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "test" },
    ///             },
    ///         });
    ///         var grafanaPlugin = new Tencentcloud.Monitor.GrafanaPlugin("grafanaPlugin", new Tencentcloud.Monitor.GrafanaPluginArgs
    ///         {
    ///             InstanceId = foo.Id,
    ///             PluginId = "grafana-piechart-panel",
    ///             Version = "1.6.2",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// monitor grafanaPlugin can be imported using the instance_id#plugin_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Monitor/grafanaPlugin:GrafanaPlugin grafanaPlugin grafana-50nj6v00#grafana-piechart-panel
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/grafanaPlugin:GrafanaPlugin")]
    public partial class GrafanaPlugin : Pulumi.CustomResource
    {
        /// <summary>
        /// Grafana instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Plugin id.
        /// </summary>
        [Output("pluginId")]
        public Output<string> PluginId { get; private set; } = null!;

        /// <summary>
        /// Plugin version.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaPlugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaPlugin(string name, GrafanaPluginArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaPlugin:GrafanaPlugin", name, args ?? new GrafanaPluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaPlugin(string name, Input<string> id, GrafanaPluginState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaPlugin:GrafanaPlugin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaPlugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaPlugin Get(string name, Input<string> id, GrafanaPluginState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaPlugin(name, id, state, options);
        }
    }

    public sealed class GrafanaPluginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Grafana instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Plugin id.
        /// </summary>
        [Input("pluginId", required: true)]
        public Input<string> PluginId { get; set; } = null!;

        /// <summary>
        /// Plugin version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GrafanaPluginArgs()
        {
        }
    }

    public sealed class GrafanaPluginState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Grafana instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Plugin id.
        /// </summary>
        [Input("pluginId")]
        public Input<string>? PluginId { get; set; }

        /// <summary>
        /// Plugin version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GrafanaPluginState()
        {
        }
    }
}
