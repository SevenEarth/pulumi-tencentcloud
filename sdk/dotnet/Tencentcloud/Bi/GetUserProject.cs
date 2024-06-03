// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Bi
{
    public static class GetUserProject
    {
        /// <summary>
        /// Use this data source to query detailed information of bi user_project
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
        ///     var userProject = Tencentcloud.Bi.GetUserProject.Invoke(new()
        ///     {
        ///         AllPage = true,
        ///         ProjectId = 123,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUserProjectResult> InvokeAsync(GetUserProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserProjectResult>("tencentcloud:Bi/getUserProject:getUserProject", args ?? new GetUserProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of bi user_project
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
        ///     var userProject = Tencentcloud.Bi.GetUserProject.Invoke(new()
        ///     {
        ///         AllPage = true,
        ///         ProjectId = 123,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUserProjectResult> Invoke(GetUserProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserProjectResult>("tencentcloud:Bi/getUserProject:getUserProject", args ?? new GetUserProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to display all, if true, ignore paging.
        /// </summary>
        [Input("allPage")]
        public bool? AllPage { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetUserProjectArgs()
        {
        }
        public static new GetUserProjectArgs Empty => new GetUserProjectArgs();
    }

    public sealed class GetUserProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to display all, if true, ignore paging.
        /// </summary>
        [Input("allPage")]
        public Input<bool>? AllPage { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetUserProjectInvokeArgs()
        {
        }
        public static new GetUserProjectInvokeArgs Empty => new GetUserProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserProjectResult
    {
        public readonly bool? AllPage;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Array(Note: This field may return null, indicating that no valid value can be obtained).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserProjectListResult> Lists;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetUserProjectResult(
            bool? allPage,

            string id,

            ImmutableArray<Outputs.GetUserProjectListResult> lists,

            int? projectId,

            string? resultOutputFile)
        {
            AllPage = allPage;
            Id = id;
            Lists = lists;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
