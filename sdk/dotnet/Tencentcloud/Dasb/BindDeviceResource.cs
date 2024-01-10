// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dasb
{
    /// <summary>
    /// Provides a resource to create a dasb bind_device_resource
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
    ///         var example = new Tencentcloud.Dasb.BindDeviceResource("example", new Tencentcloud.Dasb.BindDeviceResourceArgs
    ///         {
    ///             DeviceIdSets = 
    ///             {
    ///                 17,
    ///                 18,
    ///             },
    ///             ResourceId = "bh-saas-ocmzo6lgxiv",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dasb/bindDeviceResource:BindDeviceResource")]
    public partial class BindDeviceResource : Pulumi.CustomResource
    {
        /// <summary>
        /// Asset ID collection.
        /// </summary>
        [Output("deviceIdSets")]
        public Output<ImmutableArray<int>> DeviceIdSets { get; private set; } = null!;

        /// <summary>
        /// Bastion host service ID.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a BindDeviceResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BindDeviceResource(string name, BindDeviceResourceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/bindDeviceResource:BindDeviceResource", name, args ?? new BindDeviceResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BindDeviceResource(string name, Input<string> id, BindDeviceResourceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/bindDeviceResource:BindDeviceResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BindDeviceResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BindDeviceResource Get(string name, Input<string> id, BindDeviceResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new BindDeviceResource(name, id, state, options);
        }
    }

    public sealed class BindDeviceResourceArgs : Pulumi.ResourceArgs
    {
        [Input("deviceIdSets", required: true)]
        private InputList<int>? _deviceIdSets;

        /// <summary>
        /// Asset ID collection.
        /// </summary>
        public InputList<int> DeviceIdSets
        {
            get => _deviceIdSets ?? (_deviceIdSets = new InputList<int>());
            set => _deviceIdSets = value;
        }

        /// <summary>
        /// Bastion host service ID.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        public BindDeviceResourceArgs()
        {
        }
    }

    public sealed class BindDeviceResourceState : Pulumi.ResourceArgs
    {
        [Input("deviceIdSets")]
        private InputList<int>? _deviceIdSets;

        /// <summary>
        /// Asset ID collection.
        /// </summary>
        public InputList<int> DeviceIdSets
        {
            get => _deviceIdSets ?? (_deviceIdSets = new InputList<int>());
            set => _deviceIdSets = value;
        }

        /// <summary>
        /// Bastion host service ID.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public BindDeviceResourceState()
        {
        }
    }
}
