// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    /// <summary>
    /// Provides a resource to create a mariadb renew_instance
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
    ///         var renewInstance = new Tencentcloud.Mariadb.RenewInstance("renewInstance", new Tencentcloud.Mariadb.RenewInstanceArgs
    ///         {
    ///             InstanceId = "tdsql-9vqvls95",
    ///             Period = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mariadb/renewInstance:RenewInstance")]
    public partial class RenewInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Renewal duration, unit: month.
        /// </summary>
        [Output("period")]
        public Output<int> Period { get; private set; } = null!;


        /// <summary>
        /// Create a RenewInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RenewInstance(string name, RenewInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/renewInstance:RenewInstance", name, args ?? new RenewInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RenewInstance(string name, Input<string> id, RenewInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/renewInstance:RenewInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RenewInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RenewInstance Get(string name, Input<string> id, RenewInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new RenewInstance(name, id, state, options);
        }
    }

    public sealed class RenewInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Renewal duration, unit: month.
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        public RenewInstanceArgs()
        {
        }
    }

    public sealed class RenewInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Renewal duration, unit: month.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        public RenewInstanceState()
        {
        }
    }
}
