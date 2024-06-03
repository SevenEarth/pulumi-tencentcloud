// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch
{
    /// <summary>
    /// Provides a resource to stop a elasticsearch logstash pipeline
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
    ///     var stopLogstashPipelineOperation = new Tencentcloud.Elasticsearch.StopLogstashPipelineOperation("stopLogstashPipelineOperation", new()
    ///     {
    ///         InstanceId = "ls-xxxxxx",
    ///         PipelineId = "xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Elasticsearch/stopLogstashPipelineOperation:StopLogstashPipelineOperation")]
    public partial class StopLogstashPipelineOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Pipeline id.
        /// </summary>
        [Output("pipelineId")]
        public Output<string> PipelineId { get; private set; } = null!;


        /// <summary>
        /// Create a StopLogstashPipelineOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StopLogstashPipelineOperation(string name, StopLogstashPipelineOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/stopLogstashPipelineOperation:StopLogstashPipelineOperation", name, args ?? new StopLogstashPipelineOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StopLogstashPipelineOperation(string name, Input<string> id, StopLogstashPipelineOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Elasticsearch/stopLogstashPipelineOperation:StopLogstashPipelineOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StopLogstashPipelineOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StopLogstashPipelineOperation Get(string name, Input<string> id, StopLogstashPipelineOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new StopLogstashPipelineOperation(name, id, state, options);
        }
    }

    public sealed class StopLogstashPipelineOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Pipeline id.
        /// </summary>
        [Input("pipelineId", required: true)]
        public Input<string> PipelineId { get; set; } = null!;

        public StopLogstashPipelineOperationArgs()
        {
        }
        public static new StopLogstashPipelineOperationArgs Empty => new StopLogstashPipelineOperationArgs();
    }

    public sealed class StopLogstashPipelineOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Pipeline id.
        /// </summary>
        [Input("pipelineId")]
        public Input<string>? PipelineId { get; set; }

        public StopLogstashPipelineOperationState()
        {
        }
        public static new StopLogstashPipelineOperationState Empty => new StopLogstashPipelineOperationState();
    }
}
