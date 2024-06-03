// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    /// <summary>
    /// Provides a resource to create a redis backup_download_restriction
    /// 
    /// ## Example Usage
    /// 
    /// ### Modify the network information and address of the current region backup file download
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
    ///     var foo = new Tencentcloud.Redis.BackupDownloadRestriction("foo", new()
    ///     {
    ///         LimitType = "Customize",
    ///         VpcComparisonSymbol = "In",
    ///         IpComparisonSymbol = "In",
    ///         LimitVpcs = new[]
    ///         {
    ///             new Tencentcloud.Redis.Inputs.BackupDownloadRestrictionLimitVpcArgs
    ///             {
    ///                 Region = "ap-guangzhou",
    ///                 VpcLists = new[]
    ///                 {
    ///                     @var.Vpc_id,
    ///                 },
    ///             },
    ///         },
    ///         LimitIps = new[]
    ///         {
    ///             "10.1.1.12",
    ///             "10.1.1.13",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// redis backup_download_restriction can be imported using the region, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction foo ap-guangzhou
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction")]
    public partial class BackupDownloadRestriction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
        /// </summary>
        [Output("ipComparisonSymbol")]
        public Output<string?> IpComparisonSymbol { get; private set; } = null!;

        /// <summary>
        /// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        [Output("limitIps")]
        public Output<ImmutableArray<string>> LimitIps { get; private set; } = null!;

        /// <summary>
        /// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
        /// </summary>
        [Output("limitType")]
        public Output<string> LimitType { get; private set; } = null!;

        /// <summary>
        /// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        [Output("limitVpcs")]
        public Output<ImmutableArray<Outputs.BackupDownloadRestrictionLimitVpc>> LimitVpcs { get; private set; } = null!;

        /// <summary>
        /// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
        /// </summary>
        [Output("vpcComparisonSymbol")]
        public Output<string?> VpcComparisonSymbol { get; private set; } = null!;


        /// <summary>
        /// Create a BackupDownloadRestriction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupDownloadRestriction(string name, BackupDownloadRestrictionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction", name, args ?? new BackupDownloadRestrictionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupDownloadRestriction(string name, Input<string> id, BackupDownloadRestrictionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackupDownloadRestriction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupDownloadRestriction Get(string name, Input<string> id, BackupDownloadRestrictionState? state = null, CustomResourceOptions? options = null)
        {
            return new BackupDownloadRestriction(name, id, state, options);
        }
    }

    public sealed class BackupDownloadRestrictionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
        /// </summary>
        [Input("ipComparisonSymbol")]
        public Input<string>? IpComparisonSymbol { get; set; }

        [Input("limitIps")]
        private InputList<string>? _limitIps;

        /// <summary>
        /// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public InputList<string> LimitIps
        {
            get => _limitIps ?? (_limitIps = new InputList<string>());
            set => _limitIps = value;
        }

        /// <summary>
        /// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
        /// </summary>
        [Input("limitType", required: true)]
        public Input<string> LimitType { get; set; } = null!;

        [Input("limitVpcs")]
        private InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs>? _limitVpcs;

        /// <summary>
        /// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public Input<string>? VpcComparisonSymbol { get; set; }

        public BackupDownloadRestrictionArgs()
        {
        }
        public static new BackupDownloadRestrictionArgs Empty => new BackupDownloadRestrictionArgs();
    }

    public sealed class BackupDownloadRestrictionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
        /// </summary>
        [Input("ipComparisonSymbol")]
        public Input<string>? IpComparisonSymbol { get; set; }

        [Input("limitIps")]
        private InputList<string>? _limitIps;

        /// <summary>
        /// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public InputList<string> LimitIps
        {
            get => _limitIps ?? (_limitIps = new InputList<string>());
            set => _limitIps = value;
        }

        /// <summary>
        /// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
        /// </summary>
        [Input("limitType")]
        public Input<string>? LimitType { get; set; }

        [Input("limitVpcs")]
        private InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs>? _limitVpcs;

        /// <summary>
        /// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public Input<string>? VpcComparisonSymbol { get; set; }

        public BackupDownloadRestrictionState()
        {
        }
        public static new BackupDownloadRestrictionState Empty => new BackupDownloadRestrictionState();
    }
}
