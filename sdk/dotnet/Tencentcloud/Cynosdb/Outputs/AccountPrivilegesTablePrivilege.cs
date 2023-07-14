// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class AccountPrivilegesTablePrivilege
    {
        /// <summary>
        /// Database name.
        /// </summary>
        public readonly string Db;
        /// <summary>
        /// Table privileges.
        /// </summary>
        public readonly ImmutableArray<string> Privileges;
        /// <summary>
        /// Table name.
        /// </summary>
        public readonly string TableName;

        [OutputConstructor]
        private AccountPrivilegesTablePrivilege(
            string db,

            ImmutableArray<string> privileges,

            string tableName)
        {
            Db = db;
            Privileges = privileges;
            TableName = tableName;
        }
    }
}
