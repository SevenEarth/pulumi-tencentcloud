// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb
{
    /// <summary>
    /// Provides a resource to create a CLB redirection.
    /// 
    /// ## Example Usage
    /// 
    /// Manual Rewrite
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Clb.Redirection("foo", new Tencentcloud.Clb.RedirectionArgs
    ///         {
    ///             ClbId = "lb-p7olt9e5",
    ///             SourceListenerId = "lbl-jc1dx6ju",
    ///             SourceRuleId = "loc-ft8fmngv",
    ///             TargetListenerId = "lbl-asj1hzuo",
    ///             TargetRuleId = "loc-4xxr2cy7",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Auto Rewrite
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Clb.Redirection("foo", new Tencentcloud.Clb.RedirectionArgs
    ///         {
    ///             ClbId = "lb-p7olt9e5",
    ///             IsAutoRewrite = true,
    ///             TargetListenerId = "lbl-asj1hzuo",
    ///             TargetRuleId = "loc-4xxr2cy7",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CLB redirection can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Clb/redirection:Redirection foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clb/redirection:Redirection")]
    public partial class Redirection : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Output("clbId")]
        public Output<string> ClbId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        /// </summary>
        [Output("deleteAllAutoRewrite")]
        public Output<bool?> DeleteAllAutoRewrite { get; private set; } = null!;

        /// <summary>
        /// Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        /// </summary>
        [Output("isAutoRewrite")]
        public Output<bool?> IsAutoRewrite { get; private set; } = null!;

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Output("sourceListenerId")]
        public Output<string> SourceListenerId { get; private set; } = null!;

        /// <summary>
        /// Rule ID of source listener.
        /// </summary>
        [Output("sourceRuleId")]
        public Output<string> SourceRuleId { get; private set; } = null!;

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Output("targetListenerId")]
        public Output<string> TargetListenerId { get; private set; } = null!;

        /// <summary>
        /// Rule ID of target listener.
        /// </summary>
        [Output("targetRuleId")]
        public Output<string> TargetRuleId { get; private set; } = null!;


        /// <summary>
        /// Create a Redirection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Redirection(string name, RedirectionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/redirection:Redirection", name, args ?? new RedirectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Redirection(string name, Input<string> id, RedirectionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clb/redirection:Redirection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Redirection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Redirection Get(string name, Input<string> id, RedirectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Redirection(name, id, state, options);
        }
    }

    public sealed class RedirectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Input("clbId", required: true)]
        public Input<string> ClbId { get; set; } = null!;

        /// <summary>
        /// Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        /// </summary>
        [Input("deleteAllAutoRewrite")]
        public Input<bool>? DeleteAllAutoRewrite { get; set; }

        /// <summary>
        /// Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        /// </summary>
        [Input("isAutoRewrite")]
        public Input<bool>? IsAutoRewrite { get; set; }

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Input("sourceListenerId")]
        public Input<string>? SourceListenerId { get; set; }

        /// <summary>
        /// Rule ID of source listener.
        /// </summary>
        [Input("sourceRuleId")]
        public Input<string>? SourceRuleId { get; set; }

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Input("targetListenerId", required: true)]
        public Input<string> TargetListenerId { get; set; } = null!;

        /// <summary>
        /// Rule ID of target listener.
        /// </summary>
        [Input("targetRuleId", required: true)]
        public Input<string> TargetRuleId { get; set; } = null!;

        public RedirectionArgs()
        {
        }
    }

    public sealed class RedirectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of CLB instance.
        /// </summary>
        [Input("clbId")]
        public Input<string>? ClbId { get; set; }

        /// <summary>
        /// Indicates whether delete all auto redirection. Default is `false`. It will take effect only when this redirection is auto-rewrite and this auto-rewrite auto redirected more than one rules. All the auto-rewrite relations will be deleted when this parameter set true.
        /// </summary>
        [Input("deleteAllAutoRewrite")]
        public Input<bool>? DeleteAllAutoRewrite { get; set; }

        /// <summary>
        /// Indicates whether automatic forwarding is enable, default is `false`. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.
        /// </summary>
        [Input("isAutoRewrite")]
        public Input<bool>? IsAutoRewrite { get; set; }

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Input("sourceListenerId")]
        public Input<string>? SourceListenerId { get; set; }

        /// <summary>
        /// Rule ID of source listener.
        /// </summary>
        [Input("sourceRuleId")]
        public Input<string>? SourceRuleId { get; set; }

        /// <summary>
        /// ID of source listener.
        /// </summary>
        [Input("targetListenerId")]
        public Input<string>? TargetListenerId { get; set; }

        /// <summary>
        /// Rule ID of target listener.
        /// </summary>
        [Input("targetRuleId")]
        public Input<string>? TargetRuleId { get; set; }

        public RedirectionState()
        {
        }
    }
}
