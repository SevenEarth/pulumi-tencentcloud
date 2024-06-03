// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eip
{
    /// <summary>
    /// Provides a resource to create a eip address_transform
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
    ///     var addressTransform = new Tencentcloud.Eip.AddressTransform("addressTransform", new()
    ///     {
    ///         InstanceId = "",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// eip address_transform can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Eip/addressTransform:AddressTransform address_transform address_transform_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Eip/addressTransform:AddressTransform")]
    public partial class AddressTransform : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the instance ID of a normal public network IP to be operated. eg:ins-23mk45jn.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a AddressTransform resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddressTransform(string name, AddressTransformArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Eip/addressTransform:AddressTransform", name, args ?? new AddressTransformArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AddressTransform(string name, Input<string> id, AddressTransformState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Eip/addressTransform:AddressTransform", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AddressTransform resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddressTransform Get(string name, Input<string> id, AddressTransformState? state = null, CustomResourceOptions? options = null)
        {
            return new AddressTransform(name, id, state, options);
        }
    }

    public sealed class AddressTransformArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the instance ID of a normal public network IP to be operated. eg:ins-23mk45jn.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AddressTransformArgs()
        {
        }
        public static new AddressTransformArgs Empty => new AddressTransformArgs();
    }

    public sealed class AddressTransformState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the instance ID of a normal public network IP to be operated. eg:ins-23mk45jn.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AddressTransformState()
        {
        }
        public static new AddressTransformState Empty => new AddressTransformState();
    }
}
