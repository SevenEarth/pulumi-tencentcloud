// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetSgSnapshotFileContentOriginalDataAddressTemplateResult
    {
        /// <summary>
        /// The ID of the IP address group, such as `ipmg-2uw6ujo6`.
        /// </summary>
        public readonly string AddressGroupId;
        /// <summary>
        /// The ID of the IP address, such as `ipm-2uw6ujo6`.
        /// </summary>
        public readonly string AddressId;

        [OutputConstructor]
        private GetSgSnapshotFileContentOriginalDataAddressTemplateResult(
            string addressGroupId,

            string addressId)
        {
            AddressGroupId = addressGroupId;
            AddressId = addressId;
        }
    }
}
