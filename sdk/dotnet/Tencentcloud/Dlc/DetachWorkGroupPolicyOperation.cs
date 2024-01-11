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
    [TencentcloudResourceType("tencentcloud:Dlc/detachWorkGroupPolicyOperation:DetachWorkGroupPolicyOperation")]
    public partial class DetachWorkGroupPolicyOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// The set of policies to be bound.
        /// </summary>
        [Output("policySets")]
        public Output<ImmutableArray<Outputs.DetachWorkGroupPolicyOperationPolicySet>> PolicySets { get; private set; } = null!;

        /// <summary>
        /// Work group id.
        /// </summary>
        [Output("workGroupId")]
        public Output<int> WorkGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a DetachWorkGroupPolicyOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DetachWorkGroupPolicyOperation(string name, DetachWorkGroupPolicyOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/detachWorkGroupPolicyOperation:DetachWorkGroupPolicyOperation", name, args ?? new DetachWorkGroupPolicyOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DetachWorkGroupPolicyOperation(string name, Input<string> id, DetachWorkGroupPolicyOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/detachWorkGroupPolicyOperation:DetachWorkGroupPolicyOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DetachWorkGroupPolicyOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DetachWorkGroupPolicyOperation Get(string name, Input<string> id, DetachWorkGroupPolicyOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DetachWorkGroupPolicyOperation(name, id, state, options);
        }
    }

    public sealed class DetachWorkGroupPolicyOperationArgs : Pulumi.ResourceArgs
    {
        [Input("policySets")]
        private InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetArgs>? _policySets;

        /// <summary>
        /// The set of policies to be bound.
        /// </summary>
        public InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetArgs> PolicySets
        {
            get => _policySets ?? (_policySets = new InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetArgs>());
            set => _policySets = value;
        }

        /// <summary>
        /// Work group id.
        /// </summary>
        [Input("workGroupId", required: true)]
        public Input<int> WorkGroupId { get; set; } = null!;

        public DetachWorkGroupPolicyOperationArgs()
        {
        }
    }

    public sealed class DetachWorkGroupPolicyOperationState : Pulumi.ResourceArgs
    {
        [Input("policySets")]
        private InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetGetArgs>? _policySets;

        /// <summary>
        /// The set of policies to be bound.
        /// </summary>
        public InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetGetArgs> PolicySets
        {
            get => _policySets ?? (_policySets = new InputList<Inputs.DetachWorkGroupPolicyOperationPolicySetGetArgs>());
            set => _policySets = value;
        }

        /// <summary>
        /// Work group id.
        /// </summary>
        [Input("workGroupId")]
        public Input<int>? WorkGroupId { get; set; }

        public DetachWorkGroupPolicyOperationState()
        {
        }
    }
}