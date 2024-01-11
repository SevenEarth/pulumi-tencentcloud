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
    /// Provides a resource to create a waf cc
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
    ///         var example = new Tencentcloud.Waf.Cc("example", new Tencentcloud.Waf.CcArgs
    ///         {
    ///             ActionType = "22",
    ///             Advance = "0",
    ///             Domain = "www.demo.com",
    ///             Edition = "sparta-waf",
    ///             Interval = "60",
    ///             Limit = "60",
    ///             MatchFunc = 0,
    ///             Priority = 50,
    ///             Status = 1,
    ///             Type = 1,
    ///             Url = "/cc_demo",
    ///             ValidTime = 600,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Waf/cc:Cc")]
    public partial class Cc : Pulumi.CustomResource
    {
        /// <summary>
        /// Rule Action, 20 log, 21 captcha, 22 deny, 23 accurate deny.
        /// </summary>
        [Output("actionType")]
        public Output<string> ActionType { get; private set; } = null!;

        /// <summary>
        /// Session match mode, 0 use session, 1 use ip.
        /// </summary>
        [Output("advance")]
        public Output<string> Advance { get; private set; } = null!;

        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// Event ID.
        /// </summary>
        [Output("eventId")]
        public Output<string?> EventId { get; private set; } = null!;

        /// <summary>
        /// Interval.
        /// </summary>
        [Output("interval")]
        public Output<string> Interval { get; private set; } = null!;

        /// <summary>
        /// CC detection threshold.
        /// </summary>
        [Output("limit")]
        public Output<string> Limit { get; private set; } = null!;

        /// <summary>
        /// Match method, 0 equal, 1 contains, 2 prefix.
        /// </summary>
        [Output("matchFunc")]
        public Output<int> MatchFunc { get; private set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Rule Priority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;

        /// <summary>
        /// Advance mode use session id.
        /// </summary>
        [Output("sessionApplieds")]
        public Output<ImmutableArray<int>> SessionApplieds { get; private set; } = null!;

        /// <summary>
        /// Rule Status, 0 rule close, 1 rule open.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Operate Type.
        /// </summary>
        [Output("type")]
        public Output<int?> Type { get; private set; } = null!;

        /// <summary>
        /// Check URL.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Action ValidTime, minute unit. Min: 60, Max: 604800.
        /// </summary>
        [Output("validTime")]
        public Output<int> ValidTime { get; private set; } = null!;


        /// <summary>
        /// Create a Cc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cc(string name, CcArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/cc:Cc", name, args ?? new CcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cc(string name, Input<string> id, CcState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/cc:Cc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cc Get(string name, Input<string> id, CcState? state = null, CustomResourceOptions? options = null)
        {
            return new Cc(name, id, state, options);
        }
    }

    public sealed class CcArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule Action, 20 log, 21 captcha, 22 deny, 23 accurate deny.
        /// </summary>
        [Input("actionType", required: true)]
        public Input<string> ActionType { get; set; } = null!;

        /// <summary>
        /// Session match mode, 0 use session, 1 use ip.
        /// </summary>
        [Input("advance", required: true)]
        public Input<string> Advance { get; set; } = null!;

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        /// </summary>
        [Input("edition", required: true)]
        public Input<string> Edition { get; set; } = null!;

        /// <summary>
        /// Event ID.
        /// </summary>
        [Input("eventId")]
        public Input<string>? EventId { get; set; }

        /// <summary>
        /// Interval.
        /// </summary>
        [Input("interval", required: true)]
        public Input<string> Interval { get; set; } = null!;

        /// <summary>
        /// CC detection threshold.
        /// </summary>
        [Input("limit", required: true)]
        public Input<string> Limit { get; set; } = null!;

        /// <summary>
        /// Match method, 0 equal, 1 contains, 2 prefix.
        /// </summary>
        [Input("matchFunc", required: true)]
        public Input<int> MatchFunc { get; set; } = null!;

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Rule Priority.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("sessionApplieds")]
        private InputList<int>? _sessionApplieds;

        /// <summary>
        /// Advance mode use session id.
        /// </summary>
        public InputList<int> SessionApplieds
        {
            get => _sessionApplieds ?? (_sessionApplieds = new InputList<int>());
            set => _sessionApplieds = value;
        }

        /// <summary>
        /// Rule Status, 0 rule close, 1 rule open.
        /// </summary>
        [Input("status", required: true)]
        public Input<int> Status { get; set; } = null!;

        /// <summary>
        /// Operate Type.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        /// <summary>
        /// Check URL.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Action ValidTime, minute unit. Min: 60, Max: 604800.
        /// </summary>
        [Input("validTime", required: true)]
        public Input<int> ValidTime { get; set; } = null!;

        public CcArgs()
        {
        }
    }

    public sealed class CcState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule Action, 20 log, 21 captcha, 22 deny, 23 accurate deny.
        /// </summary>
        [Input("actionType")]
        public Input<string>? ActionType { get; set; }

        /// <summary>
        /// Session match mode, 0 use session, 1 use ip.
        /// </summary>
        [Input("advance")]
        public Input<string>? Advance { get; set; }

        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Event ID.
        /// </summary>
        [Input("eventId")]
        public Input<string>? EventId { get; set; }

        /// <summary>
        /// Interval.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// CC detection threshold.
        /// </summary>
        [Input("limit")]
        public Input<string>? Limit { get; set; }

        /// <summary>
        /// Match method, 0 equal, 1 contains, 2 prefix.
        /// </summary>
        [Input("matchFunc")]
        public Input<int>? MatchFunc { get; set; }

        /// <summary>
        /// Rule Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Rule Priority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        [Input("sessionApplieds")]
        private InputList<int>? _sessionApplieds;

        /// <summary>
        /// Advance mode use session id.
        /// </summary>
        public InputList<int> SessionApplieds
        {
            get => _sessionApplieds ?? (_sessionApplieds = new InputList<int>());
            set => _sessionApplieds = value;
        }

        /// <summary>
        /// Rule Status, 0 rule close, 1 rule open.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Operate Type.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        /// <summary>
        /// Check URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Action ValidTime, minute unit. Min: 60, Max: 604800.
        /// </summary>
        [Input("validTime")]
        public Input<int>? ValidTime { get; set; }

        public CcState()
        {
        }
    }
}