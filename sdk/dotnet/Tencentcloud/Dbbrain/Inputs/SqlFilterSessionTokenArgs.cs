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

    public sealed class SqlFilterSessionTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// password.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// user name.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public SqlFilterSessionTokenArgs()
        {
        }
        public static new SqlFilterSessionTokenArgs Empty => new SqlFilterSessionTokenArgs();
    }
}
