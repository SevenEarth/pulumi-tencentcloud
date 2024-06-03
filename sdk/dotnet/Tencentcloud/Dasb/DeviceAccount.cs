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
    /// Provides a resource to create a dasb device_account
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
    ///     var exampleDevice = new Tencentcloud.Dasb.Device("exampleDevice", new()
    ///     {
    ///         OsName = "Linux",
    ///         Ip = "192.168.0.1",
    ///         Port = 80,
    ///     });
    /// 
    ///     var exampleDeviceAccount = new Tencentcloud.Dasb.DeviceAccount("exampleDeviceAccount", new()
    ///     {
    ///         DeviceId = exampleDevice.Id,
    ///         Account = "root",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// dasb device_account can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Dasb/deviceAccount:DeviceAccount example 11
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dasb/deviceAccount:DeviceAccount")]
    public partial class DeviceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Device account.
        /// </summary>
        [Output("account")]
        public Output<string> Account { get; private set; } = null!;

        /// <summary>
        /// Device ID.
        /// </summary>
        [Output("deviceId")]
        public Output<int> DeviceId { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceAccount(string name, DeviceAccountArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/deviceAccount:DeviceAccount", name, args ?? new DeviceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceAccount(string name, Input<string> id, DeviceAccountState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dasb/deviceAccount:DeviceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceAccount Get(string name, Input<string> id, DeviceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceAccount(name, id, state, options);
        }
    }

    public sealed class DeviceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device account.
        /// </summary>
        [Input("account", required: true)]
        public Input<string> Account { get; set; } = null!;

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<int> DeviceId { get; set; } = null!;

        public DeviceAccountArgs()
        {
        }
        public static new DeviceAccountArgs Empty => new DeviceAccountArgs();
    }

    public sealed class DeviceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Device account.
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// Device ID.
        /// </summary>
        [Input("deviceId")]
        public Input<int>? DeviceId { get; set; }

        public DeviceAccountState()
        {
        }
        public static new DeviceAccountState Empty => new DeviceAccountState();
    }
}
