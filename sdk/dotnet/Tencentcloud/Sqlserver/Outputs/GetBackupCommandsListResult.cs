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
    public sealed class GetBackupCommandsListResult
    {
        /// <summary>
        /// Create backup command.
        /// </summary>
        public readonly string Command;
        /// <summary>
        /// Request ID.
        /// </summary>
        public readonly string RequestId;

        [OutputConstructor]
        private GetBackupCommandsListResult(
            string command,

            string requestId)
        {
            Command = command;
            RequestId = requestId;
        }
    }
}
