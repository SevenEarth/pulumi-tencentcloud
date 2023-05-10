// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf microservice
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
    ///         var microservice = new Tencentcloud.Tsf.Microservice("microservice", new Tencentcloud.Tsf.MicroserviceArgs
    ///         {
    ///             MicroserviceDesc = "desc-microservice",
    ///             MicroserviceName = "test-microservice",
    ///             NamespaceId = "namespace-vjlkzkgy",
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tsf microservice can be imported using the namespaceId#microserviceId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tsf/microservice:Microservice microservice namespace-vjlkzkgy#ms-vjeb43lw
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/microservice:Microservice")]
    public partial class Microservice : Pulumi.CustomResource
    {
        /// <summary>
        /// Microservice description information.
        /// </summary>
        [Output("microserviceDesc")]
        public Output<string?> MicroserviceDesc { get; private set; } = null!;

        /// <summary>
        /// Microservice name.
        /// </summary>
        [Output("microserviceName")]
        public Output<string> MicroserviceName { get; private set; } = null!;

        /// <summary>
        /// Namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Microservice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Microservice(string name, MicroserviceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/microservice:Microservice", name, args ?? new MicroserviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Microservice(string name, Input<string> id, MicroserviceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/microservice:Microservice", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Microservice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Microservice Get(string name, Input<string> id, MicroserviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Microservice(name, id, state, options);
        }
    }

    public sealed class MicroserviceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Microservice description information.
        /// </summary>
        [Input("microserviceDesc")]
        public Input<string>? MicroserviceDesc { get; set; }

        /// <summary>
        /// Microservice name.
        /// </summary>
        [Input("microserviceName", required: true)]
        public Input<string> MicroserviceName { get; set; } = null!;

        /// <summary>
        /// Namespace ID.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public MicroserviceArgs()
        {
        }
    }

    public sealed class MicroserviceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Microservice description information.
        /// </summary>
        [Input("microserviceDesc")]
        public Input<string>? MicroserviceDesc { get; set; }

        /// <summary>
        /// Microservice name.
        /// </summary>
        [Input("microserviceName")]
        public Input<string>? MicroserviceName { get; set; }

        /// <summary>
        /// Namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public MicroserviceState()
        {
        }
    }
}