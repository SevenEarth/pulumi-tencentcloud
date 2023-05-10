// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb.Outputs
{

    [OutputType]
    public sealed class GetInstanceConnectionsClientResult
    {
        /// <summary>
        /// client connection count.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// is internal.
        /// </summary>
        public readonly bool InternalService;
        /// <summary>
        /// client connection ip.
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private GetInstanceConnectionsClientResult(
            int count,

            bool internalService,

            string ip)
        {
            Count = count;
            InternalService = internalService;
            Ip = ip;
        }
    }
}
