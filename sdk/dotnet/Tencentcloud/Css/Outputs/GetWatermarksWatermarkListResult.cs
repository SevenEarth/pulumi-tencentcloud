// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css.Outputs
{

    [OutputType]
    public sealed class GetWatermarksWatermarkListResult
    {
        /// <summary>
        /// The time when the watermark was added.Note: Beijing time (UTC+8) is used.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Watermark height.
        /// </summary>
        public readonly int Height;
        /// <summary>
        /// Watermark image URL.
        /// </summary>
        public readonly string PictureUrl;
        /// <summary>
        /// Current status. 0: not used. 1: in use.
        /// </summary>
        public readonly int Status;
        /// <summary>
        /// Watermark ID.
        /// </summary>
        public readonly int WatermarkId;
        /// <summary>
        /// Watermark name.
        /// </summary>
        public readonly string WatermarkName;
        /// <summary>
        /// Watermark width.
        /// </summary>
        public readonly int Width;
        /// <summary>
        /// Display position: X-axis offset.
        /// </summary>
        public readonly int XPosition;
        /// <summary>
        /// Display position: Y-axis offset.
        /// </summary>
        public readonly int YPosition;

        [OutputConstructor]
        private GetWatermarksWatermarkListResult(
            string createTime,

            int height,

            string pictureUrl,

            int status,

            int watermarkId,

            string watermarkName,

            int width,

            int xPosition,

            int yPosition)
        {
            CreateTime = createTime;
            Height = height;
            PictureUrl = pictureUrl;
            Status = status;
            WatermarkId = watermarkId;
            WatermarkName = watermarkName;
            Width = width;
            XPosition = xPosition;
            YPosition = yPosition;
        }
    }
}
