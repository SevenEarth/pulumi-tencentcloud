// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cfs
{
    /// <summary>
    /// Provides a resource to create a cfs sign_up_cfs_service
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
    ///         var signUpCfsService = new Tencentcloud.Cfs.SignUpCfsService("signUpCfsService", new Tencentcloud.Cfs.SignUpCfsServiceArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cfs/signUpCfsService:SignUpCfsService")]
    public partial class SignUpCfsService : Pulumi.CustomResource
    {
        /// <summary>
        /// Current status of the CFS service for this user. Valid values: creating (activating); created (activated).
        /// </summary>
        [Output("cfsServiceStatus")]
        public Output<string> CfsServiceStatus { get; private set; } = null!;


        /// <summary>
        /// Create a SignUpCfsService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SignUpCfsService(string name, SignUpCfsServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/signUpCfsService:SignUpCfsService", name, args ?? new SignUpCfsServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SignUpCfsService(string name, Input<string> id, SignUpCfsServiceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cfs/signUpCfsService:SignUpCfsService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SignUpCfsService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SignUpCfsService Get(string name, Input<string> id, SignUpCfsServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new SignUpCfsService(name, id, state, options);
        }
    }

    public sealed class SignUpCfsServiceArgs : Pulumi.ResourceArgs
    {
        public SignUpCfsServiceArgs()
        {
        }
    }

    public sealed class SignUpCfsServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current status of the CFS service for this user. Valid values: creating (activating); created (activated).
        /// </summary>
        [Input("cfsServiceStatus")]
        public Input<string>? CfsServiceStatus { get; set; }

        public SignUpCfsServiceState()
        {
        }
    }
}
