// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provide a resource to attach an existing subnet to Network ACL.
    /// 
    /// ## Import
    /// 
    /// Acl attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/aclAttachment:AclAttachment attachment acl-eotx5qsg#subnet-91x0geu6
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/aclAttachment:AclAttachment")]
    public partial class AclAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the attached ACL.
        /// </summary>
        [Output("aclId")]
        public Output<string> AclId { get; private set; } = null!;

        /// <summary>
        /// The Subnet instance ID.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a AclAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AclAttachment(string name, AclAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/aclAttachment:AclAttachment", name, args ?? new AclAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AclAttachment(string name, Input<string> id, AclAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/aclAttachment:AclAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AclAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AclAttachment Get(string name, Input<string> id, AclAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AclAttachment(name, id, state, options);
        }
    }

    public sealed class AclAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the attached ACL.
        /// </summary>
        [Input("aclId", required: true)]
        public Input<string> AclId { get; set; } = null!;

        /// <summary>
        /// The Subnet instance ID.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public AclAttachmentArgs()
        {
        }
    }

    public sealed class AclAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the attached ACL.
        /// </summary>
        [Input("aclId")]
        public Input<string>? AclId { get; set; }

        /// <summary>
        /// The Subnet instance ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public AclAttachmentState()
        {
        }
    }
}
