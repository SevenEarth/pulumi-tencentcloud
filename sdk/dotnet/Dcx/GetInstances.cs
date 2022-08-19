// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dcx
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of dedicated tunnels instances.
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
        ///         var nameSelect = Output.Create(Tencentcloud.Dcx.GetInstances.InvokeAsync(new Tencentcloud.Dcx.GetInstancesArgs
        ///         {
        ///             Name = "main",
        ///         }));
        ///         var id = Output.Create(Tencentcloud.Dcx.GetInstances.InvokeAsync(new Tencentcloud.Dcx.GetInstancesArgs
        ///         {
        ///             DcxId = "dcx-3ikuw30k",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:Dcx/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dedicated tunnels instances.
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
        ///         var nameSelect = Output.Create(Tencentcloud.Dcx.GetInstances.InvokeAsync(new Tencentcloud.Dcx.GetInstancesArgs
        ///         {
        ///             Name = "main",
        ///         }));
        ///         var id = Output.Create(Tencentcloud.Dcx.GetInstances.InvokeAsync(new Tencentcloud.Dcx.GetInstancesArgs
        ///         {
        ///             DcxId = "dcx-3ikuw30k",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:Dcx/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the dedicated tunnels to be queried.
        /// </summary>
        [Input("dcxId")]
        public string? DcxId { get; set; }

        /// <summary>
        /// Name of the dedicated tunnels to be queried.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstancesArgs()
        {
        }
    }

    public sealed class GetInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the dedicated tunnels to be queried.
        /// </summary>
        [Input("dcxId")]
        public Input<string>? DcxId { get; set; }

        /// <summary>
        /// Name of the dedicated tunnels to be queried.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// ID of the dedicated tunnel.
        /// </summary>
        public readonly string? DcxId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Information list of the dedicated tunnels.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesInstanceListResult> InstanceLists;
        /// <summary>
        /// Name of the dedicated tunnel.
        /// </summary>
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstancesResult(
            string? dcxId,

            string id,

            ImmutableArray<Outputs.GetInstancesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile)
        {
            DcxId = dcxId;
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
