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
    public sealed class WatermarkTemplateImageTemplate
    {
        /// <summary>
        /// The height of the watermark. Support %, px two formats:When the string ends with %, it means that the watermark Height is the percentage size of the video height, such as 10% means that the Height is 10% of the video height.When the string ends with px, it means that the watermark Height unit is pixel, such as 100px means that the Height is 100 pixels. The value range is 0 or [8, 4096].Default value: 0px. Indicates that Height is scaled according to the aspect ratio of the original watermark image.
        /// </summary>
        public readonly string? Height;
        /// <summary>
        /// Watermark image[Base64](https://tools.ietf.org/html/rfc4648) encoded string. Support jpeg, png image format.
        /// </summary>
        public readonly string ImageContent;
        /// <summary>
        /// Watermark repeat type. Usage scenario: The watermark is a dynamic image. Ranges:once: After the dynamic watermark is played, it will no longer appear.repeat_last_frame: After the watermark is played, stay on the last frame.repeat: the watermark loops until the end of the video (default).
        /// </summary>
        public readonly string? RepeatType;
        /// <summary>
        /// The width of the watermark. Support %, px two formats:When the string ends with %, it means that the watermark Width is a percentage of the video width, such as 10% means that the Width is 10% of the video width.When the string ends with px, it means that the watermark Width unit is pixel, such as 100px means that the Width is 100 pixels. The value range is [8, 4096].Default value: 10%.
        /// </summary>
        public readonly string? Width;

        [OutputConstructor]
        private WatermarkTemplateImageTemplate(
            string? height,

            string imageContent,

            string? repeatType,

            string? width)
        {
            Height = height;
            ImageContent = imageContent;
            RepeatType = repeatType;
            Width = width;
        }
    }
}
