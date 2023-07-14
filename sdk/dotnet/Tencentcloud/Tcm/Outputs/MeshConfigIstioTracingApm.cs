// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Outputs
{

    [OutputType]
    public sealed class MeshConfigIstioTracingApm
    {
        /// <summary>
        /// Whether enable APM.
        /// </summary>
        public readonly bool Enable;
        /// <summary>
        /// Instance id of the APM.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// Region.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private MeshConfigIstioTracingApm(
            bool enable,

            string? instanceId,

            string? region)
        {
            Enable = enable;
            InstanceId = instanceId;
            Region = region;
        }
    }
}
