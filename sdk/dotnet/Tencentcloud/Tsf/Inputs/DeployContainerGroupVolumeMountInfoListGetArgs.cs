// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class DeployContainerGroupVolumeMountInfoListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Read and write access mode. 1: Read-only. 2: Read-write.
        /// </summary>
        [Input("readOrWrite")]
        public Input<string>? ReadOrWrite { get; set; }

        /// <summary>
        /// mount volume name.
        /// </summary>
        [Input("volumeMountName", required: true)]
        public Input<string> VolumeMountName { get; set; } = null!;

        /// <summary>
        /// mount path.
        /// </summary>
        [Input("volumeMountPath", required: true)]
        public Input<string> VolumeMountPath { get; set; } = null!;

        /// <summary>
        /// mount subPath.
        /// </summary>
        [Input("volumeMountSubPath")]
        public Input<string>? VolumeMountSubPath { get; set; }

        public DeployContainerGroupVolumeMountInfoListGetArgs()
        {
        }
        public static new DeployContainerGroupVolumeMountInfoListGetArgs Empty => new DeployContainerGroupVolumeMountInfoListGetArgs();
    }
}
