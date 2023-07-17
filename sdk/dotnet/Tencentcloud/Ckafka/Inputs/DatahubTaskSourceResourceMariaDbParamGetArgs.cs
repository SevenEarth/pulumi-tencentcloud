// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskSourceResourceMariaDbParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// MariaDB database name, * for all database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// If the value is all, DDL data and DML data will also be written to the selected topic; if the value is dml, only DML data will be written to the selected topic.
        /// </summary>
        [Input("includeContentChanges")]
        public Input<string>? IncludeContentChanges { get; set; }

        /// <summary>
        /// If the value is true, and the value of the binlog rows query log events configuration item in My SQL is ON, the data flowing into the topic contains the original SQL statement; if the value is false, the data flowing into the topic does not contain Original SQL statement.
        /// </summary>
        [Input("includeQuery")]
        public Input<bool>? IncludeQuery { get; set; }

        /// <summary>
        /// When the Table input is a prefix, the value of this item is true, otherwise it is false.
        /// </summary>
        [Input("isTablePrefix")]
        public Input<bool>? IsTablePrefix { get; set; }

        /// <summary>
        /// Format  library 1. table 1: field 1, field 2; library 2. table 2: field 2, between tables; (semicolon) separated, between fields, (comma) separated. The table that is not specified defaults to the primary key of the table.
        /// </summary>
        [Input("keyColumns")]
        public Input<string>? KeyColumns { get; set; }

        /// <summary>
        /// output format, DEFAULT, CANAL_1, CANAL_2.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        /// <summary>
        /// If the value is true, the message will carry the schema corresponding to the message structure, if the value is false, it will not carry.
        /// </summary>
        [Input("recordWithSchema")]
        public Input<bool>? RecordWithSchema { get; set; }

        /// <summary>
        /// MariaDB connection Id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// schema_only|initial, default initial.
        /// </summary>
        [Input("snapshotMode")]
        public Input<string>? SnapshotMode { get; set; }

        /// <summary>
        /// MariaDB db name, *is the non-system table in all the monitored databases, you can use, to monitor multiple data tables, but the data table needs to be filled in the format of data database name.data table name.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public DatahubTaskSourceResourceMariaDbParamGetArgs()
        {
        }
    }
}
