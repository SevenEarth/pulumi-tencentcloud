// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ha.Outputs
{

    [OutputType]
    public sealed class GetVipEipAttachmentsHaVipEipAttachmentListResult
    {
        /// <summary>
        /// Public IP address of EIP to be queried.
        /// </summary>
        public readonly string AddressIp;
        /// <summary>
        /// ID of the attached HA VIP to be queried.
        /// </summary>
        public readonly string HavipId;

        [OutputConstructor]
        private GetVipEipAttachmentsHaVipEipAttachmentListResult(
            string addressIp,

            string havipId)
        {
            AddressIp = addressIp;
            HavipId = havipId;
        }
    }
}
