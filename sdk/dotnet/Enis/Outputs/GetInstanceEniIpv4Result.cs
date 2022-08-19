// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Enis.Outputs
{

    [OutputType]
    public sealed class GetInstanceEniIpv4Result
    {
        /// <summary>
        /// Description of the ENI. Conflict with `ids`.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Intranet IP.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Indicates whether the IP is primary.
        /// </summary>
        public readonly bool Primary;

        [OutputConstructor]
        private GetInstanceEniIpv4Result(
            string description,

            string ip,

            bool primary)
        {
            Description = description;
            Ip = ip;
            Primary = primary;
        }
    }
}
