// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a resource to create a mysql backup_download_restriction
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
    ///         var backupDownloadRestriction = new Tencentcloud.Mysql.BackupDownloadRestriction("backupDownloadRestriction", new Tencentcloud.Mysql.BackupDownloadRestrictionArgs
    ///         {
    ///             IpComparisonSymbol = "In",
    ///             LimitIps = 
    ///             {
    ///                 "127.0.0.1",
    ///             },
    ///             LimitType = "Customize",
    ///             LimitVpcs = 
    ///             {
    ///                 new Tencentcloud.Mysql.Inputs.BackupDownloadRestrictionLimitVpcArgs
    ///                 {
    ///                     Region = "ap-guangzhou",
    ///                     VpcLists = 
    ///                     {
    ///                         "vpc-4owdpnwr",
    ///                     },
    ///                 },
    ///             },
    ///             VpcComparisonSymbol = "In",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mysql backup_download_restriction can be imported using the "BackupDownloadRestriction", as follows.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction backup_download_restriction BackupDownloadRestriction
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction")]
    public partial class BackupDownloadRestriction : Pulumi.CustomResource
    {
        /// <summary>
        /// In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
        /// </summary>
        [Output("ipComparisonSymbol")]
        public Output<string?> IpComparisonSymbol { get; private set; } = null!;

        /// <summary>
        /// ip settings to limit downloads.
        /// </summary>
        [Output("limitIps")]
        public Output<ImmutableArray<string>> LimitIps { get; private set; } = null!;

        /// <summary>
        /// NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
        /// </summary>
        [Output("limitType")]
        public Output<string> LimitType { get; private set; } = null!;

        /// <summary>
        /// vpc settings to limit downloads.
        /// </summary>
        [Output("limitVpcs")]
        public Output<ImmutableArray<Outputs.BackupDownloadRestrictionLimitVpc>> LimitVpcs { get; private set; } = null!;

        /// <summary>
        /// This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
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
            : base("tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction", name, args ?? new BackupDownloadRestrictionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupDownloadRestriction(string name, Input<string> id, BackupDownloadRestrictionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/backupDownloadRestriction:BackupDownloadRestriction", name, state, MakeResourceOptions(options, id))
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

    public sealed class BackupDownloadRestrictionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
        /// </summary>
        [Input("ipComparisonSymbol")]
        public Input<string>? IpComparisonSymbol { get; set; }

        [Input("limitIps")]
        private InputList<string>? _limitIps;

        /// <summary>
        /// ip settings to limit downloads.
        /// </summary>
        public InputList<string> LimitIps
        {
            get => _limitIps ?? (_limitIps = new InputList<string>());
            set => _limitIps = value;
        }

        /// <summary>
        /// NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
        /// </summary>
        [Input("limitType", required: true)]
        public Input<string> LimitType { get; set; } = null!;

        [Input("limitVpcs")]
        private InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs>? _limitVpcs;

        /// <summary>
        /// vpc settings to limit downloads.
        /// </summary>
        public InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new InputList<Inputs.BackupDownloadRestrictionLimitVpcArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public Input<string>? VpcComparisonSymbol { get; set; }

        public BackupDownloadRestrictionArgs()
        {
        }
    }

    public sealed class BackupDownloadRestrictionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// In: The specified ip can be downloaded; NotIn: The specified ip cannot be downloaded. The default is In.
        /// </summary>
        [Input("ipComparisonSymbol")]
        public Input<string>? IpComparisonSymbol { get; set; }

        [Input("limitIps")]
        private InputList<string>? _limitIps;

        /// <summary>
        /// ip settings to limit downloads.
        /// </summary>
        public InputList<string> LimitIps
        {
            get => _limitIps ?? (_limitIps = new InputList<string>());
            set => _limitIps = value;
        }

        /// <summary>
        /// NoLimit No limit, both internal and external networks can be downloaded; LimitOnlyIntranet Only intranet can be downloaded; Customize user-defined vpc:ip can be downloaded. LimitVpc and LimitIp can be set only when the value is Customize.
        /// </summary>
        [Input("limitType")]
        public Input<string>? LimitType { get; set; }

        [Input("limitVpcs")]
        private InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs>? _limitVpcs;

        /// <summary>
        /// vpc settings to limit downloads.
        /// </summary>
        public InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new InputList<Inputs.BackupDownloadRestrictionLimitVpcGetArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// This parameter only supports In, which means that the vpc specified by LimitVpc can be downloaded. The default is In.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public Input<string>? VpcComparisonSymbol { get; set; }

        public BackupDownloadRestrictionState()
        {
        }
    }
}
