// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssm
{
    public static class GetRotationHistory
    {
        /// <summary>
        /// Use this data source to query detailed information of ssm rotation_history
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
        ///         var example = Output.Create(Tencentcloud.Ssm.GetRotationHistory.InvokeAsync(new Tencentcloud.Ssm.GetRotationHistoryArgs
        ///         {
        ///             SecretName = "keep_terraform",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRotationHistoryResult> InvokeAsync(GetRotationHistoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRotationHistoryResult>("tencentcloud:Ssm/getRotationHistory:getRotationHistory", args ?? new GetRotationHistoryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssm rotation_history
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
        ///         var example = Output.Create(Tencentcloud.Ssm.GetRotationHistory.InvokeAsync(new Tencentcloud.Ssm.GetRotationHistoryArgs
        ///         {
        ///             SecretName = "keep_terraform",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRotationHistoryResult> Invoke(GetRotationHistoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRotationHistoryResult>("tencentcloud:Ssm/getRotationHistory:getRotationHistory", args ?? new GetRotationHistoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRotationHistoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Secret name.
        /// </summary>
        [Input("secretName", required: true)]
        public string SecretName { get; set; } = null!;

        public GetRotationHistoryArgs()
        {
        }
    }

    public sealed class GetRotationHistoryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Secret name.
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        public GetRotationHistoryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRotationHistoryResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string SecretName;
        /// <summary>
        /// The number of version numbers. The maximum number of version numbers that can be displayed to users is 10.
        /// </summary>
        public readonly ImmutableArray<string> VersionIds;

        [OutputConstructor]
        private GetRotationHistoryResult(
            string id,

            string? resultOutputFile,

            string secretName,

            ImmutableArray<string> versionIds)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretName = secretName;
            VersionIds = versionIds;
        }
    }
}
