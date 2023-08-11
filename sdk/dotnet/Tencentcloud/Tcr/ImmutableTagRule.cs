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
    /// Provides a resource to create a tcr immutable tag rule.
    /// 
    /// ## Example Usage
    /// ### Create a immutable tag rule with specified tags and exclude specified repositories
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleInstance = new Tencentcloud.Tcr.Instance("exampleInstance", new Tencentcloud.Tcr.InstanceArgs
    ///         {
    ///             InstanceType = "premium",
    ///             DeleteBucket = true,
    ///         });
    ///         var exampleNamespace = new Tencentcloud.Tcr.Namespace("exampleNamespace", new Tencentcloud.Tcr.NamespaceArgs
    ///         {
    ///             InstanceId = exampleInstance.Id,
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
    ///         var exampleImmutableTagRule = new Tencentcloud.Tcr.ImmutableTagRule("exampleImmutableTagRule", new Tencentcloud.Tcr.ImmutableTagRuleArgs
    ///         {
    ///             RegistryId = exampleInstance.Id,
    ///             NamespaceName = exampleNamespace.Name,
    ///             Rule = new Tencentcloud.Tcr.Inputs.ImmutableTagRuleRuleArgs
    ///             {
    ///                 RepositoryPattern = "deprecated_repo",
    ///                 TagPattern = "**",
    ///                 RepositoryDecoration = "repoExcludes",
    ///                 TagDecoration = "matches",
    ///                 Disabled = false,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With specified repositories and exclude specified version tag
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Tcr.ImmutableTagRule("example", new Tencentcloud.Tcr.ImmutableTagRuleArgs
    ///         {
    ///             RegistryId = tencentcloud_tcr_instance.Example.Id,
    ///             NamespaceName = tencentcloud_tcr_namespace.Example.Name,
    ///             Rule = new Tencentcloud.Tcr.Inputs.ImmutableTagRuleRuleArgs
    ///             {
    ///                 RepositoryPattern = "**",
    ///                 TagPattern = "v1",
    ///                 RepositoryDecoration = "repoMatches",
    ///                 TagDecoration = "excludes",
    ///                 Disabled = false,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Disabled the specified rule
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleRuleA = new Tencentcloud.Tcr.ImmutableTagRule("exampleRuleA", new Tencentcloud.Tcr.ImmutableTagRuleArgs
    ///         {
    ///             RegistryId = tencentcloud_tcr_instance.Example.Id,
    ///             NamespaceName = tencentcloud_tcr_namespace.Example.Name,
    ///             Rule = new Tencentcloud.Tcr.Inputs.ImmutableTagRuleRuleArgs
    ///             {
    ///                 RepositoryPattern = "deprecated_repo",
    ///                 TagPattern = "**",
    ///                 RepositoryDecoration = "repoExcludes",
    ///                 TagDecoration = "matches",
    ///                 Disabled = false,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///         var exampleRuleB = new Tencentcloud.Tcr.ImmutableTagRule("exampleRuleB", new Tencentcloud.Tcr.ImmutableTagRuleArgs
    ///         {
    ///             RegistryId = tencentcloud_tcr_instance.Example.Id,
    ///             NamespaceName = tencentcloud_tcr_namespace.Example.Name,
    ///             Rule = new Tencentcloud.Tcr.Inputs.ImmutableTagRuleRuleArgs
    ///             {
    ///                 RepositoryPattern = "**",
    ///                 TagPattern = "v1",
    ///                 RepositoryDecoration = "repoMatches",
    ///                 TagDecoration = "excludes",
    ///                 Disabled = true,
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tcr immutable_tag_rule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tcr/immutableTagRule:ImmutableTagRule immutable_tag_rule immutable_tag_rule_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/immutableTagRule:ImmutableTagRule")]
    public partial class ImmutableTagRule : Pulumi.CustomResource
    {
        /// <summary>
        /// namespace name.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// rule.
        /// </summary>
        [Output("rule")]
        public Output<Outputs.ImmutableTagRuleRule> Rule { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ImmutableTagRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImmutableTagRule(string name, ImmutableTagRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/immutableTagRule:ImmutableTagRule", name, args ?? new ImmutableTagRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImmutableTagRule(string name, Input<string> id, ImmutableTagRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/immutableTagRule:ImmutableTagRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImmutableTagRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImmutableTagRule Get(string name, Input<string> id, ImmutableTagRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ImmutableTagRule(name, id, state, options);
        }
    }

    public sealed class ImmutableTagRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// rule.
        /// </summary>
        [Input("rule", required: true)]
        public Input<Inputs.ImmutableTagRuleRuleArgs> Rule { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ImmutableTagRuleArgs()
        {
        }
    }

    public sealed class ImmutableTagRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// rule.
        /// </summary>
        [Input("rule")]
        public Input<Inputs.ImmutableTagRuleRuleGetArgs>? Rule { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ImmutableTagRuleState()
        {
        }
    }
}
