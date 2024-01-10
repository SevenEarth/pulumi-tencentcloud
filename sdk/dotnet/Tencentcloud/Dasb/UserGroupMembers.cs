// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dasb
{
    /// <summary>
    /// Provides a resource to create a dasb user_group_members
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Dasb.UserGroupMembers("example", new Tencentcloud.Dasb.UserGroupMembersArgs
    ///         {
    ///             MemberIdSets = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///             },
    ///             UserGroupId = 3,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dasb user_group_members can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dasb/userGroupMembers:UserGroupMembers example 3#1,2,3
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dasb/userGroupMembers:UserGroupMembers")]
    public partial class UserGroupMembers : Pulumi.CustomResource
    {
        /// <summary>
        /// Collection of member user IDs.
        /// </summary>
        [Output("memberIdSets")]
        public Output<ImmutableArray<int>> MemberIdSets { get; private set; } = null!;

        /// <summary>
        /// User Group ID.
        /// </summary>
        [Output("userGroupId")]
        public Output<int> UserGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroupMembers resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroupMembers(string name, UserGroupMembersArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/userGroupMembers:UserGroupMembers", name, args ?? new UserGroupMembersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroupMembers(string name, Input<string> id, UserGroupMembersState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/userGroupMembers:UserGroupMembers", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserGroupMembers resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroupMembers Get(string name, Input<string> id, UserGroupMembersState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroupMembers(name, id, state, options);
        }
    }

    public sealed class UserGroupMembersArgs : Pulumi.ResourceArgs
    {
        [Input("memberIdSets", required: true)]
        private InputList<int>? _memberIdSets;

        /// <summary>
        /// Collection of member user IDs.
        /// </summary>
        public InputList<int> MemberIdSets
        {
            get => _memberIdSets ?? (_memberIdSets = new InputList<int>());
            set => _memberIdSets = value;
        }

        /// <summary>
        /// User Group ID.
        /// </summary>
        [Input("userGroupId", required: true)]
        public Input<int> UserGroupId { get; set; } = null!;

        public UserGroupMembersArgs()
        {
        }
    }

    public sealed class UserGroupMembersState : Pulumi.ResourceArgs
    {
        [Input("memberIdSets")]
        private InputList<int>? _memberIdSets;

        /// <summary>
        /// Collection of member user IDs.
        /// </summary>
        public InputList<int> MemberIdSets
        {
            get => _memberIdSets ?? (_memberIdSets = new InputList<int>());
            set => _memberIdSets = value;
        }

        /// <summary>
        /// User Group ID.
        /// </summary>
        [Input("userGroupId")]
        public Input<int>? UserGroupId { get; set; }

        public UserGroupMembersState()
        {
        }
    }
}
