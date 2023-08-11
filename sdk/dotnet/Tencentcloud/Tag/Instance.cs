// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tag
{
    /// <summary>
    /// Provides a resource to create a tag
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
    ///         var tag = new Tencentcloud.Tag.Instance("tag", new Tencentcloud.Tag.InstanceArgs
    ///         {
    ///             TagKey = "test",
    ///             TagValue = "Terraform",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tag tag can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tag/instance:Instance tag tag_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tag/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// tag key.
        /// </summary>
        [Output("tagKey")]
        public Output<string> TagKey { get; private set; } = null!;

        /// <summary>
        /// tag value.
        /// </summary>
        [Output("tagValue")]
        public Output<string> TagValue { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tag/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tag/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// tag key.
        /// </summary>
        [Input("tagKey", required: true)]
        public Input<string> TagKey { get; set; } = null!;

        /// <summary>
        /// tag value.
        /// </summary>
        [Input("tagValue", required: true)]
        public Input<string> TagValue { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// tag key.
        /// </summary>
        [Input("tagKey")]
        public Input<string>? TagKey { get; set; }

        /// <summary>
        /// tag value.
        /// </summary>
        [Input("tagValue")]
        public Input<string>? TagValue { get; set; }

        public InstanceState()
        {
        }
    }
}
