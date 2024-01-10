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
    /// Provides a resource to create a dlc work_group
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
    ///         var workGroup = new Tencentcloud.Dlc.WorkGroup("workGroup", new Tencentcloud.Dlc.WorkGroupArgs
    ///         {
    ///             WorkGroupDescription = "dlc workgroup test",
    ///             WorkGroupName = "tf-demo",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dlc work_group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dlc/workGroup:WorkGroup work_group work_group_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/workGroup:WorkGroup")]
    public partial class WorkGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// A collection of user IDs that has been bound to the workgroup.
        /// </summary>
        [Output("userIds")]
        public Output<ImmutableArray<string>> UserIds { get; private set; } = null!;

        /// <summary>
        /// Description of Work Group.
        /// </summary>
        [Output("workGroupDescription")]
        public Output<string?> WorkGroupDescription { get; private set; } = null!;

        /// <summary>
        /// Name of Work Group.
        /// </summary>
        [Output("workGroupName")]
        public Output<string> WorkGroupName { get; private set; } = null!;


        /// <summary>
        /// Create a WorkGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkGroup(string name, WorkGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/workGroup:WorkGroup", name, args ?? new WorkGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkGroup(string name, Input<string> id, WorkGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/workGroup:WorkGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkGroup Get(string name, Input<string> id, WorkGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkGroup(name, id, state, options);
        }
    }

    public sealed class WorkGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of Work Group.
        /// </summary>
        [Input("workGroupDescription")]
        public Input<string>? WorkGroupDescription { get; set; }

        /// <summary>
        /// Name of Work Group.
        /// </summary>
        [Input("workGroupName", required: true)]
        public Input<string> WorkGroupName { get; set; } = null!;

        public WorkGroupArgs()
        {
        }
    }

    public sealed class WorkGroupState : Pulumi.ResourceArgs
    {
        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// A collection of user IDs that has been bound to the workgroup.
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        /// <summary>
        /// Description of Work Group.
        /// </summary>
        [Input("workGroupDescription")]
        public Input<string>? WorkGroupDescription { get; set; }

        /// <summary>
        /// Name of Work Group.
        /// </summary>
        [Input("workGroupName")]
        public Input<string>? WorkGroupName { get; set; }

        public WorkGroupState()
        {
        }
    }
}
