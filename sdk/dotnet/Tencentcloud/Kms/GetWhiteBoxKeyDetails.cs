// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kms
{
    public static class GetWhiteBoxKeyDetails
    {
        /// <summary>
        /// Use this data source to query detailed information of kms white_box_key_details
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
        ///     var example = Tencentcloud.Kms.GetWhiteBoxKeyDetails.Invoke(new()
        ///     {
        ///         KeyStatus = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetWhiteBoxKeyDetailsResult> InvokeAsync(GetWhiteBoxKeyDetailsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWhiteBoxKeyDetailsResult>("tencentcloud:Kms/getWhiteBoxKeyDetails:getWhiteBoxKeyDetails", args ?? new GetWhiteBoxKeyDetailsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kms white_box_key_details
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
        ///     var example = Tencentcloud.Kms.GetWhiteBoxKeyDetails.Invoke(new()
        ///     {
        ///         KeyStatus = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetWhiteBoxKeyDetailsResult> Invoke(GetWhiteBoxKeyDetailsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWhiteBoxKeyDetailsResult>("tencentcloud:Kms/getWhiteBoxKeyDetails:getWhiteBoxKeyDetails", args ?? new GetWhiteBoxKeyDetailsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWhiteBoxKeyDetailsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter condition: status of the key, 0: disabled, 1: enabled.
        /// </summary>
        [Input("keyStatus")]
        public int? KeyStatus { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetWhiteBoxKeyDetailsArgs()
        {
        }
        public static new GetWhiteBoxKeyDetailsArgs Empty => new GetWhiteBoxKeyDetailsArgs();
    }

    public sealed class GetWhiteBoxKeyDetailsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter condition: status of the key, 0: disabled, 1: enabled.
        /// </summary>
        [Input("keyStatus")]
        public Input<int>? KeyStatus { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetWhiteBoxKeyDetailsInvokeArgs()
        {
        }
        public static new GetWhiteBoxKeyDetailsInvokeArgs Empty => new GetWhiteBoxKeyDetailsInvokeArgs();
    }


    [OutputType]
    public sealed class GetWhiteBoxKeyDetailsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of white box key information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetWhiteBoxKeyDetailsKeyInfoResult> KeyInfos;
        public readonly int? KeyStatus;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetWhiteBoxKeyDetailsResult(
            string id,

            ImmutableArray<Outputs.GetWhiteBoxKeyDetailsKeyInfoResult> keyInfos,

            int? keyStatus,

            string? resultOutputFile)
        {
            Id = id;
            KeyInfos = keyInfos;
            KeyStatus = keyStatus;
            ResultOutputFile = resultOutputFile;
        }
    }
}
