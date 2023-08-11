// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class AccountPrivilegesColumnPrivilege
    {
        /// <summary>
        /// Column name.
        /// </summary>
        public readonly string Column;
        /// <summary>
        /// Database name.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Permission information.
        /// </summary>
        public readonly ImmutableArray<string> Privileges;
        /// <summary>
        /// Table name.
        /// </summary>
        public readonly string Table;

        [OutputConstructor]
        private AccountPrivilegesColumnPrivilege(
            string column,

            string database,

            ImmutableArray<string> privileges,

            string table)
        {
            Column = column;
            Database = database;
            Privileges = privileges;
            Table = table;
        }
    }
}