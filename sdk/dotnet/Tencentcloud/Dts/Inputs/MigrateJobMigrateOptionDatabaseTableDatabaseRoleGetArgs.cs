// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Inputs
{

    public sealed class MigrateJobMigrateOptionDatabaseTableDatabaseRoleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NewRoleName.
        /// </summary>
        [Input("newRoleName")]
        public Input<string>? NewRoleName { get; set; }

        /// <summary>
        /// RoleName.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        public MigrateJobMigrateOptionDatabaseTableDatabaseRoleGetArgs()
        {
        }
        public static new MigrateJobMigrateOptionDatabaseTableDatabaseRoleGetArgs Empty => new MigrateJobMigrateOptionDatabaseTableDatabaseRoleGetArgs();
    }
}
