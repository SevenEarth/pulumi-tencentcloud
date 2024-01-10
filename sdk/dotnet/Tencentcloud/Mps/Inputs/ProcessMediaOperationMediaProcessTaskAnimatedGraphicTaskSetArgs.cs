// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Animated image generating template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// End time of an animated image in a video in seconds.
        /// </summary>
        [Input("endTimeOffset", required: true)]
        public Input<double> EndTimeOffset { get; set; } = null!;

        /// <summary>
        /// Output path to a generated animated image file, which can be a relative path or an absolute path. If this parameter is left empty, the following relative path will be used by default: `{inputName}_animatedGraphic_{definition}.{format}`.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// Target bucket of a generated animated image file. If this parameter is left empty, the `OutputStorage` value of the upper folder will be inherited.Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetOutputStorageArgs>? OutputStorage { get; set; }

        /// <summary>
        /// Start time of an animated image in a video in seconds.
        /// </summary>
        [Input("startTimeOffset", required: true)]
        public Input<double> StartTimeOffset { get; set; } = null!;

        public ProcessMediaOperationMediaProcessTaskAnimatedGraphicTaskSetArgs()
        {
        }
    }
}
