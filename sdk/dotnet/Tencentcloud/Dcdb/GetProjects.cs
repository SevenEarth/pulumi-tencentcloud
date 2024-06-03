// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    public static class GetProjects
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb projects
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
        ///     var projects = Tencentcloud.Dcdb.GetProjects.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetProjectsResult> InvokeAsync(GetProjectsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectsResult>("tencentcloud:Dcdb/getProjects:getProjects", args ?? new GetProjectsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb projects
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
        ///     var projects = Tencentcloud.Dcdb.GetProjects.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetProjectsResult> Invoke(GetProjectsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectsResult>("tencentcloud:Dcdb/getProjects:getProjects", args ?? new GetProjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProjectsArgs()
        {
        }
        public static new GetProjectsArgs Empty => new GetProjectsArgs();
    }

    public sealed class GetProjectsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProjectsInvokeArgs()
        {
        }
        public static new GetProjectsInvokeArgs Empty => new GetProjectsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Project list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectsProjectResult> DcdbProjects;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetProjectsResult(
            string id,

            ImmutableArray<Outputs.GetProjectsProjectResult> projects,

            string? resultOutputFile)
        {
            Id = id;
            DcdbProjects = projects;
            ResultOutputFile = resultOutputFile;
        }
    }
}
