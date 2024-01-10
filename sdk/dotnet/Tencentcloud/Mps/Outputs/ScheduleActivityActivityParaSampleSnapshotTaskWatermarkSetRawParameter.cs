// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameter
    {
        /// <summary>
        /// Origin position, which currently can only be: TopLeft: the origin of coordinates is in the top-left corner of the video, and the origin of the watermark is in the top-left corner of the image or text.Default value: TopLeft.
        /// </summary>
        public readonly string? CoordinateOrigin;
        /// <summary>
        /// Image watermark template. This field is required when `Type` is `image` and is invalid when `Type` is `text`.
        /// </summary>
        public readonly Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameterImageTemplate? ImageTemplate;
        /// <summary>
        /// Watermark type. Valid values: image: image watermark.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The horizontal position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported: If the string ends in %, the `XPos` of the watermark will be the specified percentage of the video width; for example, `10%` means that `XPos` is 10% of the video width; If the string ends in px, the `XPos` of the watermark will be the specified px; for example, `100px` means that `XPos` is 100 px.Default value: 0 px.
        /// </summary>
        public readonly string? XPos;
        /// <summary>
        /// The vertical position of the origin of the watermark relative to the origin of coordinates of the video. % and px formats are supported: If the string ends in %, the `YPos` of the watermark will be the specified percentage of the video height; for example, `10%` means that `YPos` is 10% of the video height; If the string ends in px, the `YPos` of the watermark will be the specified px; for example, `100px` means that `YPos` is 100 px.Default value: 0 px.
        /// </summary>
        public readonly string? YPos;

        [OutputConstructor]
        private ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameter(
            string? coordinateOrigin,

            Outputs.ScheduleActivityActivityParaSampleSnapshotTaskWatermarkSetRawParameterImageTemplate? imageTemplate,

            string type,

            string? xPos,

            string? yPos)
        {
            CoordinateOrigin = coordinateOrigin;
            ImageTemplate = imageTemplate;
            Type = type;
            XPos = xPos;
            YPos = yPos;
        }
    }
}
