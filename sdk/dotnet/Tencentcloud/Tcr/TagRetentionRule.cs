// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    /// <summary>
    /// Provides a resource to create a tcr tag_retention_rule
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
    ///         var myNs = new Tencentcloud.Tcr.Namespace("myNs", new Tencentcloud.Tcr.NamespaceArgs
    ///         {
    ///             InstanceId = local.Tcr_id,
    ///             IsPublic = true,
    ///             IsAutoScan = true,
    ///             IsPreventVul = true,
    ///             Severity = "medium",
    ///             CveWhitelistItems = 
    ///             {
    ///                 new Tencentcloud.Tcr.Inputs.NamespaceCveWhitelistItemArgs
    ///                 {
    ///                     CveId = "cve-xxxxx",
    ///                 },
    ///             },
    ///         });
    ///         var myRule = new Tencentcloud.Tcr.TagRetentionRule("myRule", new Tencentcloud.Tcr.TagRetentionRuleArgs
    ///         {
    ///             RegistryId = local.Tcr_id,
    ///             NamespaceName = myNs.Name,
    ///             RetentionRule = new Tencentcloud.Tcr.Inputs.TagRetentionRuleRetentionRuleArgs
    ///             {
    ///                 Key = "nDaysSinceLastPush",
    ///                 Value = 2,
    ///             },
    ///             CronSetting = "daily",
    ///             Disabled = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/tagRetentionRule:TagRetentionRule")]
    public partial class TagRetentionRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        /// </summary>
        [Output("cronSetting")]
        public Output<string> CronSetting { get; private set; } = null!;

        /// <summary>
        /// Whether to disable the rule, with the default value of false.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// The Name of the namespace.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// The main instance ID.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Retention Policy.
        /// </summary>
        [Output("retentionRule")]
        public Output<Outputs.TagRetentionRuleRetentionRule> RetentionRule { get; private set; } = null!;


        /// <summary>
        /// Create a TagRetentionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagRetentionRule(string name, TagRetentionRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/tagRetentionRule:TagRetentionRule", name, args ?? new TagRetentionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagRetentionRule(string name, Input<string> id, TagRetentionRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/tagRetentionRule:TagRetentionRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagRetentionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagRetentionRule Get(string name, Input<string> id, TagRetentionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new TagRetentionRule(name, id, state, options);
        }
    }

    public sealed class TagRetentionRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        /// </summary>
        [Input("cronSetting", required: true)]
        public Input<string> CronSetting { get; set; } = null!;

        /// <summary>
        /// Whether to disable the rule, with the default value of false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The Name of the namespace.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The main instance ID.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// Retention Policy.
        /// </summary>
        [Input("retentionRule", required: true)]
        public Input<Inputs.TagRetentionRuleRetentionRuleArgs> RetentionRule { get; set; } = null!;

        public TagRetentionRuleArgs()
        {
        }
    }

    public sealed class TagRetentionRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
        /// </summary>
        [Input("cronSetting")]
        public Input<string>? CronSetting { get; set; }

        /// <summary>
        /// Whether to disable the rule, with the default value of false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The Name of the namespace.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// The main instance ID.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// Retention Policy.
        /// </summary>
        [Input("retentionRule")]
        public Input<Inputs.TagRetentionRuleRetentionRuleGetArgs>? RetentionRule { get; set; }

        public TagRetentionRuleState()
        {
        }
    }
}
