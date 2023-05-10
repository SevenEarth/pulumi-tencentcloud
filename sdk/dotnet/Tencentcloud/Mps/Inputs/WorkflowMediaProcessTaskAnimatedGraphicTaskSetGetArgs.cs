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

    public sealed class WorkflowMediaProcessTaskAnimatedGraphicTaskSetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video turntable template id.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// The end time of the animation in the video, in seconds.
        /// </summary>
        [Input("endTimeOffset", required: true)]
        public Input<double> EndTimeOffset { get; set; } = null!;

        /// <summary>
        /// The output path of the file after rotating the image, which can be a relative path or an absolute path. If not filled, the default is a relative path: {inputName}_animatedGraphic_{definition}.{format}.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// The target storage of the transcoded file, if not filled, it will inherit the OutputStorage value of the upper layer.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.WorkflowMediaProcessTaskAnimatedGraphicTaskSetOutputStorageGetArgs>? OutputStorage { get; set; }

        /// <summary>
        /// The start time of the animation in the video, in seconds.
        /// </summary>
        [Input("startTimeOffset", required: true)]
        public Input<double> StartTimeOffset { get; set; } = null!;

        public WorkflowMediaProcessTaskAnimatedGraphicTaskSetGetArgs()
        {
        }
    }
}
