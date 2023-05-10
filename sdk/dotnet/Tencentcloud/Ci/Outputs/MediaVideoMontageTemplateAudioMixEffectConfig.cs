// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Outputs
{

    [OutputType]
    public sealed class MediaVideoMontageTemplateAudioMixEffectConfig
    {
        /// <summary>
        /// bgm transition fade-in duration, support floating point numbers.
        /// </summary>
        public readonly string? BgmFadeTime;
        /// <summary>
        /// Enable bgm conversion fade in.
        /// </summary>
        public readonly string? EnableBgmFade;
        /// <summary>
        /// enable fade out.
        /// </summary>
        public readonly string? EnableEndFadeout;
        /// <summary>
        /// enable fade in.
        /// </summary>
        public readonly string? EnableStartFadein;
        /// <summary>
        /// fade out time, greater than 0, support floating point numbers.
        /// </summary>
        public readonly string? EndFadeoutTime;
        /// <summary>
        /// Fade in duration, greater than 0, support floating point numbers.
        /// </summary>
        public readonly string? StartFadeinTime;

        [OutputConstructor]
        private MediaVideoMontageTemplateAudioMixEffectConfig(
            string? bgmFadeTime,

            string? enableBgmFade,

            string? enableEndFadeout,

            string? enableStartFadein,

            string? endFadeoutTime,

            string? startFadeinTime)
        {
            BgmFadeTime = bgmFadeTime;
            EnableBgmFade = enableBgmFade;
            EnableEndFadeout = enableEndFadeout;
            EnableStartFadein = enableStartFadein;
            EndFadeoutTime = endFadeoutTime;
            StartFadeinTime = startFadeinTime;
        }
    }
}