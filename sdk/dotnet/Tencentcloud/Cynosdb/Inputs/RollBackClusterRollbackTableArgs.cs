// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Inputs
{

    public sealed class RollBackClusterRollbackTableArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// New database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("tables", required: true)]
        private InputList<Inputs.RollBackClusterRollbackTableTableArgs>? _tables;

        /// <summary>
        /// Tables.
        /// </summary>
        public InputList<Inputs.RollBackClusterRollbackTableTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.RollBackClusterRollbackTableTableArgs>());
            set => _tables = value;
        }

        public RollBackClusterRollbackTableArgs()
        {
        }
        public static new RollBackClusterRollbackTableArgs Empty => new RollBackClusterRollbackTableArgs();
    }
}
