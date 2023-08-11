// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Inputs
{

    public sealed class EventTransformTransformationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describe how to filter data.
        /// </summary>
        [Input("etlFilter")]
        public Input<Inputs.EventTransformTransformationEtlFilterArgs>? EtlFilter { get; set; }

        /// <summary>
        /// Describe how to extract data.
        /// </summary>
        [Input("extraction")]
        public Input<Inputs.EventTransformTransformationExtractionArgs>? Extraction { get; set; }

        /// <summary>
        /// Describe how to convert data.
        /// </summary>
        [Input("transform")]
        public Input<Inputs.EventTransformTransformationTransformArgs>? Transform { get; set; }

        public EventTransformTransformationArgs()
        {
        }
    }
}
