// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Inputs
{

    public sealed class AccountPrivilegesColumnPrivilegeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Column name.
        /// </summary>
        [Input("column", required: true)]
        public Input<string> Column { get; set; } = null!;

        /// <summary>
        /// Database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("privileges", required: true)]
        private InputList<string>? _privileges;

        /// <summary>
        /// Permission information.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// Table name.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public AccountPrivilegesColumnPrivilegeGetArgs()
        {
        }
    }
}