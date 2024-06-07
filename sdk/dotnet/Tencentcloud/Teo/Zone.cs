// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo
{
    /// <summary>
    /// Provides a resource to create a teo zone
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
    ///     var zone = new Tencentcloud.Teo.Zone("zone", new()
    ///     {
    ///         AliasZoneName = "teo-test",
    ///         Area = "overseas",
    ///         Paused = false,
    ///         PlanId = "edgeone-2kfv1h391n6w",
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///         Type = "partial",
    ///         ZoneName = "tf-teo.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// teo zone can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Teo/zone:Zone zone zone_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Teo/zone:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alias site identifier. Limit the input to a combination of numbers, English, - and _, within 20 characters. For details, refer to the alias site identifier. If there is no such usage scenario, leave this field empty.
        /// </summary>
        [Output("aliasZoneName")]
        public Output<string?> AliasZoneName { get; private set; } = null!;

        /// <summary>
        /// When the Type value is partial/full, the acceleration region of the L7 domain name. The following are the values of this parameter, and the default value is overseas if not filled in. When the Type value is noDomainAccess, please leave this value empty:
        /// - global: Global availability zone.
        /// - mainland: Chinese mainland availability zone.
        /// - overseas: Global availability zone (excluding Chinese mainland).
        /// </summary>
        [Output("area")]
        public Output<string> Area { get; private set; } = null!;

        /// <summary>
        /// NS list allocated by Tencent Cloud.
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// Ownership verification information. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Output("ownershipVerifications")]
        public Output<ImmutableArray<Outputs.ZoneOwnershipVerification>> OwnershipVerifications { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the site is disabled.
        /// </summary>
        [Output("paused")]
        public Output<bool> Paused { get; private set; } = null!;

        /// <summary>
        /// The target Plan ID to be bound. When you have an existing Plan in your account, you can fill in this parameter to directly bind the site to the Plan. If you do not have a Plan that can be bound at the moment, please go to the console to purchase a Plan to complete the site creation.
        /// </summary>
        [Output("planId")]
        public Output<string> PlanId { get; private set; } = null!;

        /// <summary>
        /// Site status. Valid values: `active`: NS is switched; `pending`: NS is not switched; `moved`: NS is moved; `deactivated`: this site is blocked.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Site access type. The value of this parameter is as follows, and the default is partial if not filled in:partial: CNAME access; full: NS access; noDomainAccess: No domain access.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Site name. When accessing CNAME/NS, please pass the second-level domain (example.com) as the site name; when accessing without a domain name, please leave this value empty.
        /// </summary>
        [Output("zoneName")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias site identifier. Limit the input to a combination of numbers, English, - and _, within 20 characters. For details, refer to the alias site identifier. If there is no such usage scenario, leave this field empty.
        /// </summary>
        [Input("aliasZoneName")]
        public Input<string>? AliasZoneName { get; set; }

        /// <summary>
        /// When the Type value is partial/full, the acceleration region of the L7 domain name. The following are the values of this parameter, and the default value is overseas if not filled in. When the Type value is noDomainAccess, please leave this value empty:
        /// - global: Global availability zone.
        /// - mainland: Chinese mainland availability zone.
        /// - overseas: Global availability zone (excluding Chinese mainland).
        /// </summary>
        [Input("area", required: true)]
        public Input<string> Area { get; set; } = null!;

        /// <summary>
        /// Indicates whether the site is disabled.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// The target Plan ID to be bound. When you have an existing Plan in your account, you can fill in this parameter to directly bind the site to the Plan. If you do not have a Plan that can be bound at the moment, please go to the console to purchase a Plan to complete the site creation.
        /// </summary>
        [Input("planId", required: true)]
        public Input<string> PlanId { get; set; } = null!;

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

        /// <summary>
        /// Site access type. The value of this parameter is as follows, and the default is partial if not filled in:partial: CNAME access; full: NS access; noDomainAccess: No domain access.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Site name. When accessing CNAME/NS, please pass the second-level domain (example.com) as the site name; when accessing without a domain name, please leave this value empty.
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }

    public sealed class ZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias site identifier. Limit the input to a combination of numbers, English, - and _, within 20 characters. For details, refer to the alias site identifier. If there is no such usage scenario, leave this field empty.
        /// </summary>
        [Input("aliasZoneName")]
        public Input<string>? AliasZoneName { get; set; }

        /// <summary>
        /// When the Type value is partial/full, the acceleration region of the L7 domain name. The following are the values of this parameter, and the default value is overseas if not filled in. When the Type value is noDomainAccess, please leave this value empty:
        /// - global: Global availability zone.
        /// - mainland: Chinese mainland availability zone.
        /// - overseas: Global availability zone (excluding Chinese mainland).
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// NS list allocated by Tencent Cloud.
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        [Input("ownershipVerifications")]
        private InputList<Inputs.ZoneOwnershipVerificationGetArgs>? _ownershipVerifications;

        /// <summary>
        /// Ownership verification information. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.ZoneOwnershipVerificationGetArgs> OwnershipVerifications
        {
            get => _ownershipVerifications ?? (_ownershipVerifications = new InputList<Inputs.ZoneOwnershipVerificationGetArgs>());
            set => _ownershipVerifications = value;
        }

        /// <summary>
        /// Indicates whether the site is disabled.
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// The target Plan ID to be bound. When you have an existing Plan in your account, you can fill in this parameter to directly bind the site to the Plan. If you do not have a Plan that can be bound at the moment, please go to the console to purchase a Plan to complete the site creation.
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        /// <summary>
        /// Site status. Valid values: `active`: NS is switched; `pending`: NS is not switched; `moved`: NS is moved; `deactivated`: this site is blocked.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

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

        /// <summary>
        /// Site access type. The value of this parameter is as follows, and the default is partial if not filled in:partial: CNAME access; full: NS access; noDomainAccess: No domain access.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Site name. When accessing CNAME/NS, please pass the second-level domain (example.com) as the site name; when accessing without a domain name, please leave this value empty.
        /// </summary>
        [Input("zoneName")]
        public Input<string>? ZoneName { get; set; }

        public ZoneState()
        {
        }
        public static new ZoneState Empty => new ZoneState();
    }
}
