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
    /// Provides a resource to create a monitor grafanaInstance
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-6";
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         VpcId = vpc.Id,
    ///         AvailabilityZone = availabilityZone,
    ///         CidrBlock = "10.0.1.0/24",
    ///     });
    /// 
    ///     var foo = new Tencentcloud.Monitor.GrafanaInstance("foo", new()
    ///     {
    ///         InstanceName = "test-grafana",
    ///         VpcId = vpc.Id,
    ///         SubnetIds = new[]
    ///         {
    ///             subnet.Id,
    ///         },
    ///         GrafanaInitPassword = "1234567890",
    ///         EnableInternet = false,
    ///         IsDestroy = true,
    ///         Tags = 
    ///         {
    ///             { "createdBy", "test" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// monitor grafanaInstance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Monitor/grafanaInstance:GrafanaInstance foo grafanaInstance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/grafanaInstance:GrafanaInstance")]
    public partial class GrafanaInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to automatically use vouchers.
        /// </summary>
        [Output("autoVoucher")]
        public Output<bool?> AutoVoucher { get; private set; } = null!;

        /// <summary>
        /// Control whether grafana could be accessed by internet.
        /// </summary>
        [Output("enableInternet")]
        public Output<bool> EnableInternet { get; private set; } = null!;

        /// <summary>
        /// Grafana server admin password.
        /// </summary>
        [Output("grafanaInitPassword")]
        public Output<string> GrafanaInitPassword { get; private set; } = null!;

        /// <summary>
        /// Grafana instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
        /// </summary>
        [Output("instanceStatus")]
        public Output<int> InstanceStatus { get; private set; } = null!;

        /// <summary>
        /// Grafana public address.
        /// </summary>
        [Output("internalUrl")]
        public Output<string> InternalUrl { get; private set; } = null!;

        /// <summary>
        /// Grafana intranet address.
        /// </summary>
        [Output("internetUrl")]
        public Output<string> InternetUrl { get; private set; } = null!;

        /// <summary>
        /// Whether to clean up completely, the default is false.
        /// </summary>
        [Output("isDestroy")]
        public Output<bool?> IsDestroy { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
        /// </summary>
        [Output("isDistroy")]
        public Output<bool?> IsDistroy { get; private set; } = null!;

        /// <summary>
        /// Grafana external url which could be accessed by user.
        /// </summary>
        [Output("rootUrl")]
        public Output<string> RootUrl { get; private set; } = null!;

        /// <summary>
        /// Subnet Id array.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Vpc Id.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaInstance(string name, GrafanaInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaInstance:GrafanaInstance", name, args ?? new GrafanaInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaInstance(string name, Input<string> id, GrafanaInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaInstance:GrafanaInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaInstance Get(string name, Input<string> id, GrafanaInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaInstance(name, id, state, options);
        }
    }

    public sealed class GrafanaInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically use vouchers.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Control whether grafana could be accessed by internet.
        /// </summary>
        [Input("enableInternet")]
        public Input<bool>? EnableInternet { get; set; }

        /// <summary>
        /// Grafana server admin password.
        /// </summary>
        [Input("grafanaInitPassword")]
        public Input<string>? GrafanaInitPassword { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// Whether to clean up completely, the default is false.
        /// </summary>
        [Input("isDestroy")]
        public Input<bool>? IsDestroy { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
        /// </summary>
        [Input("isDistroy")]
        public Input<bool>? IsDistroy { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Subnet Id array.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

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
        /// Vpc Id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GrafanaInstanceArgs()
        {
        }
        public static new GrafanaInstanceArgs Empty => new GrafanaInstanceArgs();
    }

    public sealed class GrafanaInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically use vouchers.
        /// </summary>
        [Input("autoVoucher")]
        public Input<bool>? AutoVoucher { get; set; }

        /// <summary>
        /// Control whether grafana could be accessed by internet.
        /// </summary>
        [Input("enableInternet")]
        public Input<bool>? EnableInternet { get; set; }

        /// <summary>
        /// Grafana server admin password.
        /// </summary>
        [Input("grafanaInitPassword")]
        public Input<string>? GrafanaInitPassword { get; set; }

        /// <summary>
        /// Grafana instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
        /// </summary>
        [Input("instanceStatus")]
        public Input<int>? InstanceStatus { get; set; }

        /// <summary>
        /// Grafana public address.
        /// </summary>
        [Input("internalUrl")]
        public Input<string>? InternalUrl { get; set; }

        /// <summary>
        /// Grafana intranet address.
        /// </summary>
        [Input("internetUrl")]
        public Input<string>? InternetUrl { get; set; }

        /// <summary>
        /// Whether to clean up completely, the default is false.
        /// </summary>
        [Input("isDestroy")]
        public Input<bool>? IsDestroy { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
        /// </summary>
        [Input("isDistroy")]
        public Input<bool>? IsDistroy { get; set; }

        /// <summary>
        /// Grafana external url which could be accessed by user.
        /// </summary>
        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Subnet Id array.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

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
        /// Vpc Id.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public GrafanaInstanceState()
        {
        }
        public static new GrafanaInstanceState Empty => new GrafanaInstanceState();
    }
}
