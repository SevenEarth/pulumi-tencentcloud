// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clb.Outputs
{

    [OutputType]
    public sealed class GetAttachmentsAttachmentListTargetResult
    {
        /// <summary>
        /// Id of the backend server.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Port of the backend server.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Forwarding weight of the backend service, the range of [0, 100], defaults to `10`.
        /// </summary>
        public readonly int Weight;

        [OutputConstructor]
        private GetAttachmentsAttachmentListTargetResult(
            string instanceId,

            int port,

            int weight)
        {
            InstanceId = instanceId;
            Port = port;
            Weight = weight;
        }
    }
}
