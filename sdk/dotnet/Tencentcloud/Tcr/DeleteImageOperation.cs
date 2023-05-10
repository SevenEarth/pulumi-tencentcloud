// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    /// <summary>
    /// Provides a resource to create a tcr delete_image_operation
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
    ///         var deleteImageOperation = new Tencentcloud.Tcr.DeleteImageOperation("deleteImageOperation", new Tencentcloud.Tcr.DeleteImageOperationArgs
    ///         {
    ///             ImageVersion = "v1",
    ///             NamespaceName = "ns",
    ///             RegistryId = "tcr-xxx",
    ///             RepositoryName = "repo",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/deleteImageOperation:DeleteImageOperation")]
    public partial class DeleteImageOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// image version name.
        /// </summary>
        [Output("imageVersion")]
        public Output<string> ImageVersion { get; private set; } = null!;

        /// <summary>
        /// namespace name.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// repository name.
        /// </summary>
        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;


        /// <summary>
        /// Create a DeleteImageOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeleteImageOperation(string name, DeleteImageOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/deleteImageOperation:DeleteImageOperation", name, args ?? new DeleteImageOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeleteImageOperation(string name, Input<string> id, DeleteImageOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/deleteImageOperation:DeleteImageOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeleteImageOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeleteImageOperation Get(string name, Input<string> id, DeleteImageOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new DeleteImageOperation(name, id, state, options);
        }
    }

    public sealed class DeleteImageOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// image version name.
        /// </summary>
        [Input("imageVersion", required: true)]
        public Input<string> ImageVersion { get; set; } = null!;

        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// repository name.
        /// </summary>
        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        public DeleteImageOperationArgs()
        {
        }
    }

    public sealed class DeleteImageOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// image version name.
        /// </summary>
        [Input("imageVersion")]
        public Input<string>? ImageVersion { get; set; }

        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// repository name.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        public DeleteImageOperationState()
        {
        }
    }
}
