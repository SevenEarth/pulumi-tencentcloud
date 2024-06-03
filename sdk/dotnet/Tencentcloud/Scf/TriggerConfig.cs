// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf
{
    /// <summary>
    /// Provides a resource to create a scf trigger_config
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
    ///     var triggerConfig = new Tencentcloud.Scf.TriggerConfig("triggerConfig", new()
    ///     {
    ///         CustomArgument = "Information",
    ///         Description = "func",
    ///         Enable = "OPEN",
    ///         FunctionName = "keep-1676351130",
    ///         Namespace = "default",
    ///         Qualifier = "$DEFAULT",
    ///         TriggerDesc = "* 1 2 * * * *",
    ///         TriggerName = "SCF-timer-1685540160",
    ///         Type = "timer",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// scf trigger_config can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Scf/triggerConfig:TriggerConfig trigger_config functionName#namespace#triggerName
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Scf/triggerConfig:TriggerConfig")]
    public partial class TriggerConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User Additional Information.
        /// </summary>
        [Output("customArgument")]
        public Output<string?> CustomArgument { get; private set; } = null!;

        /// <summary>
        /// Trigger description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Status of trigger. Values: OPEN (enabled); CLOSE disabled).
        /// </summary>
        [Output("enable")]
        public Output<string?> Enable { get; private set; } = null!;

        /// <summary>
        /// Function name.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        /// </summary>
        [Output("qualifier")]
        public Output<string?> Qualifier { get; private set; } = null!;

        /// <summary>
        /// TriggerDesc parameter.
        /// </summary>
        [Output("triggerDesc")]
        public Output<string> TriggerDesc { get; private set; } = null!;

        /// <summary>
        /// Trigger Name.
        /// </summary>
        [Output("triggerName")]
        public Output<string> TriggerName { get; private set; } = null!;

        /// <summary>
        /// Trigger type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerConfig(string name, TriggerConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/triggerConfig:TriggerConfig", name, args ?? new TriggerConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerConfig(string name, Input<string> id, TriggerConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/triggerConfig:TriggerConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TriggerConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerConfig Get(string name, Input<string> id, TriggerConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerConfig(name, id, state, options);
        }
    }

    public sealed class TriggerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User Additional Information.
        /// </summary>
        [Input("customArgument")]
        public Input<string>? CustomArgument { get; set; }

        /// <summary>
        /// Trigger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of trigger. Values: OPEN (enabled); CLOSE disabled).
        /// </summary>
        [Input("enable")]
        public Input<string>? Enable { get; set; }

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// TriggerDesc parameter.
        /// </summary>
        [Input("triggerDesc")]
        public Input<string>? TriggerDesc { get; set; }

        /// <summary>
        /// Trigger Name.
        /// </summary>
        [Input("triggerName", required: true)]
        public Input<string> TriggerName { get; set; } = null!;

        /// <summary>
        /// Trigger type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TriggerConfigArgs()
        {
        }
        public static new TriggerConfigArgs Empty => new TriggerConfigArgs();
    }

    public sealed class TriggerConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User Additional Information.
        /// </summary>
        [Input("customArgument")]
        public Input<string>? CustomArgument { get; set; }

        /// <summary>
        /// Trigger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of trigger. Values: OPEN (enabled); CLOSE disabled).
        /// </summary>
        [Input("enable")]
        public Input<string>? Enable { get; set; }

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Function version. It defaults to `$LATEST`. It's recommended to use `[$DEFAULT](https://intl.cloud.tencent.com/document/product/583/36149?from_cn_redirect=1#.E9.BB.98.E8.AE.A4.E5.88.AB.E5.90.8D)` for canary release.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// TriggerDesc parameter.
        /// </summary>
        [Input("triggerDesc")]
        public Input<string>? TriggerDesc { get; set; }

        /// <summary>
        /// Trigger Name.
        /// </summary>
        [Input("triggerName")]
        public Input<string>? TriggerName { get; set; }

        /// <summary>
        /// Trigger type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TriggerConfigState()
        {
        }
        public static new TriggerConfigState Empty => new TriggerConfigState();
    }
}
