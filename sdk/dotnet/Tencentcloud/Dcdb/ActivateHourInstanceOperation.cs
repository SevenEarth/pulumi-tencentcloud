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
    /// Provides a resource to create a dcdb activate_hour_instance_operation
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
    ///     var activateHourInstanceOperation = new Tencentcloud.Dcdb.ActivateHourInstanceOperation("activateHourInstanceOperation", new()
    ///     {
    ///         InstanceId = local.Dcdb_id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/activateHourInstanceOperation:ActivateHourInstanceOperation")]
    public partial class ActivateHourInstanceOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ActivateHourInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActivateHourInstanceOperation(string name, ActivateHourInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/activateHourInstanceOperation:ActivateHourInstanceOperation", name, args ?? new ActivateHourInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActivateHourInstanceOperation(string name, Input<string> id, ActivateHourInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/activateHourInstanceOperation:ActivateHourInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActivateHourInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActivateHourInstanceOperation Get(string name, Input<string> id, ActivateHourInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ActivateHourInstanceOperation(name, id, state, options);
        }
    }

    public sealed class ActivateHourInstanceOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public ActivateHourInstanceOperationArgs()
        {
        }
        public static new ActivateHourInstanceOperationArgs Empty => new ActivateHourInstanceOperationArgs();
    }

    public sealed class ActivateHourInstanceOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance ID in the format of dcdbt-ow728lmc, which can be obtained through the `DescribeDCDBInstances` API.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public ActivateHourInstanceOperationState()
        {
        }
        public static new ActivateHourInstanceOperationState Empty => new ActivateHourInstanceOperationState();
    }
}
