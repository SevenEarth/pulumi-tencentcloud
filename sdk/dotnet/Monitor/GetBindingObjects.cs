// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    public static class GetBindingObjects
    {
        /// <summary>
        /// Use this data source to query policy group binding objects.
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
        ///         var name = Output.Create(Tencentcloud.Monitor.GetPolicyGroups.InvokeAsync(new Tencentcloud.Monitor.GetPolicyGroupsArgs
        ///         {
        ///             Name = "test",
        ///         }));
        ///         var objects = name.Apply(name =&gt; Output.Create(Tencentcloud.Monitor.GetBindingObjects.InvokeAsync(new Tencentcloud.Monitor.GetBindingObjectsArgs
        ///         {
        ///             GroupId = name.Lists?[0]?.GroupId,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBindingObjectsResult> InvokeAsync(GetBindingObjectsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBindingObjectsResult>("tencentcloud:Monitor/getBindingObjects:getBindingObjects", args ?? new GetBindingObjectsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query policy group binding objects.
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
        ///         var name = Output.Create(Tencentcloud.Monitor.GetPolicyGroups.InvokeAsync(new Tencentcloud.Monitor.GetPolicyGroupsArgs
        ///         {
        ///             Name = "test",
        ///         }));
        ///         var objects = name.Apply(name =&gt; Output.Create(Tencentcloud.Monitor.GetBindingObjects.InvokeAsync(new Tencentcloud.Monitor.GetBindingObjectsArgs
        ///         {
        ///             GroupId = name.Lists?[0]?.GroupId,
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBindingObjectsResult> Invoke(GetBindingObjectsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBindingObjectsResult>("tencentcloud:Monitor/getBindingObjects:getBindingObjects", args ?? new GetBindingObjectsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBindingObjectsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Policy group ID for query.
        /// </summary>
        [Input("groupId", required: true)]
        public int GroupId { get; set; }

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBindingObjectsArgs()
        {
        }
    }

    public sealed class GetBindingObjectsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Policy group ID for query.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// Used to store results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBindingObjectsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBindingObjectsResult
    {
        public readonly int GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBindingObjectsListResult> Lists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBindingObjectsResult(
            int groupId,

            string id,

            ImmutableArray<Outputs.GetBindingObjectsListResult> lists,

            string? resultOutputFile)
        {
            GroupId = groupId;
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
