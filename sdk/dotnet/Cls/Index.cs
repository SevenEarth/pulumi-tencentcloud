// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls
{
    /// <summary>
    /// Provides a resource to create a cls index.
    /// 
    /// ## Import
    /// 
    /// cls cos index can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cls/index:Index index 0937e56f-4008-49d2-ad2d-69c52a9f11cc
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/index:Index")]
    public partial class Index : Pulumi.CustomResource
    {
        /// <summary>
        /// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
        /// </summary>
        [Output("includeInternalFields")]
        public Output<bool?> IncludeInternalFields { get; private set; } = null!;

        /// <summary>
        /// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
        /// </summary>
        [Output("metadataFlag")]
        public Output<int?> MetadataFlag { get; private set; } = null!;

        /// <summary>
        /// Index rule.
        /// </summary>
        [Output("rule")]
        public Output<Outputs.IndexRule> Rule { get; private set; } = null!;

        /// <summary>
        /// Whether to take effect. Default value: true.
        /// </summary>
        [Output("status")]
        public Output<bool> Status { get; private set; } = null!;

        /// <summary>
        /// Log topic ID.
        /// </summary>
        [Output("topicId")]
        public Output<string> TopicId { get; private set; } = null!;


        /// <summary>
        /// Create a Index resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Index(string name, IndexArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/index:Index", name, args ?? new IndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Index(string name, Input<string> id, IndexState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/index:Index", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Index resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Index Get(string name, Input<string> id, IndexState? state = null, CustomResourceOptions? options = null)
        {
            return new Index(name, id, state, options);
        }
    }

    public sealed class IndexArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
        /// </summary>
        [Input("includeInternalFields")]
        public Input<bool>? IncludeInternalFields { get; set; }

        /// <summary>
        /// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
        /// </summary>
        [Input("metadataFlag")]
        public Input<int>? MetadataFlag { get; set; }

        /// <summary>
        /// Index rule.
        /// </summary>
        [Input("rule")]
        public Input<Inputs.IndexRuleArgs>? Rule { get; set; }

        /// <summary>
        /// Whether to take effect. Default value: true.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        /// <summary>
        /// Log topic ID.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public IndexArgs()
        {
        }
    }

    public sealed class IndexState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.
        /// </summary>
        [Input("includeInternalFields")]
        public Input<bool>? IncludeInternalFields { get; set; }

        /// <summary>
        /// Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields)..
        /// </summary>
        [Input("metadataFlag")]
        public Input<int>? MetadataFlag { get; set; }

        /// <summary>
        /// Index rule.
        /// </summary>
        [Input("rule")]
        public Input<Inputs.IndexRuleGetArgs>? Rule { get; set; }

        /// <summary>
        /// Whether to take effect. Default value: true.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        /// <summary>
        /// Log topic ID.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public IndexState()
        {
        }
    }
}
