// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf
{
    public static class GetUserClbRegions
    {
        /// <summary>
        /// Use this data source to query detailed information of waf user_clb_regions
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
        ///     var example = Tencentcloud.Waf.GetUserClbRegions.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUserClbRegionsResult> InvokeAsync(GetUserClbRegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserClbRegionsResult>("tencentcloud:Waf/getUserClbRegions:getUserClbRegions", args ?? new GetUserClbRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of waf user_clb_regions
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
        ///     var example = Tencentcloud.Waf.GetUserClbRegions.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUserClbRegionsResult> Invoke(GetUserClbRegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserClbRegionsResult>("tencentcloud:Waf/getUserClbRegions:getUserClbRegions", args ?? new GetUserClbRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserClbRegionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetUserClbRegionsArgs()
        {
        }
        public static new GetUserClbRegionsArgs Empty => new GetUserClbRegionsArgs();
    }

    public sealed class GetUserClbRegionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetUserClbRegionsInvokeArgs()
        {
        }
        public static new GetUserClbRegionsInvokeArgs Empty => new GetUserClbRegionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserClbRegionsResult
    {
        /// <summary>
        /// Region list(ap-xxx format).
        /// </summary>
        public readonly ImmutableArray<string> Datas;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Detail info for region.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserClbRegionsRichDataResult> RichDatas;

        [OutputConstructor]
        private GetUserClbRegionsResult(
            ImmutableArray<string> datas,

            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetUserClbRegionsRichDataResult> richDatas)
        {
            Datas = datas;
            Id = id;
            ResultOutputFile = resultOutputFile;
            RichDatas = richDatas;
        }
    }
}
