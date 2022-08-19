// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Inputs
{

    public sealed class PolicyBindingObjectDimensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'.
        /// </summary>
        [Input("dimensionsJson", required: true)]
        public Input<string> DimensionsJson { get; set; } = null!;

        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        public PolicyBindingObjectDimensionArgs()
        {
        }
    }
}
