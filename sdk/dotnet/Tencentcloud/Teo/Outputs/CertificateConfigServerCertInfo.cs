// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class CertificateConfigServerCertInfo
    {
        /// <summary>
        /// Alias of the certificate.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// ID of the server certificate.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string CertId;
        /// <summary>
        /// Domain name of the certificate. Note: This field may return `null`, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? CommonName;
        /// <summary>
        /// Time when the certificate is deployed. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? DeployTime;
        /// <summary>
        /// Time when the certificate expires. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? ExpireTime;
        /// <summary>
        /// Signature algorithm. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string? SignAlgo;
        /// <summary>
        /// Type of the certificate. Values: `default`: Default certificate; `upload`: Specified certificate; `managed`: Tencent Cloud-managed certificate. Note: This field may return `null`, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private CertificateConfigServerCertInfo(
            string? alias,

            string certId,

            string? commonName,

            string? deployTime,

            string? expireTime,

            string? signAlgo,

            string? type)
        {
            Alias = alias;
            CertId = certId;
            CommonName = commonName;
            DeployTime = deployTime;
            ExpireTime = expireTime;
            SignAlgo = signAlgo;
            Type = type;
        }
    }
}