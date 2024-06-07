// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.PrivateDns
{
    /// <summary>
    /// Provide a resource to create a Private Dns Zone.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a basic Private Dns Zone
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
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var example = new Tencentcloud.PrivateDns.Zone("example", new()
    ///     {
    ///         Domain = "domain.com",
    ///         Remark = "remark.",
    ///         VpcSets = new[]
    ///         {
    ///             new Tencentcloud.PrivateDns.Inputs.ZoneVpcSetArgs
    ///             {
    ///                 Region = "ap-guangzhou",
    ///                 UniqVpcId = vpc.Id,
    ///             },
    ///         },
    ///         DnsForwardStatus = "DISABLED",
    ///         CnameSpeedupStatus = "ENABLED",
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create a Private Dns Zone domain and bind associated accounts'VPC
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
    ///     var example = new Tencentcloud.PrivateDns.Zone("example", new()
    ///     {
    ///         Domain = "domain.com",
    ///         Remark = "remark.",
    ///         VpcSets = new[]
    ///         {
    ///             new Tencentcloud.PrivateDns.Inputs.ZoneVpcSetArgs
    ///             {
    ///                 Region = "ap-guangzhou",
    ///                 UniqVpcId = tencentcloud_vpc.Vpc.Id,
    ///             },
    ///         },
    ///         AccountVpcSets = new[]
    ///         {
    ///             new Tencentcloud.PrivateDns.Inputs.ZoneAccountVpcSetArgs
    ///             {
    ///                 Uin = "123456789",
    ///                 UniqVpcId = "vpc-adsebmya",
    ///                 Region = "ap-guangzhou",
    ///                 VpcName = "vpc-name",
    ///             },
    ///         },
    ///         DnsForwardStatus = "DISABLED",
    ///         CnameSpeedupStatus = "ENABLED",
    ///         Tags = 
    ///         {
    ///             { "createdBy", "terraform" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Private Dns Zone can be imported, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:PrivateDns/zone:Zone example zone-6xg5xgky
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:PrivateDns/zone:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of authorized accounts' VPCs to associate with the private domain.
        /// </summary>
        [Output("accountVpcSets")]
        public Output<ImmutableArray<Outputs.ZoneAccountVpcSet>> AccountVpcSets { get; private set; } = null!;

        /// <summary>
        /// CNAME acceleration: ENABLED, DISABLED, Default value is ENABLED.
        /// </summary>
        [Output("cnameSpeedupStatus")]
        public Output<string?> CnameSpeedupStatus { get; private set; } = null!;

        /// <summary>
        /// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        /// </summary>
        [Output("dnsForwardStatus")]
        public Output<string?> DnsForwardStatus { get; private set; } = null!;

        /// <summary>
        /// Domain name, which must be in the format of standard TLD.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Remarks.
        /// </summary>
        [Output("remark")]
        public Output<string?> Remark { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        /// </summary>
        [Output("tagSets")]
        public Output<ImmutableArray<Outputs.ZoneTagSet>> TagSets { get; private set; } = null!;

        /// <summary>
        /// Tags of the private dns zone.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Associates the private domain to a VPC when it is created.
        /// </summary>
        [Output("vpcSets")]
        public Output<ImmutableArray<Outputs.ZoneVpcSet>> VpcSets { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:PrivateDns/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:PrivateDns/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        [Input("accountVpcSets")]
        private InputList<Inputs.ZoneAccountVpcSetArgs>? _accountVpcSets;

        /// <summary>
        /// List of authorized accounts' VPCs to associate with the private domain.
        /// </summary>
        public InputList<Inputs.ZoneAccountVpcSetArgs> AccountVpcSets
        {
            get => _accountVpcSets ?? (_accountVpcSets = new InputList<Inputs.ZoneAccountVpcSetArgs>());
            set => _accountVpcSets = value;
        }

        /// <summary>
        /// CNAME acceleration: ENABLED, DISABLED, Default value is ENABLED.
        /// </summary>
        [Input("cnameSpeedupStatus")]
        public Input<string>? CnameSpeedupStatus { get; set; }

        /// <summary>
        /// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        /// </summary>
        [Input("dnsForwardStatus")]
        public Input<string>? DnsForwardStatus { get; set; }

        /// <summary>
        /// Domain name, which must be in the format of standard TLD.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("tagSets")]
        private InputList<Inputs.ZoneTagSetArgs>? _tagSets;

        /// <summary>
        /// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        /// </summary>
        [Obsolete(@"It has been deprecated from version 1.72.4. Use `tags` instead.")]
        public InputList<Inputs.ZoneTagSetArgs> TagSets
        {
            get => _tagSets ?? (_tagSets = new InputList<Inputs.ZoneTagSetArgs>());
            set => _tagSets = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the private dns zone.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcSets")]
        private InputList<Inputs.ZoneVpcSetArgs>? _vpcSets;

        /// <summary>
        /// Associates the private domain to a VPC when it is created.
        /// </summary>
        public InputList<Inputs.ZoneVpcSetArgs> VpcSets
        {
            get => _vpcSets ?? (_vpcSets = new InputList<Inputs.ZoneVpcSetArgs>());
            set => _vpcSets = value;
        }

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }

    public sealed class ZoneState : global::Pulumi.ResourceArgs
    {
        [Input("accountVpcSets")]
        private InputList<Inputs.ZoneAccountVpcSetGetArgs>? _accountVpcSets;

        /// <summary>
        /// List of authorized accounts' VPCs to associate with the private domain.
        /// </summary>
        public InputList<Inputs.ZoneAccountVpcSetGetArgs> AccountVpcSets
        {
            get => _accountVpcSets ?? (_accountVpcSets = new InputList<Inputs.ZoneAccountVpcSetGetArgs>());
            set => _accountVpcSets = value;
        }

        /// <summary>
        /// CNAME acceleration: ENABLED, DISABLED, Default value is ENABLED.
        /// </summary>
        [Input("cnameSpeedupStatus")]
        public Input<string>? CnameSpeedupStatus { get; set; }

        /// <summary>
        /// Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
        /// </summary>
        [Input("dnsForwardStatus")]
        public Input<string>? DnsForwardStatus { get; set; }

        /// <summary>
        /// Domain name, which must be in the format of standard TLD.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Remarks.
        /// </summary>
        [Input("remark")]
        public Input<string>? Remark { get; set; }

        [Input("tagSets")]
        private InputList<Inputs.ZoneTagSetGetArgs>? _tagSets;

        /// <summary>
        /// It has been deprecated from version 1.72.4. Use `tags` instead. Tags the private domain when it is created.
        /// </summary>
        [Obsolete(@"It has been deprecated from version 1.72.4. Use `tags` instead.")]
        public InputList<Inputs.ZoneTagSetGetArgs> TagSets
        {
            get => _tagSets ?? (_tagSets = new InputList<Inputs.ZoneTagSetGetArgs>());
            set => _tagSets = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the private dns zone.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcSets")]
        private InputList<Inputs.ZoneVpcSetGetArgs>? _vpcSets;

        /// <summary>
        /// Associates the private domain to a VPC when it is created.
        /// </summary>
        public InputList<Inputs.ZoneVpcSetGetArgs> VpcSets
        {
            get => _vpcSets ?? (_vpcSets = new InputList<Inputs.ZoneVpcSetGetArgs>());
            set => _vpcSets = value;
        }

        public ZoneState()
        {
        }
        public static new ZoneState Empty => new ZoneState();
    }
}
