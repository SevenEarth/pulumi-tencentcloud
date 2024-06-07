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

    public sealed class MediaConcatTemplateConcatTemplateAudioMixEffectConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// bgm transition fade-in duration, support floating point numbers.
        /// </summary>
        [Input("bgmFadeTime")]
        public Input<string>? BgmFadeTime { get; set; }

        /// <summary>
        /// Enable bgm conversion fade in.
        /// </summary>
        [Input("enableBgmFade")]
        public Input<string>? EnableBgmFade { get; set; }

        /// <summary>
        /// enable fade out.
        /// </summary>
        [Input("enableEndFadeout")]
        public Input<string>? EnableEndFadeout { get; set; }

        /// <summary>
        /// enable fade in.
        /// </summary>
        [Input("enableStartFadein")]
        public Input<string>? EnableStartFadein { get; set; }

        /// <summary>
        /// fade out time, greater than 0, support floating point numbers.
        /// </summary>
        [Input("endFadeoutTime")]
        public Input<string>? EndFadeoutTime { get; set; }

        /// <summary>
        /// Fade in duration, greater than 0, support floating point numbers.
        /// </summary>
        [Input("startFadeinTime")]
        public Input<string>? StartFadeinTime { get; set; }

        public MediaConcatTemplateConcatTemplateAudioMixEffectConfigGetArgs()
        {
        }
        public static new MediaConcatTemplateConcatTemplateAudioMixEffectConfigGetArgs Empty => new MediaConcatTemplateConcatTemplateAudioMixEffectConfigGetArgs();
    }
}
