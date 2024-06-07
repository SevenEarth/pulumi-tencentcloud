// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    /// <summary>
    /// Provides a resource to create a redis security_group_attachment
    /// 
    /// ## Import
    /// 
    /// redis security_group_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Redis/securityGroupAttachment:SecurityGroupAttachment security_group_attachment instance_id#security_group_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Redis/securityGroupAttachment:SecurityGroupAttachment")]
    public partial class SecurityGroupAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Security group ID.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupAttachment(string name, SecurityGroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/securityGroupAttachment:SecurityGroupAttachment", name, args ?? new SecurityGroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupAttachment(string name, Input<string> id, SecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/securityGroupAttachment:SecurityGroupAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupAttachment Get(string name, Input<string> id, SecurityGroupAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupAttachment(name, id, state, options);
        }
    }

    public sealed class SecurityGroupAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Security group ID.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public SecurityGroupAttachmentArgs()
        {
        }
        public static new SecurityGroupAttachmentArgs Empty => new SecurityGroupAttachmentArgs();
    }

    public sealed class SecurityGroupAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Security group ID.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public SecurityGroupAttachmentState()
        {
        }
        public static new SecurityGroupAttachmentState Empty => new SecurityGroupAttachmentState();
    }
}
