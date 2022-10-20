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
    /// Provide a resource to create a KMS external key.
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
    ///         var foo = new Tencentcloud.Kms.ExternalKey("foo", new Tencentcloud.Kms.ExternalKeyArgs
    ///         {
    ///             Alias = "test",
    ///             Description = "describe key test message.",
    ///             IsEnabled = true,
    ///             KeyMaterialBase64 = "MTIzMTIzMTIzMTIzMTIzQQ==",
    ///             ValidTo = 2147443200,
    ///             WrappingAlgorithm = "RSAES_PKCS1_V1_5",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// KMS external keys can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Kms/externalKey:ExternalKey foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kms/externalKey:ExternalKey")]
    public partial class ExternalKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Description of CMK. The maximum is 1024 bytes.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Output("isArchived")]
        public Output<bool?> IsArchived { get; private set; } = null!;

        /// <summary>
        /// Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool?> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        /// </summary>
        [Output("keyMaterialBase64")]
        public Output<string?> KeyMaterialBase64 { get; private set; } = null!;

        /// <summary>
        /// State of CMK.
        /// </summary>
        [Output("keyState")]
        public Output<string> KeyState { get; private set; } = null!;

        /// <summary>
        /// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        /// </summary>
        [Output("pendingDeleteWindowInDays")]
        public Output<int?> PendingDeleteWindowInDays { get; private set; } = null!;

        /// <summary>
        /// Tags of CMK.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        /// </summary>
        [Output("validTo")]
        public Output<int?> ValidTo { get; private set; } = null!;

        /// <summary>
        /// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        /// </summary>
        [Output("wrappingAlgorithm")]
        public Output<string?> WrappingAlgorithm { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalKey(string name, ExternalKeyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/externalKey:ExternalKey", name, args ?? new ExternalKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalKey(string name, Input<string> id, ExternalKeyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kms/externalKey:ExternalKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExternalKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalKey Get(string name, Input<string> id, ExternalKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalKey(name, id, state, options);
        }
    }

    public sealed class ExternalKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Description of CMK. The maximum is 1024 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Input("isArchived")]
        public Input<bool>? IsArchived { get; set; }

        /// <summary>
        /// Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        /// </summary>
        [Input("keyMaterialBase64")]
        public Input<string>? KeyMaterialBase64 { get; set; }

        /// <summary>
        /// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        /// </summary>
        [Input("pendingDeleteWindowInDays")]
        public Input<int>? PendingDeleteWindowInDays { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of CMK.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        /// </summary>
        [Input("validTo")]
        public Input<int>? ValidTo { get; set; }

        /// <summary>
        /// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        /// </summary>
        [Input("wrappingAlgorithm")]
        public Input<string>? WrappingAlgorithm { get; set; }

        public ExternalKeyArgs()
        {
        }
    }

    public sealed class ExternalKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Description of CMK. The maximum is 1024 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Input("isArchived")]
        public Input<bool>? IsArchived { get; set; }

        /// <summary>
        /// Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
        /// </summary>
        [Input("keyMaterialBase64")]
        public Input<string>? KeyMaterialBase64 { get; set; }

        /// <summary>
        /// State of CMK.
        /// </summary>
        [Input("keyState")]
        public Input<string>? KeyState { get; set; }

        /// <summary>
        /// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
        /// </summary>
        [Input("pendingDeleteWindowInDays")]
        public Input<int>? PendingDeleteWindowInDays { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of CMK.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
        /// </summary>
        [Input("validTo")]
        public Input<int>? ValidTo { get; set; }

        /// <summary>
        /// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
        /// </summary>
        [Input("wrappingAlgorithm")]
        public Input<string>? WrappingAlgorithm { get; set; }

        public ExternalKeyState()
        {
        }
    }
}
