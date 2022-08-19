// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod
{
    /// <summary>
    /// Provide a resource to create a VOD sub application.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Vod.SubApplication("foo", new Tencentcloud.Vod.SubApplicationArgs
    ///         {
    ///             Description = "this is sub application",
    ///             Status = "On",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VOD super player config can be imported using the name+, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vod/subApplication:SubApplication foo name+"#"+id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vod/subApplication:SubApplication")]
    public partial class SubApplication : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the sub application was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Sub application description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sub appliaction status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SubApplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubApplication(string name, SubApplicationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/subApplication:SubApplication", name, args ?? new SubApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubApplication(string name, Input<string> id, SubApplicationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vod/subApplication:SubApplication", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubApplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubApplication Get(string name, Input<string> id, SubApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new SubApplication(name, id, state, options);
        }
    }

    public sealed class SubApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sub application description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sub appliaction status.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public SubApplicationArgs()
        {
        }
    }

    public sealed class SubApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the sub application was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Sub application description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Sub application name, which can contain up to 64 letters, digits, underscores, and hyphens (such as test_ABC-123) and must be unique under a user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sub appliaction status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public SubApplicationState()
        {
        }
    }
}
