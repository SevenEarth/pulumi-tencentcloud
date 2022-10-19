// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eks.Inputs
{

    public sealed class ContainerInstanceNfsVolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of NFS volume.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// NFS volume path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Indicates whether the volume is read only. Default is `false`.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// NFS server address.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        public ContainerInstanceNfsVolumeArgs()
        {
        }
    }
}