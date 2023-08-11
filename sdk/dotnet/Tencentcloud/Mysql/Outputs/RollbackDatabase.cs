// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class RollbackDatabase
    {
        /// <summary>
        /// The original database name before rollback.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The new database name after rollback.
        /// </summary>
        public readonly string NewDatabaseName;

        [OutputConstructor]
        private RollbackDatabase(
            string databaseName,

            string newDatabaseName)
        {
            DatabaseName = databaseName;
            NewDatabaseName = newDatabaseName;
        }
    }
}