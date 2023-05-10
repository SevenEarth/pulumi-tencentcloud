// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcmq
{
    /// <summary>
    /// Provides a resource to create a tcmq topic
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
    ///         var topic = new Tencentcloud.Tcmq.Topic("topic", new Tencentcloud.Tcmq.TopicArgs
    ///         {
    ///             TopicName = "topic_name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tcmq topic can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tcmq/topic:Topic topic topic_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcmq/topic:Topic")]
    public partial class Topic : Pulumi.CustomResource
    {
        /// <summary>
        /// Used to specify the message match policy for the topic. `1`: tag match policy (default value); `2`: routing match policy.
        /// </summary>
        [Output("filterType")]
        public Output<int?> FilterType { get; private set; } = null!;

        /// <summary>
        /// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
        /// </summary>
        [Output("maxMsgSize")]
        public Output<int?> MaxMsgSize { get; private set; } = null!;

        /// <summary>
        /// Message retention period. Value range: 60-86400 seconds (i.e., 1 minute-1 day). Default value: 86400.
        /// </summary>
        [Output("msgRetentionSeconds")]
        public Output<int?> MsgRetentionSeconds { get; private set; } = null!;

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Output("topicName")]
        public Output<string> TopicName { get; private set; } = null!;

        /// <summary>
        /// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
        /// </summary>
        [Output("trace")]
        public Output<bool?> Trace { get; private set; } = null!;


        /// <summary>
        /// Create a Topic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Topic(string name, TopicArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcmq/topic:Topic", name, args ?? new TopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Topic(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcmq/topic:Topic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Topic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Topic Get(string name, Input<string> id, TopicState? state = null, CustomResourceOptions? options = null)
        {
            return new Topic(name, id, state, options);
        }
    }

    public sealed class TopicArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to specify the message match policy for the topic. `1`: tag match policy (default value); `2`: routing match policy.
        /// </summary>
        [Input("filterType")]
        public Input<int>? FilterType { get; set; }

        /// <summary>
        /// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
        /// </summary>
        [Input("maxMsgSize")]
        public Input<int>? MaxMsgSize { get; set; }

        /// <summary>
        /// Message retention period. Value range: 60-86400 seconds (i.e., 1 minute-1 day). Default value: 86400.
        /// </summary>
        [Input("msgRetentionSeconds")]
        public Input<int>? MsgRetentionSeconds { get; set; }

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        /// <summary>
        /// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
        /// </summary>
        [Input("trace")]
        public Input<bool>? Trace { get; set; }

        public TopicArgs()
        {
        }
    }

    public sealed class TopicState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to specify the message match policy for the topic. `1`: tag match policy (default value); `2`: routing match policy.
        /// </summary>
        [Input("filterType")]
        public Input<int>? FilterType { get; set; }

        /// <summary>
        /// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
        /// </summary>
        [Input("maxMsgSize")]
        public Input<int>? MaxMsgSize { get; set; }

        /// <summary>
        /// Message retention period. Value range: 60-86400 seconds (i.e., 1 minute-1 day). Default value: 86400.
        /// </summary>
        [Input("msgRetentionSeconds")]
        public Input<int>? MsgRetentionSeconds { get; set; }

        /// <summary>
        /// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
        /// </summary>
        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        /// <summary>
        /// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
        /// </summary>
        [Input("trace")]
        public Input<bool>? Trace { get; set; }

        public TopicState()
        {
        }
    }
}
