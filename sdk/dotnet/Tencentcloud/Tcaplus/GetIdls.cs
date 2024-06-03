// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcaplus
{
    public static class GetIdls
    {
        /// <summary>
        /// Use this data source to query  IDL information of the TcaplusDB table.
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
        ///     var idTest = Tencentcloud.Tcaplus.GetIdls.Invoke(new()
        ///     {
        ///         ClusterId = "19162256624",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetIdlsResult> InvokeAsync(GetIdlsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdlsResult>("tencentcloud:Tcaplus/getIdls:getIdls", args ?? new GetIdlsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query  IDL information of the TcaplusDB table.
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
        ///     var idTest = Tencentcloud.Tcaplus.GetIdls.Invoke(new()
        ///     {
        ///         ClusterId = "19162256624",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetIdlsResult> Invoke(GetIdlsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdlsResult>("tencentcloud:Tcaplus/getIdls:getIdls", args ?? new GetIdlsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdlsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the TcaplusDB cluster to be query.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// File for saving results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetIdlsArgs()
        {
        }
        public static new GetIdlsArgs Empty => new GetIdlsArgs();
    }

    public sealed class GetIdlsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the TcaplusDB cluster to be query.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// File for saving results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetIdlsInvokeArgs()
        {
        }
        public static new GetIdlsInvokeArgs Empty => new GetIdlsInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdlsResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of TcaplusDB table IDL. Each element contains the following attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIdlsListResult> Lists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetIdlsResult(
            string clusterId,

            string id,

            ImmutableArray<Outputs.GetIdlsListResult> lists,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
