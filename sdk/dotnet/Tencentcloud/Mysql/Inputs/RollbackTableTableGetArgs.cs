// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Inputs
{

    public sealed class RollbackTableTableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// New database table name after rollback.
        /// </summary>
        [Input("newTableName", required: true)]
        public Input<string> NewTableName { get; set; } = null!;

        /// <summary>
        /// The original database table name before rollback.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public RollbackTableTableGetArgs()
        {
        }
        public static new RollbackTableTableGetArgs Empty => new RollbackTableTableGetArgs();
    }
}
