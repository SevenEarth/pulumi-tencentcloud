// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Outputs
{

    [OutputType]
    public sealed class MigrationRenameRestore
    {
        /// <summary>
        /// When the new name of the library is used for offline migration, if it is not filled in, it will be named according to OldName. OldName and NewName cannot be filled in at the same time. OldName and NewName must be filled in and cannot be duplicate when used for cloning database.
        /// </summary>
        public readonly string? NewName;
        /// <summary>
        /// The name of the library. If oldName does not exist, a failure is returned.It can be left blank when used for offline migration tasks.
        /// </summary>
        public readonly string? OldName;

        [OutputConstructor]
        private MigrationRenameRestore(
            string? newName,

            string? oldName)
        {
            NewName = newName;
            OldName = oldName;
        }
    }
}
