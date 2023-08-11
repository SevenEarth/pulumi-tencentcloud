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
    /// Provides a resource to configure a tcr tag retention execution.
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
    ///         var exampleInstance = new Tencentcloud.Tcr.Instance("exampleInstance", new Tencentcloud.Tcr.InstanceArgs
    ///         {
    ///             InstanceType = "basic",
    ///             DeleteBucket = true,
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
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
    ///         var exampleTagRetentionRule = new Tencentcloud.Tcr.TagRetentionRule("exampleTagRetentionRule", new Tencentcloud.Tcr.TagRetentionRuleArgs
    ///         {
    ///             RegistryId = exampleInstance.Id,
    ///             NamespaceName = exampleNamespace.Name,
    ///             RetentionRule = new Tencentcloud.Tcr.Inputs.TagRetentionRuleRetentionRuleArgs
    ///             {
    ///                 Key = "nDaysSinceLastPush",
    ///                 Value = 2,
    ///             },
    ///             CronSetting = "manual",
    ///             Disabled = true,
    ///         });
    ///         var exampleTagRetentionExecutionConfig = new Tencentcloud.Tcr.TagRetentionExecutionConfig("exampleTagRetentionExecutionConfig", new Tencentcloud.Tcr.TagRetentionExecutionConfigArgs
    ///         {
    ///             RegistryId = exampleTagRetentionRule.RegistryId,
    ///             RetentionId = exampleTagRetentionRule.RetentionId,
    ///             DryRun = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/tagRetentionExecutionConfig:TagRetentionExecutionConfig")]
    public partial class TagRetentionExecutionConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to simulate execution, the default value is false, that is, non-simulation execution.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// execution id.
        /// </summary>
        [Output("executionId")]
        public Output<int> ExecutionId { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// retention id.
        /// </summary>
        [Output("retentionId")]
        public Output<int> RetentionId { get; private set; } = null!;


        /// <summary>
        /// Create a TagRetentionExecutionConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagRetentionExecutionConfig(string name, TagRetentionExecutionConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/tagRetentionExecutionConfig:TagRetentionExecutionConfig", name, args ?? new TagRetentionExecutionConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagRetentionExecutionConfig(string name, Input<string> id, TagRetentionExecutionConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/tagRetentionExecutionConfig:TagRetentionExecutionConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagRetentionExecutionConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagRetentionExecutionConfig Get(string name, Input<string> id, TagRetentionExecutionConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TagRetentionExecutionConfig(name, id, state, options);
        }
    }

    public sealed class TagRetentionExecutionConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to simulate execution, the default value is false, that is, non-simulation execution.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// retention id.
        /// </summary>
        [Input("retentionId", required: true)]
        public Input<int> RetentionId { get; set; } = null!;

        public TagRetentionExecutionConfigArgs()
        {
        }
    }

    public sealed class TagRetentionExecutionConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to simulate execution, the default value is false, that is, non-simulation execution.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// execution id.
        /// </summary>
        [Input("executionId")]
        public Input<int>? ExecutionId { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// retention id.
        /// </summary>
        [Input("retentionId")]
        public Input<int>? RetentionId { get; set; }

        public TagRetentionExecutionConfigState()
        {
        }
    }
}
