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
    /// Provides a resource to create a monitor grafana_sso_config
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
    ///     var grafanaSsoConfig = new Tencentcloud.Monitor.GrafanaSsoConfig("grafanaSsoConfig", new()
    ///     {
    ///         EnableSso = false,
    ///         InstanceId = "grafana-dp2hnnfa",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// monitor grafana_sso_config can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig grafana_sso_config instance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig")]
    public partial class GrafanaSsoConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable SSO: `true` for enabling; `false` for disabling.
        /// </summary>
        [Output("enableSso")]
        public Output<bool> EnableSso { get; private set; } = null!;

        /// <summary>
        /// Grafana instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a GrafanaSsoConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrafanaSsoConfig(string name, GrafanaSsoConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig", name, args ?? new GrafanaSsoConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrafanaSsoConfig(string name, Input<string> id, GrafanaSsoConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/grafanaSsoConfig:GrafanaSsoConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrafanaSsoConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrafanaSsoConfig Get(string name, Input<string> id, GrafanaSsoConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new GrafanaSsoConfig(name, id, state, options);
        }
    }

    public sealed class GrafanaSsoConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable SSO: `true` for enabling; `false` for disabling.
        /// </summary>
        [Input("enableSso", required: true)]
        public Input<bool> EnableSso { get; set; } = null!;

        /// <summary>
        /// Grafana instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public GrafanaSsoConfigArgs()
        {
        }
        public static new GrafanaSsoConfigArgs Empty => new GrafanaSsoConfigArgs();
    }

    public sealed class GrafanaSsoConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable SSO: `true` for enabling; `false` for disabling.
        /// </summary>
        [Input("enableSso")]
        public Input<bool>? EnableSso { get; set; }

        /// <summary>
        /// Grafana instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public GrafanaSsoConfigState()
        {
        }
        public static new GrafanaSsoConfigState Empty => new GrafanaSsoConfigState();
    }
}
