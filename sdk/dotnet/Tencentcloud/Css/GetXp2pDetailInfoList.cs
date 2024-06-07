// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    public static class GetXp2pDetailInfoList
    {
        /// <summary>
        /// Use this data source to query detailed information of css xp2p_detail_info_list
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var xp2pDetailInfoList = Tencentcloud.Css.GetXp2pDetailInfoList.Invoke(new()
        ///     {
        ///         QueryTime = "2023-11-01T14:55:01+08:00",
        ///         Types = new[]
        ///         {
        ///             "live",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetXp2pDetailInfoListResult> InvokeAsync(GetXp2pDetailInfoListArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetXp2pDetailInfoListResult>("tencentcloud:Css/getXp2pDetailInfoList:getXp2pDetailInfoList", args ?? new GetXp2pDetailInfoListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of css xp2p_detail_info_list
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var xp2pDetailInfoList = Tencentcloud.Css.GetXp2pDetailInfoList.Invoke(new()
        ///     {
        ///         QueryTime = "2023-11-01T14:55:01+08:00",
        ///         Types = new[]
        ///         {
        ///             "live",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetXp2pDetailInfoListResult> Invoke(GetXp2pDetailInfoListInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetXp2pDetailInfoListResult>("tencentcloud:Css/getXp2pDetailInfoList:getXp2pDetailInfoList", args ?? new GetXp2pDetailInfoListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetXp2pDetailInfoListArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimensions")]
        private List<string>? _dimensions;

        /// <summary>
        /// The dimension parameter can be used to specify the dimension for the query. If this parameter is not passed, the query will default to stream-level data. If you pass this parameter, it will only retrieve data for the specified dimension. The available dimension currently supported is AppId dimension, which allows you to query data based on the application ID. Please note that the returned fields will be related to the specified dimension.
        /// </summary>
        public List<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new List<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The UTC minute granularity query time for querying usage data for a specific minute is in the format: yyyy-mm-ddTHH:MM:00Z. Please refer to the link https://cloud.tencent.com/document/product/266/11732#I.For example, if the local time is 2019-01-08 10:00:00 in Beijing, the corresponding UTC time would be 2019-01-08T10:00:00+08:00.This query supports data from the past six months.
        /// </summary>
        [Input("queryTime")]
        public string? QueryTime { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("streamNames")]
        private List<string>? _streamNames;

        /// <summary>
        /// The stream array can be used to specify the streams to be queried. If no stream is specified, the query will include all streams by default.
        /// </summary>
        public List<string> StreamNames
        {
            get => _streamNames ?? (_streamNames = new List<string>());
            set => _streamNames = value;
        }

        [Input("types")]
        private List<string>? _types;

        /// <summary>
        /// The type array can be used to specify the type of media content to be queried. The two available options are live for live streaming and vod for video on demand. If no type is specified, the query will include both live and VOD content by default.
        /// </summary>
        public List<string> Types
        {
            get => _types ?? (_types = new List<string>());
            set => _types = value;
        }

        public GetXp2pDetailInfoListArgs()
        {
        }
        public static new GetXp2pDetailInfoListArgs Empty => new GetXp2pDetailInfoListArgs();
    }

    public sealed class GetXp2pDetailInfoListInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dimensions")]
        private InputList<string>? _dimensions;

        /// <summary>
        /// The dimension parameter can be used to specify the dimension for the query. If this parameter is not passed, the query will default to stream-level data. If you pass this parameter, it will only retrieve data for the specified dimension. The available dimension currently supported is AppId dimension, which allows you to query data based on the application ID. Please note that the returned fields will be related to the specified dimension.
        /// </summary>
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// The UTC minute granularity query time for querying usage data for a specific minute is in the format: yyyy-mm-ddTHH:MM:00Z. Please refer to the link https://cloud.tencent.com/document/product/266/11732#I.For example, if the local time is 2019-01-08 10:00:00 in Beijing, the corresponding UTC time would be 2019-01-08T10:00:00+08:00.This query supports data from the past six months.
        /// </summary>
        [Input("queryTime")]
        public Input<string>? QueryTime { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("streamNames")]
        private InputList<string>? _streamNames;

        /// <summary>
        /// The stream array can be used to specify the streams to be queried. If no stream is specified, the query will include all streams by default.
        /// </summary>
        public InputList<string> StreamNames
        {
            get => _streamNames ?? (_streamNames = new InputList<string>());
            set => _streamNames = value;
        }

        [Input("types")]
        private InputList<string>? _types;

        /// <summary>
        /// The type array can be used to specify the type of media content to be queried. The two available options are live for live streaming and vod for video on demand. If no type is specified, the query will include both live and VOD content by default.
        /// </summary>
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        public GetXp2pDetailInfoListInvokeArgs()
        {
        }
        public static new GetXp2pDetailInfoListInvokeArgs Empty => new GetXp2pDetailInfoListInvokeArgs();
    }


    [OutputType]
    public sealed class GetXp2pDetailInfoListResult
    {
        /// <summary>
        /// P2P streaming statistical information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetXp2pDetailInfoListDataInfoListResult> DataInfoLists;
        public readonly ImmutableArray<string> Dimensions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? QueryTime;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<string> StreamNames;
        /// <summary>
        /// Type, divided into two categories: live and vod.Note: This field may return null, indicating that no valid value is available.
        /// </summary>
        public readonly ImmutableArray<string> Types;

        [OutputConstructor]
        private GetXp2pDetailInfoListResult(
            ImmutableArray<Outputs.GetXp2pDetailInfoListDataInfoListResult> dataInfoLists,

            ImmutableArray<string> dimensions,

            string id,

            string? queryTime,

            string? resultOutputFile,

            ImmutableArray<string> streamNames,

            ImmutableArray<string> types)
        {
            DataInfoLists = dataInfoLists;
            Dimensions = dimensions;
            Id = id;
            QueryTime = queryTime;
            ResultOutputFile = resultOutputFile;
            StreamNames = streamNames;
            Types = types;
        }
    }
}
