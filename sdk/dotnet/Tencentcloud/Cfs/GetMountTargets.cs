// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cfs
{
    public static class GetMountTargets
    {
        /// <summary>
        /// Use this data source to query detailed information of cfs mount_targets
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
        ///         var mountTargets = Output.Create(Tencentcloud.Cfs.GetMountTargets.InvokeAsync(new Tencentcloud.Cfs.GetMountTargetsArgs
        ///         {
        ///             FileSystemId = "cfs-iobiaxtj",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMountTargetsResult> InvokeAsync(GetMountTargetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMountTargetsResult>("tencentcloud:Cfs/getMountTargets:getMountTargets", args ?? new GetMountTargetsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cfs mount_targets
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
        ///         var mountTargets = Output.Create(Tencentcloud.Cfs.GetMountTargets.InvokeAsync(new Tencentcloud.Cfs.GetMountTargetsArgs
        ///         {
        ///             FileSystemId = "cfs-iobiaxtj",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMountTargetsResult> Invoke(GetMountTargetsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMountTargetsResult>("tencentcloud:Cfs/getMountTargets:getMountTargets", args ?? new GetMountTargetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMountTargetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// File system ID.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public string FileSystemId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetMountTargetsArgs()
        {
        }
    }

    public sealed class GetMountTargetsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// File system ID.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetMountTargetsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMountTargetsResult
    {
        /// <summary>
        /// File system ID.
        /// </summary>
        public readonly string FileSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Mount target details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMountTargetsMountTargetResult> CfsMountTargets;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetMountTargetsResult(
            string fileSystemId,

            string id,

            ImmutableArray<Outputs.GetMountTargetsMountTargetResult> mountTargets,

            string? resultOutputFile)
        {
            FileSystemId = fileSystemId;
            Id = id;
            CfsMountTargets = mountTargets;
            ResultOutputFile = resultOutputFile;
        }
    }
}
