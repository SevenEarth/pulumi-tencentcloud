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
    public static class GetBackupDownloadInfo
    {
        /// <summary>
        /// Use this data source to query detailed information of redis backup_download_info
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var backupDownloadInfo = Output.Create(Tencentcloud.Redis.GetBackupDownloadInfo.InvokeAsync(new Tencentcloud.Redis.GetBackupDownloadInfoArgs
        ///         {
        ///             BackupId = "641186639-8362913-1516672770",
        ///             InstanceId = "crs-iw7d9wdd",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackupDownloadInfoResult> InvokeAsync(GetBackupDownloadInfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackupDownloadInfoResult>("tencentcloud:Redis/getBackupDownloadInfo:getBackupDownloadInfo", args ?? new GetBackupDownloadInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of redis backup_download_info
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var backupDownloadInfo = Output.Create(Tencentcloud.Redis.GetBackupDownloadInfo.InvokeAsync(new Tencentcloud.Redis.GetBackupDownloadInfoArgs
        ///         {
        ///             BackupId = "641186639-8362913-1516672770",
        ///             InstanceId = "crs-iw7d9wdd",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackupDownloadInfoResult> Invoke(GetBackupDownloadInfoInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackupDownloadInfoResult>("tencentcloud:Redis/getBackupDownloadInfo:getBackupDownloadInfo", args ?? new GetBackupDownloadInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupDownloadInfoArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backup ID, which can be accessed via [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) interface returns the parameter RedisBackupSet to get.
        /// </summary>
        [Input("backupId", required: true)]
        public string BackupId { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
        /// </summary>
        [Input("ipComparisonSymbol")]
        public string? IpComparisonSymbol { get; set; }

        [Input("limitIps")]
        private List<string>? _limitIps;

        /// <summary>
        /// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public List<string> LimitIps
        {
            get => _limitIps ?? (_limitIps = new List<string>());
            set => _limitIps = value;
        }

        /// <summary>
        /// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
        /// </summary>
        [Input("limitType")]
        public string? LimitType { get; set; }

        [Input("limitVpcs")]
        private List<Inputs.GetBackupDownloadInfoLimitVpcArgs>? _limitVpcs;

        /// <summary>
        /// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public List<Inputs.GetBackupDownloadInfoLimitVpcArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new List<Inputs.GetBackupDownloadInfoLimitVpcArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public string? VpcComparisonSymbol { get; set; }

        public GetBackupDownloadInfoArgs()
        {
        }
    }

    public sealed class GetBackupDownloadInfoInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The backup ID, which can be accessed via [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) interface returns the parameter RedisBackupSet to get.
        /// </summary>
        [Input("backupId", required: true)]
        public Input<string> BackupId { get; set; } = null!;

        /// <summary>
        /// The ID of instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

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
        private InputList<Inputs.GetBackupDownloadInfoLimitVpcInputArgs>? _limitVpcs;

        /// <summary>
        /// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
        /// </summary>
        public InputList<Inputs.GetBackupDownloadInfoLimitVpcInputArgs> LimitVpcs
        {
            get => _limitVpcs ?? (_limitVpcs = new InputList<Inputs.GetBackupDownloadInfoLimitVpcInputArgs>());
            set => _limitVpcs = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
        /// </summary>
        [Input("vpcComparisonSymbol")]
        public Input<string>? VpcComparisonSymbol { get; set; }

        public GetBackupDownloadInfoInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackupDownloadInfoResult
    {
        public readonly string BackupId;
        /// <summary>
        /// A list of backup file information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBackupDownloadInfoBackupInfoResult> BackupInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? IpComparisonSymbol;
        public readonly ImmutableArray<string> LimitIps;
        public readonly string? LimitType;
        public readonly ImmutableArray<Outputs.GetBackupDownloadInfoLimitVpcResult> LimitVpcs;
        public readonly string? ResultOutputFile;
        public readonly string? VpcComparisonSymbol;

        [OutputConstructor]
        private GetBackupDownloadInfoResult(
            string backupId,

            ImmutableArray<Outputs.GetBackupDownloadInfoBackupInfoResult> backupInfos,

            string id,

            string instanceId,

            string? ipComparisonSymbol,

            ImmutableArray<string> limitIps,

            string? limitType,

            ImmutableArray<Outputs.GetBackupDownloadInfoLimitVpcResult> limitVpcs,

            string? resultOutputFile,

            string? vpcComparisonSymbol)
        {
            BackupId = backupId;
            BackupInfos = backupInfos;
            Id = id;
            InstanceId = instanceId;
            IpComparisonSymbol = ipComparisonSymbol;
            LimitIps = limitIps;
            LimitType = limitType;
            LimitVpcs = limitVpcs;
            ResultOutputFile = resultOutputFile;
            VpcComparisonSymbol = vpcComparisonSymbol;
        }
    }
}
