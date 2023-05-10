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
    public static class GetFileSystemClients
    {
        /// <summary>
        /// Use this data source to query detailed information of cfs file_system_clients
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
        ///         var fileSystemClients = Output.Create(Tencentcloud.Cfs.GetFileSystemClients.InvokeAsync(new Tencentcloud.Cfs.GetFileSystemClientsArgs
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
        public static Task<GetFileSystemClientsResult> InvokeAsync(GetFileSystemClientsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemClientsResult>("tencentcloud:Cfs/getFileSystemClients:getFileSystemClients", args ?? new GetFileSystemClientsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cfs file_system_clients
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
        ///         var fileSystemClients = Output.Create(Tencentcloud.Cfs.GetFileSystemClients.InvokeAsync(new Tencentcloud.Cfs.GetFileSystemClientsArgs
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
        public static Output<GetFileSystemClientsResult> Invoke(GetFileSystemClientsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFileSystemClientsResult>("tencentcloud:Cfs/getFileSystemClients:getFileSystemClients", args ?? new GetFileSystemClientsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileSystemClientsArgs : Pulumi.InvokeArgs
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

        public GetFileSystemClientsArgs()
        {
        }
    }

    public sealed class GetFileSystemClientsInvokeArgs : Pulumi.InvokeArgs
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

        public GetFileSystemClientsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFileSystemClientsResult
    {
        /// <summary>
        /// Client list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFileSystemClientsClientListResult> ClientLists;
        public readonly string FileSystemId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetFileSystemClientsResult(
            ImmutableArray<Outputs.GetFileSystemClientsClientListResult> clientLists,

            string fileSystemId,

            string id,

            string? resultOutputFile)
        {
            ClientLists = clientLists;
            FileSystemId = fileSystemId;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
