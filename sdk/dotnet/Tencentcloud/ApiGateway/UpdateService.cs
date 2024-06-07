// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    /// <summary>
    /// Provides a resource to create a apigateway update_service
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
    ///     var example = new Tencentcloud.ApiGateway.UpdateService("example", new()
    ///     {
    ///         EnvironmentName = "test",
    ///         ServiceId = "service-oczq2nyk",
    ///         VersionName = "20240204142759-b5a4f741-adc0-4964-b01b-2a4a04ff6964",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:ApiGateway/updateService:UpdateService")]
    public partial class UpdateService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the environment to be switched, currently supporting three environments: test (test environment), prepub (pre release environment), and release (release environment).
        /// </summary>
        [Output("environmentName")]
        public Output<string> EnvironmentName { get; private set; } = null!;

        /// <summary>
        /// Service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The version number of the switch.
        /// </summary>
        [Output("versionName")]
        public Output<string> VersionName { get; private set; } = null!;


        /// <summary>
        /// Create a UpdateService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UpdateService(string name, UpdateServiceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/updateService:UpdateService", name, args ?? new UpdateServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UpdateService(string name, Input<string> id, UpdateServiceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/updateService:UpdateService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UpdateService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UpdateService Get(string name, Input<string> id, UpdateServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new UpdateService(name, id, state, options);
        }
    }

    public sealed class UpdateServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the environment to be switched, currently supporting three environments: test (test environment), prepub (pre release environment), and release (release environment).
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Service ID.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// The version number of the switch.
        /// </summary>
        [Input("versionName", required: true)]
        public Input<string> VersionName { get; set; } = null!;

        public UpdateServiceArgs()
        {
        }
        public static new UpdateServiceArgs Empty => new UpdateServiceArgs();
    }

    public sealed class UpdateServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the environment to be switched, currently supporting three environments: test (test environment), prepub (pre release environment), and release (release environment).
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// Service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The version number of the switch.
        /// </summary>
        [Input("versionName")]
        public Input<string>? VersionName { get; set; }

        public UpdateServiceState()
        {
        }
        public static new UpdateServiceState Empty => new UpdateServiceState();
    }
}
