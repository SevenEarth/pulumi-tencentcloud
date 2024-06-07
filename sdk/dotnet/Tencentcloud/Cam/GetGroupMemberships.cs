// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    public static class GetGroupMemberships
    {
        /// <summary>
        /// Use this data source to query detailed information of CAM group memberships
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
        ///     var foo = Tencentcloud.Cam.GetGroupMemberships.Invoke(new()
        ///     {
        ///         GroupId = tencentcloud_cam_group.Foo.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetGroupMembershipsResult> InvokeAsync(GetGroupMembershipsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupMembershipsResult>("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of CAM group memberships
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
        ///     var foo = Tencentcloud.Cam.GetGroupMemberships.Invoke(new()
        ///     {
        ///         GroupId = tencentcloud_cam_group.Foo.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetGroupMembershipsResult> Invoke(GetGroupMembershipsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupMembershipsResult>("tencentcloud:Cam/getGroupMemberships:getGroupMemberships", args ?? new GetGroupMembershipsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupMembershipsArgs : global::Pulumi.InvokeArgs
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
        public static new GetGroupMembershipsArgs Empty => new GetGroupMembershipsArgs();
    }

    public sealed class GetGroupMembershipsInvokeArgs : global::Pulumi.InvokeArgs
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
        public static new GetGroupMembershipsInvokeArgs Empty => new GetGroupMembershipsInvokeArgs();
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
