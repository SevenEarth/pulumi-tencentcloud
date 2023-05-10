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
    [TencentcloudResourceType("tencentcloud:Cam/policyByName:PolicyByName")]
    public partial class PolicyByName : Pulumi.CustomResource
    {
        /// <summary>
        /// Create time of the CAM policy.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the CAM policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Document of the CAM policy. The syntax refers to [CAM
        /// POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        /// terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        /// Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        /// </summary>
        [Output("document")]
        public Output<string> Document { get; private set; } = null!;

        /// <summary>
        /// Name of CAM policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        /// </summary>
        [Output("type")]
        public Output<int> Type { get; private set; } = null!;

        /// <summary>
        /// The last update time of the CAM policy.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyByName resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyByName(string name, PolicyByNameArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/policyByName:PolicyByName", name, args ?? new PolicyByNameArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyByName(string name, Input<string> id, PolicyByNameState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/policyByName:PolicyByName", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyByName resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyByName Get(string name, Input<string> id, PolicyByNameState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyByName(name, id, state, options);
        }
    }

    public sealed class PolicyByNameArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the CAM policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Document of the CAM policy. The syntax refers to [CAM
        /// POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        /// terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        /// Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        /// </summary>
        [Input("document", required: true)]
        public Input<string> Document { get; set; } = null!;

        /// <summary>
        /// Name of CAM policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PolicyByNameArgs()
        {
        }
    }

    public sealed class PolicyByNameState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time of the CAM policy.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the CAM policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Document of the CAM policy. The syntax refers to [CAM
        /// POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in
        /// terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2.
        /// Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.
        /// </summary>
        [Input("document")]
        public Input<string>? Document { get; set; }

        /// <summary>
        /// Name of CAM policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of the policy strategy. Valid values: `1`, `2`. `1` means customer strategy and `2` means preset strategy.
        /// </summary>
        [Input("type")]
        public Input<int>? Type { get; set; }

        /// <summary>
        /// The last update time of the CAM policy.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public PolicyByNameState()
        {
        }
    }
}
