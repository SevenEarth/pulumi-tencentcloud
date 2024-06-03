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

    public sealed class DatahubTaskSourceResourceClickHouseParamSchemaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the column item is allowed to be empty.
        /// </summary>
        [Input("allowNull", required: true)]
        public Input<bool> AllowNull { get; set; } = null!;

        /// <summary>
        /// column name.
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        /// <summary>
        /// The json Key name corresponding to this column.
        /// </summary>
        [Input("jsonKey", required: true)]
        public Input<string> JsonKey { get; set; } = null!;

        /// <summary>
        /// type of table column.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatahubTaskSourceResourceClickHouseParamSchemaGetArgs()
        {
        }
        public static new DatahubTaskSourceResourceClickHouseParamSchemaGetArgs Empty => new DatahubTaskSourceResourceClickHouseParamSchemaGetArgs();
    }
}
