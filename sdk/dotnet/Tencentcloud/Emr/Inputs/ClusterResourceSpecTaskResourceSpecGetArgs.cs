// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Emr.Inputs
{

    public sealed class ClusterResourceSpecTaskResourceSpecGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        [Input("rootSize")]
        public Input<int>? RootSize { get; set; }

        [Input("spec")]
        public Input<string>? Spec { get; set; }

        [Input("storageType")]
        public Input<int>? StorageType { get; set; }

        public ClusterResourceSpecTaskResourceSpecGetArgs()
        {
        }
        public static new ClusterResourceSpecTaskResourceSpecGetArgs Empty => new ClusterResourceSpecTaskResourceSpecGetArgs();
    }
}
