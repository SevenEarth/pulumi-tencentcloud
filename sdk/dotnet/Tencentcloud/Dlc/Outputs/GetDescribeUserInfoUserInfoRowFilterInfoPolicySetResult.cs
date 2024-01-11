// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc.Outputs
{

    [OutputType]
    public sealed class GetDescribeUserInfoUserInfoRowFilterInfoPolicySetResult
    {
        /// <summary>
        /// For the data source name that requires authorization, only * (representing all resources at this level) is supported under the administrator level; in the case of data source level and database level authentication, only COSDataCatalog or * is supported; in data table level authentication, it is possible Fill in the user-defined data source. If left blank, it defaults to DataLakeCatalog. note: If a user-defined data source is authenticated, the permissions that dlc can manage are a subset of the accounts provided by the user when accessing the data source.
        /// </summary>
        public readonly string Catalog;
        /// <summary>
        /// For columns that require authorization, fill in * to represent all current columns. When the authorization type is administrator level, only * is allowed.
        /// </summary>
        public readonly string Column;
        /// <summary>
        /// The time the workgroup was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Data engines that require authorization, fill in * to represent all current engines. when the authorization type is administrator level, only * is allowed.
        /// </summary>
        public readonly string DataEngine;
        /// <summary>
        /// Database name that requires authorization, fill in * to represent all databases under the current catalog. When the authorization type is administrator level, only * is allowed to be filled in. when the authorization type is data connection level, only blanks are allowed to be filled in. For other types, the database can be specified arbitrarily.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// For the function name that requires authorization, fill in * to represent all functions under the current catalog. when the authorization type is administrator level, only * is allowed to be filled in. When the authorization type is data connection level, only blanks are allowed to be filled in. in other types, functions can be specified arbitrarily.
        /// </summary>
        public readonly string Function;
        /// <summary>
        /// Policy id.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Authorization mode, please leave this parameter blank. COMMON: normal mode; SENIOR: advanced mode.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Authorized permission operations provide different operations for different levels of authentication. administrator permissions: ALL, default is ALL if left blank; data connection level authentication: CREATE; database level authentication: ALL, CREATE, ALTER, DROP; data table permissions: ALL, SELECT, INSERT, ALTER, DELETE, DROP, UPDATE. note: under data table permissions, only SELECT operations are supported when the specified data source is not COSDataCatalog.
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// Operator, do not fill in the input parameters.
        /// </summary>
        public readonly string Operator;
        /// <summary>
        /// Authorization type, currently supports eight authorization types: ADMIN: Administrator level authentication DATASOURCE: data connection level authentication DATABASE: database level authentication TABLE: Table level authentication VIEW: view level authentication FUNCTION: Function level authentication COLUMN: Column level authentication ENGINE: Data engine authentication. if left blank, the default is administrator level authentication.
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// Whether the user can perform secondary authorization. when it is true, the authorized user can re-authorize the permissions obtained this time to other sub-users. default is false.
        /// </summary>
        public readonly bool ReAuth;
        /// <summary>
        /// Permission source, please leave it blank. USER: permissions come from the user itself; WORKGROUP: permissions come from the bound workgroup.
        /// </summary>
        public readonly string Source;
        /// <summary>
        /// The id of the workgroup to which the permission belongs. this value only exists when the source of the permission is a workgroup. that is, this field has a value only when the value of the Source field is WORKGROUP.
        /// </summary>
        public readonly int SourceId;
        /// <summary>
        /// The name of the workgroup to which the permission belongs. this value only exists when the source of the permission is a workgroup. that is, this field has a value only when the value of the source field is WORKGROUP.
        /// </summary>
        public readonly string SourceName;
        /// <summary>
        /// For the table name that requires authorization, fill in * to represent all tables under the current database. when the authorization type is administrator level, only * is allowed to be filled in. when the authorization type is data connection level or database level, only blanks are allowed to be filled in. For other types, data tables can be specified arbitrarily.
        /// </summary>
        public readonly string Table;
        /// <summary>
        /// For views that require authorization, fill in * to represent all views under the current database. When the authorization type is administrator level, only * is allowed to be filled in. when the authorization type is data connection level or database level, only blanks are allowed to be filled in. for other types, the view can be specified arbitrarily.
        /// </summary>
        public readonly string View;

        [OutputConstructor]
        private GetDescribeUserInfoUserInfoRowFilterInfoPolicySetResult(
            string catalog,

            string column,

            string createTime,

            string dataEngine,

            string database,

            string function,

            int id,

            string mode,

            string operation,

            string @operator,

            string policyType,

            bool reAuth,

            string source,

            int sourceId,

            string sourceName,

            string table,

            string view)
        {
            Catalog = catalog;
            Column = column;
            CreateTime = createTime;
            DataEngine = dataEngine;
            Database = database;
            Function = function;
            Id = id;
            Mode = mode;
            Operation = operation;
            Operator = @operator;
            PolicyType = policyType;
            ReAuth = reAuth;
            Source = source;
            SourceId = sourceId;
            SourceName = sourceName;
            Table = table;
            View = view;
        }
    }
}