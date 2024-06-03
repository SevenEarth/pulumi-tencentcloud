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
    /// Provides a resource to create a dlc upgrade_data_engine_image_operation
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
    ///     var upgradeDataEngineImageOperation = new Tencentcloud.Dlc.UpgradeDataEngineImageOperation("upgradeDataEngineImageOperation", new()
    ///     {
    ///         DataEngineId = "DataEngine-g5ds87d8",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dlc/upgradeDataEngineImageOperation:UpgradeDataEngineImageOperation")]
    public partial class UpgradeDataEngineImageOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Output("dataEngineId")]
        public Output<string> DataEngineId { get; private set; } = null!;


        /// <summary>
        /// Create a UpgradeDataEngineImageOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpgradeDataEngineImageOperation(string name, UpgradeDataEngineImageOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/upgradeDataEngineImageOperation:UpgradeDataEngineImageOperation", name, args ?? new UpgradeDataEngineImageOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpgradeDataEngineImageOperation(string name, Input<string> id, UpgradeDataEngineImageOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dlc/upgradeDataEngineImageOperation:UpgradeDataEngineImageOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpgradeDataEngineImageOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpgradeDataEngineImageOperation Get(string name, Input<string> id, UpgradeDataEngineImageOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new UpgradeDataEngineImageOperation(name, id, state, options);
        }
    }

    public sealed class UpgradeDataEngineImageOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Input("dataEngineId", required: true)]
        public Input<string> DataEngineId { get; set; } = null!;

        public UpgradeDataEngineImageOperationArgs()
        {
        }
        public static new UpgradeDataEngineImageOperationArgs Empty => new UpgradeDataEngineImageOperationArgs();
    }

    public sealed class UpgradeDataEngineImageOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine unique id.
        /// </summary>
        [Input("dataEngineId")]
        public Input<string>? DataEngineId { get; set; }

        public UpgradeDataEngineImageOperationState()
        {
        }
        public static new UpgradeDataEngineImageOperationState Empty => new UpgradeDataEngineImageOperationState();
    }
}
