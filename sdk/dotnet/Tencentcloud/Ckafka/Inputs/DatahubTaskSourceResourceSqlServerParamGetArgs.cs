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

    public sealed class DatahubTaskSourceResourceSqlServerParamGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SQLServer database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// SQLServer connection Id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// schema_only|initial default initial.
        /// </summary>
        [Input("snapshotMode")]
        public Input<string>? SnapshotMode { get; set; }

        /// <summary>
        /// SQLServer table, *is the non-system table in all the monitored databases, you can use, to monitor multiple data tables, but the data table needs to be filled in the format of data database name.data table name.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public DatahubTaskSourceResourceSqlServerParamGetArgs()
        {
        }
        public static new DatahubTaskSourceResourceSqlServerParamGetArgs Empty => new DatahubTaskSourceResourceSqlServerParamGetArgs();
    }
}
