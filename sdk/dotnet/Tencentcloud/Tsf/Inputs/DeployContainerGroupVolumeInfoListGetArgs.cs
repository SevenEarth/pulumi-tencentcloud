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

    public sealed class DeployContainerGroupVolumeInfoListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// volume config.
        /// </summary>
        [Input("volumeConfig")]
        public Input<string>? VolumeConfig { get; set; }

        /// <summary>
        /// volume name.
        /// </summary>
        [Input("volumeName", required: true)]
        public Input<string> VolumeName { get; set; } = null!;

        /// <summary>
        /// volume type.
        /// </summary>
        [Input("volumeType", required: true)]
        public Input<string> VolumeType { get; set; } = null!;

        public DeployContainerGroupVolumeInfoListGetArgs()
        {
        }
        public static new DeployContainerGroupVolumeInfoListGetArgs Empty => new DeployContainerGroupVolumeInfoListGetArgs();
    }
}
