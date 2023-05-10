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
    public static class GetDataBackupOverview
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql data_backup_overview
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
        ///         var dataBackupOverview = Output.Create(Tencentcloud.Mysql.GetDataBackupOverview.InvokeAsync(new Tencentcloud.Mysql.GetDataBackupOverviewArgs
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
        public static Task<GetDataBackupOverviewResult> InvokeAsync(GetDataBackupOverviewArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataBackupOverviewResult>("tencentcloud:Mysql/getDataBackupOverview:getDataBackupOverview", args ?? new GetDataBackupOverviewArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql data_backup_overview
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
        ///         var dataBackupOverview = Output.Create(Tencentcloud.Mysql.GetDataBackupOverview.InvokeAsync(new Tencentcloud.Mysql.GetDataBackupOverviewArgs
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
        public static Output<GetDataBackupOverviewResult> Invoke(GetDataBackupOverviewInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataBackupOverviewResult>("tencentcloud:Mysql/getDataBackupOverview:getDataBackupOverview", args ?? new GetDataBackupOverviewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataBackupOverviewArgs : Pulumi.InvokeArgs
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

        public GetDataBackupOverviewArgs()
        {
        }
    }

    public sealed class GetDataBackupOverviewInvokeArgs : Pulumi.InvokeArgs
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

        public GetDataBackupOverviewInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataBackupOverviewResult
    {
        /// <summary>
        /// The total number of automatic backups in the current region.
        /// </summary>
        public readonly int AutoBackupCount;
        /// <summary>
        /// The total automatic backup capacity of the current region.
        /// </summary>
        public readonly int AutoBackupVolume;
        /// <summary>
        /// The total number of archive backups in the current region.
        /// </summary>
        public readonly int DataBackupArchiveCount;
        /// <summary>
        /// The total capacity of the current regional archive backup.
        /// </summary>
        public readonly int DataBackupArchiveVolume;
        /// <summary>
        /// The total number of data backups in the current region.
        /// </summary>
        public readonly int DataBackupCount;
        /// <summary>
        /// The total number of standard storage backups in the current region.
        /// </summary>
        public readonly int DataBackupStandbyCount;
        /// <summary>
        /// The total backup capacity of the current regional standard storage.
        /// </summary>
        public readonly int DataBackupStandbyVolume;
        /// <summary>
        /// Total data backup capacity of the current region (including automatic backup and manual backup, in bytes).
        /// </summary>
        public readonly int DataBackupVolume;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The total number of manual backups in the current region.
        /// </summary>
        public readonly int ManualBackupCount;
        /// <summary>
        /// The total manual backup capacity of the current region.
        /// </summary>
        public readonly int ManualBackupVolume;
        public readonly string Product;
        /// <summary>
        /// The total number of remote backups.
        /// </summary>
        public readonly int RemoteBackupCount;
        /// <summary>
        /// The total capacity of remote backup.
        /// </summary>
        public readonly int RemoteBackupVolume;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDataBackupOverviewResult(
            int autoBackupCount,

            int autoBackupVolume,

            int dataBackupArchiveCount,

            int dataBackupArchiveVolume,

            int dataBackupCount,

            int dataBackupStandbyCount,

            int dataBackupStandbyVolume,

            int dataBackupVolume,

            string id,

            int manualBackupCount,

            int manualBackupVolume,

            string product,

            int remoteBackupCount,

            int remoteBackupVolume,

            string? resultOutputFile)
        {
            AutoBackupCount = autoBackupCount;
            AutoBackupVolume = autoBackupVolume;
            DataBackupArchiveCount = dataBackupArchiveCount;
            DataBackupArchiveVolume = dataBackupArchiveVolume;
            DataBackupCount = dataBackupCount;
            DataBackupStandbyCount = dataBackupStandbyCount;
            DataBackupStandbyVolume = dataBackupStandbyVolume;
            DataBackupVolume = dataBackupVolume;
            Id = id;
            ManualBackupCount = manualBackupCount;
            ManualBackupVolume = manualBackupVolume;
            Product = product;
            RemoteBackupCount = remoteBackupCount;
            RemoteBackupVolume = remoteBackupVolume;
            ResultOutputFile = resultOutputFile;
        }
    }
}
