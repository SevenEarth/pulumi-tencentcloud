// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cam
{
    public static class GetGroupMemberships
    {
        /// <summary>
        /// Use this data source to query detailed information of CAM group memberships
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
        ///         var foo = Output.Create(Tencentcloud.Cam.GetGroupMemberships.InvokeAsync(new Tencentcloud.Cam.GetGroupMembershipsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupMembershipsResult> InvokeAsync(GetGroupMembershipsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupMembershipsResult>("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of CAM group memberships
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
        ///         var foo = Output.Create(Tencentcloud.Cam.GetGroupMemberships.InvokeAsync(new Tencentcloud.Cam.GetGroupMembershipsArgs
        ///         {
        ///             GroupId = tencentcloud_cam_group.Foo.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupMembershipsResult> Invoke(GetGroupMembershipsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupMembershipsResult>("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMembershipsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of CAM group to be queried.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetGroupMembershipsArgs()
        {
        }
    }

    public sealed class GetGroupMembershipsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of CAM group to be queried.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetGroupMembershipsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupMembershipsResult
    {
        /// <summary>
        /// ID of CAM group.
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of CAM group membership. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupMembershipsMembershipListResult> MembershipLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetGroupMembershipsResult(
            string? groupId,

            string id,

            ImmutableArray<Outputs.GetGroupMembershipsMembershipListResult> membershipLists,

            string? resultOutputFile)
        {
            GroupId = groupId;
            Id = id;
            MembershipLists = membershipLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
