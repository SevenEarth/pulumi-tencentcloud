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

    public sealed class ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameterImageTemplateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Watermark height. % and px formats are supported: If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height; If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px.Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// Input content of watermark image. JPEG and PNG images are supported.
        /// </summary>
        [Input("imageContent", required: true)]
        public Input<Inputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameterImageTemplateImageContentGetArgs> ImageContent { get; set; } = null!;

        /// <summary>
        /// Repeat type of an animated watermark. Valid values: `once`: no longer appears after watermark playback ends. `repeat_last_frame`: stays on the last frame after watermark playback ends. `repeat` (default): repeats the playback until the video ends.
        /// </summary>
        [Input("repeatType")]
        public Input<string>? RepeatType { get; set; }

        /// <summary>
        /// Watermark width. % and px formats are supported: If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width; for example, `10%` means that `Width` is 10% of the video width; If the string ends in px, the `Width` of the watermark will be in px; for example, `100px` means that `Width` is 100 px.Default value: 10%.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameterImageTemplateGetArgs()
        {
        }
    }
}
