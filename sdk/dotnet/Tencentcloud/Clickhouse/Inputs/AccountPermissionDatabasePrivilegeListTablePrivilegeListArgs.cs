// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse.Inputs
{

    public sealed class AccountPermissionDatabasePrivilegeListTablePrivilegeListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Table name.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("tablePrivileges", required: true)]
        private InputList<string>? _tablePrivileges;

        /// <summary>
        /// Table privileges. Valid values: SELECT, INSERT_ALL, ALTER, TRUNCATE, DROP_TABLE.
        /// </summary>
        public InputList<string> TablePrivileges
        {
            get => _tablePrivileges ?? (_tablePrivileges = new InputList<string>());
            set => _tablePrivileges = value;
        }

        public AccountPermissionDatabasePrivilegeListTablePrivilegeListArgs()
        {
        }
    }
}
