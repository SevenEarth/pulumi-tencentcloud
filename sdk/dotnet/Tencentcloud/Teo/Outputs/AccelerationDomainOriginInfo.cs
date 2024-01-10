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
    public sealed class AccelerationDomainOriginInfo
    {
        /// <summary>
        /// ID of the secondary origin group (valid when `OriginType=ORIGIN_GROUP`). If it is not specified, it indicates that secondary origins are not used.
        /// </summary>
        public readonly string? BackupOrigin;
        /// <summary>
        /// The origin address. Enter the origin group ID if `OriginType=ORIGIN_GROUP`.
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// The origin type. Values: `IP_DOMAIN`: IPv4/IPv6 address or domain name; `COS`: COS bucket address; `ORIGIN_GROUP`: Origin group; `AWS_S3`: AWS S3 bucket address; `SPACE`: EdgeOne Shield Space.
        /// </summary>
        public readonly string OriginType;
        /// <summary>
        /// Whether to authenticate access to the private object storage origin (valid when `OriginType=COS/AWS_S3`). Values: `on`: Enable private authentication; `off`: Disable private authentication. If this field is not specified, the default value `off` is used.
        /// </summary>
        public readonly string? PrivateAccess;
        /// <summary>
        /// The private authentication parameters. This field is valid when `PrivateAccess=on`.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccelerationDomainOriginInfoPrivateParameter> PrivateParameters;

        [OutputConstructor]
        private AccelerationDomainOriginInfo(
            string? backupOrigin,

            string origin,

            string originType,

            string? privateAccess,

            ImmutableArray<Outputs.AccelerationDomainOriginInfoPrivateParameter> privateParameters)
        {
            BackupOrigin = backupOrigin;
            Origin = origin;
            OriginType = originType;
            PrivateAccess = privateAccess;
            PrivateParameters = privateParameters;
        }
    }
}
