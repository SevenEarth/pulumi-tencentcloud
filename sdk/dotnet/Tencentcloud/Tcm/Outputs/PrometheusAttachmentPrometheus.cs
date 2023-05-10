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
    public sealed class PrometheusAttachmentPrometheus
    {
        /// <summary>
        /// Third party prometheus.
        /// </summary>
        public readonly Outputs.PrometheusAttachmentPrometheusCustomProm? CustomProm;
        /// <summary>
        /// Existed TMP id, auto create TMP if empty.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// Region for TMP.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// Subnet id for TMP.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Vpc id for TMP.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private PrometheusAttachmentPrometheus(
            Outputs.PrometheusAttachmentPrometheusCustomProm? customProm,

            string? instanceId,

            string? region,

            string? subnetId,

            string? vpcId)
        {
            CustomProm = customProm;
            InstanceId = instanceId;
            Region = region;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
