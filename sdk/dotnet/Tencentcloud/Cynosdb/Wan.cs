// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    /// <summary>
    /// Provides a resource to create a cynosdb wan
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
    ///         var wan = new Tencentcloud.Cynosdb.Wan("wan", new Tencentcloud.Cynosdb.WanArgs
    ///         {
    ///             ClusterId = "cynosdbmysql-bws8h88b",
    ///             InstanceGrpId = "cynosdbmysql-grp-lxav0p9z",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cynosdb wan can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cynosdb/wan:Wan wan cynosdbmysql-bws8h88b#cynosdbmysql-grp-lxav0p9z
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cynosdb/wan:Wan")]
    public partial class Wan : Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Instance Group ID.
        /// </summary>
        [Output("instanceGrpId")]
        public Output<string> InstanceGrpId { get; private set; } = null!;

        /// <summary>
        /// Domain name.
        /// </summary>
        [Output("wanDomain")]
        public Output<string> WanDomain { get; private set; } = null!;

        /// <summary>
        /// Network ip.
        /// </summary>
        [Output("wanIp")]
        public Output<string> WanIp { get; private set; } = null!;

        /// <summary>
        /// Internet port.
        /// </summary>
        [Output("wanPort")]
        public Output<int> WanPort { get; private set; } = null!;

        /// <summary>
        /// Internet status.
        /// </summary>
        [Output("wanStatus")]
        public Output<string> WanStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Wan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wan(string name, WanArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/wan:Wan", name, args ?? new WanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wan(string name, Input<string> id, WanState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cynosdb/wan:Wan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Wan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wan Get(string name, Input<string> id, WanState? state = null, CustomResourceOptions? options = null)
        {
            return new Wan(name, id, state, options);
        }
    }

    public sealed class WanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Instance Group ID.
        /// </summary>
        [Input("instanceGrpId", required: true)]
        public Input<string> InstanceGrpId { get; set; } = null!;

        public WanArgs()
        {
        }
    }

    public sealed class WanState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Instance Group ID.
        /// </summary>
        [Input("instanceGrpId")]
        public Input<string>? InstanceGrpId { get; set; }

        /// <summary>
        /// Domain name.
        /// </summary>
        [Input("wanDomain")]
        public Input<string>? WanDomain { get; set; }

        /// <summary>
        /// Network ip.
        /// </summary>
        [Input("wanIp")]
        public Input<string>? WanIp { get; set; }

        /// <summary>
        /// Internet port.
        /// </summary>
        [Input("wanPort")]
        public Input<int>? WanPort { get; set; }

        /// <summary>
        /// Internet status.
        /// </summary>
        [Input("wanStatus")]
        public Input<string>? WanStatus { get; set; }

        public WanState()
        {
        }
    }
}
