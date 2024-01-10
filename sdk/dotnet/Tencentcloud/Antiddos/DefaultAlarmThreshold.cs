// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos
{
    /// <summary>
    /// Provides a resource to create a antiddos default alarm threshold
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
    ///         var defaultAlarmThreshold = new Tencentcloud.Antiddos.DefaultAlarmThreshold("defaultAlarmThreshold", new Tencentcloud.Antiddos.DefaultAlarmThresholdArgs
    ///         {
    ///             DefaultAlarmConfig = new Tencentcloud.Antiddos.Inputs.DefaultAlarmThresholdDefaultAlarmConfigArgs
    ///             {
    ///                 AlarmThreshold = 2000,
    ///                 AlarmType = 1,
    ///             },
    ///             InstanceType = "bgp",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// antiddos default_alarm_threshold can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold default_alarm_threshold ${instanceType}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold")]
    public partial class DefaultAlarmThreshold : Pulumi.CustomResource
    {
        /// <summary>
        /// Alarm threshold configuration.
        /// </summary>
        [Output("defaultAlarmConfig")]
        public Output<Outputs.DefaultAlarmThresholdDefaultAlarmConfig> DefaultAlarmConfig { get; private set; } = null!;

        /// <summary>
        /// Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultAlarmThreshold resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultAlarmThreshold(string name, DefaultAlarmThresholdArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold", name, args ?? new DefaultAlarmThresholdArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultAlarmThreshold(string name, Input<string> id, DefaultAlarmThresholdState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/defaultAlarmThreshold:DefaultAlarmThreshold", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultAlarmThreshold resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultAlarmThreshold Get(string name, Input<string> id, DefaultAlarmThresholdState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultAlarmThreshold(name, id, state, options);
        }
    }

    public sealed class DefaultAlarmThresholdArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm threshold configuration.
        /// </summary>
        [Input("defaultAlarmConfig", required: true)]
        public Input<Inputs.DefaultAlarmThresholdDefaultAlarmConfigArgs> DefaultAlarmConfig { get; set; } = null!;

        /// <summary>
        /// Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        public DefaultAlarmThresholdArgs()
        {
        }
    }

    public sealed class DefaultAlarmThresholdState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm threshold configuration.
        /// </summary>
        [Input("defaultAlarmConfig")]
        public Input<Inputs.DefaultAlarmThresholdDefaultAlarmConfigGetArgs>? DefaultAlarmConfig { get; set; }

        /// <summary>
        /// Product type, value [bgp (represents advanced defense package product) bgpip (represents advanced defense IP product)].
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public DefaultAlarmThresholdState()
        {
        }
    }
}
