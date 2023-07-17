// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    /// <summary>
    /// Provides a resource to create a tcr service_account
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// tcr service_account can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tcr/serviceAccount:ServiceAccount service_account registry_id#name
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/serviceAccount:ServiceAccount")]
    public partial class ServiceAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// Service account description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// whether to disable Service accounts.
        /// </summary>
        [Output("disable")]
        public Output<bool?> Disable { get; private set; } = null!;

        /// <summary>
        /// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
        /// </summary>
        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        /// <summary>
        /// Service account expiration time (time stamp, unit: milliseconds).
        /// </summary>
        [Output("expiresAt")]
        public Output<int> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// Service account name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password of the service account.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// strategy list.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.ServiceAccountPermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAccount(string name, ServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/serviceAccount:ServiceAccount", name, args ?? new ServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAccount(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/serviceAccount:ServiceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAccount Get(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceAccount(name, id, state, options);
        }
    }

    public sealed class ServiceAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to disable Service accounts.
        /// </summary>
        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        /// <summary>
        /// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Service account expiration time (time stamp, unit: milliseconds).
        /// </summary>
        [Input("expiresAt")]
        public Input<int>? ExpiresAt { get; set; }

        /// <summary>
        /// Service account name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions", required: true)]
        private InputList<Inputs.ServiceAccountPermissionArgs>? _permissions;

        /// <summary>
        /// strategy list.
        /// </summary>
        public InputList<Inputs.ServiceAccountPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.ServiceAccountPermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ServiceAccountArgs()
        {
        }
    }

    public sealed class ServiceAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to disable Service accounts.
        /// </summary>
        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        /// <summary>
        /// expiration date (unit: day), calculated from the current time, priority is higher than ExpiresAt Service account description.
        /// </summary>
        [Input("duration")]
        public Input<int>? Duration { get; set; }

        /// <summary>
        /// Service account expiration time (time stamp, unit: milliseconds).
        /// </summary>
        [Input("expiresAt")]
        public Input<int>? ExpiresAt { get; set; }

        /// <summary>
        /// Service account name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password of the service account.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("permissions")]
        private InputList<Inputs.ServiceAccountPermissionGetArgs>? _permissions;

        /// <summary>
        /// strategy list.
        /// </summary>
        public InputList<Inputs.ServiceAccountPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.ServiceAccountPermissionGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ServiceAccountState()
        {
        }
    }
}
