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

    public sealed class EventTransformTransformationExtractionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// JsonPath, if not specified, the default value $.
        /// </summary>
        [Input("extractionInputPath", required: true)]
        public Input<string> ExtractionInputPath { get; set; } = null!;

        /// <summary>
        /// Value: `TEXT`, `JSON`.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// Only Text needs to be passed.
        /// </summary>
        [Input("textParams")]
        public Input<Inputs.EventTransformTransformationExtractionTextParamsGetArgs>? TextParams { get; set; }

        public EventTransformTransformationExtractionGetArgs()
        {
        }
    }
}
