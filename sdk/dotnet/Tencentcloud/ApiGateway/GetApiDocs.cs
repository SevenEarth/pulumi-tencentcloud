// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    public static class GetApiDocs
    {
        /// <summary>
        /// Use this data source to query list information of api_gateway api_doc
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
        ///     var myApiDoc = Tencentcloud.ApiGateway.GetApiDocs.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetApiDocsResult> InvokeAsync(GetApiDocsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApiDocsResult>("tencentcloud:ApiGateway/getApiDocs:getApiDocs", args ?? new GetApiDocsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query list information of api_gateway api_doc
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
        ///     var myApiDoc = Tencentcloud.ApiGateway.GetApiDocs.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetApiDocsResult> Invoke(GetApiDocsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApiDocsResult>("tencentcloud:ApiGateway/getApiDocs:getApiDocs", args ?? new GetApiDocsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApiDocsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetApiDocsArgs()
        {
        }
        public static new GetApiDocsArgs Empty => new GetApiDocsArgs();
    }

    public sealed class GetApiDocsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetApiDocsInvokeArgs()
        {
        }
        public static new GetApiDocsInvokeArgs Empty => new GetApiDocsInvokeArgs();
    }


    [OutputType]
    public sealed class GetApiDocsResult
    {
        /// <summary>
        /// List of ApiDocs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetApiDocsApiDocListResult> ApiDocLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetApiDocsResult(
            ImmutableArray<Outputs.GetApiDocsApiDocListResult> apiDocLists,

            string id,

            string? resultOutputFile)
        {
            ApiDocLists = apiDocLists;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
