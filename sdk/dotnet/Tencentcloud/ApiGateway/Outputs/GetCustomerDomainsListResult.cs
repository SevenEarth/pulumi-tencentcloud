// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetCustomerDomainsListResult
    {
        /// <summary>
        /// The certificate ID.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// Domain name.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// Whether to use default path mapping. Valid values: `true`, `false`. `true` means to use default path mapping, `false` means to use custom path mapping.
        /// </summary>
        public readonly bool IsDefaultMapping;
        /// <summary>
        /// Domain name resolution status. Valid values: `true`, `false`. `true` means normal parsing, `false` means parsing failed.
        /// </summary>
        public readonly bool IsStatusOn;
        /// <summary>
        /// Network type.
        /// </summary>
        public readonly string NetType;
        /// <summary>
        /// Domain name mapping path and environment list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCustomerDomainsListPathMappingResult> PathMappings;
        /// <summary>
        /// Custom domain name agreement type.
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private GetCustomerDomainsListResult(
            string certificateId,

            string domainName,

            bool isDefaultMapping,

            bool isStatusOn,

            string netType,

            ImmutableArray<Outputs.GetCustomerDomainsListPathMappingResult> pathMappings,

            string protocol)
        {
            CertificateId = certificateId;
            DomainName = domainName;
            IsDefaultMapping = isDefaultMapping;
            IsStatusOn = isStatusOn;
            NetType = netType;
            PathMappings = pathMappings;
            Protocol = protocol;
        }
    }
}