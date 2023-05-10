// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class MigrateJobMigrateOptionDatabaseTableDatabaseTable
    {
        /// <summary>
        /// new table name.
        /// </summary>
        public readonly string? NewTableName;
        /// <summary>
        /// table edit mode.
        /// </summary>
        public readonly string? TableEditMode;
        /// <summary>
        /// table name.
        /// </summary>
        public readonly string? TableName;
        /// <summary>
        /// temporary tables.
        /// </summary>
        public readonly ImmutableArray<string> TmpTables;

        [OutputConstructor]
        private MigrateJobMigrateOptionDatabaseTableDatabaseTable(
            string? newTableName,

            string? tableEditMode,

            string? tableName,

            ImmutableArray<string> tmpTables)
        {
            NewTableName = newTableName;
            TableEditMode = tableEditMode;
            TableName = tableName;
            TmpTables = tmpTables;
        }
    }
}
