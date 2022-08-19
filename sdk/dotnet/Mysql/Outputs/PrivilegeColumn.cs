// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class PrivilegeColumn
    {
        /// <summary>
        /// Column name.
        /// </summary>
        public readonly string ColumnName;
        /// <summary>
        /// Database name.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.
        /// </summary>
        public readonly ImmutableArray<string> Privileges;
        /// <summary>
        /// Table name.
        /// </summary>
        public readonly string TableName;

        [OutputConstructor]
        private PrivilegeColumn(
            string columnName,

            string databaseName,

            ImmutableArray<string> privileges,

            string tableName)
        {
            ColumnName = columnName;
            DatabaseName = databaseName;
            Privileges = privileges;
            TableName = tableName;
        }
    }
}
