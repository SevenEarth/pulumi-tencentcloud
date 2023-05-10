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

    public sealed class AccountPrivilegesAccountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// account host.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// account name.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public AccountPrivilegesAccountGetArgs()
        {
        }
    }
}
