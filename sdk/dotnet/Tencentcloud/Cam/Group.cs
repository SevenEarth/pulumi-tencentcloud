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
    /// <summary>
    /// Provides a resource to create a CAM group.
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
    ///         var foo = new Tencentcloud.Cam.Group("foo", new Tencentcloud.Cam.GroupArgs
    ///         {
    ///             Remark = "test",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CAM group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cam/group:Group foo 90496
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cam/group:Group")]
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// Create time of the CAM group.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Name of CAM group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Description of the CAM group.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of CAM group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Description of the CAM group.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create time of the CAM group.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Name of CAM group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Description of the CAM group.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        public GroupState()
        {
        }
    }
}
