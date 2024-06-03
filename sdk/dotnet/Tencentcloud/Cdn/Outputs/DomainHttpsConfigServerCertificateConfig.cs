// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainHttpsConfigServerCertificateConfig
    {
        /// <summary>
        /// Server certificate information. This is required when uploading an external certificate, which should contain the complete certificate chain.
        /// </summary>
        public readonly string? CertificateContent;
        /// <summary>
        /// Server certificate ID.
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// Server certificate name.
        /// </summary>
        public readonly string? CertificateName;
        /// <summary>
        /// Deploy time of server certificate.
        /// </summary>
        public readonly string? DeployTime;
        /// <summary>
        /// Signature expiration time in second. The maximum value is 630720000.
        /// </summary>
        public readonly string? ExpireTime;
        /// <summary>
        /// Certificate remarks.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Server key information. This is required when uploading an external certificate.
        /// </summary>
        public readonly string? PrivateKey;

        [OutputConstructor]
        private DomainHttpsConfigServerCertificateConfig(
            string? certificateContent,

            string? certificateId,

            string? certificateName,

            string? deployTime,

            string? expireTime,

            string? message,

            string? privateKey)
        {
            CertificateContent = certificateContent;
            CertificateId = certificateId;
            CertificateName = certificateName;
            DeployTime = deployTime;
            ExpireTime = expireTime;
            Message = message;
            PrivateKey = privateKey;
        }
    }
}
