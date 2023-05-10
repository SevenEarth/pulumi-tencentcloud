// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class GetInstancesModificationInstanceTypeConfigStatusListInstanceTypeConfigResult
    {
        /// <summary>
        /// The number of CPU kernels, in cores.
        /// </summary>
        public readonly int Cpu;
        /// <summary>
        /// The number of FPGA kernels, in cores.
        /// </summary>
        public readonly int Fpga;
        /// <summary>
        /// The number of GPU kernels, in cores.
        /// </summary>
        public readonly int Gpu;
        /// <summary>
        /// Instance family.
        /// </summary>
        public readonly string InstanceFamily;
        /// <summary>
        /// Instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Memory capacity (in GB).
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// Availability zone.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetInstancesModificationInstanceTypeConfigStatusListInstanceTypeConfigResult(
            int cpu,

            int fpga,

            int gpu,

            string instanceFamily,

            string instanceType,

            int memory,

            string zone)
        {
            Cpu = cpu;
            Fpga = fpga;
            Gpu = gpu;
            InstanceFamily = instanceFamily;
            InstanceType = instanceType;
            Memory = memory;
            Zone = zone;
        }
    }
}