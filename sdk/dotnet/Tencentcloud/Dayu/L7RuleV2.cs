// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu
{
    /// <summary>
    /// Use this resource to create dayu new layer 7 rule
    /// 
    /// &gt; **NOTE:** This resource only support resource Anti-DDoS of type `bgpip`
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
    ///         var tencentcloudDayuL7RuleV2 = new Tencentcloud.Dayu.L7RuleV2("tencentcloudDayuL7RuleV2", new Tencentcloud.Dayu.L7RuleV2Args
    ///         {
    ///             ResourceId = "bgpip-000004xe",
    ///             ResourceIp = "119.28.217.162",
    ///             ResourceType = "bgpip",
    ///             Rule = new Tencentcloud.Dayu.Inputs.L7RuleV2RuleArgs
    ///             {
    ///                 Domain = "github.com",
    ///                 KeepEnable = false,
    ///                 Keeptime = 0,
    ///                 LbType = 1,
    ///                 Protocol = "http",
    ///                 SourceLists = 
    ///                 {
    ///                     new Tencentcloud.Dayu.Inputs.L7RuleV2RuleSourceListArgs
    ///                     {
    ///                         Source = "1.2.3.5",
    ///                         Weight = 100,
    ///                     },
    ///                     new Tencentcloud.Dayu.Inputs.L7RuleV2RuleSourceListArgs
    ///                     {
    ///                         Source = "1.2.3.6",
    ///                         Weight = 100,
    ///                     },
    ///                 },
    ///                 SourceType = 2,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dayu/l7RuleV2:L7RuleV2")]
    public partial class L7RuleV2 : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the resource that the layer 7 rule works for.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// Ip of the resource that the layer 7 rule works for.
        /// </summary>
        [Output("resourceIp")]
        public Output<string> ResourceIp { get; private set; } = null!;

        /// <summary>
        /// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// A list of layer 7 rules. Each element contains the following attributes:
        /// </summary>
        [Output("rule")]
        public Output<Outputs.L7RuleV2Rule> Rule { get; private set; } = null!;


        /// <summary>
        /// Create a L7RuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public L7RuleV2(string name, L7RuleV2Args args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dayu/l7RuleV2:L7RuleV2", name, args ?? new L7RuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private L7RuleV2(string name, Input<string> id, L7RuleV2State? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dayu/l7RuleV2:L7RuleV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing L7RuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static L7RuleV2 Get(string name, Input<string> id, L7RuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new L7RuleV2(name, id, state, options);
        }
    }

    public sealed class L7RuleV2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the resource that the layer 7 rule works for.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// Ip of the resource that the layer 7 rule works for.
        /// </summary>
        [Input("resourceIp", required: true)]
        public Input<string> ResourceIp { get; set; } = null!;

        /// <summary>
        /// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        /// <summary>
        /// A list of layer 7 rules. Each element contains the following attributes:
        /// </summary>
        [Input("rule", required: true)]
        public Input<Inputs.L7RuleV2RuleArgs> Rule { get; set; } = null!;

        public L7RuleV2Args()
        {
        }
    }

    public sealed class L7RuleV2State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the resource that the layer 7 rule works for.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Ip of the resource that the layer 7 rule works for.
        /// </summary>
        [Input("resourceIp")]
        public Input<string>? ResourceIp { get; set; }

        /// <summary>
        /// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        /// <summary>
        /// A list of layer 7 rules. Each element contains the following attributes:
        /// </summary>
        [Input("rule")]
        public Input<Inputs.L7RuleV2RuleGetArgs>? Rule { get; set; }

        public L7RuleV2State()
        {
        }
    }
}
