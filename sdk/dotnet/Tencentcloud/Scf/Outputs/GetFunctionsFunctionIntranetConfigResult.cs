// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetFunctionsFunctionIntranetConfigResult
    {
        /// <summary>
        /// If fixed intranet IP is enabled, this field returns the IP list used.
        /// </summary>
        public readonly ImmutableArray<string> IpAddresses;
        /// <summary>
        /// Whether to enable fixed intranet IP, ENABLE is enabled, DISABLE is disabled.
        /// </summary>
        public readonly string IpFixed;

        [OutputConstructor]
        private GetFunctionsFunctionIntranetConfigResult(
            ImmutableArray<string> ipAddresses,

            string ipFixed)
        {
            IpAddresses = ipAddresses;
            IpFixed = ipFixed;
        }
    }
}
