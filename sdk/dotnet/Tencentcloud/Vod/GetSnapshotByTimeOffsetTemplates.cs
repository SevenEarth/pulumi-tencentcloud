// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod
{
    public static class GetSnapshotByTimeOffsetTemplates
    {
        /// <summary>
        /// Use this data source to query detailed information of VOD snapshot by time offset templates.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooSnapshotByTimeOffsetTemplate = new Tencentcloud.Vod.SnapshotByTimeOffsetTemplate("fooSnapshotByTimeOffsetTemplate", new()
        ///     {
        ///         Width = 130,
        ///         Height = 128,
        ///         ResolutionAdaptive = false,
        ///         Format = "png",
        ///         Comment = "test",
        ///         FillType = "white",
        ///     });
        /// 
        ///     var fooSnapshotByTimeOffsetTemplates = Tencentcloud.Vod.GetSnapshotByTimeOffsetTemplates.Invoke(new()
        ///     {
        ///         Type = "Custom",
        ///         Definition = fooSnapshotByTimeOffsetTemplate.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSnapshotByTimeOffsetTemplatesResult> InvokeAsync(GetSnapshotByTimeOffsetTemplatesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotByTimeOffsetTemplatesResult>("tencentcloud:Vod/getSnapshotByTimeOffsetTemplates:getSnapshotByTimeOffsetTemplates", args ?? new GetSnapshotByTimeOffsetTemplatesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of VOD snapshot by time offset templates.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fooSnapshotByTimeOffsetTemplate = new Tencentcloud.Vod.SnapshotByTimeOffsetTemplate("fooSnapshotByTimeOffsetTemplate", new()
        ///     {
        ///         Width = 130,
        ///         Height = 128,
        ///         ResolutionAdaptive = false,
        ///         Format = "png",
        ///         Comment = "test",
        ///         FillType = "white",
        ///     });
        /// 
        ///     var fooSnapshotByTimeOffsetTemplates = Tencentcloud.Vod.GetSnapshotByTimeOffsetTemplates.Invoke(new()
        ///     {
        ///         Type = "Custom",
        ///         Definition = fooSnapshotByTimeOffsetTemplate.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSnapshotByTimeOffsetTemplatesResult> Invoke(GetSnapshotByTimeOffsetTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotByTimeOffsetTemplatesResult>("tencentcloud:Vod/getSnapshotByTimeOffsetTemplates:getSnapshotByTimeOffsetTemplates", args ?? new GetSnapshotByTimeOffsetTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotByTimeOffsetTemplatesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique ID filter of snapshot by time offset template.
        /// </summary>
        [Input("definition")]
        public string? Definition { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public int? SubAppId { get; set; }

        /// <summary>
        /// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetSnapshotByTimeOffsetTemplatesArgs()
        {
        }
        public static new GetSnapshotByTimeOffsetTemplatesArgs Empty => new GetSnapshotByTimeOffsetTemplatesArgs();
    }

    public sealed class GetSnapshotByTimeOffsetTemplatesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique ID filter of snapshot by time offset template.
        /// </summary>
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Subapplication ID in VOD. If you need to access a resource in a subapplication, enter the subapplication ID in this field; otherwise, leave it empty.
        /// </summary>
        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        /// <summary>
        /// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetSnapshotByTimeOffsetTemplatesInvokeArgs()
        {
        }
        public static new GetSnapshotByTimeOffsetTemplatesInvokeArgs Empty => new GetSnapshotByTimeOffsetTemplatesInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotByTimeOffsetTemplatesResult
    {
        /// <summary>
        /// Unique ID of snapshot by time offset template.
        /// </summary>
        public readonly string? Definition;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        /// <summary>
        /// A list of snapshot by time offset templates. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSnapshotByTimeOffsetTemplatesTemplateListResult> TemplateLists;
        /// <summary>
        /// Template type filter. Valid values: `Preset`, `Custom`. `Preset`: preset template; `Custom`: custom template.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private GetSnapshotByTimeOffsetTemplatesResult(
            string? definition,

            string id,

            string? resultOutputFile,

            int? subAppId,

            ImmutableArray<Outputs.GetSnapshotByTimeOffsetTemplatesTemplateListResult> templateLists,

            string? type)
        {
            Definition = definition;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            TemplateLists = templateLists;
            Type = type;
        }
    }
}
