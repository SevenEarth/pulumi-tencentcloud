// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver.Outputs
{

    [OutputType]
    public sealed class GetInsAttributeSslConfigResult
    {
        /// <summary>
        /// TDE encryption, 'enable' - enabled, 'disable' - not enabled.
        /// </summary>
        public readonly string Encryption;
        /// <summary>
        /// SSL certificate validity, 0-invalid, 1-valid Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int SslValidity;
        /// <summary>
        /// SSL certificate validity period, time format YYYY-MM-DD HH:MM:SS Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string SslValidityPeriod;

        [OutputConstructor]
        private GetInsAttributeSslConfigResult(
            string encryption,

            int sslValidity,

            string sslValidityPeriod)
        {
            Encryption = encryption;
            SslValidity = sslValidity;
            SslValidityPeriod = sslValidityPeriod;
        }
    }
}
