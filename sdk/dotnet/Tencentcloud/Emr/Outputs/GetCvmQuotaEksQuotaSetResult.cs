// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Emr.Outputs
{

    [OutputType]
    public sealed class GetCvmQuotaEksQuotaSetResult
    {
        /// <summary>
        /// Cpu cores.
        /// </summary>
        public readonly int Cpu;
        /// <summary>
        /// Memory quantity (unit: GB).
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The specifications of the marketable resource are as follows: `TASK`, `CORE`, `MASTER`, `ROUTER`.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// Specifies the maximum number of resources that can be applied for.
        /// </summary>
        public readonly int Number;

        [OutputConstructor]
        private GetCvmQuotaEksQuotaSetResult(
            int cpu,

            int memory,

            string nodeType,

            int number)
        {
            Cpu = cpu;
            Memory = memory;
            NodeType = nodeType;
            Number = number;
        }
    }
}
