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

    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video Transcoding Template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// End time offset of video after transcoding, unit: second.Do not fill in or fill in 0, indicating that the transcoded video continues until the end of the original video.When the value is greater than 0 (assumed to be n), it means that the transcoded video lasts until the nth second of the original video and terminates.When the value is less than 0 (assumed to be -n), it means that the transcoded video lasts until n seconds before the end of the original video.
        /// </summary>
        [Input("endTimeOffset")]
        public Input<double>? EndTimeOffset { get; set; }

        /// <summary>
        /// Opening and ending parameters.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("headTailParameter")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs>? HeadTailParameter { get; set; }

        [Input("mosaicSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs>? _mosaicSets;

        /// <summary>
        /// Mosaic list, up to 10 sheets can be supported.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs> MosaicSets
        {
            get => _mosaicSets ?? (_mosaicSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetMosaicSetGetArgs>());
            set => _mosaicSets = value;
        }

        /// <summary>
        /// Rules for the `{number}` variable in the output path after transcoding.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("objectNumberFormat")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetObjectNumberFormatGetArgs>? ObjectNumberFormat { get; set; }

        /// <summary>
        /// The output path of the main file after transcoding can be a relative path or an absolute path. If not filled, the default is a relative path: {inputName}_transcode_{definition}.{format}.
        /// </summary>
        [Input("outputObjectPath")]
        public Input<string>? OutputObjectPath { get; set; }

        /// <summary>
        /// The target storage of the transcoded file, if not filled, it will inherit the OutputStorage value of the upper layer.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("outputStorage")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOutputStorageGetArgs>? OutputStorage { get; set; }

        /// <summary>
        /// Video transcoding custom parameters, valid when Definition is not filled with 0.When some transcoding parameters in this structure are filled in, the parameters in the transcoding template will be overwritten with the filled parameters.This parameter is used in highly customized scenarios, it is recommended that you only use Definition to specify transcoding parameters.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("overrideParameter")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetOverrideParameterGetArgs>? OverrideParameter { get; set; }

        /// <summary>
        /// Video transcoding custom parameters, valid when Definition is filled with 0.This parameter is used in highly customized scenarios. It is recommended that you use Definition to specify transcoding parameters first.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("rawParameter")]
        public Input<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetRawParameterGetArgs>? RawParameter { get; set; }

        /// <summary>
        /// The output path of the transcoded fragment file (the path of ts when transcoding HLS), can only be a relative path. If not filled, the default is: `{inputName}_transcode_{definition}_{number}.{format}.
        /// </summary>
        [Input("segmentObjectName")]
        public Input<string>? SegmentObjectName { get; set; }

        /// <summary>
        /// The start time offset of the transcoded video, unit: second.Do not fill in or fill in 0, indicating that the transcoded video starts from the beginning of the original video.When the value is greater than 0 (assumed to be n), it means that the transcoded video starts from the nth second position of the original video.When the value is less than 0 (assumed to be -n), it means that the transcoded video starts from the position n seconds before the end of the original video.
        /// </summary>
        [Input("startTimeOffset")]
        public Input<double>? StartTimeOffset { get; set; }

        [Input("watermarkSets")]
        private InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs>? _watermarkSets;

        /// <summary>
        /// Watermark list, support multiple pictures or text watermarks, up to 10.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs> WatermarkSets
        {
            get => _watermarkSets ?? (_watermarkSets = new InputList<Inputs.WorkflowMediaProcessTaskTranscodeTaskSetWatermarkSetGetArgs>());
            set => _watermarkSets = value;
        }

        public WorkflowMediaProcessTaskTranscodeTaskSetGetArgs()
        {
        }
        public static new WorkflowMediaProcessTaskTranscodeTaskSetGetArgs Empty => new WorkflowMediaProcessTaskTranscodeTaskSetGetArgs();
    }
}
