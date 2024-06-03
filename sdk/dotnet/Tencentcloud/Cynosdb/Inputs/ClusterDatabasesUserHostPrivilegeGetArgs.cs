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

    public sealed class ClusterDatabasesUserHostPrivilegeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// .
        /// </summary>
        [Input("dbHost", required: true)]
        public Input<string> DbHost { get; set; } = null!;

        /// <summary>
        /// .
        /// </summary>
        [Input("dbPrivilege", required: true)]
        public Input<string> DbPrivilege { get; set; } = null!;

        /// <summary>
        /// Authorized Users.
        /// </summary>
        [Input("dbUserName", required: true)]
        public Input<string> DbUserName { get; set; } = null!;

        public ClusterDatabasesUserHostPrivilegeGetArgs()
        {
        }
        public static new ClusterDatabasesUserHostPrivilegeGetArgs Empty => new ClusterDatabasesUserHostPrivilegeGetArgs();
    }
}
