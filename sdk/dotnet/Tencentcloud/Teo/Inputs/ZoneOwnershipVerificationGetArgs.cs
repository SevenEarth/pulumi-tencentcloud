// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class ZoneOwnershipVerificationGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsVerifications")]
        private InputList<Inputs.ZoneOwnershipVerificationDnsVerificationGetArgs>? _dnsVerifications;

        /// <summary>
        /// CNAME access, using DNS to resolve the information required for authentication. For details, please refer to [Site/Domain Name Ownership Verification ](https://cloud.tencent.com/document/product/1552/70789#7af6ecf8-afca-4e35-8811-b5797ed1bde5). Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.ZoneOwnershipVerificationDnsVerificationGetArgs> DnsVerifications
        {
            get => _dnsVerifications ?? (_dnsVerifications = new InputList<Inputs.ZoneOwnershipVerificationDnsVerificationGetArgs>());
            set => _dnsVerifications = value;
        }

        public ZoneOwnershipVerificationGetArgs()
        {
        }
    }
}