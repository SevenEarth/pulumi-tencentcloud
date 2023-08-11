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
    public sealed class RollBackClusterRollbackTableTable
    {
        /// <summary>
        /// New table name.
        /// </summary>
        public readonly string NewTable;
        /// <summary>
        /// Old table name.
        /// </summary>
        public readonly string OldTable;

        [OutputConstructor]
        private RollBackClusterRollbackTableTable(
            string newTable,

            string oldTable)
        {
            NewTable = newTable;
            OldTable = oldTable;
        }
    }
}