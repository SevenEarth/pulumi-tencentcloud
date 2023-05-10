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
    public sealed class GetSyncJobsListObjectDatabaseResult
    {
        /// <summary>
        /// database mode.
        /// </summary>
        public readonly string DbMode;
        /// <summary>
        /// database name.
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// function mode.
        /// </summary>
        public readonly string FunctionMode;
        /// <summary>
        /// functions.
        /// </summary>
        public readonly ImmutableArray<string> Functions;
        /// <summary>
        /// new database name.
        /// </summary>
        public readonly string NewDbName;
        /// <summary>
        /// new schema name.
        /// </summary>
        public readonly string NewSchemaName;
        /// <summary>
        /// procedure mode.
        /// </summary>
        public readonly string ProcedureMode;
        /// <summary>
        /// procedures.
        /// </summary>
        public readonly ImmutableArray<string> Procedures;
        /// <summary>
        /// schema name.
        /// </summary>
        public readonly string SchemaName;
        /// <summary>
        /// table mode.
        /// </summary>
        public readonly string TableMode;
        /// <summary>
        /// table list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSyncJobsListObjectDatabaseTableResult> Tables;
        /// <summary>
        /// view mode.
        /// </summary>
        public readonly string ViewMode;
        /// <summary>
        /// view list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSyncJobsListObjectDatabaseViewResult> Views;

        [OutputConstructor]
        private GetSyncJobsListObjectDatabaseResult(
            string dbMode,

            string dbName,

            string functionMode,

            ImmutableArray<string> functions,

            string newDbName,

            string newSchemaName,

            string procedureMode,

            ImmutableArray<string> procedures,

            string schemaName,

            string tableMode,

            ImmutableArray<Outputs.GetSyncJobsListObjectDatabaseTableResult> tables,

            string viewMode,

            ImmutableArray<Outputs.GetSyncJobsListObjectDatabaseViewResult> views)
        {
            DbMode = dbMode;
            DbName = dbName;
            FunctionMode = functionMode;
            Functions = functions;
            NewDbName = newDbName;
            NewSchemaName = newSchemaName;
            ProcedureMode = procedureMode;
            Procedures = procedures;
            SchemaName = schemaName;
            TableMode = tableMode;
            Tables = tables;
            ViewMode = viewMode;
            Views = views;
        }
    }
}