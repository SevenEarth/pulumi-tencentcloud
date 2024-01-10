// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc
{
    /// <summary>
    /// Provides a resource to create a dlc user
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
    ///         var user = new Tencentcloud.Dlc.User("user", new Tencentcloud.Dlc.UserArgs
    ///         {
    ///             UserAlias = "terraform-test",
    ///             UserDescription = "for terraform test",
    ///             UserId = "100027012454",
    ///             UserType = "COMMON",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dlc user can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dlc/user:User user user_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// User alias, the character length is less than 50.
        /// </summary>
        [Output("userAlias")]
        public Output<string?> UserAlias { get; private set; } = null!;

        /// <summary>
        /// User description information, easy to distinguish between different users.
        /// </summary>
        [Output("userDescription")]
        public Output<string?> UserDescription { get; private set; } = null!;

        /// <summary>
        /// The sub-user uin that needs to be authorized.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// User Type. `ADMIN` or `COMMONN`.
        /// </summary>
        [Output("userType")]
        public Output<string?> UserType { get; private set; } = null!;

        /// <summary>
        /// A collection of workgroup IDs bound to the user.
        /// </summary>
        [Output("workGroupIds")]
        public Output<ImmutableArray<int>> WorkGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User alias, the character length is less than 50.
        /// </summary>
        [Input("userAlias")]
        public Input<string>? UserAlias { get; set; }

        /// <summary>
        /// User description information, easy to distinguish between different users.
        /// </summary>
        [Input("userDescription")]
        public Input<string>? UserDescription { get; set; }

        /// <summary>
        /// The sub-user uin that needs to be authorized.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        /// <summary>
        /// User Type. `ADMIN` or `COMMONN`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User alias, the character length is less than 50.
        /// </summary>
        [Input("userAlias")]
        public Input<string>? UserAlias { get; set; }

        /// <summary>
        /// User description information, easy to distinguish between different users.
        /// </summary>
        [Input("userDescription")]
        public Input<string>? UserDescription { get; set; }

        /// <summary>
        /// The sub-user uin that needs to be authorized.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// User Type. `ADMIN` or `COMMONN`.
        /// </summary>
        [Input("userType")]
        public Input<string>? UserType { get; set; }

        [Input("workGroupIds")]
        private InputList<int>? _workGroupIds;

        /// <summary>
        /// A collection of workgroup IDs bound to the user.
        /// </summary>
        public InputList<int> WorkGroupIds
        {
            get => _workGroupIds ?? (_workGroupIds = new InputList<int>());
            set => _workGroupIds = value;
        }

        public UserState()
        {
        }
    }
}
