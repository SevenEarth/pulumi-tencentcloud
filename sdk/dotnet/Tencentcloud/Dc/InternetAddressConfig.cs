// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dc
{
    /// <summary>
    /// Provides a resource to create a dc internet_address_config
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
    ///         var internetAddress = new Tencentcloud.Dc.InternetAddress("internetAddress", new Tencentcloud.Dc.InternetAddressArgs
    ///         {
    ///             MaskLen = 30,
    ///             AddrType = 2,
    ///             AddrProto = 0,
    ///         });
    ///         var internetAddressConfig = new Tencentcloud.Dc.InternetAddressConfig("internetAddressConfig", new Tencentcloud.Dc.InternetAddressConfigArgs
    ///         {
    ///             InstanceId = internetAddress.Id,
    ///             Enable = false,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dc internet_address_config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dc/internetAddressConfig:InternetAddressConfig internet_address_config internet_address_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dc/internetAddressConfig:InternetAddressConfig")]
    public partial class InternetAddressConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// whether enable internet address.
        /// </summary>
        [Output("enable")]
        public Output<bool> Enable { get; private set; } = null!;

        /// <summary>
        /// internet public address id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a InternetAddressConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InternetAddressConfig(string name, InternetAddressConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/internetAddressConfig:InternetAddressConfig", name, args ?? new InternetAddressConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InternetAddressConfig(string name, Input<string> id, InternetAddressConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dc/internetAddressConfig:InternetAddressConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InternetAddressConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InternetAddressConfig Get(string name, Input<string> id, InternetAddressConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new InternetAddressConfig(name, id, state, options);
        }
    }

    public sealed class InternetAddressConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether enable internet address.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// internet public address id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public InternetAddressConfigArgs()
        {
        }
    }

    public sealed class InternetAddressConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether enable internet address.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// internet public address id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public InternetAddressConfigState()
        {
        }
    }
}
