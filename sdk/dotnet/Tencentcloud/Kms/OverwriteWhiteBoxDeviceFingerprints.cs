// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kms
{
    /// <summary>
    /// Provides a resource to create a kms overwrite_white_box_device_fingerprints
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
    ///         var example = new Tencentcloud.Kms.OverwriteWhiteBoxDeviceFingerprints("example", new Tencentcloud.Kms.OverwriteWhiteBoxDeviceFingerprintsArgs
    ///         {
    ///             KeyId = "23e80852-1e38-11e9-b129-5cb9019b4b01",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kms/overwriteWhiteBoxDeviceFingerprints:OverwriteWhiteBoxDeviceFingerprints")]
    public partial class OverwriteWhiteBoxDeviceFingerprints : Pulumi.CustomResource
    {
        /// <summary>
        /// Device fingerprint list.
        /// </summary>
        [Output("deviceFingerprints")]
        public Output<ImmutableArray<Outputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprint>> DeviceFingerprints { get; private set; } = null!;

        /// <summary>
        /// CMK unique identifier.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;


        /// <summary>
        /// Create a OverwriteWhiteBoxDeviceFingerprints resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OverwriteWhiteBoxDeviceFingerprints(string name, OverwriteWhiteBoxDeviceFingerprintsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/overwriteWhiteBoxDeviceFingerprints:OverwriteWhiteBoxDeviceFingerprints", name, args ?? new OverwriteWhiteBoxDeviceFingerprintsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OverwriteWhiteBoxDeviceFingerprints(string name, Input<string> id, OverwriteWhiteBoxDeviceFingerprintsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/overwriteWhiteBoxDeviceFingerprints:OverwriteWhiteBoxDeviceFingerprints", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OverwriteWhiteBoxDeviceFingerprints resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OverwriteWhiteBoxDeviceFingerprints Get(string name, Input<string> id, OverwriteWhiteBoxDeviceFingerprintsState? state = null, CustomResourceOptions? options = null)
        {
            return new OverwriteWhiteBoxDeviceFingerprints(name, id, state, options);
        }
    }

    public sealed class OverwriteWhiteBoxDeviceFingerprintsArgs : Pulumi.ResourceArgs
    {
        [Input("deviceFingerprints")]
        private InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs>? _deviceFingerprints;

        /// <summary>
        /// Device fingerprint list.
        /// </summary>
        public InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs> DeviceFingerprints
        {
            get => _deviceFingerprints ?? (_deviceFingerprints = new InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintArgs>());
            set => _deviceFingerprints = value;
        }

        /// <summary>
        /// CMK unique identifier.
        /// </summary>
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public OverwriteWhiteBoxDeviceFingerprintsArgs()
        {
        }
    }

    public sealed class OverwriteWhiteBoxDeviceFingerprintsState : Pulumi.ResourceArgs
    {
        [Input("deviceFingerprints")]
        private InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintGetArgs>? _deviceFingerprints;

        /// <summary>
        /// Device fingerprint list.
        /// </summary>
        public InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintGetArgs> DeviceFingerprints
        {
            get => _deviceFingerprints ?? (_deviceFingerprints = new InputList<Inputs.OverwriteWhiteBoxDeviceFingerprintsDeviceFingerprintGetArgs>());
            set => _deviceFingerprints = value;
        }

        /// <summary>
        /// CMK unique identifier.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public OverwriteWhiteBoxDeviceFingerprintsState()
        {
        }
    }
}
