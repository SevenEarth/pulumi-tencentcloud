// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class FunctionCfsConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File system instance ID.
        /// </summary>
        [Input("cfsId", required: true)]
        public Input<string> CfsId { get; set; } = null!;

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Local mount directory.
        /// </summary>
        [Input("localMountDir", required: true)]
        public Input<string> LocalMountDir { get; set; } = null!;

        /// <summary>
        /// File system mount instance ID.
        /// </summary>
        [Input("mountInsId", required: true)]
        public Input<string> MountInsId { get; set; } = null!;

        [Input("mountSubnetId")]
        public Input<string>? MountSubnetId { get; set; }

        [Input("mountVpcId")]
        public Input<string>? MountVpcId { get; set; }

        /// <summary>
        /// Remote mount directory.
        /// </summary>
        [Input("remoteMountDir", required: true)]
        public Input<string> RemoteMountDir { get; set; } = null!;

        /// <summary>
        /// ID of user group.
        /// </summary>
        [Input("userGroupId", required: true)]
        public Input<string> UserGroupId { get; set; } = null!;

        /// <summary>
        /// ID of user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public FunctionCfsConfigArgs()
        {
        }
    }
}
