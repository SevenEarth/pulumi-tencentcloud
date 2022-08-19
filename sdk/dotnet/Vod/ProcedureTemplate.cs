// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod
{
    /// <summary>
    /// Provide a resource to create a VOD procedure template.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooAdaptiveDynamicStreamingTemplate = new Tencentcloud.Vod.AdaptiveDynamicStreamingTemplate("fooAdaptiveDynamicStreamingTemplate", new Tencentcloud.Vod.AdaptiveDynamicStreamingTemplateArgs
    ///         {
    ///             Format = "HLS",
    ///             DrmType = "SimpleAES",
    ///             DisableHigherVideoBitrate = false,
    ///             DisableHigherVideoResolution = false,
    ///             Comment = "test",
    ///             StreamInfos = 
    ///             {
    ///                 new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs
    ///                 {
    ///                     Video = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs
    ///                     {
    ///                         Codec = "libx265",
    ///                         Fps = 4,
    ///                         Bitrate = 129,
    ///                         ResolutionAdaptive = false,
    ///                         Width = 128,
    ///                         Height = 128,
    ///                         FillType = "stretch",
    ///                     },
    ///                     Audio = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs
    ///                     {
    ///                         Codec = "libmp3lame",
    ///                         Bitrate = 129,
    ///                         SampleRate = 44100,
    ///                         AudioChannel = "dual",
    ///                     },
    ///                     RemoveAudio = false,
    ///                 },
    ///                 new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoArgs
    ///                 {
    ///                     Video = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoVideoArgs
    ///                     {
    ///                         Codec = "libx264",
    ///                         Fps = 4,
    ///                         Bitrate = 256,
    ///                     },
    ///                     Audio = new Tencentcloud.Vod.Inputs.AdaptiveDynamicStreamingTemplateStreamInfoAudioArgs
    ///                     {
    ///                         Codec = "libfdk_aac",
    ///                         Bitrate = 256,
    ///                         SampleRate = 44100,
    ///                     },
    ///                     RemoveAudio = true,
    ///                 },
    ///             },
    ///         });
    ///         var fooSnapshotByTimeOffsetTemplate = new Tencentcloud.Vod.SnapshotByTimeOffsetTemplate("fooSnapshotByTimeOffsetTemplate", new Tencentcloud.Vod.SnapshotByTimeOffsetTemplateArgs
    ///         {
    ///             Width = 130,
    ///             Height = 128,
    ///             ResolutionAdaptive = false,
    ///             Format = "png",
    ///             Comment = "test",
    ///             FillType = "white",
    ///         });
    ///         var fooImageSpriteTemplate = new Tencentcloud.Vod.ImageSpriteTemplate("fooImageSpriteTemplate", new Tencentcloud.Vod.ImageSpriteTemplateArgs
    ///         {
    ///             SampleType = "Percent",
    ///             SampleInterval = 10,
    ///             RowCount = 3,
    ///             ColumnCount = 3,
    ///             Comment = "test",
    ///             FillType = "stretch",
    ///             Width = 128,
    ///             Height = 128,
    ///             ResolutionAdaptive = false,
    ///         });
    ///         var fooProcedureTemplate = new Tencentcloud.Vod.ProcedureTemplate("fooProcedureTemplate", new Tencentcloud.Vod.ProcedureTemplateArgs
    ///         {
    ///             Comment = "test",
    ///             MediaProcessTask = new Tencentcloud.Vod.Inputs.ProcedureTemplateMediaProcessTaskArgs
    ///             {
    ///                 AdaptiveDynamicStreamingTaskLists = 
    ///                 {
    ///                     new Tencentcloud.Vod.Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListArgs
    ///                     {
    ///                         Definition = fooAdaptiveDynamicStreamingTemplate.Id,
    ///                     },
    ///                 },
    ///                 SnapshotByTimeOffsetTaskLists = 
    ///                 {
    ///                     new Tencentcloud.Vod.Inputs.ProcedureTemplateMediaProcessTaskSnapshotByTimeOffsetTaskListArgs
    ///                     {
    ///                         Definition = fooSnapshotByTimeOffsetTemplate.Id,
    ///                         ExtTimeOffsetLists = 
    ///                         {
    ///                             "3.5s",
    ///                         },
    ///                     },
    ///                 },
    ///                 ImageSpriteTaskLists = 
    ///                 {
    ///                     new Tencentcloud.Vod.Inputs.ProcedureTemplateMediaProcessTaskImageSpriteTaskListArgs
    ///                     {
    ///                         Definition = fooImageSpriteTemplate.Id,
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VOD procedure template can be imported using the name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vod/procedureTemplate:ProcedureTemplate foo tf-procedure
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vod/procedureTemplate:ProcedureTemplate")]
    public partial class ProcedureTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Creation time of template in ISO date format.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Parameter of video processing task.
        /// </summary>
        [Output("mediaProcessTask")]
        public Output<Outputs.ProcedureTemplateMediaProcessTask?> MediaProcessTask { get; private set; } = null!;

        /// <summary>
        /// Task flow name (up to 20 characters).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Output("subAppId")]
        public Output<int?> SubAppId { get; private set; } = null!;

        /// <summary>
        /// Last modified time of template in ISO date format.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ProcedureTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProcedureTemplate(string name, ProcedureTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/procedureTemplate:ProcedureTemplate", name, args ?? new ProcedureTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProcedureTemplate(string name, Input<string> id, ProcedureTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/procedureTemplate:ProcedureTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProcedureTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProcedureTemplate Get(string name, Input<string> id, ProcedureTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ProcedureTemplate(name, id, state, options);
        }
    }

    public sealed class ProcedureTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Parameter of video processing task.
        /// </summary>
        [Input("mediaProcessTask")]
        public Input<Inputs.ProcedureTemplateMediaProcessTaskArgs>? MediaProcessTask { get; set; }

        /// <summary>
        /// Task flow name (up to 20 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        public ProcedureTemplateArgs()
        {
        }
    }

    public sealed class ProcedureTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template description. Length limit: 256 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Creation time of template in ISO date format.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Parameter of video processing task.
        /// </summary>
        [Input("mediaProcessTask")]
        public Input<Inputs.ProcedureTemplateMediaProcessTaskGetArgs>? MediaProcessTask { get; set; }

        /// <summary>
        /// Task flow name (up to 20 characters).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        /// <summary>
        /// Last modified time of template in ISO date format.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ProcedureTemplateState()
        {
        }
    }
}
