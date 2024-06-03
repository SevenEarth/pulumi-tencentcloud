// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Inputs
{

    public sealed class MigrationSourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the migration source Cvm, used when MigrateType=2 (cloud server self-built SQL Server database).
        /// </summary>
        [Input("cvmId")]
        public Input<string>? CvmId { get; set; }

        /// <summary>
        /// The ID of the migration source instance, which is used when MigrateType=1 (TencentDB for SQLServers). The format is mssql-si2823jyl.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Migrate the intranet IP of the self-built database of the source Cvm, and use it when MigrateType=2 (self-built SQL Server database of the cloud server).
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Password, MigrateType=1 or MigrateType=2.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The port number of the self-built database of the migration source Cvm, which is used when MigrateType=2 (self-built SQL Server database of the cloud server).
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The subnet ID under the Vpc of the source Cvm is used when MigrateType=2 (ECS self-built SQL Server database). The format is as follows subnet-h9extioi.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The source backup password for offline migration, MigrateType=4 or MigrateType=5.
        /// </summary>
        [Input("urlPassword")]
        public Input<string>? UrlPassword { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// The source backup address for offline migration. MigrateType=4 or MigrateType=5.
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        /// <summary>
        /// User name, MigrateType=1 or MigrateType=2.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// The Vpc network ID of the migration source Cvm is used when MigrateType=2 (cloud server self-built SQL Server database). The format is as follows vpc-6ys9ont9.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public MigrationSourceGetArgs()
        {
        }
        public static new MigrationSourceGetArgs Empty => new MigrationSourceGetArgs();
    }
}
