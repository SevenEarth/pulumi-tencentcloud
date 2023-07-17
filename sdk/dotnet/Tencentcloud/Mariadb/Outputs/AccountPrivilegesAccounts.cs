// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class AccountPrivilegesAccounts
    {
        /// <summary>
        /// user host.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// user name.
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private AccountPrivilegesAccounts(
            string host,

            string user)
        {
            Host = host;
            User = user;
        }
    }
}
