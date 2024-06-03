// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    public static class GetBackupByFlowId
    {
        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_backup_by_flow_id
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var zones = Tencentcloud.Availability.GetZonesByProduct.Invoke(new()
        ///     {
        ///         Product = "sqlserver",
        ///     });
        /// 
        ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
        ///     {
        ///         CidrBlock = "10.0.0.0/16",
        ///     });
        /// 
        ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
        ///     {
        ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[4]?.Name),
        ///         VpcId = vpc.Id,
        ///         CidrBlock = "10.0.0.0/16",
        ///         IsMulticast = false,
        ///     });
        /// 
        ///     var securityGroup = new Tencentcloud.Security.Group("securityGroup", new()
        ///     {
        ///         Description = "desc.",
        ///     });
        /// 
        ///     var exampleBasicInstance = new Tencentcloud.Sqlserver.BasicInstance("exampleBasicInstance", new()
        ///     {
        ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[4]?.Name),
        ///         ChargeType = "POSTPAID_BY_HOUR",
        ///         VpcId = vpc.Id,
        ///         SubnetId = subnet.Id,
        ///         ProjectId = 0,
        ///         Memory = 4,
        ///         Storage = 100,
        ///         Cpu = 2,
        ///         MachineType = "CLOUD_PREMIUM",
        ///         MaintenanceWeekSets = new[]
        ///         {
        ///             1,
        ///             2,
        ///             3,
        ///         },
        ///         MaintenanceStartTime = "09:00",
        ///         MaintenanceTimeSpan = 3,
        ///         SecurityGroups = new[]
        ///         {
        ///             securityGroup.Id,
        ///         },
        ///         Tags = 
        ///         {
        ///             { "test", "test" },
        ///         },
        ///     });
        /// 
        ///     var exampleDb = new Tencentcloud.Sqlserver.Db("exampleDb", new()
        ///     {
        ///         InstanceId = exampleBasicInstance.Id,
        ///         Charset = "Chinese_PRC_BIN",
        ///         Remark = "test-remark",
        ///     });
        /// 
        ///     var exampleGeneralBackup = new Tencentcloud.Sqlserver.GeneralBackup("exampleGeneralBackup", new()
        ///     {
        ///         InstanceId = exampleDb.Id,
        ///         BackupName = "tf_example_backup",
        ///         Strategy = 0,
        ///     });
        /// 
        ///     var exampleBackupByFlowId = Tencentcloud.Sqlserver.GetBackupByFlowId.Invoke(new()
        ///     {
        ///         InstanceId = exampleGeneralBackup.InstanceId,
        ///         FlowId = exampleGeneralBackup.FlowId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBackupByFlowIdResult> InvokeAsync(GetBackupByFlowIdArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackupByFlowIdResult>("tencentcloud:Sqlserver/getBackupByFlowId:getBackupByFlowId", args ?? new GetBackupByFlowIdArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of sqlserver datasource_backup_by_flow_id
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var zones = Tencentcloud.Availability.GetZonesByProduct.Invoke(new()
        ///     {
        ///         Product = "sqlserver",
        ///     });
        /// 
        ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
        ///     {
        ///         CidrBlock = "10.0.0.0/16",
        ///     });
        /// 
        ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
        ///     {
        ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[4]?.Name),
        ///         VpcId = vpc.Id,
        ///         CidrBlock = "10.0.0.0/16",
        ///         IsMulticast = false,
        ///     });
        /// 
        ///     var securityGroup = new Tencentcloud.Security.Group("securityGroup", new()
        ///     {
        ///         Description = "desc.",
        ///     });
        /// 
        ///     var exampleBasicInstance = new Tencentcloud.Sqlserver.BasicInstance("exampleBasicInstance", new()
        ///     {
        ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[4]?.Name),
        ///         ChargeType = "POSTPAID_BY_HOUR",
        ///         VpcId = vpc.Id,
        ///         SubnetId = subnet.Id,
        ///         ProjectId = 0,
        ///         Memory = 4,
        ///         Storage = 100,
        ///         Cpu = 2,
        ///         MachineType = "CLOUD_PREMIUM",
        ///         MaintenanceWeekSets = new[]
        ///         {
        ///             1,
        ///             2,
        ///             3,
        ///         },
        ///         MaintenanceStartTime = "09:00",
        ///         MaintenanceTimeSpan = 3,
        ///         SecurityGroups = new[]
        ///         {
        ///             securityGroup.Id,
        ///         },
        ///         Tags = 
        ///         {
        ///             { "test", "test" },
        ///         },
        ///     });
        /// 
        ///     var exampleDb = new Tencentcloud.Sqlserver.Db("exampleDb", new()
        ///     {
        ///         InstanceId = exampleBasicInstance.Id,
        ///         Charset = "Chinese_PRC_BIN",
        ///         Remark = "test-remark",
        ///     });
        /// 
        ///     var exampleGeneralBackup = new Tencentcloud.Sqlserver.GeneralBackup("exampleGeneralBackup", new()
        ///     {
        ///         InstanceId = exampleDb.Id,
        ///         BackupName = "tf_example_backup",
        ///         Strategy = 0,
        ///     });
        /// 
        ///     var exampleBackupByFlowId = Tencentcloud.Sqlserver.GetBackupByFlowId.Invoke(new()
        ///     {
        ///         InstanceId = exampleGeneralBackup.InstanceId,
        ///         FlowId = exampleGeneralBackup.FlowId,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBackupByFlowIdResult> Invoke(GetBackupByFlowIdInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackupByFlowIdResult>("tencentcloud:Sqlserver/getBackupByFlowId:getBackupByFlowId", args ?? new GetBackupByFlowIdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackupByFlowIdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Create a backup process ID, which can be obtained through the [CreateBackup](https://cloud.tencent.com/document/product/238/19946) interface.
        /// </summary>
        [Input("flowId", required: true)]
        public string FlowId { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetBackupByFlowIdArgs()
        {
        }
        public static new GetBackupByFlowIdArgs Empty => new GetBackupByFlowIdArgs();
    }

    public sealed class GetBackupByFlowIdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Create a backup process ID, which can be obtained through the [CreateBackup](https://cloud.tencent.com/document/product/238/19946) interface.
        /// </summary>
        [Input("flowId", required: true)]
        public Input<string> FlowId { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetBackupByFlowIdInvokeArgs()
        {
        }
        public static new GetBackupByFlowIdInvokeArgs Empty => new GetBackupByFlowIdInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackupByFlowIdResult
    {
        /// <summary>
        /// Backup task name, customizable.
        /// </summary>
        public readonly string BackupName;
        /// <summary>
        /// Backup method, 0-scheduled backup; 1-manual temporary backup; instance status is 0-creating, this field is the default value 0, meaningless.
        /// </summary>
        public readonly int BackupWay;
        /// <summary>
        /// For the DB list, only the library name contained in the first record is returned for a single-database backup file; for a single-database backup file, the library names of all records need to be obtained through the DescribeBackupFiles interface.
        /// </summary>
        public readonly ImmutableArray<string> Dbs;
        /// <summary>
        /// backup end time.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// External network download address, for a single database backup file, only the external network download address of the first record is returned; single database backup files need to obtain the download addresses of all records through the DescribeBackupFiles interface.
        /// </summary>
        public readonly string ExternalAddr;
        /// <summary>
        /// File name. For a single-database backup file, only the file name of the first record is returned; for a single-database backup file, the file names of all records need to be obtained through the DescribeBackupFiles interface.
        /// </summary>
        public readonly string FileName;
        public readonly string FlowId;
        /// <summary>
        /// Aggregate Id, this value is not returned for packaged backup files. Use this value to call the DescribeBackupFiles interface to obtain the detailed information of a single database backup file.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// Intranet download address, for a single database backup file, only the intranet download address of the first record is returned; single database backup files need to obtain the download addresses of all records through the DescribeBackupFiles interface.
        /// </summary>
        public readonly string InternalAddr;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// backup start time.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Backup file status, 0-creating; 1-success; 2-failure.
        /// </summary>
        public readonly int Status;
        /// <summary>
        /// Backup strategy, 0-instance backup; 1-multi-database backup; when the instance status is 0-creating, this field is the default value 0, meaningless.
        /// </summary>
        public readonly int Strategy;

        [OutputConstructor]
        private GetBackupByFlowIdResult(
            string backupName,

            int backupWay,

            ImmutableArray<string> dbs,

            string endTime,

            string externalAddr,

            string fileName,

            string flowId,

            string groupId,

            string id,

            string instanceId,

            string internalAddr,

            string? resultOutputFile,

            string startTime,

            int status,

            int strategy)
        {
            BackupName = backupName;
            BackupWay = backupWay;
            Dbs = dbs;
            EndTime = endTime;
            ExternalAddr = externalAddr;
            FileName = fileName;
            FlowId = flowId;
            GroupId = groupId;
            Id = id;
            InstanceId = instanceId;
            InternalAddr = internalAddr;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            Status = status;
            Strategy = strategy;
        }
    }
}
