// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcaplus.Outputs
{

    [OutputType]
    public sealed class IdlTableInfo
    {
        /// <summary>
        /// Error messages for creating IDL file.
        /// </summary>
        public readonly string? Error;
        /// <summary>
        /// Index key set of the TcaplusDB table.
        /// </summary>
        public readonly string? IndexKeySet;
        /// <summary>
        /// Primary key fields of the TcaplusDB table.
        /// </summary>
        public readonly string? KeyFields;
        /// <summary>
        /// Total size of primary key field of the TcaplusDB table.
        /// </summary>
        public readonly int? SumKeyFieldSize;
        /// <summary>
        /// Total size of non-primary key fields of the TcaplusDB table.
        /// </summary>
        public readonly int? SumValueFieldSize;
        /// <summary>
        /// Name of the TcaplusDB table.
        /// </summary>
        public readonly string? TableName;
        /// <summary>
        /// Non-primary key fields of the TcaplusDB table.
        /// </summary>
        public readonly string? ValueFields;

        [OutputConstructor]
        private IdlTableInfo(
            string? error,

            string? indexKeySet,

            string? keyFields,

            int? sumKeyFieldSize,

            int? sumValueFieldSize,

            string? tableName,

            string? valueFields)
        {
            Error = error;
            IndexKeySet = indexKeySet;
            KeyFields = keyFields;
            SumKeyFieldSize = sumKeyFieldSize;
            SumValueFieldSize = sumValueFieldSize;
            TableName = tableName;
            ValueFields = valueFields;
        }
    }
}
