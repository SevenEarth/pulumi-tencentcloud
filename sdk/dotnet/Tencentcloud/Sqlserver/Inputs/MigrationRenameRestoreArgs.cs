// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Inputs
{

    public sealed class MigrationRenameRestoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the new name of the library is used for offline migration, if it is not filled in, it will be named according to OldName. OldName and NewName cannot be filled in at the same time. OldName and NewName must be filled in and cannot be duplicate when used for cloning database.
        /// </summary>
        [Input("newName")]
        public Input<string>? NewName { get; set; }

        /// <summary>
        /// The name of the library. If oldName does not exist, a failure is returned.It can be left blank when used for offline migration tasks.
        /// </summary>
        [Input("oldName")]
        public Input<string>? OldName { get; set; }

        public MigrationRenameRestoreArgs()
        {
        }
    }
}