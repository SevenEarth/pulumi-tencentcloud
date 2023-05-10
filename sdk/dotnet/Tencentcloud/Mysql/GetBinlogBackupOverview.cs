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
    public static class GetBinlogBackupOverview
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql binlog_backup_overview
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
        ///         var binlogBackupOverview = Output.Create(Tencentcloud.Mysql.GetBinlogBackupOverview.InvokeAsync(new Tencentcloud.Mysql.GetBinlogBackupOverviewArgs
        ///         {
        ///             Product = "mysql",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBinlogBackupOverviewResult> InvokeAsync(GetBinlogBackupOverviewArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBinlogBackupOverviewResult>("tencentcloud:Mysql/getBinlogBackupOverview:getBinlogBackupOverview", args ?? new GetBinlogBackupOverviewArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql binlog_backup_overview
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
        ///         var binlogBackupOverview = Output.Create(Tencentcloud.Mysql.GetBinlogBackupOverview.InvokeAsync(new Tencentcloud.Mysql.GetBinlogBackupOverviewArgs
        ///         {
        ///             Product = "mysql",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBinlogBackupOverviewResult> Invoke(GetBinlogBackupOverviewInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBinlogBackupOverviewResult>("tencentcloud:Mysql/getBinlogBackupOverview:getBinlogBackupOverview", args ?? new GetBinlogBackupOverviewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBinlogBackupOverviewArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of cloud database product to be queried, currently only supports `mysql`.
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBinlogBackupOverviewArgs()
        {
        }
    }

    public sealed class GetBinlogBackupOverviewInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The type of cloud database product to be queried, currently only supports `mysql`.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBinlogBackupOverviewInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBinlogBackupOverviewResult
    {
        /// <summary>
        /// The number of archived log backups.
        /// </summary>
        public readonly int BinlogArchiveCount;
        /// <summary>
        /// Archived log backup capacity (in bytes).
        /// </summary>
        public readonly int BinlogArchiveVolume;
        /// <summary>
        /// The total number of log backups, including remote log backups.
        /// </summary>
        public readonly int BinlogBackupCount;
        /// <summary>
        /// Total log backup capacity, including off-site log backup (unit is byte).
        /// </summary>
        public readonly int BinlogBackupVolume;
        /// <summary>
        /// The number of standard storage log backups.
        /// </summary>
        public readonly int BinlogStandbyCount;
        /// <summary>
        /// Standard storage log backup capacity (in bytes).
        /// </summary>
        public readonly int BinlogStandbyVolume;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Product;
        /// <summary>
        /// The number of remote log backups.
        /// </summary>
        public readonly int RemoteBinlogCount;
        /// <summary>
        /// Remote log backup capacity (in bytes).
        /// </summary>
        public readonly int RemoteBinlogVolume;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetBinlogBackupOverviewResult(
            int binlogArchiveCount,

            int binlogArchiveVolume,

            int binlogBackupCount,

            int binlogBackupVolume,

            int binlogStandbyCount,

            int binlogStandbyVolume,

            string id,

            string product,

            int remoteBinlogCount,

            int remoteBinlogVolume,

            string? resultOutputFile)
        {
            BinlogArchiveCount = binlogArchiveCount;
            BinlogArchiveVolume = binlogArchiveVolume;
            BinlogBackupCount = binlogBackupCount;
            BinlogBackupVolume = binlogBackupVolume;
            BinlogStandbyCount = binlogStandbyCount;
            BinlogStandbyVolume = binlogStandbyVolume;
            Id = id;
            Product = product;
            RemoteBinlogCount = remoteBinlogCount;
            RemoteBinlogVolume = remoteBinlogVolume;
            ResultOutputFile = resultOutputFile;
        }
    }
}
