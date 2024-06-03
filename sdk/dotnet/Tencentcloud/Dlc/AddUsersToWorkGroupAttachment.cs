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
    /// Provides a resource to create a dlc add_users_to_work_group_attachment
    /// 
    /// ## Import
    /// 
    /// dlc add_users_to_work_group_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment add_users_to_work_group_attachment add_users_to_work_group_attachment_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment")]
    public partial class AddUsersToWorkGroupAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Work group and user information to operate on.
        /// </summary>
        [Output("addInfo")]
        public Output<Outputs.AddUsersToWorkGroupAttachmentAddInfo> AddInfo { get; private set; } = null!;


        /// <summary>
        /// Create a AddUsersToWorkGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddUsersToWorkGroupAttachment(string name, AddUsersToWorkGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment", name, args ?? new AddUsersToWorkGroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AddUsersToWorkGroupAttachment(string name, Input<string> id, AddUsersToWorkGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/addUsersToWorkGroupAttachment:AddUsersToWorkGroupAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AddUsersToWorkGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddUsersToWorkGroupAttachment Get(string name, Input<string> id, AddUsersToWorkGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AddUsersToWorkGroupAttachment(name, id, state, options);
        }
    }

    public sealed class AddUsersToWorkGroupAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Work group and user information to operate on.
        /// </summary>
        [Input("addInfo", required: true)]
        public Input<Inputs.AddUsersToWorkGroupAttachmentAddInfoArgs> AddInfo { get; set; } = null!;

        public AddUsersToWorkGroupAttachmentArgs()
        {
        }
        public static new AddUsersToWorkGroupAttachmentArgs Empty => new AddUsersToWorkGroupAttachmentArgs();
    }

    public sealed class AddUsersToWorkGroupAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Work group and user information to operate on.
        /// </summary>
        [Input("addInfo")]
        public Input<Inputs.AddUsersToWorkGroupAttachmentAddInfoGetArgs>? AddInfo { get; set; }

        public AddUsersToWorkGroupAttachmentState()
        {
        }
        public static new AddUsersToWorkGroupAttachmentState Empty => new AddUsersToWorkGroupAttachmentState();
    }
}
