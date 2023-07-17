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
    public sealed class GetInstanceSlowQueriesSlowQueryResult
    {
        /// <summary>
        /// Database name.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Lock duration in seconds.
        /// </summary>
        public readonly double LockTime;
        /// <summary>
        /// Execution time in seconds.
        /// </summary>
        public readonly double QueryTime;
        /// <summary>
        /// Scan Rows.
        /// </summary>
        public readonly int RowsExamined;
        /// <summary>
        /// Return the number of rows.
        /// </summary>
        public readonly int RowsSent;
        /// <summary>
        /// SQL statement md5.
        /// </summary>
        public readonly string SqlMd5;
        /// <summary>
        /// SQL template.
        /// </summary>
        public readonly string SqlTemplate;
        /// <summary>
        /// SQL statement.
        /// </summary>
        public readonly string SqlText;
        /// <summary>
        /// Execution timestamp.
        /// </summary>
        public readonly int Timestamp;
        /// <summary>
        /// Client host.
        /// </summary>
        public readonly string UserHost;
        /// <summary>
        /// user name.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetInstanceSlowQueriesSlowQueryResult(
            string database,

            double lockTime,

            double queryTime,

            int rowsExamined,

            int rowsSent,

            string sqlMd5,

            string sqlTemplate,

            string sqlText,

            int timestamp,

            string userHost,

            string userName)
        {
            Database = database;
            LockTime = lockTime;
            QueryTime = queryTime;
            RowsExamined = rowsExamined;
            RowsSent = rowsSent;
            SqlMd5 = sqlMd5;
            SqlTemplate = sqlTemplate;
            SqlText = sqlText;
            Timestamp = timestamp;
            UserHost = userHost;
            UserName = userName;
        }
    }
}
