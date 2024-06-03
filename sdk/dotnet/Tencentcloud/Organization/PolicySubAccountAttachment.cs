// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Organization
{
    /// <summary>
    /// Provides a resource to create a organization policy_sub_account_attachment
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var policySubAccountAttachment = new Tencentcloud.Organization.PolicySubAccountAttachment("policySubAccountAttachment", new()
    ///     {
    ///         MemberUin = 100028582828,
    ///         OrgSubAccountUin = 100028223737,
    ///         PolicyId = 144256499,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// organization policy_sub_account_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment policy_sub_account_attachment policyId#memberUin#orgSubAccountUin
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment")]
    public partial class PolicySubAccountAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Manage Identity ID.
        /// </summary>
        [Output("identityId")]
        public Output<int> IdentityId { get; private set; } = null!;

        /// <summary>
        /// Identity role alias name.
        /// </summary>
        [Output("identityRoleAliasName")]
        public Output<string> IdentityRoleAliasName { get; private set; } = null!;

        /// <summary>
        /// Identity role name.
        /// </summary>
        [Output("identityRoleName")]
        public Output<string> IdentityRoleName { get; private set; } = null!;

        /// <summary>
        /// Organization member uin.
        /// </summary>
        [Output("memberUin")]
        public Output<int> MemberUin { get; private set; } = null!;

        /// <summary>
        /// Organization administrator sub account name.
        /// </summary>
        [Output("orgSubAccountName")]
        public Output<string> OrgSubAccountName { get; private set; } = null!;

        /// <summary>
        /// Organization administrator sub account uin list.
        /// </summary>
        [Output("orgSubAccountUin")]
        public Output<int> OrgSubAccountUin { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("policyId")]
        public Output<int> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Policy name.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// Update time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a PolicySubAccountAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicySubAccountAttachment(string name, PolicySubAccountAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment", name, args ?? new PolicySubAccountAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicySubAccountAttachment(string name, Input<string> id, PolicySubAccountAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Organization/policySubAccountAttachment:PolicySubAccountAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicySubAccountAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicySubAccountAttachment Get(string name, Input<string> id, PolicySubAccountAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicySubAccountAttachment(name, id, state, options);
        }
    }

    public sealed class PolicySubAccountAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Organization member uin.
        /// </summary>
        [Input("memberUin", required: true)]
        public Input<int> MemberUin { get; set; } = null!;

        /// <summary>
        /// Organization administrator sub account uin list.
        /// </summary>
        [Input("orgSubAccountUin", required: true)]
        public Input<int> OrgSubAccountUin { get; set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<int> PolicyId { get; set; } = null!;

        public PolicySubAccountAttachmentArgs()
        {
        }
        public static new PolicySubAccountAttachmentArgs Empty => new PolicySubAccountAttachmentArgs();
    }

    public sealed class PolicySubAccountAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Manage Identity ID.
        /// </summary>
        [Input("identityId")]
        public Input<int>? IdentityId { get; set; }

        /// <summary>
        /// Identity role alias name.
        /// </summary>
        [Input("identityRoleAliasName")]
        public Input<string>? IdentityRoleAliasName { get; set; }

        /// <summary>
        /// Identity role name.
        /// </summary>
        [Input("identityRoleName")]
        public Input<string>? IdentityRoleName { get; set; }

        /// <summary>
        /// Organization member uin.
        /// </summary>
        [Input("memberUin")]
        public Input<int>? MemberUin { get; set; }

        /// <summary>
        /// Organization administrator sub account name.
        /// </summary>
        [Input("orgSubAccountName")]
        public Input<string>? OrgSubAccountName { get; set; }

        /// <summary>
        /// Organization administrator sub account uin list.
        /// </summary>
        [Input("orgSubAccountUin")]
        public Input<int>? OrgSubAccountUin { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyId")]
        public Input<int>? PolicyId { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public PolicySubAccountAttachmentState()
        {
        }
        public static new PolicySubAccountAttachmentState Empty => new PolicySubAccountAttachmentState();
    }
}
