// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As
{
    /// <summary>
    /// Provides a resource to create a as remove_instances
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
    ///         var removeInstances = new Tencentcloud.As.RemoveInstances("removeInstances", new Tencentcloud.As.RemoveInstancesArgs
    ///         {
    ///             AutoScalingGroupId = tencentcloud_as_scaling_group.Scaling_group.Id,
    ///             InstanceIds = 
    ///             {
    ///                 "ins-xxxxxx",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:As/removeInstances:RemoveInstances")]
    public partial class RemoveInstances : Pulumi.CustomResource
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        [Output("autoScalingGroupId")]
        public Output<string> AutoScalingGroupId { get; private set; } = null!;

        /// <summary>
        /// List of cvm instances to remove.
        /// </summary>
        [Output("instanceIds")]
        public Output<ImmutableArray<string>> InstanceIds { get; private set; } = null!;


        /// <summary>
        /// Create a RemoveInstances resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RemoveInstances(string name, RemoveInstancesArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:As/removeInstances:RemoveInstances", name, args ?? new RemoveInstancesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RemoveInstances(string name, Input<string> id, RemoveInstancesState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:As/removeInstances:RemoveInstances", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RemoveInstances resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RemoveInstances Get(string name, Input<string> id, RemoveInstancesState? state = null, CustomResourceOptions? options = null)
        {
            return new RemoveInstances(name, id, state, options);
        }
    }

    public sealed class RemoveInstancesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        [Input("autoScalingGroupId", required: true)]
        public Input<string> AutoScalingGroupId { get; set; } = null!;

        [Input("instanceIds", required: true)]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// List of cvm instances to remove.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        public RemoveInstancesArgs()
        {
        }
    }

    public sealed class RemoveInstancesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Launch configuration ID.
        /// </summary>
        [Input("autoScalingGroupId")]
        public Input<string>? AutoScalingGroupId { get; set; }

        [Input("instanceIds")]
        private InputList<string>? _instanceIds;

        /// <summary>
        /// List of cvm instances to remove.
        /// </summary>
        public InputList<string> InstanceIds
        {
            get => _instanceIds ?? (_instanceIds = new InputList<string>());
            set => _instanceIds = value;
        }

        public RemoveInstancesState()
        {
        }
    }
}