// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos
{
    /// <summary>
    /// Provides a resource to create a antiddos ddos_geo_ip_block_config
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
    ///     var ddosGeoIpBlockConfig = new Tencentcloud.Antiddos.DdosGeoIpBlockConfig("ddosGeoIpBlockConfig", new()
    ///     {
    ///         AntiddosDdosGeoIpBlockConfig = new Tencentcloud.Antiddos.Inputs.DdosGeoIpBlockConfigDdosGeoIpBlockConfigArgs
    ///         {
    ///             Action = "drop",
    ///             AreaLists = new[]
    ///             {
    ///                 100002,
    ///             },
    ///             RegionType = "customized",
    ///         },
    ///         InstanceId = "bgp-xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// antiddos ddos_geo_ip_block_config can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig ddos_geo_ip_block_config ${instanceId}#${configId}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig")]
    public partial class DdosGeoIpBlockConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Output("ddosGeoIpBlockConfig")]
        public Output<Outputs.DdosGeoIpBlockConfigDdosGeoIpBlockConfig> AntiddosDdosGeoIpBlockConfig { get; private set; } = null!;

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a DdosGeoIpBlockConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosGeoIpBlockConfig(string name, DdosGeoIpBlockConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig", name, args ?? new DdosGeoIpBlockConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosGeoIpBlockConfig(string name, Input<string> id, DdosGeoIpBlockConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosGeoIpBlockConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosGeoIpBlockConfig Get(string name, Input<string> id, DdosGeoIpBlockConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosGeoIpBlockConfig(name, id, state, options);
        }
    }

    public sealed class DdosGeoIpBlockConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Input("ddosGeoIpBlockConfig", required: true)]
        public Input<Inputs.DdosGeoIpBlockConfigDdosGeoIpBlockConfigArgs> AntiddosDdosGeoIpBlockConfig { get; set; } = null!;

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DdosGeoIpBlockConfigArgs()
        {
        }
        public static new DdosGeoIpBlockConfigArgs Empty => new DdosGeoIpBlockConfigArgs();
    }

    public sealed class DdosGeoIpBlockConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
        /// </summary>
        [Input("ddosGeoIpBlockConfig")]
        public Input<Inputs.DdosGeoIpBlockConfigDdosGeoIpBlockConfigGetArgs>? AntiddosDdosGeoIpBlockConfig { get; set; }

        /// <summary>
        /// InstanceId.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DdosGeoIpBlockConfigState()
        {
        }
        public static new DdosGeoIpBlockConfigState Empty => new DdosGeoIpBlockConfigState();
    }
}
