// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc
{
    /// <summary>
    /// Provides a resource to create a dlc restart_data_engine
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
    ///     var restartDataEngine = new Tencentcloud.Dlc.RestartDataEngineOperation("restartDataEngine", new()
    ///     {
    ///         DataEngineId = "DataEngine-g5ds87d8",
    ///         ForcedOperation = false,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/restartDataEngineOperation:RestartDataEngineOperation")]
    public partial class RestartDataEngineOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Output("dataEngineId")]
        public Output<string> DataEngineId { get; private set; } = null!;

        /// <summary>
        /// Whether to force restart and ignore tasks.
        /// </summary>
        [Output("forcedOperation")]
        public Output<bool?> ForcedOperation { get; private set; } = null!;


        /// <summary>
        /// Create a RestartDataEngineOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartDataEngineOperation(string name, RestartDataEngineOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/restartDataEngineOperation:RestartDataEngineOperation", name, args ?? new RestartDataEngineOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartDataEngineOperation(string name, Input<string> id, RestartDataEngineOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/restartDataEngineOperation:RestartDataEngineOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestartDataEngineOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartDataEngineOperation Get(string name, Input<string> id, RestartDataEngineOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RestartDataEngineOperation(name, id, state, options);
        }
    }

    public sealed class RestartDataEngineOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Input("dataEngineId", required: true)]
        public Input<string> DataEngineId { get; set; } = null!;

        /// <summary>
        /// Whether to force restart and ignore tasks.
        /// </summary>
        [Input("forcedOperation")]
        public Input<bool>? ForcedOperation { get; set; }

        public RestartDataEngineOperationArgs()
        {
        }
        public static new RestartDataEngineOperationArgs Empty => new RestartDataEngineOperationArgs();
    }

    public sealed class RestartDataEngineOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Input("dataEngineId")]
        public Input<string>? DataEngineId { get; set; }

        /// <summary>
        /// Whether to force restart and ignore tasks.
        /// </summary>
        [Input("forcedOperation")]
        public Input<bool>? ForcedOperation { get; set; }

        public RestartDataEngineOperationState()
        {
        }
        public static new RestartDataEngineOperationState Empty => new RestartDataEngineOperationState();
    }
}
