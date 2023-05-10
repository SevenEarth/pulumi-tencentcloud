// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain.Inputs
{

    public sealed class TdsqlAuditLogFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of affected rows. Indicates filtering audit logs whose affected rows are greater than this value.
        /// </summary>
        [Input("affectRows")]
        public Input<int>? AffectRows { get; set; }

        [Input("dbNames")]
        private InputList<string>? _dbNames;

        /// <summary>
        /// Database name.
        /// </summary>
        public InputList<string> DbNames
        {
            get => _dbNames ?? (_dbNames = new InputList<string>());
            set => _dbNames = value;
        }

        /// <summary>
        /// Execution time. The unit is: us. It means to filter the audit logs whose execution time is greater than this value.
        /// </summary>
        [Input("execTime")]
        public Input<int>? ExecTime { get; set; }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// Client Address.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// Return the number of rows. It means to filter the audit log with the number of returned rows greater than this value.
        /// </summary>
        [Input("sentRows")]
        public Input<int>? SentRows { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// Username.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public TdsqlAuditLogFilterGetArgs()
        {
        }
    }
}
