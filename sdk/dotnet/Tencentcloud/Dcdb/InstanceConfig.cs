// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    /// <summary>
    /// Provides a resource to create a dcdb instance_config
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
    ///         var instanceConfig = new Tencentcloud.Dcdb.InstanceConfig("instanceConfig", new Tencentcloud.Dcdb.InstanceConfigArgs
    ///         {
    ///             InstanceId = local.Dcdb_id,
    ///             RsAccessStrategy = 0,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dcdb instance_config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dcdb/instanceConfig:InstanceConfig instance_config instance_config_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/instanceConfig:InstanceConfig")]
    public partial class InstanceConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// RS nearest access mode, 0-no policy, 1-nearest access.
        /// </summary>
        [Output("rsAccessStrategy")]
        public Output<int> RsAccessStrategy { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceConfig(string name, InstanceConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/instanceConfig:InstanceConfig", name, args ?? new InstanceConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceConfig(string name, Input<string> id, InstanceConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/instanceConfig:InstanceConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceConfig Get(string name, Input<string> id, InstanceConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceConfig(name, id, state, options);
        }
    }

    public sealed class InstanceConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// RS nearest access mode, 0-no policy, 1-nearest access.
        /// </summary>
        [Input("rsAccessStrategy", required: true)]
        public Input<int> RsAccessStrategy { get; set; } = null!;

        public InstanceConfigArgs()
        {
        }
    }

    public sealed class InstanceConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// RS nearest access mode, 0-no policy, 1-nearest access.
        /// </summary>
        [Input("rsAccessStrategy")]
        public Input<int>? RsAccessStrategy { get; set; }

        public InstanceConfigState()
        {
        }
    }
}
