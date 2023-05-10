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
    /// Provides a resource to create a monitor grafanaNotificationChannel
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
    ///         var grafanaNotificationChannel = new Tencentcloud.Monitor.GrafanaNotificationChannel("grafanaNotificationChannel", new Tencentcloud.Monitor.GrafanaNotificationChannelArgs
    ///         {
    ///             ChannelName = "create-channel",
    ///             ExtraOrgIds = {},
    ///             InstanceId = "grafana-50nj6v00",
    ///             OrgId = 1,
    ///             Receivers = 
    ///             {
    ///                 "Consumer-6vkna7pevq",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/grafanaNotificationChannel:GrafanaNotificationChannel")]
    public partial class GrafanaNotificationChannel : Pulumi.CustomResource
    {
        /// <summary>
        /// plugin id.
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// channel name.
        /// </summary>
        [Output("channelName")]
        public Output<string> ChannelName { get; private set; } = null!;

        /// <summary>
        /// extra grafana organization id list, default to 1 representing Main Org.
        /// </summary>
        [Output("extraOrgIds")]
        public Output<ImmutableArray<string>> ExtraOrgIds { get; private set; } = null!;

        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Grafana organization which channel will be installed, default to 1 representing Main Org.
        /// </summary>
        [Output("orgId")]
        public Output<int> OrgId { get; private set; } = null!;

        /// <summary>
        /// cloud monitor notification template notice-id list.
        /// </summary>
        [Output("receivers")]
        public Output<ImmutableArray<string>> Receivers { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaNotificationChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaNotificationChannel(string name, GrafanaNotificationChannelArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaNotificationChannel:GrafanaNotificationChannel", name, args ?? new GrafanaNotificationChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaNotificationChannel(string name, Input<string> id, GrafanaNotificationChannelState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaNotificationChannel:GrafanaNotificationChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaNotificationChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaNotificationChannel Get(string name, Input<string> id, GrafanaNotificationChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaNotificationChannel(name, id, state, options);
        }
    }

    public sealed class GrafanaNotificationChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// channel name.
        /// </summary>
        [Input("channelName")]
        public Input<string>? ChannelName { get; set; }

        [Input("extraOrgIds")]
        private InputList<string>? _extraOrgIds;

        /// <summary>
        /// extra grafana organization id list, default to 1 representing Main Org.
        /// </summary>
        public InputList<string> ExtraOrgIds
        {
            get => _extraOrgIds ?? (_extraOrgIds = new InputList<string>());
            set => _extraOrgIds = value;
        }

        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Grafana organization which channel will be installed, default to 1 representing Main Org.
        /// </summary>
        [Input("orgId")]
        public Input<int>? OrgId { get; set; }

        [Input("receivers")]
        private InputList<string>? _receivers;

        /// <summary>
        /// cloud monitor notification template notice-id list.
        /// </summary>
        public InputList<string> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<string>());
            set => _receivers = value;
        }

        public GrafanaNotificationChannelArgs()
        {
        }
    }

    public sealed class GrafanaNotificationChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// plugin id.
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// channel name.
        /// </summary>
        [Input("channelName")]
        public Input<string>? ChannelName { get; set; }

        [Input("extraOrgIds")]
        private InputList<string>? _extraOrgIds;

        /// <summary>
        /// extra grafana organization id list, default to 1 representing Main Org.
        /// </summary>
        public InputList<string> ExtraOrgIds
        {
            get => _extraOrgIds ?? (_extraOrgIds = new InputList<string>());
            set => _extraOrgIds = value;
        }

        /// <summary>
        /// grafana instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Grafana organization which channel will be installed, default to 1 representing Main Org.
        /// </summary>
        [Input("orgId")]
        public Input<int>? OrgId { get; set; }

        [Input("receivers")]
        private InputList<string>? _receivers;

        /// <summary>
        /// cloud monitor notification template notice-id list.
        /// </summary>
        public InputList<string> Receivers
        {
            get => _receivers ?? (_receivers = new InputList<string>());
            set => _receivers = value;
        }

        public GrafanaNotificationChannelState()
        {
        }
    }
}
