// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf
{
    /// <summary>
    /// Provides a resource to create a waf anti_fake
    /// 
    /// &gt; **NOTE:** Uri: Please configure static resources such as. html,. shtml,. txt,. js,. css,. jpg,. png, or access paths for static resources..
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
    ///         var example = new Tencentcloud.Waf.AntiFake("example", new Tencentcloud.Waf.AntiFakeArgs
    ///         {
    ///             Domain = "www.waf.com",
    ///             Status = 1,
    ///             Uri = "/anti_fake_url.html",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// waf anti_fake can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Waf/antiFake:AntiFake example 3200035516#www.waf.com
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Waf/antiFake:AntiFake")]
    public partial class AntiFake : Pulumi.CustomResource
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// protocol.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// rule id.
        /// </summary>
        [Output("ruleId")]
        public Output<int> RuleId { get; private set; } = null!;

        /// <summary>
        /// status. 0: Turn off rules and log switches, 1: Turn on the rule switch and Turn off the log switch; 2: Turn off the rule switch and turn on the log switch;3: Turn on the log switch.
        /// </summary>
        [Output("status")]
        public Output<int?> Status { get; private set; } = null!;

        /// <summary>
        /// Uri.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a AntiFake resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntiFake(string name, AntiFakeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/antiFake:AntiFake", name, args ?? new AntiFakeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntiFake(string name, Input<string> id, AntiFakeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/antiFake:AntiFake", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AntiFake resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntiFake Get(string name, Input<string> id, AntiFakeState? state = null, CustomResourceOptions? options = null)
        {
            return new AntiFake(name, id, state, options);
        }
    }

    public sealed class AntiFakeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// status. 0: Turn off rules and log switches, 1: Turn on the rule switch and Turn off the log switch; 2: Turn off the rule switch and turn on the log switch;3: Turn on the log switch.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Uri.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public AntiFakeArgs()
        {
        }
    }

    public sealed class AntiFakeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// rule id.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// status. 0: Turn off rules and log switches, 1: Turn on the rule switch and Turn off the log switch; 2: Turn off the rule switch and turn on the log switch;3: Turn on the log switch.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Uri.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public AntiFakeState()
        {
        }
    }
}
