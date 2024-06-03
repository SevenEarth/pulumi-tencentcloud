// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Rum
{
    public static class GetOfflineLogConfig
    {
        /// <summary>
        /// Use this data source to query detailed information of rum offlineLogConfig
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
        ///     var offlineLogConfig = Tencentcloud.Rum.GetOfflineLogConfig.Invoke(new()
        ///     {
        ///         ProjectKey = "ZEYrYfvaYQ30jRdmPx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetOfflineLogConfigResult> InvokeAsync(GetOfflineLogConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOfflineLogConfigResult>("tencentcloud:Rum/getOfflineLogConfig:getOfflineLogConfig", args ?? new GetOfflineLogConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rum offlineLogConfig
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
        ///     var offlineLogConfig = Tencentcloud.Rum.GetOfflineLogConfig.Invoke(new()
        ///     {
        ///         ProjectKey = "ZEYrYfvaYQ30jRdmPx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetOfflineLogConfigResult> Invoke(GetOfflineLogConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOfflineLogConfigResult>("tencentcloud:Rum/getOfflineLogConfig:getOfflineLogConfig", args ?? new GetOfflineLogConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOfflineLogConfigArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique project key for reporting.
        /// </summary>
        [Input("projectKey", required: true)]
        public string ProjectKey { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetOfflineLogConfigArgs()
        {
        }
        public static new GetOfflineLogConfigArgs Empty => new GetOfflineLogConfigArgs();
    }

    public sealed class GetOfflineLogConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique project key for reporting.
        /// </summary>
        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetOfflineLogConfigInvokeArgs()
        {
        }
        public static new GetOfflineLogConfigInvokeArgs Empty => new GetOfflineLogConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetOfflineLogConfigResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// API call information.
        /// </summary>
        public readonly string Msg;
        public readonly string ProjectKey;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Unique identifier of the user to be listened on(aid or uin).
        /// </summary>
        public readonly ImmutableArray<string> UniqueIdSets;

        [OutputConstructor]
        private GetOfflineLogConfigResult(
            string id,

            string msg,

            string projectKey,

            string? resultOutputFile,

            ImmutableArray<string> uniqueIdSets)
        {
            Id = id;
            Msg = msg;
            ProjectKey = projectKey;
            ResultOutputFile = resultOutputFile;
            UniqueIdSets = uniqueIdSets;
        }
    }
}
