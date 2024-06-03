// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Inputs
{

    public sealed class AccountPrivilegesViewPrivilegesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of database.
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
        /// Database view name.
        /// </summary>
        [Input("view", required: true)]
        public Input<string> View { get; set; } = null!;

        public AccountPrivilegesViewPrivilegesArgs()
        {
        }
        public static new AccountPrivilegesViewPrivilegesArgs Empty => new AccountPrivilegesViewPrivilegesArgs();
    }
}
