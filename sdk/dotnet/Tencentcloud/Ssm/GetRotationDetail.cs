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
    public static class GetRotationDetail
    {
        /// <summary>
        /// Use this data source to query detailed information of ssm rotation_detail
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
        ///         var example = Output.Create(Tencentcloud.Ssm.GetRotationDetail.InvokeAsync(new Tencentcloud.Ssm.GetRotationDetailArgs
        ///         {
        ///             SecretName = "tf_example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRotationDetailResult> InvokeAsync(GetRotationDetailArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRotationDetailResult>("tencentcloud:Ssm/getRotationDetail:getRotationDetail", args ?? new GetRotationDetailArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssm rotation_detail
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
        ///         var example = Output.Create(Tencentcloud.Ssm.GetRotationDetail.InvokeAsync(new Tencentcloud.Ssm.GetRotationDetailArgs
        ///         {
        ///             SecretName = "tf_example",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRotationDetailResult> Invoke(GetRotationDetailInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRotationDetailResult>("tencentcloud:Ssm/getRotationDetail:getRotationDetail", args ?? new GetRotationDetailInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRotationDetailArgs : Pulumi.InvokeArgs
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

        public GetRotationDetailArgs()
        {
        }
    }

    public sealed class GetRotationDetailInvokeArgs : Pulumi.InvokeArgs
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

        public GetRotationDetailInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRotationDetailResult
    {
        /// <summary>
        /// Whether to allow rotation.
        /// </summary>
        public readonly bool EnableRotation;
        /// <summary>
        /// The rotation frequency, in days, defaults to 1 day.
        /// </summary>
        public readonly int Frequency;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Time of last rotation.
        /// </summary>
        public readonly string LatestRotateTime;
        /// <summary>
        /// The time to start the next rotation.
        /// </summary>
        public readonly string NextRotateBeginTime;
        public readonly string? ResultOutputFile;
        public readonly string SecretName;

        [OutputConstructor]
        private GetRotationDetailResult(
            bool enableRotation,

            int frequency,

            string id,

            string latestRotateTime,

            string nextRotateBeginTime,

            string? resultOutputFile,

            string secretName)
        {
            EnableRotation = enableRotation;
            Frequency = frequency;
            Id = id;
            LatestRotateTime = latestRotateTime;
            NextRotateBeginTime = nextRotateBeginTime;
            ResultOutputFile = resultOutputFile;
            SecretName = secretName;
        }
    }
}