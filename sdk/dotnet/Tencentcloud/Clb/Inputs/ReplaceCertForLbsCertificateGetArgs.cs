// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb.Inputs
{

    public sealed class ReplaceCertForLbsCertificateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Content of the uploaded client certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
        /// </summary>
        [Input("certCaContent")]
        public Input<string>? CertCaContent { get; set; }

        /// <summary>
        /// ID of a client certificate. When the listener adopts mutual authentication (i.e., SSLMode = mutual), if you leave this parameter empty, you must upload the client certificate, including CertCaContent and CertCaName.
        /// </summary>
        [Input("certCaId")]
        public Input<string>? CertCaId { get; set; }

        /// <summary>
        /// Name of the uploaded client CA certificate. When SSLMode = mutual, if there is no CertCaId, this parameter is required.
        /// </summary>
        [Input("certCaName")]
        public Input<string>? CertCaName { get; set; }

        /// <summary>
        /// Content of the uploaded server certificate. If there is no CertId, this parameter is required.
        /// </summary>
        [Input("certContent")]
        public Input<string>? CertContent { get; set; }

        /// <summary>
        /// ID of a server certificate. If you leave this parameter empty, you must upload the certificate, including CertContent, CertKey, and CertName.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// Key of the uploaded server certificate. If there is no CertId, this parameter is required.
        /// </summary>
        [Input("certKey")]
        public Input<string>? CertKey { get; set; }

        /// <summary>
        /// Name of the uploaded server certificate. If there is no CertId, this parameter is required.
        /// </summary>
        [Input("certName")]
        public Input<string>? CertName { get; set; }

        /// <summary>
        /// Authentication type. Value range: UNIDIRECTIONAL (unidirectional authentication), MUTUAL (mutual authentication).
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        public ReplaceCertForLbsCertificateGetArgs()
        {
        }
    }
}
