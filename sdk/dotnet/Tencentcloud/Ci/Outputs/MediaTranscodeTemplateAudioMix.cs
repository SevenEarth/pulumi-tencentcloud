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
    public sealed class MediaTranscodeTemplateAudioMix
    {
        /// <summary>
        /// The media address of the audio track that needs to be mixed.
        /// </summary>
        public readonly string AudioSource;
        /// <summary>
        /// Mix Fade Configuration.
        /// </summary>
        public readonly Outputs.MediaTranscodeTemplateAudioMixEffectConfig? EffectConfig;
        /// <summary>
        /// Mixing mode Repeat: background sound loop, Once: The background sound is played once.
        /// </summary>
        public readonly string? MixMode;
        /// <summary>
        /// Whether to replace the original audio of the Input media file with the mixed audio track media.
        /// </summary>
        public readonly string? Replace;

        [OutputConstructor]
        private MediaTranscodeTemplateAudioMix(
            string audioSource,

            Outputs.MediaTranscodeTemplateAudioMixEffectConfig? effectConfig,

            string? mixMode,

            string? replace)
        {
            AudioSource = audioSource;
            EffectConfig = effectConfig;
            MixMode = mixMode;
            Replace = replace;
        }
    }
}
