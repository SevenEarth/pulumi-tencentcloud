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
    [TencentcloudResourceType("tencentcloud:Cam/rolePolicyAttachmentByName:RolePolicyAttachmentByName")]
    public partial class RolePolicyAttachmentByName : Pulumi.CustomResource
    {
        /// <summary>
        /// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and
        /// the others indicate syntax strategy ways.
        /// </summary>
        [Output("createMode")]
        public Output<int> CreateMode { get; private set; } = null!;

        /// <summary>
        /// The create time of the CAM role policy attachment.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        /// <summary>
        /// Name of the attached CAM role.
        /// </summary>
        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;


        /// <summary>
        /// Create a RolePolicyAttachmentByName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RolePolicyAttachmentByName(string name, RolePolicyAttachmentByNameArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/rolePolicyAttachmentByName:RolePolicyAttachmentByName", name, args ?? new RolePolicyAttachmentByNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RolePolicyAttachmentByName(string name, Input<string> id, RolePolicyAttachmentByNameState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/rolePolicyAttachmentByName:RolePolicyAttachmentByName", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RolePolicyAttachmentByName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RolePolicyAttachmentByName Get(string name, Input<string> id, RolePolicyAttachmentByNameState? state = null, CustomResourceOptions? options = null)
        {
            return new RolePolicyAttachmentByName(name, id, state, options);
        }
    }

    public sealed class RolePolicyAttachmentByNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// Name of the attached CAM role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        public RolePolicyAttachmentByNameArgs()
        {
        }
    }

    public sealed class RolePolicyAttachmentByNameState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode of Creation of the CAM role policy attachment. `1` means the CAM policy attachment is created by production, and
        /// the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public Input<int>? CreateMode { get; set; }

        /// <summary>
        /// The create time of the CAM role policy attachment.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Name of the policy.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// Type of the policy strategy. `User` means customer strategy and `QCS` means preset strategy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Name of the attached CAM role.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        public RolePolicyAttachmentByNameState()
        {
        }
    }
}
