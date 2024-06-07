// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eip
{
    public static class GetNetworkAccountType
    {
        /// <summary>
        /// Use this data source to query detailed information of eip network_account_type
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
        ///     var networkAccountType = Tencentcloud.Eip.GetNetworkAccountType.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetNetworkAccountTypeResult> InvokeAsync(GetNetworkAccountTypeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkAccountTypeResult>("tencentcloud:Eip/getNetworkAccountType:getNetworkAccountType", args ?? new GetNetworkAccountTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of eip network_account_type
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
        ///     var networkAccountType = Tencentcloud.Eip.GetNetworkAccountType.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetNetworkAccountTypeResult> Invoke(GetNetworkAccountTypeInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkAccountTypeResult>("tencentcloud:Eip/getNetworkAccountType:getNetworkAccountType", args ?? new GetNetworkAccountTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkAccountTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetNetworkAccountTypeArgs()
        {
        }
        public static new GetNetworkAccountTypeArgs Empty => new GetNetworkAccountTypeArgs();
    }

    public sealed class GetNetworkAccountTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetNetworkAccountTypeInvokeArgs()
        {
        }
        public static new GetNetworkAccountTypeInvokeArgs Empty => new GetNetworkAccountTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkAccountTypeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The network type of the user account, STANDARD is a standard user, LEGACY is a traditional user.
        /// </summary>
        public readonly string EipNetworkAccountType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetNetworkAccountTypeResult(
            string id,

            string networkAccountType,

            string? resultOutputFile)
        {
            Id = id;
            EipNetworkAccountType = networkAccountType;
            ResultOutputFile = resultOutputFile;
        }
    }
}
