// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaAnimationTemplateVideoGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Animation per second frame number, Priority: AnimateFramesPerSecond &amp;gt; AnimateOnlyKeepKeyFrame &amp;gt; AnimateTimeIntervalOfFrame.
        /// </summary>
        [Input("animateFramesPerSecond")]
        public Input<string>? AnimateFramesPerSecond { get; set; }

        /// <summary>
        /// GIFs are kept only Keyframe, Priority: AnimateFramesPerSecond &amp;gt; AnimateOnlyKeepKeyFrame &amp;gt; AnimateTimeIntervalOfFrame.
        /// </summary>
        [Input("animateOnlyKeepKeyFrame")]
        public Input<string>? AnimateOnlyKeepKeyFrame { get; set; }

        /// <summary>
        /// Animation frame extraction every time, (0, video duration], Animation frame extraction time interval, If TimeInterval.Duration is set, it is less than this value.
        /// </summary>
        [Input("animateTimeIntervalOfFrame")]
        public Input<string>? AnimateTimeIntervalOfFrame { get; set; }

        /// <summary>
        /// Codec format `gif`, `webp`.
        /// </summary>
        [Input("codec", required: true)]
        public Input<string> Codec { get; set; } = null!;

        /// <summary>
        /// Frame rate, value range: (0, 60], Unit: fps.
        /// </summary>
        [Input("fps")]
        public Input<string>? Fps { get; set; }

        /// <summary>
        /// High, value range: [128, 4096], Unit: px, If only Height is set, Width is calculated according to the original ratio of the video, must be even.
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// Set relative quality, [1, 100), webp image quality setting takes effect, gif has no quality parameter.
        /// </summary>
        [Input("quality")]
        public Input<string>? Quality { get; set; }

        /// <summary>
        /// width, value range: [128, 4096], Unit: px, If only Width is set, Height is calculated according to the original ratio of the video, must be even.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public MediaAnimationTemplateVideoGetArgs()
        {
        }
        public static new MediaAnimationTemplateVideoGetArgs Empty => new MediaAnimationTemplateVideoGetArgs();
    }
}
