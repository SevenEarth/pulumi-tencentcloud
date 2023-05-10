// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketDomainCertificateAttachmentDomainCertificateCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate type.
        /// </summary>
        [Input("certType", required: true)]
        public Input<string> CertType { get; set; } = null!;

        /// <summary>
        /// Custom certificate.
        /// </summary>
        [Input("customCert", required: true)]
        public Input<Inputs.BucketDomainCertificateAttachmentDomainCertificateCertificateCustomCertArgs> CustomCert { get; set; } = null!;

        public BucketDomainCertificateAttachmentDomainCertificateCertificateArgs()
        {
        }
    }
}
