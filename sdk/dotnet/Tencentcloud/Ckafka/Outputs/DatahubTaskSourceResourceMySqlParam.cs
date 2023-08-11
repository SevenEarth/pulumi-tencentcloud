// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskSourceResourceMySqlParam
    {
        /// <summary>
        /// the name of the column to be monitored.
        /// </summary>
        public readonly string? DataSourceIncrementColumn;
        /// <summary>
        /// TIMESTAMP indicates that the incremental column is of timestamp type, INCREMENT indicates that the incremental column is of self-incrementing id type.
        /// </summary>
        public readonly string? DataSourceIncrementMode;
        /// <summary>
        /// TABLE indicates that the read item is a table, QUERY indicates that the read item is a query.
        /// </summary>
        public readonly string? DataSourceMonitorMode;
        /// <summary>
        /// When DataMonitorMode=TABLE, pass in the Table that needs to be read; when DataMonitorMode=QUERY, pass in the query sql statement that needs to be read.
        /// </summary>
        public readonly string? DataSourceMonitorResource;
        /// <summary>
        /// HEAD means copy stock + incremental data, TAIL means copy only incremental data.
        /// </summary>
        public readonly string? DataSourceStartFrom;
        /// <summary>
        /// INSERT means insert using Insert mode, UPSERT means insert using Upsert mode.
        /// </summary>
        public readonly string? DataTargetInsertMode;
        /// <summary>
        /// When DataInsertMode=UPSERT, pass in the primary key that the current upsert depends on.
        /// </summary>
        public readonly string? DataTargetPrimaryKeyField;
        /// <summary>
        /// Mapping relationship between tables and messages.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatahubTaskSourceResourceMySqlParamDataTargetRecordMapping> DataTargetRecordMappings;
        /// <summary>
        /// MySQL database name, * is the whole database.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// The Topic that stores the Ddl information of My SQL, if it is empty, it will not be stored by default.
        /// </summary>
        public readonly string? DdlTopic;
        /// <summary>
        /// When the member parameter Drop Invalid Message To Cls is set to true, the Drop Invalid Message parameter is invalid.
        /// </summary>
        public readonly Outputs.DatahubTaskSourceResourceMySqlParamDropCls? DropCls;
        /// <summary>
        /// Whether to discard messages that fail to parse, the default is true.
        /// </summary>
        public readonly bool? DropInvalidMessage;
        /// <summary>
        /// If the value is all, DDL data and DML data will also be written to the selected topic; if the value is dml, only DML data will be written to the selected topic.
        /// </summary>
        public readonly string? IncludeContentChanges;
        /// <summary>
        /// If the value is true, and the value of the binlog rows query log events configuration item in My SQL is ON, the data flowing into the topic contains the original SQL statement; if the value is false, the data flowing into the topic does not contain Original SQL statement.
        /// </summary>
        public readonly bool? IncludeQuery;
        /// <summary>
        /// When the Table input is a prefix, the value of this item is true, otherwise it is false.
        /// </summary>
        public readonly bool? IsTablePrefix;
        /// <summary>
        /// Whether the input table is a regular expression, if this option and Is Table Prefix are true at the same time, the judgment priority of this option is higher than Is Table Prefix.
        /// </summary>
        public readonly bool? IsTableRegular;
        /// <summary>
        /// Format library1.table1 field 1,field 2;library 2.table2 field 2, between tables; (semicolon) separated, between fields, (comma) separated. The table that is not specified defaults to the primary key of the table.
        /// </summary>
        public readonly string? KeyColumns;
        /// <summary>
        /// output format, DEFAULT, CANAL_1, CANAL_2.
        /// </summary>
        public readonly string? OutputFormat;
        /// <summary>
        /// If the value is true, the message will carry the schema corresponding to the message structure, if the value is false, it will not carry.
        /// </summary>
        public readonly bool? RecordWithSchema;
        /// <summary>
        /// MySQL connection Id.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// database name of signal table.
        /// </summary>
        public readonly string? SignalDatabase;
        /// <summary>
        /// whether to Copy inventory information (schema_only does not copy, initial full amount), the default is initial.
        /// </summary>
        public readonly string? SnapshotMode;
        /// <summary>
        /// The name of the MySQL data table,  is the non-system table in all the monitored databases, which can be separated by, to monitor multiple data tables, but the data table needs to be filled in the format of data database name.data table name, when a regular expression needs to be filled in, the format is data database name.data table name.
        /// </summary>
        public readonly string Table;
        /// <summary>
        /// Regular expression for routing events to specific topics, defaults to (.*).
        /// </summary>
        public readonly string? TopicRegex;
        /// <summary>
        /// TopicRegex, $1, $2.
        /// </summary>
        public readonly string? TopicReplacement;

        [OutputConstructor]
        private DatahubTaskSourceResourceMySqlParam(
            string? dataSourceIncrementColumn,

            string? dataSourceIncrementMode,

            string? dataSourceMonitorMode,

            string? dataSourceMonitorResource,

            string? dataSourceStartFrom,

            string? dataTargetInsertMode,

            string? dataTargetPrimaryKeyField,

            ImmutableArray<Outputs.DatahubTaskSourceResourceMySqlParamDataTargetRecordMapping> dataTargetRecordMappings,

            string database,

            string? ddlTopic,

            Outputs.DatahubTaskSourceResourceMySqlParamDropCls? dropCls,

            bool? dropInvalidMessage,

            string? includeContentChanges,

            bool? includeQuery,

            bool? isTablePrefix,

            bool? isTableRegular,

            string? keyColumns,

            string? outputFormat,

            bool? recordWithSchema,

            string resource,

            string? signalDatabase,

            string? snapshotMode,

            string table,

            string? topicRegex,

            string? topicReplacement)
        {
            DataSourceIncrementColumn = dataSourceIncrementColumn;
            DataSourceIncrementMode = dataSourceIncrementMode;
            DataSourceMonitorMode = dataSourceMonitorMode;
            DataSourceMonitorResource = dataSourceMonitorResource;
            DataSourceStartFrom = dataSourceStartFrom;
            DataTargetInsertMode = dataTargetInsertMode;
            DataTargetPrimaryKeyField = dataTargetPrimaryKeyField;
            DataTargetRecordMappings = dataTargetRecordMappings;
            Database = database;
            DdlTopic = ddlTopic;
            DropCls = dropCls;
            DropInvalidMessage = dropInvalidMessage;
            IncludeContentChanges = includeContentChanges;
            IncludeQuery = includeQuery;
            IsTablePrefix = isTablePrefix;
            IsTableRegular = isTableRegular;
            KeyColumns = keyColumns;
            OutputFormat = outputFormat;
            RecordWithSchema = recordWithSchema;
            Resource = resource;
            SignalDatabase = signalDatabase;
            SnapshotMode = snapshotMode;
            Table = table;
            TopicRegex = topicRegex;
            TopicReplacement = topicReplacement;
        }
    }
}