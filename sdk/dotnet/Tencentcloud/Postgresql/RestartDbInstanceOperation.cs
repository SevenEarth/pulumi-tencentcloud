// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Provides a resource to create a postgresql restart_db_instance_operation
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
    ///     var restartDbInstanceOperation = new Tencentcloud.Postgresql.RestartDbInstanceOperation("restartDbInstanceOperation", new()
    ///     {
    ///         DbInstanceId = local.Pgsql_id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation")]
    public partial class RestartDbInstanceOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// dbInstance ID.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a RestartDbInstanceOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RestartDbInstanceOperation(string name, RestartDbInstanceOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation", name, args ?? new RestartDbInstanceOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RestartDbInstanceOperation(string name, Input<string> id, RestartDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/restartDbInstanceOperation:RestartDbInstanceOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RestartDbInstanceOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RestartDbInstanceOperation Get(string name, Input<string> id, RestartDbInstanceOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new RestartDbInstanceOperation(name, id, state, options);
        }
    }

    public sealed class RestartDbInstanceOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// dbInstance ID.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        public RestartDbInstanceOperationArgs()
        {
        }
        public static new RestartDbInstanceOperationArgs Empty => new RestartDbInstanceOperationArgs();
    }

    public sealed class RestartDbInstanceOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// dbInstance ID.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        public RestartDbInstanceOperationState()
        {
        }
        public static new RestartDbInstanceOperationState Empty => new RestartDbInstanceOperationState();
    }
}
