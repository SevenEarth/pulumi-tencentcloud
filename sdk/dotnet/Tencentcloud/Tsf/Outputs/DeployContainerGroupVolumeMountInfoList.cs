// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class DeployContainerGroupVolumeMountInfoList
    {
        /// <summary>
        /// Read and write access mode. 1: Read-only. 2: Read-write.
        /// </summary>
        public readonly string? ReadOrWrite;
        /// <summary>
        /// mount volume name.
        /// </summary>
        public readonly string VolumeMountName;
        /// <summary>
        /// mount path.
        /// </summary>
        public readonly string VolumeMountPath;
        /// <summary>
        /// mount subPath.
        /// </summary>
        public readonly string? VolumeMountSubPath;

        [OutputConstructor]
        private DeployContainerGroupVolumeMountInfoList(
            string? readOrWrite,

            string volumeMountName,

            string volumeMountPath,

            string? volumeMountSubPath)
        {
            ReadOrWrite = readOrWrite;
            VolumeMountName = volumeMountName;
            VolumeMountPath = volumeMountPath;
            VolumeMountSubPath = volumeMountSubPath;
        }
    }
}
