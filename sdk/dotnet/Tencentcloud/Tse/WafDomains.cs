// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse
{
    /// <summary>
    /// Provides a resource to create a tse waf_domains
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
    ///     var wafDomains = new Tencentcloud.Tse.WafDomains("wafDomains", new()
    ///     {
    ///         Domain = "tse.exmaple.com",
    ///         GatewayId = "gateway-ed63e957",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// tse waf_domains can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Tse/wafDomains:WafDomains waf_domains waf_domains_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tse/wafDomains:WafDomains")]
    public partial class WafDomains : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The waf protected domain name.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;


        /// <summary>
        /// Create a WafDomains resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WafDomains(string name, WafDomainsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/wafDomains:WafDomains", name, args ?? new WafDomainsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WafDomains(string name, Input<string> id, WafDomainsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/wafDomains:WafDomains", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WafDomains resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WafDomains Get(string name, Input<string> id, WafDomainsState? state = null, CustomResourceOptions? options = null)
        {
            return new WafDomains(name, id, state, options);
        }
    }

    public sealed class WafDomainsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The waf protected domain name.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        public WafDomainsArgs()
        {
        }
        public static new WafDomainsArgs Empty => new WafDomainsArgs();
    }

    public sealed class WafDomainsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The waf protected domain name.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        public WafDomainsState()
        {
        }
        public static new WafDomainsState Empty => new WafDomainsState();
    }
}
