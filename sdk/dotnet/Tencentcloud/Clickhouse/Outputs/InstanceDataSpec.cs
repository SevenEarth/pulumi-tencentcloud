// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse.Outputs
{

    [OutputType]
    public sealed class InstanceDataSpec
    {
        /// <summary>
        /// Data spec count.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// Disk size.
        /// </summary>
        public readonly int DiskSize;
        /// <summary>
        /// Spec name.
        /// </summary>
        public readonly string SpecName;

        [OutputConstructor]
        private InstanceDataSpec(
            int count,

            int diskSize,

            string specName)
        {
            Count = count;
            DiskSize = diskSize;
            SpecName = specName;
        }
    }
}