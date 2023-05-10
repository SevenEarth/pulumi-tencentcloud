// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a resource to create a mysql security_groups_attachment
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
    ///         var securityGroupsAttachment = new Tencentcloud.Mysql.SecurityGroupsAttachment("securityGroupsAttachment", new Tencentcloud.Mysql.SecurityGroupsAttachmentArgs
    ///         {
    ///             InstanceId = "cdb-fitq5t9h",
    ///             SecurityGroupId = "sg-baxfiao5",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mysql security_groups_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mysql/securityGroupsAttachment:SecurityGroupsAttachment security_groups_attachment securityGroupId#instanceId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/securityGroupsAttachment:SecurityGroupsAttachment")]
    public partial class SecurityGroupsAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupsAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupsAttachment(string name, SecurityGroupsAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/securityGroupsAttachment:SecurityGroupsAttachment", name, args ?? new SecurityGroupsAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupsAttachment(string name, Input<string> id, SecurityGroupsAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/securityGroupsAttachment:SecurityGroupsAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupsAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupsAttachment Get(string name, Input<string> id, SecurityGroupsAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupsAttachment(name, id, state, options);
        }
    }

    public sealed class SecurityGroupsAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public SecurityGroupsAttachmentArgs()
        {
        }
    }

    public sealed class SecurityGroupsAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of security group.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public SecurityGroupsAttachmentState()
        {
        }
    }
}
