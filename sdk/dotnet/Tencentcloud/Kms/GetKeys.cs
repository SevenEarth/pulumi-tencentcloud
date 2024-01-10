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
    public static class GetKeys
    {
        /// <summary>
        /// Use this data source to query detailed information of KMS key
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Tencentcloud.Kms.GetKeys.InvokeAsync(new Tencentcloud.Kms.GetKeysArgs
        ///         {
        ///             KeyState = 0,
        ///             KeyUsage = "ALL",
        ///             Origin = "TENCENT_KMS",
        ///             SearchKeyAlias = "tf_example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKeysResult> InvokeAsync(GetKeysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeysResult>("tencentcloud:Kms/getKeys:getKeys", args ?? new GetKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of KMS key
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Tencentcloud.Kms.GetKeys.InvokeAsync(new Tencentcloud.Kms.GetKeysArgs
        ///         {
        ///             KeyState = 0,
        ///             KeyUsage = "ALL",
        ///             Origin = "TENCENT_KMS",
        ///             SearchKeyAlias = "tf_example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetKeysResult> Invoke(GetKeysInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetKeysResult>("tencentcloud:Kms/getKeys:getKeys", args ?? new GetKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter by state of CMK. `0` - all CMKs are queried, `1` - only Enabled CMKs are queried, `2` - only Disabled CMKs are queried, `3` - only PendingDelete CMKs are queried, `4` - only PendingImport CMKs are queried, `5` - only Archived CMKs are queried.
        /// </summary>
        [Input("keyState")]
        public int? KmsKeyState { get; set; }

        /// <summary>
        /// Filter by usage of CMK. Available values include `ALL`, `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        /// </summary>
        [Input("keyUsage")]
        public string? KeyUsage { get; set; }

        /// <summary>
        /// Order to sort the CMK create time. `0` - desc, `1` - asc. Default value is `0`.
        /// </summary>
        [Input("orderType")]
        public int? OrderType { get; set; }

        /// <summary>
        /// Filter by origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user, `ALL` - all CMKs. Default value is `ALL`.
        /// </summary>
        [Input("origin")]
        public string? Origin { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Filter by role of the CMK creator. `0` - created by user, `1` - created by cloud product. Default value is `0`.
        /// </summary>
        [Input("role")]
        public int? Role { get; set; }

        /// <summary>
        /// Words used to match the results, and the words can be: key_id and alias.
        /// </summary>
        [Input("searchKeyAlias")]
        public string? SearchKeyAlias { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// Tags to filter CMK.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetKeysArgs()
        {
        }
    }

    public sealed class GetKeysInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter by state of CMK. `0` - all CMKs are queried, `1` - only Enabled CMKs are queried, `2` - only Disabled CMKs are queried, `3` - only PendingDelete CMKs are queried, `4` - only PendingImport CMKs are queried, `5` - only Archived CMKs are queried.
        /// </summary>
        [Input("keyState")]
        public Input<int>? KeyState { get; set; }

        /// <summary>
        /// Filter by usage of CMK. Available values include `ALL`, `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        /// </summary>
        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        /// <summary>
        /// Order to sort the CMK create time. `0` - desc, `1` - asc. Default value is `0`.
        /// </summary>
        [Input("orderType")]
        public Input<int>? OrderType { get; set; }

        /// <summary>
        /// Filter by origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user, `ALL` - all CMKs. Default value is `ALL`.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Filter by role of the CMK creator. `0` - created by user, `1` - created by cloud product. Default value is `0`.
        /// </summary>
        [Input("role")]
        public Input<int>? Role { get; set; }

        /// <summary>
        /// Words used to match the results, and the words can be: key_id and alias.
        /// </summary>
        [Input("searchKeyAlias")]
        public Input<string>? SearchKeyAlias { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags to filter CMK.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetKeysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of KMS keys.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKeysKeyListResult> KeyLists;
        /// <summary>
        /// State of CMK.
        /// </summary>
        public readonly int? KmsKeyState;
        /// <summary>
        /// Usage of CMK.
        /// </summary>
        public readonly string? KeyUsage;
        public readonly int? OrderType;
        /// <summary>
        /// Origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user.
        /// </summary>
        public readonly string? Origin;
        public readonly string? ResultOutputFile;
        public readonly int? Role;
        public readonly string? SearchKeyAlias;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetKeysResult(
            string id,

            ImmutableArray<Outputs.GetKeysKeyListResult> keyLists,

            int? keyState,

            string? keyUsage,

            int? orderType,

            string? origin,

            string? resultOutputFile,

            int? role,

            string? searchKeyAlias,

            ImmutableDictionary<string, object>? tags)
        {
            Id = id;
            KeyLists = keyLists;
            KmsKeyState = keyState;
            KeyUsage = keyUsage;
            OrderType = orderType;
            Origin = origin;
            ResultOutputFile = resultOutputFile;
            Role = role;
            SearchKeyAlias = searchKeyAlias;
            Tags = tags;
        }
    }
}
