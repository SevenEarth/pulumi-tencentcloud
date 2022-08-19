// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class ClusterRwGroupAddr
    {
        /// <summary>
        /// IP address for read-write connection.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Port of CynosDB cluster.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private ClusterRwGroupAddr(
            string? ip,

            int? port)
        {
            Ip = ip;
            Port = port;
        }
    }
}
