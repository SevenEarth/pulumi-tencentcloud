// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Chdfs
{
    /// <summary>
    /// Provides a resource to create a chdfs mount_point
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
    ///     var mountPoint = new Tencentcloud.Chdfs.MountPoint("mountPoint", new()
    ///     {
    ///         FileSystemId = "f14mpfy5lh4e",
    ///         MountPointName = "terraform-test",
    ///         MountPointStatus = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// chdfs mount_point can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Chdfs/mountPoint:MountPoint mount_point mount_point_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Chdfs/mountPoint:MountPoint")]
    public partial class MountPoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// file system id you want to mount.
        /// </summary>
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        /// <summary>
        /// mount point name.
        /// </summary>
        [Output("mountPointName")]
        public Output<string> MountPointName { get; private set; } = null!;

        /// <summary>
        /// mount status 1:open, 2:close.
        /// </summary>
        [Output("mountPointStatus")]
        public Output<int> MountPointStatus { get; private set; } = null!;


        /// <summary>
        /// Create a MountPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MountPoint(string name, MountPointArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Chdfs/mountPoint:MountPoint", name, args ?? new MountPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MountPoint(string name, Input<string> id, MountPointState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Chdfs/mountPoint:MountPoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MountPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MountPoint Get(string name, Input<string> id, MountPointState? state = null, CustomResourceOptions? options = null)
        {
            return new MountPoint(name, id, state, options);
        }
    }

    public sealed class MountPointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// file system id you want to mount.
        /// </summary>
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        /// <summary>
        /// mount point name.
        /// </summary>
        [Input("mountPointName", required: true)]
        public Input<string> MountPointName { get; set; } = null!;

        /// <summary>
        /// mount status 1:open, 2:close.
        /// </summary>
        [Input("mountPointStatus", required: true)]
        public Input<int> MountPointStatus { get; set; } = null!;

        public MountPointArgs()
        {
        }
        public static new MountPointArgs Empty => new MountPointArgs();
    }

    public sealed class MountPointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// file system id you want to mount.
        /// </summary>
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        /// <summary>
        /// mount point name.
        /// </summary>
        [Input("mountPointName")]
        public Input<string>? MountPointName { get; set; }

        /// <summary>
        /// mount status 1:open, 2:close.
        /// </summary>
        [Input("mountPointStatus")]
        public Input<int>? MountPointStatus { get; set; }

        public MountPointState()
        {
        }
        public static new MountPointState Empty => new MountPointState();
    }
}
