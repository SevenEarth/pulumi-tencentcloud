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

    public sealed class RestoreInstanceRenameRestoreGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// New database name. In offline migration, OldName will be used if NewName is left empty (OldName and NewName cannot be both empty). In database cloning, OldName and NewName must be both specified and cannot have the same value.
        /// </summary>
        [Input("newName", required: true)]
        public Input<string> NewName { get; set; } = null!;

        /// <summary>
        /// Database name. If the OldName database does not exist, a failure will be returned.It can be left empty in offline migration tasks.
        /// </summary>
        [Input("oldName", required: true)]
        public Input<string> OldName { get; set; } = null!;

        public RestoreInstanceRenameRestoreGetArgs()
        {
        }
        public static new RestoreInstanceRenameRestoreGetArgs Empty => new RestoreInstanceRenameRestoreGetArgs();
    }
}
