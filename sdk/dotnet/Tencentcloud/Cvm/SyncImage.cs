// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm
{
    /// <summary>
    /// Provides a resource to create a cvm sync_image
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleInstance = Tencentcloud.Images.GetInstance.Invoke(new()
    ///     {
    ///         ImageTypes = new[]
    ///         {
    ///             "PRIVATE_IMAGE",
    ///         },
    ///         ImageNameRegex = "MyImage",
    ///     });
    /// 
    ///     var exampleSyncImage = new Tencentcloud.Cvm.SyncImage("exampleSyncImage", new()
    ///     {
    ///         ImageId = exampleInstance.Apply(getInstanceResult =&gt; getInstanceResult.Images[0]?.ImageId),
    ///         DestinationRegions = new[]
    ///         {
    ///             "ap-guangzhou",
    ///             "ap-shanghai",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cvm/syncImage:SyncImage")]
    public partial class SyncImage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        /// </summary>
        [Output("destinationRegions")]
        public Output<ImmutableArray<string>> DestinationRegions { get; private set; } = null!;

        /// <summary>
        /// Checks whether image synchronization can be initiated.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// Destination image name.
        /// </summary>
        [Output("imageName")]
        public Output<string?> ImageName { get; private set; } = null!;

        /// <summary>
        /// Whether to return the ID of image created in the destination region.
        /// </summary>
        [Output("imageSetRequired")]
        public Output<bool?> ImageSetRequired { get; private set; } = null!;


        /// <summary>
        /// Create a SyncImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SyncImage(string name, SyncImageArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/syncImage:SyncImage", name, args ?? new SyncImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SyncImage(string name, Input<string> id, SyncImageState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/syncImage:SyncImage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SyncImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SyncImage Get(string name, Input<string> id, SyncImageState? state = null, CustomResourceOptions? options = null)
        {
            return new SyncImage(name, id, state, options);
        }
    }

    public sealed class SyncImageArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationRegions", required: true)]
        private InputList<string>? _destinationRegions;

        /// <summary>
        /// List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        /// </summary>
        public InputList<string> DestinationRegions
        {
            get => _destinationRegions ?? (_destinationRegions = new InputList<string>());
            set => _destinationRegions = value;
        }

        /// <summary>
        /// Checks whether image synchronization can be initiated.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// Destination image name.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// Whether to return the ID of image created in the destination region.
        /// </summary>
        [Input("imageSetRequired")]
        public Input<bool>? ImageSetRequired { get; set; }

        public SyncImageArgs()
        {
        }
        public static new SyncImageArgs Empty => new SyncImageArgs();
    }

    public sealed class SyncImageState : global::Pulumi.ResourceArgs
    {
        [Input("destinationRegions")]
        private InputList<string>? _destinationRegions;

        /// <summary>
        /// List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
        /// </summary>
        public InputList<string> DestinationRegions
        {
            get => _destinationRegions ?? (_destinationRegions = new InputList<string>());
            set => _destinationRegions = value;
        }

        /// <summary>
        /// Checks whether image synchronization can be initiated.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Destination image name.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// Whether to return the ID of image created in the destination region.
        /// </summary>
        [Input("imageSetRequired")]
        public Input<bool>? ImageSetRequired { get; set; }

        public SyncImageState()
        {
        }
        public static new SyncImageState Empty => new SyncImageState();
    }
}
