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
    public static class GetDescribeKeys
    {
        /// <summary>
        /// Use this data source to query detailed information of kms key_lists
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
        ///     var example = Tencentcloud.Kms.GetDescribeKeys.Invoke(new()
        ///     {
        ///         KeyIds = new[]
        ///         {
        ///             "9ffacc8b-6461-11ee-a54e-525400dd8a7d",
        ///             "bffae4ed-6465-11ee-90b2-5254000ef00e",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDescribeKeysResult> InvokeAsync(GetDescribeKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeKeysResult>("tencentcloud:Kms/getDescribeKeys:getDescribeKeys", args ?? new GetDescribeKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of kms key_lists
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
        ///     var example = Tencentcloud.Kms.GetDescribeKeys.Invoke(new()
        ///     {
        ///         KeyIds = new[]
        ///         {
        ///             "9ffacc8b-6461-11ee-a54e-525400dd8a7d",
        ///             "bffae4ed-6465-11ee-90b2-5254000ef00e",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDescribeKeysResult> Invoke(GetDescribeKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeKeysResult>("tencentcloud:Kms/getDescribeKeys:getDescribeKeys", args ?? new GetDescribeKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeKeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyIds", required: true)]
        private List<string>? _keyIds;

        /// <summary>
        /// Query the ID list of CMK, batch query supports up to 100 KeyIds at a time.
        /// </summary>
        public List<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new List<string>());
            set => _keyIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeKeysArgs()
        {
        }
        public static new GetDescribeKeysArgs Empty => new GetDescribeKeysArgs();
    }

    public sealed class GetDescribeKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("keyIds", required: true)]
        private InputList<string>? _keyIds;

        /// <summary>
        /// Query the ID list of CMK, batch query supports up to 100 KeyIds at a time.
        /// </summary>
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeKeysInvokeArgs()
        {
        }
        public static new GetDescribeKeysInvokeArgs Empty => new GetDescribeKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> KeyIds;
        /// <summary>
        /// A list of KMS keys.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeKeysKeyListResult> KeyLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeKeysResult(
            string id,

            ImmutableArray<string> keyIds,

            ImmutableArray<Outputs.GetDescribeKeysKeyListResult> keyLists,

            string? resultOutputFile)
        {
            Id = id;
            KeyIds = keyIds;
            KeyLists = keyLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
