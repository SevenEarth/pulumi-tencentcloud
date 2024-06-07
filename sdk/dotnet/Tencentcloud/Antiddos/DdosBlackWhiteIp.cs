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
    /// Provides a resource to create a antiddos ddos black white ip
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
    ///     var ddosBlackWhiteIp = new Tencentcloud.Antiddos.DdosBlackWhiteIp("ddosBlackWhiteIp", new()
    ///     {
    ///         InstanceId = "bgp-xxxxxx",
    ///         Ip = "1.2.3.5",
    ///         Mask = 0,
    ///         Type = "black",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// antiddos ddos_black_white_ip can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Antiddos/ddosBlackWhiteIp:DdosBlackWhiteIp ddos_black_white_ip ${instanceId}#${ip}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Antiddos/ddosBlackWhiteIp:DdosBlackWhiteIp")]
    public partial class DdosBlackWhiteIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// ip list.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// ip mask.
        /// </summary>
        [Output("mask")]
        public Output<int> Mask { get; private set; } = null!;

        /// <summary>
        /// ip type, black: black ip list, white: white ip list.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DdosBlackWhiteIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosBlackWhiteIp(string name, DdosBlackWhiteIpArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosBlackWhiteIp:DdosBlackWhiteIp", name, args ?? new DdosBlackWhiteIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosBlackWhiteIp(string name, Input<string> id, DdosBlackWhiteIpState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ddosBlackWhiteIp:DdosBlackWhiteIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosBlackWhiteIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosBlackWhiteIp Get(string name, Input<string> id, DdosBlackWhiteIpState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosBlackWhiteIp(name, id, state, options);
        }
    }

    public sealed class DdosBlackWhiteIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// ip list.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// ip mask.
        /// </summary>
        [Input("mask", required: true)]
        public Input<int> Mask { get; set; } = null!;

        /// <summary>
        /// ip type, black: black ip list, white: white ip list.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DdosBlackWhiteIpArgs()
        {
        }
        public static new DdosBlackWhiteIpArgs Empty => new DdosBlackWhiteIpArgs();
    }

    public sealed class DdosBlackWhiteIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ip list.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// ip mask.
        /// </summary>
        [Input("mask")]
        public Input<int>? Mask { get; set; }

        /// <summary>
        /// ip type, black: black ip list, white: white ip list.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DdosBlackWhiteIpState()
        {
        }
        public static new DdosBlackWhiteIpState Empty => new DdosBlackWhiteIpState();
    }
}
