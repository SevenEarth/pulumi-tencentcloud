// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class NativeNodePoolLabelGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name in the map table.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value in map table.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NativeNodePoolLabelGetArgs()
        {
        }
        public static new NativeNodePoolLabelGetArgs Empty => new NativeNodePoolLabelGetArgs();
    }
}
