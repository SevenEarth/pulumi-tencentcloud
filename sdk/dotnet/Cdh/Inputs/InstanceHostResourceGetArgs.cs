// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdh.Inputs
{

    public sealed class InstanceHostResourceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of available CPU cores of the instance.
        /// </summary>
        [Input("cpuAvailableNum")]
        public Input<int>? CpuAvailableNum { get; set; }

        /// <summary>
        /// The number of total CPU cores of the instance.
        /// </summary>
        [Input("cpuTotalNum")]
        public Input<int>? CpuTotalNum { get; set; }

        /// <summary>
        /// Instance disk available capacity, unit in GB.
        /// </summary>
        [Input("diskAvailableSize")]
        public Input<int>? DiskAvailableSize { get; set; }

        /// <summary>
        /// Instance disk total capacity, unit in GB.
        /// </summary>
        [Input("diskTotalSize")]
        public Input<int>? DiskTotalSize { get; set; }

        /// <summary>
        /// Type of the disk.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Instance memory available capacity, unit in GB.
        /// </summary>
        [Input("memoryAvailableSize")]
        public Input<double>? MemoryAvailableSize { get; set; }

        /// <summary>
        /// Instance memory total capacity, unit in GB.
        /// </summary>
        [Input("memoryTotalSize")]
        public Input<double>? MemoryTotalSize { get; set; }

        public InstanceHostResourceGetArgs()
        {
        }
    }
}
