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
    /// Provides a resource to create a dlc bind_work_groups_to_user
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
    ///         var bindWorkGroupsToUser = new Tencentcloud.Dlc.BindWorkGroupsToUserAttachment("bindWorkGroupsToUser", new Tencentcloud.Dlc.BindWorkGroupsToUserAttachmentArgs
    ///         {
    ///             AddInfo = new Tencentcloud.Dlc.Inputs.BindWorkGroupsToUserAttachmentAddInfoArgs
    ///             {
    ///                 UserId = "100032772113",
    ///                 WorkGroupIds = 
    ///                 {
    ///                     23184,
    ///                     23181,
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
    /// dlc bind_work_groups_to_user can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment bind_work_groups_to_user bind_work_groups_to_user_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment")]
    public partial class BindWorkGroupsToUserAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Bind user and workgroup information.
        /// </summary>
        [Output("addInfo")]
        public Output<Outputs.BindWorkGroupsToUserAttachmentAddInfo> AddInfo { get; private set; } = null!;


        /// <summary>
        /// Create a BindWorkGroupsToUserAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BindWorkGroupsToUserAttachment(string name, BindWorkGroupsToUserAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment", name, args ?? new BindWorkGroupsToUserAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BindWorkGroupsToUserAttachment(string name, Input<string> id, BindWorkGroupsToUserAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BindWorkGroupsToUserAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BindWorkGroupsToUserAttachment Get(string name, Input<string> id, BindWorkGroupsToUserAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new BindWorkGroupsToUserAttachment(name, id, state, options);
        }
    }

    public sealed class BindWorkGroupsToUserAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bind user and workgroup information.
        /// </summary>
        [Input("addInfo", required: true)]
        public Input<Inputs.BindWorkGroupsToUserAttachmentAddInfoArgs> AddInfo { get; set; } = null!;

        public BindWorkGroupsToUserAttachmentArgs()
        {
        }
    }

    public sealed class BindWorkGroupsToUserAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bind user and workgroup information.
        /// </summary>
        [Input("addInfo")]
        public Input<Inputs.BindWorkGroupsToUserAttachmentAddInfoGetArgs>? AddInfo { get; set; }

        public BindWorkGroupsToUserAttachmentState()
        {
        }
    }
}
