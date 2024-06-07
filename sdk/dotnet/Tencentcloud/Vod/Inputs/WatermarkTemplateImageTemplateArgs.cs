// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Inputs
{

    public sealed class WatermarkTemplateImageTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Watermark height. % and px formats are supported: If the string ends in %, the `Height` of the watermark will be the specified percentage of the video height; for example, `10%` means that `Height` is 10% of the video height;  If the string ends in px, the `Height` of the watermark will be in px; for example, `100px` means that `Height` is 100 px. Valid values: 0 or [8,4096]. Default value: 0 px, which means that `Height` will be proportionally scaled according to the aspect ratio of the original watermark image.
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// The [Base64](https://tools.ietf.org/html/rfc4648) encoded string of a watermark image. Only JPEG, PNG, and GIF images are supported.
        /// </summary>
        [Input("imageContent", required: true)]
        public Input<string> ImageContent { get; set; } = null!;

        /// <summary>
        /// Repeat type of an animated watermark. Valid values: once: no longer appears after watermark playback ends.  repeat_last_frame: stays on the last frame after watermark playback ends.  repeat (default): repeats the playback until the video ends.
        /// </summary>
        [Input("repeatType")]
        public Input<string>? RepeatType { get; set; }

        /// <summary>
        /// Image watermark transparency: 0: completely opaque  100: completely transparent Default value: 0.
        /// </summary>
        [Input("transparency")]
        public Input<int>? Transparency { get; set; }

        /// <summary>
        /// Watermark width. % and px formats are supported: If the string ends in %, the `Width` of the watermark will be the specified percentage of the video width. For example, `10%` means that `Width` is 10% of the video width;  If the string ends in px, the `Width` of the watermark will be in pixels. For example, `100px` means that `Width` is 100 pixels. Value range: [8, 4096]. Default value: 10%.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public WatermarkTemplateImageTemplateArgs()
        {
        }
        public static new WatermarkTemplateImageTemplateArgs Empty => new WatermarkTemplateImageTemplateArgs();
    }
}
