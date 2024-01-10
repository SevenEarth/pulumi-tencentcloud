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
    /// <summary>
    /// Provides a resource to create a cam tag_role
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
    ///         var tagRole = new Tencentcloud.Cam.TagRoleAttachment("tagRole", new Tencentcloud.Cam.TagRoleAttachmentArgs
    ///         {
    ///             RoleId = "test-cam-tag",
    ///             Tags = 
    ///             {
    ///                 new Tencentcloud.Cam.Inputs.TagRoleAttachmentTagArgs
    ///                 {
    ///                     Key = "test1",
    ///                     Value = "test1",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cam tag_role can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cam/tagRoleAttachment:TagRoleAttachment tag_role tag_role_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cam/tagRoleAttachment:TagRoleAttachment")]
    public partial class TagRoleAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Character ID, at least one input with the character name.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// Character name, at least one input with the character ID.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        /// <summary>
        /// Label.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TagRoleAttachmentTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a TagRoleAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagRoleAttachment(string name, TagRoleAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/tagRoleAttachment:TagRoleAttachment", name, args ?? new TagRoleAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagRoleAttachment(string name, Input<string> id, TagRoleAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/tagRoleAttachment:TagRoleAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TagRoleAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagRoleAttachment Get(string name, Input<string> id, TagRoleAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new TagRoleAttachment(name, id, state, options);
        }
    }

    public sealed class TagRoleAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Character ID, at least one input with the character name.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// Character name, at least one input with the character ID.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("tags", required: true)]
        private InputList<Inputs.TagRoleAttachmentTagArgs>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputList<Inputs.TagRoleAttachmentTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagRoleAttachmentTagArgs>());
            set => _tags = value;
        }

        public TagRoleAttachmentArgs()
        {
        }
    }

    public sealed class TagRoleAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Character ID, at least one input with the character name.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// Character name, at least one input with the character ID.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("tags")]
        private InputList<Inputs.TagRoleAttachmentTagGetArgs>? _tags;

        /// <summary>
        /// Label.
        /// </summary>
        public InputList<Inputs.TagRoleAttachmentTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TagRoleAttachmentTagGetArgs>());
            set => _tags = value;
        }

        public TagRoleAttachmentState()
        {
        }
    }
}
