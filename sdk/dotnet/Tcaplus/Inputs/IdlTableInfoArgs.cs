// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcaplus.Inputs
{

    public sealed class IdlTableInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Error messages for creating IDL file.
        /// </summary>
        [Input("error")]
        public Input<string>? Error { get; set; }

        /// <summary>
        /// Index key set of the TcaplusDB table.
        /// </summary>
        [Input("indexKeySet")]
        public Input<string>? IndexKeySet { get; set; }

        /// <summary>
        /// Primary key fields of the TcaplusDB table.
        /// </summary>
        [Input("keyFields")]
        public Input<string>? KeyFields { get; set; }

        /// <summary>
        /// Total size of primary key field of the TcaplusDB table.
        /// </summary>
        [Input("sumKeyFieldSize")]
        public Input<int>? SumKeyFieldSize { get; set; }

        /// <summary>
        /// Total size of non-primary key fields of the TcaplusDB table.
        /// </summary>
        [Input("sumValueFieldSize")]
        public Input<int>? SumValueFieldSize { get; set; }

        /// <summary>
        /// Name of the TcaplusDB table.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        /// <summary>
        /// Non-primary key fields of the TcaplusDB table.
        /// </summary>
        [Input("valueFields")]
        public Input<string>? ValueFields { get; set; }

        public IdlTableInfoArgs()
        {
        }
    }
}
