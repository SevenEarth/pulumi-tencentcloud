// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provides a resource to create a vpc flow_log_config
    /// 
    /// ## Example Usage
    /// ### If enable FlowLogs
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Tencentcloud.Vpc.FlowLogConfig("config", new Tencentcloud.Vpc.FlowLogConfigArgs
    ///         {
    ///             FlowLogId = tencentcloud_vpc_flow_log.Example.Id,
    ///             Enable = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc flow_log_config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/flowLogConfig:FlowLogConfig flow_log_config flow_log_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/flowLogConfig:FlowLogConfig")]
    public partial class FlowLogConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// If enable snapshot policy.
        /// </summary>
        [Output("enable")]
        public Output<bool> Enable { get; private set; } = null!;

        /// <summary>
        /// Flow log ID.
        /// </summary>
        [Output("flowLogId")]
        public Output<string> FlowLogId { get; private set; } = null!;


        /// <summary>
        /// Create a FlowLogConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowLogConfig(string name, FlowLogConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/flowLogConfig:FlowLogConfig", name, args ?? new FlowLogConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlowLogConfig(string name, Input<string> id, FlowLogConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/flowLogConfig:FlowLogConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FlowLogConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowLogConfig Get(string name, Input<string> id, FlowLogConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new FlowLogConfig(name, id, state, options);
        }
    }

    public sealed class FlowLogConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If enable snapshot policy.
        /// </summary>
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        /// <summary>
        /// Flow log ID.
        /// </summary>
        [Input("flowLogId", required: true)]
        public Input<string> FlowLogId { get; set; } = null!;

        public FlowLogConfigArgs()
        {
        }
    }

    public sealed class FlowLogConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If enable snapshot policy.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Flow log ID.
        /// </summary>
        [Input("flowLogId")]
        public Input<string>? FlowLogId { get; set; }

        public FlowLogConfigState()
        {
        }
    }
}