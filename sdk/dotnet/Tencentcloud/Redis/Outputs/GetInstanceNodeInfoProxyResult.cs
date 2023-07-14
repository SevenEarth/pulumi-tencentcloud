// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis.Outputs
{

    [OutputType]
    public sealed class GetInstanceNodeInfoProxyResult
    {
        /// <summary>
        /// Node ID.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// Zone ID.
        /// </summary>
        public readonly int ZoneId;

        [OutputConstructor]
        private GetInstanceNodeInfoProxyResult(
            string nodeId,

            int zoneId)
        {
            NodeId = nodeId;
            ZoneId = zoneId;
        }
    }
}
