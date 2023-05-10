// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class PayCertificateDvAuth
    {
        /// <summary>
        /// DV authentication key.
        /// </summary>
        public readonly string? DvAuthKey;
        /// <summary>
        /// DV authentication value.
        /// </summary>
        public readonly string? DvAuthValue;
        /// <summary>
        /// DV authentication type.
        /// </summary>
        public readonly string? DvAuthVerifyType;

        [OutputConstructor]
        private PayCertificateDvAuth(
            string? dvAuthKey,

            string? dvAuthValue,

            string? dvAuthVerifyType)
        {
            DvAuthKey = dvAuthKey;
            DvAuthValue = dvAuthValue;
            DvAuthVerifyType = dvAuthVerifyType;
        }
    }
}
