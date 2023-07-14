// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigSidecarResourcesLimitGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource type name, `cpu/memory`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource quantity, example: cpu-`100m`, memory-`1Gi`.
        /// </summary>
        [Input("quantity")]
        public Input<string>? Quantity { get; set; }

        public MeshConfigSidecarResourcesLimitGetArgs()
        {
        }
    }
}
