// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap.Outputs
{

    [OutputType]
    public sealed class GetHttpDomainsDomainResult
    {
        /// <summary>
        /// Indicates whether basic authentication is enable.
        /// </summary>
        public readonly bool BasicAuth;
        /// <summary>
        /// ID of the basic authentication.
        /// </summary>
        public readonly string BasicAuthId;
        /// <summary>
        /// ID of the server certificate.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// (**Deprecated**) It has been deprecated from version 1.26.0. Use `client_certificate_ids` instead. ID of the client certificate.
        /// </summary>
        public readonly string ClientCertificateId;
        /// <summary>
        /// ID list of the client certificate.
        /// </summary>
        public readonly ImmutableArray<string> ClientCertificateIds;
        /// <summary>
        /// Forward domain of the layer7 listener to be queried.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Indicates whether SSL certificate authentication is enable.
        /// </summary>
        public readonly bool GaapAuth;
        /// <summary>
        /// ID of the SSL certificate.
        /// </summary>
        public readonly string GaapAuthId;
        /// <summary>
        /// Indicates whether realserver authentication is enable.
        /// </summary>
        public readonly bool RealserverAuth;
        /// <summary>
        /// CA certificate domain of the realserver.
        /// </summary>
        public readonly string RealserverCertificateDomain;
        /// <summary>
        /// (**Deprecated**) It has been deprecated from version 1.28.0. Use `realserver_certificate_ids` instead. CA certificate ID of the realserver.
        /// </summary>
        public readonly string RealserverCertificateId;
        /// <summary>
        /// CA certificate ID list of the realserver.
        /// </summary>
        public readonly ImmutableArray<string> RealserverCertificateIds;

        [OutputConstructor]
        private GetHttpDomainsDomainResult(
            bool basicAuth,

            string basicAuthId,

            string certificateId,

            string clientCertificateId,

            ImmutableArray<string> clientCertificateIds,

            string domain,

            bool gaapAuth,

            string gaapAuthId,

            bool realserverAuth,

            string realserverCertificateDomain,

            string realserverCertificateId,

            ImmutableArray<string> realserverCertificateIds)
        {
            BasicAuth = basicAuth;
            BasicAuthId = basicAuthId;
            CertificateId = certificateId;
            ClientCertificateId = clientCertificateId;
            ClientCertificateIds = clientCertificateIds;
            Domain = domain;
            GaapAuth = gaapAuth;
            GaapAuthId = gaapAuthId;
            RealserverAuth = realserverAuth;
            RealserverCertificateDomain = realserverCertificateDomain;
            RealserverCertificateId = realserverCertificateId;
            RealserverCertificateIds = realserverCertificateIds;
        }
    }
}
